<template>
  <div class="worker-resource-wrap">
    <tm-layout height="100%">
      <tm-layout height="100%">
        <tm-layout height="90px">
          <div class="pdlr-10">
            <div class="worker-panel-title">搜索</div>
            <el-divider class="mg-0"></el-divider>
            <el-form
              class="pdt-10"
              :inline="true"
              :model="searchForm"
              size="mini"
              @submit.native.prevent
            >
              <el-form-item label="搜索">
                <el-input
                  v-model="searchForm.name"
                  placeholder="name"
                ></el-input>
              </el-form-item>
              <el-form-item>
                <a
                  class="tm-btn tm-btn-sm color-grey mgl-5"
                  :class="{ 'tm-disabled': loading }"
                  @click="doSearch"
                >
                  搜索
                </a>
                <a
                  class="tm-btn tm-btn-sm color-green mgl-5"
                  @click="doCreateDir"
                >
                  新建文件夹
                </a>
              </el-form-item>
              <el-form-item>
                <el-upload
                  class="upload-demo"
                  :action="source.ROOT_URL + 'server/upload/resource'"
                  :on-success="handleSuccess"
                  :limit="1"
                  :show-file-list="false"
                  :file-list="fileList"
                  accept="*"
                  list-type="text"
                  :data="{ relPath: relPath }"
                >
                  <el-link slot="trigger" style="line-height: initial">
                    <a class="tm-btn tm-btn-sm color-blue"> 上传 </a>
                  </el-link>
                </el-upload>
              </el-form-item>
            </el-form>
          </div>
        </tm-layout>
        <tm-layout-bar bottom></tm-layout-bar>
        <tm-layout height="40px">
          <div class="pdlr-10">
            <div class="worker-panel-title pdt-10">
              <el-breadcrumb separator="/">
                <template v-for="(one, index) in parents">
                  <el-breadcrumb-item :key="index">
                    <a class="color-green" @click="toShowFolder(one)">
                      {{ one.name }}
                    </a>
                  </el-breadcrumb-item>
                </template>
              </el-breadcrumb>
            </div>
          </div>
        </tm-layout>
        <tm-layout height="auto" v-loading="loading">
          <div class="worker-resource-list">
            <el-table
              :data="files"
              height="100%"
              style="width: 100%"
              size="mini"
            >
              <el-table-column label="文件">
                <template slot-scope="scope">
                  <template v-if="scope.row.isDir">
                    <a
                      class="tm-link color-orange ft-14"
                      :title="'打开目录：' + scope.row.name"
                      @click="toShowFolder(scope.row)"
                    >
                      {{ scope.row.name }}
                    </a>
                  </template>
                  <template v-else>
                    <a
                      class="tm-link color-green ft-14"
                      :title="'下载：' + scope.row.name"
                      :href="
                        source.ROOT_URL +
                        'server/download/resource?relPath=' +
                        scope.row.relPath
                      "
                    >
                      {{ scope.row.name }}
                    </a>
                  </template>
                </template>
              </el-table-column>
              <el-table-column prop="size" label="size" width="60">
              </el-table-column>
              <el-table-column width="60">
                <template slot-scope="scope">
                  <div class="worker-btn-group">
                    <a
                      class="tm-link color-green ft-15 mgr-2"
                      title="下载"
                      @click="toDownload(scope.row)"
                    >
                      <i class="mdi mdi-playlist-edit"></i>
                    </a>
                    <a
                      class="tm-link color-orange ft-15 mgr-2"
                      title="删除"
                      @click="toDelete(scope.row)"
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
    </tm-layout>
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
      loading: false,
      searchForm: { name: "" },
      files: null,
      relPath: null,
      fileList: [],
      parents: [],
    };
  },
  watch: {
    relPath() {
      this.parents = [];
      let path = this.relPath;
      let paths = path.split("/");
      if (path.indexOf("\\") > 0) {
        paths = path.split("\\");
      }
      this.parents.push({ relPath: "", name: ".." });
      let parentPath = "";
      paths.forEach((one) => {
        let relPath = one;
        if (parentPath != "") {
          relPath = parentPath + "/" + one;
        }
        this.parents.push({ relPath: relPath, name: one });

        parentPath = relPath;
      });
      console.log(this.parents);
      this.doSearch();
    },
  },
  computed: {},
  methods: {
    handleSuccess(file) {
      this.fileList = [];
      if (file.code == 0) {
        tool.success("上传成功");
        this.doSearch();
      } else {
        tool.error(file.msg);
      }
    },
    loadList(path, name) {
      let data = {};
      data.name = name;
      data.path = path;
      return server.resource.list(data);
    },
    toShowFolder(one) {
      if (window.event) {
        window.event.stopPropagation && window.event.stopPropagation();
      }
      this.relPath = one.relPath;
    },
    toDelete(one) {
      if (window.event) {
        window.event.stopPropagation && window.event.stopPropagation();
      }
      let data = {};
      data.relPath = one.relPath;
      if (tool.isEmpty(data.relPath)) {
        tool.error("relPath不能为空！");
        return;
      }
      tool
        .confirm("将删除[" + one.name + "]，确认删除？")
        .then(() => {
          server.resource.delete(data).then((res) => {
            if (res.code != 0) {
              tool.error(res.msg);
            } else {
              tool.success("删除成功");
              this.reload();
            }
          });
        })
        .catch(() => {});
    },
    doSearch() {
      this.load();
    },
    reload() {
      this.load();
    },
    load() {
      let data = {};
      Object.assign(data, this.searchForm);
      data.relPath = this.relPath;
      this.files = null;
      this.loading = true;
      return new Promise((resolve, reject) => {
        this.loadList(data.relPath, data.name)
          .then((res) => {
            this.loading = false;
            if (res.code != 0) {
              tool.error(res.msg);
            } else {
              let value = res.value || {};
              let files = value.files || [];

              this.files = files;
            }
            resolve && resolve(res);
          })
          .catch((e) => {
            this.loading = false;
            reject && reject(e);
          });
      });
    },
    toUpload() {
      if (window.event) {
        window.event.stopPropagation && window.event.stopPropagation();
      }
    },
    init() {
      this.relPath = "";
      // this.doSearch();
    },
  },
  mounted() {
    this.init();
  },
};
</script>

<style>
.worker-resource-wrap {
  height: 100%;
  margin: 0px;
  padding: 0px;
  position: relative;
}
.worker-resource-wrap .worker-resource-list {
  height: 100%;
}
.worker-resource-wrap .el-tree {
  /* border: 1px solid #f3f3f3; */
  border-bottom: 0px;
}
.worker-resource-wrap .el-tree-node__content {
  border-bottom: 1px dotted #696969;
}
</style>
