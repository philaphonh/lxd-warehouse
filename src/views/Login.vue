<template>
  <v-container>
    <v-card clas="pa-4">
      <v-card-title>ເຂົ້າສູ່ລະບົບຈັດການສາງລ້ານຊ້າງດິຈິຕອນ</v-card-title>
      <v-container>
        <v-form>
          <v-text-field v-model="username" label="ຊື່ຜູ້ໃຊ້"></v-text-field>
          <v-text-field
            v-model="password"
            label="ລະຫັດຜ່ານ"
            type="password"
          ></v-text-field>
          <p v-if="isError" class="red--text">{{ errorMessage }}</p>
          <v-row class="pa-4">
            <v-btn class="primary" @click="onLogin">ເຂົ້າສູ່ລະບົບ</v-btn>
          </v-row>
        </v-form>
      </v-container>
    </v-card>
  </v-container>
</template>
<script lang="ts">
import { Vue, Component } from "vue-property-decorator";
import { namespace } from "vuex-class";
import { AuthData } from "@/models/User";

const App = namespace("App");
const Auth = namespace("Auth");

@Component({
  name: "LoginPage",
  metaInfo() {
    return {
      title: "ເຂົ້າສູ່ລະບົບ"
    };
  }
})
export default class LoginPage extends Vue {
  private pageTitle = "ເຂົ້າສູ່ລະບົບ";

  private username = "";
  private password = "";
  private isError = false;
  private errorMessage = "";

  @Auth.State("authData") authData!: AuthData;
  @App.Mutation("SET_APP_NAV_TITLE") APP_NAV_TITLE!: (title: string) => void;
  @Auth.Mutation("LOGIN_STATUS") SET_LOGIN_STATUS!: (status: boolean) => void;
  @Auth.Mutation("SET_AUTH_DATA") AUTH_DATA!: (data: AuthData) => void;

  mounted() {
    this.APP_NAV_TITLE(this.pageTitle);
  }

  onLogin() {
    const form = new FormData();
    form.append("username", this.username);
    form.append("password", this.password);
    this.$http
      .post("/user/login", form)
      .then(res => {
        if (res.status == 200) {
          localStorage.setItem("token", res.data.token);
          localStorage.setItem("id", res.data.userId);
          localStorage.setItem("username", res.data.username);
          localStorage.setItem("role", res.data.role);
          localStorage.setItem("image", res.data.imageUrl);
          this.AUTH_DATA(res.data);
          this.SET_LOGIN_STATUS(true);
          console.log(this.authData);
          this.$router.push("/");
        } else {
          console.log(res.data);
          this.isError = true;
          this.errorMessage = `ບໍ່ສໍາເລັດ`;
        }
      })
      .catch(err => {
        console.log(err);
        this.isError = true;
        this.errorMessage = `ບໍ່ສໍາເລັດ`;
      });
  }
}
</script>
