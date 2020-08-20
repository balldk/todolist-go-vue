import Vue from 'vue'
import Buefy from 'buefy'
import Axios from 'axios'
import App from './App.vue'
import router from './router'
import store from './store'

import 'buefy/dist/buefy.css'
import '@mdi/font/css/materialdesignicons.css'

Vue.prototype.$axios = Axios

Vue.use(Buefy)
new Vue({
  router,
  store,
  watch: {
    $route(to, from) {
      document.title = to.meta.title || "Todo List"
    }
  },
  render: function (h) { return h(App) }
}).$mount('#app')
