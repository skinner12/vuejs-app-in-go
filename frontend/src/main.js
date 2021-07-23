import Vue from "vue";
import App from "./App.vue";
import axios from "axios";
import VueAxios from "vue-axios";

Vue.config.productionTip = false;

const client = axios.create({
  // Opzione 1: Usare .env files per distinguere prod e dev server
  // endpoints.
  // baseURL: process.env.FRONTEND_API_BASE_URL,
  //
  // Opzione 2: Usare lo stesso endpoint, ma istruire Vue di effettuare
  // il proxy al vero API URL
  baseURL: "/api/v1"
  // Test CORS:
  //baseURL: "http://localhost:8081/api/v1",
});
Vue.use(VueAxios, client);

new Vue({
  render: h => h(App)
}).$mount("#app");
