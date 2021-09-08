package com.teamide.toolbox.redis.worker;

import com.teamide.toolbox.worker.ToolboxWork;
import com.teamide.toolbox.worker.ToolboxWorker;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Component;

import java.util.HashMap;
import java.util.Map;

/**
 * Redis Worker
 *
 * @author 朱亮
 * @date 2021/09/08
 */
@Component
public class RedisWorker implements ToolboxWorker {

    @Override
    public String name() {
        return "redis";
    }

    @Override
    public String text() {
        return "Redis";
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
    RedisWorkKeys keys;

    @Autowired
    RedisWorkGet get;

    @Autowired
    RedisWorkSave save;

    @Autowired
    RedisWorkDelete delete;

    @Override
    public Map<String, ToolboxWork<?, ?>> workMap() {
        Map<String, ToolboxWork<?, ?>> workMap = new HashMap<>(4);
        workMap.put("keys", keys);
        workMap.put("get", get);
        workMap.put("save", save);
        workMap.put("delete", delete);
        return workMap;
    }
}
