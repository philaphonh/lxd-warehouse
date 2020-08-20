<template>
  <v-container fluid>
    <v-snackbar
      v-model="snackbar.show"
      top
      :color="snackbar.type"
      :timeout="snackbar.timeout"
      >{{ snackbar.text }}</v-snackbar
    >
    <v-data-table :headers="tableHeaders" :items="imports">
      <template v-slot:top>
        <v-toolbar flat>
          <v-toolbar-title>ຕາຕະລາງການນໍາເຂົ້າສິນຄ້າ</v-toolbar-title>
          <v-spacer></v-spacer>
          <v-spacer></v-spacer>
          <v-btn color="primary" class="mr-2" dark to="/import-editor"
            >ສັ່ງນໍາເຂົ້າສິນຄ້າ</v-btn
          >
        </v-toolbar>
      </template>
      <template v-slot:item.status="props">
        {{ props.item.status == "ORDERED" ? "ສັ່ງຈອງແລ້ວ" : "ນໍາເຂົ້າແລ້ວ" }}
      </template>
      <template v-slot:item.actions="props">
        <v-btn text @click.stop="getImportDetail(props.item.id)">
          <v-icon small color="primary" class="mr-2">mdi-view-list</v-icon
          >ເບິ່ງລາຍລະອຽດ
        </v-btn>
        <v-btn
          text
          @click="onConfirmImport(props.item)"
          v-if="props.item.status !== 'IMPORTED'"
        >
          <v-icon
            small
            color="primary"
            class="mr-2"
            :disabled="props.item.status == 'IMPORTED'"
            >mdi-pencil</v-icon
          >{{
            props.item.status == "ORDERED"
              ? "ຢືນຢັນການນໍາເຂົ້າ"
              : "ນໍາເຂົ້າແລ້ວ"
          }}
        </v-btn>
        <v-dialog v-model="detailDialog">
          <v-card>
            <v-card-title>ລາຍລະອຽດການນໍາເຂົ້າ</v-card-title>
            <v-list-item v-for="item in importDetail" :key="item.id">
              <v-list-item-title>{{ item.product }}</v-list-item-title>
              <v-list-item-action>{{ item.qty }} ອັນ</v-list-item-action>
            </v-list-item>
          </v-card>
        </v-dialog>
      </template>
    </v-data-table>
  </v-container>
</template>
<script lang="ts">
import { Vue, Component } from "vue-property-decorator";
import { namespace } from "vuex-class";

import { Import, ImportDetail } from "@/models/Import";

const App = namespace("App");

@Component({
  name: "ImportManagerPage",
  metaInfo() {
    return {
      title: "ຈັດການນໍາເຂົ້າສິນຄ້າ"
    };
  }
})
export default class ImportManager extends Vue {
  private pageTitle = "ຈັດການນໍາເຂົ້າສິນຄ້າ";

  private snackbar = {
    show: false,
    text: "",
    type: "",
    timeout: 3000
  };

  private tableHeaders = [
    {
      text: "ລະຫັດ",
      value: "id"
    },
    {
      text: "ຜູ້ສະໜອງສິນຄ້າ",
      value: "supplier"
    },
    {
      text: "ຜູ້ຮັບຜິດຊອບການນໍາເຂົ້າ",
      value: "user"
    },
    {
      text: "ເວລານໍາເຂົ້າ",
      value: "importTime"
    },
    {
      text: "ສະຖານະການນໍາເຂົ້າ",
      value: "status"
    },
    {
      text: "ຕົວເລືອກ",
      value: "actions",
      sortable: false
    }
  ];

  private imports: Array<Import> = [
    {
      id: 0,
      supplier: "Nothing",
      user: "Nothing",
      importTime: Date(),
      status: "Ordered"
    }
  ];

  private detailDialog = false;

  private importDetail: Array<ImportDetail> = [];

  @App.Mutation("SET_APP_NAV_TITLE") APP_NAV_TITLE!: (title: string) => void;

  getImportInfo() {
    this.$http
      .get("/product/import/all")
      .then(res => {
        if (res.data.length > 0) {
          this.imports = res.data;
        }
      })
      .catch(err => {
        console.log(err);
      });
  }

  getImportDetail(id: number) {
    this.detailDialog = true;
    this.$http
      .get("/product/import/detail/" + id)
      .then(res => {
        if (res.data.length > 0) {
          this.importDetail = res.data;
        }
      })
      .catch(err => {
        console.log(err);
      });
  }

  onConfirmImport(data: Import) {
    if (data.status !== "ORDERED") {
      return;
    }
    this.$http
      .put(`/product/import/confirm/${data.id}`)
      .then(res => {
        if (res.status == 200) {
          this.snackbar = {
            show: true,
            text: "ຢືນຢັນການນໍາເຂົ້າແລ້ວ",
            type: "success",
            timeout: 3000
          };
          this.getImportInfo();
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
      });
  }

  created() {
    this.getImportInfo();
  }

  mounted() {
    this.APP_NAV_TITLE(this.pageTitle);
  }
}
</script>
