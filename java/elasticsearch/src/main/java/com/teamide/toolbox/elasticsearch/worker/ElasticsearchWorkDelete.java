package com.teamide.toolbox.elasticsearch.worker;

import com.teamide.toolbox.elasticsearch.service.ElasticsearchCurator;
import com.teamide.toolbox.elasticsearch.service.ElasticsearchService;
import com.teamide.toolbox.worker.ToolboxWork;
import lombok.Data;
import lombok.EqualsAndHashCode;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

/**
 * @author 朱亮
 * @date 2021/08/30
 */
@Service
public class ElasticsearchWorkDelete implements ToolboxWork<ElasticsearchWorkDelete.Request, ElasticsearchWorkDelete.Response> {

    @Autowired
    ElasticsearchService elasticsearchService;

    @Override
    public Class<Request> getRequestClass() {
        return Request.class;
    }

    /**
     * 工作
     *
     * @param request 请求
     * @return {@link Request}
     * @throws Exception 异常
     */
    @Override
    public Response work(Request request) throws Exception {
        Response response = new Response();
        ElasticsearchCurator curator = elasticsearchService.curator(request.getUrl(), request.getAutomaticShutdown());
        curator.work();
        return response;
    }


    /**
     * @author 朱亮
     * @date 2021/08/30
     */
    @EqualsAndHashCode(callSuper = true)
    @Data
    public static class Request extends ElasticsearchRequestBase {

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
    public static class Response extends ElasticsearchResponseBase {

    }
}
