<template>
  <el-card class="box-card">
    <el-page-header @back="goToHome" content="添加公告"></el-page-header>
    <div>
      <el-input v-model="addBulletinForm.bulletinTitle" placeholder="请输入标题"></el-input>
      <el-divider></el-divider>
      <div class="myeditor">
        <div class="home">
          <tinymce ref="editor" v-model="addBulletinForm.bulletinDetail"/>
        </div>
      </div>
    </div>
    <div class="submitButton">
      <el-button type="primary" :loading="isLoading" @click="submitBulletin">提交</el-button>
    </div>
  </el-card>
</template>

<script>
  import tinymce from "../../tinymce";
  export default {
    props: ['userTipInfo'],
    name: "AddBulletin",
    components: {
      tinymce
    },
    data() {
      return {
        isLoading: false,
        addBulletinForm: {
          bulletinTitle: '',
          bulletinDetail: '',
        }
      }
    },
    methods: {
      goToHome: function () {
        this.$router.push('/console_home');
      },
      submitBulletin: async function () {
        this.isLoading = true;
        const { data:result } = await this.$http.post('/console/bulletin_add', this.addBulletinForm);
        if (result.code !== 200) {
          this.$message.error(result.msg);
        } else {
          this.$message.success('添加公告成功');
          this.addBulletinForm.bulletinTitle = '';
          this.addBulletinForm.bulletinDetail = '';
        }
        this.isLoading = false;
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
  }
  .submitButton {
    margin-top: 10px;
    text-align: end;
  }
</style>
