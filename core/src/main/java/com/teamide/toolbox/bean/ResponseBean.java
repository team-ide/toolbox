package com.teamide.toolbox.bean;


import com.teamide.toolbox.component.CodeMsg;

/**
 * 响应对象
 * <p>
 * 输出code状态码、msg信息、value值
 *
 * @author Model Coder
 */
public class ResponseBean extends ResultBean<Object> {

    /**
     *
     */
    private static final long serialVersionUID = 1L;

    /**
     * 返回成功
     */
    public static ResponseBean success() {

        return new ResponseBean(CodeMsg.SUCCESS);
    }

    /**
     * 返回成功
     */
    public static ResponseBean success(Object value) {

        ResponseBean res = new ResponseBean(CodeMsg.SUCCESS);
        res.setValue(value);
        return res;
    }

    /**
     * 返回失败
     */
    public static ResponseBean fail() {

        return new ResponseBean(CodeMsg.FAIL);
    }

    /**
     * 返回失败
     */
    public static ResponseBean fail(String msg) {

        return new ResponseBean(CodeMsg.FAIL.getCode(), msg);
    }

    public ResponseBean() {

    }

    public ResponseBean(CodeMsg codeMsg) {
        this(codeMsg.getCode(), codeMsg.getMsg());
    }

    public ResponseBean(String code, String msg) {
        this.code = code;
        this.msg = msg;
    }

    public ResponseBean(int code, String msg) {
        this(Integer.toString(code), msg);
    }

}
