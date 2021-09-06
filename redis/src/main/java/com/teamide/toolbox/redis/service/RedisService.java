package com.teamide.toolbox.redis.service;

import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;
import redis.clients.jedis.HostAndPort;

import java.util.*;

@Service
@Slf4j
public class RedisService {

    private final Map<String, RedisDo> cache = new HashMap<>();

    public RedisDo redis(String address, String auth, boolean cluster) throws Exception {

        return redis(address, auth, cluster, null);
    }

    public RedisDo redis(String address, String auth, boolean cluster, Long automaticShutdown) throws Exception {
        if (automaticShutdown == null) {
            automaticShutdown = 60 * 10L; // 默认10分钟自动关闭
        }

        String name = address;
        String key = name + "-" + auth + "-" + cluster;
        RedisDo redis = cache.get(key);
        if (redis == null || !redis.isStarted()) {
            synchronized (cache) {
                redis = cache.get(key);
                if (redis == null || !redis.isStarted()) {
                    if (redis == null) {
                        log.debug("redis [" + name + "] is null,now create redis");
                    } else {
                        log.warn("redis [" + name + "] is closed,now recreate redis");
                    }
                    if (cluster) {
                        redis = new RedisCluster(address, auth, automaticShutdown);
                    } else {
                        redis = new RedisJedis(address, auth, automaticShutdown);
                    }
                    // 随便获取一个key，测试连接是否正常
                    redis.get("test");
                    cache.put(key, redis);
                }
            }
        }
        return redis;
    }

}
