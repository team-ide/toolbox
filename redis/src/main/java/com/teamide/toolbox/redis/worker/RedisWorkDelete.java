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
public class RedisWorkDelete implements ToolboxWork<RedisWorkDelete.RedisDeleteRequest, RedisWorkDelete.RedisDeleteResponse> {

    @Autowired
    private RedisService redisService;

    public Class<RedisDeleteRequest> getRequestClass() {
        return RedisDeleteRequest.class;
    }

    /**
     * 工作
     *
     * @param request 请求
     * @return {@link RedisDeleteResponse}
     * @throws Exception 异常
     */
    @Override
    public RedisDeleteResponse work(RedisDeleteRequest request) throws Exception {
        RedisDeleteResponse response = new RedisDeleteResponse();
        RedisDo redis = redisService.redis(request.getAddress(), request.getAuth(), request.isCluster(), request.getAutomaticShutdown());
        if (StringUtils.isNoneEmpty(request.getKey())) {
            redis.delete(request.getKey());
        }
        return response;
    }


    /**
     * @author 朱亮
     * @date 2021/08/30
     */
    @Data
    public static class RedisDeleteRequest extends RedisRequestBase {
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
    public static class RedisDeleteResponse extends RedisResponseBase {

    }
}
