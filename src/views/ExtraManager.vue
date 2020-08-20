<template>
  <v-container>
    <v-tabs v-model="tab" grow background-color="primary" dark>
      <v-tab v-for="tab of tabs" :key="tab.key" :to="`#${tab.key}`" exact>{{
        tab.title
      }}</v-tab>
    </v-tabs>
    <v-tabs-items v-model="tab">
      <v-tab-item v-for="tab of tabs" :key="tab.key" :value="tab.key">
        <component :is="tab.component"></component>
      </v-tab-item>
    </v-tabs-items>
  </v-container>
</template>
<script lang="ts">
import { Vue, Component } from "vue-property-decorator";
import { namespace } from "vuex-class";

const ProductBrandTab = () => import("@/components/ProductBrandExtra.vue");
const ProductTypeTab = () => import("@/components/ProductTypeExtra.vue");
const SupplierTab = () => import("@/components/SupplierExtra.vue");
const DistributorTab = () => import("@/components/DistributorExtra.vue");

const App = namespace("App");

@Component({
  name: "ExtraManagerPage",
  metaInfo() {
    return {
      title: "ຈັດການຂໍ້ມູນເສີມ"
    };
  },
  components: {
    ProductBrandTab,
    ProductTypeTab,
    SupplierTab,
    DistributorTab
  }
})
export default class Import extends Vue {
  private pageTitle = "ຈັດການຂໍ້ມູນເສີມ";

  private tab = null;

  private tabs = [
    {
      title: "ຍີ່ຫໍ້ສິນຄ້າ",
      key: "product-brand",
      component: ProductBrandTab
    },
    {
      title: "ປະເພດສິນຄ້າ",
      key: "product-type",
      component: ProductTypeTab
    },
    {
      title: "ຜູ້ສະໜອງສິນຄ້າ",
      key: "supplier",
      component: SupplierTab
    },
    {
      title: "ຜູ້ຈໍາໜ່າຍສິນຄ້າ",
      key: "distributor",
      component: DistributorTab
    }
  ];

  @App.Mutation("SET_APP_NAV_TITLE") APP_NAV_TITLE!: (title: string) => void;

  mounted() {
    this.APP_NAV_TITLE(this.pageTitle);
  }
}
</script>
