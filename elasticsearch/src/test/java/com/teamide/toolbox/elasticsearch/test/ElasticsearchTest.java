package com.teamide.toolbox.elasticsearch.test;

import com.teamide.toolbox.elasticsearch.ElasticsearchAutoConfiguration;
import com.teamide.toolbox.elasticsearch.service.ElasticsearchCurator;
import com.teamide.toolbox.elasticsearch.service.ElasticsearchService;
import com.teamide.toolbox.worker.ToolboxWorkerCache;
import com.teamide.toolbox.worker.WorkerAutoConfiguration;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

/**
 * @description: TODO 类描述
 * @author: 朱亮
 * @date: 2021/8/27 17:04
 **/
@SpringBootTest(classes = {WorkerAutoConfiguration.class, ElasticsearchAutoConfiguration.class})
public class ElasticsearchTest {

    static final String url = "127.0.0.1:2181";

    @Autowired
    private ElasticsearchService elasticsearchService;

    @Autowired
    private ToolboxWorkerCache workerCache;


    @Test
    public void testCurator() throws Exception {
        System.out.println(workerCache.getWorkers().size());
        new Thread(() -> {
            try {
                ElasticsearchCurator curator = elasticsearchService.curator(url);

            } catch (Exception e) {
                e.printStackTrace();
            }
        }).start();
        Thread.sleep(1000 * 60 * 10);
    }

}
