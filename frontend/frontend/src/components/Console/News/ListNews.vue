<template>
  <el-card class="box-card">
    <el-page-header @back="goToHome" content="新闻列表"></el-page-header>
    <el-table
            :data="lists"
            style="width: 100%">
      <el-table-column prop="date" label="发布日期" width="100"></el-table-column>
      <el-table-column prop="title" label="标题" width="650"></el-table-column>
      <el-table-column prop="promulgator" label="发布者" width="200">
      </el-table-column>
      <el-table-column prop="options" label="操作">
        <template slot-scope="scope">
          <el-tooltip effect="dark" content="查看新闻全文" placement="top" :enterable="false">
            <el-button type="info" icon="el-icon-files" size="mini" @click="viewDetail(scope.row.n_id)"></el-button>
          </el-tooltip>
          <el-tooltip effect="dark" content="编辑新闻" placement="top" :enterable="false">
            <el-button type="primary" icon="el-icon-edit" size="mini" @click="editNews(scope.row.n_id)"></el-button>
          </el-tooltip>
          <el-tooltip effect="dark" content="删除此条新闻" placement="top" :enterable="false">
            <el-popconfirm confirmButtonText='确认删除' cancelButtonText='我手滑了' icon="el-icon-info" iconColor="red"
                           title="确认删除该新闻吗？" style="margin-left:10px" @onConfirm='removeNewsByNid(scope.row.n_id)'>
              <el-button slot="reference" type="danger" icon="el-icon-delete" size="mini"></el-button>
            </el-popconfirm>
          </el-tooltip>
        </template>
      </el-table-column>
    </el-table>
    <el-pagination @current-change="handleCurrentChange" :current-page.sync="paginationInfo.currentPage"
                   :page-size="paginationInfo.pageSize" layout="prev, pager, next, jumper" :total="paginationInfo.totalRecords">
    </el-pagination>
  </el-card>
</template>

<script>
  export default {
    name: "ListNews",
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
      goToHome: function () {
        this.$router.push('/console_home');
      },
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
      removeNewsByNid: async function(nid) {
        const {data: res} = await this.$http.delete('console/news_delete', {params: {nid: nid}});
        if (res.code !== 200)
          return this.$message.error(res.msg);
        else {
          this.$message.success('删除新闻成功');
          this.getNewsBriefList();
        }
      },
      editNews: function (nid) {
        this.$router.push('/console_editNews/' + nid)
      },
      viewDetail: function (nid) {
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
  .box-card {
    text-align: center;
    position: relative;
    width: 96%;
    left: 2%;
    padding: 0;
  }
</style>
