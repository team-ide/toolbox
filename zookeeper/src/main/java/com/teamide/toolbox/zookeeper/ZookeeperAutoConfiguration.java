package com.teamide.toolbox.zookeeper;

import com.teamide.toolbox.worker.ToolboxWorkerCache;
import com.teamide.toolbox.zookeeper.worker.ZookeeperWorker;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.annotation.Configuration;


/**
 * Zookeeper组件自动配置
 *
 * @author 朱亮
 * @date 2021/08/29
 */
@Configuration
@ComponentScan
public class ZookeeperAutoConfiguration {

    @Autowired
    private ToolboxWorkerCache cache;

    @Autowired
    private ZookeeperWorker worker;

    @Bean
    public ZookeeperWorker getWorker() {
        cache.add(worker);
        return worker;
    }

}
