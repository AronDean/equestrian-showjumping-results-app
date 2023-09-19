<template>
  <div class="container">
    <header class="p-3 m-4 border-bottom">
      <h1 class="text-begin" v-if="!showModal">
        Ihre Turnierergebnisse
        <a
          v-if="!showModal"
          class="btn btn-outline-primary btn-outline float-end"
          @click="showModal = true"
          >Create</a
        >
      </h1>
      <h1 class="text-begin" v-if="showModal">
        Fügen Sie einen neuen Spielstand hinzu
        <a v-if="showModal" class="btn btn-outline-secondary float-end" @click="showModal = false"
          >Geh zurück</a
        >
      </h1>
    </header>

    <div v-if="showModal">
      <ScoreForm @score-created="onScoreCreated" />
    </div>

    <div v-if="!showModal">
      <div class="p-3 m-4">
        <div class="row">
          <div class="col-3">
            <label for="horse_id">Pferd</label>
            <select class="form-select mb-3" id="horse_id" v-model="horseId" required>
              <option v-for="horse in horses" :key="horse.id" :value="horse.id">
                {{ horse.name }}
              </option>
            </select>
          </div>
          <div class="col-3">
            <label for="jump_height_id">Hindernisshöhe</label>
            <select class="form-select mb-3" id="jump_height_id" v-model="jumpHeightId" required>
              <option v-for="jumpHeight in jumpHeights" :key="jumpHeight.id" :value="jumpHeight.id">
                {{ jumpHeight.title }}
              </option>
            </select>
          </div>
          <div class="col-3">
            <label for="rank_id">Rang</label>
            <select class="form-select mb-3" id="rank_id" v-model="rankId" required>
              <option v-for="rank in ranks" :key="rank.id" :value="rank.id">
                {{ rank.id }}
              </option>
            </select>
          </div>
          <div class="col-3 position-relative">
            <button
              @click="resetFilter"
              id="btn"
              class="btn btn-warning position-absolute top-50 start-50 translate-middle"
            >
              Reset
            </button>
          </div>
        </div>
        <div class="bg-light rounded-3">
          <div class="container-fluid py-2">
            <table class="table table-striped table-hover border align-middle">
              <thead class="table-primary">
                <tr>
                  <th>Pferd</th>
                  <th>Sprunghöhe</th>
                  <th>Rang</th>
                  <th>Punkte erhalten</th>
                  <th>Datum</th>
                </tr>
              </thead>
              <tbody>
                <tr v-for="score in scores" :key="score.id">
                  <td>{{ score.horse.name }}</td>
                  <td>{{ score.jump_height.title }}</td>
                  <td>{{ score.rank.id }}</td>
                  <td>{{ score.points_earned }}</td>
                  <td>{{ formattedCreatedAt(score.created_at) }}</td>
                </tr>

                <tr v-for="(obj, index) in dummy" :key="index">
                  <td>-</td>
                  <td>-</td>
                  <td>-</td>
                  <td>-</td>
                  <td>-</td>
                </tr>
              </tbody>
            </table>
          </div>
        </div>
        <nav class="p-2 float-end" aria-label="...">
          <ul class="pagination">
            <li class="page-item">
              <button class="page-link" @click="page = Math.max(page - 1, 1)" tabindex="-1">
                Vorherige
              </button>
            </li>
            <li
              class="page-item"
              v-for="(i, index) in total"
              v-bind:class="{ active: page === i }"
              :key="index"
            >
              <button class="page-link" @click="page = i">{{ i }}</button>
            </li>

            <li class="page-item">
              <button class="page-link" @click="page = Math.min(page + 1, total)">Nächste</button>
            </li>
          </ul>
        </nav>
      </div>
    </div>
  </div>
</template>
<style>
.container {
  margin-top: 50px;
}
</style>
<script setup lang="ts">
import api from "@/stores/api";
import { useAuthStore } from "@/stores/auth";
import { ref, watch } from "vue";

import ScoreForm from "@/views/ScoreForm.vue";
const showModal = ref(false);

const scores = ref([] as any);
const pageSize = ref(5);
const page = ref(1);
const total = ref(0);
const dummy = ref([] as any);
const authStore = useAuthStore();

const horseId = ref(null);
const rankId = ref(null);
const jumpHeightId = ref(null);

const horses = ref([] as any);
const ranks = ref([] as any);
const jumpHeights = ref([] as any);

function formattedCreatedAt(created_at: any): string {
  const options = {
    weekday: "short",
    year: "numeric",
    month: "short",
    day: "numeric"
  } as any;
  const date = new Date(created_at);
  return date.toLocaleString("de-DE", options).replace(".,", ",");
}

async function getScores() {
  console.log("fetching");
  let url = `/score/?page=${page.value}&pageSize=${pageSize.value}`;

  if (horseId.value) {
    url += `&horse=${horseId.value}`;
  }

  if (rankId.value) {
    url += `&rank=${rankId.value}`;
  }

  if (jumpHeightId.value) {
    url += `&jump=${jumpHeightId.value}`;
  }

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
      total.value = Math.ceil(response.data.count / pageSize.value);
      data = response.data.scores;
      dummy.value = Array(pageSize.value - data.length).fill({});
    })
    .catch((error) => {
      console.log(error);
    });
  return data;
}

async function onScoreCreated() {
  showModal.value = false;
  scores.value = await getScores();
}

scores.value = await getScores();

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

async function loadData() {
  horses.value = await GetData("/horse/");
  ranks.value = await GetData("/other/ranks/");
  jumpHeights.value = await GetData("/other/jump-heights/");

  horses.value.unshift({ id: null, name: "" });
  ranks.value.unshift({ id: null, name: "" });
  jumpHeights.value.unshift({ id: null, name: "" });
}

function resetFilter() {
  horseId.value = null;
  rankId.value = null;
  jumpHeightId.value = null;
}

await loadData();
watch([page, horseId, rankId, jumpHeightId], (newValues, oldValues) => {
  if (newValues.some((value, index) => value !== oldValues[index])) {
    getScores().then((data) => {
      scores.value = data;
    });
  }
});
</script>
