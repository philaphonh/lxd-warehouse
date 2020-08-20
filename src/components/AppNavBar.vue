<template>
  <v-app-bar app color="primary" dark>
    <div class="d-flex align-center">
      <v-app-bar-nav-icon
        v-if="isLoggedIn"
        @click="APP_NAV_DRAWER(!AppNavDrawer)"
      ></v-app-bar-nav-icon>

      <v-toolbar-title>{{ AppNavTitle }}</v-toolbar-title>
    </div>

    <v-spacer></v-spacer>

    <v-btn icon v-if="isLoggedIn">
      <v-avatar color="secondary">
        <img :src="authData.imageUrl" alt="user" />
      </v-avatar>
    </v-btn>
  </v-app-bar>
</template>
<script lang="ts">
import { Vue, Component } from "vue-property-decorator";
import { namespace } from "vuex-class";
import { AuthData } from "@/models/User";

const App = namespace("App");
const Auth = namespace("Auth");

@Component({
  name: "AppNavBar"
})
export default class AppNavBar extends Vue {
  @App.State("appNavDrawer") AppNavDrawer!: boolean;
  @App.State("appNavTitle") AppNavTitle!: string;
  @Auth.State("authData") authData!: AuthData;
  @App.Mutation("SET_APP_NAV_DRAWER") APP_NAV_DRAWER!: (
    status: boolean
  ) => void;
  @Auth.State("isLoggedIn") isLoggedIn!: boolean;
}

// export default Vue.extend({
//   name: "AppNavBar",
//   data: () => ({}),
//   computed: {
//     ...mapState(["appNavTitle", "appNavDrawer"])
//   },
//   methods: {
//     ...mapMutations({
//       setAppNavDrawer: "SET_APP_NAV_DRAWER"
//     })
//   }
// });
</script>
