<template>
  <div>
    <el-card class="box-card">
      <el-page-header @back="goToHome" content="竞赛列表"></el-page-header>
      <el-table
              :data="lists"
              style="width: 100%">
        <el-table-column prop="c_time" label="日期" width="100"></el-table-column>
        <el-table-column prop="c_name_zh" label="名称" width="500"></el-table-column>
        <el-table-column prop="achievements" label="成绩" width="350">
          <template slot-scope="scope">
            <div>
              <span :class="scope.row['r_type']==='1'?'hiddenspan':''">
                <svg class="icon" aria-hidden="true"><use xlink:href="#icon-tedengjiang"></use></svg>
                <span style="margin-left: 5px">{{"X " + scope.row["num_0"]}}</span>
              </span>
              <span style="margin-left: 20px">
                <svg class="icon" aria-hidden="true"><use xlink:href="#icon-jinpai"></use></svg>
                <span style="margin-left: 5px">{{"X " + scope.row["num_1"]}}</span>
              </span>
              <span style="margin-left: 20px">
                <svg class="icon" aria-hidden="true"><use xlink:href="#icon-yinpai"></use></svg>
                <span style="margin-left: 5px">{{"X " + scope.row["num_2"]}}</span>
              </span>
              <span style="margin-left: 20px">
                <svg class="icon" aria-hidden="true"><use xlink:href="#icon-tongpai"></use></svg>
                <span style="margin-left: 5px">{{"X " + scope.row["num_3"]}}</span>
              </span>
              <span style="margin-left: 20px">
                <svg class="icon" aria-hidden="true"><use xlink:href="#icon-tiepai"></use></svg>
                <span style="margin-left: 5px">{{"X " + scope.row["num_4"]}}</span>
              </span>
            </div>
          </template>
        </el-table-column>
        <el-table-column prop="options" label="操作">
          <template slot-scope="scope">
            <el-tooltip effect="dark" content="查看详细信息" placement="top" :enterable="false">
              <el-button type="info" icon="el-icon-files" size="mini"></el-button>
            </el-tooltip>
            <el-tooltip effect="dark" content="编辑额外信息" placement="top" :enterable="false">
              <el-button type="primary" icon="el-icon-edit" size="mini" @click='editContestExtras(scope.row.c_id)'></el-button>
            </el-tooltip>
            <el-tooltip effect="dark" content="删除此条记录" placement="top" :enterable="false">
              <el-popconfirm confirmButtonText='确认删除' cancelButtonText='我手滑了' icon="el-icon-info" iconColor="red"
                             title="确认删除该记录吗？" style="margin-left:10px" @onConfirm='removeContestByCid(scope.row.c_id)'>
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
    <el-dialog title="编辑额外信息" :visible.sync="editVisible" :modal="false" width="80%">
      <div class="myeditor">
        <div class="home">
          <tinymce ref="editor" v-model="editForm.extras"/>
        </div>
      </div>
      <div class="submitButton">
        <el-button @click="cancelEditExtras">取 消</el-button>
        <el-button type="primary" @click="comfirmEditExtras">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
  import tinymce from "../../tinymce";
  import icon from "../../../assets/font_awards/iconfont.js"
  export default {
    name: "ListContest",
    components: {
      tinymce
    },
    data() {
      return {
        lists: [],
        paginationInfo: {
          currentPage: 1,
          totalRecords: 1,
          pageSize: 8
        },
        queryContestBriefInfo: {
          size: 0,
          offset: 0,
        },
        editForm: {
          cid: 0,
          extras: ''
        },
        editVisible: false
      }
    },
    methods: {
      goToHome: function () {
        this.$router.push('/console_home');
      },
      handleCurrentChange: function () {
        this.getContestBriefList();
      },
      getContestBriefList: async function () {
        this.queryContestBriefInfo.size = this.paginationInfo.pageSize;
        this.queryContestBriefInfo.offset = (this.paginationInfo.currentPage-1)*this.paginationInfo.pageSize;
        const {data: res} = await this.$http.get('contest_list_brief',
            {params: {conf: JSON.stringify(this.queryContestBriefInfo)}});
        if (res.code !== 200)
          return this.$message.error(res.msg);
        this.lists = res.data.records;
        this.paginationInfo.totalRecords = res.data.total;
      },
      removeContestByCid: async function(cid) {
        const {data: res} = await this.$http.delete('contest/contest_delete', {params: {cid: cid}});
        if (res.code !== 200)
          return this.$message.error(res.msg);
        else {
          this.$message.success('删除竞赛记录成功');
          this.getContestBriefList();
        }
      },
      editContestExtras: async function(cid) {
        this.editVisible = true;
        const {data: res} = await this.$http.get('contest_extras', {params: {cid: cid}});
        if (res.code !== 200)
          return this.$message.error(res.msg);
        else {
          this.editForm.extras = res.data;
          this.editForm.cid = cid;
        }
      },
      cancelEditExtras: function() {
        this.editForm.extras = "";
        this.editForm.cid = 0;
        this.editVisible = false;
      },
      comfirmEditExtras: async function() {
        const {data: res} = await this.$http.post('contest/contest_extras_edit', this.editForm);
        if (res.code !== 200)
          return this.$message.error(res.msg);
        else {
          this.$message.success("修改成功");
          return this.editVisible = false;
        }
      }
    },
    created() {
      this.getContestBriefList()
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
  .icon {
    width: 1em;
    height: 1em;
    vertical-align: -0.15em;
    fill: currentColor;
    overflow: hidden;
  }
  .submitButton {
    margin-top: 10px;
    text-align: end;
  }
</style>
