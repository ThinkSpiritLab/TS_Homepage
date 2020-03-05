// 弃用
import Vue from 'vue'
import Vuex from 'vuex'
import VueXAlong from 'vuex-along'
import fa from "element-ui/src/locale/lang/fa";

Vue.use(Vuex)

export default new Vuex.Store({
  state: {
    Authorization: "",
    MyInfo: {}
  },
  mutations: {
    set_authorization: function (state, token) {
      state.Authorization = token
    },
    set_myinfo: function (state, my_info) {
      state.MyInfo = my_info
    }
  },
  actions: {
    setAuthorization: function ({commit, state},token) {
      commit("set_authorization", token)
    },
    setMyInfo: function ({commit, state},my_info) {
      commit("set_myinfo", my_info)
    }
  },
  getters: {
    getAuthorization: function (state) {
      return state.Authorization
    },
    getMyInfo: function (state) {
      return state.MyInfo
    }
  },
  plugins: [VueXAlong({
    name: 'along',     //存放在localStroage或者sessionStroage 中的名字
    local: false,      //是否存放在local中  false 不存放 如果存放按照下面session的配置配
    session: { list: [], isFilter: false } //如果值不为false 那么可以传递对象 其中 当isFilter设置为true时， list 数组中的值就会被过滤调,这些值不会存放在seesion或者local中
  })]
})
