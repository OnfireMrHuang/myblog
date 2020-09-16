import Vue from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import ElementUI from 'element-ui'
import MavonEditor from 'mavon-editor'
import 'mavon-editor/dist/css/index.css'
import 'element-ui/lib/theme-chalk/index.css'
import './mock'

// 引用API文件
import axios from 'axios'
import api from './api/http.js'
// 将API方法绑定到全局
Vue.prototype.$http = axios
Vue.prototype.$api = api

Vue.config.productionTip = false
Vue.use(ElementUI)
Vue.use(MavonEditor)

new Vue({
  router,
  store,
  render: h => h(App)
}).$mount('#app')
