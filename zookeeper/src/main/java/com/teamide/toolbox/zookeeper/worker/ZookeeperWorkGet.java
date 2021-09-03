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
public class ZookeeperWorkGet implements ToolboxWork<ZookeeperWorkGet.ZookeeperGetRequest, ZookeeperWorkGet.ZookeeperGetResponse> {

    @Autowired
    private ZookeeperService zookeeperService;

    public Class<ZookeeperGetRequest> getRequestClass() {
        return ZookeeperGetRequest.class;
    }

    /**
     * 工作
     *
     * @param request 请求
     * @return {@link ZookeeperGetResponse}
     * @throws Exception 异常
     */
    @Override
    public ZookeeperGetResponse work(ZookeeperGetRequest request) throws Exception {
        ZookeeperGetResponse response = new ZookeeperGetResponse();
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
    @Data
    public static class ZookeeperGetRequest extends ZookeeperRequestBase {

        /**
         * 需要查询数据的路径
         */
        private String path;
    }

    /**
     * @author 朱亮
     * @date 2021/08/30
     */
    @Data
    public static class ZookeeperGetResponse extends ZookeeperResponseBase {

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
