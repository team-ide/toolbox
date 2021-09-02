<template>
  <div class="body-main-wrap">
    <el-tabs v-model="activeKey" @tab-click="tabClick">
      <template v-for="(tab, index) in tabs">
        <el-tab-pane :key="index" :label="tab.name" :name="tab.key">
          <template v-if="tab.workerType == 'zookeeper'">
            <WorkerZookeeper :workerKey="tab.key"></WorkerZookeeper>
          </template>
        </el-tab-pane>
      </template>
    </el-tabs>
  </div>
</template>

<script>
import server from "@/server";
import tool from "@/tool";
import source from "@/source";

import WorkerZookeeper from "@/views/worker/WorkerZookeeper";
import { on } from "element-ui/lib/utils/dom";
export default {
  components: { WorkerZookeeper },
  data() {
    return {
      tool,
      source,
      tabs: [],
      activeKey: null,
    };
  },
  watch: {},
  computed: {
    treeStyleObject: function () {
      return {};
    },
  },
  methods: {
    onOpen(data) {
      let tabKey = this.getTabKey(data);
      let tab = this.getTab(tabKey);
      if (tab == null) {
        tab = this.createTab(data);
        this.addTab(tab);
      }
      this.$nextTick(() => {
        this.active(tab);
      });
    },
    active(tab) {
      this.activeKey = tab.key;
    },
    createTab(data) {
      let tab = {};
      let tabKey = this.getTabKey(data);
      tab.key = tabKey;
      tab.name = data.name;
      tab.title = data.title;
      tab.workerType = data.workerType;
      return tab;
    },
    getTabKey(data) {
      let key = data.key;

      return key;
    },
    addTab(tab) {
      this.tabs.push(tab);
    },
    getTab(key) {
      let res = null;
      this.tabs.forEach((one) => {
        if (one.key == key) {
          res = one;
        }
      });
      return res;
    },
    getTabIndex(key) {
      let tab = this.getTab(key);
      return this.tabs.indexOf(tab);
    },
    tabClick() {},
    init() {},
  },
  mounted() {
    this.init();
  },
};
</script>

<style>
.body-main-wrap {
  height: 100%;
  width: 100%;
  margin: 0px;
  padding: 0px;
  position: relative;
}
.body-main-wrap .el-tabs {
  height: 100%;
  width: 100%;
  position: relative;
}
.body-main-wrap .el-tabs__header {
  height: 30px;
  width: 100%;
  margin: 0px;
  padding: 0px;
  position: relative;
}
.body-main-wrap .el-tabs__header .el-tabs__active-bar {
  background-color: #6f6f6f;
}
.body-main-wrap .el-tabs__nav-wrap::after {
  height: 1px;
  background-color: transparent;
}
.body-main-wrap .el-tabs__header .el-tabs__item {
  height: 28px;
  line-height: 28px;
  padding: 0 10px;
  color: unset;
}
.body-main-wrap .el-tabs__content {
  height: calc(100% - 30px);
  width: 100%;
  margin: 0px;
  padding: 0px;
  position: relative;
}
.body-main-wrap .el-tabs__content .el-tab-pane {
  height: 100%;
  width: 100%;
  position: relative;
}
</style>
