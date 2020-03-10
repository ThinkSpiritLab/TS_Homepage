<template>
  <el-container class="home-container" v-cloak>
    <el-header>
      <div>
        <img src="../assets/TS_logo.png" alt="TS_logo" class="topIMG">
        <span>Think Spirit 控制台</span>
      </div>
      <div>
        <div class="userTip">{{UserTipInfo}}</div>
        <el-button type="danger" round @click="log_out">退出</el-button>
        <el-button type="primary" @click="goToIndex">前台</el-button>
      </div>
    </el-header>
    <el-container>
      <el-aside width="210px">
        <el-menu el-menu-vertical-demo :unique-opened="false" :router="true" :default-active="$route.path">
          <el-submenu :index="'/console_'+item.mid" v-for="item in menuList" :key="item.mid">
            <template slot="title">
              <span>{{item.mname}}</span>
            </template>
            <el-menu-item :index="'/console_'+subItem.path" v-for="subItem in item.children"
                          :key="subItem.mid">
              <template slot="title">
                <i class="el-icon-menu"></i>
                <span>{{subItem.mname}}</span>
              </template>
            </el-menu-item>
          </el-submenu>
        </el-menu>
      </el-aside>
      <el-main>
<!--        <transition name="el-fade-in">-->
          <router-view :userTipInfo="userTipInfo"></router-view>
<!--        </transition>-->
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
      }
    },
    created() {
      this.checkLoginState();
      this.getMenuList();
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
          this.log_out();
        }
        this.menuList = res.data;
      },
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
    background-color: #EAEDF1;
    .el-menu {
      margin-top: 10px;
      margin-left: 10px;
      margin-bottom: 10px;
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
  .topIMG {
    width: 50px;
    height: 50px;
    border-radius: 50%;
  }
  .userTip {
    margin-right: 10px;
    font-size: 15px;
  }
</style>
