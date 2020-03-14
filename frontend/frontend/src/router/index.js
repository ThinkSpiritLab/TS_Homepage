import Vue from 'vue'
import VueRouter from 'vue-router'
import Index from "../components/Index";
import Main from "../components/Index/Main";
import Login from "../components/Index/Login";
import Console from "../components/Console";
import accessDeny from "../components/accessDeny";
import Home from "../components/Console/Home";
import AddUser from "../components/Console/Users/AddUser";
import ListUser from "../components/Console/Users/ListUser";
import AddContest from "../components/Console/Contests/AddContest";
import ListContest from "../components/Console/Contests/ListContest";
import AddBulletin from "../components/Console/Bulletins/AddBulletin";
import ListBulletin from "../components/Console/Bulletins/ListBulletin";
import EditBulletin from "../components/Console/Bulletins/EditBulletin";
import AddNews from "../components/Console/News/AddNews";
import EditNews from "../components/Console/News/EditNews";
import ListNews from "../components/Console/News/ListNews";
import EditUserInfo from "../components/Index/Members/EditUserInfo";
import Members from "../components/Index/Members/Members";
import MemberView from "../components/Index/Members/MemberView";

Vue.use(VueRouter);

const routes = [
  { path: '/', redirect: '/index' },
  { path: '/index', component: Index, redirect: '/main', children:[
      { path: '/main', component: Main},
      { path: '/login', component: Login},
      { path: '/indexDeny', component: accessDeny},
      { path: '/editPersonalInfo/:uid', component: EditUserInfo},
      { path: '/members', component: Members},
      { path: '/member/:uid', component: MemberView}
    ] },
  { path: '/console', component: Console, redirect: '/console_home', children:[
      { path: '/console_home', component: Home},
      { path: '/console_addUser', component: AddUser},
      { path: '/console_listUser', component: ListUser},
      { path: '/console_addContest', component: AddContest},
      { path: '/console_listContest', component: ListContest},
      { path: '/console_addBulletin', component: AddBulletin},
      { path: '/console_listBulletin', component: ListBulletin},
      { path: '/console_editBulletin/:bid', component: EditBulletin},
      { path: '/console_addNews', component: AddNews},
      { path: '/console_listNews', component: ListNews},
      { path: '/console_editNews/:nid', component: EditNews},
    ] },
];

const router = new VueRouter({
  mode: 'history',
  base: process.env.BASE_URL,
  routes
});

router.beforeEach((to, from, next) => {
  let index = 0;
  if ((index=to.path.indexOf('console')) !== -1) {
    let Myinfo = JSON.parse(window.sessionStorage.getItem('MyInfo'));
    if (!Myinfo) {
      return next('/login');
    } else if (Myinfo.privilege !== 1 && Myinfo.privilege !== 2 &&
        Myinfo.privilege !== 3 && Myinfo.privilege !== 4) {
      return next('/indexDeny');
    } else {
      next();
    }
  } else if ((index=to.path.indexOf('editPersonalInfo/')) !== -1) {
    let Myinfo = JSON.parse(window.sessionStorage.getItem('MyInfo'));
    if (!Myinfo) {
      return next('/login');
    } else if (Number(to.params.uid)!==Myinfo.uid && Myinfo.privilege !== 1 &&
        Myinfo.privilege !== 2 && Myinfo.privilege !== 3){
      return next('/indexDeny');
    } else {
      next();
    }
  }
  next();
});

export default router
