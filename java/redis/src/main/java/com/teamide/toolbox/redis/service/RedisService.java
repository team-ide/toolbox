package com.teamide.toolbox.redis.service;

import com.teamide.toolbox.worker.ToolboxAutomaticShutdownTimer;
import lombok.extern.slf4j.Slf4j;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.*;

/**
 * Redis服务
 *
 * @author 朱亮
 * @date 2021/09/08
 */
@Service
@Slf4j
public class RedisService {

    @Autowired
    ToolboxAutomaticShutdownTimer toolboxAutomaticShutdownTimer;

    private final Map<String, RedisDo> cache = new HashMap<>();

    public RedisDo redis(String address, String auth) throws Exception {

        return redis(address, auth, null);
    }

    public RedisDo redis(String address, String auth, Long automaticShutdown) throws Exception {
        if (automaticShutdown == null) {
            // 默认10分钟自动关闭
            automaticShutdown = 60 * 10L;
        }

        String key = address + "-" + auth;
        RedisDo redis = cache.get(key);
        if (redis == null || !redis.started()) {
            synchronized (cache) {
                redis = cache.get(key);
                if (redis == null || !redis.started()) {
                    if (redis == null) {
                        log.debug("redis [" + address + "] is null,now create redis");
                    } else {
                        log.warn("redis [" + address + "] is closed,now recreate redis");
                    }
                    boolean cluster = address.contains(";");
                    if (cluster) {
                        redis = new RedisCluster(address, auth, automaticShutdown);
                    } else {
                        redis = new RedisJedis(address, auth, automaticShutdown);
                    }
                    // 随便获取一个key，测试连接是否正常
                    redis.get("test");
                    cache.put(key, redis);
                    toolboxAutomaticShutdownTimer.add(redis);
                }
            }
        }
        return redis;
    }

}
