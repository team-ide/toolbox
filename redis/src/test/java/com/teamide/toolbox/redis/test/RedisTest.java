package com.teamide.toolbox.redis.test;

import com.teamide.toolbox.redis.RedisAutoConfiguration;
import com.teamide.toolbox.redis.service.RedisJedis;
import com.teamide.toolbox.redis.service.RedisService;
import com.teamide.toolbox.worker.ToolboxWorkerCache;
import com.teamide.toolbox.worker.WorkerAutoConfiguration;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

/**
 * @description: TODO 类描述
 * @author: 朱亮
 * @date: 2021/8/27 17:04
 **/
@SpringBootTest(classes = {WorkerAutoConfiguration.class, RedisAutoConfiguration.class})
public class RedisTest {

    static final String host = "proxy.teamide.com";

    static final int port = 6379;

    static final String auth = null;

    @Autowired
    private RedisService redisService;

    @Autowired
    private ToolboxWorkerCache workerCache;


    @Test
    public void testRedis() throws Exception {
        System.out.println(workerCache.getWorkers().size());
        new Thread(() -> {
            try {
                RedisJedis redis = redisService.redis(host, port, auth);
                String code = redis.set("test-1", "111111");
                System.out.println("code:" + code);
                Thread.sleep(1000 * 3);

                redis = redisService.redis(host, port, auth);
                String value = redis.get("test-1");
                Thread.sleep(1000 * 6);

                redis = redisService.redis(host, port, auth);
                redis.delete("test-1");

                Thread.sleep(1000 * 6);
                redis = redisService.redis(host, port, auth);
                value = redis.get("test-1");
                System.out.println("value:" + value);

            } catch (Exception e) {
                e.printStackTrace();
            }
        }).start();
        Thread.sleep(1000 * 60 * 10);
    }

}
