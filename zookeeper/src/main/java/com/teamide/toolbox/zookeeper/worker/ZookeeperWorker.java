package com.teamide.toolbox.zookeeper.worker;

import com.teamide.toolbox.worker.ToolboxWorker;
import com.teamide.toolbox.worker.ToolboxWork;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import java.util.HashMap;
import java.util.Map;

@Component
public class ZookeeperWorker implements ToolboxWorker {

    @Override
    public String name() {
        return "zookeeper";
    }

    @Override
    public String text() {
        return "Zookeeper";
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
    ZookeeperWorkGet get;

    @Autowired
    ZookeeperWorkGetChildren getChildren;

    @Autowired
    ZookeeperWorkSave save;

    @Autowired
    ZookeeperWorkDelete delete;

    @Override
    public Map<String, ToolboxWork<?, ?>> workMap() {
        Map<String, ToolboxWork<?, ?>> workMap = new HashMap<>();
        workMap.put("get", get);
        workMap.put("getChildren", getChildren);
        workMap.put("save", save);
        workMap.put("delete", delete);
        return workMap;
    }
}
