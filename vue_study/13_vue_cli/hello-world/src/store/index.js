import Vue from 'vue'
import Vuex from 'vuex'
import { xd } from './xd.js'

Vue.use(Vuex)



export default new Vuex.Store({
  modules: {
    xd: xd
  },

})
