<template>
  <el-container class="index-container">
    <el-header>
      <el-row type="flex">
        <el-col :span="4"><div>LOGO</div></el-col>
        <el-col :span="15"><div>Menu</div></el-col>
        <el-col :span="5">
              <div v-if="hasLogIn">
            <el-dropdown @command="handleUserTipCommand">
              <el-button type="info">
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
        hasLogIn: false
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
    background-color: #99a9bf;
    .el-row {
      height: 60px;
      div {
        display: flex;
        align-items: center;
      }
    }
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
