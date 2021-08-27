package com.teamide.toolbox.zookeeper.service;

import com.teamide.toolbox.zookeeper.bean.ZookeeperContext;
import com.teamide.toolbox.zookeeper.bean.ZookeeperNode;
import com.teamide.toolbox.zookeeper.bean.ZookeeperTree;
import lombok.extern.slf4j.Slf4j;
import org.apache.curator.framework.CuratorFramework;
import org.apache.curator.framework.CuratorFrameworkFactory;
import org.apache.curator.framework.api.CreateBuilder;
import org.apache.curator.retry.ExponentialBackoffRetry;
import org.apache.zookeeper.CreateMode;
import org.apache.zookeeper.WatchedEvent;
import org.apache.zookeeper.Watcher;
import org.apache.zookeeper.data.Stat;

import java.util.ArrayList;
import java.util.List;
import java.util.concurrent.Executor;

/**
 * @description: TODO 类描述
 * @author: 朱亮
 * @date: 2021/8/27 15:44
 **/
@Slf4j
public class ZookeeperListener {

    private final String url;
    private CuratorFramework curator;
    private final List<ZookeeperContext> contexts = new ArrayList<>();
    private final ZookeeperTree tree = new ZookeeperTree();
    private Watcher watcher;

    public ZookeeperListener(String url) {
        this.url = url;
    }

    public boolean started = false;

    public final void start(Executor executor) throws Exception {
        if (this.started) {
            return;
        }
        synchronized (this.contexts) {
            if (this.started) {
                return;
            }
            this.stop();
            // 重试策略
            ExponentialBackoffRetry retry = new ExponentialBackoffRetry(1000, 10);

            log.info("curator connect to [" + this.url + "] start");
            this.curator = CuratorFrameworkFactory.builder().connectString(this.url)// zkClint连接地址
                    .connectionTimeoutMs(30 * 1000)// 连接超时时间
                    .sessionTimeoutMs(30 * 1000)// 会话超时时间
                    .retryPolicy(retry)
                    .build();
            this.curator.start();
            log.info("curator connect to [" + this.url + "] end");
            executor.execute(() -> listen());
            this.started = true;
        }
    }

    private final void listen() {
        ZookeeperNode root = new ZookeeperNode();
        root.setPath("/");

        this.watcher = new Watcher() {
            @Override
            public void process(WatchedEvent event) {
                log.debug("listen node [" + event.getPath() + "] event [{}]", event.getType());
            }
        };

        listen(root);
        this.tree.setRoot(root);
    }


    private final void listen(ZookeeperNode node) {
        final String path = node.getPath();
        try {
            log.debug("listen node [" + path + "] start");
            if (!checkExists(path)) {
                log.warn("listen node [" + path + "] not exists");
                return;
            }

            byte[] date = this.curator.getData().usingWatcher(this.watcher).forPath(path);
            if (date == null) {
                log.debug("listen node [" + path + "] data [" + null + "] ");
            } else {
                log.debug("listen node [" + path + "] data [" + date.length + "] ");
            }
            List<String> children = this.curator.getChildren().usingWatcher(this.watcher).forPath(path);
            if (children != null && children.size() > 0) {
                log.debug("listen node [" + path + "] children [" + children.size() + "] ");
                children.forEach(child -> {
                    ZookeeperNode childNode = new ZookeeperNode();
                    String childPath = path;
                    if (childPath.endsWith("/")) {
                        childPath += child;
                    } else {
                        childPath += "/" + child;
                    }
                    childNode.setPath(childPath);
                    listen(childNode);
                });
            }
        } catch (Exception e) {
            log.error("listen node [" + path + "] error", e);
        }
    }


    public void create(String path, String data) throws Exception {
        create(path, data, null);
    }

    public void create(String path, String data, CreateMode mode) throws Exception {
        log.debug("create node [" + path + "] data [" + data + "] mode [" + mode + "] start ");
        CreateBuilder builder = this.curator.create();
        if (mode != null) {
            builder.withMode(mode);
        }
        if (data == null) {
            builder.creatingParentsIfNeeded().forPath(path);
        } else {
            builder.creatingParentsIfNeeded().forPath(path, data.getBytes("UTF-8"));
        }
        log.debug("create node [" + path + "] data [" + data + "] mode [" + mode + "] end ");
    }

    public void setData(String path, String data) throws Exception {
        log.debug("setData node [" + path + "] data [" + data + "] start ");
        if (data == null) {
            this.curator.setData().forPath(path);
        } else {
            this.curator.setData().forPath(path, data.getBytes("UTF-8"));
        }
        log.debug("setData node [" + path + "] data [" + data + "] end ");
    }

    public String getData(String path) throws Exception {
        log.debug("getData node [" + path + "] start ");
        byte[] bytes = this.curator.getData().forPath(path);
        log.debug("getData node [" + path + "] end ");
        if (bytes == null) {
            return null;
        }
        return new String(bytes, "UTF-8");
    }

    public void delete(String path) throws Exception {
        log.debug("delete node [" + path + "] start ");
        Stat stat = this.curator.checkExists().forPath(path);
        log.debug("delete node [" + path + "] end ");
        if (stat == null) {
            return;
        }
        curator.delete().deletingChildrenIfNeeded().withVersion(stat.getVersion()).forPath(path);
    }

    public boolean checkExists(String path) throws Exception {
        log.debug("checkExists node [" + path + "] start ");
        Stat stat = this.curator.checkExists().forPath(path);
        log.debug("checkExists node [" + path + "] end ");
        if (stat == null) {
            return false;
        }
        return true;
    }

    public void createNotExists(String path) throws Exception {
        createNotExists(path, null, null);
    }

    public void createNotExists(String path, CreateMode mode) throws Exception {
        createNotExists(path, null, mode);
    }

    public void createNotExists(String path, String data) throws Exception {
        createNotExists(path, data, null);
    }

    public void createNotExists(String path, String data, CreateMode mode) throws Exception {
        if (checkExists(path)) {
            return;
        }
        create(path, data, mode);
    }

    public final void stop() throws Exception {
        if (this.curator != null) {
            synchronized (this.contexts) {
                if (this.curator != null) {
                    log.info("curator close start");
                    this.curator.close();
                    log.info("curator close end");
                    this.curator = null;
                }
            }
        }
        this.started = false;
    }


    public final void destroy() throws Exception {
        this.stop();
    }

    public ZookeeperListener addContext(ZookeeperContext context) {
        this.contexts.add(context);
        return this;
    }

}
