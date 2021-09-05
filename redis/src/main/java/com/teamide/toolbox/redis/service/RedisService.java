package com.teamide.toolbox.redis.service;

import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;

import java.util.HashMap;
import java.util.Map;

@Service
@Slf4j
public class RedisService {

    private final Map<String, RedisJedis> cache = new HashMap<>();

    public RedisJedis redis(String host, int port, String auth) throws Exception {

        return redis(host, port, auth, null);
    }

    public RedisJedis redis(String host, int port, String auth, Long automaticShutdown) throws Exception {
        if (automaticShutdown == null) {
            automaticShutdown = 60 * 10L; // 默认10分钟自动关闭
        }
        String name = host + ":" + port;
        String key = host + "-" + port + "-auth";
        RedisJedis redis = cache.get(key);
        if (redis == null || !redis.isStarted()) {
            synchronized (redis) {
                redis = cache.get(key);
                if (redis == null || !redis.isStarted()) {
                    if (redis == null) {
                        log.debug("redis [" + name + "] is null,now create redis");
                    } else {
                        log.warn("redis [" + name + "] is closed,now recreate redis");
                    }
                    redis = new RedisJedis(host, port, auth, automaticShutdown);
                    // 随便获取一个key，测试连接是否正常
                    redis.get("test");
                    cache.put(key, redis);
                }
            }
        }
        return redis;
    }

}
