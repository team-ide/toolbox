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
public class ZookeeperWorkGet implements ToolboxWork<ZookeeperWorkGet.Request, ZookeeperWorkGet.Response> {

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
            final String path = request.getPath();
            if (curator.checkExists(path)) {
                String data = curator.getData(path);
                response.setPath(path);
                response.setData(data);
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
         * 需要查询数据的路径
         */
        private String path;
    }

    /**
     * @author 朱亮
     * @date 2021/08/30
     */
    @EqualsAndHashCode(callSuper = true)
    @Data
    public static class Response extends ZookeeperResponseBase {

        /**
         * 路径
         */
        private String path;

        /**
         * 路径数据
         */
        private String data;

    }
}
