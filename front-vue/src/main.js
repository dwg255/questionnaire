import Vue from 'vue'
import App from './App.vue'

import weui from 'weui.js'
import 'weui'
Vue.prototype.$weui = weui

Vue.config.productionTip = false

// 1.1 导入路由的包
import VueRouter from 'vue-router'
// 1.2 安装路由
Vue.use(VueRouter)
    // 1.3 导入自己的 router.js 路由模块
import router from './router.js'
new Vue({
    render: h => h(App),
    router,
}).$mount('#app')