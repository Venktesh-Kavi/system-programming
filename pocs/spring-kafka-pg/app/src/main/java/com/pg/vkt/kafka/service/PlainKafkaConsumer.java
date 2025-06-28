package com.pg.vkt.kafka.service;

import org.springframework.kafka.annotation.KafkaListener;
import org.springframework.stereotype.Service;
import lombok.RequiredArgsConstructor;

@Service
@RequiredArgsConstructor
public class PlainKafkaConsumer {
    private final ComputationService computationService;

    @KafkaListener(topics = "test", groupId = "plain-kafka-consumer")
    public void consume(String message) throws InterruptedException {
        System.out.println("Received message: " + message);
        computationService.performComputation(message);
        throw new IllegalArgumentException("test");
    }
}
