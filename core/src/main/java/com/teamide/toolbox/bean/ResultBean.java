package com.teamide.toolbox.bean;

import java.io.Serializable;

/**
 * 结果对象
 * <p>
 * 输出code状态码、msg信息、value值
 *
 * @author Model Coder
 */
public class ResultBean<T> implements Serializable {

    /**
     *
     */
    private static final long serialVersionUID = 1L;

    /**
     * 结果 值
     */
    private T value;

    /**
     * 结果 状态码
     */
    protected String code = "0";

    /**
     * 结果 信息
     */
    protected String msg;

    public T getValue() {

        return value;
    }

    public void setValue(T value) {

        this.value = value;
    }

    public String getCode() {

        return code;
    }

    public void setCode(int code) {

        this.setCode(Integer.toString(code));
    }

    public void setCode(String code) {

        this.code = code;
    }

    public String getMsg() {

        return msg;
    }

    public void setMsg(String msg) {

        this.msg = msg;
    }

}
