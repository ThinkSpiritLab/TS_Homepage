<template>
    <el-container class="index-container">
      <el-main>
        <div class="header">
          <el-row type="flex" style="margin-top: 20px; height: 55px">
            <el-col :span="3">
              <div>
                <img src="../assets/TS_logo.png" alt="TS_logo" class="topIMG1">
                <span style="margin-left: 5px" class="web-font">Think Spirit Lab</span>
              </div>
            </el-col>
            <el-col :span="4">
              <div style="margin-left: 20px">
                <img src="../assets/icpc.png" alt="TS_logo" class="topIMG2">
                <span style="margin-left: 10px" class="web-font">NUIST CPC Team</span>
              </div>
            </el-col>
            <el-col :span="12">
              <div style="margin-right: auto; margin-left: auto">
                <el-menu :default-active="$route.path" class="menuBar" mode="horizontal" router>
                  <el-menu-item index="/main">首页</el-menu-item>
                  <el-menu-item index="/members">成员</el-menu-item>
                  <el-menu-item index="/contests">竞赛历史</el-menu-item>
                  <el-menu-item index="/announcement">公告</el-menu-item>
                  <el-menu-item index="/news">活动与新闻</el-menu-item>
                  <el-menu-item index="/about">关于</el-menu-item>
                  <el-submenu index="platform" style="height: 40px">
                    <template slot="title">Think Spirit 网站群</template>
                    <el-menu-item><a href="https://oj.vicz.cn/" target="_blank">TSOJ v5</a></el-menu-item>
                    <el-menu-item><a href="https://nuistcpc.club/" target="_blank">NUIST CPC 集训队内部网站</a></el-menu-item>
                    <el-menu-item disabled>语言学习暨计算机等级考试学习系统（张斌杰毕设，还没写）</el-menu-item>
                    <el-menu-item disabled>课程辅助学习平台（曹梦琪毕设，开发中）</el-menu-item>
                    <el-menu-item disabled>实验报告管理平台（倪文卿毕设，开发中）</el-menu-item>
                  </el-submenu>
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
        </div>
        <div class="index-main-body">
          <router-view @LoginComplete="checkLoginState"></router-view>
        </div>
        <div class="footer">
          <el-row>
            <el-col :span="8">
              © 2020 PenguinZhang @ Think Spirit Lab 版权所有
            </el-col>
            <el-col :span="8">
              <a href="http://www.beian.miit.gov.cn/">
                苏ICP备18001271号-2
              </a>
            </el-col>
            <el-col :span="8" style="height: 20px">
              <a target="_blank" href="" style="height: 20px">
                <img src="../assets/police.png">
                <span> 苏公网安备 还没批下来占个位置 号</span>
              </a>
            </el-col>
          </el-row>
        </div>
      </el-main>
    </el-container>
</template>

<script>
  export default {
    name: "Index",
    data() {
      return {
        userTipInfo: {
          uid: 0,
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
          this.$router.push('/editPersonalInfo/'+this.userTipInfo.uid)
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
      },
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
  .header {
    background-color: #F4F4F4;
    margin-bottom: 15px;
    .el-row {
      height: 60px;
      div {
        display: flex;
        align-items: center;
      }
    }
  }
  .menuBar {
    height: 40px;
    .el-menu-item {
      height: 40px;
      line-height: 40px;
    }
    .el-submenu ::v-deep .el-submenu__title {
      height: 40px !important;
      line-height: 40px !important;
    }
  }
  .el-main {
    background-color: #F4F4F4;
    padding-top: 0;
  }
  .index-container {
    height: 100% !important;
  }
  .topIMG1 {
    width: 50px;
    height: 50px;
    border-radius: 50%;
  }
  .topIMG2 {
    width: 84px;
    height: 50px;
    /*border-radius: 20%;*/
  }
  .index-main-body {
    border-radius: 10px;
  }
  .btnConsole {
    margin-left: 10px;
  }
  .textstyle {
    font-family: "Helvetica Neue",Helvetica,"PingFang SC","Hiragino Sans GB","Microsoft YaHei","微软雅黑",Arial,sans-serif;
  }
  @font-face {
    font-family: 'webfont';
    font-display: swap;
    src: url('//at.alicdn.com/t/webfont_snovg14vkr.eot'); /* IE9*/
    src: url('//at.alicdn.com/t/webfont_snovg14vkr.eot?#iefix') format('embedded-opentype'), /* IE6-IE8 */
    url('//at.alicdn.com/t/webfont_snovg14vkr.woff2') format('woff2'),
    url('//at.alicdn.com/t/webfont_snovg14vkr.woff') format('woff'), /* chrome、firefox */
    url('//at.alicdn.com/t/webfont_snovg14vkr.ttf') format('truetype'), /* chrome、firefox、opera、Safari, Android, iOS 4.2+*/
    url('//at.alicdn.com/t/webfont_snovg14vkr.svg#AlibabaPuHuiTiM') format('svg'); /* iOS 4.1- */
  }
  .web-font{
    font-family:"webfont" !important;
    font-size:16px;font-style:normal;
    -webkit-font-smoothing: antialiased;
    -webkit-text-stroke-width: 0.2px;
    -moz-osx-font-smoothing: grayscale;
  }
  .footer {
    margin-top: 20px;
    margin-bottom: 5px;
    text-align: center;
    color: #999;
    font-size: 15px;
  }
</style>
<style>
  .el-scrollbar__wrap
  {
    overflow-x: hidden !important;
  }
  .el-scrollbar__view {
    height: 100% !important;
  }
  a {
    color: #999;
    text-decoration: none
  }
  a:visited {
    color: #999;
    text-decoration: none
  }

  ::-webkit-scrollbar {
    /*滚动条整体样式*/
    width : 10px;  /*高宽分别对应横竖滚动条的尺寸*/
    height: 1px;
  }
  ::-webkit-scrollbar-thumb {
    /*滚动条里面小方块*/
    border-radius: 10px;
    box-shadow   : inset 0 0 5px rgba(0, 0, 0, 0.2);
    background   : #535353;
  }
  ::-webkit-scrollbar-track {
    /*滚动条里面轨道*/
    box-shadow   : inset 0 0 5px rgba(0, 0, 0, 0.2);
    border-radius: 10px;
    background   : #ededed;
  }
</style>
