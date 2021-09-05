package com.teamide.toolbox.zookeeper.service;

import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;

import java.util.HashMap;
import java.util.Map;
import java.util.concurrent.*;


/**
 * Zookeeper服务
 *
 * @author 朱亮
 * @date 2021/08/29
 */
@Service
@Slf4j
public class ZookeeperService {


    private final Map<String, ZookeeperCurator> cache = new HashMap<>();

    public ZookeeperCurator curator(String url) throws Exception {
        return curator(url, null);
    }

    public ZookeeperCurator curator(String url, Long automaticShutdown) throws Exception {
        if (automaticShutdown == null) {
            automaticShutdown = 60 * 10L; // 默认10分钟自动关闭
        }

        ZookeeperCurator curator = cache.get(url);
        if (curator == null || !curator.isStarted()) {
            synchronized (cache) {
                curator = cache.get(url);
                if (curator == null || !curator.isStarted()) {
                    if (curator == null) {
                        log.debug("curator url [" + url + "] is null,now create curator");
                    } else {
                        log.warn("curator url [" + url + "] is closed,now recreate curator");
                    }
                    curator = new ZookeeperCurator(url, automaticShutdown);
                    // 先检查某个节点，确定能连接成功
                    curator.checkExists("/");
                    cache.put(url, curator);
                }
            }
        }
        return curator;
    }


    /**
     * 创建线程池
     * <p>
     * ThreadPoolExecutor + LinkedBlockingQueue
     * <p>
     * 阻塞队列
     *
     * @return Executor 线程池
     */
    public Executor createExecutor() {
        // 核心工作线程数量
        int workerCore = Runtime.getRuntime().availableProcessors();
        // 最大工作线程数量
        int workerMax = workerCore + 2;

        // 空闲时长
        long keepAliveTime = 60;
        TimeUnit timeUnit = TimeUnit.SECONDS;

        // 如果队列大小小于0 则配置无界队列
        BlockingQueue<Runnable> queue = new LinkedBlockingQueue<>(workerMax);

        return new ThreadPoolExecutor(
                workerCore,
                workerMax,
                keepAliveTime, timeUnit,
                queue,
                Executors.defaultThreadFactory(),
                new CustomRejectedExecutionHandler());
    }

    /**
     * 自定义 线程池对拒绝任务的处理策略
     * <p>
     * 阻塞队列
     *
     * @author ZhuLiang
     * @date 2021/08/06
     */
    public static class CustomRejectedExecutionHandler implements RejectedExecutionHandler {
        @Override
        public void rejectedExecution(Runnable r, ThreadPoolExecutor executor) {
            try {
                // 核心改造点，由blockingqueue的offer改成put阻塞方法
                executor.getQueue().put(r);
            } catch (InterruptedException e) {
                e.printStackTrace();
            }
        }
    }
}
