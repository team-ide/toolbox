package com.teamide.toolbox.zookeeper.test;

import com.teamide.toolbox.worker.ToolboxWorkerCache;
import com.teamide.toolbox.worker.WorkerAutoConfiguration;
import com.teamide.toolbox.zookeeper.ZookeeperAutoConfiguration;
import com.teamide.toolbox.zookeeper.service.ZookeeperCurator;
import com.teamide.toolbox.zookeeper.service.ZookeeperService;
import org.junit.jupiter.api.Test;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.boot.test.context.SpringBootTest;

/**
 * @description: TODO 类描述
 * @author: 朱亮
 * @date: 2021/8/27 17:04
 **/
@SpringBootTest(classes = {WorkerAutoConfiguration.class, ZookeeperAutoConfiguration.class})
public class ZookeeperTest {

    static final String url = "127.0.0.1:2181";

    @Autowired
    private ZookeeperService zookeeperService;

    @Autowired
    private ToolboxWorkerCache workerCache;


    @Test
    public void testCurator() throws Exception {
        System.out.println(workerCache.getWorkers().size());
        new Thread(() -> {
            try {
                ZookeeperCurator curator = zookeeperService.curator(url);
                curator.createNotExists("/data", "s");
                Thread.sleep(1000 * 3);

                curator = zookeeperService.curator(url);
                curator.delete("/data");

                Thread.sleep(1000 * 6);
                curator = zookeeperService.curator(url);
                curator.createNotExists("/data", "s");

                Thread.sleep(1000 * 5);
                curator = zookeeperService.curator(url);
                curator.delete("/data");
            } catch (Exception e) {
                e.printStackTrace();
            }
        }).start();
        Thread.sleep(1000 * 60 * 10);
    }

}
