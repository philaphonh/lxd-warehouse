<template>
  <v-container>
    <h1>ກ່ຽວກັບລະບົບ</h1>
    <p>
      ເວັບນີ້ແມ່ນບົດໂຄງການໃນວິຊາ ວິສະວະກໍາຊອບແວ ເຊິ່ງພັດທະນາໂດຍ ພິລະພົນ
      ອິນທະວົງສາ ໂດຍນໍາໃຊ້ພາສາ TypeScript ແລະ Go ເປັນເຄື່ອງມື.
    </p>
  </v-container>
</template>
<script lang="ts">
import { Vue, Component } from "vue-property-decorator";
import { namespace } from "vuex-class";

const App = namespace("App");

@Component({
  name: "AboutPage",
  metaInfo() {
    return {
      title: "ກ່ຽວກັບ"
    };
  }
})
export default class About extends Vue {
  private pageTitle = "ກ່ຽວກັບ";

  @App.Mutation("SET_APP_NAV_TITLE") APP_NAV_TITLE!: (title: string) => void;

  mounted() {
    this.APP_NAV_TITLE(this.pageTitle);
  }
}

// import Vue from "vue";
// export default Vue.extend({
//   name: "AboutPage",
//   data: () => ({
//     pageTitle: "About This Project"
//   }),
//   mounted() {
//     this.$store.commit("SET_APP_NAV_TITLE", this.pageTitle);
//   }
// });
//
</script>
