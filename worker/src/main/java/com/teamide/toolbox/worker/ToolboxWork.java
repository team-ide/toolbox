package com.teamide.toolbox.worker;


import com.teamide.toolbox.util.BeanMapUtil;

import java.util.Map;

/**
 * 工具箱的工作者工作
 *
 * @author 朱亮
 * @date 2021/08/30
 */
public interface ToolboxWork<P extends ToolboxWorkRequest, R extends ToolboxWorkResponse> {


    Class<P> getRequestClass();

    default P getRequest(Map<String, Object> data) throws Exception {
        if (data == null) {
            return getRequestClass().newInstance();
        } else {
            return BeanMapUtil.toBean(data, getRequestClass());
        }
    }


    /**
     * 工作
     *
     * @param request 请求
     * @return {@link ToolboxWorkResponse}
     * @throws Exception 异常
     */
    R work(P request) throws Exception;
}
