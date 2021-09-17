<template>
  <div class="worker-kafka-wrap">
    <tm-layout height="100%">
      <tm-layout height="50px">
        <el-form
          class="pd-10"
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
            <a class="tm-btn tm-btn-sm color-green" @click="doConnect">
              连接
            </a>
          </el-form-item>
        </el-form>
      </tm-layout>
      <tm-layout-bar bottom></tm-layout-bar>
      <tm-layout height="auto">
        <tm-layout height="100%" width="100%">
          <tm-layout width="300px" v-loading="topics_loading">
            <tm-layout height="40px">
              <div class="pdlr-10">
                <div class="worker-panel-title" v-if="connect.open">
                  Topics列表（{{ connect.form.address }}）
                </div>
              </div>
            </tm-layout>
            <tm-layout height="auto">
              <div class="worker-kafka-topic-list" v-if="connect.open">
                <el-table
                  :data="topics"
                  height="100%"
                  style="width: 100%"
                  size="mini"
                >
                  <el-table-column prop="topic" label="topic">
                  </el-table-column>
                  <el-table-column width="60">
                    <template slot-scope="scope">
                      <a
                        class="tm-link color-green ft-15 mgr-2"
                        title="查看消息"
                        @click="selectTopic(scope.row)"
                      >
                        <i class="mdi mdi-eye-outline"></i>
                      </a>
                      <a
                        class="tm-link color-orange ft-15 mgr-2"
                        title="删除"
                        @click="toDeleteTopic(scope.row)"
                      >
                        <i class="mdi mdi-delete-outline"></i>
                      </a>
                    </template>
                  </el-table-column>
                </el-table>
              </div>
            </tm-layout>
          </tm-layout>
          <tm-layout-bar right></tm-layout-bar>
          <tm-layout width="auto">
            <tm-layout height="100%">
              <tm-layout height="140px">
                <div class="pdlr-10" v-if="connect.open">
                  <div class="worker-panel-title">Msgs搜索</div>
                  <el-divider class="mg-0"></el-divider>
                  <el-form
                    class="pdt-10"
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
                      <a
                        class="tm-btn tm-btn-sm color-green mgl-5"
                        @click="loadDatas"
                        :class="{ 'tm-disabled': datas_loading }"
                      >
                        拉取
                      </a>
                      <a
                        class="tm-btn tm-btn-sm color-blue mgl-5"
                        @click="toInsert"
                      >
                        新增
                      </a>
                    </el-form-item>
                  </el-form>
                </div>
              </tm-layout>
              <tm-layout-bar bottom></tm-layout-bar>
              <tm-layout height="40px">
                <div class="pdlr-10">
                  <div class="worker-panel-title" v-if="connect.open">
                    Msgs列表（{{ datasTopic }}）
                  </div>
                </div>
              </tm-layout>
              <tm-layout height="auto" v-loading="datas_loading">
                <div
                  class="worker-kafka-data-list"
                  v-if="connect.open && datas != null"
                >
                  <el-table
                    :data="datas"
                    height="100%"
                    style="width: 100%"
                    size="mini"
                  >
                    <el-table-column
                      prop="partition"
                      label="Partition"
                      width="80"
                    >
                    </el-table-column>
                    <el-table-column prop="offset" label="Offset" width="80">
                    </el-table-column>
                    <el-table-column prop="key" label="Key" width="110">
                    </el-table-column>
                    <el-table-column prop="value" label="Value">
                    </el-table-column>
                    <el-table-column width="90">
                      <template slot-scope="scope">
                        <a
                          class="tm-link color-grey-1 ft-15 mgr-5"
                          @click="toShow(scope.row)"
                          title="查看"
                        >
                          <i class="mdi mdi-eye"></i>
                        </a>
                        <a
                          class="tm-link color-orange ft-15 mgr-5"
                          @click="toCommit(scope.row)"
                          title="提交"
                        >
                          <i class="mdi mdi-send-check"></i>
                        </a>
                        <a
                          class="tm-link color-red-3 ft-15 mgr-5"
                          title="删除数据"
                          @click="toDeleteData(scope.row)"
                        >
                          <i class="mdi mdi-delete-outline"></i>
                        </a>
                      </template>
                    </el-table-column>
                  </el-table>
                </div>
              </tm-layout>
            </tm-layout>
          </tm-layout>
          <tm-layout-bar left></tm-layout-bar>
          <tm-layout width="300px">
            <div class="pdlr-10" v-if="connect.open">
              <div class="worker-panel-title">
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
              <el-divider class="mg-0"></el-divider>
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
                  <a
                    v-if="!readonlyOne"
                    class="tm-btn color-green"
                    @click="push"
                  >
                    推送
                  </a>
                  <a
                    v-if="readonlyOne"
                    class="tm-btn color-orange"
                    @click="toCommit(oneForm)"
                  >
                    提交
                  </a>
                </el-form-item>
              </el-form>
            </div>
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
      topics_loading: false,
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
        partition: null,
        offset: null,
        json: null,
      },
      topics: null,
      datasTopic: null,
      datas: null,
      datas_loading: false,
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
      if (tool.isNotEmpty(data.partition)) {
        data.partition = Number(data.partition);
      } else {
        data.partition = -1;
      }
      if (tool.isNotEmpty(data.offset)) {
        data.offset = Number(data.offset);
      } else {
        data.offset = -1;
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
    toCreatePartitions(data) {},
    toDeleteTopic(one) {
      let data = {};
      Object.assign(data, this.connect.form);
      Object.assign(data, this.searchForm);
      if (window.event) {
        window.event.stopPropagation && window.event.stopPropagation();
      }
      data.topic = one.topic;
      if (tool.isEmpty(data.topic)) {
        tool.error("topic不能为空！");
        return;
      }
      tool
        .confirm("将删除Topic[" + data.topic + "]且无法恢复，确认删除？")
        .then(() => {
          server.kafka.deleteTopic(data).then((res) => {
            if (res.code != 0) {
              tool.error(res.msg);
            } else {
              tool.success("删除成功");
              this.initTopics();
            }
          });
        })
        .catch(() => {});
    },
    toDeleteData(one) {
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
      data.offset = one.offset + 1;
      if (tool.isEmpty(data.offset)) {
        tool.error("offset不能为空！");
        return;
      }
      data.partition = Number(data.partition);
      data.offset = Number(data.offset);
      tool
        .confirm(
          "将删除Topic[" +
            data.topic +
            "]分区[" +
            data.partition +
            "]位置[" +
            data.offset +
            "]且无法恢复，确认提交？"
        )
        .then(() => {
          server.kafka.deleteRecords(data).then((res) => {
            if (res.code != 0) {
              tool.error(res.msg);
            } else {
              tool.success("删除成功");
              this.loadDatas();
            }
          });
        })
        .catch(() => {});
    },
    toCommit(one) {
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
      data.offset = one.offset + 1;
      if (tool.isEmpty(data.offset)) {
        tool.error("offset不能为空！");
        return;
      }
      data.partition = Number(data.partition);
      data.offset = Number(data.offset);
      tool
        .confirm(
          "将使用Group[" +
            data.groupId +
            "]提交Topic[" +
            data.topic +
            "]分区[" +
            data.partition +
            "]位置[" +
            data.offset +
            "]，确认提交？"
        )
        .then(() => {
          server.kafka.commit(data).then((res) => {
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
    doSearch() {
      this.initTopics();
    },
    reload() {
      this.initTopics();
    },
    initTopics() {
      this.topics = null;
      this.datas = null;
      this.topics_loading = true;
      this.loadTopics()
        .then((res) => {
          this.topics_loading = false;
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
          this.topics_loading = false;
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
      this.datasTopic = data.topic;
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
        this.initTopics();
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
      this.oneForm.partition = null;
      this.oneForm.offset = null;
      this.updateOne = false;
      this.insertOne = true;
      this.readonlyOne = false;
    },
    selectTopic(data) {
      this.searchForm.topic = data.topic;
      this.loadDatas();
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
    toShow(data) {
      this.oneForm.key = data.key;
      this.oneForm.value = data.value;
      this.oneForm.topic = data.topic;
      this.oneForm.partition = data.partition;
      this.oneForm.offset = data.offset;

      this.insertOne = false;
      this.readonlyOne = true;
      this.updateOne = false;
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
  margin: 0px;
  padding: 0px;
  position: relative;
}
.worker-kafka-wrap .worker-kafka-topic-list {
  height: 100%;
}
.worker-kafka-wrap .worker-kafka-data-list {
  height: 100%;
}
.worker-kafka-wrap .el-tree {
  /* border: 1px solid #f3f3f3; */
  border-bottom: 0px;
}
.worker-kafka-wrap .el-tree-node__content {
  border-bottom: 1px dotted #696969;
}
</style>
