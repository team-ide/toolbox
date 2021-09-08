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
    String name();

    /**
     * 文本
     *
     * @return {@link String}
     */
    String text();

    /**
     * 图标
     *
     * @return {@link String}
     */
    String icon();

    /**
     * 评论
     *
     * @return {@link String}
     */
    String comment();

    /**
     * 工作地图
     *
     * @return {@link Map}<{@link String}, {@link ToolboxWork}>
     */
    Map<String, ToolboxWork<?, ?>> workMap();

}
