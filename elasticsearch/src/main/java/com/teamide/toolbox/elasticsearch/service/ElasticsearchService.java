package com.teamide.toolbox.elasticsearch.service;

import lombok.extern.slf4j.Slf4j;
import org.springframework.stereotype.Service;

import java.util.HashMap;
import java.util.Map;


/**
 * Zookeeper服务
 *
 * @author 朱亮
 * @date 2021/08/29
 */
@Service
@Slf4j
public class ElasticsearchService {


    private final Map<String, ElasticsearchCurator> cache = new HashMap<>();

    public ElasticsearchCurator curator(String url) throws Exception {
        return curator(url, null);
    }

    public ElasticsearchCurator curator(String url, Long automaticShutdown) throws Exception {
        if (automaticShutdown == null) {
            automaticShutdown = 60 * 10L; // 默认10分钟自动关闭
        }

        ElasticsearchCurator curator = cache.get(url);
        if (curator == null || !curator.isStarted()) {
            synchronized (cache) {
                curator = cache.get(url);
                if (curator == null || !curator.isStarted()) {
                    if (curator == null) {
                        log.debug("curator url [" + url + "] is null,now create curator");
                    } else {
                        log.warn("curator url [" + url + "] is closed,now recreate curator");
                    }
                    curator = new ElasticsearchCurator(url, automaticShutdown);
                    // 先检查某个节点，确定能连接成功
//                    curator.checkExists("/");
                    cache.put(url, curator);
                }
            }
        }
        return curator;
    }


}
