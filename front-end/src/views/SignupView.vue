<template>
  <div class="container py-2">
    <header class="pb-3 mb-4 border-bottom">
      <h1 class="d-flex align-items-center text-dark text-decoration-none">Ein Konto erstellen</h1>
    </header>
    <div class="p-2 mb-2 bg-light rounded-3">
      <div class="container-fluid py-5">
        <form @submit.prevent="signup">
          <div class="form-floating mb-3">
            <input
              type="text"
              class="form-control"
              id="nameInput"
              placeholder="CR7"
              v-model="name"
              required
            />
            <label for="nameInput">Nutzername</label>
          </div>
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
          <div class="row">
            <div class="col-6">
              <div class="form-floating">
                <input
                  type="password"
                  class="form-control"
                  id="floatingPassword"
                  placeholder="Password"
                  v-model="password"
                />
                <label for="floatingPassword">Passwort</label>
              </div>
            </div>

            <div class="col-6">
              <div class="form-floating">
                <input
                  type="password"
                  class="form-control"
                  id="floatingPassword2"
                  placeholder="Password"
                  v-model="password2"
                />
                <label for="floatingPassword2">Bestätige das Passwort</label>
              </div>
            </div>
          </div>
          <div id="exampleInputPassword1Help" class="form-text mb-3" v-bind:class="helpTextClass">
            {{ passwordMatch }}
          </div>
          <button type="submit" class="btn btn-primary">Submit</button>
        </form>
      </div>
    </div>
  </div>
</template>

<script setup lang="ts">
import router from "@/router";
import api from "@/stores/api";
import { useAuthStore } from "@/stores/auth";
import { ref, computed } from "vue";
import { useToast } from "vue-toast-notification";
import "vue-toast-notification/dist/theme-sugar.css";

const name = ref("");
const email = ref("");
const password = ref("");
const password2 = ref("");

const passwordMatch = computed(() => {
  if (password.value === "" || password2.value === "") {
    return "";
  }
  if (password.value === password2.value) {
    return "Passwörter stimmen überein 😃😃😃";
  } else {
    return "Passwörter stimmen nicht überein 😟😟😟";
  }
});

const helpTextClass = computed(() => {
  if (password.value === "" || password2.value === "") {
    return "";
  }
  if (password.value === password2.value) {
    return "text-success";
  } else {
    return "text-danger";
  }
});

async function signup() {
  if (password.value !== password2.value) return;

  const authStore = useAuthStore();
  let data = JSON.stringify({
    name: name.value,
    email: email.value,
    password: password.value
  });

  let config = {
    method: "post",
    url: "/user/sign-up",
    headers: {
      "Content-Type": "application/json"
    },
    data: data
  };
  let isSuccess = false;
  const toast = useToast();
  await api
    .request(config)
    .then((response) => {
      if (response.status !== 200) return;
      authStore.afterUserCreate(response.data);
      isSuccess = true;
      toast.open({
        message: "Account erfolgreich erstellt",
        type: "success"
      });
    })
    .catch((error) => {
      console.log(error);
      toast.open({
        message: "Account konnte nicht erstellt werden",
        type: "error"
      });
    });

  if (isSuccess) {
    await router.push("/");
  }
}
</script>
