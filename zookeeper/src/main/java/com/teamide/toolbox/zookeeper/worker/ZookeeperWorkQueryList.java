package com.teamide.toolbox.zookeeper.worker;

import com.teamide.toolbox.worker.ToolboxWork;
import com.teamide.toolbox.zookeeper.service.ZookeeperService;
import lombok.Builder;
import lombok.Data;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

/**
 * @author 朱亮
 * @date 2021/08/30
 */
@Service
public class ZookeeperWorkQueryList implements ToolboxWork<ZookeeperWorkQueryList.ZookeeperQueryListRequest, ZookeeperWorkQueryList.ZookeeperQueryListResponse> {

    @Autowired
    private ZookeeperService zookeeperService;

    public Class<ZookeeperQueryListRequest> getRequestClass() {
        return ZookeeperQueryListRequest.class;
    }

    /**
     * 工作
     *
     * @param request 请求
     * @return {@link ZookeeperQueryListResponse}
     * @throws Exception 异常
     */
    @Override
    public ZookeeperQueryListResponse work(ZookeeperQueryListRequest request) throws Exception {
        return null;
    }


    /**
     * @author 朱亮
     * @date 2021/08/30
     */
    @Data
    public static class ZookeeperQueryListRequest extends ZookeeperRequestBase {

        /**
         * 搜索文案
         */
        private String search;
    }

    /**
     * @author 朱亮
     * @date 2021/08/30
     */
    @Data
    public static class ZookeeperQueryListResponse extends ZookeeperResponseBase {

    }
}
