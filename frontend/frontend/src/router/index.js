import Vue from 'vue'
import VueRouter from 'vue-router'
import Index from "../components/Index";
import Main from "../components/Index/Main";
import Login from "../components/Index/Login";
import Console from "../components/Console";
import accessDeny from "../components/Index/accessDeny";

Vue.use(VueRouter)

const routes = [
  { path: '/', redirect: '/index' },
  { path: '/index', component: Index, children:[
      { path: '/main', component: Main},
      { path: '/login', component: Login},
      { path: '/IndexDeny', component: accessDeny}
    ], redirect:'/main' },
  { path: '/console', component: Console },
]

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
})

router.beforeEach((to, from, next) => {
  if (to.path.indexOf('console') !== -1) {
    let Myinfo = JSON.parse(window.sessionStorage.getItem('MyInfo'))
    if (!Myinfo) {
      return next('/login');
    } else if (Myinfo.privilege !== 1 && Myinfo.privilege !== 2 &&
        Myinfo.privilege !== 3 && Myinfo.privilege !== 4) {
      return next('/IndexDeny');
    } else {
      next();
    }
  }
  next();
})

export default router
