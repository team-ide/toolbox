<template>
  <div>
    <h1 class="text-center mgt-30 color-orange">Team IDE Toolbox</h1>
  </div>
</template>

<script>
import server from "@/server";
import tool from "@/tool";
import source from "@/source";

export default {
  name: "App",
  components: {
  },
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
</style>
