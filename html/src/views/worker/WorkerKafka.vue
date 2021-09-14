<template>
  <div class="worker-kafka-wrap">
    <div class="worker-kafka-config pd-10 pdb-0">
      <el-form
        :inline="true"
        :model="configForm"
        size="mini"
        @submit.native.prevent
      >
        <el-form-item label="address（多个使用“,”隔开）">
          <el-input
            v-model="configForm.address"
            placeholder="address"
          ></el-input>
        </el-form-item>
        <el-form-item label="groupId（用于消费）">
          <el-input
            v-model="configForm.groupId"
            placeholder="groupId"
          ></el-input>
        </el-form-item>
        <el-form-item>
          <a class="tm-btn tm-btn-sm color-green" @click="doConnect"> 连接 </a>
        </el-form-item>
      </el-form>
    </div>
    <div
      class="worker-kafka-topic-list worker-scrollbar"
      v-if="connect.open"
      ref="treeBox"
    >
      <el-tree
        ref="topic_tree"
        :props="defaultProps"
        :data="topics"
        node-key="key"
        @node-click="topicNodeClick"
        @current-change="topicCurrentChange"
        :expand-on-click-node="false"
      >
        <span class="worker-box-tree-span" slot-scope="{ node, data }">
          <span>{{ node.label }}</span>
          <span class="mgl-20">
            <a class="tm-link color-orange ft-15 mgr-2" @click="toDelete(data)">
              <i class="mdi mdi-delete-outline"></i>
            </a>
          </span>
        </span>
      </el-tree>
    </div>
    <div
      class="worker-kafka-data-list worker-scrollbar"
      v-if="connect.open"
      ref="treeBox"
    >
      <el-tree
        ref="data_tree"
        :props="defaultProps"
        :data="datas"
        node-key="key"
        @node-click="dataNodeClick"
        @current-change="dataCurrentChange"
        :expand-on-click-node="false"
      >
        <span class="worker-box-tree-span" slot-scope="{ node, data }">
          <span>{{ node.label }}</span>
          <span class="mgl-20">
            <a class="tm-link color-orange ft-15 mgr-2" @click="toDelete(data)">
              <i class="mdi mdi-delete-outline"></i>
            </a>
          </span>
        </span>
      </el-tree>
    </div>
    <div class="worker-kafka-form worker-scrollbar" v-if="connect.open">
      <template v-if="readonlyOne">
        <h3>查看</h3>
      </template>
      <template v-else-if="insertOne">
        <h3>新增</h3>
      </template>
      <template v-else-if="updateOne">
        <h3>修改</h3>
      </template>
      <el-form :model="oneForm" size="lg" @submit.native.prevent>
        <el-form-item label="key">
          <el-input
            v-model="oneForm.key"
            placeholder="key"
            :readonly="readonlyOne || updateOne"
          ></el-input>
        </el-form-item>
        <el-form-item label="value">
          <el-input
            type="textarea"
            v-model="oneForm.value"
            placeholder="value"
            :autosize="{ minRows: 3, maxRows: 3 }"
            :readonly="readonlyOne"
          ></el-input>
        </el-form-item>
        <el-form-item label="格式化JSON（只做查看，不保存）">
          <el-input
            type="textarea"
            v-model="oneForm.json"
            placeholder="格式化JSON"
            :autosize="{ minRows: 6, maxRows: 6 }"
          ></el-input>
        </el-form-item>
        <el-form-item>
          <a v-if="!readonlyOne" class="tm-btn color-green" @click="push">
            推送
          </a>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
import server from "@/server";
import tool from "@/tool";
import source from "@/source";

export default {
  components: {},
  props: ["workerKey"],
  data() {
    return {
      tool,
      source,
      configForm: {
        address: "127.0.0.1:9092",
        groupId: "test-group",
      },
      connect: {
        open: false,
        form: null,
      },
      loading: false,
      searchForm: { pattern: "*", size: 20 },
      readonlyOne: true,
      insertOne: true,
      updateOne: true,
      oneForm: {
        key: null,
        value: null,
        json: null,
      },
      topics: null,
      datas: null,
      defaultProps: {
        children: "children",
        label: "name",
        isLeaf: "leaf",
      },
    };
  },
  watch: {
    "oneForm.value"(value) {
      this.oneForm.json = null;
      if (tool.isNotEmpty(value)) {
        try {
          if (
            (value.startsWith("{") && value.endsWith("}")) ||
            (value.startsWith("[") && value.endsWith("]"))
          ) {
            let data = JSON.parse(value);
            this.oneForm.json = JSON.stringify(data, null, "    ");
          }
        } catch (e) {
          this.oneForm.json = e;
        }
      }
    },
  },
  computed: {
    treeStyleObject: function () {
      return {};
    },
  },
  methods: {
    loadTopics(pattern, size) {
      let data = {};
      Object.assign(data, this.connect.form);
      return server.kafka.topics(data);
    },
    poll() {
      let data = {};
      Object.assign(data, this.connect.form);
      Object.assign(data, this.searchForm);
      return server.kafka.poll(data);
    },
    push() {
      let data = {};
      Object.assign(data, this.connect.form);
      Object.assign(data, this.oneForm);
      if (tool.isEmpty(data.value)) {
        tool.error("value不能为空！");
        return;
      }
      server.kafka.push(data).then((res) => {
        if (res.code != 0) {
          tool.error(res.msg);
        } else {
          tool.success("推送成功");
          this.reload();
        }
      });
    },
    toSubmit(one) {
      if (window.event) {
        window.event.stopPropagation && window.event.stopPropagation();
      }
      let data = {};
      Object.assign(data, this.connect.form);
      Object.assign(data, this.searchForm);
      data.topic = one.topic;
      if (tool.isEmpty(data.topic)) {
        tool.error("topic不能为空！");
        return;
      }
      data.partition = one.partition;
      if (tool.isEmpty(data.partition)) {
        tool.error("partition不能为空！");
        return;
      }
      data.offset = one.offset;
      if (tool.isEmpty(data.offset)) {
        tool.error("offset不能为空！");
        return;
      }
      data.partition = Number(data.partition);
      data.offset = Number(data.offset);
      tool
        .confirm(
          "将使用组[" +
            data.groupId +
            "]提交[" +
            data.topic +
            "]分区[" +
            data.partition +
            "]位置[" +
            data.offset +
            "]，确认提交？"
        )
        .then(() => {
          server.kafka.submit(data).then((res) => {
            if (res.code != 0) {
              tool.error(res.msg);
            } else {
              tool.success("提交成功");
              this.reloadData();
            }
          });
        })
        .catch(() => {});
    },
    topicNodeClick() {},
    dataNodeClick() {},
    doSearch() {
      this.loadKafka();
    },
    reload() {
      this.loadKafka();
    },
    loadKafka() {
      let data = {};
      Object.assign(data, this.searchForm);
      this.data = null;
      this.count = 0;
      this.loading = true;
      this.loadTopics()
        .then((res) => {
          this.loading = false;
          if (res.code != 0) {
            tool.error(res.msg);
          } else {
            let value = res.value || {};
            let topics = value.topics || [];
            let datas = [];

            topics.forEach((name) => {
              datas.push({ key: name, name: name });
            });
            this.topics = datas;
          }
        })
        .catch(() => {
          this.loading = false;
        });
    },
    doConnect() {
      this.connect.open = false;
      tool.trimList(this.expands);
      tool.trimList(this.opens);

      this.toInsert();

      this.$nextTick(() => {
        this.connect.form = Object.assign({}, this.configForm);
        this.connect.open = true;
        this.loadKafka();
        tool.setCache(this.getCacheKey(), JSON.stringify(this.connect.form));
      });
    },
    initConnect() {
      let value = tool.getCache(this.getCacheKey());
      if (tool.isNotEmpty(value)) {
        let data = JSON.parse(value);
        for (var key in this.configForm) {
          this.configForm[key] = data[key];
        }
        //this.doConnect();
      }
    },
    getCacheKey() {
      return "teamide-toolbox-" + this.workerKey;
    },
    toInsert() {
      if (window.event) {
        window.event.stopPropagation && window.event.stopPropagation();
      }
      this.oneForm.key = "";
      this.oneForm.value = "";
      this.updateOne = false;
      this.insertOne = true;
      this.readonlyOne = false;
    },
    topicCurrentChange(data) {
      this.toUpdate(data);
    },
    dataCurrentChange(data) {
      this.toUpdate(data);
    },
    toUpdate(data) {
      this.oneForm.key = data.key;
      this.insertOne = false;
      this.updateOne = false;
      this.readonlyOne = true;

      this.get(this.oneForm.key).then((res) => {
        if (res.code != 0) {
          tool.error(res.msg);
        } else {
          let value = res.value || {};
          this.oneForm.value = value.value;
          this.readonlyOne = false;
          this.updateOne = true;
        }
      });
    },
    toInfo(data) {
      this.oneForm.key = data.key;
      this.insertOne = false;
      this.readonlyOne = true;

      this.get(this.oneForm.key).then((res) => {
        if (res.code != 0) {
          tool.error(res.msg);
        } else {
          let value = res.value || {};
          this.oneForm.value = value.value;
        }
      });
    },
    init() {
      this.initConnect();
    },
  },
  mounted() {
    this.init();
  },
};
</script>

<style>
.worker-kafka-wrap {
  height: 100%;
  width: 100%;
  margin: 0px;
  padding: 0px;
  position: relative;
}
.worker-kafka-wrap .worker-kafka-topic-list {
  height: calc(100% - 150px);
  width: 300px;
  margin: 0px;
  padding: 0px;
  position: relative;
  float: left;
  overflow-x: hidden !important;
}
.worker-kafka-wrap .worker-kafka-data-list {
  height: calc(100% - 150px);
  width: 400px;
  margin: 0px;
  padding: 0px;
  position: relative;
  float: left;
  overflow-x: hidden !important;
}
.worker-kafka-wrap .worker-kafka-form {
  height: calc(100% - 100px);
  width: calc(100% - 760px);
  margin: 0px;
  padding: 0px;
  position: relative;
  float: left;
  padding: 10px;
  overflow: auto;
}
.worker-kafka-wrap .el-tree {
  /* border: 1px solid #f3f3f3; */
  border-bottom: 0px;
}
.worker-kafka-wrap .el-tree-node__content {
  border-bottom: 1px dotted #696969;
}
</style>
