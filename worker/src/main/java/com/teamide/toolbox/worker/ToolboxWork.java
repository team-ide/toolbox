package com.teamide.toolbox.worker;


/**
 * 工具箱的工作者工作
 *
 * @author 朱亮
 * @date 2021/08/30
 */
public interface ToolboxWork<P extends ToolboxWorkRequest, R extends ToolboxWorkResponse> {

    /**
     * 工作
     *
     * @param request 请求
     * @return {@link ToolboxWorkResponse}
     * @throws Exception 异常
     */
    public R work(P request) throws Exception;
}
