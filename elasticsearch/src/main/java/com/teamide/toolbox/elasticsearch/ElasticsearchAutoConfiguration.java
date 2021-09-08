package com.teamide.toolbox.elasticsearch;

import com.teamide.toolbox.elasticsearch.worker.ElasticsearchWorker;
import com.teamide.toolbox.worker.ToolboxWorkerCache;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.context.annotation.Bean;
import org.springframework.context.annotation.ComponentScan;
import org.springframework.context.annotation.Configuration;

/**
 * elasticsearch 自动配置
 *
 * @author 朱亮
 * @date 2021/09/08
 */
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
