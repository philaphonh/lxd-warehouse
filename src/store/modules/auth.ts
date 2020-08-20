import { VuexModule, Module, Mutation } from "vuex-module-decorators";
import { AuthData } from "@/models/User";
@Module({
  namespaced: true,
  name: "Auth"
})
export default class Auth extends VuexModule {
  private authData!: AuthData;
  private isLoggedIn = false;

  @Mutation
  public LOGIN_STATUS(status: boolean) {
    this.isLoggedIn = status;
  }

  @Mutation
  public SET_AUTH_DATA(data: AuthData) {
    this.authData = {
      userId: data.userId,
      username: data.username,
      role: data.role,
      imageUrl: data.imageUrl,
      token: data.token
    };
  }

  @Mutation
  public RETRIVE_AUTH_DATA() {
    if (
      localStorage.getItem("username") !== null &&
      localStorage.getItem("role") !== null &&
      localStorage.getItem("image") !== null &&
      localStorage.getItem("token") !== null
    ) {
      this.authData = {
        userId: localStorage.getItem("id") as string,
        username: localStorage.getItem("username") as string,
        role: Number.parseInt(localStorage.getItem("role") as string),
        imageUrl: localStorage.getItem("image") as string,
        token: localStorage.getItem("token") as string
      };
      console.log("Retrieved!");
    }
  }

  @Mutation
  public CLEAR_AUTH_DATA() {
    this.authData = {
      userId: "",
      username: "",
      role: 0,
      imageUrl: "",
      token: ""
    };
  }
}
