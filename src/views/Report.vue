<template>
  <v-container></v-container>
</template>
<script lang="ts">
import { Vue, Component } from "vue-property-decorator";
import { namespace } from "vuex-class";

const App = namespace("App");

@Component({
  name: "ReportPage",
  metaInfo() {
    return {
      title: "ລາຍງານ"
    };
  }
})
export default class Report extends Vue {
  private pageTitle = "ລາຍງານ";

  @App.Mutation("SET_APP_NAV_TITLE") APP_NAV_TITLE!: (title: string) => void;

  mounted() {
    this.APP_NAV_TITLE(this.pageTitle);
  }
}
</script>
