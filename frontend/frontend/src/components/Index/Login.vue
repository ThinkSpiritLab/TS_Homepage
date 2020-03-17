<template>
  <el-card class="login-card">
    <div class="login_box">
      <div class="avatar_box">
        <img src="../../assets/TS_logo.png" alt="TS_logo">
      </div>
      <el-form ref="login_form_ref" :model="login_form" :rules="login_rules" class="login_form">
        <el-form-item prop="Stid">
          <el-input v-model="login_form.Stid" prefix-icon="el-icon-user"></el-input>
        </el-form-item>
        <el-form-item prop="Psw">
          <el-input v-model="login_form.Psw" prefix-icon="el-icon-lock" type="password"></el-input>
        </el-form-item>
        <el-form-item class="btns">
          <el-button type="primary" round @click="login" :loading="login_button_loading">登录</el-button>
          <el-button type="info" round @click="reset_form" >重置</el-button>
        </el-form-item>
      </el-form>
    </div>
  </el-card>
</template>

<script>
  export default {
    name: "Login",
    data() {
      return {
        login_form: {
          Stid: '',
          Psw: ''
        },
        login_rules: {
          Stid: [
            { required: true, message: '请输入学号/工号', trigger: 'blur' }
          ],
          Psw: [
            { required: true, message: '请输入密码', trigger: 'blur' }
          ]
        },
        login_button_loading: false
      }
    },
    methods: {
      reset_form: function () {
        this.$refs.login_form_ref.resetFields();
      },
      login: function () {
        this.$refs.login_form_ref.validate(async (valid)=> {
          if (!valid) {
            return;
          }
          else {
            this.login_button_loading  = true;
            const { data:result } = await this.$http.post('/auth_login', this.login_form);
            this.login_button_loading  = false;
            if (result.code !== 200) {
              this.$message.error('用户名或密码错误');
            } else {
              this.$message.success('登陆成功');
              window.sessionStorage.setItem('Authorization', result.data.token);
              window.sessionStorage.setItem('MyInfo', JSON.stringify(result.data.userinfo));
              this.$emit('LoginComplete');
              await this.$router.push('/main')
            }
          }
        })
      },
    },
  }
</script>

<style lang="less" scoped>
  .login-card {
    text-align: center;
    position: relative;
    width: 40%;
    height: 400px;
    left: 30%;
    margin-top: 100px;
    margin-bottom: 100px;
    border-radius: 10%;
  }
  .avatar_box {
    height: 130px;
    width: 130px;
    border: 1px solid #eee;
    border-radius: 50%;
    padding: 10px;
    box-shadow: 0 0 10px #ddd;
    background-color: #fff;
    margin:0 auto;
    img {
      width: 100%;
      height: 100%;
      border-radius: 50%;
      background-color: #eee;
    }
  }

  .btns {
    display: flex;
    justify-content: flex-end;
  }
  .login_form {
    margin:0 auto;
    width: 90%;
    padding-top: 30px;
    box-sizing: border-box;
  }
</style>
