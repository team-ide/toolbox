package com.teamide.toolbox.start;

import lombok.extern.slf4j.Slf4j;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.EnableAutoConfiguration;
import org.springframework.boot.autoconfigure.SpringBootApplication;

/**
 * @author 朱亮
 * @version 1.0.0
 * @ClassName ZookeeperAutoConfiguration.java
 * @Description TODO
 * @createTime 2021年08月26日 16:59:00
 */
@SpringBootApplication
@EnableAutoConfiguration
@Slf4j
public class ToolboxStartup {

    public static void main(String[] args) {
        log.info("Team IDE Toolbox starting...");
        SpringApplication.run(ToolboxStartup.class, args);
        log.info("Team IDE Toolbox started");
    }
}
