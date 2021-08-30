package com.teamide.toolbox.zookeeper.worker;

import com.teamide.toolbox.worker.ToolboxWork;
import com.teamide.toolbox.zookeeper.service.ZookeeperCurator;
import com.teamide.toolbox.zookeeper.service.ZookeeperService;
import lombok.Builder;
import lombok.Data;
import org.apache.commons.lang3.StringUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

/**
 * @author 朱亮
 * @date 2021/08/30
 */
@Service
public class ZookeeperWorkDelete implements ToolboxWork<ZookeeperWorkDelete.ZookeeperDeleteRequest, ZookeeperWorkDelete.ZookeeperDeleteResponse> {

    @Autowired
    private ZookeeperService zookeeperService;

    /**
     * 工作
     *
     * @param request 请求
     * @return {@link ZookeeperDeleteResponse}
     * @throws Exception 异常
     */
    @Override
    public ZookeeperDeleteResponse work(ZookeeperDeleteRequest request) throws Exception {
        ZookeeperDeleteResponse response = ZookeeperDeleteResponse.builder().build();
        ZookeeperCurator curator = zookeeperService.curator(request.getUrl());
        if (StringUtils.isNoneEmpty(request.getPath())) {
            curator.delete(request.getPath());
        }
        if (request.getPaths() != null && request.getPaths().length > 0) {
            for (String path : request.getPaths()) {
                curator.delete(path);
            }
        }
        return response;
    }


    /**
     * @author 朱亮
     * @date 2021/08/30
     */
    @Data
    @Builder
    public static class ZookeeperDeleteRequest extends ZookeeperRequestBase {

        /**
         * 需要删除的路径
         */
        private String path;

        /**
         * 需要删除的路径
         */
        private String[] paths;
    }

    /**
     * @author 朱亮
     * @date 2021/08/30
     */
    @Data
    @Builder
    public static class ZookeeperDeleteResponse extends ZookeeperResponseBase {

    }
}
