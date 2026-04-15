import { createRouter, createWebHistory } from "vue-router";
import HomeView from "../views/HomeView.vue";
import AboutView from "../views/AboutView.vue";
import LoginView from "@/views/auth/login-view.vue";
const routes = [
  { path: '/', component: HomeView },
  { path: '/about', component: AboutView },
  {
    path: "/login",
    component: LoginView,
  }
]
export const router = createRouter({
  history: createWebHistory(),
  routes,
})
