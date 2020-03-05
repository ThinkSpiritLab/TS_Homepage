import Vue from 'vue'
import App from './App.vue'
import router from './router'
// import store from './store'
import './plugins/element.js'
import './assets/css/global.css'
import './assets/fonts/iconfont.css'
import axios from 'axios'
axios.interceptors.request.use(function (config) {
  config.headers.Authorization = window.sessionStorage.getItem('Authorization');
  return config;
});

Vue.prototype.$http = axios;
// axios.defaults.baseURL = 'http://localhost:8000/api/v1/';

Vue.config.productionTip = false

new Vue({
  router,
  render: h => h(App)
}).$mount('#app')
