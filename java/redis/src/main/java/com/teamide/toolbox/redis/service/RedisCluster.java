package com.teamide.toolbox.redis.service;

import lombok.extern.slf4j.Slf4j;
import org.apache.commons.lang3.StringUtils;
import org.apache.commons.pool2.impl.GenericObjectPoolConfig;
import redis.clients.jedis.HostAndPort;
import redis.clients.jedis.Jedis;
import redis.clients.jedis.JedisCluster;
import redis.clients.jedis.JedisPool;
import redis.clients.jedis.exceptions.JedisConnectionException;

import java.util.*;


/**
 * Redis集群
 *
 * @author 朱亮
 * @date 2021/09/08
 */
@Slf4j
public class RedisCluster implements RedisDo {

    /**
     * 名称，用作日志区分
     */
    private final String name;

    /**
     * 集群地址
     */
    private final String address;

    /**
     * 鉴权
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
     * jedis cluster 客户端管理工具
     */
    private final JedisCluster cluster;

    /**
     * 已启动标识
     */
    private boolean started = true;

    /**
     * 用于块级锁
     */
    private final Object lock = new Object();


    public RedisCluster(String address, String auth, long automaticShutdown) {
        if (StringUtils.isEmpty(auth)) {
            auth = null;
        }
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
        GenericObjectPoolConfig<Jedis> config = new GenericObjectPoolConfig<>();
        this.cluster = new JedisCluster(hostAndPortSet, connectionTimeout, soTimeout, maxAttempts, auth, config);
        log.info("redis [" + this.name + "] create cluster end");
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
    @Override
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
    @Override
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
                    Set<String> ksB = jedis.keys(pattern);
                    if (ksB.size() > 0) {
                        ks.addAll(ksB);
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
    @Override
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
    @Override
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
