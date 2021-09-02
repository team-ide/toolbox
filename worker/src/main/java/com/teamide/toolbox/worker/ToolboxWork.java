package com.teamide.toolbox.worker;


import com.teamide.toolbox.util.BeanMapUtil;

import java.lang.reflect.ParameterizedType;
import java.util.Map;

/**
 * 工具箱的工作者工作
 *
 * @author 朱亮
 * @date 2021/08/30
 */
public interface ToolboxWork<P extends ToolboxWorkRequest, R extends ToolboxWorkResponse> {


    public Class<P> getRequestClass();

    public default P getRequest(Map<String, Object> data) throws Exception {
        P p = null;
        if (data == null) {
            p = getRequestClass().newInstance();
        } else {
            p = BeanMapUtil.toBean(data, getRequestClass());
        }
        return p;
    }

    public default Class<R> getResponseClass() {
        Class<R> rClass = (Class<R>) ((ParameterizedType) getClass().getGenericSuperclass()).getActualTypeArguments()[1];
        return rClass;
    }

    /**
     * 工作
     *
     * @param request 请求
     * @return {@link ToolboxWorkResponse}
     * @throws Exception 异常
     */
    public R work(P request) throws Exception;
}
