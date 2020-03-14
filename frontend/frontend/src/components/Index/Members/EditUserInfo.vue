<template>
  <el-card class="box-card">
    <div slot="header" class="clearfix" style="text-align: left">
      <span>修改个人资料</span>
    </div>
    <div>
      <el-form ref="infoForm" :model="infoForm" label-width="80px" :disabled="isLoading">
        <el-row>
          <el-col :span="8" style="text-align: center">
            <el-upload class="avatar-uploader" action="default" :http-request = "customUpload"
                       :show-file-list="false" :before-upload="beforeAvatarUpload"
                       :multiple="false" :auto-upload="true">
              <img v-if="infoForm.avatarUrl" :src="infoForm.avatarUrl" class="avatar" alt="img">
              <i v-else class="el-icon-plus avatar-uploader-icon"></i>
            </el-upload>
            <div class="divText">
              <div>点击上方图片上传或修改个人证件照</div>
              <div>只能上传 jpg/jpeg/png 文件，且不超过 1MB</div>
            </div>
          </el-col>
          <el-col :span="8" style="text-align: left">
            <el-form-item label="学号/工号">
              <div>{{infoForm.stid}}</div>
            </el-form-item>
            <el-form-item label="姓名">
              <div>{{infoForm.name}}</div>
            </el-form-item>
            <el-form-item label="身份">
              <div class="banner">
                <el-alert title="教师/教练" type="error" effect="dark" :closable="false" center
                          v-if="infoForm.identity==='1'"></el-alert>
                <el-alert title="正式成员" type="success" effect="dark" :closable="false" center
                          v-else-if="infoForm.identity==='2'"></el-alert>
                <el-alert title="预备成员" type="info" effect="dark" :closable="false" center
                          v-else-if="infoForm.identity==='3'"></el-alert>
              </div>
            </el-form-item>
            <el-form-item label="权限">
              <div class="banner">
                <el-alert title="root" type="error" effect="dark" :closable="false" center
                          v-if="infoForm.privilege==='1'"></el-alert>
                <el-alert title="超级管理员" type="success" effect="dark" :closable="false" center
                          v-else-if="infoForm.privilege==='2'"></el-alert>
                <el-alert title="团队建设管理员" type="warning" effect="dark" :closable="false" center
                          v-else-if="infoForm.privilege==='3'"></el-alert>
                <el-alert title="竞赛训练管理员" type="warning" effect="dark" :closable="false"
                          v-else-if="infoForm.privilege==='4'"></el-alert>
                <el-alert title="普通权限" type="info" effect="dark" :closable="false"
                          v-else-if="infoForm.privilege==='5'"></el-alert>
              </div>
            </el-form-item>
          </el-col>
          <el-col :span="8">
            <el-form-item label="Email" style="width: 400px">
              <el-input v-model="infoForm.email"></el-input>
            </el-form-item>
            <el-form-item label="地址" style="width: 400px">
              <el-input v-model="infoForm.address"></el-input>
            </el-form-item>
            <el-form-item label="QQ" style="width: 400px">
              <el-input v-model="infoForm.QQ"></el-input>
            </el-form-item>
            <el-form-item label="URL" style="width: 400px">
              <el-input v-model="infoForm.URL"></el-input>
            </el-form-item>
          </el-col>
        </el-row>
        <el-divider content-position="left">硕士及以上教育经历</el-divider>
        <div>
          <el-row v-for="(val, index) in education" :key='index'>
            <el-col :span="4">
              <el-form-item label="开始" label-width="50px">
                <el-date-picker v-model="val.begin" type="year" placeholder="开始年份"
                                value-format="yyyy" style="width: 120px"></el-date-picker>
              </el-form-item>
            </el-col>
            <el-col :span="4">
              <el-form-item label="结束" label-width="50px">
                <el-date-picker v-model="val.end" type="year" placeholder="结束年份"
                                value-format="yyyy" style="width: 120px"></el-date-picker>
              </el-form-item>
            </el-col>
            <el-col :span="5">
              <el-form-item label="学校" label-width="50px">
                <el-input v-model="val.university"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="5" label-width="50px">
              <el-form-item label="学院" label-width="50px">
                <el-input v-model="val.college"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="4" label-width="50px">
              <el-form-item label="学历" label-width="50px">
                <el-select v-model="val.level" placeholder="请选择">
                  <el-option key="1" label="硕士研究生" value="1"></el-option>
                  <el-option key="2" label="博士研究生" value="2"></el-option>
                  <el-option key="3" label="博士以上" value="3"></el-option>
                </el-select>
              </el-form-item>
            </el-col>
            <el-col :span="2">
              <el-button type="danger" icon="el-icon-delete"
                         style="width: 80px" @click.prevent="deleteEducation(val)"></el-button>
            </el-col>
          </el-row>
          <div>
            <el-alert title="无硕士及以上教育记录" type="info" center show-icon :closable="false"
                      v-if="education.length===0" style="margin-bottom: 10px"></el-alert>
          </div>
          <div style="text-align: end; margin-right: 20px">
            <el-button type="primary" @click="addEducation">添加硕士及以上教育经历</el-button>
          </div>
        </div>
        <el-divider content-position="left">职业经历</el-divider>
        <div>
          <el-row v-for="(val, index) in career" :key='index'>
            <el-col :span="4">
              <el-form-item label="开始" label-width="50px">
                <el-date-picker v-model="val.begin" type="year" placeholder="开始年份"
                                value-format="yyyy" style="width: 120px"></el-date-picker>
              </el-form-item>
            </el-col>
            <el-col :span="4">
              <el-form-item label="结束" label-width="50px">
                <el-date-picker v-model="val.end" type="year" placeholder="结束年份"
                                value-format="yyyy" style="width: 120px"></el-date-picker>
              </el-form-item>
            </el-col>
            <el-col :span="7">
              <el-form-item label="公司" label-width="50px">
                <el-input v-model="val.company"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="7" label-width="50px">
              <el-form-item label="部门" label-width="50px">
                <el-input v-model="val.department"></el-input>
              </el-form-item>
            </el-col>
            <el-col :span="2">
              <el-button type="danger" icon="el-icon-delete"
                         style="width: 80px" @click.prevent="deleteCareer(val)"></el-button>
            </el-col>
          </el-row>
          <div>
            <el-alert title="无职业记录" type="info" center show-icon :closable="false"
                      v-if="career.length===0" style="margin-bottom: 10px"></el-alert>
          </div>
          <div style="text-align: end; margin-right: 20px">
            <el-button type="primary" @click="addCareer">添加职业经历</el-button>
          </div>
        </div>
      </el-form>
      <el-divider content-position="left">个人介绍</el-divider>
      <div class="home">
        <tinymce ref="editor" v-model="infoForm.introduction"/>
      </div>
    </div>
    <div class="submitBox">
      <el-button type="primary" @click="updateInfoForm">提交修改</el-button>
    </div>
  </el-card>
</template>

<script>
  import tinymce from "../../tinymce";
  export default {
    name: "EditUserInfo",
    components: {
      tinymce
    },
    data() {
      return {
        isLoading: false,
        education: [],
        career: [],
        infoForm: {
          uid: 0,
          stid: '',
          name: '',
          identity: '0',
          privilege: '0',
          email: '',
          address: '',
          QQ: '',
          URL: '',
          introduction: '',
          avatarUrl: '',
          education: [{
            begin: '',
            end: '',
            university: '',
            college: '',
            level: ''
          }],
          career: [{
            begin: '',
            end: '',
            company: '',
            department: ''
          }]
        },
      };
    },
    methods: {
      beforeAvatarUpload(file) {
        const isJPG = file.type === 'image/jpg';
        const isPNG = file.type === 'image/png';
        const isJPEG = file.type === 'image/jpeg';
        const isLt2M = file.size / 1024 / 1024 < 1;
        if (!isJPG && !isPNG && !isJPEG) {
          this.$message.error('上传头像图片只能是 JPG 格式!');
        }
        if (!isLt2M) {
          this.$message.error('上传头像图片大小不能超过 1MB!');
        }
        return (isJPG || isPNG || isJPEG) && isLt2M;
      },
      async customUpload(file) {
        let formData = new FormData();
        // 服务端接收文件的参数名，文件数据，文件名
        formData.append('upfile', file.file, file.file.name);
        const {data: res} = await this.$http.post('upload_image', formData);
        if (res.code === 200) {
          this.infoForm.avatarUrl = res.data;
        } else {
          this.$message(res.msg)
        }
      },
      addEducation: function () {
        this.education.push({begin: '', end: '', university: '', college: '', level: ''})
      },
      deleteEducation: function (item) {
        const index = this.education.indexOf(item);
        if (index !== -1) {
          this.education.splice(index, 1)
        }
      },
      addCareer: function () {
        this.career.push({begin: '', end: '', company: '', department: ''})
      },
      deleteCareer: function (item) {
        const index = this.career.indexOf(item);
        if (index !== -1) {
          this.career.splice(index, 1)
        }
      },
      feachInfoForm: async function () {
        const {data: res} = await this.$http.get('user_all_info',{params: {uid: this.$route.params.uid}});
        if (res.code !== 200)
          return this.$message.error(res.msg);
        else {
          this.infoForm = res.data.userAllInfo;
          this.infoForm.identity  = this.infoForm.identity + "";
          this.infoForm.privilege  = this.infoForm.privilege + "";
          if (this.infoForm.education.length !== 0) {
            this.education = JSON.parse(this.infoForm.education);
          } else {
            this.education = []
          }
          if (this.career.length !== 0) {
            this.career = JSON.parse(this.infoForm.career);
          } else {
            this.infoForm.career = []
          }
        }
      },
      updateInfoForm: async function () {
        this.isLoading = true;
        this.infoForm.uid = Number(this.$route.params.uid);
        this.infoForm.identity  = Number(this.infoForm.identity);
        this.infoForm.privilege  = Number(this.infoForm.privilege);
        this.infoForm.education = JSON.stringify(this.education);
        this.infoForm.career = JSON.stringify(this.career);
        const { data:result } = await this.$http.post('/console/user_all_edit', this.infoForm);
        if (result.code !== 200) {
          this.$message.error(result.msg);
        } else {
          this.$message.success('修改个人资料成功');
        }
        this.isLoading = false;
      }
    },
    created() {
      this.infoForm.uid = Number(this.$route.params.uid);
      this.feachInfoForm()
    }
  }
</script>

<style lang="less">
  .avatar-uploader .el-upload {
    border: 1px dashed #d9d9d9;
    border-radius: 6px;
    cursor: pointer;
    position: relative;
    overflow: hidden;
  }
  .avatar-uploader .el-upload:hover {
    border-color: #409EFF;
  }
  .avatar-uploader-icon {
    font-size: 28px;
    color: #8c939d;
    width: 150px;
    height: 200px;
    line-height: 200px !important;
    text-align: center;
  }
  .avatar {
    width: 150px;
    height: 200px;
    display: block;
  }
</style>

<style lang="less" scoped>
  .box-card {
    text-align: center;
    position: relative;
    width: 96%;
    left: 2%;
    top: 5%;
  }
  .banner {
    width: 150px;
    .el-alert {
      text-align: center;
      padding: 0;
    }
  }
  .divText {
    font-size: 13px;
  }
  .submitBox {
    margin-top: 15px;
    text-align: end;
    margin-right: 20px;
  }
</style>
