<template>
  <el-card class="box-card">
    <el-page-header @back="goToHome" content="竞赛列表"></el-page-header>
    <el-table
            :data="lists"
            style="width: 100%">
      <el-table-column prop="c_time" label="日期" width="100"></el-table-column>
      <el-table-column prop="c_name_zh" label="名称" width="500"></el-table-column>
      <el-table-column prop="achievements" label="成绩" width="350">
        <template slot-scope="scope">
          <el-alert title="教师/教练" type="error" effect="dark" :closable="false"
                    v-if="scope.row.identity===1"></el-alert>
          <el-alert title="正式成员" type="success" effect="dark" :closable="false"
                    v-else-if="scope.row.identity===2"></el-alert>
          <el-alert title="预备成员" type="info" effect="dark" :closable="false"
                    v-else-if="scope.row.identity===3"></el-alert>
        </template>
      </el-table-column>
      <el-table-column prop="options" label="操作">
        <template slot-scope="scope">
          <el-tooltip effect="dark" content="查看详细信息" placement="top" :enterable="false">
            <el-button type="info" icon="el-icon-files" size="mini"></el-button>
          </el-tooltip>
          <el-tooltip effect="dark" content="编辑详细信息" placement="top" :enterable="false">
            <el-button type="primary" icon="el-icon-edit" size="mini"></el-button>
          </el-tooltip>
          <el-tooltip effect="dark" content="删除此条记录" placement="top" :enterable="false">
            <el-popconfirm confirmButtonText='确认删除' cancelButtonText='我手滑了' icon="el-icon-info" iconColor="red"
                           title="确认删除该成员吗？" style="margin-left:10px" @onConfirm="removeContestByCid(scope.row.uid)">
              <el-button slot="reference" type="danger" icon="el-icon-delete" size="mini"></el-button>
            </el-popconfirm>
          </el-tooltip>
        </template>
      </el-table-column>
    </el-table>
  </el-card>
</template>

<script>
  export default {
    name: "ListContest",
    data() {
      return {
        lists: [{
          c_time: '2016-05-02',
          c_name_zh: '第43届 ICPC 国际大学生程序设计竞赛亚洲区域赛（南京站）',
          achievements: '上海市普陀区金沙江路 1518 弄'
        }]
      }
    },
    methods: {
      goToHome: function () {
        this.$router.push('/console_home');
      },
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
