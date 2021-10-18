<template>
  <tm-layout height="100%">
    <tm-layout height="40px">
      <div class="pdlr-10" v-if="database != null && table != null">
        <div class="worker-panel-title">
          Database：<span class="mgr-10 color-orange">{{ database.name }}</span>
          Table：<span class="mgr-10 color-orange">{{ table.name }}</span>
        </div>
      </div>
    </tm-layout>
    <tm-layout height="200px">
      <tm-layout width="400px">
        <el-divider class="mg-0"></el-divider>
        <ul class="part-box" v-if="database != null && table != null">
          <template v-for="(one, index) in form.wheres">
            <li :key="index">
              <input v-model="one.checked" type="checkbox" />
              <template v-if="one.sqlConditionalOperation == 'custom'">
              </template>
              <template v-else>
                <select
                  v-model="one.name"
                  @change="tool.initInputWidth"
                  class="part-form-input"
                >
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
                @change="tool.initInputWidth"
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
                  @input="tool.initInputWidth"
                  @change="tool.initInputWidth"
                  class="part-form-input"
                />
                <span class="mglr-5">,</span>
                <input
                  v-model="one.after"
                  type="text"
                  @input="tool.initInputWidth"
                  @change="tool.initInputWidth"
                  class="part-form-input"
                />
              </template>
              <template v-else-if="one.sqlConditionalOperation == 'custom'">
                <input
                  v-model="one.customSql"
                  type="text"
                  @input="tool.initInputWidth"
                  @change="tool.initInputWidth"
                  class="part-form-input"
                />
              </template>
              <template v-else>
                <input
                  v-model="one.value"
                  type="text"
                  @input="tool.initInputWidth"
                  @change="tool.initInputWidth"
                  class="part-form-input"
                />
              </template>

              <select
                v-model="one.andOr"
                @change="tool.initInputWidth"
                class="part-form-input"
              >
                <option value="and">and</option>
                <option value="or">or</option>
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
        <el-divider class="mg-0"></el-divider>
        <ul class="part-box" v-if="database != null && table != null">
          <li></li>
        </ul>
      </tm-layout>
      <tm-layout-bar right></tm-layout-bar>
      <tm-layout>
        <ul class="part-box" v-if="database != null && table != null">
          <li></li>
        </ul>
      </tm-layout>
    </tm-layout>
    <tm-layout-bar bottom></tm-layout-bar>
    <tm-layout height="auto" v-loading="datas_loading">
      <div
        class="worker-database-data-list"
        v-if="database != null && table != null && result != null"
      >
        <el-table
          :data="result.datas"
          :border="true"
          height="100%"
          style="width: 100%"
          size="mini"
        >
          <template v-for="(column, index) in result.columns">
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
                    :placeholder="scope.row[column.name] == null ? 'null' : ''"
                    type="text"
                  />
                </div>
              </template>
            </el-table-column>
          </template>
        </el-table>
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
      result: null,
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
        return;
      }
      where.name = column.name;

      this.form.wheres.push(where);
      this.initFormInputWidth();
    },
    doSearch() {
      let wheres = [];
      let orders = [];
      let columns = [];
      this.form.wheres.forEach((where) => {
        if (!where.checked) {
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
      data.pageIndex = 1;
      data.pageSize = 100;
      this.datas_loading = true;
      this.loadDatas(data)
        .then((res) => {
          this.datas_loading = false;
          if (res.code != 0) {
            tool.error(res.msg);
          }
          this.result = res.value;
          if (this.result != null) {
            this.result.columns = columns;
          }
        })
        .catch(() => {
          this.datas_loading = false;
        });
    },
    initFormInputWidth() {
      this.$nextTick(() => {
        let es = this.$el.getElementsByClassName("part-form-input");
        if (es) {
          es.forEach((one) => {
            this.initWidth(one);
          });
        }
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
</style>
