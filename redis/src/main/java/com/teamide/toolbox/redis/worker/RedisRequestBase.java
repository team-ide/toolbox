package com.teamide.toolbox.redis.worker;

import com.teamide.toolbox.worker.ToolboxWorkRequest;
import lombok.Data;

/**
 * @author 朱亮
 * @date 2021/08/30
 */
@Data
public class RedisRequestBase extends ToolboxWorkRequest {

    // Redis host
    private String host;

    // Redis port
    private int port;

    // Redis auth
    private String auth;

    /**
     * 自动关闭
     */
    private Long automaticShutdown;
}
