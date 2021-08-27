package com.teamide.toolbox.web.controller;

import com.teamide.toolbox.util.ResponseBean;
import org.springframework.web.bind.annotation.RequestMapping;
import org.springframework.web.bind.annotation.RequestMethod;
import org.springframework.web.bind.annotation.RestController;

import java.util.HashMap;
import java.util.Map;

/**
 * @author 朱亮
 * @version 1.0.0
 * @ClassName ServerController.java
 * @Description TODO
 * @createTime 2021年08月27日 14:20:00
 */
@RestController
@RequestMapping("/server")
public class ServerController {

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
