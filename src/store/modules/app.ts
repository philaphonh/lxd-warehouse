import { VuexModule, Module, Mutation } from "vuex-module-decorators";
@Module({
  namespaced: true,
  name: "App"
})
export default class App extends VuexModule {
  private appNavTitle = "Warehouse";
  private appNavDrawer = false;

  @Mutation
  public SET_APP_NAV_TITLE(title: string) {
    this.appNavTitle = title;
  }

  @Mutation
  public SET_APP_NAV_DRAWER(status: boolean) {
    this.appNavDrawer = status;
  }
}
