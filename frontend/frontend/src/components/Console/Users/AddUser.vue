<template>
    <el-card class="box-card">
      <el-page-header @back="goToHome" content="添加成员"></el-page-header>
      <el-divider></el-divider>
      <el-row :gutter="100">
        <el-col :span="12">
          <div class="divForm">
            <el-form :model="addForm.user" :rules="addFormRules" :disabled="loading"
                     ref="addFormRefUser" label-width="90px">
              <el-form-item label="学号/工号" prop="stid">
                <el-input v-model="addForm.user.stid"></el-input>
              </el-form-item>
              <el-form-item label="姓名" prop="name">
                <el-input v-model="addForm.user.name"></el-input>
              </el-form-item>
              <el-form-item label="身份" prop="identity">
                <el-radio-group v-model="addForm.user.identity">
                  <el-radio label="1">教师/教练</el-radio>
                  <el-radio label="2">正式成员</el-radio>
                  <el-radio label="3">预备成员</el-radio>
                </el-radio-group>
              </el-form-item>
              <transition name="el-zoom-in-top">
                <el-form-item label="年级" prop="grade" v-if="addForm.user.identity==='2'||addForm.user.identity==='3'">
                  <el-input v-model="addForm.user.grade"></el-input>
                </el-form-item>
              </transition>
              <el-form-item label="权限" prop="privilege">
                <el-select v-model="addForm.user.privilege" placeholder="请选择">
                  <el-option key="p1" label="root" value="1" v-if="userTipInfo.privilege===1"></el-option>
                  <el-option key="p2" label="超级管理员" value="2" v-if="userTipInfo.privilege===1"></el-option>
                  <el-option key="p3" label="团队建设管理员" value="3" v-if="userTipInfo.privilege<=2"></el-option>
                  <el-option key="p4" label="竞赛训练管理员" value="4" v-if="userTipInfo.privilege<=2"></el-option>
                  <el-option key="p5" label="普通权限" value="5" v-if="userTipInfo.privilege<=3"></el-option>
                </el-select>
              </el-form-item>
            </el-form>
            <el-form :model="addForm.user_password" :rules="addFormRules" :disabled="loading"
                     ref="addFormRefUserPassword" label-width="90px">
              <el-form-item label="密码" prop="psw">
                <el-input v-model="addForm.user_password.psw" show-password></el-input>
              </el-form-item>
            </el-form>
          </div>
        </el-col>
        <el-col :span="12">
          <div class="divForm">
            <el-form :model="addForm.user_detail" :rules="addFormRules" :disabled="loading"
                     ref="addFormRefUserDetail" label-width="90px">
              <el-form-item label="邮箱" prop="email">
                <el-input v-model="addForm.user_detail.email"></el-input>
              </el-form-item>
              <el-form-item label="手机" prop="mobile">
                <el-input v-model="addForm.user_detail.mobile"></el-input>
              </el-form-item>
              <el-form-item label="QQ" prop="QQ">
                <el-input v-model="addForm.user_detail.QQ"></el-input>
              </el-form-item>
              <el-form-item label="URL" prop="URL">
                <el-input v-model="addForm.user_detail.URL"></el-input>
              </el-form-item>
              <el-form-item label="简介" prop="introduction">
                <el-input type="textarea" :rows="5" v-model="addForm.user_detail.introduction">
                </el-input>
              </el-form-item>
            </el-form>
          </div>
        </el-col>
      </el-row>
      <div>
        <el-button style="margin-top: 12px;" @click="addUser" :loading="loading">提交</el-button>
        <el-button style="margin-top: 12px;" @click="formReset" :disabled="loading">重置</el-button>
      </div>
    </el-card>
</template>

<script>
  export default {
    props: ['userTipInfo'],
    name: "AddUser",
    data() {
      const checkEmail = (rule, value, cb) => {
        const regEmail = /^([a-zA-Z0-9_-])+@([a-zA-Z0-9_-])+(\.[a-zA-Z0-9_-])+/;
        if (regEmail.test(value) || !value) {
          return cb()
        }
        cb(new Error('请输入合法的邮箱'))
      };
      const checkMobile = (rule, value, cb) => {
        const regMobile = /^(0|86|17951)?(13[0-9]|15[012356789]|17[678]|18[0-9]|14[57])[0-9]{8}$/;
        if (regMobile.test(value) || !value) {
          return cb()
        }
        cb(new Error('请输入合法的手机号'))
      };
      const checkGrade = (rule, value, cb) => {
        if (value || this.addForm.user.identity==='1') {
          return cb()
        }
        cb(new Error('请输入年级'))
      };
      return {
        addForm: {
          user: {
            stid: '',
            name: '',
            identity: '2',
            grade: '',
            privilege: '5',
          },
          user_password: {
            psw: ''
          },
          user_detail: {
            email: '',
            QQ: '',
            URL: '',
            mobile: '',
            introduction: ''
          }
        },
        addFormRules: {
          stid: [
            {required: true, message: '请输入学号/工号', trigger: 'blur'},
          ],
          name: [
            {required: true, message: '请输入姓名', trigger: 'blur'},
          ],
          identity: [
            {required: true, message: '请选择身份', trigger: 'blur'},
          ],
          grade: [
            {validator: checkGrade, trigger: 'blur'}
          ],
          psw: [
            {required: true, message: '请输入密码', trigger: 'blur'},
            {min: 6, max: 20, message: '长度在 6 到 20 个字符', trigger: 'blur'}
          ],
          email: [
            {validator: checkEmail, trigger: 'blur'}
          ],
          mobile: [
            {validator: checkMobile, trigger: 'blur'}
          ]
        },
        loading: false,
      };
    },
    methods: {
      addUser: async function () {
        this.loading = true;
        let formValid = true;
        this.$refs.addFormRefUser.validate((valid) => {
          if (!valid) {
            formValid = false
          }
        });
        this.$refs.addFormRefUserPassword.validate((valid) => {
          if(!valid) {
            formValid = false
          }
        });
        this.$refs.addFormRefUserDetail.validate((valid) => {
          if(!valid) {
            formValid = false
          }
        });
        if (formValid) {
          const { data:result } = await this.$http.post('/console/user_regist', this.addForm);
          if (result.code !== 200) {
            this.$message.error(result.msg);
          } else {
            this.$message.success('添加成员成功');
            this.formReset();
          }
        }
        this.loading = false;
      },
      formReset: function () {
        this.$refs.addFormRefUser.clearValidate();
        this.$refs.addFormRefUserPassword.resetFields();
        this.$refs.addFormRefUserDetail.clearValidate();
        this.$refs.addFormRefUser.resetFields();
        this.$refs.addFormRefUserPassword.clearValidate();
        this.$refs.addFormRefUserDetail.resetFields();
      },
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
  .divForm {
    margin-top: 20px;
    text-align: left;
  }
</style>
