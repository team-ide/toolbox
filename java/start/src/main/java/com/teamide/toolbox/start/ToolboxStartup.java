package com.teamide.toolbox.start;

import lombok.extern.slf4j.Slf4j;
import org.springframework.boot.SpringApplication;
import org.springframework.boot.autoconfigure.EnableAutoConfiguration;
import org.springframework.boot.autoconfigure.SpringBootApplication;
import org.springframework.boot.autoconfigure.mongo.MongoAutoConfiguration;

/**
 * 工具箱的启动
 *
 * @author 朱亮
 * @date 2021/09/08
 */
@SpringBootApplication
@EnableAutoConfiguration(exclude = {MongoAutoConfiguration.class})
@Slf4j
public class ToolboxStartup {

    public static void main(String[] args) {
        log.info("Team IDE Toolbox starting...");
        SpringApplication.run(ToolboxStartup.class, args);
        log.info("Team IDE Toolbox started");
    }
}
