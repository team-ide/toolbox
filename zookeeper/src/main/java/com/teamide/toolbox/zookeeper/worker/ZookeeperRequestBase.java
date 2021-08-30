package com.teamide.toolbox.zookeeper.worker;

import com.teamide.toolbox.worker.ToolboxWorkRequest;
import lombok.Data;

/**
 * @author 朱亮
 * @date 2021/08/30
 */
@Data
public class ZookeeperRequestBase extends ToolboxWorkRequest {

    /**
     * Zookeeper连接地址
     */
    private String url;
}
