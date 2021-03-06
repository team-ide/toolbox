<template>
  <div class="worker-database-wrap">
    <tm-layout height="100%">
      <tm-layout height="50px">
        <WorkerConfig
          :workerKey="workerKey"
          workerType="database"
          :connect="connect"
          @connect="doConnect"
        ></WorkerConfig>
      </tm-layout>
      <tm-layout-bar bottom></tm-layout-bar>
      <tm-layout height="auto">
        <tm-layout height="100%" width="100%">
          <tm-layout width="300px">
            <tm-layout height="40px">
              <div class="pdlr-10">
                <div class="worker-panel-title" v-if="connect.open">
                  Databases列表（{{ connect.title }}）
                </div>
              </div>
            </tm-layout>
            <tm-layout
              height="auto"
              v-loading="databases_loading"
              @change="treeLayoutChange"
            >
              <div
                class="worker-database-database-list worker-scrollbar"
                style="overflow-x: hidden"
                ref="treeBox"
                v-if="connect.open"
              >
                <el-tree
                  ref="tree"
                  lazy
                  :load="loadNode"
                  :props="treeProps"
                  node-key="key"
                  :expand-on-click-node="false"
                  @node-click="databaseNodeClick"
                >
                  <template slot-scope="{ data }">
                    <span
                      class="worker-box-tree-span"
                      :style="{
                        'margin-left': data.isTable ? '-18px' : 'unset',
                      }"
                    >
                      <span
                        >{{ data.name }}
                        <template v-if="tool.isNotEmpty(data.comment)">
                          <span class="ft-12 mgl-5"
                            >（{{ data.comment }}）</span
                          >
                        </template>
                      </span>
                      <span class="worker-btn-group">
                        <template v-if="data.isDatabase || data.isTable">
                          <a
                            class="tm-link color-grey ft-14 mgr-2"
                            @click="toShowCreateSql(data)"
                            title="查看创建SQL"
                          >
                            <i class="mdi mdi-eye"></i>
                          </a>
                        </template>
                        <template v-if="data.isDatabase || data.isTableFolder">
                          <a
                            class="tm-link color-grey ft-14 mgr-2"
                            @click="toReloadChildren(data)"
                            title="刷新"
                          >
                            <i class="mdi mdi-reload"></i>
                          </a>
                        </template>
                        <template v-if="data.isTableFolder">
                          <a
                            class="tm-link color-blue ft-16 mgr-2"
                            @click="toInsertTable(data)"
                          >
                            <i class="mdi mdi-plus"></i>
                          </a>
                        </template>
                        <template v-if="data.isTable">
                          <a
                            class="tm-link color-green ft-15 mgr-2"
                            @click="toShowTableData(data)"
                          >
                            <i class="mdi mdi-database-outline"></i>
                          </a>
                          <a
                            class="tm-link color-orange ft-15 mgr-2"
                            @click="toDelete(data)"
                          >
                            <i class="mdi mdi-delete-outline"></i>
                          </a>
                        </template>
                      </span>
                    </span>
                  </template>
                </el-tree>
              </div>
            </tm-layout>
          </tm-layout>
          <tm-layout-bar right></tm-layout-bar>
          <tm-layout width="auto">
            <TableData ref="TableData"></TableData>
          </tm-layout>
        </tm-layout>
      </tm-layout>
    </tm-layout>
    <el-dialog
      title="Show Create Database"
      destroy-on-close
      :visible.sync="showCreateDatabase"
      width="900"
    >
      <el-input
        type="textarea"
        class="color-orange"
        v-model="showCreateDatabaseSql"
        :autosize="{ minRows: 6, maxRows: 15 }"
      ></el-input>
    </el-dialog>
    <el-dialog
      title="Show Create Table"
      destroy-on-close
      :visible.sync="showCreateTable"
      width="900"
    >
      <el-input
        type="textarea"
        class="color-orange"
        v-model="showCreateTableSql"
        :autosize="{ minRows: 6, maxRows: 15 }"
      ></el-input>
    </el-dialog>
  </div>
</template>

<script>
import server from "@/server";
import tool from "@/tool";
import source from "@/source";

import WorkerConfig from "./WorkerConfig";
import TableData from "./database/TableData";

export default {
  props: ["workerKey"],
  components: { WorkerConfig, TableData },
  data() {
    return {
      tool,
      source,
      connect: {
        open: false,
        config: null,
        title: null,
      },
      tableForm: {},
      databases_loading: false,
      treeProps: {
        children: "children",
        label: "name",
        isLeaf: "leaf",
      },
      showCreateDatabase: false,
      showCreateDatabaseSql: null,
      showCreateTable: false,
      showCreateTableSql: null,
    };
  },
  watch: {},
  computed: {
    treeStyleObject: function () {
      return {};
    },
  },
  methods: {
    getConfig() {
      let config = {};
      Object.assign(config, this.connect.config);
      config.port = Number(config.port);
      return config;
    },
    loadDatabases() {
      let data = {};
      data.config = this.getConfig();
      return server.database.databases(data);
    },
    checkConnect() {
      let data = {};
      data.config = this.getConfig();
      return server.database.checkConnect(data);
    },
    loadTables(database) {
      let data = {};
      data.config = this.getConfig();
      data.database = database;
      return server.database.tables(data);
    },
    loadTableDetail(database, table) {
      let data = {};
      data.config = this.getConfig();
      data.database = database;
      data.table = table;
      return server.database.tableDetail(data);
    },

    initDatabases(resolve) {
      this.databases_loading = true;
      this.loadDatabases()
        .then((res) => {
          this.databases_loading = false;
          let list = [];
          if (res.code != 0) {
            tool.error(res.msg);
          } else {
            let value = res.value || {};
            list = value.databases || [];
          }
          list.forEach((one) => {
            one.isDatabase = true;
            one.leaf = false;
            one.key = "database-" + one.name;
          });
          this.resolveData(resolve, list);
        })
        .catch(() => {
          this.databases_loading = false;
          this.resolveData(resolve, []);
        });
    },
    initTables(database, resolve) {
      if (database == null || tool.isEmpty(database.name)) {
        tool.error("database不能为空！");
        return;
      }
      this.tables_loading = true;
      this.loadTables(database.name)
        .then((res) => {
          this.tables_loading = false;
          let list = [];
          if (res.code != 0) {
            tool.error(res.msg);
          } else {
            let value = res.value || {};
            list = value.tables || [];
          }
          list.forEach((one) => {
            one.isTable = true;
            one.leaf = true;
            one.key = database.key + "-table-" + one.name;
            one.database = database;
          });
          this.resolveData(resolve, list);
        })
        .catch(() => {
          this.tables_loading = false;
          this.resolveData(resolve, []);
        });
    },
    toShowCreateSql(data) {
      let param = {};
      param.config = this.getConfig();
      if (data.isDatabase) {
        this.showCreateDatabaseSql = null;
        param.database = data.name;
        server.database
          .showCreateDatabase(param)
          .then((res) => {
            if (res.code != 0) {
              tool.error(res.msg);
            } else {
              let value = res.value || {};
              this.showCreateDatabaseSql = value.create;
              this.showCreateDatabase = true;
            }
          })
          .catch(() => {});
      } else if (data.isTable) {
        this.showCreateTableSql = null;
        param.database = data.database.name;
        param.table = data.name;
        server.database
          .showCreateTable(param)
          .then((res) => {
            if (res.code != 0) {
              tool.error(res.msg);
            } else {
              let value = res.value || {};
              this.showCreateTableSql = value.create;
              this.showCreateTable = true;
            }
          })
          .catch(() => {});
      }
    },
    doConnect(config, callback) {
      this.checkConnect()
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
    toReloadChildren(data) {
      if (window.event) {
        window.event.stopPropagation && window.event.stopPropagation();
      }
      this.reloadChildren(data);
    },
    reloadChildren(data) {
      if (window.event) {
        window.event.stopPropagation && window.event.stopPropagation();
      }
      let node = this.$refs.tree.getNode(data.key || data);
      if (node) {
        node.loaded = false;
        node.expand();
      }
    },
    loadNode(node, resolve) {
      if (node.level === 0) {
        this.initDatabases(resolve);
        return;
      }
      if (node.data.isDatabase) {
        let list = [];
        list.push({
          database: node.data,
          key: node.data.key + "-folder-table",
          name: "表",
          leaf: false,
          isTableFolder: true,
        });
        this.resolveData(resolve, list);
        return;
      } else if (node.data.isTableFolder) {
        this.initTables(node.data.database, resolve);
      }
    },
    resolveData(resolve, data) {
      resolve && resolve(data);
      this.initTreeWidth();
    },
    treeLayoutChange() {
      this.initTreeWidth();
    },
    initTreeWidth() {
      // setTimeout(() => {
      //   this.$nextTick(() => {
      //     tool.initTreeWidth(this.$refs.tree, this.$refs.treeBox);
      //   });
      // }, 100);
    },
    databaseNodeClick(data, node, nodeView) {
      let nowTime = new Date().getTime();
      let clickTime = node.clickTime;
      node.clickTime = nowTime;
      if (clickTime) {
        let timeout = nowTime - clickTime;
        if (timeout < 300) {
          node.clickTime = null;
          this.databaseNodeDbClick(data, node, nodeView);
        }
      } else {
      }
    },
    databaseNodeDbClick(data, node, nodeView) {
      if (data.isDatabase) {
        if (node.expanded) {
          node.collapse();
        } else {
          node.expand();
        }
      } else if (data.isTableFolder) {
        if (node.expanded) {
          node.collapse();
        } else {
          node.expand();
        }
      }
    },
    toShowTableData(table) {
      let database = table.database;
      this.loadTableDetail(database.name, table.name)
        .then((res) => {
          if (res.code != 0) {
            tool.error(res.msg);
          } else {
            let value = res.value || {};
            let table = value.table;
            this.$refs.TableData.init(this.getConfig(), database, table);
          }
        })
        .catch(() => {});
    },
    init() {},
  },
  mounted() {
    this.init();
  },
};
</script>

<style>
.worker-database-wrap {
  height: 100%;
  margin: 0px;
  padding: 0px;
  position: relative;
}
.worker-database-wrap .worker-database-database-list {
  height: 100%;
}
.worker-database-wrap .el-tree {
  /* border: 1px solid #f3f3f3; */
  border-bottom: 0px;
}
.worker-database-wrap .el-tree-node__content {
  border-bottom: 1px dotted #696969;
}
</style>
