package com.teamide.toolbox.elasticsearch;

import com.teamide.toolbox.elasticsearch.worker.ElasticsearchWorker;
import com.teamide.toolbox.worker.ToolboxWorkerCache;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.annotation.Configuration;

/**
 * @author 朱亮
 * @version 1.0.0
 * @ClassName ElasticsearchAutoConfiguration.java
 * @Description TODO
 * @createTime 2021年08月26日 17:01:00
 */
@Configuration
@ComponentScan
public class ElasticsearchAutoConfiguration {

    @Autowired
    private ToolboxWorkerCache cache;

    @Autowired
    private ElasticsearchWorker worker;

    @Bean
    public ElasticsearchWorker getElasticsearchWorker() {
        cache.add(worker);
        return worker;
    }
}
