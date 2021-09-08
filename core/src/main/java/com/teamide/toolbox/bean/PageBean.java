package com.teamide.toolbox.bean;

import java.util.List;

/**
 * 分页结果对象
 * <p>
 * 用于分页查询结果输出
 *
 * @author Model Coder
 */
public class PageBean<T> extends ResultBean<List<T>> {

    /**
     *
     */
    private static final long serialVersionUID = 1L;

    /**
     * 当前页
     */
    private int pageNum = 1;

    /**
     * 每页记录
     */
    private int pageSize = 10;

    /**
     * 总页数
     */
    private int pages = 0;

    /**
     * 总记录数
     */
    private long total = 0;

    /**
     * 上一页
     */
    private int prePage = 1;

    /**
     * 下一页
     */
    private int nextPage = 1;

    public int getPageNum() {
        return pageNum;
    }

    public void setPageNum(int pageNum) {
        this.pageNum = pageNum;
    }

    public int getPageSize() {
        return pageSize;
    }

    public void setPageSize(int pageSize) {
        this.pageSize = pageSize;
    }

    public int getPages() {
        return pages;
    }

    public void setPages(int pages) {
        this.pages = pages;
    }

    public long getTotal() {
        return total;
    }

    public void setTotal(long total) {
        this.total = total;
    }

    public int getPrePage() {
        return prePage;
    }

    public void setPrePage(int prePage) {
        this.prePage = prePage;
    }

    public int getNextPage() {
        return nextPage;
    }

    public void setNextPage(int nextPage) {
        this.nextPage = nextPage;
    }
}
