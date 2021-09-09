package com.teamide.toolbox.component;

/**
 * CodeMsg
 * <p>
 * 应用错误码
 *
 * @author Model Coder
 */
public class CodeMsg {

    /**
     * 成功
     **/
    public static final CodeMsg SUCCESS = new CodeMsg("0", "成功");

    /**
     * 失败
     **/
    public static final CodeMsg FAIL = new CodeMsg("-1", "失败");

    private final String code;

    private final String msg;

    public CodeMsg(String code, String msg) {
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
