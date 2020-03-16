<template>
  <el-card class="box-card">
    <div class="clearfix" style="text-align: left">
      <span>新闻与活动详情</span>
    </div>
    <div class="bulletinHead">
      <div style="font-size: 32px">{{detail.title}}</div>
      <el-row style="font-size: 12px; color: #999999">
        <el-col :span="8" :offset="4">
          日期：{{detail.date}}
        </el-col>
        <el-col :span="8">
          发布者：{{detail.promulgator}}
        </el-col>
      </el-row>
      <el-divider><i class="el-icon-help"></i></el-divider>
    </div>
    <div v-html="detail.detail"></div>
  </el-card>
</template>

<script>
  export default {
    name: "AnnouncementView",
    data() {
      return {
        detail: {
          nid: 0,
          title: '',
          detail: '',
          date: '',
          promulgator: ''
        }
      }
    },
    methods: {
      getNewsDetail: async function () {
        const {data: res} = await this.$http.get('news_detail', {params: {nid: this.$route.params.nid}});
        if (res.code !== 200)
          return this.$message.error(res.msg);
        else {
          this.detail = res.data;
        }
      }
    },
    created() {
      this.detail.nid = Number(this.$route.params.nid);
      this.getNewsDetail()
    }
  }
</script>

<style scoped>
  .bulletinHead {
    text-align: center;
    font-family: "Helvetica Neue",Helvetica,"PingFang SC","Hiragino Sans GB","Microsoft YaHei","微软雅黑",Arial,sans-serif;
  }
</style>
