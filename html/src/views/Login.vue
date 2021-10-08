<template>
  <div class="login-page">
    <div class="login-content">
      <h3 class="title">Team IDE Toolbox 登录</h3>
      <el-form
        class="login-form"
        :model="form"
        status-icon
        :rules="rules"
        ref="loginForm"
        label-width="100px"
      >
        <el-form-item label="用户名称" prop="name">
          <el-input
            type="text"
            @keyup.enter.native="submit('loginForm')"
            v-model="form.name"
            autocomplete="off"
          >
          </el-input>
        </el-form-item>
        <el-form-item label="用户密钥" prop="auth">
          <el-input
            type="password"
            @keyup.enter.native="submit('loginForm')"
            v-model="form.auth"
            autocomplete="off"
          >
          </el-input>
        </el-form-item>
        <el-form-item>
          <el-button type="primary" @click="submitForm('loginForm')">
            登录
          </el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
import tool from "@/tool/index.js";
import server from "@/server/index.js";
import source from "@/source/index.js";

export default {
  name: "Index",
  components: {},
  data() {
    return {
      tool,
      source,
      form: { name: null, auth: null },
      rules: {
        name: [
          { required: true, message: "请输入用户名称！", trigger: "blur" },
        ],
        auth: [
          { required: true, message: "请输入用户密钥！", trigger: "blur" },
        ],
      },
    };
  },
  watch: {},
  methods: {
    submitForm(formName) {
      this.$refs[formName].validate((valid) => {
        if (valid) {
          let data = Object.assign({}, this.form);
          data.auth = tool.md5(data.auth);
          server.login(data).then((res) => {
            if (res.code == 0) {
              tool.success("登录成功！");
              tool.initSession();
            } else {
              tool.error(res.msg);
            }
          });
        }
      });
    },
    submit(formName) {
      this.submitForm(formName);
    },
  },
  mounted() {
    document.onkeydown = (e) => {
      let key = window.event.keyCode;
      if (key == 13) {
        this.submit();
      }
    };
  },
  beforeCreate() {},
};
</script>

<style >
.login-page {
  position: fixed;
  width: 100%;
  height: 100%;
  background-color: #ddd;
}
.login-page .login-content {
  width: 500px;
  padding: 30px 20px;
  position: fixed;
  left: 50%;
  margin-left: -250px;
  top: 50%;
  margin-top: -190px;
  background: #fff;
}
.login-page .login-content .title {
  margin-bottom: 20px;
  text-align: center;
  color: #ff9066;
}
.login-page .login-form {
  overflow: hidden;
}
</style>
