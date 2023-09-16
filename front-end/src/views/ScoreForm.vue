<template>
  <div class="bg-light rounded-3">
    <div class="container-fluid py-2">
      <form @submit.prevent="createScore">
        <div class="row">
          <div class="col-4">
            <label for="horse_id">Pferd</label>
            <select class="form-select mb-3" id="horse_id" v-model="horseId" required>
              <option v-for="horse in horses" :key="horse.id" :value="horse.id">
                {{ horse.name }}
              </option>
            </select>
          </div>
          <div class="col-4">
            <label for="jump_height_id">Hindernisshöhe</label>
            <select class="form-select mb-3" id="jump_height_id" v-model="jumpHeightId" required>
              <option v-for="jumpHeight in jumpHeights" :key="jumpHeight.id" :value="jumpHeight.id">
                {{ jumpHeight.title }}
              </option>
            </select>
          </div>
          <div class="col-2">
            <label for="rank_id">Rang</label>
            <select class="form-select mb-3" id="rank_id" v-model="rankId" required>
              <option v-for="rank in ranks" :key="rank.id" :value="rank.id">
                {{ rank.id }}
              </option>
            </select>
          </div>
          <div class="col-2 position-relative">
            <button
              id="btn"
              type="submit"
              class="btn btn-primary position-absolute top-50 start-50 translate-middle"
            >
              Submit
            </button>
          </div>
        </div>
      </form>
    </div>
  </div>
</template>

<script setup lang="ts">
import api from "@/stores/api";
import { useAuthStore } from "@/stores/auth";
import { ref } from "vue";
import { useToast } from "vue-toast-notification";
import "vue-toast-notification/dist/theme-sugar.css";

const horseId = ref(null);
const rankId = ref(null);
const jumpHeightId = ref(null);

const horses = ref([] as any);
const ranks = ref([] as any);
const jumpHeights = ref([] as any);

const authStore = useAuthStore();
const emit = defineEmits(["score-created"]);

async function GetData(url: string) {
  let config = {
    method: "get",
    url: url,
    headers: {
      Authorization: authStore.accessToken
    }
  };

  let data = [] as any;
  await api
    .request(config)
    .then((response) => {
      data = response.data;
    })
    .catch((error) => {
      console.log(error);
    });

  return data;
}

async function createScore() {
  let data = JSON.stringify({
    horse_id: horseId.value,
    rank_id: rankId.value,
    jump_height_id: jumpHeightId.value
  });
  const toast = useToast();
  let config = {
    method: "post",
    url: "/score/",
    headers: {
      Authorization: authStore.accessToken,
      "Content-Type": "application/json"
    },
    data: data
  };
  let isSuccessful = false;
  await api
    .request(config)
    .then((response) => {
      if (response.status == 200) {
        isSuccessful = true;
        toast.open({
          message: "Score erstellt",
          type: "success"
        });
        emit("score-created", response.data.id);
      }
    })
    .catch((error) => {
      toast.open({
        message: "Score konnte nicht erstellt werden",
        type: "error"
      });
      console.log(error);
    });

  return isSuccessful;
}

async function loadData() {
  horses.value = await GetData("/horse/");
  ranks.value = await GetData("/other/ranks/");
  jumpHeights.value = await GetData("/other/jump-heights/");
}

await loadData();
</script>
