<template>
  <div class="worker-database-wrap">
    <div class="color-orange ft-20 mgt-40 text-center">
      正在研发中，敬请期待！
    </div>
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
.worker-database-wrap {
  height: 100%;
  width: 100%;
  margin: 0px;
  padding: 0px;
  position: relative;
}
</style>
