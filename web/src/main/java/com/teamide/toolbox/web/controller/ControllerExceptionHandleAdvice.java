package com.teamide.toolbox.web.controller;

import com.teamide.toolbox.bean.ResponseBean;
import com.teamide.toolbox.exception.BaseException;
import lombok.extern.slf4j.Slf4j;
import org.apache.commons.lang3.StringUtils;
import org.springframework.web.bind.annotation.ExceptionHandler;
import org.springframework.web.bind.annotation.RestControllerAdvice;

import javax.servlet.http.HttpServletRequest;

@Slf4j
@RestControllerAdvice
public class ControllerExceptionHandleAdvice {

    @ExceptionHandler
    public Object error(HttpServletRequest req, Throwable e) {
        log.error("request path [" + req.getPathInfo() + "] error {}", e);
        return toResult(e);
    }

    /**
     * 将异常转成control返回对象
     *
     * @param e 异常
     * @return result 转后的返回
     */
    public ResponseBean toResult(Throwable e) {
        ResponseBean error = ResponseBean.fail();
        error.setMsg(getErrMsg(e));
        if (e instanceof BaseException) {
            BaseException base = (BaseException) e;
            error.setCode(base.getCode());
        }
        return error;
    }

    /**
     * 递归获取异常信息
     *
     * @param e error
     * @return errMsg
     */
    public String getErrMsg(Throwable e) {
        String errMsg = e.getMessage();
        if (StringUtils.isEmpty(errMsg)) {
            if (e instanceof NullPointerException) {
                errMsg = "NullPointerException";
            }
        }
        if (StringUtils.isEmpty(errMsg)) {
            if (e.getCause() != null && e.getCause() != e) {
                errMsg = getErrMsg(e.getCause());
            }
        }
        return errMsg;
    }
}
