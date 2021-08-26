// The Vue build version to load with the `import` command
// (runtime-only or standalone) has been set in webpack.base.conf with an alias.
import Vue from 'vue'
import router from "./router";
import ElementUI from 'element-ui';
import 'element-ui/lib/theme-chalk/index.css';
import 'bootstrap/dist/css/bootstrap.min.css'

import App from './App'
import axios from "axios";

Vue.use(ElementUI);

Vue.config.productionTip = false
Vue.prototype.$axios = axios
axios.defaults.baseURL = process.env.API_BASE_URL
axios.defaults.timeout = 10000;
axios.defaults.withCredentials = false
axios.defaults.crossDomain = true
// axios.defaults.headers.post = {
//   'Content-Type': 'application/json'
// } ;
axios.interceptors.response.use(response => {
  // debugger
  return response.data;
}, error => {
  // debugger
  return error
});

/* eslint-disable no-new */
new Vue({
  router,
  axios,
  el: '#app',
  components: {App},
  template: '<App/>'
})
