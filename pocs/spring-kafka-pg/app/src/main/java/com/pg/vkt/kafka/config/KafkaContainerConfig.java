package com.pg.vkt.kafka.config;

import org.apache.kafka.clients.consumer.ConsumerRecord;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.Configuration;
import org.springframework.kafka.config.ConcurrentKafkaListenerContainerFactory;
import org.springframework.kafka.core.ConsumerFactory;
import org.springframework.kafka.listener.ConsumerRecordRecoverer;
import org.springframework.kafka.listener.DefaultErrorHandler;
import org.springframework.util.backoff.FixedBackOff;
import com.pg.vkt.kafka.service.FailedEventService;

@Configuration
public class KafkaContainerConfig {

    @Bean
    public ConcurrentKafkaListenerContainerFactory<String, String> kafkaListenerContainerFactory(
            ConsumerFactory<String, String> consumerFactory, DefaultErrorHandler errorHandler) {
        ConcurrentKafkaListenerContainerFactory<String, String> factory =
                new ConcurrentKafkaListenerContainerFactory<>();
        factory.setConsumerFactory(consumerFactory);
        factory.setCommonErrorHandler(errorHandler);
        return factory;
    }

    @Bean
    public DefaultErrorHandler errorHandler(FailedEventService failedEventService) {
        FixedBackOff fixedBackOff = new FixedBackOff(5000L, 3L);

        ConsumerRecordRecoverer recoverer = (consumerRecord, exception) -> {
            ConsumerRecord<String, String> record = (ConsumerRecord<String, String>) consumerRecord;
            failedEventService.handleFailedRecord(record, exception, 3);
        };

        DefaultErrorHandler errorHandler = new DefaultErrorHandler(recoverer, fixedBackOff);

        errorHandler.setRetryListeners((record, ex, deliveryAttempt) -> {
            System.out.printf(
                    "Failed to process record. Topic: %s, Partition: %d, Offset: %d, Attempt: %d, Error: %s%n",
                    record.topic(), record.partition(), record.offset(), deliveryAttempt,
                    ex.getMessage());
        });

        errorHandler.addNotRetryableExceptions(NullPointerException.class);

        return errorHandler;
    }
}
