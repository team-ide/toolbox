package com.teamide.toolbox.elasticsearch.service;

import com.teamide.toolbox.worker.ToolboxAutomaticShutdown;
import lombok.extern.slf4j.Slf4j;

import java.net.ConnectException;


/**
 * zookeeper 操作管理
 * 查询节点
 *
 * @author 朱亮
 * @date 2021/08/29
 */
@Slf4j
public class ElasticsearchCurator implements ToolboxAutomaticShutdown {

    /**
     * ES地址
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
    public ElasticsearchCurator(String url, long automaticShutdown) {
        this.url = url;
        this.automaticShutdown = automaticShutdown;

        log.info("es [" + url + "] curator connect to [" + this.url + "] start");


        log.info("es [" + url + "] curator connect to [" + this.url + "] end");
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
            boolean shouldClose = exception instanceof ConnectException;
            if (shouldClose) {
                stop();
            }
        }
    }


    @Override
    public void stop() {
        synchronized (lock) {
            if (!this.started) {
                return;
            }
            log.debug("es [" + url + "] curator close ");
            this.started = false;
//            this.curator.close();
        }
    }

    public void work() throws Exception {
        Exception err = null;
        userStart();
        try {
            getClass().newInstance();
        } catch (Exception e) {
            err = e;
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

    @Override
    public boolean started() {
        return this.started;
    }
}
