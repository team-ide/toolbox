package com.teamide.toolbox.elasticsearch.worker;

import com.teamide.toolbox.worker.ToolboxWork;
import com.teamide.toolbox.worker.ToolboxWorker;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import java.util.HashMap;
import java.util.Map;

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
    ElasticsearchWorkDelete delete;

    @Override
    public Map<String, ToolboxWork<?, ?>> workMap() {
        Map<String, ToolboxWork<?, ?>> workMap = new HashMap<>();
        workMap.put("delete", delete);
        return workMap;
    }
}
