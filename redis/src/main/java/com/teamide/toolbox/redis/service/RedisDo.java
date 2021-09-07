package com.teamide.toolbox.redis.service;

import lombok.extern.slf4j.Slf4j;
import org.apache.commons.lang3.StringUtils;
import redis.clients.jedis.HostAndPort;
import redis.clients.jedis.Jedis;
import redis.clients.jedis.JedisPool;
import redis.clients.jedis.JedisPoolConfig;

import java.util.Set;
import java.util.Timer;
import java.util.TimerTask;


public interface RedisDo {

    public boolean isStarted() throws Exception;

    public String set(String key, String value) throws Exception;

    public Set<String> keys(String pattern, int size) throws Exception;

    public String get(String key) throws Exception;

    public void delete(String key) throws Exception;
}
