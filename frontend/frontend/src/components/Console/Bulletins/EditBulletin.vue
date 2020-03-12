<template>
  <el-card class="box-card">
    <el-page-header @back="goback" content="编辑公告"></el-page-header>
    <div>
      <el-input v-model="editBulletinForm.bulletinTitle" placeholder="请输入标题"></el-input>
      <el-divider></el-divider>
      <div class="myeditor">
        <div class="home">
          <tinymce ref="editor" v-model="editBulletinForm.bulletinDetail"/>
        </div>
      </div>
    </div>
    <div class="submitButton">
      <el-button type="primary" :loading="isLoading" @click="submitEditBulletin">提交修改</el-button>
    </div>
  </el-card>
</template>

<script>
  import tinymce from "../../tinymce";

  export default {
    props: ['userTipInfo'],
    name: "EditBulletin",
    components: {
      tinymce
    },
    data() {
      return {
        isLoading: false,
        editBulletinForm: {
          bid: 0,
          bulletinTitle: '',
          bulletinDetail: '',
        },
      }
    },
    methods: {
      goback: function () {
        this.$router.go(-1);
      },
      getBulletinInfo: async function() {
        const {data: res} = await this.$http.get('bulletin_info', {params: {bid: this.$route.params.bid}});
        if (res.code !== 200) {
          this.$message.error(res.msg);
          this.$router.push('/console_listBulletin')
        }
        else {
          this.editBulletinForm.bulletinTitle = res.data.bulletinTitle;
          this.editBulletinForm.bulletinDetail = res.data.bulletinDetail;
        }
      },
      submitEditBulletin: async function () {
        this.isLoading = true;
        this.editBulletinForm.bid = Number(this.$route.params.bid);
        const { data:result } = await this.$http.post('/console/bulletin_edit', this.editBulletinForm);
        if (result.code !== 200) {
          this.$message.error(result.msg);
        } else {
          this.$message.success('修改公告成功');
          this.$router.push('/console_listBulletin')
        }
        this.isLoading = false;
      }
    },
    created() {
      this.getBulletinInfo()
    }
  }
</script>

<style scoped>
  .submitButton {
    margin-top: 10px;
    text-align: end;
  }
</style>
