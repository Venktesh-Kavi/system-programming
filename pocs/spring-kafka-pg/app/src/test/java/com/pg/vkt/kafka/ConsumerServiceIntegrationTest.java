package com.pg.vkt.kafka;

import static org.assertj.core.api.Assertions.assertThat;
import static org.junit.jupiter.api.Assertions.assertEquals;
import static org.junit.jupiter.api.Assertions.assertTrue;
import java.time.Duration;
import java.util.Collections;
import java.util.HashMap;
import java.util.Map;
import java.util.concurrent.CountDownLatch;
import java.util.concurrent.TimeUnit;
import java.util.concurrent.atomic.AtomicBoolean;
import java.util.concurrent.atomic.AtomicInteger;
import org.apache.kafka.clients.consumer.Consumer;
import org.apache.kafka.clients.consumer.ConsumerConfig;
import org.apache.kafka.clients.consumer.ConsumerRecord;
import org.apache.kafka.clients.consumer.ConsumerRecords;
import org.apache.kafka.clients.producer.Producer;
import org.apache.kafka.clients.producer.ProducerConfig;
import org.apache.kafka.clients.producer.ProducerRecord;
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

@ExtendWith(SpringExtension.class)
@SpringBootTest
@DirtiesContext
@EmbeddedKafka(topics = {"test-topic", "test-ack-topic", "test-timeout-topic"}, partitions = 1,
        bootstrapServersProperty = "spring.kafka.bootstrap-servers")
public class ConsumerServiceIntegrationTest {

    private static final String TEST_TOPIC = "test-topic";
    private static final String TEST_ACK_TOPIC = "test-ack-topic";
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

        // Produce 10 messages
        for (int i = 0; i < 10; i++) {
            producer.send(new ProducerRecord<>(TEST_TOPIC, "key-" + i, "value-" + i)).get();
        }

        // Poll for records
        ConsumerRecords<String, String> records =
                KafkaTestUtils.getRecords(consumer, Duration.ofSeconds(10));

        // Verify we received all messages
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
     * Tests manual acknowledgment by committing offsets explicitly
     */
    // @Test
    // public void testManualAcknowledgment() throws Exception {
    //     // Create consumer with auto-commit disabled
    //     Map<String, Object> consumerProps = new HashMap<>(
    //             KafkaTestUtils.consumerProps(GROUP_ID + "-manual", "false", embeddedKafkaBroker));
    //     consumerProps.put(ConsumerConfig.AUTO_OFFSET_RESET_CONFIG, "earliest");
    //     consumerProps.put(ConsumerConfig.ENABLE_AUTO_COMMIT_CONFIG, "false");

    //     ConsumerFactory<String, String> cf = new DefaultKafkaConsumerFactory<>(consumerProps,
    //             new StringDeserializer(), new StringDeserializer());

    //     consumer = cf.createConsumer();
    //     consumer.subscribe(Collections.singletonList(TEST_ACK_TOPIC));

    //     // Produce 5 messages
    //     for (int i = 0; i < 5; i++) {
    //         producer.send(new ProducerRecord<>(TEST_ACK_TOPIC, "key-" + i, "value-" + i)).get();
    //     }

    //     // Poll for records
    //     ConsumerRecords<String, String> records =
    //             KafkaTestUtils.getRecords(consumer, Duration.ofSeconds(10));

    //     // Process records and manually commit after processing
    //     List<String> processedMessages = new ArrayList<>();
    //     for (ConsumerRecord<String, String> record : records) {
    //         processedMessages.add(record.value());
    //     }

    //     // Manually commit offsets
    //     consumer.commitSync();

    //     assertEquals(5, processedMessages.size(), "Should have processed 5 messages");

    //     // Close and create a new consumer with the same group id
    //     consumer.close();

    //     // Create a new consumer with the same group id
    //     consumer = cf.createConsumer();
    //     consumer.subscribe(Collections.singletonList(TEST_ACK_TOPIC));

    //     // Poll again - should not receive any records since offsets were committed
    //     ConsumerRecords<String, String> newRecords =
    //             KafkaTestUtils.getRecords(consumer, Duration.ofMillis(500));
    //     assertEquals(0, newRecords.count(), "Should not receive any records after manual commit");
    // }

    /**
     * Tests behavior when max.poll.interval.ms is exceeded
     */
    @Test
    @Timeout(value = 20, unit = TimeUnit.SECONDS)
    public void testMaxPollIntervalTimeout() throws Exception {
        // Create consumer with a very short max poll interval
        Map<String, Object> consumerProps = new HashMap<>(
                KafkaTestUtils.consumerProps(GROUP_ID + "-timeout", "false", embeddedKafkaBroker));
        consumerProps.put(ConsumerConfig.AUTO_OFFSET_RESET_CONFIG, "earliest");
        consumerProps.put(ConsumerConfig.MAX_POLL_INTERVAL_MS_CONFIG, "3000");
        consumerProps.put(ConsumerConfig.HEARTBEAT_INTERVAL_MS_CONFIG, "1000");
        consumerProps.put(ConsumerConfig.SESSION_TIMEOUT_MS_CONFIG, "5000");

        ConsumerFactory<String, String> cf = new DefaultKafkaConsumerFactory<>(consumerProps,
                new StringDeserializer(), new StringDeserializer());

        consumer = cf.createConsumer();
        consumer.subscribe(Collections.singletonList(TEST_TIMEOUT_TOPIC));

        for (int i = 0; i < 10; i++) {
            producer.send(new ProducerRecord<>(TEST_TIMEOUT_TOPIC, "key-" + i, "value-" + i)).get();
        }

        ConsumerRecords<String, String> records = consumer.poll(Duration.ofSeconds(5));
        assertThat(records.count()).isGreaterThan(0);

        Thread.sleep(4000); // Sleep longer than max.poll.interval.ms

        // Create a thread to call poll() which should detect the timeout
        final AtomicBoolean timeoutDetected = new AtomicBoolean(false);
        final CountDownLatch latch = new CountDownLatch(1);

        Thread pollingThread = new Thread(() -> {
            try {
                // This poll should trigger a rebalance due to max.poll.interval.ms being exceeded
                consumer.poll(Duration.ofSeconds(1));
            } catch (Exception e) {
                // The consumer group will rebalance and this consumer might get kicked out
                System.out.println("Exception during poll: " + e.getMessage());
                timeoutDetected.set(true);
            } finally {
                latch.countDown();
            }
        });

        pollingThread.start();
        latch.await(10, TimeUnit.SECONDS);

        // Verify that a rebalance occurred or the consumer was affected
        assertTrue(timeoutDetected.get() || consumer.assignment().isEmpty(),
                "Consumer should detect timeout or lose assignments");
    }

    /**
     * Tests consumer behavior when processing messages with varying processing times
     */
    @Test
    public void testConsumerWithVaryingProcessingTimes() throws Exception {
        // Create consumer with specific settings
        Map<String, Object> consumerProps = new HashMap<>(
                KafkaTestUtils.consumerProps(GROUP_ID + "-varying", "false", embeddedKafkaBroker));
        consumerProps.put(ConsumerConfig.AUTO_OFFSET_RESET_CONFIG, "earliest");
        consumerProps.put(ConsumerConfig.ENABLE_AUTO_COMMIT_CONFIG, "false");
        consumerProps.put(ConsumerConfig.MAX_POLL_RECORDS_CONFIG, "5"); // Process max 5 records per
                                                                        // poll

        ConsumerFactory<String, String> cf = new DefaultKafkaConsumerFactory<>(consumerProps,
                new StringDeserializer(), new StringDeserializer());

        consumer = cf.createConsumer();
        consumer.subscribe(Collections.singletonList(TEST_TOPIC));

        // Produce 20 messages
        for (int i = 0; i < 20; i++) {
            producer.send(new ProducerRecord<>(TEST_TOPIC, "batch-key-" + i, "batch-value-" + i))
                    .get();
        }

        final AtomicInteger processedCount = new AtomicInteger(0);
        final int expectedMessages = 20;

        // Process messages in batches with manual commits
        while (processedCount.get() < expectedMessages) {
            ConsumerRecords<String, String> records = consumer.poll(Duration.ofSeconds(5));

            if (records.isEmpty()) {
                continue;
            }

            // Process each record with varying processing times
            for (ConsumerRecord<String, String> record : records) {
                // Simulate varying processing times
                int processingTime = (processedCount.get() % 3 == 0) ? 500 : 100; // Longer every
                                                                                  // 3rd message
                Thread.sleep(processingTime);

                System.out.println("Processed: " + record.value() + ", processing time: "
                        + processingTime + "ms");
                processedCount.incrementAndGet();
            }

            // Commit offsets after processing the batch
            consumer.commitSync();
            System.out.println("Committed offsets after processing batch. Total processed: "
                    + processedCount.get());

            // Break if we've processed all expected messages
            if (processedCount.get() >= expectedMessages) {
                break;
            }
        }

        assertEquals(expectedMessages, processedCount.get(), "Should have processed all messages");
    }
}
