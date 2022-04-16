import Vue from 'vue'
import Vuex from 'vuex'

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    myNum: 0
  },
  mutations: {
    addNum(state, params) {
      // state.myNum++
      state.myNum = state.myNum + params.num
    },
    subNum(state) {
      state.myNum--
    }
  },
  actions: {
    asyncAdd(context, params) {
      setTimeout(() => {
        context.commit('addNum', params)
      }, 5000);
    }
  },
  modules: {
  },
  //在state之前做数据处理
  getters: {
    getNum: (state) => state.myNum + 10
  }
})
