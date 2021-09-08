package com.teamide.toolbox.zookeeper.worker;

import com.teamide.toolbox.worker.ToolboxWork;
import com.teamide.toolbox.zookeeper.service.ZookeeperCurator;
import com.teamide.toolbox.zookeeper.service.ZookeeperService;
import lombok.Data;
import lombok.EqualsAndHashCode;
import org.apache.commons.lang3.StringUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

/**
 * @author 朱亮
 * @date 2021/08/30
 */
@Service
public class ZookeeperWorkDelete implements ToolboxWork<ZookeeperWorkDelete.Request, ZookeeperWorkDelete.Response> {

    @Autowired
    ZookeeperService zookeeperService;

    public Class<Request> getRequestClass() {
        return Request.class;
    }

    /**
     * 工作
     *
     * @param request 请求
     * @return {@link Response}
     * @throws Exception 异常
     */
    @Override
    public Response work(Request request) throws Exception {
        Response response = new Response();
        ZookeeperCurator curator = zookeeperService.curator(request.getUrl(), request.getAutomaticShutdown());
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
    @EqualsAndHashCode(callSuper = true)
    @Data
    public static class Request extends ZookeeperRequestBase {

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
    @EqualsAndHashCode(callSuper = true)
    @Data
    public static class Response extends ZookeeperResponseBase {

    }
}
