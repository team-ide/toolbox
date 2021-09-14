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
        <el-form-item>
          <a class="tm-btn tm-btn-sm color-green" @click="doConnect"> 连接 </a>
        </el-form-item>
      </el-form>
      <el-divider class="mg-0"></el-divider>
    </div>
    <div class="worker-kafka-topic-list-box" v-if="connect.open">
      <div class="ft-16 pdb-15 color-orange">Topics列表</div>
      <el-divider class="mgb-5 mg-0"></el-divider>
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
          <span class="worker-box-tree-span" slot-scope="{ data }">
            <span>{{ data.topic }}</span>
            <span class="mgl-20"> </span>
          </span>
        </el-tree>
      </div>
    </div>
    <div class="worker-kafka-data-list-box" v-if="connect.open">
      <el-form
        :inline="true"
        :model="searchForm"
        size="mini"
        @submit.native.prevent
      >
        <el-form-item label="topic">
          <el-input
            v-model="searchForm.topic"
            placeholder="topic"
            style="width: 120px"
          ></el-input>
        </el-form-item>
        <el-form-item label="groupId">
          <el-input
            v-model="searchForm.groupId"
            placeholder="groupId"
            style="width: 120px"
          ></el-input>
        </el-form-item>
        <el-form-item label="key类型">
          <el-select
            v-model="searchForm.keyType"
            placeholder="keyType"
            style="width: 80px"
          >
            <el-option label="String" value="String"> </el-option>
            <el-option label="Long" value="Long"> </el-option>
            <el-option label="Integer" value="Integer"> </el-option>
          </el-select>
        </el-form-item>
        <el-form-item label="value类型">
          <el-select
            v-model="searchForm.valueType"
            placeholder="valueType"
            style="width: 80px"
          >
            <el-option label="String" value="String"> </el-option>
            <el-option label="Long" value="Long"> </el-option>
            <el-option label="Integer" value="Integer"> </el-option>
          </el-select>
        </el-form-item>
        <el-form-item>
          <a class="tm-btn tm-btn-sm color-green mgl-5" @click="loadDatas">
            拉取
          </a>
          <a class="tm-btn tm-btn-sm color-blue mgl-5" @click="toInsert">
            新增
          </a>
        </el-form-item>
      </el-form>
      <div class="ft-16 pdb-15 color-orange">Msgs列表</div>
      <el-divider class="mgb-5 mg-0"></el-divider>
      <div class="worker-kafka-data-list worker-scrollbar" ref="treeBox">
        <el-tree
          ref="data_tree"
          :props="defaultProps"
          :data="datas"
          node-key="key"
          @node-click="dataNodeClick"
          @current-change="dataCurrentChange"
          :expand-on-click-node="false"
        >
          <span class="worker-box-tree-span" slot-scope="{ data }">
            <span>
              <template v-if="tool.isNotEmpty(data.key)">
                {{ data.key }}
                :
              </template>
              {{ data.value }}
            </span>
            <span class="mgl-20">
              <a
                class="tm-link color-orange ft-15 mgr-2"
                @click="toSubmit(data)"
                title="提交"
              >
                <i class="mdi mdi-send-check"></i>
              </a>
            </span>
          </span>
        </el-tree>
      </div>
    </div>
    <div class="worker-kafka-form" v-if="connect.open">
      <div class="ft-16 pdb-15 color-orange">
        <template v-if="readonlyOne">
          <span>查看</span>
        </template>
        <template v-else-if="insertOne">
          <span>新增</span>
        </template>
        <template v-else-if="updateOne">
          <span>修改</span>
        </template>
      </div>
      <el-form :model="oneForm" size="lg" @submit.native.prevent>
        <el-form-item label="topic">
          <el-input
            v-model="oneForm.topic"
            placeholder="topic"
            :readonly="readonlyOne || updateOne"
          ></el-input>
        </el-form-item>
        <el-form-item label="partition（分区）">
          <el-input
            v-model="oneForm.partition"
            placeholder="partition"
            :readonly="readonlyOne || updateOne"
          ></el-input>
        </el-form-item>
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
            :readonly="readonlyOne || updateOne"
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
      },
      connect: {
        open: false,
        form: null,
      },
      loading: false,
      searchForm: {
        topic: "test-topic",
        groupId: "test-group",
        keyType: "String",
        valueType: "String",
      },
      readonlyOne: true,
      insertOne: true,
      updateOne: true,
      oneForm: {
        topic: null,
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
    pull(data) {
      return server.kafka.pull(data);
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
          this.loadDatas();
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
              this.loadDatas();
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
      this.topics = null;
      this.datas = null;
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
              datas.push({ key: name, name: name, topic: name });
            });
            this.topics = datas;
          }
        })
        .catch(() => {
          this.loading = false;
        });
    },
    loadDatas() {
      let data = {};
      Object.assign(data, this.connect.form);
      Object.assign(data, this.searchForm);
      if (tool.isEmpty(data.topic)) {
        tool.error("topic不能为空！");
        return;
      }
      if (tool.isEmpty(data.groupId)) {
        tool.error("groupId不能为空！");
        return;
      }
      this.datas = null;
      this.datas_loading = true;
      this.pull(data)
        .then((res) => {
          this.datas_loading = false;
          if (res.code != 0) {
            tool.error(res.msg);
          } else {
            let value = res.value || {};
            let msgs = value.msgs || [];
            let datas = [];

            msgs.forEach((msg) => {
              datas.push(msg);
            });
            this.datas = datas;
          }
        })
        .catch(() => {
          this.datas_loading = false;
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
      this.oneForm.topic = this.searchForm.topic;
      this.oneForm.key = "";
      this.oneForm.value = "";
      this.oneForm.partition = 0;
      this.oneForm.offset = 0;
      this.updateOne = false;
      this.insertOne = true;
      this.readonlyOne = false;
    },
    topicCurrentChange(data) {
      for (var key in data) {
        this.searchForm[key] = data[key];
      }
    },
    dataCurrentChange(data) {
      this.toUpdate(data);
    },
    toUpdate(data) {
      this.oneForm.key = data.key;
      this.oneForm.value = data.value;
      this.oneForm.topic = data.topic;
      this.oneForm.partition = data.partition;
      this.oneForm.offset = data.offset;

      this.insertOne = false;
      this.readonlyOne = false;
      this.updateOne = true;
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
.worker-kafka-wrap .worker-kafka-topic-list-box {
  height: calc(100% - 100px);
  width: 300px;
  margin: 10px;
  padding: 0px;
  position: relative;
  float: left;
  overflow: hidden !important;
}
.worker-kafka-wrap .worker-kafka-topic-list {
  height: calc(100% - 50px);
  margin: 0px;
  padding: 0px;
  position: relative;
  overflow-x: hidden !important;
}
.worker-kafka-wrap .worker-kafka-data-list-box {
  height: calc(100% - 100px);
  width: calc(100% - 760px);
  min-width: 400px;
  margin: 10px;
  padding: 0px;
  position: relative;
  float: left;
  overflow: hidden;
}
.worker-kafka-wrap .worker-kafka-data-list {
  height: calc(100% - 150px);
  margin: 0px;
  padding: 0px;
  position: relative;
  overflow-x: hidden !important;
}
.worker-kafka-wrap .worker-kafka-form {
  height: calc(100% - 100px);
  width: 400px;
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
