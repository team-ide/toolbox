<template>
  <div class="worker-redis-wrap">
    <tm-layout height="100%">
      <tm-layout height="50px">
        <WorkerConfig
          :workerKey="workerKey"
          workerType="redis"
          :connect="connect"
          @connect="doConnect"
        ></WorkerConfig>
      </tm-layout>
      <tm-layout-bar bottom></tm-layout-bar>
      <tm-layout height="auto">
        <tm-layout height="100%" width="100%">
          <tm-layout width="auto">
            <tm-layout height="100%">
              <tm-layout height="140px">
                <div class="pdlr-10">
                  <div class="worker-panel-title" v-if="connect.open">
                    Keys搜索（{{ connect.title }}）
                  </div>
                  <el-divider class="mg-0"></el-divider>
                  <el-form
                    v-if="connect.open"
                    class="pdt-10"
                    :inline="true"
                    :model="searchForm"
                    size="mini"
                    @submit.native.prevent
                  >
                    <el-form-item label="搜索（*模糊匹配）">
                      <el-input
                        v-model="searchForm.pattern"
                        placeholder="pattern"
                      ></el-input>
                    </el-form-item>
                    <el-form-item label="展示数量">
                      <el-input
                        v-model="searchForm.size"
                        placeholder="size"
                      ></el-input>
                    </el-form-item>
                    <el-form-item>
                      <a
                        class="tm-btn tm-btn-sm color-green mgl-5"
                        :class="{ 'tm-disabled': loading }"
                        @click="doSearch"
                      >
                        搜索
                      </a>
                      <a
                        class="tm-btn tm-btn-sm color-red mgl-5"
                        @click="deletePattern"
                      >
                        删除匹配
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
                    Keys列表
                  </div>
                </div>
              </tm-layout>
              <tm-layout height="auto" v-loading="loading">
                <div class="worker-redis-list" v-if="connect.open">
                  <el-table
                    :data="keys"
                    height="100%"
                    style="width: 100%"
                    size="mini"
                  >
                    <el-table-column prop="key" label="Key"> </el-table-column>
                    <el-table-column width="60">
                      <template slot-scope="scope">
                        <div class="worker-btn-group">
                          <a
                            class="tm-link color-green ft-15 mgr-2"
                            title="编辑"
                            @click="toUpdate(scope.row)"
                          >
                            <i class="mdi mdi-playlist-edit"></i>
                          </a>
                          <a
                            class="tm-link color-orange ft-15 mgr-2"
                            title="删除"
                            @click="toDelete(scope.row)"
                          >
                            <i class="mdi mdi-delete-outline"></i>
                          </a>
                        </div>
                      </template>
                    </el-table-column>
                  </el-table>
                </div>
              </tm-layout>
              <tm-layout-bar top></tm-layout-bar>
              <tm-layout height="30px">
                <div
                  v-if="connect.open && keys != null"
                  class="text-center pd-5 ft-12"
                >
                  共搜索
                  <span class="color-orange ft-15 mglr-5">{{ count }}</span>
                  条记录
                </div>
              </tm-layout>
            </tm-layout>
          </tm-layout>
          <tm-layout-bar left></tm-layout-bar>
          <tm-layout width="500px">
            <div
              class="pdlr-10"
              v-if="connect.open"
              style="overflow: auto; height: 100%"
            >
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
                <el-form-item label="key">
                  <el-input
                    v-model="oneForm.key"
                    placeholder="key"
                    :readonly="readonlyOne || updateOne"
                  ></el-input>
                </el-form-item>
                <el-form-item label="type">
                  <el-select
                    v-model="oneForm.type"
                    placeholder="选择类型"
                    style="width: 100%"
                  >
                    <el-option label="string" value="string"></el-option>
                    <el-option label="list" value="list"></el-option>
                    <el-option label="set" value="set"></el-option>
                    <el-option label="hash" value="hash"></el-option>
                  </el-select>
                </el-form-item>
                <template v-if="oneForm.type == 'string'">
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
                    <a
                      v-if="!readonlyOne"
                      class="tm-btn color-green"
                      @click="toDo('set', { value: oneForm.value })"
                    >
                      保存
                    </a>
                  </el-form-item>
                </template>
                <template v-else-if="oneForm.type == 'list'">
                  <div class="">
                    List
                    <a
                      class="tm-link color-green ft-15 mgl-20"
                      @click="
                        addOne('list', {
                          value: null,
                          isNew: true,
                          index: oneForm.list.length,
                        })
                      "
                      title="添加"
                      v-if="!readonlyOne"
                    >
                      <i class="mdi mdi-plus"></i>
                    </a>
                  </div>
                  <div
                    v-if="oneForm.list.length > 0"
                    class="pd-5 mgt-5 bd-1 bd-grey-7"
                  >
                    <template v-for="(one, index) in oneForm.list">
                      <el-row :key="index">
                        <el-col :span="20">
                          <el-form-item
                            :label="'Value' + '（' + one.index + '）'"
                          >
                            <el-input
                              v-model="one.value"
                              :readonly="readonlyOne"
                            ></el-input>
                          </el-form-item>
                        </el-col>
                        <el-col :span="4">
                          <a
                            class="tm-link color-green ft-15 mgt-30"
                            @click="
                              toDo(one.isNew ? 'rpush' : 'lset', {
                                value: one.value,
                                index: one.index,
                              })
                            "
                            title="保存"
                            v-if="!readonlyOne"
                          >
                            <i class="mdi mdi-content-save-outline"></i>
                          </a>
                          <a
                            class="tm-link color-red ft-15 mgt-30"
                            @click="
                              toDo('lrem', { value: one.value, count: 0 })
                            "
                            title="删除"
                            v-if="!readonlyOne && one.canDel"
                          >
                            <i class="mdi mdi-delete-outline"></i>
                          </a>
                        </el-col>
                      </el-row>
                    </template>
                  </div>
                </template>
                <template v-else-if="oneForm.type == 'set'">
                  <div class="">
                    Set
                    <a
                      class="tm-link color-green ft-15 mgl-20"
                      @click="
                        addOne('set', {
                          value: null,
                          isNew: true,
                          index: oneForm.set.length,
                        })
                      "
                      title="添加"
                      v-if="!readonlyOne"
                    >
                      <i class="mdi mdi-plus"></i>
                    </a>
                  </div>
                  <div
                    v-if="oneForm.set.length > 0"
                    class="pd-5 mgt-5 bd-1 bd-grey-7"
                  >
                    <template v-for="(one, index) in oneForm.set">
                      <el-row :key="index">
                        <el-col :span="20">
                          <el-form-item
                            :label="'Value' + '（' + one.index + '）'"
                          >
                            <el-input
                              v-model="one.value"
                              :readonly="readonlyOne"
                            ></el-input>
                          </el-form-item>
                        </el-col>
                        <el-col :span="4">
                          <a
                            class="tm-link color-green ft-15 mgt-30"
                            @click="
                              toDo('sadd', {
                                value: one.value,
                                index: one.index,
                              })
                            "
                            title="保存"
                            v-if="!readonlyOne"
                          >
                            <i class="mdi mdi-content-save-outline"></i>
                          </a>
                          <a
                            class="tm-link color-red ft-15 mgt-30"
                            @click="toDo('srem', { value: one.value })"
                            title="删除"
                            v-if="!readonlyOne && !one.isNew"
                          >
                            <i class="mdi mdi-delete-outline"></i>
                          </a>
                        </el-col>
                      </el-row>
                    </template>
                  </div>
                </template>
                <template v-else-if="oneForm.type == 'hash'">
                  <div class="">
                    Hash
                    <a
                      class="tm-link color-green ft-15 mgl-20"
                      @click="
                        addOne('hash', {
                          field: null,
                          value: null,
                          isNew: true,
                          index: oneForm.hash.length,
                        })
                      "
                      title="添加"
                      v-if="!readonlyOne"
                    >
                      <i class="mdi mdi-plus"></i>
                    </a>
                  </div>
                  <div
                    v-if="oneForm.hash.length > 0"
                    class="pd-5 mgt-5 bd-1 bd-grey-7"
                  >
                    <template v-for="(one, index) in oneForm.hash">
                      <el-row :key="index">
                        <el-col :span="10">
                          <el-form-item
                            :label="'Field' + '（' + one.index + '）'"
                          >
                            <el-input
                              v-model="one.field"
                              :readonly="readonlyOne"
                            ></el-input>
                          </el-form-item>
                        </el-col>
                        <el-col :span="10">
                          <el-form-item :label="'Value'">
                            <el-input
                              v-model="one.value"
                              :readonly="readonlyOne"
                            ></el-input>
                          </el-form-item>
                        </el-col>
                        <el-col :span="4">
                          <a
                            class="tm-link color-green ft-15 mgt-30"
                            @click="
                              toDo('hset', {
                                field: one.field,
                                value: one.value,
                                index: one.index,
                              })
                            "
                            title="保存"
                            v-if="!readonlyOne"
                          >
                            <i class="mdi mdi-content-save-outline"></i>
                          </a>
                          <a
                            class="tm-link color-red ft-15 mgt-30"
                            @click="toDo('hdel', { field: one.field })"
                            title="删除"
                            v-if="!readonlyOne && !one.isNew"
                          >
                            <i class="mdi mdi-delete-outline"></i>
                          </a>
                        </el-col>
                      </el-row>
                    </template>
                  </div>
                </template>
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
      loading: false,
      searchForm: { pattern: "*", size: 20 },
      readonlyOne: true,
      insertOne: true,
      updateOne: true,
      oneForm: {
        key: null,
        value: null,
        set: [],
        list: [],
        hash: [],
        type: null,
        json: null,
      },
      keys: null,
      count: 0,
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
    addOne(type, one) {
      this.oneForm[type] = this.oneForm[type] || {};
      this.oneForm[type].push(one);
    },
    loadKeys(pattern, size) {
      let data = {};
      Object.assign(data, this.connect.config);
      data.pattern = pattern;
      data.size = Number(size);
      return server.redis.keys(data);
    },
    get(key) {
      let data = {};
      Object.assign(data, this.connect.config);
      data.key = key;
      return server.redis.get(data);
    },
    toDo(type, data) {
      data = data || {};
      Object.assign(data, this.connect.config);
      data.key = this.oneForm.key;
      data.type = type;
      if (tool.isEmpty(type)) {
        tool.error("操作类型不能为空！");
        return;
      }
      if (tool.isEmpty(data.key)) {
        tool.error("Key不能为空！");
        return;
      }
      server.redis.do(data).then((res) => {
        if (res.code != 0) {
          tool.error(res.msg);
        } else {
          tool.success("操作成功");
          this.reload();
          this.toUpdate({ key: data.key });
        }
      });
    },
    toDelete(one) {
      if (window.event) {
        window.event.stopPropagation && window.event.stopPropagation();
      }
      let data = {};
      Object.assign(data, this.connect.config);
      data.key = one.key;
      if (tool.isEmpty(data.key)) {
        tool.error("Key不能为空！");
        return;
      }
      tool
        .confirm("将删除[" + data.key + "]，确认删除？")
        .then(() => {
          server.redis.delete(data).then((res) => {
            if (res.code != 0) {
              tool.error(res.msg);
            } else {
              tool.success("删除成功");
              this.reload();
            }
          });
        })
        .catch(() => {});
    },
    deletePattern() {
      if (window.event) {
        window.event.stopPropagation && window.event.stopPropagation();
      }
      let data = {};
      Object.assign(data, this.connect.config);
      Object.assign(data, this.searchForm);
      if (tool.isEmpty(data.pattern)) {
        tool.error("匹配字符不能为空！");
        return;
      }
      tool
        .confirm("将删除匹配[" + data.pattern + "]所有Key，确认删除？")
        .then(() => {
          server.redis.deletePattern(data).then((res) => {
            res.value = res.value || {};
            if (res.code != 0) {
              if (res.value.count > 0) {
                tool.error(
                  "共删除" + res.value.count + "条记录出现异常：" + res.msg
                );
                this.reload();
              } else {
                tool.error("删除失败：" + res.msg);
              }
            } else {
              tool.success("删除成功，共删除" + res.value.count + "条记录");
              this.reload();
            }
          });
        })
        .catch(() => {});
    },
    doSearch() {
      this.load();
    },
    reload(key) {
      this.load();
    },
    load() {
      let data = {};
      Object.assign(data, this.searchForm);
      this.data = null;
      this.count = 0;
      this.loading = true;
      return new Promise((resolve, reject) => {
        this.loadKeys(data.pattern, data.size)
          .then((res) => {
            this.loading = false;
            if (res.code != 0) {
              tool.error(res.msg);
            } else {
              let value = res.value || {};
              let keys = value.keys || [];
              this.count = value.count;
              let datas = [];

              keys.forEach((name) => {
                datas.push({ key: name, name: name });
              });
              this.keys = datas;
            }
            resolve && resolve(res);
          })
          .catch((e) => {
            this.loading = false;
            reject && reject(e);
          });
      });
    },
    doConnect(config, callback) {
      this.toInsert();
      this.load()
        .then((res) => {
          callback(res);
        })
        .catch((e) => {
          callback(e);
        });
    },
    toInsert() {
      if (window.event) {
        window.event.stopPropagation && window.event.stopPropagation();
      }
      this.oneForm.key = "";
      this.oneForm.type = "string";
      this.oneForm.value = "";
      this.oneForm.list = [];
      this.oneForm.set = [];
      this.oneForm.hash = [];
      this.updateOne = false;
      this.insertOne = true;
      this.readonlyOne = false;
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
          this.oneForm.type = value.type;
          if (this.oneForm.type == "string") {
            this.oneForm.value = value.value;
          } else if (this.oneForm.type == "list") {
            value.value = value.value || [];
            this.oneForm.list = [];
            value.value.forEach((one, index) => {
              let canDel = true;
              value.value.forEach((one_, index_) => {
                if (one_ == one && index != index_) {
                  canDel = false;
                }
              });
              this.oneForm.list.push({
                value: one,
                index: index,
                canDel: canDel,
              });
            });
          } else if (this.oneForm.type == "set") {
            value.value = value.value || [];
            this.oneForm.set = [];
            value.value.forEach((one, index) => {
              this.oneForm.set.push({
                value: one,
                index: index,
              });
            });
          } else if (this.oneForm.type == "hash") {
            value.value = value.value || [];
            this.oneForm.hash = [];
            let index = 0;
            for (var field in value.value) {
              this.oneForm.hash.push({
                field: field,
                value: value.value[field],
                index: index,
              });
              index++;
            }
          }
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
          this.oneForm.type = value.type;
          this.oneForm.value = value.value;
        }
      });
    },
    init() {},
  },
  mounted() {
    this.init();
  },
};
</script>

<style>
.worker-redis-wrap {
  height: 100%;
  margin: 0px;
  padding: 0px;
  position: relative;
}
.worker-redis-wrap .worker-redis-list {
  height: 100%;
}
.worker-redis-wrap .el-tree {
  /* border: 1px solid #f3f3f3; */
  border-bottom: 0px;
}
.worker-redis-wrap .el-tree-node__content {
  border-bottom: 1px dotted #696969;
}
</style>
