<template>
  <div id="app">
    <button class="example" type="button" @click="getFrase()">
      Voglio una nuova frase
    </button>
    <div v-if="frase != null">
      <h1>{{ frase.titolo }}</h1>
      <p>{{ frase.enunciato }}</p>
    </div>
  </div>
</template>

<script>
import Vue from "vue";

export default {
  name: "App",
  components: {},
  data() {
    return {
      frase: null,
    };
  },
  methods: {
    getFrase() {
      console.log("Clicked, adesso ottengo una nuova frase");
      Vue.axios.get("/embed").then((response) => (this.frase = response.data));
    },
  },
};
</script>
<style scoped>
.container {
  display: flex;
  flex-direction: column;
  flex-grow: 1;
  margin: 0 auto;
}

.row {
  width: 100%;
  display: flex;
  flex-direction: row;
  flex-grow: 1;
  justify-content: space-around;
}
</style>
