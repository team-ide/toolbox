<template>
  <div class="worker-box" :style="boxStyleObject">
    <tm-layout height="100%">
      <tm-layout :height="style.header.height">
        <div class="worker-header" :style="headerStyleObject">
          <h2 class="mg-0 pd-0 pdlr-20" style="line-height: 40px">
            Team IDE Toolbox
            <a
              class="tm-link color-green ft-12"
              href="https://github.com/team-ide/toolbox"
              target="_blank"
              >https://github.com/team-ide/toolbox</a
            >
          </h2>
          <div
            style="float: right; margin-top: -37px; margin-right: 10px"
            v-if="source.login_user != null"
          >
            <el-dropdown
              @command="checkUserBtn"
              trigger="click"
              class="color-grey-3"
            >
              <div class="el-dropdown-link">
                {{ source.login_user.name }}
                <i class="el-icon-arrow-down el-icon--right"></i>
              </div>
              <el-dropdown-menu slot="dropdown">
                <el-dropdown-item command="logout">退出</el-dropdown-item>
              </el-dropdown-menu>
            </el-dropdown>
          </div>
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
    checkUserBtn(command) {
      if (command == "logout") {
        tool.doLogout();
      }
    },
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
.worker-box tr .worker-btn-group {
  transform: scale(0);
}
.worker-box tr:hover .worker-btn-group {
  transform: scale(1);
}

.worker-box .el-tree-node__content {
  position: relative;
}

.worker-box .el-tree-node__content .worker-btn-group {
  display: none;
}
.worker-box .el-tree-node__content:hover .worker-btn-group {
  display: block;
}
.worker-box .worker-box-tree-span .worker-btn-group {
  position: absolute;
  right: 5px;
  top: 0px;
}
.worker-box .worker-box-tree-span {
  /* flex: 1;
  display: flex;
  align-items: center;
  justify-content: space-between; */
  font-size: 14px;
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
.worker-box .el-table--border,
.worker-box .el-table--group {
  border: 1px solid #404040;
}
.worker-box .el-table td,
.worker-box .el-table th.is-leaf {
  border-bottom: 1px solid #404040;
}
.worker-box .el-table--border .el-table__cell,
.worker-box
  .el-table__body-wrapper
  .el-table--border.is-scrolling-left
  ~ .el-table__fixed {
  border-right: 1px solid #404040;
}
.worker-box .el-table--border th.el-table__cell.gutter:last-of-type {
  border-bottom: 1px solid #404040;
}
.worker-box .el-divider {
  background-color: #404040;
}
.worker-box .el-loading-mask {
  background-color: rgba(255, 255, 255, 0.1);
}

.part-box {
  line-height: 20px;
  font-size: 12px;
  overflow: auto;
  width: 100%;
  height: 100%;
}
.part-box,
.part-box li {
  padding: 0px;
  margin: 0px;
  list-style: none;
}
.part-box li {
  text-overflow: ellipsis;
  white-space: nowrap;
  word-break: keep-all;
}

.part-box input,
.part-box select {
  color: #ffffff;
  width: 40px;
  min-width: 40px;
  border: 1px dashed transparent;
  background-color: transparent;
  height: 20px;
  max-width: 100%;
  padding: 0px;
  padding-left: 2px;
  padding-right: 2px;
  box-sizing: border-box;
  outline: none;
  font-size: 12px;
}

.part-box input {
  border-bottom: 1px dashed #636363;
}
.part-box select {
  -moz-appearance: auto;
  -webkit-appearance: auto;
}
.part-box option {
  background-color: #ffffff;
  color: #3e3e3e;
}
.part-box input[type="checkbox"] {
  width: 10px;
  min-width: 10px;
  height: 13px;
  vertical-align: bottom;
  margin-left: 6px;
}

.part-box textarea {
  color: #ffffff;
  height: 70px;
  border: 1px dashed #636363;
  text-align: left;
  padding: 5px;
  min-width: 500px;
  background-color: transparent;
  font-size: 12px;
  vertical-align: text-top;
}
</style>
