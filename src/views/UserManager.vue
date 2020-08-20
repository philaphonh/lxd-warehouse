<template>
  <v-container fluid>
    <v-data-table :headers="tableHeaders" :items="users" :search="searchText">
      <template v-slot:top>
        <v-toolbar flat>
          <v-toolbar-title>ຕາຕະລາງຂໍ້ມູນຜູ້ໃຊ້ງານລະບົບ</v-toolbar-title>
          <v-spacer></v-spacer>
          <v-text-field
            v-model="searchText"
            append-icon="mdi-magnify"
            label="ຄົ້ນຫາຜູ້ໃຊ້ງານລະບົບ"
            single-line
            hide-details
          ></v-text-field>
          <v-spacer></v-spacer>
          <v-btn color="primary" class="mr-2" dark to="/add-user"
            >ເພີ່ມຜູ້ໃຊ້ງານລະບົບ</v-btn
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
        <v-btn text :to="`/edit-user/${item.username}`">
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
import { User } from "@/models/User";

const App = namespace("App");

@Component({
  name: "UserManagerPage",
  metaInfo() {
    return {
      title: "ຈັດການຂໍ້ມູນຜູ້ໃຊ້ງານ"
    };
  }
})
export default class Import extends Vue {
  private pageTitle = "ຈັດການຜູ້ໃຊ້ງານ";

  private tableHeaders = [
    {
      text: "ລະຫັດ",
      value: "id"
    },
    {
      text: "ຊື່ຜູ້ໃຊ້",
      value: "username"
    },
    {
      text: "ຊື່ແທ້",
      value: "firstname"
    },
    {
      text: "ນາມສະກຸນ",
      value: "lastname"
    },
    {
      text: "ອີເມວ",
      value: "email"
    },
    {
      text: "ລະດັບຜູ້ໃຊ້",
      value: "role"
    },
    {
      text: "ຕົວເລືອກ",
      value: "actions",
      sortable: false
    }
  ];

  private users: Array<User> = [
    {
      id: 0,
      username: "N/A",
      firstname: "N/A",
      lastname: "N/A",
      email: "N/A",
      role: "N/A",
      imageUrl: ""
    }
  ];

  private searchText = "";
  private imageDialog = false;
  private imagePreview = "";

  @App.Mutation("SET_APP_NAV_TITLE") APP_NAV_TITLE!: (title: string) => void;

  getUsersData() {
    this.$http
      .get("/user/all")
      .then(res => {
        this.users = res.data;
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
    this.getUsersData();
  }
}
</script>
