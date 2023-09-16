<template>
  <div class="bg-light rounded-3">
    <div class="container-fluid py-2">
      <form @submit.prevent="createOrUpdate(isUpdate, localHorse)">
        <div class="form-floating mb-3">
          <input
            type="text"
            class="form-control"
            id="horseName"
            placeholder="Ferrarie"
            v-model="localHorse.name"
            required
          />
          <label for="horseName">Name</label>
        </div>

        <div class="row">
          <div class="col-6">
            <div class="form-floating mb-3">
              <input
                type="text"
                class="form-control"
                id="horseColor"
                placeholder="Red"
                v-model="localHorse.color"
                required
              />
              <label for="horseColor">Farbe</label>
            </div>
          </div>

          <div class="col-6">
            <div class="form-floating mb-3">
              <input
                type="number"
                class="form-control"
                id="horseAge"
                v-model="localHorse.age"
                required
              />
              <label for="horseAge">Alter</label>
            </div>
          </div>
        </div>
        <button v-if="!isUpdate" type="submit" class="btn btn-primary">Create</button>
        <button v-if="isUpdate" type="submit" class="btn btn-primary">Update</button>
      </form>
    </div>
  </div>
</template>

<script lang="ts">
import { defineComponent } from "vue";
import type { Horse } from "@/types/interface";
import api from "@/stores/api";
import { useAuthStore } from "@/stores/auth";
import { useToast } from "vue-toast-notification";
import "vue-toast-notification/dist/theme-sugar.css";

export default defineComponent({
  props: {
    horse: {
      type: Object as () => Horse,
      required: true
    },
    isUpdate: {
      type: Boolean,
      required: true
    }
  },
  data() {
    return {
      localHorse: { ...this.horse }
    };
  },
  methods: {
    async createOrUpdate(isUpdate: boolean, body: Horse) {
      const authStore = useAuthStore();
      let data = JSON.stringify(body);
      const toast = useToast();
      let config = {
        method: isUpdate ? "put" : "post",
        url: isUpdate ? `/horse/${body.id}` : "/horse/",
        headers: {
          Authorization: authStore.accessToken,
          "Content-Type": "application/json"
        },
        data: data
      };

      await api
        .request(config)
        .then((response) => {
          toast.open({
            message: isUpdate ? "Pferd aktualisiert" : "Pferd erstellt",
            type: "success"
          });
          this.$emit("horse-created", response.data.id);
        })
        .catch((error) => {
          toast.open({
            message: "Pferd konnte nicht erstellt werden",
            type: "error"
          });
          console.log(error);
        });
    }
  }
});
</script>
