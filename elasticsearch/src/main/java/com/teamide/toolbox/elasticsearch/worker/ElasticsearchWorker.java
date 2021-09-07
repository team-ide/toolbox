package com.teamide.toolbox.elasticsearch.worker;

import com.teamide.toolbox.worker.ToolboxWork;
import com.teamide.toolbox.worker.ToolboxWorker;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import java.util.HashMap;
import java.util.Map;

/**
 * @description: TODO 类描述
 * @author: 朱亮
 * @date: 2021/8/30 10:53
 **/
@Component
public class ElasticsearchWorker implements ToolboxWorker {

    @Override
    public String name() {
        return "elasticsearch";
    }

    @Override
    public String text() {
        return "Elasticsearch";
    }

    @Override
    public String icon() {
        return null;
    }

    @Override
    public String comment() {
        return null;
    }


    @Autowired
    private ElasticsearchWorkDelete delete;

    @Override
    public Map<String, ToolboxWork> workMap() {
        Map<String, ToolboxWork> workMap = new HashMap<>();
        workMap.put("delete", delete);
        return workMap;
    }
}
