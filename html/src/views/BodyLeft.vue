<template>
  <div class="body-left-wrap">
    <div class="worker-tree" :style="treeStyleObject">
      <el-tree
        :data="data"
        :props="defaultProps"
        :default-expanded-keys="expands"
        :expand-on-click-node="false"
        node-key="key"
        @node-click="nodeClick"
      ></el-tree>
    </div>
  </div>
</template>

<script>
import server from "@/server";
import tool from "@/tool";
import source from "@/source";

export default {
  components: {},
  data() {
    return {
      tool,
      source,
      expands: [],
      opens: [],
      data: [
        {
          name: "Redis",
          key: "redis",
          children: [
            {
              name: "default",
              title: "Redis",
              workerType: "redis",
              key: "redis/default",
            },
          ],
        },
        {
          name: "Zookeeper",
          key: "zookeeper",
          children: [
            {
              name: "default",
              title: "Zookeeper",
              workerType: "zookeeper",
              key: "zookeeper/default",
            },
          ],
        },
        {
          name: "Database",
          key: "database",
          children: [
            {
              name: "default",
              title: "Database",
              workerType: "database",
              key: "database/default",
            },
          ],
        },
        {
          name: "Elasticsearch",
          key: "elasticsearch",
          children: [
            {
              name: "default",
              title: "Elasticsearch",
              workerType: "elasticsearch",
              key: "elasticsearch/default",
            },
          ],
        },
        {
          name: "Mongo",
          key: "mongo",
          children: [
            {
              name: "default",
              title: "Mongo",
              workerType: "mongo",
              key: "mongo/default",
            },
          ],
        },
        {
          name: "Kafka",
          key: "kafka",
          children: [
            {
              name: "default",
              title: "Kafka Default",
              workerType: "kafka",
              key: "kafka/default",
            },
          ],
        },
        {
          name: "Rabbit",
          key: "rabbit",
          children: [
            {
              name: "default",
              title: "Rabbit Default",
              workerType: "rabbit",
              key: "rabbit/default",
            },
          ],
        },
      ],
      defaultProps: {
        children: "children",
        label: "name",
      },
    };
  },
  watch: {},
  computed: {
    treeStyleObject: function () {
      return {};
    },
  },
  methods: {
    nodeClick(data, node, nodeView) {
      let nowTime = new Date().getTime();
      let clickTime = node.clickTime;
      node.clickTime = nowTime;
      if (clickTime) {
        let timeout = nowTime - clickTime;
        if (timeout < 300) {
          node.clickTime = null;
          this.nodeDbclick(data, node, nodeView);
        }
      } else {
      }
    },
    nodeDbclick(data, node, nodeView) {
      if (data.children && data.children.length > 0) {
        let key = data.key;
        let index = this.expands.indexOf(key);
        if (index < 0) {
          this.expands.push(key);
        } else {
          this.expands.splice(index, 1);
        }
        return;
      }
      this.open(data);
    },
    initExpands() {
      this.expand("redis");
      this.expand("zookeeper");
      this.expand("kafka");
    },
    expand(key) {
      if (this.expands.indexOf(key) < 0) {
        this.expands.push(key);
      }
    },
    open(data) {
      if (this.opens.indexOf(data.key) < 0) {
        this.opens.push(data.key);
      }
      this.$emit("open", data);
    },
    initOpens() {},
    init() {
      this.initExpands();
    },
  },
  mounted() {
    this.init();
  },
};
</script>

<style>
.body-left-wrap {
  height: 100%;
  width: 100%;
  margin: 0px;
  padding: 0px;
  position: relative;
}
.worker-tree {
  height: 100%;
  width: 100%;
  margin: 0px;
  padding: 0px;
  position: relative;
}
</style>
