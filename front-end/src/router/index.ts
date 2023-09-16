import { useAuthStore } from "@/stores/auth";
import HomeViewVue from "@/views/HomeView.vue";
import { createRouter, createWebHistory } from "vue-router";

const router = createRouter({
  history: createWebHistory(import.meta.env.BASE_URL),
  routes: [
    {
      path: "/",
      name: "home",
      component: HomeViewVue
    },
    {
      path: "/horse",
      name: "horse",
      component: () => import("../views/HorseView.vue")
    },
    {
      path: "/score",
      name: "score",
      component: () => import("../views/ScoreView.vue")
    },
    {
      path: "/about",
      name: "about",
      component: () => import("../views/AboutView.vue")
    },
    {
      path: "/login",
      name: "login",
      component: () => import("../views/LoginView.vue")
    },
    {
      path: "/signup",
      name: "signup",
      component: () => import("../views/SignupView.vue")
    }
  ]
});

router.beforeEach(async (to) => {
  const authStore = useAuthStore();
  if (!authStore.isAuthenticated && to.name !== "login" && to.name !== "signup" && to.name !== "home") {
    return { name: "login" };
  }
});

export default router;
