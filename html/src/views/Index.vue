<template>
  <div class="index-page">
    <WorkerBox></WorkerBox>
  </div>
</template>

<script>
import server from "@/server";
import tool from "@/tool";
import source from "@/source";

import WorkerBox from "@/views/WorkerBox";
export default {
  components: { WorkerBox },
  data() {
    return {
      tool,
      source,
    };
  },
  watch: {},
  computed: {},
  methods: {
    init() {},
  },
  mounted() {
    this.init();
  },
};
</script>

<style>
.index-page {
  height: 100%;
  width: 100%;
  margin: 0px;
  padding: 0px;
  position: relative;
}
</style>
