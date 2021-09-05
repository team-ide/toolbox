/* Model Coder File Mark */
/**
 * The current file is generated by Model Coder, please do not modify and move!
 **/
package com.teamide.toolbox.bean;


import com.teamide.toolbox.component.Errcode;

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
    public static final ResponseBean success() {

        return new ResponseBean(Errcode.SUCCESS);
    }

    /**
     * 返回成功
     */
    public static final ResponseBean success(Object value) {

        ResponseBean res = new ResponseBean(Errcode.SUCCESS);
        res.setValue(value);
        return res;
    }

    /**
     * 返回失败
     */
    public static final ResponseBean fail() {

        return new ResponseBean(Errcode.FAIL);
    }

    /**
     * 返回失败
     */
    public static final ResponseBean fail(String msg) {

        return new ResponseBean(Errcode.FAIL.getCode(), msg);
    }

    public ResponseBean() {

    }

    public ResponseBean(Errcode errcode) {
        this(errcode.getCode(), errcode.getMsg());
    }

    public ResponseBean(String code, String msg) {
        this.code = code;
        this.msg = msg;
    }

    public ResponseBean(int code, String msg) {
        this(Integer.toString(code), msg);
    }

}