<template>
  <div class="login_contaniner">
    <div class="login_box">
      <!--      头像-->
      <div class="avatar_box">
        <img src="../../assets/TS_logo.png" alt="TS_logo">
      </div>
      <!--      登陆表单-->
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
            const { data:result } = await this.$http.post('common/auth_login', this.login_form);
            this.login_button_loading  = false;
            if (result.code !== 200) {
              this.$message.error('用户名或密码错误');
            } else {
              this.$message.success('登陆成功');
              window.sessionStorage.setItem('Authorization', result.data.token);
              window.sessionStorage.setItem('MyInfo', JSON.stringify(result.data.userinfo));
              this.$emit('LoginComplete');
              await this.$router.push('/index')
            }
          }
        })
      }
    }
  }
</script>

<style lang="less" scoped>
  .login_contaniner {
    background-color: #2b4b6b;
    height: 100%;
  }

  .login_box {
    width: 450px;
    height: 300px;
    background-color: #fff;
    border-radius: 3px;
    position: absolute;
    left: 50%;
    top: 50%;
    transform: translate(-50%, -50%);
    .avatar_box {
      height: 130px;
      width: 130px;
      border: 1px solid #eee;
      border-radius: 50%;
      padding: 10px;
      box-shadow: 0 0 10px #ddd;
      position: absolute;
      left: 50%;
      transform: translate(-50%, -50%);
      background-color: #fff;
      img {
        width: 100%;
        height: 100%;
        border-radius: 50%;
        background-color: #eee;
      }
    }
  }

  .btns {
    display: flex;
    justify-content: flex-end;
  }
  .login_form {
    position: absolute;
    bottom: 0;
    width: 100%;
    padding: 0 20px;
    box-sizing: border-box;
  }
</style>
