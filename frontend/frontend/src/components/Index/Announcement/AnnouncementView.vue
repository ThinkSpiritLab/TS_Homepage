<template>
  <el-card class="box-card">
    <div class="clearfix" style="text-align: left">
      <span>公告详情</span>
    </div>
    <div class="bulletinHead">
      <div style="font-size: 32px">{{detail.title}}</div>
      <el-row style="font-size: 12px; color: #999999">
        <el-col :span="8" :offset="4">
          发布日期：{{detail.date}}
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
          bid: 0,
          title: '',
          detail: '',
          date: '',
          promulgator: ''
        }
      }
    },
    methods: {
      getBulletinDetail: async function () {
        const {data: res} = await this.$http.get('bulletin_detail', {params: {bid: this.$route.params.bid}});
        if (res.code !== 200)
          return this.$message.error(res.msg);
        else {
          this.detail = res.data;
        }
      }
    },
    created() {
      this.detail.bid = Number(this.$route.params.bid);
      this.getBulletinDetail()
    }
  }
</script>

<style scoped>
  .bulletinHead {
    text-align: center;
    font-family: "Helvetica Neue",Helvetica,"PingFang SC","Hiragino Sans GB","Microsoft YaHei","微软雅黑",Arial,sans-serif;
  }
</style>
