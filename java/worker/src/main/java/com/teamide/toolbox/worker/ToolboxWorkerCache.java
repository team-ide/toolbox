package com.teamide.toolbox.worker;

import org.springframework.stereotype.Component;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * 工具箱 Worker 缓存
 *
 * @author 朱亮
 * @date 2021/09/08
 */
@Component
public class ToolboxWorkerCache {

    private final Map<String, ToolboxWorker> workerMap = new HashMap<>();

    public ToolboxWorkerCache add(ToolboxWorker worker) {
        put(worker.name(), worker);
        return this;
    }

    public ToolboxWorkerCache put(String name, ToolboxWorker worker) {
        workerMap.put(name, worker);
        return this;
    }

    public ToolboxWorker get(String name) {
        return workerMap.get(name);
    }

    public List<ToolboxWorker> getWorkers() {
        return new ArrayList<>(workerMap.values());
    }
}