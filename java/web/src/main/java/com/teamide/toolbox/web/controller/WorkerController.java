package com.teamide.toolbox.web.controller;

import com.teamide.toolbox.bean.ResponseBean;
import com.teamide.toolbox.worker.ToolboxWork;
import com.teamide.toolbox.worker.ToolboxWorker;
import com.teamide.toolbox.worker.ToolboxWorkerCache;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.web.bind.annotation.*;

import java.util.Map;

/**
 * 工人控制器
 *
 * @author 朱亮
 * @date 2021/09/02
 */
@RestController
@RequestMapping("/server/worker")
public class WorkerController {

    /**
     * 工人缓存
     */
    @Autowired
    ToolboxWorkerCache workerCache;


    /**
     * 工人工作
     *
     * @param workerName 工人的名字
     * @param workerWork 工人工作
     * @param body       身体
     * @return {@link Object}
     * @throws Exception 异常
     */
    @RequestMapping(path = "/{workerName}/{workerWork}", method = RequestMethod.POST)
    public Object workerWork(@PathVariable("workerName") String workerName, @PathVariable("workerWork") String workerWork, @RequestBody(required = false) Map<String, Object> body) throws Exception {
        ToolboxWorker worker = workerCache.get(workerName);
        if (worker == null) {
            return ResponseBean.fail("worker [" + workerName + "] not exist.");
        }
        ToolboxWork work = null;
        if (worker.workMap() != null) {
            work = worker.workMap().get(workerWork);
        }
        if (work == null) {
            return ResponseBean.fail("worker [" + workerName + "]  work [" + workerWork + "] not exist.");
        }

        Object response = work.work(work.getRequest(body));
        return ResponseBean.success(response);
    }

}
