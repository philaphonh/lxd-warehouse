<template>
  <v-container>
    <h1 class="text-center">ຍິນດີຕ້ອນຮັບສູ່ສາງລ້ານຊ້າງດິຈິຕອລ</h1>
  </v-container>
</template>

<script lang="ts">
// @ is an alias to /src
// import HelloWorld from "@/components/HelloWorld.vue";
import { Vue, Component } from "vue-property-decorator";
import { namespace } from "vuex-class";

const App = namespace("App");

@Component({
  name: "HomePage",
  metaInfo() {
    return {
      title: "ໜ້າຫຼັກ"
    };
  }
})
export default class Home extends Vue {
  private pageTitle = "ໜ້າຫຼັກ";

  @App.Mutation("SET_APP_NAV_TITLE") APP_NAV_TITLE!: (title: string) => void;

  mounted() {
    this.APP_NAV_TITLE(this.pageTitle);
  }
}
</script>
