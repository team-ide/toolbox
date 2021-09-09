package com.teamide.toolbox.redis.service;

import lombok.extern.slf4j.Slf4j;
import org.apache.commons.lang3.StringUtils;
import redis.clients.jedis.HostAndPort;
import redis.clients.jedis.Jedis;
import redis.clients.jedis.JedisPool;
import redis.clients.jedis.JedisPoolConfig;
import redis.clients.jedis.exceptions.JedisConnectionException;

import java.util.*;


/**
 * Redis单机操作
 *
 * @author 朱亮
 * @date 2021/09/08
 */
@Slf4j
public class RedisJedis implements RedisDo {

    /**
     * 名称，用于日志区分
     */
    private final String name;

    /**
     * 连接地址
     */
    private final String address;

    /**
     * 鉴权密钥
     */
    private final String auth;

    /**
     * 自动关闭时长 单位秒
     */
    private final long automaticShutdown;

    /**
     * 最后使用时间戳
     */
    private long lastUseTimestamp = System.currentTimeMillis();

    /**
     * jedis 连接池 客户端管理工具
     */
    private final JedisPool pool;

    /**
     * 已启动标识
     */
    private boolean started = true;

    /**
     * 用于块级锁
     */
    private final Object lock = new Object();

    /**
     * 可用连接实例的最大数目，默认值为8；
     * 如果赋值为-1，则表示不限制；如果pool已经分配了maxActive个jedis实例，则此时pool的状态为exhausted(耗尽)。
     */
    private static final int MAX_TOTAL = 10;

    /**
     * 控制一个pool最多有多少个状态为idle(空闲的)的jedis实例，默认值也是8。
     */
    private static final int MAX_IDLE = 5;

    /**
     * 在borrow一个jedis实例时，是否提前进行validate操作；如果为true，则得到的jedis实例均是可用的；
     */
    private static final boolean TEST_ON_BORROW = true;

    private static final int TIMEOUT = 10000;

    public RedisJedis(String address, String auth, long automaticShutdown) {
        this.address = address;
        if (StringUtils.isEmpty(auth)) {
            auth = null;
        }
        this.auth = auth;
        this.name = address;
        this.automaticShutdown = automaticShutdown;

        JedisPoolConfig config = new JedisPoolConfig();
        config.setMaxIdle(MAX_IDLE);
        config.setMaxTotal(MAX_TOTAL);
        config.setTestOnBorrow(TEST_ON_BORROW);

        log.info("redis [" + this.name + "] create pool start");
        HostAndPort hostAndPort = HostAndPort.from(address);
        this.pool = new JedisPool(config, hostAndPort.getHost(), hostAndPort.getPort(), TIMEOUT, auth);

        log.info("redis [" + this.name + "] create pool end");

    }


    private void userStart() {
        if (!this.started) {
            return;
        }
        lastUseTimestamp = System.currentTimeMillis();
    }

    private void userEnd(Exception exception) {
        lastUseTimestamp = System.currentTimeMillis();
        if (exception != null) {
            boolean shouldClose = exception instanceof JedisConnectionException;

            if (shouldClose) {
                stop();
            }
        }
    }

    /**
     * 是否启动
     *
     * @return 是否启动
     */
    @Override
    public boolean started() {
        return this.started;
    }

    @Override
    public void stop() {
        synchronized (lock) {
            if (!this.started) {
                return;
            }
            log.debug("redis [" + name + "] curator close ");
            this.started = false;
            this.pool.close();
        }
    }

    /**
     * 设值
     *
     * @param key   key
     * @param value value
     * @throws Exception 异常
     */
    @Override
    public String set(String key, String value) throws Exception {
        Exception err = null;
        Jedis jedis = null;
        userStart();
        try {
            log.debug("redis [" + name + "] set key [" + key + "] value [" + value + "] start ");
            jedis = pool.getResource();
            String code = jedis.set(key, value);
            log.debug("redis [" + name + "] set key [" + key + "] value [" + value + "] code [" + code + "] end ");
            return code;
        } catch (Exception e) {
            log.error("redis [" + name + "] set key [" + key + "] value [" + value + "] error {} ", e);
            err = e;
            throw e;
        } finally {
            if (jedis != null) {
                jedis.close();
            }
            userEnd(err);
        }
    }

    /**
     * 查keys
     *
     * @param pattern pattern
     * @throws Exception 异常
     */
    @Override
    public Set<String> keys(String pattern, int size) throws Exception {
        Exception err = null;
        Jedis jedis = null;
        userStart();
        Set<String> keys = new HashSet<>();
        try {
            log.debug("redis [" + name + "] keys pattern [" + pattern + "] start ");
            jedis = pool.getResource();
            Set<String> ks = jedis.keys(pattern);
            if (size <= 0 || ks.size() <= size) {
                keys.addAll(ks);
            } else {
                int index = 0;
                for (String k : ks) {
                    keys.add(k);
                    index++;
                    if (index > size) {
                        break;
                    }
                }
            }
            log.debug("redis [" + name + "] keys pattern [" + pattern + "] end ");
            return keys;
        } catch (Exception e) {
            log.error("redis [" + name + "] keys pattern [" + pattern + "] error {} ", e);
            err = e;
            throw e;
        } finally {
            if (jedis != null) {
                jedis.close();
            }
            userEnd(err);
        }
    }

    /**
     * 取值
     *
     * @param key key
     * @throws Exception 异常
     */
    @Override
    public String get(String key) throws Exception {
        Exception err = null;
        Jedis jedis = null;
        userStart();
        try {
            log.debug("redis [" + name + "] get key [" + key + "] start ");
            jedis = pool.getResource();
            String value = jedis.get(key);
            log.debug("redis [" + name + "] get key [" + key + "] end ");
            return value;
        } catch (Exception e) {
            log.error("redis [" + name + "] get key [" + key + "] error {} ", e);
            err = e;
            throw e;
        } finally {
            if (jedis != null) {
                jedis.close();
            }
            userEnd(err);
        }
    }

    /**
     * 删除
     *
     * @param key key
     * @throws Exception 异常
     */
    @Override
    public void delete(String key) throws Exception {
        Exception err = null;
        Jedis jedis = null;
        userStart();
        try {
            log.debug("redis [" + name + "] delete key [" + key + "] start ");
            jedis = pool.getResource();
            jedis.del(key);
            log.debug("redis [" + name + "] delete key [" + key + "] end ");
        } catch (Exception e) {
            log.error("redis [" + name + "] delete key [" + key + "] error {} ", e);
            err = e;
            throw e;
        } finally {
            if (jedis != null) {
                jedis.close();
            }
            userEnd(err);
        }
    }

    public String getName() {
        return name;
    }

    public String getAddress() {
        return address;
    }

    public String getAuth() {
        return auth;
    }

    @Override
    public long automaticShutdown() {
        return this.automaticShutdown;
    }

    @Override
    public long lastUseTimestamp() {
        return this.lastUseTimestamp;
    }
}
