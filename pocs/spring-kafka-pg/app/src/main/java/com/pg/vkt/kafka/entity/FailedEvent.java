package com.pg.vkt.kafka.entity;

import java.time.LocalDateTime;
import lombok.AllArgsConstructor;
import lombok.Builder;
import lombok.Data;
import lombok.NoArgsConstructor;

@Data
@Builder
@NoArgsConstructor
@AllArgsConstructor
public class FailedEvent {
    private String topic;
    private Integer partition;
    private Long offset;
    private String key;
    private String value;
    private String exception;
    private LocalDateTime failedAt;
    private Integer retryCount;
}
