package com.pg.vkt.kafka.service;

import java.time.LocalDateTime;
import org.apache.kafka.clients.consumer.ConsumerRecord;
import org.springframework.stereotype.Service;
import com.pg.vkt.kafka.entity.FailedEvent;
import lombok.RequiredArgsConstructor;

@Service
@RequiredArgsConstructor
public class FailedEventService {

    public void handleFailedRecord(ConsumerRecord<String, String> record, Exception exception,
            int retryCount) {
        FailedEvent failedEvent = FailedEvent.builder().topic(record.topic())
                .partition(record.partition()).offset(record.offset()).key(record.key())
                .value(record.value()).exception(exception.getMessage())
                .failedAt(LocalDateTime.now()).retryCount(retryCount).build();

        // Here you can add additional error handling:
        // 1. Send notifications
        // 2. Log to monitoring system
        // 3. Trigger alerts
        // 4. etc.
    }
}
