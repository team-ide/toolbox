package com.teamide.toolbox.elasticsearch.worker;

import com.teamide.toolbox.worker.ToolboxWorkRequest;
import lombok.Data;

/**
 * @author 朱亮
 * @date 2021/08/30
 */
@Data
public class ElasticsearchRequestBase extends ToolboxWorkRequest {

    /**
     * Zookeeper连接地址
     */
    private String url;

    /**
     * 自动关闭
     */
    private Long automaticShutdown;
}
