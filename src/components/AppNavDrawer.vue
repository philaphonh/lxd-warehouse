<template>
  <v-navigation-drawer app temporary v-model="drawer">
    <v-divider></v-divider>
    <v-list nav>
      <v-list-item
        v-for="drawer of drawerItems"
        :key="drawer.name"
        :to="drawer.path"
      >
        <v-list-item-title>{{ drawer.name }}</v-list-item-title>
      </v-list-item>
      <v-list-item>
        <v-btn block color="accent" dark @click="logout">ອອກຈາກລະບົບ</v-btn>
      </v-list-item>
    </v-list>
  </v-navigation-drawer>
</template>
<script lang="ts">
import { Vue, Component } from "vue-property-decorator";
import { namespace } from "vuex-class";
import { AuthData } from "@/models/User";

const App = namespace("App");
const Auth = namespace("Auth");

@Component({
  name: "AppNavDrawer"
})
export default class AppNavDrawer extends Vue {
  @App.State("appNavDrawer") AppNavDrawer!: boolean;
  @Auth.State("authData") authData!: AuthData;
  @App.Mutation("SET_APP_NAV_DRAWER") APP_NAV_DRAWER!: (
    status: boolean
  ) => void;
  @Auth.Mutation("LOGIN_STATUS") SET_LOGIN_STATUS!: (status: boolean) => void;
  @Auth.Mutation("CLEAR_AUTH_DATA") LOGOUT!: () => void;

  drawerItems = [
    {
      name: "ໜ້າຫຼັກ",
      path: "/"
    },
    {
      name: "ກະດານວຽກ",
      path: "/dashboard"
    },
    {
      name: "ຈັດການສິນຄ້າ",
      path: "/products"
    },
    {
      name: "ຈັດການການນໍາເຂົ້າສິນຄ້າ",
      path: "/import"
    },
    {
      name: "ຈັດການການສົ່ງອອກສິນຄ້າ",
      path: "/export"
    },
    {
      name: "ຈັດການຂໍ້ມູນເສີມ",
      path: "/extras"
    },
    {
      name: "ຈັດການຜູ້ໃຊ້ງານ",
      path: "/users"
    },
    {
      name: "ລາຍງານ",
      path: "/report"
    },
    {
      name: "ກ່ຽວກັບ",
      path: "/about"
    }
  ];

  get drawer() {
    return this.AppNavDrawer;
  }

  set drawer(value: boolean) {
    this.APP_NAV_DRAWER(value);
  }

  logout() {
    localStorage.clear();
    this.LOGOUT();
    this.SET_LOGIN_STATUS(false);
    console.log(this.authData);
    this.$router.push("/login");
  }
}
</script>
