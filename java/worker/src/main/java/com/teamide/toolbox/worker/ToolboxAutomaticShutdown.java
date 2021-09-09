package com.teamide.toolbox.worker;


/**
 * @description: 自动关闭
 * @author: 朱亮
 * @date: 2021/9/8 15:22
 **/
public interface ToolboxAutomaticShutdown {

    /**
     * 自动关闭时长
     *
     * @return 自动关闭时长 单位秒
     */
    long automaticShutdown();

    /**
     * 最后使用时间戳
     *
     * @return 最后使用时间戳
     */
    long lastUseTimestamp();

    /**
     * 通知停止
     */
    void stop();

    /**
     * 是否启动
     *
     * @return 是否启动
     */
    boolean started();

}
