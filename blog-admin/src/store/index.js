import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

import dialog_store from "./dialog_store.js";

export default new Vuex.Store({
  modules: {
    dialog: dialog_store
  }
})