import Vue from "vue";
import Vuex from "vuex";

import App from "@/store/modules/app";
import Auth from "@/store/modules/auth";

Vue.use(Vuex);

export default new Vuex.Store({
  modules: {
    App,
    Auth
  }
});

// export default new Vuex.Store({
//   state: {
//     appNavTitle: "Warehouse",
//     appNavDrawer: false,
//   },
//   mutations: {
//     SET_APP_NAV_TITLE(state, payload: string) {
//       state.appNavTitle = payload;
//     },
//     SET_APP_NAV_DRAWER(state, payload: boolean) {
//       state.appNavDrawer = payload;
//     },
//   },
//   actions: {},
//   modules: {},
// });
