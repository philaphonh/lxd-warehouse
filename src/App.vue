<template>
  <v-app>
    <AppNavBar></AppNavBar>
    <AppNavDrawer></AppNavDrawer>
    <v-main>
      <router-view></router-view>
    </v-main>
  </v-app>
</template>

<script lang="ts">
import { Vue, Component } from "vue-property-decorator";
import { namespace } from "vuex-class";

const Auth = namespace("Auth");

import AppNavBar from "@/components/AppNavBar.vue";
import AppNavDrawer from "@/components/AppNavDrawer.vue";

@Component({
  name: "App",
  metaInfo() {
    return {
      titleTemplate: "%s | ສາງລ້ານຊ້າງດິຈິຕອລ"
    };
  },
  components: {
    AppNavBar,
    AppNavDrawer
  }
})
export default class App extends Vue {
  @Auth.Mutation("LOGIN_STATUS") isLoggedIn!: (status: boolean) => void;
  @Auth.Mutation("RETRIVE_AUTH_DATA") GET_AUTH!: () => void;

  created() {
    if (localStorage.getItem("token") == null) {
      this.isLoggedIn(false);
    } else {
      this.GET_AUTH();
      this.isLoggedIn(true);
    }
  }
}
</script>
