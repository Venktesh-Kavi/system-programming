package com.pg.vkt.kafka;

import static org.assertj.core.api.Assertions.assertThat;
import static org.junit.jupiter.api.Assertions.assertEquals;
import java.time.Duration;
import java.util.Collection;
import java.util.Collections;
import java.util.HashMap;
import java.util.HashSet;
import java.util.Map;
import java.util.Set;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.TimeUnit;
import org.apache.kafka.clients.admin.NewTopic;
import org.apache.kafka.clients.consumer.Consumer;
import org.apache.kafka.clients.consumer.ConsumerConfig;
import org.apache.kafka.clients.consumer.ConsumerRebalanceListener;
import org.apache.kafka.clients.consumer.ConsumerRecord;
import org.apache.kafka.clients.consumer.ConsumerRecords;
import org.apache.kafka.clients.producer.Producer;
import org.apache.kafka.clients.producer.ProducerConfig;
import org.apache.kafka.clients.producer.ProducerRecord;
import org.apache.kafka.common.TopicPartition;
import org.apache.kafka.common.serialization.StringDeserializer;
import org.apache.kafka.common.serialization.StringSerializer;
import org.junit.jupiter.api.AfterEach;
import org.junit.jupiter.api.BeforeEach;
import org.junit.jupiter.api.Test;
import org.junit.jupiter.api.Timeout;
import org.junit.jupiter.api.extension.ExtendWith;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;
import org.springframework.kafka.core.ConsumerFactory;
import org.springframework.kafka.core.DefaultKafkaConsumerFactory;
import org.springframework.kafka.core.DefaultKafkaProducerFactory;
import org.springframework.kafka.core.ProducerFactory;
import org.springframework.kafka.test.EmbeddedKafkaBroker;
import org.springframework.kafka.test.context.EmbeddedKafka;
import org.springframework.kafka.test.utils.KafkaTestUtils;
import org.springframework.test.annotation.DirtiesContext;
import org.springframework.test.context.junit.jupiter.SpringExtension;
import jakarta.inject.Named;

@ExtendWith(SpringExtension.class)
@SpringBootTest
@DirtiesContext
@EmbeddedKafka(topics = {"test-topic", "test-timeout-topic"}, partitions = 1,
        bootstrapServersProperty = "spring.kafka.bootstrap-servers")
public class ConsumerServiceIntegrationTest {

    private static final String TEST_TOPIC = "test-topic";
    private static final String TEST_TIMEOUT_TOPIC = "test-timeout-topic";
    private static final String GROUP_ID = "test-group";

    @Autowired
    private EmbeddedKafkaBroker embeddedKafkaBroker;

    private Consumer<String, String> consumer;
    private Producer<String, String> producer;

    @BeforeEach
    public void setup() {
        // Set up the producer
        Map<String, Object> producerProps = KafkaTestUtils.producerProps(embeddedKafkaBroker);
        producerProps.put(ProducerConfig.KEY_SERIALIZER_CLASS_CONFIG, StringSerializer.class);
        producerProps.put(ProducerConfig.VALUE_SERIALIZER_CLASS_CONFIG, StringSerializer.class);
        ProducerFactory<String, String> pf = new DefaultKafkaProducerFactory<>(producerProps);
        producer = pf.createProducer();
    }

    @AfterEach
    public void tearDown() {
        if (consumer != null) {
            consumer.close();
        }
        if (producer != null) {
            producer.close();
        }
    }

    /**
     * Tests basic consumer functionality with auto-commit enabled
     */
    @Test
    @Named("test basic consumer with auto offset commit mechanism")
    public void testBasicConsumerWithAutoCommit() throws Exception {
        // Create consumer with auto-commit enabled
        Map<String, Object> consumerProps =
                new HashMap<>(KafkaTestUtils.consumerProps(GROUP_ID, "true", embeddedKafkaBroker));
        consumerProps.put(ConsumerConfig.AUTO_OFFSET_RESET_CONFIG, "earliest");
        consumerProps.put(ConsumerConfig.ENABLE_AUTO_COMMIT_CONFIG, "true");
        consumerProps.put(ConsumerConfig.AUTO_COMMIT_INTERVAL_MS_CONFIG, "100");

        ConsumerFactory<String, String> cf = new DefaultKafkaConsumerFactory<>(consumerProps,
                new StringDeserializer(), new StringDeserializer());

        consumer = cf.createConsumer();
        consumer.subscribe(Collections.singletonList(TEST_TOPIC));

        for (int i = 0; i < 10; i++) {
            producer.send(new ProducerRecord<>(TEST_TOPIC, "key-" + i, "value-" + i)).get();
        }

        ConsumerRecords<String, String> records =
                KafkaTestUtils.getRecords(consumer, Duration.ofSeconds(10));

        int count = 0;
        for (ConsumerRecord<String, String> record : records) {
            System.out.println("Received: " + record.value());
            count++;
        }

        assertEquals(10, count, "Should have received 10 messages");

        // Close and create a new consumer with the same group id to verify commits
        consumer.close();

        // Create a new consumer with the same group id
        consumer = cf.createConsumer();
        consumer.subscribe(Collections.singletonList(TEST_TOPIC));

        // Poll again - should not receive any records since offsets were committed
        ConsumerRecords<String, String> newRecords =
                KafkaTestUtils.getRecords(consumer, Duration.ofMillis(500));
        assertEquals(0, newRecords.count(), "Should not receive any records after commit");
    }

    /**
     * Tests behavior when max.poll.interval.ms is exceeded
     */
    @Test
    @Timeout(value = 30, unit = TimeUnit.SECONDS)
    @Named("test consumer behaviour on exceeding max poll interval ms")
    public void testMaxPollIntervalTimeout() throws Exception {
        String topic = TEST_TIMEOUT_TOPIC + "-multi";
        embeddedKafkaBroker.addTopics(new NewTopic(topic, 2, (short) 1));

        // Make sure topic is discoverable
        producer.partitionsFor(topic);

        Map<String, Object> consumer1Props = new HashMap<>(
                KafkaTestUtils.consumerProps(GROUP_ID + "-timeout", "false", embeddedKafkaBroker));
        consumer1Props.put(ConsumerConfig.CLIENT_ID_CONFIG, "vkt-test-timeout-client1");
        consumer1Props.put(ConsumerConfig.AUTO_OFFSET_RESET_CONFIG, "earliest");
        consumer1Props.put(ConsumerConfig.MAX_POLL_INTERVAL_MS_CONFIG, "5000");
        consumer1Props.put(ConsumerConfig.HEARTBEAT_INTERVAL_MS_CONFIG, "1000");
        consumer1Props.put(ConsumerConfig.SESSION_TIMEOUT_MS_CONFIG, "6000");

        Map<String, Object> consumer2Props = new HashMap<>(
                KafkaTestUtils.consumerProps(GROUP_ID + "-timeout", "false", embeddedKafkaBroker));
        consumer2Props.put(ConsumerConfig.CLIENT_ID_CONFIG, "vkt-test-timeout-client2");
        consumer2Props.put(ConsumerConfig.AUTO_OFFSET_RESET_CONFIG, "earliest");
        consumer2Props.put(ConsumerConfig.MAX_POLL_INTERVAL_MS_CONFIG, "5000");
        consumer2Props.put(ConsumerConfig.HEARTBEAT_INTERVAL_MS_CONFIG, "1000");
        consumer2Props.put(ConsumerConfig.SESSION_TIMEOUT_MS_CONFIG, "6000");

        ConsumerFactory<String, String> cf1 = new DefaultKafkaConsumerFactory<>(consumer1Props,
                new StringDeserializer(), new StringDeserializer());
        ConsumerFactory<String, String> cf2 = new DefaultKafkaConsumerFactory<>(consumer2Props,
                new StringDeserializer(), new StringDeserializer());
        Consumer<String, String> c1 = cf1.createConsumer();
        Consumer<String, String> c2 = cf2.createConsumer();

        try {
            CountDownLatch c1Assigned = new CountDownLatch(1);
            CountDownLatch c2Assigned = new CountDownLatch(1);

            // Acquiring couple of latches till the consumer partitions are assigned.
            c1.subscribe(Collections.singletonList(topic), new ConsumerRebalanceListener() {
                public void onPartitionsRevoked(Collection<TopicPartition> partitions) {}

                public void onPartitionsAssigned(Collection<TopicPartition> partitions) {
                    if (!partitions.isEmpty())
                        c1Assigned.countDown();
                }
            });
            c2.subscribe(Collections.singletonList(topic), new ConsumerRebalanceListener() {
                public void onPartitionsRevoked(Collection<TopicPartition> partitions) {}

                public void onPartitionsAssigned(Collection<TopicPartition> partitions) {
                    if (!partitions.isEmpty())
                        c2Assigned.countDown();
                }
            });

            for (int i = 0; i < 10; i++) {
                producer.send(new ProducerRecord<>(topic, 0, "key-" + i, "value-" + i)).get();
                producer.send(new ProducerRecord<>(topic, 1, "key-" + i, "value-" + i)).get();
            }

            // Force a rebalance & assignments by repeated poll, and wait for listener. Not
            // commiting explicitly.
            while (c1Assigned.getCount() > 0 || c2Assigned.getCount() > 0) {
                c1.poll(Duration.ofMillis(100));
                c2.poll(Duration.ofMillis(100));
            }

            // Just in case, give the assignment a moment to stabilize
            Thread.sleep(200);

            Set<TopicPartition> initialC1Assignment = new HashSet<>(c1.assignment());
            Set<TopicPartition> initialC2Assignment = new HashSet<>(c2.assignment());
            System.out.println("Consumer 1 initial assignment: " + initialC1Assignment);
            System.out.println("Consumer 2 initial assignment: " + initialC2Assignment);
            assertThat(initialC1Assignment.size() + initialC2Assignment.size()).isEqualTo(2);
            assertThat(initialC1Assignment).doesNotContainAnyElementsOf(initialC2Assignment);

            // c1: drain records (establish poll loop)
            c1.poll(Duration.ofSeconds(1));

            // Sleep longer than max.poll.interval.ms for c1 --> to simulate consumer stuck
            System.out.println("Sleeping to trigger max.poll.interval.ms for Consumer 1...");
            Thread.sleep(7000);

            // c2: poll after that to trigger rebalance & process reassignment
            System.out.println("Polling with Consumer 2 to trigger rebalance...");
            ConsumerRecords<String, String> cr = c2.poll(Duration.ofSeconds(10));
            System.out.println("Consumer 2 records after rebalance poll: " + cr.count());

            // Wait until c2 has both partitions assigned (must be explicit; rebalance is async)
            CountDownLatch c2Reassigned = new CountDownLatch(1);
            // If partitions already both present, latch is satisfied
            if (c2.assignment().size() == 2) {
                c2Reassigned.countDown();
            } else {
                // If not, set up a listener for the rebalance
                c2.assign(Collections.emptyList());
                c2.unsubscribe();
                c2.subscribe(Collections.singletonList(topic), new ConsumerRebalanceListener() {
                    @Override
                    public void onPartitionsRevoked(Collection<TopicPartition> partitions) {}

                    @Override
                    public void onPartitionsAssigned(Collection<TopicPartition> partitions) {
                        System.out.println("Rebalance: c2 assigned = " + partitions);
                        if (partitions.size() == 2)
                            c2Reassigned.countDown();
                    }
                });
                // Poll until we're reassigned both partitions
                while (c2Reassigned.getCount() > 0) {
                    c2.poll(Duration.ofMillis(100));
                }
            }

            Set<TopicPartition> postTimeoutAssignment = new HashSet<>(c2.assignment());
            System.out.println("Consumer 2 assignment after c1 timeout: " + postTimeoutAssignment);

            // c1's old partitions should be assigned to c2
            assertThat(postTimeoutAssignment).containsAll(initialC1Assignment);

        } finally {
            c1.close();
            c2.close();
        }
    }

}
