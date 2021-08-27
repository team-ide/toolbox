package com.teamide.toolbox.zookeeper.service;

import com.teamide.toolbox.zookeeper.bean.ZookeeperContext;
import org.springframework.stereotype.Service;

import java.util.Map;
import java.util.concurrent.*;

/**
 * @description: TODO 类描述
 * @author: 朱亮
 * @date: 2021/8/27 15:44
 **/
@Service
public class ZookeeperService {

    private final Map<String, ZookeeperListener> context_listener_cache = new ConcurrentHashMap<>();

    private final Executor executor = createExecutor();


    public ZookeeperListener listen(ZookeeperContext context) throws Exception {
        ZookeeperListener listener = context_listener_cache.get(context.getUrl());
        if (listener == null) {
            listener = context_listener_cache.computeIfAbsent(context.getUrl(), ZookeeperListener::new);
        }
        listener.addContext(context);
        listener.start(executor);
        return listener;
    }


    /**
     * 根据配置获取线程池
     * <p>
     * ThreadPoolExecutor + LinkedBlockingQueue
     * <p>
     * 阻塞队列
     *
     * @return Executor
     */
    public Executor createExecutor() {
        // 核心工作线程数量
        int workeCore = Runtime.getRuntime().availableProcessors();
        // 最大工作线程数量
        int workeMax = workeCore + 2;
        // 队列大小
        int queueSize = workeMax;
        // 每次拉取多少数据
        int maxPollRecords = 30;

        // 空闲时长
        long keepAliveTime = 60;
        TimeUnit timeUnit = TimeUnit.SECONDS;

        // 如果队列大小小于0 则配置无界队列
        BlockingQueue<Runnable> queue = null;
        if (queueSize >= 0) {
            queue = new LinkedBlockingQueue<Runnable>(queueSize);
        } else {
            queue = new LinkedBlockingQueue<Runnable>();
        }
        ThreadPoolExecutor pool = new ThreadPoolExecutor(
                workeCore,
                workeMax,
                keepAliveTime, timeUnit,
                queue,
                Executors.defaultThreadFactory(),
                new CustomRejectedExecutionHandler());

        return pool;
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
