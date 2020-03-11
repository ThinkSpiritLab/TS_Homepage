<template>
  <el-card class="box-card">
    <el-page-header @back="goToHome" content="添加公告"></el-page-header>
  </el-card>
</template>

<script>
  export default {
    name: "AddBulletin",
    methods: {
      goToHome: function () {
        this.$router.push('/console_home');
      }
    }
  }
</script>

<style scoped>
  .box-card {
    text-align: center;
    position: relative;
    width: 96%;
    left: 2%;
    top: 5%;
  }
</style>
