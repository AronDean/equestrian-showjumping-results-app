<template>
  <nav class="navbar navbar-expand-lg navbar-dark nav-bg">
    <div class="container-fluid">
      <router-link to="/" class="navbar-brand">Equestrian Showjumping Results</router-link>
      <button
        class="navbar-toggler"
        type="button"
        data-bs-toggle="collapse"
        data-bs-target="#navbarSupportedContent"
        aria-controls="navbarSupportedContent"
        aria-expanded="false"
        aria-label="Toggle navigation"
      >
        <span class="navbar-toggler-icon"></span>
      </button>
      <div class="collapse navbar-collapse" id="navbarSupportedContent">
        <ul class="navbar-nav me-auto mb-2 mb-lg-0">
          <li v-if="authStore.isAuthenticated" class="nav-item">
            <router-link
              to="/horse"
              class="nav-link"
              :class="{ active: $route.path === '/horse' }"
              aria-current="page"
              >Pferd</router-link
            >
          </li>
          <li v-if="authStore.isAuthenticated" class="nav-item">
            <router-link
              to="/score"
              class="nav-link"
              :class="{ active: $route.path === '/score' }"
              aria-current="page"
              >Punktzahl</router-link
            >
          </li>
        </ul>

        <div v-if="authStore.isAuthenticated" class="d-flex align-items-center">
          <p class="mb-0 me-3 text-bg">Willkommen, {{ authStore.user }}!</p>
          <button class="btn btn-outline-primary text-bg" @click="logout">Ausloggen</button>
        </div>
        <div v-else>
          <router-link to="/login" class="separate btn btn-outline-primary float-end text-bg"
            >Anmeldung</router-link
          >
          <router-link to="/signup" class="separate btn btn-outline-secondary float-end text-bg"
            >Registrieren</router-link
          >
        </div>
      </div>
    </div>
  </nav>
</template>

<style scoped>
.separate {
  margin-left: 10px;
}

router-link-active {
  color: #0f4162;
}
.nav-bg {
  background-color: #114163;
}

.text-bg {
  color: #f0f3f5;
}
</style>

<script setup lang="ts">
import { useAuthStore } from "@/stores/auth";
import { useRouter } from "vue-router";

const authStore = useAuthStore();
const router = useRouter();

async function logout() {
  authStore.logout();
  await router.push("/login");
}
</script>
