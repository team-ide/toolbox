<template>
  <div class="worker-kafka-wrap">
    <tm-layout height="100%">
      <tm-layout height="50px">
        <WorkerConfig
          :workerKey="workerKey"
          workerType="kafka"
          :connect="connect"
          @connect="doConnect"
        ></WorkerConfig>
      </tm-layout>
      <tm-layout-bar bottom></tm-layout-bar>
      <tm-layout height="auto">
        <tm-layout height="100%" width="100%">
          <tm-layout width="300px" v-loading="topics_loading">
            <tm-layout height="40px">
              <div class="pdlr-10">
                <div class="worker-panel-title" v-if="connect.open">
                  Topics列表（{{ connect.title }}）
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
                      <div class="worker-btn-group">
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
                      </div>
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
                        新增推送
                      </a>
                      <a
                        class="tm-btn tm-btn-sm color-orange mgl-5"
                        @click="toCreatePartitions"
                      >
                        设置分区数量
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
                        <div class="worker-btn-group">
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
                        </div>
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
              <el-form :model="oneForm" size="mini" @submit.native.prevent>
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
                <div class="">
                  Header
                  <a
                    class="tm-link color-green ft-15 mgl-20"
                    @click="addHeader"
                    title="添加"
                    v-if="!(readonlyOne || updateOne)"
                  >
                    <i class="mdi mdi-plus"></i>
                  </a>
                </div>
                <div
                  v-if="oneForm.headers.length > 0"
                  class="pd-5 mgt-5 bd-1 bd-grey-7"
                >
                  <template v-for="(one, index) in oneForm.headers">
                    <el-row :key="index">
                      <el-col :span="11">
                        <el-form-item label="Key">
                          <el-input
                            v-model="one.key"
                            :readonly="readonlyOne || updateOne"
                          ></el-input>
                        </el-form-item>
                      </el-col>
                      <el-col :span="11">
                        <el-form-item label="Value">
                          <el-input
                            v-model="one.value"
                            :readonly="readonlyOne || updateOne"
                          ></el-input>
                        </el-form-item>
                      </el-col>
                      <el-col :span="2">
                        <a
                          class="tm-link color-red ft-15 mgt-30"
                          @click="removeHeader(one)"
                          title="删除"
                          v-if="!(readonlyOne || updateOne)"
                        >
                          <i class="mdi mdi-delete-outline"></i>
                        </a>
                      </el-col>
                    </el-row>
                  </template>
                </div>

                <el-form-item
                  label="key类型"
                  :show-message="false"
                  class="mgb-0"
                >
                  <el-select
                    v-model="oneForm.keyType"
                    placeholder="keyType"
                    style="width: 80px"
                  >
                    <el-option label="String" value="String"> </el-option>
                    <el-option label="Long" value="Long"> </el-option>
                  </el-select>
                </el-form-item>
                <el-form-item label="key">
                  <el-input
                    v-model="oneForm.key"
                    placeholder="key"
                    :readonly="readonlyOne || updateOne"
                  ></el-input>
                </el-form-item>
                <el-form-item
                  label="value类型"
                  :show-message="false"
                  class="mgb-0"
                >
                  <el-select
                    v-model="oneForm.valueType"
                    placeholder="valueType"
                    style="width: 80px"
                  >
                    <el-option label="String" value="String"> </el-option>
                    <el-option label="Long" value="Long"> </el-option>
                  </el-select>
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
    <el-dialog
      title="设置Topic分区"
      destroy-on-close
      :visible.sync="showCreatePartitions"
      width="600"
    >
      <el-form :model="createPartitionsForm" size="mini" @submit.native.prevent>
        <el-form-item label="topic">
          <el-input
            v-model="createPartitionsForm.topic"
            placeholder="topic"
          ></el-input>
        </el-form-item>
        <el-form-item label="分区数量">
          <el-input
            v-model="createPartitionsForm.count"
            placeholder="分区数量"
          ></el-input>
        </el-form-item>
        <el-form-item>
          <a class="tm-btn color-green" @click="doCreatePartitions"> 设置 </a>
        </el-form-item>
      </el-form>
    </el-dialog>
  </div>
</template>

<script>
import server from "@/server";
import tool from "@/tool";
import source from "@/source";

import WorkerConfig from "./WorkerConfig";

export default {
  components: { WorkerConfig },
  props: ["workerKey"],
  data() {
    return {
      tool,
      source,
      connect: {
        open: false,
        config: null,
        title: null,
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
        keyType: "String",
        valueType: "String",
        topic: null,
        key: null,
        value: null,
        partition: null,
        offset: null,
        headers: [],
        json: null,
      },
      topics: null,
      datasTopic: null,
      datas: null,
      datas_loading: false,
      showCreatePartitions: false,
      createPartitionsForm: {
        topic: null,
        count: 1,
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
  computed: {},
  methods: {
    loadTopics(pattern, size) {
      let data = {};
      Object.assign(data, this.connect.config);
      return server.kafka.topics(data);
    },
    pull(data) {
      return server.kafka.pull(data);
    },
    push() {
      let data = {};
      Object.assign(data, this.connect.config);
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
    toCreatePartitions() {
      this.createPartitionsForm.topic = this.searchForm.topic;
      this.createPartitionsForm.count = 1;
      this.showCreatePartitions = true;
    },
    doCreatePartitions() {
      let data = {};
      Object.assign(data, this.connect.config);
      Object.assign(data, this.createPartitionsForm);
      if (tool.isEmpty(data.topic)) {
        tool.error("topic不能为空！");
        return;
      }
      if (tool.isEmpty(data.count)) {
        tool.error("分区数量不能为空！");
        return;
      }
      data.count = Number(data.count);
      tool
        .confirm(
          "将设置Topic[" +
            data.topic +
            "] 分区数量[" +
            data.count +
            "]，确认设置？"
        )
        .then(() => {
          server.kafka.createPartitions(data).then((res) => {
            if (res.code != 0) {
              tool.error(res.msg);
            } else {
              this.showCreatePartitions = false;
              tool.success("设置成功");
              this.loadDatas();
            }
          });
        })
        .catch(() => {});
    },
    toDeleteTopic(one) {
      let data = {};
      Object.assign(data, this.connect.config);
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
      Object.assign(data, this.connect.config);
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
      Object.assign(data, this.connect.config);
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
      return new Promise((resolve, reject) => {
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
            resolve && resolve(res);
          })
          .catch((e) => {
            this.topics_loading = false;
            reject && reject(e);
          });
      });
    },
    loadDatas() {
      let data = {};
      Object.assign(data, this.connect.config);
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
        .catch((e) => {
          this.datas_loading = false;
        });
    },
    doConnect(config, callback) {
      this.toInsert();

      this.$nextTick(() => {
        this.initTopics()
          .then((res) => {
            callback(res);
          })
          .catch((e) => {
            callback(e);
          });
      });
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
      this.oneForm.headers = [];
      this.oneForm.keyType = this.searchForm.keyType;
      this.oneForm.valueType = this.searchForm.valueType;
      this.updateOne = false;
      this.insertOne = true;
      this.readonlyOne = false;
    },
    addHeader() {
      this.oneForm.headers.push({ key: "", value: "" });
    },
    removeHeader(one) {
      if (this.oneForm.headers.indexOf(one) >= 0) {
        this.oneForm.headers.splice(this.oneForm.headers.indexOf(one), 1);
      }
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
      this.oneForm.headers = data.headers || [];
      this.oneForm.keyType = this.searchForm.keyType;
      this.oneForm.valueType = this.searchForm.valueType;

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
      this.oneForm.headers = data.headers || [];
      this.oneForm.keyType = this.searchForm.keyType;
      this.oneForm.valueType = this.searchForm.valueType;

      this.insertOne = false;
      this.readonlyOne = true;
      this.updateOne = false;
    },
    init() {},
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
