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
        <ul class="part-box" v-if="database != null && table != null">
          <li>
            <input type="checkbox" />
            <select @change="tool.initInputWidth">
              <template v-for="(one, index) in table.columns">
                <option :key="index" :value="one.name">
                  {{ one.name }}
                  <template v-if="tool.isNotEmpty(one.comment)">
                    （{{ one.comment }}）
                  </template>
                </option>
              </template>
            </select>
            <select @change="tool.initInputWidth">
              <template
                v-for="(one, index) in source.enum_map.sqlConditionalOperations"
              >
                <option :key="index" :value="one.value">
                  {{ one.text }}
                  <template v-if="tool.isNotEmpty(one.comment)">
                    （{{ one.comment }}）
                  </template>
                </option>
              </template>
            </select>
            <input
              type="text"
              @input="tool.initInputWidth"
              @change="tool.initInputWidth"
            />
          </li>
        </ul>
      </tm-layout>
      <tm-layout-bar right></tm-layout-bar>
      <tm-layout width="400px">
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
        class="worker-kafka-data-list"
        v-if="database != null && table != null && datas != null"
      >
        <el-table :data="datas" height="100%" style="width: 100%" size="mini">
          <el-table-column prop="partition" label="Partition" width="80">
          </el-table-column>
          <el-table-column prop="offset" label="Offset" width="80">
          </el-table-column>
          <el-table-column prop="key" label="Key" width="110">
          </el-table-column>
          <el-table-column prop="value" label="Value"> </el-table-column>
          <el-table-column width="90">
            <template slot-scope="scope">
              <div class="worker-btn-group">
                <a
                  class="tm-link color-grey-1 ft-15 mgr-5"
                  @click="toShow(scope.row)"
                  title="查看"
                >
                  <i class="mdi mdi-eye"></i>
                </a>
                <a
                  class="tm-link color-orange ft-15 mgr-5"
                  @click="toCommit(scope.row)"
                  title="提交"
                >
                  <i class="mdi mdi-send-check"></i>
                </a>
                <a
                  class="tm-link color-red-3 ft-15 mgr-5"
                  title="删除数据"
                  @click="toDeleteData(scope.row)"
                >
                  <i class="mdi mdi-delete-outline"></i>
                </a>
              </div>
            </template>
          </el-table-column>
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
      datas: null,
      form: {
        wheres: [],
        orders: [],
      },
    };
  },
  watch: {},
  computed: {},
  methods: {
    initWidth(event) {
      let target = event.target;
      let value = target.value;
      tool.initInputWidth(event);
    },
    init(config, database, table) {
      this.config = config;
      this.database = database;
      this.table = table;
      this.datas = null;
      this.form.wheres = [];
      this.form.orders = [];
    },
  },
  mounted() {},
};
</script>

<style>
</style>
