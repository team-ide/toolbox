package com.teamide.toolbox.web.controller;

import com.teamide.toolbox.util.ResponseBean;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import java.util.HashMap;
import java.util.Map;

@RestController
@RequestMapping("/server")
public class ServerController {

    /**
     * @description:
     * @param: * @param:
     * @return: java.lang.Object
     * @throws:
     * @author: 朱亮
     * @date: 2021/8/27 15:22
     */
    @RequestMapping(path = "/open", method = RequestMethod.POST)
    public Object open() {
        Map<String,Object> data = new HashMap<>();


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
