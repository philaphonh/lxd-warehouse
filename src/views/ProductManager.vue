<template>
  <v-container fluid>
    <v-data-table
      :headers="tableHeaders"
      :items="products"
      :search="searchText"
    >
      <template v-slot:top>
        <v-toolbar flat>
          <v-toolbar-title>ຕາຕະລາງສິນຄ້າ</v-toolbar-title>
          <v-spacer></v-spacer>
          <v-text-field
            v-model="searchText"
            append-icon="mdi-magnify"
            label="ຄົ້ນຫາສິນຄ້າ"
            single-line
            hide-details
          ></v-text-field>
          <v-spacer></v-spacer>
          <v-btn color="primary" class="mr-2" dark to="/add-product"
            >ເພີ່ມຂໍ້ມູນສິນຄ້າ</v-btn
          >
        </v-toolbar>
      </template>
      <template v-slot:item.actions="{ item }">
        <v-btn text @click.stop="getImage(item.imageUrl)">
          <v-icon small color="primary" class="mr-2">mdi-image</v-icon>ເບິ່ງຮູບ
        </v-btn>
        <v-dialog v-model="imageDialog" max-width="600">
          <v-card>
            <v-card-actions>
              <v-btn @click.stop="imageDialog = false">X</v-btn>
            </v-card-actions>
            <v-img :src="imagePreview"></v-img>
          </v-card>
        </v-dialog>
        <v-btn text :to="`/edit-product/${item.id}`">
          <v-icon small color="primary" class="mr-2">mdi-pencil</v-icon>ແກ້ໄຂ
        </v-btn>
        <v-btn text color="accent">
          <v-icon small>mdi-delete</v-icon>ລົບ
        </v-btn>
      </template>
    </v-data-table>
  </v-container>
</template>
<script lang="ts">
import { Vue, Component } from "vue-property-decorator";
import { namespace } from "vuex-class";

import Product from "@/models/Product";

const App = namespace("App");

@Component({
  name: "ProductManagerPage",
  metaInfo() {
    return {
      title: "ຈັດການສິນຄ້າ"
    };
  }
})
export default class ProductManager extends Vue {
  private pageTitle = "ຈັດການສິນຄ້າ";
  private tableHeaders = [
    {
      text: "ລະຫັດ",
      value: "id"
    },
    {
      text: "ຊື່ສິນຄ້າ",
      value: "name"
    },
    {
      text: "ລາຍລະອຽດ",
      value: "description"
    },
    {
      text: "ລາຄາ",
      value: "price"
    },
    {
      text: "ຈໍານວນໃນສາງ",
      value: "qty"
    },
    {
      text: "ຕົວເລືອກ",
      value: "actions",
      sortable: false
    }
  ];
  private products: Array<Product> = [
    {
      id: 0,
      name: "Nothing",
      brand: "Nothing",
      type: "Nothing",
      description: "Nothing",
      price: 0,
      qty: 0,
      imageUrl: ""
    }
  ];

  private searchText = "";
  private imageDialog = false;
  private imagePreview = "";

  @App.Mutation("SET_APP_NAV_TITLE") APP_NAV_TITLE!: (title: string) => void;

  getProductsData() {
    this.$http
      .get("/product/all")
      .then(res => {
        this.products = res.data;
      })
      .catch(err => {
        console.log(err);
      });
  }

  getImage(url: string) {
    this.imageDialog = true;
    this.imagePreview = url;
  }

  mounted() {
    this.APP_NAV_TITLE(this.pageTitle);
    this.getProductsData();
  }
}
</script>
