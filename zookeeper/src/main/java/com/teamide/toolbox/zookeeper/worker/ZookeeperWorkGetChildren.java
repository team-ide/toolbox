package com.teamide.toolbox.zookeeper.worker;

import com.teamide.toolbox.worker.ToolboxWork;
import com.teamide.toolbox.zookeeper.service.ZookeeperCurator;
import com.teamide.toolbox.zookeeper.service.ZookeeperService;
import lombok.Builder;
import lombok.Data;
import org.apache.commons.lang3.StringUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.List;

/**
 * @author 朱亮
 * @date 2021/08/30
 */
@Service
public class ZookeeperWorkGetChildren implements ToolboxWork<ZookeeperWorkGetChildren.ZookeeperGetChildrenRequest, ZookeeperWorkGetChildren.ZookeeperGetChildrenResponse> {

    @Autowired
    private ZookeeperService zookeeperService;

    public Class<ZookeeperGetChildrenRequest> getRequestClass() {
        return ZookeeperGetChildrenRequest.class;
    }

    /**
     * 工作
     *
     * @param request 请求
     * @return {@link ZookeeperGetChildrenRequest}
     * @throws Exception 异常
     */
    @Override
    public ZookeeperGetChildrenResponse work(ZookeeperGetChildrenRequest request) throws Exception {
        ZookeeperGetChildrenResponse response = new ZookeeperGetChildrenResponse();
        ZookeeperCurator curator = zookeeperService.curator(request.getUrl());
        if (StringUtils.isNoneEmpty(request.getPath())) {
            final String path = request.getPath();
            if (curator.checkExists(path)) {
                List<String> children = curator.getChildren(path);
                response.setChildren(children);
            }
        }
        return response;
    }


    /**
     * @author 朱亮
     * @date 2021/08/30
     */
    @Data
    public static class ZookeeperGetChildrenRequest extends ZookeeperRequestBase {

        /**
         * 需要查询的路径
         */
        private String path;
    }

    /**
     * @author 朱亮
     * @date 2021/08/30
     */
    @Data
    public static class ZookeeperGetChildrenResponse extends ZookeeperResponseBase {

        /**
         * 子列表
         */
        private List<String> children;

    }
}
