<template>
  <div class="container">
    <header class="p-3 m-4 border-bottom">
      <h1 v-if="!showModal" class="text-begin">
        Pferdeinventar
        <a v-if="!showModal" class="btn btn-outline-primary float-end" @click="showModal = true"
          >Create</a
        >
      </h1>
      <h1 v-if="showModal && !isUpdate" class="text-begin">
        Fügen Sie ein Pferd hinzu
        <a v-if="showModal" class="btn btn-outline-secondary float-end" @click="hideModal"
          >Geh zurück</a
        >
      </h1>
      <h1 v-if="showModal && isUpdate" class="text-begin">
        Pferd aktualisieren
        <a v-if="showModal" class="btn btn-outline-secondary float-end" @click="hideModal"
          >Go Back</a
        >
      </h1>
    </header>

    <div v-if="showModal">
      <HorseFrom :horse="horse" :is-update="isUpdate" @horse-created="onHorseCreated"></HorseFrom>
    </div>

    <div v-if="!showModal">
      <div class="bg-light rounded-3">
        <div class="container-fluid py-2">
          <table class="table table-striped table-hover border">
            <thead class="table-primary">
              <tr>
                <th>Name</th>
                <th>Farbe</th>
                <th>Alter</th>
                <th>Aktionen</th>
              </tr>
            </thead>
            <tbody>
              <tr v-for="horse in horses" :key="horse.id">
                <td>{{ horse.name }}</td>
                <td>{{ horse.color }}</td>
                <td>{{ horse.age }}</td>
                <td>
                  <button class="btn btn-secondary pr-2" @click="setModal(horse)">Edit</button>
                </td>
              </tr>
            </tbody>
          </table>
        </div>
      </div>
    </div>
  </div>
</template>

<style>
.container {
  margin-top: 50px;
}
</style>

<script lang="ts">
import api from "@/stores/api";
import { defineComponent, ref } from "vue";
import { useAuthStore } from "@/stores/auth";
import type { Horse } from "@/types/interface";
import HorseFrom from "./HorseFrom.vue";

const authStore = useAuthStore();

async function getHorses(): Promise<Horse[]> {
  const config = {
    method: "get",
    maxBodyLength: Infinity,
    url: "/horse/",
    headers: {
      Authorization: authStore.accessToken
    }
  };

  let retVal: Horse[] = [];

  try {
    const response = await api.request(config);
    if (response.status !== 200) {
      throw new Error("Error: " + response.status);
    }
    retVal = response.data;
  } catch (error) {
    console.log(error);
  }

  return retVal;
}

async function deleteHorse(id: number) {
  let config = {
    method: "delete",
    url: `/horse/${id}`,
    headers: {
      Authorization: authStore.accessToken
    }
  };
  let retVal = 0;
  await api
    .request(config)
    .then((response) => {
      retVal = response.data.id;
    })
    .catch((error) => {
      console.log(error);
    });

  return retVal;
}

export default defineComponent({
  async setup() {
    const showModal = ref(false);
    const isUpdate = ref(false);
    const horse = ref<Horse>({
      id: 0,
      age: 0,
      name: "",
      weight: 0,
      color: ""
    } as Horse);
    const horses = ref<Horse[]>([]);
    horses.value = await getHorses();

    function hideModal() {
      showModal.value = false;
      isUpdate.value = false;
    }

    function setModal(_horse: Horse) {
      showModal.value = true;
      isUpdate.value = true;
      horse.value = _horse;
    }

    async function onHorseCreated(id: number) {
      console.log("horse created: " + id);
      showModal.value = false;
      isUpdate.value = false;
      horses.value = await getHorses();
    }

    return {
      horse,
      horses,
      getHorses,
      showModal,
      setModal,
      isUpdate,
      deleteHorse,
      hideModal,
      onHorseCreated
    };
  },
  components: { HorseFrom }
});
</script>
