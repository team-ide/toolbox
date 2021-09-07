<template>
  <div class="worker-redis-wrap">
    <div class="worker-redis-config pd-10 pdb-0">
      <el-form
        :inline="true"
        :model="configForm"
        size="mini"
        @submit.native.prevent
      >
        <el-form-item label="集群">
          <el-switch
            v-model="configForm.cluster"
            placeholder="cluster"
          ></el-switch>
        </el-form-item>
        <el-form-item label="address">
          <el-input
            v-model="configForm.address"
            placeholder="address"
          ></el-input>
        </el-form-item>
        <el-form-item label="auth">
          <el-input v-model="configForm.auth" placeholder="auth"></el-input>
        </el-form-item>
        <el-form-item>
          <a class="tm-btn tm-btn-sm color-green" @click="doConnect"> 连接 </a>
        </el-form-item>
      </el-form>
      <el-form
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
          <el-input v-model="searchForm.size" placeholder="size"></el-input>
        </el-form-item>
        <el-form-item>
          <a
            class="tm-btn tm-btn-sm color-green"
            :class="{ 'tm-disabled': loading }"
            @click="doSearch"
          >
            搜索
          </a>
          <a class="tm-btn tm-btn-sm color-blue mgl-5" @click="toInsert">
            新增
          </a>
        </el-form-item>
      </el-form>
    </div>
    <div class="worker-redis-list worker-scrollbar" v-if="connect.open">
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
        <span class="worker-redis-node" slot-scope="{ node, data }">
          <span>{{ node.label }}</span>
          <span>
            <a
              class="tm-link color-red ft-12 mgl-5"
              size="mini"
              @click="toDelete(data)"
            >
              删除
            </a>
          </span>
        </span>
      </el-tree>
    </div>
    <div class="worker-redis-form worker-scrollbar" v-if="connect.open">
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
          <a v-if="!readonlyOne" class="tm-btn color-green" @click="doSave">
            保存
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
        address: "127.0.0.1:6379",
        cluster: false,
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
      data.size = size;
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
      this.loading = true;
      this.keys(data.pattern, data.size)
        .then((res) => {
          this.loading = false;
          if (res.code != 0) {
            tool.error(res.msg);
          } else {
            let value = res.value || {};
            let keys = value.keys || [];
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
        this.doConnect();
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
  width: 100%;
  margin: 0px;
  padding: 0px;
  position: relative;
}
.worker-redis-wrap .worker-redis-list {
  height: calc(100% - 150px);
  width: calc(100% - 500px);
  max-width: 600px;
  min-width: 300px;
  margin: 10px;
  padding: 0px;
  position: relative;
  float: left;
  overflow-x: hidden !important;
}
.worker-redis-wrap .worker-redis-node {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-size: 14px;
  padding-right: 8px;
}
.worker-redis-wrap .worker-redis-form {
  height: calc(100% - 100px);
  width: 400px;
  margin: 0px;
  padding: 0px;
  position: relative;
  float: left;
  padding: 10px;
  overflow: auto;
}
.worker-redis-wrap .el-tree {
  /* border: 1px solid #f3f3f3; */
  border-bottom: 0px;
}
.worker-redis-wrap .el-tree-node__content {
  border-bottom: 1px dotted #696969;
}
</style>
