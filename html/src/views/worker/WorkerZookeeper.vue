<template>
  <div class="worker-zookeeper-wrap">
    <div class="worker-zookeeper-config pd-10 pdb-0">
      <el-form
        :inline="true"
        :model="configForm"
        size="mini"
        @submit.native.prevent
      >
        <el-form-item label="连接地址">
          <el-input v-model="configForm.url" placeholder="连接地址"></el-input>
        </el-form-item>
        <el-form-item>
          <a class="tm-btn tm-btn-sm color-green" @click="doConnect"> 连接 </a>
        </el-form-item>
      </el-form>
    </div>
    <div class="worker-zookeeper-list worker-scrollbar" v-if="connect.open">
      <el-tree
        ref="tree"
        :load="loadNode"
        lazy
        :props="defaultProps"
        :default-expanded-keys="expands"
        node-key="key"
        @node-click="nodeClick"
        @current-change="currentChange"
        :expand-on-click-node="false"
      >
        <span class="worker-zookeeper-node" slot-scope="{ node, data }">
          <template v-if="data.path == '/' && connect.form != null">
            <span>{{ connect.form.url }}</span>
          </template>
          <template v-else>
            <span>{{ node.label }}</span>
          </template>
          <span>
            <a
              class="tm-link color-grey ft-12 mgl-5"
              size="mini"
              @click="toReloadChildren(data)"
            >
              刷新
            </a>
            <a
              class="tm-link color-blue ft-12 mgl-5"
              size="mini"
              @click="toInsert(data)"
            >
              新增
            </a>
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
    <div class="worker-zookeeper-form worker-scrollbar" v-if="connect.open">
      <template v-if="readonlyOne">
        <h3>查看节点</h3>
      </template>
      <template v-else-if="insertOne">
        <h3>新增节点</h3>
      </template>
      <template v-else-if="updateOne">
        <h3>修改节点</h3>
      </template>
      <el-form :model="oneForm" size="lg" @submit.native.prevent>
        <el-form-item label="路径">
          <el-input
            v-model="oneForm.path"
            placeholder="路径"
            :readonly="readonlyOne || updateOne"
          ></el-input>
        </el-form-item>
        <el-form-item label="数据">
          <el-input
            type="textarea"
            v-model="oneForm.data"
            placeholder="数据"
            :autosize="{ minRows: 3, maxRows: 10 }"
            :readonly="readonlyOne"
          ></el-input>
        </el-form-item>
        <el-form-item label="格式化JSON（只做查看，不保存）">
          <el-input
            type="textarea"
            v-model="oneForm.json"
            placeholder="格式化JSON"
            :autosize="{ minRows: 6, maxRows: 10 }"
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
      configForm: { url: "127.0.0.1:2181" },
      connect: {
        open: false,
        form: null,
      },
      readonlyOne: true,
      insertOne: true,
      updateOne: true,
      oneForm: {
        path: null,
        data: null,
        json: null,
      },
      expands: [],
      opens: [],
      defaultProps: {
        children: "children",
        label: "name",
        isLeaf: "leaf",
      },
    };
  },
  watch: {
    "oneForm.data"(value) {
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
    loadChildren(path) {
      let data = {};
      Object.assign(data, this.connect.form);
      data.path = path;
      return server.zookeeper.getChildren(data);
    },
    get(path) {
      let data = {};
      Object.assign(data, this.connect.form);
      data.path = path;
      return server.zookeeper.get(data);
    },
    doSave() {
      let data = {};
      Object.assign(data, this.connect.form);
      Object.assign(data, this.oneForm);
      if (tool.isEmpty(data.path)) {
        tool.error("路径不能为空！");
        return;
      }
      if (data.path.indexOf("//") >= 0 || data.path.endsWith("/")) {
        tool.error("路径格式有误，请检查路径格式！");
        return;
      }
      server.zookeeper.save(data).then((res) => {
        if (res.code != 0) {
          tool.error(res.msg);
        } else {
          tool.success("保存成功");
          let path = data.path;
          if (path.lastIndexOf("/") < path.length - 1) {
            let key = path.substring(0, path.lastIndexOf("/"));
            if (!key.startsWith("/")) {
              key = "/" + key;
            }
            this.reloadChildren(key);
          }
        }
      });
    },
    toDelete(one) {
      if (window.event) {
        window.event.stopPropagation && window.event.stopPropagation();
      }
      let data = {};
      Object.assign(data, this.connect.form);
      data.path = one.path;
      if (tool.isEmpty(data.path)) {
        tool.error("路径不能为空！");
        return;
      }
      tool
        .confirm("将删除节点[" + data.path + "]和子节点，确认删除？")
        .then(() => {
          server.zookeeper.delete(data).then((res) => {
            if (res.code != 0) {
              tool.error(res.msg);
            } else {
              tool.success("删除成功");
              let path = data.path;
              if (path.lastIndexOf("/") < path.length - 1) {
                let key = path.substring(0, path.lastIndexOf("/"));
                if (!key.startsWith("/")) {
                  key = "/" + key;
                }
                this.reloadChildren(key);
              }
            }
          });
        })
        .catch(() => {});
      if (data.path.indexOf("//") >= 0 || data.path.endsWith("/")) {
        tool.error("路径格式有误，请检查路径格式！");
        return;
      }
    },
    toReloadChildren(data) {
      if (window.event) {
        window.event.stopPropagation && window.event.stopPropagation();
      }
      this.reloadChildren(data);
    },
    reloadChildren(key) {
      if (window.event) {
        window.event.stopPropagation && window.event.stopPropagation();
      }
      let node = this.$refs.tree.getNode(key);
      if (node) {
        node.loaded = false;
        node.expand();
      }
    },
    loadNode(node, resolve) {
      if (node.level === 0) {
        resolve([{ name: "/", path: "/", key: "/" }]);
        this.$nextTick(() => {
          let node = this.$refs.tree.getNode("/");
          node.expand();
        });
        return;
      }
      let parent = node.data;
      let path = parent.path;
      this.loadChildren(path).then((res) => {
        if (res.code != 0) {
          tool.error(res.msg);
        } else {
          let value = res.value || {};
          let list = value.children || [];
          let datas = [];
          list.forEach((name) => {
            datas.push({ name: name });
          });
          this.formatDatas(parent, datas);
          resolve(datas);
        }
      });
    },
    formatDatas(parent, datas) {
      datas = datas || [];
      datas.forEach((data) => {
        this.formatData(parent, data);
      });
    },
    formatData(parent, data) {
      data.path = "/" + data.name;
      if (parent != null && parent.path != "/") {
        data.path = parent.path + "/" + data.name;
      }
      data.key = data.path;
    },
    nodeClick() {},
    doConnect() {
      this.connect.open = false;
      tool.trimList(this.expands);
      tool.trimList(this.opens);

      this.toInsert();

      this.$nextTick(() => {
        this.connect.form = Object.assign({}, this.configForm);
        this.connect.open = true;
        tool.setCache(this.getCacheKey(), JSON.stringify(this.connect.form));
      });
    },
    initConnect() {
      let value = tool.getCache(this.getCacheKey());
      if (tool.isNotEmpty(value)) {
        let data = JSON.parse(value);
        Object.assign(this.configForm, data);
        this.doConnect();
      }
    },
    getCacheKey() {
      return "teamide-toolbox-" + this.workerKey;
    },
    toInsert(parent) {
      if (window.event) {
        window.event.stopPropagation && window.event.stopPropagation();
      }
      this.oneForm.path = "/";
      if (parent != null && parent.path != "/") {
        this.oneForm.path = parent.path + "/";
      }
      this.oneForm.data = "";
      this.updateOne = false;
      this.insertOne = true;
      this.readonlyOne = false;
    },
    currentChange(data) {
      this.toUpdate(data);
    },
    toUpdate(data) {
      this.oneForm.path = data.path;
      this.insertOne = false;
      this.updateOne = false;
      this.readonlyOne = true;

      this.get(this.oneForm.path).then((res) => {
        if (res.code != 0) {
          tool.error(res.msg);
        } else {
          let value = res.value || {};
          this.oneForm.data = value.data;
          this.readonlyOne = false;
          this.updateOne = true;
        }
      });
    },
    toInfo(data) {
      this.oneForm.path = data.path;
      this.insertOne = false;
      this.readonlyOne = true;

      this.get(this.oneForm.path).then((res) => {
        if (res.code != 0) {
          tool.error(res.msg);
        } else {
          let value = res.value || {};
          this.oneForm.data = value.data;
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
.worker-zookeeper-wrap {
  height: 100%;
  width: 100%;
  margin: 0px;
  padding: 0px;
  position: relative;
}
.worker-zookeeper-wrap .worker-zookeeper-list {
  height: calc(100% - 150px);
  width: calc(100% - 500px);
  max-width: 600px;
  min-width: 300px;
  margin: 10px;
  padding: 0px;
  position: relative;
  float: left;
}
.worker-zookeeper-wrap .worker-zookeeper-node {
  flex: 1;
  display: flex;
  align-items: center;
  justify-content: space-between;
  font-size: 14px;
  padding-right: 8px;
}
.worker-zookeeper-wrap .worker-zookeeper-form {
  height: calc(100% - 100px);
  width: 400px;
  margin: 0px;
  padding: 0px;
  position: relative;
  float: left;
  padding: 10px;
  overflow: auto;
}
.worker-zookeeper-wrap .el-tree {
  /* border: 1px solid #f3f3f3; */
  border-bottom: 0px;
}
.worker-zookeeper-wrap .el-tree-node__content {
  border-bottom: 1px dotted #696969;
}
</style>
