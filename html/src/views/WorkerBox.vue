<template>
  <div class="worker-box" :style="boxStyleObject">
    <tm-layout height="100%">
      <tm-layout :height="style.header.height">
        <div class="worker-header" :style="headerStyleObject">
          <h2 class="mg-0 pd-0 pdlr-20" style="line-height: 40px;">Team IDE Toolbox</h2>
        </div>
      </tm-layout>
      <tm-layout-bar bottom></tm-layout-bar>
      <tm-layout height="auto">
        <tm-layout height="100%">
          <tm-layout :width="style.body.left.width">
            <BodyLeft ref="BodyLeft" :style="style" @open="open"></BodyLeft>
          </tm-layout>
          <tm-layout-bar right></tm-layout-bar>
          <tm-layout width="auto">
            <BodyMain ref="BodyMain" :style="style"></BodyMain>
          </tm-layout>
          <!-- <tm-layout-bar left></tm-layout-bar> -->
          <tm-layout :width="style.body.right.width">
            <BodyRight ref="BodyRight" :style="style"></BodyRight>
          </tm-layout>
        </tm-layout>
      </tm-layout>
    </tm-layout>
  </div>
</template>

<script>
import server from "@/server";
import tool from "@/tool";
import source from "@/source";

import BodyLeft from "@/views/BodyLeft";
import BodyMain from "@/views/BodyMain";
import BodyRight from "@/views/BodyRight";

export default {
  components: { BodyLeft, BodyMain, BodyRight },
  data() {
    return {
      tool,
      source,
      style: {
        backgroundColor: "#2d2d2d",
        color: "#adadad",
        header: {
          height: "40px",
        },
        body: {
          left: {
            width: "260px",
          },
          main: {},
          right: {
            width: "0px",
          },
        },
        footer: {
          height: "30px",
        },
      },
    };
  },
  watch: {},
  computed: {
    boxStyleObject: function () {
      return {
        backgroundColor: this.style.backgroundColor,
        color: this.style.color,
      };
    },
    headerStyleObject: function () {
      return {};
    },
    bodyStyleObject: function () {
      return {};
    },
    footerStyleObject: function () {
      return {};
    },
    bodyLeftStyleObject: function () {
      return {};
    },
    bodyMainStyleObject: function () {
      return {};
    },
    bodyRightStyleObject: function () {
      return {};
    },
  },
  methods: {
    open(data) {
      this.$refs.BodyMain.onOpen(data);
    },
    init() {},
  },
  mounted() {
    this.init();
  },
};
</script>

<style>
.worker-box {
  height: 100%;
  width: 100%;
  margin: 0px;
  padding: 0px;
  position: relative;
}
.worker-box .worker-header {
  width: 100%;
  height: 100%;
  margin: 0px;
  padding: 0px;
  position: relative;
}
.worker-box .worker-body {
  width: 100%;
  height: 100%;
  margin: 0px;
  padding: 0px;
  position: relative;
}
.worker-box .worker-footer {
  width: 100%;
  height: 100%;
  margin: 0px;
  padding: 0px;
  position: relative;
}
.worker-box .worker-panel-title {
  font-size: 13px;
  overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
  color: #e0e0e0;
  line-height: 40px;
}
.worker-box .tm-layout {
  overflow: hidden;
}
.worker-box .tm-layout > .tm-layout-bar {
  background-color: #4e4e4e;
}
.worker-box .tm-layout-bar > .tm-layout-bar-part {
  background-color: #4e4e4e;
}
.worker-box .el-tree {
  background-color: transparent;
  color: unset;
  width: auto;
  font-size: 12px;
}
.worker-box .el-tree .mdi {
  vertical-align: middle;
}

.worker-box .worker-box-tree-span {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-size: 14px;
  padding-right: 8px;
}
.worker-box .el-tree__empty-block {
  display: none;
}
.worker-box .el-tree-node__children {
  overflow: visible !important;
}
.worker-box .el-tree-node__children.v-enter-active,
.worker-box .el-tree-node__children.v-leave-active {
  overflow: hidden !important;
}
.worker-box .el-tree .el-tree-node.is-current > .el-tree-node__content {
  background-color: #636363 !important;
}
.worker-box .el-tree .el-tree-node:focus > .el-tree-node__content {
  background-color: #545454;
}
.worker-box .el-tree .el-tree-node > .el-tree-node__content:hover {
  background-color: #545454;
}
.worker-box .el-form-item__label {
  color: #adadad;
}
.worker-box .el-input__inner,
.worker-box .el-textarea__inner,
.worker-box .el-select__inner {
  background-color: transparent;
  color: #adadad;
  border-style: dotted;
}
.worker-box input::-webkit-input-placeholder {
  color: #575757;
}
.worker-box input::-moz-input-placeholder {
  color: #575757;
}
.worker-box input::-ms-input-placeholder {
  color: #575757;
}
.worker-box textarea::-webkit-input-placeholder {
  color: #575757;
}
.worker-box textarea::-moz-input-placeholder {
  color: #575757;
}
.worker-box textarea::-ms-input-placeholder {
  color: #575757;
}
/* 滚动条样式*/
.worker-box .worker-scrollbar {
  overflow: scroll;
}
.worker-box textarea {
  overflow-x: hidden;
}
.worker-box .worker-scrollbar::-webkit-scrollbar,
.worker-box textarea::-webkit-scrollbar {
  width: 10px;
  height: 10px;
}
.worker-box .worker-scrollbar:hover::-webkit-scrollbar,
.worker-box textarea:hover::-webkit-scrollbar {
  width: 10px;
  height: 10px;
}
.worker-box .worker-scrollbar::-webkit-scrollbar-thumb,
.worker-box textarea::-webkit-scrollbar-thumb {
  border-radius: 0px;
}
.worker-box .worker-scrollbar::-webkit-scrollbar-track,
.worker-box textarea::-webkit-scrollbar-track {
  border-radius: 0;
}
.worker-box .worker-scrollbar:hover::-webkit-scrollbar-thumb,
.worker-box textarea:hover::-webkit-scrollbar-thumb {
  box-shadow: inset 0 0 10px #333333;
  background: #333333;
}
.worker-box .worker-scrollbar:hover::-webkit-scrollbar-track,
.worker-box textarea:hover::-webkit-scrollbar-track {
  box-shadow: inset 0 0 10px #262626;
  background: #262626;
}

.worker-box .el-table,
.worker-box .el-table__expanded-cell {
  background: transparent;
}
.worker-box .el-table th.el-table__cell {
  background: transparent;
}
.worker-box .el-table td.el-table__cell,
.worker-box .el-table th.el-table__cell.is-leaf {
  border-bottom: 1px solid #473939;
}
.worker-box .el-table th,
.worker-box .el-table tr {
  background: transparent;
}
.worker-box .el-table .el-table__row:hover td.el-table__cell {
  background-color: #473939;
}
.worker-box .el-table {
  color: unset;
}
.worker-box .el-table--mini td,
.worker-box .el-table--mini th {
  padding: 3px 0;
}
.worker-box .el-table--border::after,
.worker-box .el-table--group::after,
.worker-box .el-table::before {
  background-color: transparent;
}
.worker-box .el-table td,
.worker-box .el-table th.is-leaf {
  border-bottom: 1px solid #404040;
}
.worker-box .el-divider {
  background-color: #404040;
}
.worker-box .el-loading-mask{
    background-color: rgba(255,255,255,.1);
}
</style>
