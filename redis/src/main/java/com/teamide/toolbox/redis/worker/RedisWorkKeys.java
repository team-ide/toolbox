package com.teamide.toolbox.redis.worker;

import com.teamide.toolbox.redis.service.RedisJedis;
import com.teamide.toolbox.redis.service.RedisService;
import com.teamide.toolbox.worker.ToolboxWork;
import lombok.Data;
import org.apache.commons.lang3.StringUtils;
import org.springframework.beans.factory.annotation.Autowired;
import org.springframework.stereotype.Service;

import java.util.Set;

/**
 * @author 朱亮
 * @date 2021/08/30
 */
@Service
public class RedisWorkKeys implements ToolboxWork<RedisWorkKeys.RedisKeysRequest, RedisWorkKeys.RedisKeysResponse> {

    @Autowired
    private RedisService redisService;

    public Class<RedisKeysRequest> getRequestClass() {
        return RedisKeysRequest.class;
    }

    /**
     * 工作
     *
     * @param request 请求
     * @return {@link RedisKeysResponse}
     * @throws Exception 异常
     */
    @Override
    public RedisKeysResponse work(RedisKeysRequest request) throws Exception {
        RedisKeysResponse response = new RedisKeysResponse();
        RedisJedis redis = redisService.redis(request.getHost(), request.getPort(), request.getAuth(), request.getAutomaticShutdown());
        if (StringUtils.isNoneEmpty(request.getPattern())) {
            Set<String> keys = redis.keys(request.getPattern());
            response.setKeys(keys);
        }
        return response;
    }


    /**
     * @author 朱亮
     * @date 2021/08/30
     */
    @Data
    public static class RedisKeysRequest extends RedisRequestBase {
        /**
         * key
         */
        private String pattern;


    }

    /**
     * @author 朱亮
     * @date 2021/08/30
     */
    @Data
    public static class RedisKeysResponse extends RedisResponseBase {
        /**
         * keys
         */
        private Set<String> keys;

    }
}
