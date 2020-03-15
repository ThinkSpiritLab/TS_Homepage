<template>
  <el-card class="box-card">
    <div class="clearfix" style="text-align: left">
      <span>公告列表</span>
    </div>
    <div style="text-align: center">
      <el-table :data="lists" style="width: 100%; height: 440px;" stripe>
        <el-table-column prop="title" label="公告标题" min-width="80%">
          <template slot-scope="scope">
            <span @click="goToDetail(scope.row.b_id)">{{scope.row.title}}</span>
          </template>
        </el-table-column>
        <el-table-column prop="date" label="发布日期" min-width="10%"></el-table-column>
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
        queryBulletinBriefInfo: {
          size: 0,
          offset: 0,
        },
      }
    },
    methods: {
      handleCurrentChange: function () {
        this.getBulletinBriefList();
      },
      getBulletinBriefList: async function () {
        this.queryBulletinBriefInfo.size = this.paginationInfo.pageSize;
        this.queryBulletinBriefInfo.offset = (this.paginationInfo.currentPage-1)*this.paginationInfo.pageSize;
        const {data: res} = await this.$http.get('bulletin_list_brief',
            {params: {conf: JSON.stringify(this.queryBulletinBriefInfo)}});
        if (res.code !== 200)
          return this.$message.error(res.msg);
        this.lists = res.data.records;
        this.paginationInfo.totalRecords = res.data.total;
      },
      goToDetail: function (bid) {
        console.log(bid)
      }
    },
    created() {
      this.getBulletinBriefList()
    }
  }
</script>

<style scoped>

</style>
