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
        <v-select
          v-model="supplierId"
          :items="suppliersList"
          item-text="name"
          item-value="id"
          label="ຜູ້ສະໜອງສິນຄ້າ"
        ></v-select>
        <v-data-table
          :headers="tableHeaders"
          :items="productsList"
          item-key="id"
          show-select
          v-model="selectedProducts"
        >
          <template v-slot:item.quantity="props">
            <v-text-field
              v-model="props.item.quantity"
              type="number"
            ></v-text-field>
          </template>
        </v-data-table>
        <v-btn color="primary" class="mr-2" @click="onImport">ດໍາເນີນການ</v-btn>
        <v-btn color="accent" @click="onCancel">ຍົກເລີກ</v-btn>
      </v-form>
    </v-card>
  </v-container>
</template>
<script lang="ts">
import { Vue, Component } from "vue-property-decorator";
import { namespace } from "vuex-class";
import Supplier from "@/models/Supplier";
import Product, { ImportedProduct } from "@/models/Product";

const App = namespace("App");

@Component({
  name: "ImportEditorPage",
  metaInfo() {
    return {
      title: "ນໍາເຂົ້າສິນຄ້າ"
    };
  }
})
export default class ImportEditor extends Vue {
  private pageTitle = "ນໍາເຂົ້າສິນຄ້າ";

  private snackbar = {
    show: false,
    text: "",
    type: "",
    timeout: 3000
  };

  private supplierId = 0;

  private selectedProducts: Array<ImportedProduct> = [];

  private suppliersList: Array<Supplier> = [
    {
      id: 0,
      name: "Nothing",
      address: "Nothing"
    }
  ];

  private productsList: Array<Product> = [
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
      text: "ຈໍານວນທີ່ຕ້ອງການນໍາເຂົ້າ",
      value: "quantity",
      sortable: false
    }
  ];

  @App.Mutation("SET_APP_NAV_TITLE") APP_NAV_TITLE!: (title: string) => void;

  onImport() {
    let error = false;
    if (!this.supplierId) {
    	error = true;
    	return;
    }
    for (let i = 0; i < this.selectedProducts.length; i++) {
      if (this.selectedProducts[i].quantity <= 0) {
        error = true;
        break;
      }
    }
    if (error) {
      return;
    }
    let id = "";
    const form = new FormData();
    form.append("supplier", this.supplierId.toString());
    form.append("user", localStorage.getItem("id") as string);
    this.$http
      .post("/product/import/order", form)
      .then(res => {
        if (res.status == 201) {
          id = res.data;
          this.selectedProducts.forEach(item => {
            this.$http
              .post(`/product/import/order/detail/${id}`, item, {
                headers: { "Content-Type": "application/json" }
              })
              .then(res => {
                if (res.status == 201) {
                  this.snackbar = {
                    show: true,
                    text: "ເພີ່ມຂໍ້ມູນສໍາເລັດ",
                    type: "success",
                    timeout: 3000
                  };
                  setTimeout(() => {
                    this.$router.back();
                  }, 3500);
                }
              })
              .catch(err => {
                this.snackbar = {
                  show: true,
                  text: "ບໍ່ສໍາເລັດ!",
                  type: "error",
                  timeout: 3000
                };
                console.log(err);
                return;
              });
          });
        } else {
          return;
        }
      })
      .catch(err => {
        this.snackbar = {
          show: true,
          text: "ບໍ່ສໍາເລັດ!",
          type: "error",
          timeout: 3000
        };
        console.log(err);
        return;
      });
  }

  onCancel() {
    this.$router.back();
  }

  getSupplierData() {
    this.$http
      .get("/extras/supplier/all")
      .then(res => {
        if (res.data.length > 0) {
          this.suppliersList = res.data;
        }
      })
      .catch(err => {
        this.snackbar = {
          show: true,
          text: "ດຶງຂໍ້ມູນບໍ່ສໍາເລັດ!",
          type: "error",
          timeout: 3000
        };
        console.log(err);
      });
  }

  getProductsData() {
    this.$http
      .get("/product/all")
      .then(res => {
        this.productsList = res.data;
      })
      .catch(err => {
        this.snackbar = {
          show: true,
          text: "ດຶງຂໍ້ມູນບໍ່ສໍາເລັດ!",
          type: "error",
          timeout: 3000
        };
        console.log(err);
      });
  }

  created() {
    this.getSupplierData();
    this.getProductsData();
  }

  mounted() {
    this.APP_NAV_TITLE(this.pageTitle);
  }
}
</script>
