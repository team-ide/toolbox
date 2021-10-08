<template>
  <div class="app">
    <template v-if="source.served">
      <template v-if="source.login_user != null">
        <router-view></router-view>
      </template>
      <template v-else>
        <Login></Login>
      </template>
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

import Login from "@/views/Login";
export default {
  name: "App",
  components: { Login },
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
