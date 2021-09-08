package com.teamide.toolbox.zookeeper.service;

import com.teamide.toolbox.worker.ToolboxAutomaticShutdown;
import lombok.extern.slf4j.Slf4j;
import org.apache.curator.framework.CuratorFramework;
import org.apache.curator.framework.CuratorFrameworkFactory;
import org.apache.curator.framework.api.CreateBuilder;
import org.apache.curator.retry.ExponentialBackoffRetry;
import org.apache.zookeeper.CreateMode;
import org.apache.zookeeper.KeeperException;
import org.apache.zookeeper.data.Stat;

import java.net.ConnectException;
import java.nio.charset.StandardCharsets;
import java.util.List;
import java.util.Timer;
import java.util.TimerTask;


/**
 * zookeeper 操作管理
 * 查询节点
 *
 * @author 朱亮
 * @date 2021/08/29
 */
@Slf4j
public class ZookeeperCurator implements ToolboxAutomaticShutdown {

    /**
     * ZK地址
     */
    private final String url;

    /**
     * 自动关闭时长 单位秒
     */
    private final long automaticShutdown;

    /**
     * 最后使用时间戳
     */
    private long lastUseTimestamp = System.currentTimeMillis();

    /**
     * ZK客户端管理工具
     */
    private final CuratorFramework curator;

    /**
     * 已启动标识
     */
    private boolean started = true;

    /**
     * 用于块级锁
     */
    private final Object lock = new Object();


    /**
     * 传入Zookeeper地址和自动关闭时长构建Curator
     *
     * @param url               Zookeeper地址
     * @param automaticShutdown 自动关闭时长，单位S，超出该时间不再操作，则关闭该连接
     */
    public ZookeeperCurator(String url, long automaticShutdown) throws Exception {
        this.url = url;
        this.automaticShutdown = automaticShutdown;
        // 重试策略
        ExponentialBackoffRetry retry = new ExponentialBackoffRetry(1000, 10);

        log.info("zk [" + url + "] curator connect to [" + this.url + "] start");

        // zkClint连接地址
        this.curator = CuratorFrameworkFactory.builder().connectString(this.url)
                // 连接超时时间
                .connectionTimeoutMs(60 * 1000)
                // 会话超时时间
                .sessionTimeoutMs(60 * 1000)
                .retryPolicy(retry)
                .build();
        this.curator.start();
        try {

            this.curator.getZookeeperClient().getZooKeeper().exists("/", false);
        } catch (Exception e) {
            log.info("zk [" + url + "] curator connect to [" + this.url + "] error {}", e);
            this.stop();
            throw e;
        }
        log.info("zk [" + url + "] curator connect to [" + this.url + "] end");
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
            boolean shouldClose = false;
            // 会话过期 则停止 提供服务
            if (exception instanceof KeeperException.SessionExpiredException) {
                shouldClose = true;
            } else if (exception instanceof KeeperException.SessionClosedRequireAuthException) {
                shouldClose = true;
            } else if (exception instanceof KeeperException.ConnectionLossException) {
                shouldClose = true;
            } else if (exception instanceof ConnectException) {
                shouldClose = true;
            }
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
            log.debug("zk [" + url + "] curator close ");
            this.started = false;
            this.curator.close();
        }
    }

    /**
     * 创建节点
     *
     * @param path 路径
     * @param data 数据
     * @throws Exception 异常
     */
    public void create(String path, String data) throws Exception {
        create(path, data, null);
    }

    /**
     * 不存在则创建
     *
     * @param path 路径
     * @throws Exception 异常
     */
    public void createNotExists(String path) throws Exception {
        createNotExists(path, null, null);
    }

    /**
     * 不存在则创建
     *
     * @param path 路径
     * @param data 数据
     * @throws Exception 异常
     */
    public void createNotExists(String path, String data) throws Exception {
        createNotExists(path, data, null);
    }

    /**
     * 不存在则创建
     *
     * @param path 路径
     * @param mode 模式
     * @throws Exception 异常
     */
    public void createNotExists(String path, CreateMode mode) throws Exception {
        createNotExists(path, null, mode);
    }

    /**
     * 不存在则创建
     *
     * @param path 路径
     * @param data 数据
     * @param mode 模式
     * @throws Exception 异常
     */
    public void createNotExists(String path, String data, CreateMode mode) throws Exception {
        Exception err = null;
        userStart();
        try {
            if (checkExists(path)) {
                log.debug("zk [" + url + "] createNotExists node [" + path + "] checkExists true ");
                return;
            }
            log.debug("zk [" + url + "] createNotExists node [" + path + "] checkExists false ");
            create(path, data, mode);
        } catch (Exception e) {
            err = e;
            throw e;
        } finally {
            userEnd(err);
        }
    }

    /**
     * 验证节点是否存在
     *
     * @param path 路径
     * @return 是否存在
     * @throws Exception 异常
     */
    public boolean checkExists(String path) throws Exception {
        Exception err = null;
        userStart();
        try {
            log.debug("zk [" + url + "] checkExists node [" + path + "] start ");
            Stat stat = this.curator.checkExists().forPath(path);
            log.debug("zk [" + url + "] checkExists node [" + path + "] end ");
            return stat != null;
        } catch (Exception e) {
            log.error("zk [" + url + "] checkExists node [" + path + "] error {} ", e);
            err = e;
            throw e;
        } finally {
            userEnd(err);
        }
    }


    /**
     * 创建节点
     *
     * @param path 路径
     * @param data 数据
     * @param mode 模式
     * @throws Exception 异常
     */
    public void create(String path, String data, CreateMode mode) throws Exception {
        Exception err = null;
        userStart();
        try {
            log.debug("zk [" + url + "] create node [" + path + "] data [" + data + "] mode [" + mode + "] start ");
            CreateBuilder builder = this.curator.create();
            if (mode != null) {
                builder.withMode(mode);
            }
            if (data == null) {
                builder.creatingParentsIfNeeded().forPath(path);
            } else {
                builder.creatingParentsIfNeeded().forPath(path, data.getBytes(StandardCharsets.UTF_8));
            }
            log.debug("zk [" + url + "] create node [" + path + "] data [" + data + "] mode [" + mode + "] end ");
        } catch (Exception e) {
            log.error("zk [" + url + "] create node [" + path + "] data [" + data + "] mode [" + mode + "] error {} ", e);
            err = e;
            throw e;
        } finally {
            userEnd(err);
        }
    }

    /**
     * 修改节点数据
     *
     * @param path 路径
     * @param data 数据
     * @throws Exception 异常
     */
    public void setData(String path, String data) throws Exception {
        Exception err = null;
        userStart();
        try {

            log.debug("zk [" + url + "] setData node [" + path + "] data [" + data + "] start ");
            if (data == null) {
                this.curator.setData().forPath(path);
            } else {
                this.curator.setData().forPath(path, data.getBytes(StandardCharsets.UTF_8));
            }
            log.debug("zk [" + url + "] setData node [" + path + "] data [" + data + "] end ");
        } catch (Exception e) {
            log.error("zk [" + url + "] setData node [" + path + "] data [" + data + "] error {} ", e);
            err = e;
            throw e;
        } finally {
            userEnd(err);
        }
    }

    /**
     * 查询子节点
     *
     * @param path 路径
     * @throws Exception 异常
     */
    public List<String> getChildren(String path) throws Exception {
        Exception err = null;
        userStart();
        try {
            log.debug("zk [" + url + "] getChildren node [" + path + "] start ");
            List<String> children = this.curator.getChildren().forPath(path);
            log.debug("zk [" + url + "] getChildren node [" + path + "] end ");

            return children;
        } catch (Exception e) {
            log.error("zk [" + url + "] getData node [" + path + "] error {} ", e);
            err = e;
            throw e;
        } finally {
            userEnd(err);
        }
    }

    /**
     * 查询节点数据
     *
     * @param path 路径
     * @throws Exception 异常
     */
    public String getData(String path) throws Exception {
        Exception err = null;
        userStart();
        try {
            log.debug("zk [" + url + "] getData node [" + path + "] start ");
            byte[] bytes = this.curator.getData().forPath(path);
            log.debug("zk [" + url + "] getData node [" + path + "] end ");
            if (bytes == null) {
                return null;
            }
            return new String(bytes, StandardCharsets.UTF_8);
        } catch (Exception e) {
            log.error("zk [" + url + "] getData node [" + path + "] error {} ", e);
            err = e;
            throw e;
        } finally {
            userEnd(err);
        }
    }

    /**
     * 删除节点
     *
     * @param path 路径
     * @throws Exception 异常
     */
    public void delete(String path) throws Exception {
        Exception err = null;
        userStart();
        try {
            log.debug("zk [" + url + "] delete node [" + path + "] start ");
            Stat stat = this.curator.checkExists().forPath(path);
            log.debug("zk [" + url + "] delete node [" + path + "] end ");
            if (stat == null) {
                return;
            }
            this.curator.delete().deletingChildrenIfNeeded().withVersion(stat.getVersion()).forPath(path);
        } catch (Exception e) {
            log.error("zk [" + url + "] delete node [" + path + "] error {} ", e);
            err = e;
            throw e;
        } finally {
            userEnd(err);
        }
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
