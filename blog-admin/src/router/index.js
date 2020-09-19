import Vue from 'vue'
import VueRouter from 'vue-router'
import store from '../store'

const originalPush = VueRouter.prototype.push
VueRouter.prototype.push = function push(location) {
  return originalPush.call(this, location).catch(err => err)
}

Vue.use(VueRouter)

 const routes = [
  {
    path: '/',
    name: 'Home',
    redirect: '/admin',
    meta: {
      requireAuth: true // 添加改字段，表示进入这个路由是需要登陆的
    }
  },
  {
    path: "/login",
    name: "login",
    component: () => import('../views/Login.vue')
  },
  {
    path: "/admin",
    name: "admin",
    component: () => import("../views/Admin.vue"),
    children: [
      {
        path: "/admin/",
        name: "admin",
        component: () => import("../views/Admin-Index.vue")
      },
      {
        path: "article",
        name: "article",
        component: () => import("../views/Article.vue")
      },
      {
        path: 'article/:id',
        name: 'update',
        component: () => import(/* webpackChunkName: 'update' */ '../views/Update.vue')
      },
      {
        path: 'list',
        name: 'articleList',
        component: () => import( '../views/ArticleList.vue'),
      },
      {
        path: "version",
        name: "version",
        component: () => import("../views/Version.vue")
      },
      { 
        path: 'comment',
        name: 'comment',
        component: () => import('../views/Comment.vue'),
      }
    ]
  }
]

const router = new VueRouter({
  routes
})

var storeTemp=store;
// 添加钩子函数，拦截需要登陆验证的页面
router.beforeEach((to,from,next) => {
  // 缓存token，在http后端请求的时候添加拦截验证
  if (!storeTemp.state.Token) {
    storeTemp.commit("saveToken",window.localStorage.Token)
  }
  // 页面的拦截验证
  if (to.meta.requireAuth) {
    if (window.localStorage.Token && window.localStorage.Token.length>=128) {
      next();
    } else {
      next({
        path: '/login',
        query: { redirect: to.fullPath} //将跳转的路由path作为参数，登陆成功后跳转到该路由
      })
    }
  } else {
    next();
  }
})

export default router
