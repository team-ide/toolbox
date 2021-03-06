<template>
  <div class="worker-zookeeper-wrap">
    <tm-layout height="100%">
      <tm-layout height="50px">
        <WorkerConfig
          :workerKey="workerKey"
          workerType="zookeeper"
          :connect="connect"
          @connect="doConnect"
        ></WorkerConfig>
      </tm-layout>
      <tm-layout-bar bottom></tm-layout-bar>
      <tm-layout height="auto">
        <tm-layout height="100%" width="100%">
          <tm-layout width="auto">
            <tm-layout height="45px">
              <div class="worker-panel-title pdlr-10" v-if="connect.open">
                节点信息（{{ connect.title }}）
              </div>
              <el-divider class="mg-0"></el-divider>
            </tm-layout>
            <tm-layout height="auto">
              <div
                class="worker-zookeeper-list worker-scrollbar"
                style="overflow-x: hidden"
                ref="treeBox"
                v-if="connect.open"
              >
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
                  <span
                    class="worker-box-tree-span"
                    slot-scope="{ node, data }"
                  >
                    <template v-if="data.path == '/' && connect.form != null">
                      <span>/</span>
                    </template>
                    <template v-else>
                      <span>{{ node.label }}</span>
                    </template>
                    <div class="worker-btn-group">
                      <a
                        class="tm-link color-grey ft-14 mgr-2"
                        @click="toReloadChildren(data)"
                      >
                        <i class="mdi mdi-reload"></i>
                      </a>
                      <a
                        class="tm-link color-blue ft-16 mgr-2"
                        @click="toInsert(data)"
                      >
                        <i class="mdi mdi-plus"></i>
                      </a>
                      <a
                        class="tm-link color-orange ft-15 mgr-2"
                        @click="toDelete(data)"
                      >
                        <i class="mdi mdi-delete-outline"></i>
                      </a>
                    </div>
                  </span>
                </el-tree>
              </div>
            </tm-layout>
          </tm-layout>
          <tm-layout-bar left></tm-layout-bar>
          <tm-layout width="400px">
            <div class="pdlr-10" v-if="connect.open">
              <div class="worker-panel-title">
                <template v-if="readonlyOne">
                  <span>查看节点</span>
                </template>
                <template v-else-if="insertOne">
                  <span>新增节点</span>
                </template>
                <template v-else-if="updateOne">
                  <span>修改节点</span>
                </template>
              </div>
              <el-divider class="mg-0"></el-divider>
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
      Object.assign(data, this.connect.config);
      data.path = path;
      return server.zookeeper.getChildren(data);
    },
    get(path) {
      let data = {};
      Object.assign(data, this.connect.config);
      data.path = path;
      return server.zookeeper.get(data);
    },
    doSave() {
      let data = {};
      Object.assign(data, this.connect.config);
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
      Object.assign(data, this.connect.config);
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
      this.loadChildren(path)
        .then((res) => {
          if (res.code != 0) {
            tool.error(res.msg);
            resolve([]);
            this.initTreeWidth();
          } else {
            let value = res.value || {};
            let list = value.children || [];
            let datas = [];
            list.forEach((name) => {
              datas.push({ name: name });
            });
            this.formatDatas(parent, datas);
            resolve(datas);
            this.initTreeWidth();
          }
        })
        .catch((e) => {});
    },
    initTreeWidth() {
      // setTimeout(() => {
      //   this.$nextTick(() => {
      //     tool.initTreeWidth(this.$refs.tree, this.$refs.treeBox);
      //   });
      // }, 100);
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
    doConnect(config, callback) {
      tool.trimList(this.expands);
      tool.trimList(this.opens);

      this.toInsert();
      this.get("/")
        .then((res) => {
          if (res.code != 0) {
            tool.error(res.msg);
          }
          callback(res);
        })
        .catch((e) => {
          callback(e);
        });
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
    init() {},
  },
  mounted() {
    this.init();
  },
};
</script>

<style>
.worker-zookeeper-wrap {
  height: 100%;
  margin: 0px;
  padding: 0px;
  position: relative;
}
.worker-zookeeper-wrap .worker-zookeeper-list {
  height: 100%;
}
.worker-zookeeper-wrap .el-tree {
  /* border: 1px solid #f3f3f3; */
  border-bottom: 0px;
}
.worker-zookeeper-wrap .el-tree-node__content {
  border-bottom: 1px dotted #696969;
}
</style>
