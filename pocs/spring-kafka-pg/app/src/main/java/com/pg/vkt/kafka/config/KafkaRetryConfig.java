package com.pg.vkt.kafka.config;

import org.springframework.boot.context.properties.ConfigurationProperties;
import org.springframework.context.annotation.Configuration;
import lombok.Data;

@Data
@Configuration
@ConfigurationProperties(prefix = "spring.kafka.consumer")
public class KafkaRetryConfig {
    private int maxPollIntervalMs = 300000; // default 5 minutes
    private int maxPollRecords = 500;       // default batch size
    
    public long calculateBackoffInterval() {
        // Calculate a safe backoff interval that won't exceed max poll interval
        // Consider batch size and leave room for processing
        int maxRetries = 3;
        // Reserve 20% of max poll interval for actual processing
        long safeTimeForRetries = (long) (maxPollIntervalMs * 0.8);
        // Divide by (batch size * max retries) to get per-record retry interval
        return safeTimeForRetries / (maxPollRecords * maxRetries);
    }
    
    public int getMaxRetries() {
        return 3; // Can be made configurable if needed
    }
}
