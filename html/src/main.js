import Vue from 'vue'
import App from './App.vue'

import router from '@/router'

Vue.config.productionTip = false

import '@mdi/font/css/materialdesignicons.css'
import 'material-design-icons-iconfont/dist/material-design-icons.css'

import ElementUI from 'element-ui'
import 'element-ui/lib/theme-chalk/index.css'
Vue.use(ElementUI)

import CoosUI from 'coos-ui'
Vue.use(CoosUI)

import {
  PopupManager
} from 'element-ui/lib/utils/popup'

CoosUI.getNextZIndex = function () {
  return PopupManager.nextZIndex();
};


import server from "@/server/index.js";
import tool from "@/tool/index.js";

Vue.prototype.server = server;
Vue.prototype.tool = tool;

if (window.localStorage) {
  new Vue({
    router,
    render: h => h(App),
    mounted() {
      document.getElementById('app-loading-box').setAttribute('class', 'app-loading-box hide');
      window.onunload = function () {
        tool.close();
      };
      window.onbeforeunload = function () {
        tool.close();
      };
    }
  }).$mount('#app')
  tool.open();
} else {
  document.getElementById('app-loading-text').innerHTML = '暂不支持该浏览器，请使用其它浏览器！（推荐使用谷歌内核浏览器）';
  document.getElementById('app-loading-img').style.display = 'none';
}