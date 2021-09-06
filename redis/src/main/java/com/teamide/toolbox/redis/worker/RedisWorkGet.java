package com.teamide.toolbox.redis.worker;

import com.teamide.toolbox.redis.service.RedisDo;
import com.teamide.toolbox.redis.service.RedisService;
import com.teamide.toolbox.worker.ToolboxWork;
import lombok.Data;
import org.apache.commons.lang3.StringUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

/**
 * @author 朱亮
 * @date 2021/08/30
 */
@Service
public class RedisWorkGet implements ToolboxWork<RedisWorkGet.RedisGetRequest, RedisWorkGet.RedisGetResponse> {

    @Autowired
    private RedisService redisService;

    public Class<RedisGetRequest> getRequestClass() {
        return RedisGetRequest.class;
    }

    /**
     * 工作
     *
     * @param request 请求
     * @return {@link RedisGetResponse}
     * @throws Exception 异常
     */
    @Override
    public RedisGetResponse work(RedisGetRequest request) throws Exception {
        RedisGetResponse response = new RedisGetResponse();
        RedisDo redis = redisService.redis(request.getAddress(), request.getAuth(), request.isCluster(), request.getAutomaticShutdown());
        if (StringUtils.isNoneEmpty(request.getKey())) {
            String value = redis.get(request.getKey());
            response.setValue(value);
        }
        return response;
    }


    /**
     * @author 朱亮
     * @date 2021/08/30
     */
    @Data
    public static class RedisGetRequest extends RedisRequestBase {
        /**
         * key
         */
        private String key;


    }

    /**
     * @author 朱亮
     * @date 2021/08/30
     */
    @Data
    public static class RedisGetResponse extends RedisResponseBase {
        /**
         * value
         */
        private String value;

    }
}
