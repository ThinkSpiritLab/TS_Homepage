<template>
  <el-card class="box-card">
    <el-page-header @back="goback" content="编辑新闻"></el-page-header>
    <div>
      <el-input v-model="editNewsForm.newsTitle" placeholder="请输入标题"></el-input>
      <el-divider></el-divider>
      <div class="myeditor">
        <div class="home">
          <tinymce ref="editor" v-model="editNewsForm.newsDetail"/>
        </div>
      </div>
    </div>
    <div class="bottomArea">
      <el-row>
        <el-col :span="21">
          <span class=".web-font3" style="margin-right: 15px">活动与新闻发生时间</span>
          <el-date-picker v-model="editNewsForm.newsDate" type="date" placeholder="选择日期"
                          value-format="yyyy-MM-dd"></el-date-picker>
        </el-col>
        <el-col :span="3">
          <el-button type="primary" :loading="isLoading" @click="submitEditNews">提交修改</el-button>
        </el-col>
      </el-row>
    </div>
  </el-card>
</template>

<script>
  import tinymce from "../../tinymce";

  export default {
    props: ['userTipInfo'],
    name: "EditNews",
    components: {
      tinymce
    },
    data() {
      return {
        isLoading: false,
        editNewsForm: {
          nid: 0,
          newsTitle: '',
          newsDetail: '',
          newsDate: ''
        },
      }
    },
    methods: {
      goback: function () {
        this.$router.go(-1);
      },
      getNewsInfo: async function() {
        const {data: res} = await this.$http.get('news_info', {params: {nid: this.$route.params.nid}});
        if (res.code !== 200) {
          this.$message.error(res.msg);
          this.$router.push('/console_listNews')
        }
        else {
          this.editNewsForm.newsTitle = res.data.newsTitle;
          this.editNewsForm.newsDetail = res.data.newsDetail;
          this.editNewsForm.newsDate = res.data.newsDate;
        }
      },
      submitEditNews: async function () {
        this.isLoading = true;
        this.editNewsForm.nid = Number(this.$route.params.nid);
        const { data:result } = await this.$http.post('/console/news_edit', this.editNewsForm);
        if (result.code !== 200) {
          this.$message.error(result.msg);
        } else {
          this.$message.success('修改新闻成功');
          this.$router.push('/console_listNews')
        }
        this.isLoading = false;
      }
    },
    created() {
      this.getNewsInfo()
    }
  }
</script>

<style scoped>
  .bottomArea {
    margin-top: 10px;
    text-align: left;
  }
</style>
