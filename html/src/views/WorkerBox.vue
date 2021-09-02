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
        color: "#ffffff",
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
</style>
