package com.teamide.toolbox.redis.service;

import lombok.extern.slf4j.Slf4j;
import org.apache.commons.lang3.StringUtils;
import redis.clients.jedis.Jedis;
import redis.clients.jedis.JedisPool;
import redis.clients.jedis.JedisPoolConfig;

import java.util.List;
import java.util.Set;
import java.util.Timer;
import java.util.TimerTask;


@Slf4j
public class RedisJedis {

    // Redis name
    private final String name;

    // Redis host
    private final String host;

    // Redis port
    private final int port;

    // Redis auth
    private final String auth;

    // 自动关闭时长 单位秒
    private final long automaticShutdown;

    // 最后使用时间戳
    private long lastUseTime = System.currentTimeMillis();

    // jedis 连接池 客户端管理工具
    private final JedisPool pool;

    // 已启动标识
    private boolean started = true;

    // 用于块级锁
    private final Object lock = new Object();

    //
    private Timer timer;

    //可用连接实例的最大数目，默认值为8；
    //如果赋值为-1，则表示不限制；如果pool已经分配了maxActive个jedis实例，则此时pool的状态为exhausted(耗尽)。
    private static int MAX_TOTAL = 10;

    //控制一个pool最多有多少个状态为idle(空闲的)的jedis实例，默认值也是8。
    private static int MAX_IDLE = 5;

    //在borrow一个jedis实例时，是否提前进行validate操作；如果为true，则得到的jedis实例均是可用的；
    private static boolean TEST_ON_BORROW = true;

    private static int TIMEOUT = 10000;

    public RedisJedis(String host, int port, String auth, long automaticShutdown) {
        this.host = host;
        this.port = port;
        if (StringUtils.isEmpty(auth)) {
            auth = null;
        }
        this.auth = auth;
        this.name = host + ":" + port;
        this.automaticShutdown = automaticShutdown;

        JedisPoolConfig config = new JedisPoolConfig();
        config.setMaxIdle(MAX_IDLE);
        config.setMaxTotal(MAX_TOTAL);
        config.setTestOnBorrow(TEST_ON_BORROW);

        log.info("redis [" + this.name + "] create pool start");

        this.pool = new JedisPool(config, host, port, TIMEOUT, auth);

        log.info("redis [" + this.name + "] create pool end");

        if (this.automaticShutdown > 0) {
            timer = new Timer();
            startTimerTask();
        }
    }

    private void startTimerTask() {
//        log.debug("redis [" + name + "] timer add task ");
        timer.schedule(new TimerTask() {
            @Override
            public void run() {
                checkAutomaticShutdown();
            }
        }, 1000);
    }

    private void checkAutomaticShutdown() {
//        log.debug("redis [" + name + "] timer task checkAutomaticShutdown ");
        long now = System.currentTimeMillis();
        long timeout = now - lastUseTime;
        // 超时 关闭
        if (timeout > this.automaticShutdown * 1000) {
            log.warn("redis [" + name + "] timeout need to stop");
            stop();
            timer.cancel();
        } else {
            startTimerTask();
        }
    }

    private void userStart() {
        if (!this.started) {
            return;
        }
        lastUseTime = System.currentTimeMillis();
    }

    private void userEnd(Exception exception) {
        lastUseTime = System.currentTimeMillis();
        if (exception != null) {
            boolean shouldClose = false;

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
    public boolean isStarted() {
        return this.started;
    }

    private void stop() {
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
            pool.returnResource(jedis);
            userEnd(err);
        }
    }

    /**
     * 查keys
     *
     * @param pattern pattern
     * @throws Exception 异常
     */
    public Set<String> keys(String pattern) throws Exception {
        Exception err = null;
        Jedis jedis = null;
        userStart();
        Set<String> keys = null;
        try {
            log.debug("redis [" + name + "] keys pattern [" + pattern + "] start ");
            jedis = pool.getResource();
            keys = jedis.keys(pattern);
            log.debug("redis [" + name + "] keys pattern [" + pattern + "] end ");
            return keys;
        } catch (Exception e) {
            log.error("redis [" + name + "] keys pattern [" + pattern + "] error {} ", e);
            err = e;
            throw e;
        } finally {
            pool.returnResource(jedis);
            userEnd(err);
        }
    }

    /**
     * 取值
     *
     * @param key key
     * @throws Exception 异常
     */
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
            pool.returnResource(jedis);
            userEnd(err);
        }
    }

    /**
     * 删除
     *
     * @param key key
     * @throws Exception 异常
     */
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
            pool.returnResource(jedis);
            userEnd(err);
        }
    }
}
