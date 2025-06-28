package com.pg.vkt.kafka.service;

import org.springframework.stereotype.Service;
import lombok.extern.slf4j.Slf4j;

@Service
@Slf4j
public class ComputationService {

    public void performComputation(String message) throws InterruptedException {
        log.info("performing I/O operation");
        Thread.sleep(5000);
        log.info("completed I/O operation");
    }
}
