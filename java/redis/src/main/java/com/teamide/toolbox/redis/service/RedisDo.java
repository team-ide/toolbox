package com.teamide.toolbox.redis.service;

import com.teamide.toolbox.worker.ToolboxAutomaticShutdown;

import java.util.Set;


/**
 * Redis 基本操作
 *
 * @author 朱亮
 * @date 2021/09/08
 */
public interface RedisDo extends ToolboxAutomaticShutdown {


    /**
     * 设置值
     *
     * @param key   key
     * @param value value
     * @return 状态
     * @throws Exception err
     */
    String set(String key, String value) throws Exception;

    /**
     * 搜索key
     *
     * @param pattern pattern
     * @param size    size
     * @return 结果
     * @throws Exception err
     */
    Set<String> keys(String pattern, int size) throws Exception;

    /**
     * 获取值
     *
     * @param key key
     * @return value
     * @throws Exception err
     */
    String get(String key) throws Exception;

    /**
     * 删除Key
     *
     * @param key key
     * @throws Exception err
     */
    int delete(String key) throws Exception;

    /**
     * 匹配pattern删除所有key
     *
     * @param pattern pattern
     * @throws Exception err
     */
    int deletePattern(String pattern) throws Exception;
}
