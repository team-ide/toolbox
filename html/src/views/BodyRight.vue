<template>
  <div class="body-right-wrap"></div>
</template>

<script>
import server from "@/server";
import tool from "@/tool";
import source from "@/source";

export default {
  components: {},
  data() {
    return {
      tool,
      source,
    };
  },
  watch: {},
  computed: {
    treeStyleObject: function () {
      return {};
    },
  },
  methods: {
    init() {},
  },
  mounted() {
    this.init();
  },
};
</script>

<style>
.body-right-wrap {
  height: 100%;
  width: 100%;
  margin: 0px;
  padding: 0px;
  position: relative;
}
</style>
