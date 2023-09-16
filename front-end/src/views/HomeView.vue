<template>
  <div class="container">
    <header class="m-2 border-bottom">
      <h1 class="text-center">Bestenliste für die letzte Woche</h1>
    </header>
    <div class="bg-light rounded-3">
      <div class="container-fluid">
        <table class="table table-striped table-hover border">
          <thead class="table-primary">
            <tr>
              <th>Name</th>
              <th>Email</th>
              <th>Anzahl Umgänge</th>
              <th>Punkte erhalten</th>
            </tr>
          </thead>
          <tbody>
            <tr v-for="item in leaderBoard" :key="item.email">
              <td>{{ item.name }}</td>
              <td>{{ item.email }}</td>
              <td>{{ item.game_count }}</td>
              <td>{{ item.points_earned }}</td>
            </tr>
          </tbody>
        </table>
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

export default defineComponent({
  async setup() {
    const dataFetched = ref(false);
    const leaderBoard = ref([] as any[]);
    const error = ref("");

    let config = {
      method: "get",
      url: "/score/leader-board/",
      headers: {}
    };

    await api
      .request(config)
      .then((response: any) => {
        if (response.status != 200) {
          error.value = response?.data?.error;
          return;
        }

        leaderBoard.value = response.data;
        dataFetched.value = true;
      })
      .catch((error: any) => {
        console.log(error);
      });

    return {
      dataFetched,
      leaderBoard,
      error
    };
  }
});
</script>
