<template>
  <v-container>
    <v-snackbar
      v-model="snackbar.show"
      top
      :color="snackbar.type"
      :timeout="snackbar.timeout"
      >{{ snackbar.text }}</v-snackbar
    >
    <v-toolbar class="mb-4">
      <v-toolbar-title>{{ pageTitle }}</v-toolbar-title>
    </v-toolbar>
    <v-card class="pa-4">
      <v-form>
        <v-text-field v-model="productName" label="ຊື່ສິນຄ້າ"></v-text-field>
        <v-row>
          <v-col>
            <v-select
              v-model="productBrandId"
              label="ຍີ່ຫໍ້ສິນຄ້າ"
              :items="productBrandList"
              item-value="id"
              item-text="name"
            ></v-select>
          </v-col>
          <v-col>
            <v-select
              v-model="productTypeId"
              label="ປະເພດສິນຄ້າ"
              :items="productTypeList"
              item-value="id"
              item-text="name"
            ></v-select>
          </v-col>
        </v-row>
        <v-textarea
          v-model="productDescription"
          label="ຄໍາອະທິບາຍສິນຄ້າ"
        ></v-textarea>
        <v-row>
          <v-col>
            <v-text-field
              v-model="productPrice"
              label="ລາຄາສິນຄ້າ"
              type="number"
            ></v-text-field>
          </v-col>
          <v-col>
            <v-file-input
              accept="image/*"
              label="ຮູບສິນຄ້າ"
              @change="onSelectImage"
            ></v-file-input>
          </v-col>
        </v-row>
        <v-btn color="primary" class="mr-2" dark @click="onSave">ບັນທຶກ</v-btn>
        <v-btn color="secondary" class="mr-2" @click="onCancel">ຍົກເລີກ</v-btn>
        <v-btn color="accent" dark type="reset">ລ້າງຂໍ້ມູນ</v-btn>
      </v-form>
    </v-card>
  </v-container>
</template>
<script lang="ts">
import { Vue, Component } from "vue-property-decorator";
import { namespace } from "vuex-class";
import router from "../router";

const App = namespace("App");

@Component({
  name: "ProductEditorPage",
  metaInfo() {
    return {
      title: router.currentRoute.params.id
        ? "ແກ້ໄຂຂໍ້ມູນສິນຄ້າ"
        : "ເພີ່ມຂໍ້ມູນສິນຄ້າ"
    };
  }
})
export default class ProductEditor extends Vue {
  private pageTitle = "ເພີ່ມ - ແກ້ໄຂສິນຄ້າ";

  private mode = "create";

  private productId: number | string | undefined = 0;
  private productName: string | undefined = "";
  private productBrandId: number | string | undefined = 0;
  private productTypeId: number | string | undefined = 0;
  private productDescription: string | undefined = "";
  private productPrice: number | string | undefined = 0;
  private productImage: File | undefined;

  private productBrandList = [
    {
      id: 0,
      name: "Nothing"
    }
  ];

  private productTypeList = [
    {
      id: 0,
      name: "Nothing"
    }
  ];

  private snackbar = {
    show: false,
    text: "",
    type: "",
    timeout: 3000
  };

  @App.Mutation("SET_APP_NAV_TITLE") APP_NAV_TITLE!: (title: string) => void;

  getProductData() {
    this.$http
      .get(`/product/${this.$router.currentRoute.params.id}`)
      .then(res => {
        if (res.data) {
          this.productId = res.data.id;
          this.productName = res.data.name;
          this.productBrandId = res.data.brand_id;
          this.productTypeId = res.data.type_id;
          this.productDescription = res.data.description;
          this.productPrice = res.data.price;
        } else {
          this.snackbar = {
            show: true,
            text: "ບໍ່ມີຂໍ້ມູນສິນຄ້າ!",
            type: "error",
            timeout: 2000
          };
          setTimeout(() => {
            this.$router.back();
          }, 2500);
        }
      });
  }

  getProductBrandData() {
    this.$http
      .get("/extras/product-brand/all")
      .then(res => {
        if (res.data.length > 0) {
          this.productBrandList = res.data;
        }
      })
      .catch(err => {
        console.log(err);
      });
  }

  getProductTypeData() {
    this.$http
      .get("/extras/product-type/all")
      .then(res => {
        if (res.data.length > 0) {
          this.productTypeList = res.data;
        }
      })
      .catch(err => {
        console.log(err);
      });
  }

  onSelectImage(file: File) {
    if (file) {
      this.productImage = file;
    }
  }

  onSave() {
    const form = new FormData();
    if (this.mode === "create") {
      form.append("name", this.productName as string);
      form.append("brandId", this.productBrandId as string);
      form.append("typeId", this.productTypeId as string);
      form.append("description", this.productDescription as string);
      form.append("price", this.productPrice as string);
      form.append("image", this.productImage as File);
      this.$http.post("/product/create", form).then(res => {
        if (res.data === "Successful") {
          this.snackbar = {
            show: true,
            text: "ເພີ່ມຂໍ້ມູນສໍາເລັດ",
            type: "success",
            timeout: 3000
          };
          setTimeout(() => {
            this.$router.back();
          }, 3500);
        } else {
          this.snackbar = {
            show: true,
            text: "ບໍ່ສໍາເລັດ!",
            type: "error",
            timeout: 3000
          };
        }
      });
    } else {
      form.append("id", this.productId as string);
      form.append("name", this.productName as string);
      form.append("brandId", this.productBrandId as string);
      form.append("typeId", this.productTypeId as string);
      form.append("description", this.productDescription as string);
      form.append("price", this.productPrice as string);
      form.append("image", this.productImage as File);
      this.$http.put(`/product/edit/${this.productId}`, form).then(res => {
        if (res.data === "Successful") {
          this.snackbar = {
            show: true,
            text: "ເພີ່ມຂໍ້ມູນສໍາເລັດ",
            type: "success",
            timeout: 3000
          };
          setTimeout(() => {
            this.$router.back();
          }, 3500);
        } else {
          this.snackbar = {
            show: true,
            text: "ບໍ່ສໍາເລັດ!",
            type: "error",
            timeout: 3000
          };
        }
      });
    }
  }

  onCancel() {
    this.$router.back();
  }

  mounted() {
    this.pageTitle = this.$router.currentRoute.params.id
      ? "ແກ້ໄຂຂໍ້ມູນສິນຄ້າ"
      : "ເພີ່ມຂໍ້ມູນສິນຄ້າ";
    this.APP_NAV_TITLE(this.pageTitle);
  }

  created() {
    if (this.$router.currentRoute.params.id) {
      this.mode = "edit";
      this.getProductData();
    } else {
      this.mode = "create";
    }
    this.getProductBrandData();
    this.getProductTypeData();
  }
}
</script>
