package com.teamide.toolbox.exception;


import com.teamide.toolbox.component.CodeMsg;
import org.apache.commons.lang3.StringUtils;

/**
 * 基础异常
 * <p>
 * 异常包括错误码、异常信息等
 *
 * @author Model Coder
 */
public class BaseException extends Exception {

    /**
     *
     */
    private static final long serialVersionUID = 5851650165956394577L;

    private final String code;

    private final String msg;

    public BaseException(CodeMsg codeMsg) {
        this(codeMsg.getCode(), codeMsg.getMsg());
    }

    public BaseException(CodeMsg codeMsg, String msg) {
        this(codeMsg.getCode(), msg);
    }

    public BaseException(Throwable cause) {
        this(null, cause.getMessage(), cause);
    }

    public BaseException(String code, String msg) {
        this(code, msg, null);
    }

    public BaseException(String code, String msg, Throwable cause) {
        super(msg, cause);
        if (StringUtils.isEmpty(code)) {
            code = CodeMsg.FAIL.getCode();
        }
        this.code = code;
        this.msg = msg;
    }

    public String getCode() {
        return code;
    }

    public String getMsg() {
        return msg;
    }

}
