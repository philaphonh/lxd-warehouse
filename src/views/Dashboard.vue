<template>
  <v-container fluid>
    <h1>ກະດານວຽກ</h1>
    <v-row>
      <v-col>
        <v-card color="primary" dark>
          <v-card-title>ຈໍານວນສິນຄ້າທັງໝົດ</v-card-title>
          <v-card-text>
            <h3 class="headline white--text">6969</h3>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col>
        <v-card color="primary" dark>
          <v-card-title>ສິນຄ້າທີ່ມີຫຼາຍທີ່ສຸດ</v-card-title>
          <v-card-text>
            <h3 class="headline white--text">MSI GS66 Stealth (96)</h3>
          </v-card-text>
        </v-card>
      </v-col>
      <v-col>
        <v-card color="primary" dark>
          <v-card-title>ສິນຄ້າທີ່ມີໜ້ອຍທີ່ສຸດ</v-card-title>
          <v-card-text>
            <h3 class="headline white--text">MSI GE66 Raider (69)</h3>
          </v-card-text>
        </v-card>
      </v-col>
    </v-row>
  </v-container>
</template>
<script lang="ts">
import { Vue, Component } from "vue-property-decorator";
import { namespace } from "vuex-class";

const App = namespace("App");

@Component({
  name: "DashboardPage",
  metaInfo() {
    return {
      title: "ກະດານວຽກ"
    };
  }
})
export default class Dashboard extends Vue {
  private pageTitle = "ກະດານວຽກ";

  @App.Mutation("SET_APP_NAV_TITLE") APP_NAV_TITLE!: (title: string) => void;

  mounted() {
    this.APP_NAV_TITLE(this.pageTitle);
  }
}
</script>
