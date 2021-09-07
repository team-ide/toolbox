package com.teamide.toolbox.elasticsearch.service;

import lombok.extern.slf4j.Slf4j;

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
public class ElasticsearchCurator {

    // ES地址
    private final String url;

    // 自动关闭时长 单位秒
    private final long automaticShutdown;

    // 最后使用时间戳
    private long lastUseTime = System.currentTimeMillis();


    // 已启动标识
    private boolean started = true;

    // 用于块级锁
    private final Object lock = new Object();

    //
    private Timer timer;

    /**
     * 传入Zookeeper地址和自动关闭时长构建Curator
     *
     * @param url               Zookeeper地址
     * @param automaticShutdown 自动关闭时长，单位S，超出该时间不再操作，则关闭该连接
     */
    public ElasticsearchCurator(String url, long automaticShutdown) throws Exception {
        this.url = url;
        this.automaticShutdown = automaticShutdown;

        log.info("es [" + url + "] curator connect to [" + this.url + "] start");


        log.info("es [" + url + "] curator connect to [" + this.url + "] end");
        if (this.automaticShutdown > 0) {
            timer = new Timer();
            startTimerTask();
        }
    }

    private void startTimerTask() {
//        log.debug("es [" + url + "] timer add task ");
        timer.schedule(new TimerTask() {
            @Override
            public void run() {
                checkAutomaticShutdown();
            }
        }, 1000);
    }

    private void checkAutomaticShutdown() {
//        log.debug("es [" + url + "] timer task checkAutomaticShutdown ");
        long now = System.currentTimeMillis();
        long timeout = now - lastUseTime;
        // 超时 关闭
        if (timeout > this.automaticShutdown * 1000) {
            log.warn("es [" + url + "] curator timeout need to stop");
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
            log.debug("es [" + url + "] curator close ");
            this.started = false;
//            this.curator.close();
        }
    }

}
