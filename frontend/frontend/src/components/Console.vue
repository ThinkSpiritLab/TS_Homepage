<template>
  <el-container class="home-container">
    <el-header>
      <div>
        <img src="../assets/TS_logo.png" alt="TS_logo">
        <span>Think Spirit 控制台</span>
      </div>
      <div>
        <div class="userTip">{{UserTipInfo}}</div>
        <el-button type="danger" round @click="log_out">退出</el-button>
        <el-button type="primary" @click="goToIndex">前台</el-button>
      </div>
    </el-header>
    <el-container>
      <el-aside width="200px">
        <el-menu background-color="#333744" text-color="#fff" active-text-color="#409BFF"
                :unique-opened="false" :router="true" :default-active="activePath">
          <el-submenu :index="'console_'+item.mid" v-for="item in menuList" :key="item.mid">
            <template slot="title">
              <span>{{item.mname}}</span>
            </template>
            <el-menu-item :index="'console_'+subItem.path" v-for="subItem in item.children"
                          :key="subItem.mid" @click="savaNavState(subItem.path)">
              <template slot="title">
                <i class="el-icon-menu"></i>
                <span>{{subItem.mname}}</span>
              </template>
            </el-menu-item>
          </el-submenu>
        </el-menu>
      </el-aside>
      <el-main>
        <transition name="el-zoom-in-center">
          <router-view></router-view>
        </transition>
      </el-main>
    </el-container>
  </el-container>
</template>

<script>
  export default {
    name: "Console",
    data() {
      return {
        userTipInfo: {
          stid: '',
          name: '',
          grade: 0,
          identity: 0,
          privilege: 0
        },
        menuList: [],
        activePath: ''
      }
    },
    created() {
      this.checkLoginState();
      this.getMenuList();
      this.activePath = window.sessionStorage.getItem('activePath');
    },
    methods: {
      checkLoginState: function () {
        this.userTipInfo = JSON.parse(window.sessionStorage.getItem('MyInfo'));
      },
      goToIndex: function() {
        this.$router.push('/index');
      },
      log_out: function () {
        window.sessionStorage.clear();
        this.userTipInfo = null;
        this.$message.success('退出成功');
        this.$router.push('/index');
      },
      getMenuList: async function () {
        const {data: res} = await this.$http.get('/console/menu_list')
        if (res.code !== 200) {
          this.$message.error(res.msg);
          await this.$router.push('/main');
        }
        this.menuList = res.data;
      },
      savaNavState: function (activePath) {
        window.sessionStorage.setItem('activePath', activePath);
        this.activePath = activePath;
      }
    },
    computed: {
      /**
       * @return {string}
       */
      UserTipInfo: function() {
        if (!this.userTipInfo)
          return '----';
        return this.userTipInfo.stid + " " + this.userTipInfo.name
      }
    }
  }
</script>

<style lang="less" scoped>
  .home-container {
    height: 100%;
  }
  .el-header {
    background-color: #373d41;
    display: flex;
    justify-content: space-between;
    align-items: center;
    color: #ffffff;
    font-size: 20px;
    font-family: "Helvetica Neue",Helvetica,"PingFang SC","Hiragino Sans GB","Microsoft YaHei","微软雅黑",Arial,sans-serif;
    div {
      display: flex;
      align-items: center;
      span {
        margin-left: 15px;
      }
    }
  }
  .el-aside {
    background-color: #333744;
    .el-menu {
      border-right: 0;
    }
  }
  .el-main {
    background-color: #EAEDF1;
  }
  .iconfont {
    margin-right: 10px;
  }
  .toggle-button {
    background-color: #4A5064;
    font-size: 10px;
    line-height: 24px;
    color: #fff;
    text-align: center;
    letter-spacing: 0.2em;
    cursor: pointer;
  }
  img {
    width: 50px;
    height: 50px;
    border-radius: 50%;
  }
  .userTip {
    margin-right: 10px;
    font-size: 15px;
  }
</style>
