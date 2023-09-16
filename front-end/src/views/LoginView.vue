<template>
  <div class="container py-4">
    <header class="pb-3 mb-4 border-bottom">
      <h1 class="d-flex align-items-center text-dark text-decoration-none">
        Melden Sie sich auf der Website an
      </h1>
    </header>
    <div class="p-5 mb-4 bg-light rounded-3">
      <div class="container-fluid py-5">
        <form @submit.prevent="login">
          <div class="form-floating mb-3">
            <input
              type="email"
              class="form-control"
              id="floatingInput"
              placeholder="name@example.com"
              v-model="email"
              required
            />
            <label for="floatingInput">E-Mail-Adresse</label>
          </div>
          <div class="form-floating mb-3">
            <input
              type="password"
              class="form-control"
              id="floatingPassword"
              placeholder="Password"
              v-model="password"
            />
            <label for="floatingPassword">Passwort</label>
          </div>
          <div class="mb-3 form-check">
            <input
              @change="togglePasswordVisibility"
              type="checkbox"
              class="form-check-input"
              id="exampleCheck1"
            />
            <label class="form-check-label" for="exampleCheck1">Passwort anzeigen</label>
          </div>
          <button type="submit" class="btn btn-primary">Anmeldung</button>
        </form>
      </div>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent, ref } from "vue";
import { useAuthStore } from "@/stores/auth";
import { useRouter } from "vue-router";
import { useToast } from "vue-toast-notification";
import "vue-toast-notification/dist/theme-sugar.css";

export default defineComponent({
  setup() {
    const authStore = useAuthStore();
    const router = useRouter();

    const email = ref("");
    const password = ref("");
    const showPassword = ref(false);

    async function login() {
      const isSuccess = await authStore.login(email.value, password.value);
      const toast = useToast();
      if (isSuccess) {
        toast.open({
          message: "Login erfolgreich",
          type: "success"
        });
        await router.push("/");
      } else {
        console.log("login failed");
        toast.open({
          message: "Login fehlgeschlagen",
          type: "error"
        });
      }
    }

    function togglePasswordVisibility() {
      showPassword.value = !showPassword.value;
      const passwordInput = document.getElementById("floatingPassword") as any;
      if (showPassword.value) {
        passwordInput.type = "text";
      } else {
        passwordInput.type = "password";
      }
    }

    return {
      email,
      password,
      login,
      togglePasswordVisibility
    };
  }
});
</script>
