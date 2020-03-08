<template>
  <el-container class="index-container">
    <el-header style="height: 75px">
      <el-row type="flex" style="margin-top: 20px; height: 55px">
        <el-col :span="3">
          <div>
            <img src="../assets/TS_logo.png" alt="TS_logo" class="topIMG1">
            <span style="margin-left: 5px">Think Spirit Lab</span>
          </div>
        </el-col>
        <el-col :span="4">
          <div>
            <img src="../assets/icpc.png" alt="TS_logo" class="topIMG2">
            <span style="margin-left: 10px">NUIST CPC Team</span>
          </div>
        </el-col>
        <el-col :span="12">
          <div style="margin-right: auto; margin-left: auto">
            <el-menu :default-active="$route.path" class="el-menu-demo" mode="horizontal">
              <el-menu-item index="/main">首页</el-menu-item>
              <el-menu-item index="/announcement">公告</el-menu-item>
              <el-menu-item index="/members">成员</el-menu-item>
              <el-menu-item index="/contests">竞赛历史</el-menu-item>
              <el-menu-item index="/news">活动与新闻</el-menu-item>
              <el-menu-item index="/trains">训练与积分</el-menu-item>
              <el-menu-item index="/publicity">内部公示</el-menu-item>
              <el-menu-item index="/resource">资源</el-menu-item>
            </el-menu>
          </div>
        </el-col>
        <el-col :span="5">
              <div v-if="hasLogIn">
            <el-dropdown @command="handleUserTipCommand">
              <el-button type="primary">
                {{UserTipInfo}}<i class="el-icon-arrow-down el-icon--right"></i>
              </el-button>
              <el-dropdown-menu slot="dropdown">
                <el-dropdown-item command="1">修改资料</el-dropdown-item>
                <el-dropdown-item command="2">退出</el-dropdown-item>
              </el-dropdown-menu>
            </el-dropdown>
            <el-button v-show="consolePrivilege" type="primary" class="btnConsole" @click="goToConsole">后台</el-button>
          </div>
          <div v-else>
            <el-button type="primary" @click="goToLogin">登录</el-button>
          </div>
        </el-col>
      </el-row>
    </el-header>
    <el-main>
      <div class="index-main-body">
        <router-view @LoginComplete="checkLoginState"></router-view>
      </div>
    </el-main>
    <el-footer>
      Footer
    </el-footer>
  </el-container>
</template>

<script>
  export default {
    name: "Index",
    data() {
      return {
        userTipInfo: {
          stid: '',
          name: '',
          grade: 0,
          identity: 0,
          privilege: 0
        },
        hasLogIn: false,
      }
    },
    methods: {
      goToLogin: function () {
        this.$router.push("/login");
      },
      checkLoginState: function () {
        this.userTipInfo = JSON.parse(window.sessionStorage.getItem('MyInfo'));
        this.hasLogIn = this.userTipInfo !== null
      },
      handleUserTipCommand(command) {
        if (command === '1') {

        } else if (command === '2') {
          this.Logout();
        }
      },
      Logout: function () {
        window.sessionStorage.clear();
        this.userTipInfo = null;
        this.hasLogIn = false;
        this.$message.success('退出成功');
      },
      goToConsole: function () {
        this.$router.push('/console')
      }
    },
    computed: {
      consolePrivilege: function () {
        if (!this.userTipInfo)
          return false;
        return this.userTipInfo.privilege === 1 || this.userTipInfo.privilege === 2 ||
            this.userTipInfo.privilege === 3 || this.userTipInfo.privilege === 4;
      },
      /**
       * @return {string}
       */
      UserTipInfo: function() {
        if (!this.userTipInfo)
          return '----';
        return this.userTipInfo.stid + " " + this.userTipInfo.name
      }
    },
    mounted() {
      this.checkLoginState()
    },
  }
</script>
<!--this.$store.getters.getAuthorization-->
<!--this.$store.dispatch('setAuthorization', '123456')-->
<style lang="less" scoped>
  .index-container {
    height: 100%;
  }
  .el-header {
    background-color: #F4F4F4;
    .el-row {
      height: 60px;
      div {
        display: flex;
        align-items: center;
      }
    }
  }
  .topIMG1 {
    width: 50px;
    height: 50px;
    border-radius: 50%;
  }
  .topIMG2 {
    width: 84px;
    height: 50px;
    border-radius: 20%;
  }
  .el-main {
    background-color: #F4F4F4;
  }
  .index-main-body {
    border-radius: 10px;
  }
  .btnConsole {
    margin-left: 10px;
  }

</style>
