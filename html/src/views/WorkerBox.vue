<template>
  <div class="worker-box" :style="boxStyleObject">
    <div class="worker-header" :style="headerStyleObject"></div>
    <div class="worker-body" :style="bodyStyleObject">
      <div class="worker-left" :style="bodyLeftStyleObject">
        <BodyLeft ref="BodyLeft" :style="style" @open="open"></BodyLeft>
      </div>
      <div class="worker-main" :style="bodyMainStyleObject">
        <BodyMain ref="BodyMain" :style="style"></BodyMain>
      </div>
      <div class="worker-right" :style="bodyRightStyleObject">
        <BodyRight ref="BodyRight" :style="style"></BodyRight>
      </div>
    </div>
    <div class="worker-footer" :style="footerStyleObject"></div>
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
          height: "29px",
          borderBottom: "0px solid #6f6f6f",
          borderSize: "1px",
        },
        body: {
          left: {
            width: "259px",
            borderRight: "0px solid #6f6f6f",
            borderSize: "1px",
          },
          main: {},
          right: {
            width: "0px",
            borderLeft: "0px solid #6f6f6f",
            borderSize: "0px",
          },
        },
        footer: {
          height: "29px",
          borderTop: "0px solid #6f6f6f",
          borderSize: "1px",
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
      return {
        height: this.style.header.height,
        borderBottom: this.style.header.borderBottom,
        borderWidth: this.style.header.borderSize,
      };
    },
    bodyStyleObject: function () {
      return {
        height:
          "calc(100% - " +
          this.style.header.height +
          " - " +
          this.style.header.borderSize +
          " - " +
          this.style.footer.height +
          " - " +
          this.style.footer.borderSize +
          ")",
      };
    },
    footerStyleObject: function () {
      return {
        height: this.style.footer.height,
        borderTop: this.style.footer.borderTop,
        borderWidth: this.style.footer.borderSize,
      };
    },
    bodyLeftStyleObject: function () {
      return {
        width: this.style.body.left.width,
        borderRight: this.style.body.left.borderRight,
        borderWidth: this.style.body.left.borderSize,
      };
    },
    bodyMainStyleObject: function () {
      return {
        width:
          "calc(100% - " +
          this.style.body.left.width +
          " - " +
          this.style.body.left.borderSize +
          " - " +
          this.style.body.right.width +
          " - " +
          this.style.body.right.borderSize +
          ")",
      };
    },
    bodyRightStyleObject: function () {
      return {
        width: this.style.body.right.width,
        borderLeft: this.style.body.right.borderLeft,
        borderWidth: this.style.body.right.borderSize,
      };
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
  margin: 0px;
  padding: 0px;
  position: relative;
}
.worker-box .worker-body {
  width: 100%;
  margin: 0px;
  padding: 0px;
  position: relative;
}
.worker-box .worker-footer {
  width: 100%;
  margin: 0px;
  padding: 0px;
  position: relative;
}
.worker-box .worker-body .worker-left {
  height: 100%;
  margin: 0px;
  padding: 0px;
  float: left;
  position: relative;
}
.worker-box .worker-body .worker-main {
  height: 100%;
  margin: 0px;
  padding: 0px;
  float: left;
  position: relative;
}
.worker-box .worker-body .worker-right {
  height: 100%;
  margin: 0px;
  padding: 0px;
  float: right;
  position: relative;
}
.worker-box .el-tree {
  background-color: transparent;
  color: unset;
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
</style>
