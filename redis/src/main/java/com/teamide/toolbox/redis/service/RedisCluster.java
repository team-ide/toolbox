package com.teamide.toolbox.redis.service;

import lombok.extern.slf4j.Slf4j;
import org.apache.commons.lang3.StringUtils;
import org.apache.commons.pool2.impl.GenericObjectPoolConfig;
import redis.clients.jedis.HostAndPort;
import redis.clients.jedis.Jedis;
import redis.clients.jedis.JedisCluster;
import redis.clients.jedis.JedisPool;

import java.util.*;


@Slf4j
public class RedisCluster implements RedisDo {

    // Redis name
    private final String name;

    // Redis address
    private final String address;

    // Redis auth
    private final String auth;

    // 自动关闭时长 单位秒
    private final long automaticShutdown;

    // 最后使用时间戳
    private long lastUseTime = System.currentTimeMillis();

    // jedis cluster 客户端管理工具
    private final JedisCluster cluster;

    // 已启动标识
    private boolean started = true;

    // 用于块级锁
    private final Object lock = new Object();

    //
    private Timer timer;

    public RedisCluster(String address, String auth, long automaticShutdown) throws Exception {
        if (StringUtils.isEmpty(auth)) {
            auth = null;
        }
        StringBuffer name = new StringBuffer();

        Set<HostAndPort> hostAndPortSet = new HashSet<>();

        Arrays.stream(address.split(";")).forEach(one -> {
            HostAndPort hostAndPort = HostAndPort.from(one);
            hostAndPortSet.add(hostAndPort);
        });
        this.address = address;
        this.auth = auth;
        this.name = address;
        this.automaticShutdown = automaticShutdown;

        int connectionTimeout = 1000 * 10;
        int soTimeout = 1000 * 10;
        int maxAttempts = 10;
        log.info("redis [" + this.name + "] create cluster start");
        this.cluster = new JedisCluster(hostAndPortSet, connectionTimeout, soTimeout, maxAttempts, auth, new GenericObjectPoolConfig());
        log.info("redis [" + this.name + "] create cluster end");

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
            this.cluster.close();
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
        userStart();
        try {
            log.debug("redis [" + name + "] set key [" + key + "] value [" + value + "] start ");
            String code = cluster.set(key, value);
            log.debug("redis [" + name + "] set key [" + key + "] value [" + value + "] code [" + code + "] end ");
            return code;
        } catch (Exception e) {
            log.error("redis [" + name + "] set key [" + key + "] value [" + value + "] error {} ", e);
            err = e;
            throw e;
        } finally {
            userEnd(err);
        }
    }

    /**
     * 查keys
     *
     * @param pattern pattern
     * @throws Exception 异常
     */
    public Set<String> keys(String pattern, int size) throws Exception {
        Exception err = null;
        userStart();
        Set<String> keys = new HashSet<>();
        try {
            log.debug("redis [" + name + "] keys pattern [" + pattern + "] start ");

            Map<String, JedisPool> clusterNodes = cluster.getClusterNodes();
            Set<String> ks = new HashSet<>();
            for (Map.Entry<String, JedisPool> entry : clusterNodes.entrySet()) {
                JedisPool pool = entry.getValue();
                Jedis jedis = pool.getResource();
                // 判断非从节点(因为若主从复制，从节点会跟随主节点的变化而变化)
                if (!jedis.info("replication").contains("role:slave")) {
                    Set<String> ks_ = jedis.keys(pattern);
                    if (ks_.size() > 0) {
                        ks.addAll(ks_);
                    }
                }
                jedis.close();
            }
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
        userStart();
        try {
            log.debug("redis [" + name + "] get key [" + key + "] start ");
            String value = cluster.get(key);
            log.debug("redis [" + name + "] get key [" + key + "] end ");
            return value;
        } catch (Exception e) {
            log.error("redis [" + name + "] get key [" + key + "] error {} ", e);
            err = e;
            throw e;
        } finally {
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
        userStart();
        try {
            log.debug("redis [" + name + "] delete key [" + key + "] start ");
            cluster.del(key);
            log.debug("redis [" + name + "] delete key [" + key + "] end ");
        } catch (Exception e) {
            log.error("redis [" + name + "] delete key [" + key + "] error {} ", e);
            err = e;
            throw e;
        } finally {
            userEnd(err);
        }
    }
}
