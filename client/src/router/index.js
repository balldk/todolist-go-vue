import Vue from 'vue'
import VueRouter from 'vue-router'
import List from '../views/List.vue'
import Login from '../views/Login.vue'
import Register from '../views/Register.vue'

Vue.use(VueRouter)

  const routes = [
  {
    path: '/',
    component: List,
    meta: {
      title: "Todo List"
    }
  },
  {
    path: '/login',
    component: Login,
    meta: {
      title: "Login"
    }
  },
  {
    path: '/register',
    component: Register,
    meta: {
      title: "Register"
    }
  }
]

const router = new VueRouter({
  mode: 'history',
  routes
})

export default router
