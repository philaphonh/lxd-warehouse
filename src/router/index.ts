import Vue from "vue";
import VueRouter, { Route } from "vue-router";
import VueMeta from "vue-meta";
import Home from "../views/Home.vue";

Vue.use(VueRouter);
Vue.use(VueMeta);

function routeGuard(to: Route, from: Route, next: Function) {
  if (localStorage.getItem("token") !== null) {
    next();
  } else {
    next("/login");
  }
}

const routes = [
  {
    path: "/",
    name: "Home",
    component: Home,
    beforeEnter: routeGuard
  },
  {
    path: "/about",
    name: "About",
    // route level code-splitting
    // this generates a separate chunk (about.[hash].js) for this route
    // which is lazy-loaded when the route is visited.
    component: () =>
      import(/* webpackChunkName: "about" */ "../views/About.vue"),
    beforeEnter: routeGuard
  },
  {
    path: "/dashboard",
    name: "Dashboard",
    component: () => import("../views/Dashboard.vue"),
    beforeEnter: routeGuard
  },
  {
    path: "/products",
    name: "Products Management",
    component: () => import("../views/ProductManager.vue"),
    beforeEnter: routeGuard
  },
  {
    path: "/add-product",
    name: "Add Product",
    component: () => import("../views/ProductEditor.vue"),
    beforeEnter: routeGuard
  },
  {
    path: "/edit-product/:id",
    name: "Edit Product",
    component: () => import("../views/ProductEditor.vue"),
    beforeEnter: routeGuard
  },
  {
    path: "/report",
    name: "Report",
    component: () => import("../views/Report.vue"),
    beforeEnter: routeGuard
  },
  {
    path: "/import",
    name: "Import",
    component: () => import("../views/ImportManager.vue"),
    beforeEnter: routeGuard
  },
  {
    path: "/import-editor",
    name: "Import Products",
    component: () => import("../views/ImportEditor.vue"),
    beforeEnter: routeGuard
  },
  {
    path: "/export",
    name: "Export",
    component: () => import("../views/ExportManager.vue"),
    beforeEnter: routeGuard
  },
  {
    path: "/export-editor",
    name: "Export Editor",
    component: () => import("../views/ExportEditor.vue"),
    beforeEnter: routeGuard
  },
  {
    path: "/extras",
    name: "Extras Management",
    component: () => import("../views/ExtraManager.vue"),
    beforeEnter: routeGuard
  },
  {
    path: "/users",
    name: "Users Management",
    component: () => import("../views/UserManager.vue"),
    beforeEnter: routeGuard
  },
  {
    path: "/add-user",
    name: "Add User",
    component: () => import("../views/UserEditor.vue"),
    beforeEnter: routeGuard
  },
  {
    path: "/edit-user/:username",
    name: "Edit User",
    component: () => import("../views/UserEditor.vue"),
    beforeEnter: routeGuard
  },
  {
    path: "/login",
    name: "Login",
    component: () => import("../views/Login.vue")
  }
];

const router = new VueRouter({
  mode: "history",
  base: process.env.BASE_URL,
  routes
});

export default router;
