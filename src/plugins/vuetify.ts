import Vue from "vue";
import Vuetify from "vuetify/lib";
// import "@mdi/font/css/materialdesignicons.css";

Vue.use(Vuetify);

export default new Vuetify({
  icons: {
    iconfont: "mdiSvg"
  },
  theme: {
    themes: {
      light: {
        primary: "#364F6B",
        secondary: "#43DDE6",
        accent: "#FC5185"
      },
      dark: {
        primary: "#364F6B",
        secondary: "#43DDE6",
        accent: "#FC5185"
      }
    }
  }
});
