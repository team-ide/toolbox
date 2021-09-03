package com.teamide.toolbox.zookeeper.worker;

import com.teamide.toolbox.worker.ToolboxWork;
import com.teamide.toolbox.zookeeper.service.ZookeeperCurator;
import com.teamide.toolbox.zookeeper.service.ZookeeperService;
import lombok.Builder;
import lombok.Data;
import org.apache.commons.lang3.StringUtils;
import org.apache.zookeeper.CreateMode;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

/**
 * @author 朱亮
 * @date 2021/08/30
 */
@Service
public class ZookeeperWorkSave implements ToolboxWork<ZookeeperWorkSave.ZookeeperSaveRequest, ZookeeperWorkSave.ZookeeperSaveResponse> {

    @Autowired
    private ZookeeperService zookeeperService;

    public Class<ZookeeperSaveRequest> getRequestClass() {
        return ZookeeperSaveRequest.class;
    }

    /**
     * 工作
     *
     * @param request 请求
     * @return {@link ZookeeperSaveResponse}
     * @throws Exception 异常
     */
    @Override
    public ZookeeperSaveResponse work(ZookeeperSaveRequest request) throws Exception {
        ZookeeperSaveResponse response = new ZookeeperSaveResponse();
        ZookeeperCurator curator = zookeeperService.curator(request.getUrl(), request.getAutomaticShutdown()  );
        if (StringUtils.isNoneEmpty(request.getPath())) {
            final String path = request.getPath();
            if (curator.checkExists(path)) {
                curator.setData(path, request.getData());
            } else {
                if (request.getModel() == null) {
                    curator.create(path, request.getData());
                } else {
                    curator.create(path, request.getData(), CreateMode.fromFlag(request.getModel()));
                }
            }
        }
        return response;
    }


    /**
     * @author 朱亮
     * @date 2021/08/30
     */
    @Data
    public static class ZookeeperSaveRequest extends ZookeeperRequestBase {
        /**
         * 路径
         */
        private String path;

        /**
         * 路径数据
         */
        private String data;

        /**
         * 模式
         */
        private Integer model;
    }

    /**
     * @author 朱亮
     * @date 2021/08/30
     */
    @Data
    public static class ZookeeperSaveResponse extends ZookeeperResponseBase {

    }
}
