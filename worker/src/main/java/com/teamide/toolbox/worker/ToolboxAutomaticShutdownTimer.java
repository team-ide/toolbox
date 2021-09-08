package com.teamide.toolbox.worker;

import org.springframework.scheduling.annotation.EnableScheduling;
import org.springframework.scheduling.annotation.Scheduled;
import org.springframework.stereotype.Component;

import java.util.*;

/**
 * 工具箱自动关闭定时器
 *
 * @author 朱亮
 * @date 2021/09/08
 */
@Component
@EnableScheduling
public class ToolboxAutomaticShutdownTimer {

    private final List<ToolboxAutomaticShutdown> cache = new ArrayList<>();

    public ToolboxAutomaticShutdownTimer add(ToolboxAutomaticShutdown worker) {
        cache.add(worker);
        return this;
    }

    /**
     * 每秒执行一次
     */
    @Scheduled(cron = "0/1 * * * * ?")
    public void automaticShutdown() {
        List<ToolboxAutomaticShutdown> list = new ArrayList<>(cache);
        list.forEach(one -> {
            if (one.automaticShutdown() <= 0) {
                return;
            }
            long now = System.currentTimeMillis();
            long timeout = now - one.lastUseTimestamp();
            // 超时 关闭
            long automaticShutdownTime = one.automaticShutdown() * 1000L;
            if (timeout > automaticShutdownTime) {
                one.stop();
                cache.remove(one);
            }
        });
    }

}
