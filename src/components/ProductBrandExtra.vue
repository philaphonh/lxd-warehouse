<template>
  <v-container>
    <v-snackbar
      v-model="snackbar.show"
      top
      :color="snackbar.type"
      :timeout="snackbar.timeout"
    >
      {{ snackbar.text }}
    </v-snackbar>
    <v-data-table :headers="tableHeaders" :items="productBrands">
      <template v-slot:top>
        <v-toolbar flat>
          <v-toolbar-title>ຕາຕະລາງຂໍ້ມູນຍີ່ຫໍ້ສິນຄ້າ</v-toolbar-title>
          <v-spacer></v-spacer>
          <v-dialog v-model="dialog">
            <template v-slot:activator="{ on, attrs }">
              <v-btn color="primary" v-bind="attrs" v-on="on"
                >ເພີ່ມຍີ່ຫໍ້</v-btn
              >
            </template>
            <v-card>
              <v-card-title>
                <span class="headline">{{ dialogTitle() }}</span>
              </v-card-title>
              <v-card-text>
                <v-form>
                  <v-text-field
                    label="ລະຫັດ"
                    v-model="editItem.id"
                    disabled
                  ></v-text-field>
                  <v-text-field
                    label="ຊື່ຍີ່ຫໍ້"
                    v-model="editItem.name"
                  ></v-text-field>
                </v-form>
              </v-card-text>
              <v-card-actions>
                <v-btn color="primary" @click="onSave()" class="mr-2"
                  >ບັນທຶກ</v-btn
                >
                <v-btn color="accent" @click="closeDialog()">ຍົກເລີກ</v-btn>
              </v-card-actions>
            </v-card>
          </v-dialog>
        </v-toolbar>
      </template>
      <template v-slot:item.actions="{ item }">
        <v-btn text @click="openDialog(item)">
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
import { Vue, Component, Watch } from "vue-property-decorator";

import ProductBrand from "@/models/ProductBrand";

@Component({
  name: "ProductBrandExtraTab"
})
export default class ProductBrandExtra extends Vue {
  private tableHeaders = [
    {
      text: "ລະຫັດ",
      value: "id"
    },
    {
      text: "ຊື່ຍີ່ຫໍ້ສິນຄ້າ",
      value: "name"
    },
    {
      text: "ຕົວເລືອກ",
      value: "actions",
      sortable: false
    }
  ];

  private editIndex = -1;
  private editItem = {
    id: 0,
    name: ""
  };

  private dialog = false;

  private mode = "create";

  private dialogTitle() {
    return this.mode === "create" ? "ເພີ່ມຂໍ້ມູນ" : "ແກ້ໄຂຂໍ້ມູນ";
  }
  @Watch("dialog")
  onChanged(val: boolean) {
    val || this.closeDialog();
  }

  private snackbar = {
    show: false,
    text: "",
    type: "",
    timeout: 3000
  };

  private productBrands: Array<ProductBrand> = [
    {
      id: 0,
      name: "Nothing"
    }
  ];

  getProductBrandData() {
    this.$http
      .get("/extras/product-brand/all")
      .then(res => {
        if (res.data.length > 0) {
          this.productBrands = res.data;
        }
      })
      .catch(err => {
        console.log(err);
      });
  }

  openDialog(item: ProductBrand) {
    this.editIndex = this.productBrands.indexOf(item);
    this.editItem = Object.assign({}, item);
    this.mode = "edit";
    this.dialog = true;
  }

  onSave() {
    const form = new FormData();
    if (this.mode === "edit" && this.editItem.id !== 0) {
      form.append("name", this.editItem.name);
      this.$http
        .put(`/extras/product-brand/edit/${this.editItem.id}`, form, {
          headers: { "Content-Type": "multipart/form" }
        })
        .then(res => {
          if (res.data === "Successful") {
            this.snackbar = {
              show: true,
              text: "ແກ້ໄຂຂໍ້ມູນສໍາເລັດ!",
              type: "success",
              timeout: 3000
            };
            this.getProductBrandData();
            this.closeDialog();
          }
        })
        .catch(err => {
          console.log(err);
          this.snackbar = {
            show: true,
            text: "ບໍ່ສໍາເລັດ!",
            type: "error",
            timeout: 3000
          };
        });
    } else {
      form.append("name", this.editItem.name);
      // form.append("csrf", this.csrfToken);
      // form.append("csrf", )
      this.$http
        .post("/extras/product-brand/create", form, {
          headers: {
            "Content-Type": "multipart/form"
            // "X-CSRF-TOKEN": this.csrfToken
          }
        })
        .then(res => {
          if (res.data === "Successful") {
            this.snackbar = {
              show: true,
              text: "ເພີ່ມຂໍ້ມູນສໍາເລັດ!",
              type: "success",
              timeout: 3000
            };
            this.getProductBrandData();
            this.closeDialog();
          } else {
            this.snackbar = {
              show: true,
              text: "ບໍ່ສໍາເລັດ!",
              type: "error",
              timeout: 3000
            };
          }
        })
        .catch(err => {
          console.log(err);
          this.snackbar = {
            show: true,
            text: "ບໍ່ສໍາເລັດ!",
            type: "error",
            timeout: 3000
          };
        });
    }
  }

  closeDialog() {
    this.dialog = false;
    this.mode = "create";
    this.editItem = {
      id: 0,
      name: ""
    };
  }

  // private csrfToken = "";

  // async getCSRFToken() {
  //   const res = await this.$http.get("/util/csrftoken");
  //   this.csrfToken = res.data;
  // }

  created() {
    this.getProductBrandData();
  }
}
</script>
