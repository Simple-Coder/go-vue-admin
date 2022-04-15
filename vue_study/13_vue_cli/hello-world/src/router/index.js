import Vue from 'vue'
import VueRouter from 'vue-router'
import Home from '../views/Home.vue'

const xd = () => import('../views/xd.vue')
const NotFound = () => import('../views/NotFound.vue')
const User = () => import('../views/User.vue')

Vue.use(VueRouter)

const routes = [
  {
    path: '/',
    redirect: 'home'
  },
  {
    path: '/home',
    name: 'Home',
    component: Home,
    children: [
      {
        path: 'about',
        name: 'About',
        // route level code-splitting
        // this generates a separate chunk (about.[hash].js) for this route
        // which is lazy-loaded when the route is visited.
        component: () => import(/* webpackChunkName: "about" */ '../views/About.vue')
      },
      {
        path: 'xd',
        name: 'xd',
        component: xd
      },
      {
        path: 'user/:id',
        name: 'user',
        component: User
      }
    ]
  },
  {
    path: '/404',
    name: '404',
    component: NotFound
  },
  {
    path: '*',
    redirect: '/404'
  }

]

const router = new VueRouter({
  routes
})

export default router
