import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'
import Atencion from '../components/Atencion.vue'

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    name: 'Home',
    component: Home
  },
  {
    path: '/atencions',
    name: 'atencions',
    component: Atencion
  }
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

export default router
