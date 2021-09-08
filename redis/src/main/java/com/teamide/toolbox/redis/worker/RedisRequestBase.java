package com.teamide.toolbox.redis.worker;

import com.teamide.toolbox.worker.ToolboxWorkRequest;
import lombok.Data;
import lombok.EqualsAndHashCode;

/**
 * @author 朱亮
 * @date 2021/08/30
 */
@EqualsAndHashCode(callSuper = true)
@Data
public class RedisRequestBase extends ToolboxWorkRequest {

    // Redis address
    private String address;

    // Redis auth
    private String auth;

    // Redis cluster
    private boolean cluster;

    /**
     * 自动关闭
     */
    private Long automaticShutdown;
}
