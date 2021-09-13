package com.teamide.toolbox.redis.worker;

import com.teamide.toolbox.redis.service.RedisDo;
import com.teamide.toolbox.redis.service.RedisService;
import com.teamide.toolbox.worker.ToolboxWork;
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
public class RedisWorkDelete implements ToolboxWork<RedisWorkDelete.Request, RedisWorkDelete.Response> {

    @Autowired
    RedisService redisService;

    @Override
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
        RedisDo redis = redisService.redis(request.getAddress(), request.getAuth(), request.isCluster(), request.getAutomaticShutdown());
        if (StringUtils.isNoneEmpty(request.getKey())) {
            int count = redis.delete(request.getKey());
            response.setCount(count);
        }
        return response;
    }


    /**
     * @author 朱亮
     * @date 2021/08/30
     */
    @EqualsAndHashCode(callSuper = true)
    @Data
    public static class Request extends RedisRequestBase {
        /**
         * key
         */
        private String key;


    }

    /**
     * @author 朱亮
     * @date 2021/08/30
     */
    @EqualsAndHashCode(callSuper = true)
    @Data
    public static class Response extends RedisResponseBase {
        /**
         * count
         */
        private int count;
    }
}
