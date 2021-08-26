import Vue from 'vue'
import Router from 'vue-router'

Vue.use(Router)

export default new Router({
  routes: [
    {path: "/", component: () => import("../components/Home")},
    {path: "/list", component: () => import("../components/List")},
    {path: "/detail", component: () => import("../components/Detail")},
    {path: "/info", component: () => import("../components/Info")},
  ],
  mode: "history"
})
