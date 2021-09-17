<template>
  <div class="worker-database-wrap">
    <tm-layout height="100%">
      <tm-layout height="50px">
        <el-form
          class="pd-10"
          :inline="true"
          :model="configForm"
          size="mini"
          @submit.native.prevent
        >
          <el-form-item label="type">
            <el-select
              v-model="configForm.type"
              placeholder="type"
              style="width: 80px"
            >
              <el-option label="Mysql" value="mysql"> </el-option>
            </el-select>
          </el-form-item>
          <el-form-item label="host">
            <el-input
              v-model="configForm.host"
              placeholder="host"
              style="width: 100px"
            ></el-input>
          </el-form-item>
          <el-form-item label="port">
            <el-input
              v-model="configForm.port"
              placeholder="port"
              style="width: 60px"
            ></el-input>
          </el-form-item>
          <el-form-item label="username">
            <el-input
              v-model="configForm.username"
              placeholder="username"
              style="width: 100px"
            ></el-input>
          </el-form-item>
          <el-form-item label="password">
            <el-input
              v-model="configForm.password"
              placeholder="password"
              style="width: 100px"
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
          <tm-layout width="250px">
            <tm-layout height="40px">
              <div class="pdlr-10">
                <div class="worker-panel-title" v-if="connect.open">
                  Databases列表（{{ connect.form.host }} :
                  {{ connect.form.port }}）
                </div>
              </div>
            </tm-layout>
            <tm-layout height="auto" v-loading="databases_loading">
              <div class="worker-database-database-list" v-if="connect.open">
                <el-table
                  :data="databases"
                  height="100%"
                  style="width: 100%"
                  size="mini"
                >
                  <el-table-column prop="name" label="库名"> </el-table-column>
                  <el-table-column width="60">
                    <template slot-scope="scope">
                      <a
                        class="tm-link color-green ft-15 mgr-2"
                        title="查看库表"
                        @click="selectDatabase(scope.row)"
                      >
                        <i class="mdi mdi-eye-outline"></i>
                      </a>
                      <a
                        class="tm-link color-orange ft-15 mgr-2"
                        title="删除Database"
                        @click="toDeleteDatabase(scope.row)"
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
          <tm-layout width="430px">
            <tm-layout height="100%">
              <tm-layout height="140px">
                <div class="pdlr-10" v-if="connect.open">
                  <div class="worker-panel-title" v-if="connect.open">
                    Tables搜索
                  </div>
                  <el-divider class="mg-0"></el-divider>
                  <el-form
                    class="pdt-10"
                    :inline="true"
                    :model="tableSearchForm"
                    size="mini"
                    @submit.native.prevent
                  >
                    <el-form-item label="database">
                      <el-input
                        v-model="tableSearchForm.database"
                        placeholder="database"
                        style="width: 120px"
                      ></el-input>
                    </el-form-item>
                    <el-form-item label="name">
                      <el-input
                        v-model="tableSearchForm.name"
                        placeholder="name"
                        style="width: 120px"
                      ></el-input>
                    </el-form-item>
                    <el-form-item label="comment">
                      <el-input
                        v-model="tableSearchForm.comment"
                        placeholder="comment"
                        style="width: 120px"
                      ></el-input>
                    </el-form-item>
                    <el-form-item>
                      <a
                        class="tm-btn tm-btn-sm color-green mgl-5"
                        @click="searchTables"
                        :class="{ 'tm-disabled': tables_loading }"
                      >
                        查询
                      </a>
                      <a
                        class="tm-btn tm-btn-sm color-blue mgl-5"
                        @click="toInsertTable"
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
                  <div
                    class="worker-panel-title"
                    v-if="connect.open && tables != null"
                  >
                    Tables列表（{{ tablesDatabase }}）
                  </div>
                </div>
              </tm-layout>
              <tm-layout height="auto" v-loading="tables_loading">
                <div
                  class="worker-database-table-list"
                  v-if="connect.open && tables != null"
                >
                  <el-table
                    :data="tables"
                    height="100%"
                    style="width: 100%"
                    size="mini"
                  >
                    <el-table-column prop="name" label="表名">
                    </el-table-column>
                    <el-table-column prop="comment" label="注释">
                    </el-table-column>
                    <el-table-column width="60">
                      <template slot-scope="scope">
                        <a
                          class="tm-link color-green ft-15 mgr-2"
                          title="查看表"
                          @click="selectTable(scope.row)"
                        >
                          <i class="mdi mdi-eye-outline"></i>
                        </a>
                        <a
                          class="tm-link color-orange ft-15 mgr-2"
                          title="删除Table"
                          @click="toDeleteTable(scope.row)"
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
          <tm-layout-bar right></tm-layout-bar>
          <tm-layout width="auto">
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
              <el-form :model="tableForm" size="lg" @submit.native.prevent>
                <el-form-item label="name">
                  <el-input
                    v-model="tableForm.name"
                    placeholder="name"
                    :readonly="readonlyOne || updateOne"
                  ></el-input>
                </el-form-item>
                <el-form-item>
                  <a
                    v-if="!readonlyOne"
                    class="tm-btn color-green"
                    @click="saveTable"
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
        type: "mysql",
        host: "127.0.0.1",
        port: 3306,
        username: "root",
        password: "123456",
      },
      connect: {
        open: false,
        form: null,
      },
      tableSearchForm: {
        database: "",
        name: "",
        comment: "",
      },
      readonlyOne: true,
      insertOne: true,
      updateOne: true,
      tableForm: {},
      databases_loading: false,
      databases: null,
      tablesDatabase: null,
      tables: null,
      tables_loading: false,
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
      Object.assign(config, this.connect.form);
      config.port = Number(config.port);
      return config;
    },
    loadDatabases() {
      let data = {};
      data.config = this.getConfig();
      return server.database.databases(data);
    },
    loadTables() {
      let data = {};
      data.config = this.getConfig();
      Object.assign(data, this.tableSearchForm);
      return server.database.tables(data);
    },
    searchDatabases() {
      this.databases = null;
      this.tables = null;
      this.databases_loading = true;
      this.loadDatabases()
        .then((res) => {
          this.databases_loading = false;
          if (res.code != 0) {
            tool.error(res.msg);
          } else {
            let value = res.value || {};
            let databases = value.databases || [];
            let datas = [];

            databases.forEach((one) => {
              datas.push(one);
            });
            this.databases = datas;
          }
        })
        .catch(() => {
          this.databases_loading = false;
        });
    },
    selectDatabase(data) {
      this.tableSearchForm.database = data.name;
      this.searchTables();
    },
    searchTables() {
      if (tool.isEmpty(this.tableSearchForm.database)) {
        tool.error("database不能为空！");
        return;
      }
      this.tables = null;
      this.tables_loading = true;
      this.tablesDatabase = this.tableSearchForm.database;
      this.loadTables()
        .then((res) => {
          this.tables_loading = false;
          if (res.code != 0) {
            tool.error(res.msg);
          } else {
            let value = res.value || {};
            let tables = value.tables || [];
            let datas = [];

            tables.forEach((one) => {
              datas.push(one);
            });
            this.tables = datas;
          }
        })
        .catch(() => {
          this.tables_loading = false;
        });
    },
    doConnect() {
      this.connect.open = false;
      tool.trimList(this.expands);
      tool.trimList(this.opens);

      this.$nextTick(() => {
        this.connect.form = Object.assign({}, this.configForm);
        this.connect.open = true;
        this.searchDatabases();
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
    tableNodeClick() {},
    tableCurrentChange() {},
    toInsertTable() {},
    getCacheKey() {
      return "teamide-toolbox-" + this.workerKey;
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
.worker-database-wrap {
  height: 100%;
  margin: 0px;
  padding: 0px;
  position: relative;
}
.worker-database-wrap .worker-database-database-list {
  height: 100%;
}
.worker-database-wrap .worker-database-table-list {
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
