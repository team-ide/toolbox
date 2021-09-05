package com.teamide.toolbox.redis;

import com.teamide.toolbox.redis.worker.RedisWorker;
import com.teamide.toolbox.worker.ToolboxWorkerCache;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.annotation.Configuration;

/**
 * @author 朱亮
 * @version 1.0.0
 * @ClassName RedisAutoConfiguration.java
 * @Description TODO
 * @createTime 2021年08月26日 17:00:00
 */
@Configuration
@ComponentScan
public class RedisAutoConfiguration {

    @Autowired
    private ToolboxWorkerCache cache;

    @Autowired
    private RedisWorker worker;

    @Bean
    public RedisWorker getRedisWorker() {
        cache.add(worker);
        return worker;
    }
}
