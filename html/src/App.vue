<template>
  <div class="app">
    <template v-if="source.served">
      <router-view></router-view>
    </template>
    <template v-else>
      <div class="text-center pdtb-100 ft-20 color-grey">
        {{
          source.inited
            ? "服务器连接异常，请稍后再试！"
            : "服务器连接中，请稍后..."
        }}
      </div>
    </template>
  </div>
</template>

<script>
import server from "@/server";
import tool from "@/tool";
import source from "@/source";

export default {
  name: "App",
  components: {},
  data() {
    return {
      tool,
      source,
    };
  },
  watch: {},
  methods: {
    init() {},
  },
  mounted() {
    this.init();
  },
};
</script>

<style>
html,
body {
  height: 100%;
  width: 100%;
  margin: 0px;
  padding: 0px;
  position: relative;
}
.app {
  height: 100%;
  width: 100%;
  margin: 0px;
  padding: 0px;
  position: relative;
}
a {
  margin: 0px;
  padding: 0px;
}
</style>
