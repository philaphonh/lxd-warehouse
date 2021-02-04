<template>
  <v-container fluid>
    <v-snackbar
      v-model="snackbar.show"
      top
      :color="snackbar.type"
      :timeout="snackbar.timeout"
      >{{ snackbar.text }}</v-snackbar
    >
    <v-data-table :headers="tableHeaders" :items="exports">
      <template v-slot:top>
        <v-toolbar flat>
          <v-toolbar-title>ຕາຕະລາງການສົ່ງອອກສິນຄ້າ</v-toolbar-title>
          <v-spacer></v-spacer>
          <v-spacer></v-spacer>
          <v-btn color="primary" class="mr-2" dark to="/export-editor"
            >ນໍາອອກສິນຄ້າ</v-btn
          >
        </v-toolbar>
      </template>
      <template v-slot:item.exportStatus="props">
        {{
          props.item.exportStatus == "ORDERED"
            ? "ກໍາລັງດໍາເນີນການ"
            : "ສົ່ງອອກແລ້ວ"
        }}
      </template>
      <template v-slot:item.actions="props">
        <v-btn text @click.stop="getExportDetail(props.item.id)">
          <v-icon small color="primary" class="mr-2">mdi-view-list</v-icon
          >ເບິ່ງລາຍລະອຽດ
        </v-btn>
        <v-btn text @click="onConfirmExport(props.item)">
          <v-icon
            small
            color="primary"
            class="mr-2"
            :disabled="props.item.status == 'EXPORTED'"
            >mdi-pencil</v-icon
          >{{
            props.item.exportStatus == "ORDERED"
              ? "ຢືນຢັນການສົ່ງອອກ"
              : "ສົ່ງອອກແລ້ວ"
          }}
        </v-btn>
        <v-dialog v-model="detailDialog">
          <v-card>
            <v-card-title>ລາຍລະອຽດການນໍາອອກ</v-card-title>
            <v-list-item v-for="item in exportDetail" :key="item.id">
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

import { Export, ExportDetail } from "@/models/Export";

const App = namespace("App");

@Component({
  name: "ExportManagerPage",
  metaInfo() {
    return {
      title: "ຈັດການສົ່ງອອກສິນຄ້າ",
    };
  },
})
export default class ExportManager extends Vue {
  private pageTitle = "ຈັດການສົ່ງອອກສິນຄ້າ";

  private snackbar = {
    show: false,
    text: "",
    type: "",
    timeout: 3000,
  };

  private tableHeaders = [
    {
      text: "ລະຫັດ",
      value: "id",
    },
    {
      text: "ຕົວແທນຈໍາໜ່າຍ",
      value: "distributor",
    },
    {
      text: "ຜູ້ຮັບຜິດຊອບການນໍາອອກ",
      value: "user",
    },
    {
      text: "ເວລານໍາອອກ",
      value: "exportTime",
    },
    {
      text: "ສະຖານະ",
      value: "exportStatus",
    },
    {
      text: "ຕົວເລືອກ",
      value: "actions",
      sortable: false,
    },
  ];

  private exports: Array<Export> = [
    {
      id: 0,
      distributor: "Nothing",
      user: "Nothing",
      exportTime: Date(),
      exportStatus: "ORDERED",
    },
  ];

  private detailDialog = false;

  private exportDetail: Array<ExportDetail> = [];

  @App.Mutation("SET_APP_NAV_TITLE") APP_NAV_TITLE!: (title: string) => void;

  onConfirmExport(data: Export) {
    if (data.exportStatus !== "ORDERED") {
      return;
    }
    this.$http
      .put(`/product/export/confirm/${data.id}`)
      .then((res) => {
        if (res.status == 200) {
          this.snackbar = {
            show: true,
            text: "ຢືນຢັນການນໍາອອກແລ້ວ",
            type: "success",
            timeout: 3000,
          };
          this.getExportInfo();
        }
      })
      .catch((err) => {
        this.snackbar = {
          show: true,
          text: "ບໍ່ສໍາເລັດ!",
          type: "error",
          timeout: 3000,
        };
        console.log(err);
        return;
      });
  }

  getExportInfo() {
    this.$http
      .get("/product/export/all")
      .then((res) => {
        if (res.data.length > 0) {
          this.exports = res.data;
        }
      })
      .catch((err) => {
        console.log(err);
      });
  }

  getExportDetail(id: number) {
    this.detailDialog = true;
    this.$http
      .get("/product/export/detail/" + id)
      .then((res) => {
        if (res.data.length > 0) {
          this.exportDetail = res.data;
        }
      })
      .catch((err) => {
        console.log(err);
      });
  }

  created() {
    this.getExportInfo();
  }

  mounted() {
    this.APP_NAV_TITLE(this.pageTitle);
  }
}
</script>
