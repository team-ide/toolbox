package com.teamide.toolbox.redis.service;

import java.util.Set;


public interface RedisDo {

    boolean isStarted() throws Exception;

    String set(String key, String value) throws Exception;

    Set<String> keys(String pattern, int size) throws Exception;

    String get(String key) throws Exception;

    void delete(String key) throws Exception;
}
