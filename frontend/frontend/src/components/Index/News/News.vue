<template>
  <el-card class="box-card">
    <div class="clearfix" style="text-align: left">
      <span>活动与新闻列表</span>
    </div>
    <div style="text-align: center">
      <el-table :data="lists" style="width: 100%; height: 440px;" stripe>
        <el-table-column prop="title" label="标题" min-width="80%">
          <template slot-scope="scope">
            <span @click="goToDetail(scope.row.n_id)" class="newsItem">{{scope.row.title}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="date" label="日期" min-width="10%"></el-table-column>
        <el-table-column prop="promulgator" label="发布者" min-width="10%"></el-table-column>
      </el-table>
      <el-pagination @current-change="handleCurrentChange" :current-page.sync="paginationInfo.currentPage"
                     :page-size="paginationInfo.pageSize" layout="prev, pager, next, jumper" :total="paginationInfo.totalRecords">
      </el-pagination>
    </div>
  </el-card>
</template>

<script>
  export default {
    name: "Announcement",
    data() {
      return {
        lists: [],
        paginationInfo: {
          currentPage: 1,
          totalRecords: 1,
          pageSize: 8
        },
        queryNewsBriefInfo: {
          size: 0,
          offset: 0,
        },
      }
    },
    methods: {
      handleCurrentChange: function () {
        this.getNewsBriefList();
      },
      getNewsBriefList: async function () {
        this.queryNewsBriefInfo.size = this.paginationInfo.pageSize;
        this.queryNewsBriefInfo.offset = (this.paginationInfo.currentPage-1)*this.paginationInfo.pageSize;
        const {data: res} = await this.$http.get('news_list_brief',
            {params: {conf: JSON.stringify(this.queryNewsBriefInfo)}});
        if (res.code !== 200)
          return this.$message.error(res.msg);
        this.lists = res.data.records;
        this.paginationInfo.totalRecords = res.data.total;
      },
      goToDetail: function (nid) {
        const { href } = this.$router.resolve(`news/`+nid);
        window.open(href, "_blank");
      }
    },
    created() {
      this.getNewsBriefList()
    }
  }
</script>

<style scoped>
  .newsItem:hover {
    text-decoration: underline;
    cursor: pointer;
  }
</style>
