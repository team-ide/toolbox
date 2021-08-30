package com.teamide.toolbox.worker;

import org.springframework.stereotype.Component;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

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
        List<ToolboxWorker> workers = new ArrayList<>();
        workerMap.values().forEach(one -> {
            workers.add(one);
        });
        return workers;
    }
}
