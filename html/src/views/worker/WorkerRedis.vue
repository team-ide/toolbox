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
                <el-form-item label="key">
                  <el-input
                    v-model="oneForm.key"
                    placeholder="key"
                    :readonly="readonlyOne || updateOne"
                  ></el-input>
                </el-form-item>
                <el-form-item label="type">
                  <el-select v-model="oneForm.type" placeholder="选择类型">
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
                      @click="doSave"
                    >
                      保存
                    </a>
                  </el-form-item>
                </template>
                <template v-if="oneForm.type == 'set'"> </template>
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
        set: null,
        list: null,
        hash: null,
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
    doSave() {
      let data = {};
      Object.assign(data, this.connect.config);
      Object.assign(data, this.oneForm);
      if (tool.isEmpty(data.key)) {
        tool.error("Key不能为空！");
        return;
      }
      server.redis.save(data).then((res) => {
        if (res.code != 0) {
          tool.error(res.msg);
        } else {
          tool.success("保存成功");
          this.reload();
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
