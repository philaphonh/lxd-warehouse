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
        <v-row>
          <v-col>
            <v-text-field v-model="userFirstName" label="ຊື່ແທ້"></v-text-field>
          </v-col>
          <v-col>
            <v-text-field
              v-model="userLastName"
              label="ນາມສະກຸນ"
            ></v-text-field>
          </v-col>
        </v-row>
        <v-row>
          <v-col>
            <v-text-field
              v-model="username"
              label="ຊື່ຜູ້ໃຊ້ລະບົບ"
              :disabled="mode == 'edit'"
            ></v-text-field>
          </v-col>
          <v-col>
            <v-text-field
              type="email"
              v-model="userEmail"
              label="ອີເມວ"
            ></v-text-field>
          </v-col>
        </v-row>
        <v-row>
          <v-col>
            <v-text-field
              v-model="userPassword"
              label="ລະຫັດຜ່ານ"
              type="password"
            ></v-text-field>
          </v-col>
          <v-col>
            <v-text-field
              v-model="userConfirmPassword"
              label="ຢືນຢັນລະຫັດຜ່ານ"
              type="password"
            ></v-text-field>
          </v-col>
        </v-row>
        <v-row>
          <v-col>
            <v-select
              v-model="userRoleId"
              :items="userRoleList"
              item-value="id"
              item-text="roleName"
              label="ສິດຜູ້ໃຊ້"
            ></v-select>
          </v-col>
          <v-col>
            <v-file-input
              accept="image/*"
              label="ຮູບຜູ້ໃຊ້"
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
import { UserRole } from "@/models/User";

const App = namespace("App");

@Component({
  name: "UserEditorPage",
  metaInfo() {
    return {
      title: router.currentRoute.params.username
        ? "ແກ້ໄຂຂໍ້ມູນຜູ້ໃຊ້"
        : "ເພີ່ມຂໍ້ມູນຜູ້ໃຊ້"
    };
  }
})
export default class UserEditor extends Vue {
  private pageTitle = "ເພີ່ມ - ແກ້ໄຂຜູ້ໃຊ້ລະບົບ";

  private mode = "create";

  private userId: number | string | undefined = 0;
  private username: string | undefined = "";
  private userFirstName: string | undefined = "";
  private userLastName: string | undefined = "";
  private userEmail: string | undefined = "";
  private userPassword: string | undefined = "";
  private userConfirmPassword: string | undefined = "";
  private userRoleId: number | string | undefined = "";
  private userImage: File | undefined;

  private userRoleList: Array<UserRole> = [
    {
      id: 0,
      roleName: "N/A"
    }
  ];

  private snackbar = {
    show: false,
    text: "",
    type: "",
    timeout: 3000
  };

  @App.Mutation("SET_APP_NAV_TITLE") APP_NAV_TITLE!: (title: string) => void;

  onSelectImage(file: File) {
    if (file) {
      this.userImage = file;
    }
  }

  onSave() {
    if (this.userPassword !== this.userConfirmPassword) {
      this.snackbar = {
        show: true,
        text: "ລະຫັດຜ່ານບໍ່ກົງກັນ!",
        type: "error",
        timeout: 3000
      };
      return;
    }
    const form = new FormData();
    if (this.mode === "create") {
      form.append("firstname", this.userFirstName as string);
      form.append("lastname", this.userLastName as string);
      form.append("username", this.username as string);
      form.append("password", this.userPassword as string);
      form.append("email", this.userEmail as string);
      form.append("role", this.userRoleId as string);
      form.append("image", this.userImage as File);
      this.$http
        .post("/user/create", form)
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
          this.snackbar = {
            show: true,
            text: "ບໍ່ສໍາເລັດ!",
            type: "error",
            timeout: 3000
          };
          console.log(err);
        });
    } else {
      form.append("id", this.userId as string);
      form.append("firstname", this.userFirstName as string);
      form.append("lastname", this.userLastName as string);
      form.append("username", this.username as string);
      form.append("password", this.userPassword as string);
      form.append("email", this.userEmail as string);
      form.append("role", this.userRoleId as string);
      form.append("image", this.userImage as File);
      this.$http
        .put(`/user/update/${this.userId}`, form)
        .then(res => {
          if (res.status == 200) {
            this.snackbar = {
              show: true,
              text: "ແກ້ໄຂຂໍ້ມູນສໍາເລັດ",
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
  }

  onCancel() {
    this.$router.back();
  }

  getUserData() {
    this.$http
      .get(`/user/${this.$router.currentRoute.params.username}`)
      .then(res => {
        if (res.data) {
          this.userId = res.data.id;
          this.username = res.data.username;
          this.userFirstName = res.data.firstname;
          this.userLastName = res.data.lastname;
          this.userEmail = res.data.email;
          this.userRoleId = res.data.role;
        }
      })
      .catch(err => {
        console.log(err);
      });
  }

  getUserRoleData() {
    this.$http
      .get("/user/role/all")
      .then(res => {
        if (res.data.length > 0) {
          this.userRoleList = res.data;
        }
      })
      .catch(err => {
        console.log(err);
      });
  }

  mounted() {
    this.pageTitle = this.$router.currentRoute.params.username
      ? "ແກ້ໄຂຂໍ້ມູນຜູ້ໃຊ້ລະບົບ"
      : "ເພີ່ມຂໍ້ມູນຜູ້ໃຊ້ລະບົບ";
    this.APP_NAV_TITLE(this.pageTitle);
  }

  created() {
    if (this.$router.currentRoute.params.username) {
      this.mode = "edit";
      this.getUserData();
    } else {
      this.mode = "create";
    }
    this.getUserRoleData();
  }
}
</script>
