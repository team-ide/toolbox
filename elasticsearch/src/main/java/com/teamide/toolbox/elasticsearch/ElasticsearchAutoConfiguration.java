package com.teamide.toolbox.elasticsearch;

import com.teamide.toolbox.elasticsearch.worker.ElasticsearchWorker;
import com.teamide.toolbox.worker.ToolboxWorkerCache;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.annotation.Configuration;

@Configuration
@ComponentScan
public class ElasticsearchAutoConfiguration {

    @Autowired
    ToolboxWorkerCache cache;

    @Autowired
    ElasticsearchWorker worker;

    @Bean
    public ElasticsearchWorker getElasticsearchWorker() {
        cache.add(worker);
        return worker;
    }
}
