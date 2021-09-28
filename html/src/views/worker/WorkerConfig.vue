<template>
  <el-form
    class="pd-10"
    :inline="true"
    :model="config[workerType]"
    size="mini"
    @submit.native.prevent
  >
    <el-form-item label="">
      <el-select v-model="config.index" placeholder="自定义">
        <el-option label="自定义" :value="null"> </el-option>
        <el-option
          v-for="(one, index) in source[workerType].configs"
          :key="index"
          :label="one.address"
          :value="index"
        >
        </el-option>
      </el-select>
    </el-form-item>
    <template v-if="workerType == 'redis'">
      <el-form-item label="address（多个使用“;”隔开）">
        <el-input
          v-model="config[workerType].address"
          placeholder="address"
          :readonly="config.index != null"
        ></el-input>
      </el-form-item>
      <el-form-item label="auth">
        <el-input
          v-model="config[workerType].auth"
          placeholder="auth"
          :readonly="config.index != null"
        ></el-input>
      </el-form-item>
    </template>
    <template v-if="workerType == 'kafka' || workerType == 'zookeeper'">
      <el-form-item label="address（多个使用“,”隔开）">
        <el-input
          v-model="config[workerType].address"
          placeholder="address"
          :readonly="config.index != null"
        ></el-input>
      </el-form-item>
    </template>
    <el-form-item>
      <a class="tm-btn tm-btn-sm color-green" @click="doConnect">连接</a>
    </el-form-item>
  </el-form>
</template>

<script>
import server from "@/server";
import tool from "@/tool";
import source from "@/source";

export default {
  components: {},
  props: ["workerKey", "workerType", "connect"],
  data() {
    return {
      tool,
      source,
      config: {
        index: null,
        redis: {
          address: "127.0.0.1:6379",
          auth: "",
        },
        zookeeper: {
          address: "127.0.0.1:2181",
        },
        kafka: {
          address: "127.0.0.1:9092",
        },
      },
    };
  },
  watch: {
    "config.index"(value) {
      if (value == null) {
        return;
      }
      let data = source[this.workerType].configs[value];
      for (var key in this.config[this.workerType]) {
        this.config[this.workerType][key] = data[key];
      }
    },
  },
  computed: {},
  methods: {
    getConfig() {
      let config = {};
      Object.assign(config, this.config[this.workerType]);
      return config;
    },
    connectCallback(res) {
      if (res.code == 0) {
        let config = Object.assign(this.connect.config);
        config.index = this.config.index;
        tool.setCache(this.getCacheKey(), JSON.stringify(config));
        this.connect.open = true;
        delete this.doConnectConfig;
      }
    },
    doConnect() {
      let config = this.getConfig();
      this.connect.config = config;
      this.connect.open = false;
      this.connect.title = config.address;
      this.$emit("connect", config, this.connectCallback);
    },
    initConnect() {
      let value = tool.getCache(this.getCacheKey());
      if (tool.isNotEmpty(value)) {
        let data = JSON.parse(value);
        for (var key in this.config[this.workerType]) {
          this.config[this.workerType][key] = data[key];
        }
        this.config.index = data.index;
      }
    },
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
</style>
