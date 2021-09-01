package com.teamide.toolbox.worker;

import java.util.Map;

/**
 * 工具箱的工人
 *
 * @author 朱亮
 * @date 2021/08/30
 */
public interface ToolboxWorker {

    /**
     * 的名字
     *
     * @return {@link String}
     */
    public String name();

    /**
     * 文本
     *
     * @return {@link String}
     */
    public String text();

    /**
     * 图标
     *
     * @return {@link String}
     */
    public String icon();

    /**
     * 评论
     *
     * @return {@link String}
     */
    public String comment();

    /**
     * 工作地图
     *
     * @return {@link Map}<{@link String}, {@link ToolboxWork}>
     */
    public Map<String, ToolboxWork> workMap();

}
