package com.teamide.toolbox.redis.worker;

import com.teamide.toolbox.redis.service.RedisJedis;
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
public class RedisWorkSave implements ToolboxWork<RedisWorkSave.RedisSaveRequest, RedisWorkSave.RedisSaveResponse> {

    @Autowired
    private RedisService redisService;

    public Class<RedisSaveRequest> getRequestClass() {
        return RedisSaveRequest.class;
    }

    /**
     * 工作
     *
     * @param request 请求
     * @return {@link RedisSaveResponse}
     * @throws Exception 异常
     */
    @Override
    public RedisSaveResponse work(RedisSaveRequest request) throws Exception {
        RedisSaveResponse response = new RedisSaveResponse();
        RedisJedis redis = redisService.redis(request.getHost(), request.getPort(), request.getAuth(), request.getAutomaticShutdown());
        if (StringUtils.isNoneEmpty(request.getKey())) {
            redis.set(request.getKey(), request.getValue());
        }
        return response;
    }


    /**
     * @author 朱亮
     * @date 2021/08/30
     */
    @Data
    public static class RedisSaveRequest extends RedisRequestBase {
        /**
         * key
         */
        private String key;

        /**
         * value
         */
        private String value;

    }

    /**
     * @author 朱亮
     * @date 2021/08/30
     */
    @Data
    public static class RedisSaveResponse extends RedisResponseBase {

    }
}
