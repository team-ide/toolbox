package com.teamide.toolbox.web.controller;

import com.teamide.toolbox.bean.ResponseBean;
import com.teamide.toolbox.exception.BaseException;
import lombok.extern.slf4j.Slf4j;
import org.apache.commons.lang3.StringUtils;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;

import javax.servlet.http.HttpServletRequest;
import javax.servlet.http.HttpServletResponse;

/**
 * @description: TODO 类描述
 * @author: 朱亮
 * @date: 2021/9/6 16:42
 **/
@Slf4j
@RestControllerAdvice
public class ControllerExceptionHandleAdvice {

    @ExceptionHandler
    public Object error(HttpServletRequest req, HttpServletResponse res, Throwable e) {
        log.error("request path [" + req.getPathInfo() + "] error {}", e);
        ResponseBean result = toResult(e);
        return result;
    }

    /**
     * 将异常转成control返回对象
     *
     * @param e 异常
     * @return result 转后的返回
     */
    public ResponseBean toResult(Throwable e) {
        ResponseBean error = ResponseBean.fail();
        error.setMsg(getErrmsg(e));
        if (e instanceof BaseException) {
            BaseException base = (BaseException) e;
            error.setCode(base.getCode());
        }
        return error;
    }

    /**
     * 递归获取异常信息
     *
     * @param e
     * @return
     */
    public String getErrmsg(Throwable e) {
        String errmsg = e.getMessage();
        if (StringUtils.isEmpty(errmsg)) {
            if (e instanceof NullPointerException) {
                errmsg = "NullPointerException";
            }
        }
        if (StringUtils.isEmpty(errmsg)) {
            if (e.getCause() != null && e.getCause() != e) {
                errmsg = getErrmsg(e.getCause());
            }
        }
        return errmsg;
    }
}
