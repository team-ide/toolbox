package com.teamide.toolbox.web.controller;

import com.teamide.toolbox.bean.ResponseBean;
import com.teamide.toolbox.worker.ToolboxWorkerCache;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import java.util.ArrayList;
import java.util.HashMap;
import java.util.List;
import java.util.Map;

/**
 * 服务器控制器
 *
 * @author 朱亮
 * @date 2021/09/08
 */
@RestController
@RequestMapping("/server")
public class ServerController {
    /**
     * 工人缓存
     */
    @Autowired
    ToolboxWorkerCache workerCache;

    @RequestMapping(path = "/open", method = RequestMethod.POST)
    public Object open() {
        Map<String, Object> data = new HashMap<>(1);

        List<Map<String, Object>> workers = new ArrayList<>();
        data.put("workers", workers);
        workerCache.getWorkers().forEach(worker -> {
            Map<String, Object> workerData = new HashMap<>(4);
            workers.add(workerData);

            workerData.put("name", worker.name());
            workerData.put("comment", worker.comment());
            workerData.put("icon", worker.icon());

            List<Map<String, Object>> works = new ArrayList<>();
            workerData.put("works", works);

            worker.workMap().forEach((key, value) -> {
                Map<String, Object> workData = new HashMap<>(1);
                works.add(workData);

                workData.put("name", key);
            });

        });

        return ResponseBean.success(data);
    }

    @RequestMapping(path = "/close", method = RequestMethod.POST)
    public Object close() {

        return ResponseBean.success();
    }

    @RequestMapping(path = "/session", method = RequestMethod.POST)
    public Object session() {

        return ResponseBean.success();
    }
}
