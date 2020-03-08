<template>
  <div>
    <el-card class="box-card">
      <el-page-header @back="goToHome" content="成员列表"></el-page-header>
      <div>
        <el-row :gutter="20">
          <el-col :span="10">
            <el-input placeholder="输入 学号/工号 或 姓名 以检索" v-model="queryRecordInfo.recordRule.keyword"
                      class="input-with-select" clearable @clear="clearFilter" size="small">
              <el-button slot="append" icon="el-icon-search" @click="searchHandler"></el-button>
            </el-input>
          </el-col>
          <el-col :span="2">
            <el-button @click="clearFilter" size="small">清除筛选</el-button>
          </el-col>
        </el-row>
        <el-table ref="filterTable" :data="recordsList" style="width: 100%;" size="small" :cell-style="{padding:'5px'}"
                  @filter-change="filterHandler" @sort-change="sortHandler">
          <el-table-column prop="stid" label="学号/工号" width="200" sortable="custom"></el-table-column>
          <el-table-column prop="name" label="姓名" width="190"></el-table-column>
          <el-table-column prop="grade" label="年级" width="150" sortable="custom" column-key="grade"
                           :filters="filterList.grade"></el-table-column>
          <el-table-column prop="identity" label="身份" width="160" column-key="identity"
                           :filters="filterList.identity">
            <template slot-scope="scope">
              <el-alert title="教师/教练" type="error" effect="dark" :closable="false"
                        v-if="scope.row.identity===1"></el-alert>
              <el-alert title="正式成员" type="success" effect="dark" :closable="false"
                        v-else-if="scope.row.identity===2"></el-alert>
              <el-alert title="预备成员" type="info" effect="dark" :closable="false"
                        v-else-if="scope.row.identity===3"></el-alert>
            </template>
          </el-table-column>
          <el-table-column prop="privilege" label="权限" width="170" column-key="privilege"
                           :filters="filterList.privilege">
            <template slot-scope="scope">
              <el-alert title="root" type="error" effect="dark" :closable="false"
                        v-if="scope.row.privilege===1"></el-alert>
              <el-alert title="超级管理员" type="success" effect="dark" :closable="false"
                        v-else-if="scope.row.privilege===2"></el-alert>
              <el-alert title="团队建设管理员" type="warning" effect="dark" :closable="false"
                        v-else-if="scope.row.privilege===3"></el-alert>
              <el-alert title="竞赛训练管理员" type="warning" effect="dark" :closable="false"
                        v-else-if="scope.row.privilege===4"></el-alert>
              <el-alert title="普通权限" type="info" effect="dark" :closable="false"
                        v-else-if="scope.row.privilege===5"></el-alert>
            </template>
          </el-table-column>
          <el-table-column prop="operation" label="操作" width="300">
            <template slot-scope="scope">
              <el-tooltip effect="dark" content="查看个人主页" placement="top" :enterable="false">
                <el-button type="info" icon="el-icon-user" size="mini"></el-button>
              </el-tooltip>
              <el-tooltip effect="dark" content="编辑基本信息" placement="top" :enterable="false">
                <el-button type="warning" icon="el-icon-setting" size="mini" @click="editUserInfo(scope.row)"></el-button>
              </el-tooltip>
              <el-tooltip effect="dark" content="编辑详细信息" placement="top" :enterable="false">
                <el-button type="primary" icon="el-icon-edit" size="mini"></el-button>
              </el-tooltip>
              <el-tooltip effect="dark" content="重置密码" placement="top" :enterable="false">
                <el-popconfirm confirmButtonText='确认重置' cancelButtonText='我手滑了' icon="el-icon-info" iconColor="red"
                               title="确认将密码重置为学号吗？" style="margin-left:10px" @onConfirm="resetUserPSW(scope.row.uid)">
                  <el-button slot="reference" type="warning" icon="el-icon-unlock" size="mini"
                             :disabled="userTipInfo.stid===scope.row.stid || userTipInfo.privilege>=scope.row.privilege"></el-button>
                </el-popconfirm>
              </el-tooltip>
              <el-tooltip effect="dark" content="删除成员" placement="top" :enterable="false">
                <el-popconfirm confirmButtonText='确认删除' cancelButtonText='我手滑了' icon="el-icon-info" iconColor="red"
                               title="确认删除该成员吗？" style="margin-left:10px" @onConfirm="removeUserByUid(scope.row.uid)">
                  <el-button slot="reference" type="danger" icon="el-icon-delete" size="mini"
                             :disabled="userTipInfo.stid===scope.row.stid || userTipInfo.privilege>=scope.row.privilege"></el-button>
                </el-popconfirm>
              </el-tooltip>
            </template>
          </el-table-column>

        </el-table>
        <el-pagination @current-change="handleCurrentChange" :current-page.sync="paginationInfo.currentPage"
                       :page-size="paginationInfo.pageSize" layout="prev, pager, next, jumper" :total="paginationInfo.totalRecords">
        </el-pagination>
      </div>
    </el-card>
    <el-dialog title="编辑基本信息" :visible.sync="editForm.formVisible">
      <el-form :model="editForm.UserInfo" :disabled="editForm.formLoading" :label-width="editForm.labelWidth"
               :rules="editFormRules" ref="editFormRefUser">
        <el-form-item label="学号" prop="stid">
          <el-input v-model="editForm.UserInfo.stid" disabled></el-input>
        </el-form-item>
        <el-form-item label="姓名" prop="name">
          <el-input v-model="editForm.UserInfo.name" :disabled="editForm.formLoading"></el-input>
        </el-form-item>
        <el-form-item label="身份" prop="identity">
          <el-radio-group v-model="editForm.UserInfo.identity">
            <el-radio label="1">教师/教练</el-radio>
            <el-radio label="2">正式成员</el-radio>
            <el-radio label="3">预备成员</el-radio>
          </el-radio-group>
        </el-form-item>
        <el-form-item label="年级" prop="grade" v-if="editForm.UserInfo.identity==='2'||editForm.UserInfo.identity==='3'">
          <el-input v-model="editForm.UserInfo.grade" :disabled="editForm.formLoading"></el-input>
        </el-form-item>
        <el-form-item label="权限" prop="privilege">
          <el-select v-model="editForm.UserInfo.privilege" placeholder="请选择"
                     :disabled="(userTipInfo.privilege>=editForm.UserInfo.privilege && userTipInfo.privilege!==1)||userTipInfo.uid===editForm.UserInfo.uid">
            <el-option key="p1" label="root" value="1" v-if="userTipInfo.privilege===1"></el-option>
            <el-option key="p2" label="超级管理员" value="2" v-if="userTipInfo.privilege===1"></el-option>
            <el-option key="p3" label="团队建设管理员" value="3" v-if="userTipInfo.privilege<=2"></el-option>
            <el-option key="p4" label="竞赛训练管理员" value="4" v-if="userTipInfo.privilege<=2"></el-option>
            <el-option key="p5" label="普通权限" value="5" v-if="userTipInfo.privilege<=3"></el-option>
          </el-select>
        </el-form-item>
      </el-form>
      <div slot="footer" class="dialog-footer">
        <el-button @click="cancelEditForm">取 消</el-button>
        <el-button type="primary" @click="comfirmEditForm">确 定</el-button>
      </div>
    </el-dialog>
  </div>
</template>

<script>
  export default {
    props: ['userTipInfo'],
    name: "ListUser",
    data() {
      const checkGrade = (rule, value, cb) => {
        if (value || this.editForm.UserInfo.identity==='1') {
          return cb()
        }
        cb(new Error('请输入年级'))
      };
      return {
        paginationInfo: {
          currentPage: 1,
          totalRecords: 1,
          pageSize: 8
        },
        queryRecordInfo: {
          sortRule: {
            column: '',
            order: ''
          },
          filterRule: {
            identity: [],
            privilege: [],
            grade: []
          },
          recordRule: {
            size: 0,
            offset: 0,
            keyword: '' // 搜索用
          }
        },
        recordsList: [],
        filterList: {
          grade: [],
          identity: [
            {text: '教师/教练', value: 1},
            {text: '正式成员', value: 2},
            {text: '预备成员', value: 3},
          ],
          privilege: [
            {text: 'root', value: 1},
            {text: '超级管理员', value: 2},
            {text: '团队建设管理员', value: 3},
            {text: '竞赛训练管理员', value: 4},
            {text: '普通权限', value: 5},
          ]
        },
        editForm: {
          formVisible: false,
          formLoading: false,
          UserInfo: {
            uid: '',
            stid: '',
            name: '',
            grade: '',
            identity: "0",
            privilege: "0"
          },
          labelWidth: "70px"
        },
        editFormRules: {
          stid: [
            {required: true, message: '请输入学号/工号', trigger: 'blur'},
          ],
          name: [
            {required: true, message: '请输入姓名', trigger: 'blur'},
          ],
          identity: [
            {required: true, message: '请选择身份', trigger: 'blur'},
          ],
          grade: [
            {validator: checkGrade, trigger: 'blur'}
          ]
        },
      }
    },
    methods: {
      goToHome: function () {
        this.$router.push('/console_home');
      },
      clearFilter: function () {
        this.queryRecordInfo.sortRule.column = this.queryRecordInfo.sortRule.order = '';
        this.queryRecordInfo.filterRule.identity = this.queryRecordInfo.filterRule.privilege = this.queryRecordInfo.filterRule.grade = [];
        this.queryRecordInfo.recordRule.size = this.queryRecordInfo.recordRule.offset = 0;
        this.queryRecordInfo.recordRule.keyword = '';
        this.getRecordList()
      },
      filterHandler(obj) {
        const itemName = Object.keys(obj)[0];
        this.queryRecordInfo["filterRule"][itemName] = obj[itemName];
        this.getRecordList()
      },
      sortHandler(obj) {
        if (obj.order !== null) {
          if (obj.order === "descending")
            obj.order = "desc";
          else
            obj.order = "asc";
          this.queryRecordInfo["sortRule"]["column"] = obj.prop;
          this.queryRecordInfo["sortRule"]["order"] = obj.order;
        } else {
          this.queryRecordInfo["sortRule"]["column"] = '';
          this.queryRecordInfo["sortRule"]["order"] = '';
        }
        this.getRecordList()
      },
      searchHandler: function() {
        const tmp = this.queryRecordInfo.recordRule.keyword;
        this.clearFilter();
        this.queryRecordInfo.recordRule.keyword = tmp;
        this.getRecordList()
      },
      handleCurrentChange: function (val) {
        this.getRecordList()
      },
      getRecordList: async function () {
        this.queryRecordInfo.recordRule.offset = (this.paginationInfo.currentPage - 1)*this.paginationInfo.pageSize;
        this.queryRecordInfo.recordRule.size = this.paginationInfo.pageSize;
        const {data: res} = await this.$http.get('user_info_list',
            {params: {conf: JSON.stringify(this.queryRecordInfo)}});
        if (res.code !== 200)
          return this.$message.error(res.msg);
        this.recordsList = res.data.records;
        this.paginationInfo.totalRecords = res.data.total;
      },
      getGradeList: async function () {
        const {data: res} = await this.$http.get('user_grade_list');
        if (res.code !== 200)
          return this.$message.error(res.msg);
        this.filterList.grade = res.data
      },
      removeUserByUid: async function (uid) {
        const {data: res} = await this.$http.delete('console/user_delete', {params: {uid: uid}});
        if (res.code !== 200)
          return this.$message.error(res.msg);
        else {
          this.$message.success('删除成员成功');
          this.getRecordList();
        }
      },
      editUserInfo: function(row) {
        this.editForm.UserInfo.stid = row.stid;
        this.editForm.UserInfo.uid = row.uid;
        this.editForm.UserInfo.name = row.name;
        this.editForm.UserInfo.grade = row.grade;
        this.editForm.UserInfo.privilege = row.privilege+"";
        this.editForm.UserInfo.identity = row.identity+"";
        this.editForm.formVisible = true;
      },
      cancelEditForm: function () {
        this.editForm.formVisible = false;
        this.editForm.uid = '';
        this.editForm.name = '';
        this.editForm.grade = '';
        this.editForm.identity = 0;
        this.editForm.privilege = 0;
      },
      comfirmEditForm: function () {
        this.$refs.editFormRefUser.validate(async (valid) => {
          if (!valid) {
            return;
          }
          this.editForm.formLoading = true;
          this.editForm.UserInfo.identity = Number(this.editForm.UserInfo.identity);
          this.editForm.UserInfo.privilege = Number(this.editForm.UserInfo.privilege);
          const {data: res} = await this.$http.post('console/user_edit', this.editForm.UserInfo);
          if (res.code !== 200)
            this.$message.error(res.msg);
          else {
            this.$message.success('更新基本信息成功');
            this.getRecordList();
          }
          this.editForm.formLoading = false;
        });
        this.cancelEditForm()
      },
      resetUserPSW: async function(uid){
        const {data: res} = await this.$http.post('console/user_reset_psw', {"uid": uid});
        if (res.code !== 200)
          return this.$message.error(res.msg);
        else {
          this.$message.success('成功将密码重置为学号');
        }
      }
    },
    created() {
      this.getGradeList();
      this.getRecordList();
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
