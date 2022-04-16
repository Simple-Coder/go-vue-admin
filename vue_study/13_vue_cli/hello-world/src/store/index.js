import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    myNum: 0
  },
  mutations: {
    addNum(state) {
      state.myNum++
    },
    subNum(state) {
      state.myNum--
    }
  },
  actions: {
    asyncAdd(context) {
      setTimeout(() => {
        context.commit('addNum')
      }, 5000);
    }
  },
  modules: {
  }
})
