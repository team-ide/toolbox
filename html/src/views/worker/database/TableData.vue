<template>
  <tm-layout height="100%">
    <tm-layout height="20px">
      <div class="pdlr-10" v-if="database != null && table != null">
        <div class="ft-12" style="line-height: 19px">
          Database：<span class="mgr-10 color-orange">{{ database.name }}</span>
          Table：<span class="mgr-10 color-orange">{{ table.name }}</span>
        </div>
      </div>
      <el-divider class="mg-0"></el-divider>
    </tm-layout>
    <tm-layout height="120px">
      <tm-layout width="400px">
        <ul
          class="part-box worker-scrollbar"
          v-if="database != null && table != null"
        >
          <template v-for="(one, index) in form.wheres">
            <li :key="index">
              <input v-model="one.checked" type="checkbox" />
              <template v-if="one.sqlConditionalOperation == 'custom'">
              </template>
              <template v-else>
                <select
                  v-model="one.name"
                  @change="initInputWidth"
                  class="part-form-input"
                >
                  <option :value="null" text="请选择">请选择</option>
                  <template v-for="(one, index) in table.columns">
                    <option
                      :key="index"
                      :value="one.name"
                      :text="one.name + '&nbsp;'"
                    >
                      {{ one.name }}
                      <template v-if="tool.isNotEmpty(one.comment)">
                        （{{ one.comment }}）
                      </template>
                    </option>
                  </template>
                </select>
              </template>
              <select
                v-model="one.sqlConditionalOperation"
                @change="initInputWidth"
                class="part-form-input"
              >
                <template
                  v-for="(one, index) in source.enum_map
                    .sqlConditionalOperations"
                >
                  <option :key="index" :value="one.value" :text="one.text">
                    {{ one.text }}
                    <template v-if="tool.isNotEmpty(one.comment)">
                      （{{ one.comment }}）
                    </template>
                  </option>
                </template>
              </select>
              <template
                v-if="
                  one.sqlConditionalOperation == 'is null' ||
                  one.sqlConditionalOperation == 'is not null' ||
                  one.sqlConditionalOperation == 'is empty' ||
                  one.sqlConditionalOperation == 'is not empty'
                "
              >
              </template>
              <template
                v-else-if="
                  one.sqlConditionalOperation == 'between' ||
                  one.sqlConditionalOperation == 'not between'
                "
              >
                <input
                  v-model="one.before"
                  type="text"
                  @input="initInputWidth"
                  @change="initInputWidth"
                  class="part-form-input"
                />
                <span class="mglr-5">,</span>
                <input
                  v-model="one.after"
                  type="text"
                  @input="initInputWidth"
                  @change="initInputWidth"
                  class="part-form-input"
                />
              </template>
              <template v-else-if="one.sqlConditionalOperation == 'custom'">
                <input
                  v-model="one.customSql"
                  type="text"
                  @input="initInputWidth"
                  @change="initInputWidth"
                  class="part-form-input"
                />
              </template>
              <template v-else>
                <input
                  v-model="one.value"
                  type="text"
                  @input="initInputWidth"
                  @change="initInputWidth"
                  class="part-form-input"
                />
              </template>

              <select
                v-model="one.andOr"
                @change="initInputWidth"
                class="part-form-input"
              >
                <option value="AND">AND</option>
                <option value="OR">OR</option>
              </select>
            </li>
          </template>
          <li class="pdl-5">
            <a @click="addWhere" class="color-green tm-link mgr-10">添加条件</a>
            <a @click="doSearch" class="color-green tm-link mgr-10">查询</a>
          </li>
        </ul>
      </tm-layout>
      <tm-layout-bar right></tm-layout-bar>
      <tm-layout width="400px">
        <ul
          class="part-box worker-scrollbar"
          v-if="database != null && table != null"
        >
          <li></li>
        </ul>
      </tm-layout>
      <tm-layout-bar right></tm-layout-bar>
      <tm-layout>
        <ul
          class="part-box worker-scrollbar"
          v-if="database != null && table != null"
        >
          <li></li>
        </ul>
      </tm-layout>
    </tm-layout>
    <tm-layout-bar bottom></tm-layout-bar>
    <tm-layout height="20px">
      <div class="ft-12 pdl-10" v-if="database != null && table != null">
        <a class="color-red tm-link mgr-10">删除</a>
        <a class="color-green tm-link mgr-10">保存修改</a>
        <a class="color-blue tm-link mgr-10">新增</a>
      </div>
    </tm-layout>
    <tm-layout height="auto" v-loading="datas_loading">
      <div
        class="worker-database-data-list"
        v-if="database != null && table != null"
      >
        <el-table
          :data="datas"
          :border="true"
          height="100%"
          style="width: 100%"
          size="mini"
        >
          <el-table-column width="70">
            <template slot-scope="scope" label="">
              <input
                :checked="selects.indexOf(scope.row) >= 0"
                type="checkbox"
              />
              <template v-if="updates.indexOf(scope.row) >= 0">
                <i class="mdi mdi-text-box-outline"></i>
              </template>
              <template v-if="inserts.indexOf(scope.row) >= 0">
                <i class="mdi mdi-text-box-plus-outline"></i>
              </template>
            </template>
          </el-table-column>
          <template v-for="(column, index) in form.columns">
            <template v-if="column.checked">
              <el-table-column
                :key="index"
                :prop="column.name"
                :label="column.name"
                width="150"
              >
                <template slot-scope="scope">
                  <div class="">
                    <input
                      v-model="scope.row[column.name]"
                      :placeholder="
                        scope.row[column.name] == null ? 'null' : ''
                      "
                      type="text"
                    />
                  </div>
                </template>
              </el-table-column>
            </template>
          </template>
        </el-table>
      </div>
    </tm-layout>
    <tm-layout height="30px">
      <div
        class="ft-12 pdt-2 text-center worker-database-data-pagination"
        v-if="database != null && table != null && total > 0"
      >
        <el-pagination
          small
          @size-change="handleSizeChange"
          @current-change="handleCurrentChange"
          :current-page="pageIndex"
          :page-sizes="[10, 50, 100, 200, 500]"
          :page-size="pageSize"
          layout="total, sizes, prev, pager, next, jumper"
          :total="total"
        >
        </el-pagination>
      </div>
    </tm-layout>
    <tm-layout height="25px">
      <div
        class="ft-12 pdl-10"
        v-if="database != null && table != null && sql != null"
      >
        <div style="line-height: 25px">
          <span>{{ sql }}</span>
        </div>
      </div>
    </tm-layout>
  </tm-layout>
</template>

<script>
import server from "@/server";
import tool from "@/tool";
import source from "@/source";

export default {
  components: {},
  props: [],
  data() {
    return {
      tool,
      source,
      config: null,
      database: null,
      table: null,
      datas_loading: false,
      datas: null,
      sql: null,
      params: null,
      inserts: [],
      updates: [],
      selects: [],
      pageSize: 100,
      pageIndex: 1,
      total: 0,
      form: {
        wheres: [],
        orders: [],
        columns: [],
      },
    };
  },
  watch: {},
  computed: {},
  methods: {
    handleSizeChange(pageSize) {
      this.pageSize = pageSize;
      this.doSearch();
    },
    handleCurrentChange(pageIndex) {
      this.pageIndex = pageIndex;
      this.doSearch();
    },
    initInputWidth() {
      this.$nextTick(() => {
        let es = this.$el.getElementsByClassName("part-form-input");
        if (es) {
          es.forEach((one) => {
            this.initWidth(one);
          });
        }
      });
    },
    initWidth(event) {
      let target = event;
      if (event.target) {
        target = event.target;
      }
      let value = target.value;
      tool.initInputWidth(event);
    },
    getConfig() {
      return this.config;
    },
    loadDatas(data) {
      data.config = this.getConfig();
      return server.database.datas(data);
    },
    addWhere() {
      let where = {
        checked: true,
        name: null,
        value: null,
        before: null,
        after: null,
        sqlConditionalOperation: "=",
        andOr: "and",
      };
      let column = null;
      if (this.table.columns.length > 0) {
        this.table.columns.forEach((one) => {
          if (column != null) {
            return;
          }
          let find = false;
          this.form.wheres.forEach((w) => {
            if (w.name == one.name) {
              find = true;
            }
          });
          if (find) {
            return;
          }
          column = one;
        });
      }
      if (column == null) {
      } else {
        where.name = column.name;
      }

      this.form.wheres.push(where);
      this.initInputWidth();
    },
    doSearch() {
      let wheres = [];
      let orders = [];
      let columns = [];
      this.form.wheres.forEach((where) => {
        if (!where.checked) {
          return;
        }
        if (tool.isEmpty(where.name)) {
          return;
        }
        wheres.push(where);
      });
      this.form.orders.forEach((order) => {
        if (!order.checked) {
          return;
        }
        orders.push(order);
      });
      this.form.columns.forEach((column) => {
        if (!column.checked) {
          return;
        }
        columns.push(column);
      });
      let data = {};
      data.database = this.database.name;
      data.table = this.table.name;
      data.wheres = wheres;
      data.orders = orders;
      data.columns = columns;
      data.pageIndex = this.pageIndex;
      data.pageSize = this.pageSize;
      this.datas_loading = true;

      this.datas = null;
      this.sql = null;
      this.total = 0;
      this.params = null;
      this.updates = [];
      this.inserts = [];
      this.selects = [];

      this.loadDatas(data)
        .then((res) => {
          this.datas_loading = false;
          if (res.code != 0) {
            tool.error(res.msg);
          }
          let value = res.value || {};
          this.datas = value.datas;
          this.sql = value.sql;
          this.total = Number(value.total || 0);
          this.params = value.params;
        })
        .catch(() => {
          this.datas_loading = false;
        });
    },
    init(config, database, table) {
      this.config = config;
      this.database = database;
      this.table = table;
      this.result = null;
      this.form.wheres = [];
      this.form.orders = [];
      this.form.columns = [];
      table.columns.forEach((one) => {
        let column = Object.assign({}, one);
        column.checked = true;
        this.form.columns.push(column);
      });
    },
  },
  mounted() {},
};
</script>

<style>
.worker-database-data-list {
  width: 100%;
  height: 100%;
}
.worker-database-data-list .el-table .el-table__cell {
  padding: 0px;
}
.worker-database-data-list .el-table .cell {
  white-space: nowrap;
}
.worker-database-data-list .el-table tbody .cell {
  padding-left: 0px !important;
}
.worker-database-data-list .el-table input {
  color: #fff;
  border: 0px dashed transparent;
  background-color: transparent;
  width: 100%;
  padding: 0px 5px;
  box-sizing: border-box;
  outline: none;
  font-size: 12px;
}
.worker-database-data-pagination .el-pagination {
  color: #929292;
}
.worker-database-data-pagination .el-pagination .el-input__inner {
  height: 22px !important;
  line-height: 22px !important;
}
.worker-database-data-pagination .el-pagination .el-input__icon {
  line-height: 22px !important;
}
.worker-database-data-pagination .el-pagination button:disabled {
  background-color: transparent;
}
.worker-database-data-pagination .el-pagination .btn-next,
.worker-database-data-pagination .el-pagination .btn-prev {
  background-color: transparent;
}
.worker-database-data-pagination .el-pagination .btn-next,
.worker-database-data-pagination .el-pagination .btn-prev {
  color: #929292;
}
.worker-database-data-pagination .el-pager li.btn-quicknext,
.worker-database-data-pagination .el-pager li.btn-quickprev {
  color: #929292;
}
.worker-database-data-pagination .el-pager li {
  background: transparent;
}
</style>
