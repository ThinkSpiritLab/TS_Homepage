<template>
  <div class="login_contaniner">
    <div class="login_box">
      <!--            头像-->
      <div class="avatar_box">
        <img src="../../assets/TS_logo.png" alt="TS_logo">
      </div>
      <!--            登陆表单-->
      <el-form ref="login_form_ref" :model="login_form" :rules="login_rules" class="login_form">
        <el-form-item prop="Stid">
          <el-input v-model="login_form.Stid" prefix-icon="iconfont icon-user"></el-input>
        </el-form-item>
        <el-form-item prop="Psw">
          <el-input v-model="login_form.Psw" prefix-icon="iconfont icon-3702mima" type="password"></el-input>
        </el-form-item>
        <el-form-item class="btns">
          <el-button type="primary" round @click="login" :loading="login_button_loading">登录</el-button>
          <el-button type="info"round @click="reset_form" >重置</el-button>
        </el-form-item>
      </el-form>
    </div>
  </div>
</template>

<script>
  export default {
    name: "v"
  }
</script>

<style scoped>

</style>
