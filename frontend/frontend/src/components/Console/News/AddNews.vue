<template>
  <el-card class="box-card">
    <el-page-header @back="goToHome" content="添加活动与新闻"></el-page-header>
    <div>
      <el-input v-model="addNewsForm.newsTitle" placeholder="请输入标题"></el-input>
      <el-divider></el-divider>
      <div class="myeditor">
        <div class="home">
          <tinymce ref="editor" v-model="addNewsForm.newsDetail"/>
        </div>
      </div>
    </div>
    <div class="bottomArea">
      <el-row>
        <el-col :span="21">
          <span class=".web-font3" style="margin-right: 15px">活动与新闻发生时间</span>
          <el-date-picker v-model="addNewsForm.newsDate" type="date" placeholder="选择日期"
                          value-format="yyyy-MM-dd"></el-date-picker>
        </el-col>
        <el-col :span="3">
          <el-button type="primary" :loading="isLoading" @click="submitNews">提交</el-button>
        </el-col>
      </el-row>
    </div>
  </el-card>
</template>

<script>
  import tinymce from "../../tinymce";
  export default {
    props: ['userTipInfo'],
    name: "AddNews",
    components: {
      tinymce
    },
    data() {
      return {
        isLoading: false,
        addNewsForm: {
          newsTitle: '',
          newsDetail: '',
          newsDate: ''
        }
      }
    },
    methods: {
      goToHome: function () {
        this.$router.push('/console_home');
      },
      submitNews: async function () {
        this.isLoading = true;
        const { data:result } = await this.$http.post('/console/news_add', this.addNewsForm);
        if (result.code !== 200) {
          this.$message.error(result.msg);
        } else {
          this.$message.success('添加活动与新闻成功');
          this.addNewsForm.newsTitle = '';
          this.addNewsForm.newsDetail = '';
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
  .bottomArea {
    margin-top: 10px;
    text-align: left;
  }
  @font-face {
    font-family: 'webfont3';
    font-display: swap;
    src: url('//at.alicdn.com/t/webfont_vvdd6zsd78.eot'); /* IE9*/
    src: url('//at.alicdn.com/t/webfont_vvdd6zsd78.eot?#iefix') format('embedded-opentype'), /* IE6-IE8 */
    url('//at.alicdn.com/t/webfont_vvdd6zsd78.woff2') format('woff2'),
    url('//at.alicdn.com/t/webfont_vvdd6zsd78.woff') format('woff'), /* chrome、firefox */
    url('//at.alicdn.com/t/webfont_vvdd6zsd78.ttf') format('truetype'), /* chrome、firefox、opera、Safari, Android, iOS 4.2+*/
    url('//at.alicdn.com/t/webfont_vvdd6zsd78.svg#Alibaba-PuHuiTi-Bold') format('svg'); /* iOS 4.1- */
  }
  .web-font3{
    font-family:"webfont3" !important;
    font-size:16px;font-style:normal;
    -webkit-font-smoothing: antialiased;
    -webkit-text-stroke-width: 0.2px;
    -moz-osx-font-smoothing: grayscale;
  }
</style>
