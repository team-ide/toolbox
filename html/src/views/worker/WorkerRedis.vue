<template>
  <div class="worker-redis-wrap">
    <tm-layout height="100%">
      <tm-layout height="50px">
        <el-form
          class="pd-10"
          :inline="true"
          :model="configForm"
          size="mini"
          @submit.native.prevent
        >
          <el-form-item label="address（多个使用“;”隔开）">
            <el-input
              v-model="configForm.address"
              placeholder="address"
            ></el-input>
          </el-form-item>
          <el-form-item label="auth">
            <el-input v-model="configForm.auth" placeholder="auth"></el-input>
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
          <tm-layout width="auto">
            <tm-layout height="100%">
              <tm-layout height="140px">
                <div class="pdlr-10">
                  <div class="worker-panel-title" v-if="connect.open">
                    Keys搜索（{{ connect.form.address }}）
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
              <tm-layout height="45px">
                <div class="worker-panel-title  pdlr-10" v-if="connect.open">
                  Keys列表
                </div>
                <el-divider class="mg-0"></el-divider>
              </tm-layout>
              <tm-layout height="auto">
                <div
                  class="worker-redis-list worker-scrollbar"
                  ref="treeBox"
                  v-if="connect.open"
                >
                  <el-tree
                    ref="tree"
                    :props="defaultProps"
                    :data="data"
                    :default-expanded-keys="expands"
                    node-key="key"
                    @node-click="nodeClick"
                    @current-change="currentChange"
                    :expand-on-click-node="false"
                  >
                    <span
                      class="worker-box-tree-span"
                      slot-scope="{ node, data }"
                    >
                      <span>{{ node.label }}</span>
                      <span class="mgl-20">
                        <a
                          class="tm-link color-orange ft-15 mgr-2"
                          @click="toDelete(data)"
                        >
                          <i class="mdi mdi-delete-outline"></i>
                        </a>
                      </span>
                    </span>
                  </el-tree>
                </div>
              </tm-layout>
              <tm-layout-bar top></tm-layout-bar>
              <tm-layout height="30px">
                <div
                  v-if="connect.open && data != null"
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
          <tm-layout width="400px">
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
        address: "127.0.0.1:6379",
        auth: "",
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
      expands: [],
      opens: [],
      data: null,
      count: 0,
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
    keys(pattern, size) {
      let data = {};
      Object.assign(data, this.connect.form);
      data.pattern = pattern;
      data.size = Number(size);
      return server.redis.keys(data);
    },
    get(key) {
      let data = {};
      Object.assign(data, this.connect.form);
      data.key = key;
      return server.redis.get(data);
    },
    doSave() {
      let data = {};
      Object.assign(data, this.connect.form);
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
      Object.assign(data, this.connect.form);
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
      Object.assign(data, this.connect.form);
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
    nodeClick() {},
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
      this.keys(data.pattern, data.size)
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
            this.data = datas;
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
        this.load();
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
    currentChange(data) {
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
