<template>
  <el-card class="cardBox">
    <template>
      <el-tabs v-model="activeName" :stretch="true"  @tab-click="handleClick">
        <el-tab-pane label="教师与教练" name="faculty">
          <el-row :gutter="30">
            <el-col :span="8" v-for="(val, index) in infoList" :key="index">
              <el-card shadow="always" body-style="padding-top: 0">
                <el-row :gutter="40">
                  <el-col :span="12">
                    <img :src="val.avatarUrl" class="memberAvatar" v-if="val.avatarUrl !== ''">
                    <img src="../../../assets/defaultAvatar.png" class="memberAvatar" v-else>
                  </el-col>
                  <el-col :span="12">
                    <div style="padding: 14px;position: absolute;top: 35%; text-align: center">
                      <div style="margin-bottom: 20px">{{val.name}}</div>
                    </div>
                  </el-col>
                </el-row>
              </el-card>
            </el-col>
          </el-row>
        </el-tab-pane>
        <el-tab-pane label="在校成员" name="students">
          <el-row :gutter="30">
            <el-col :span="6" v-for="(val, index) in infoList" :key="index">
              <el-card shadow="always" @click.native="goToMemberPage(val.uid)" class="userCard">
                <el-row :gutter="40">
                  <el-col :span="12">
                    <img :src="val.avatarUrl" class="memberAvatar" v-if="val.avatarUrl !== ''">
                    <img src="../../../assets/defaultAvatar.png" class="memberAvatar" v-else>
                  </el-col>
                  <el-col :span="12">
                    <div style="padding: 14px;position: absolute;top: 30%; text-align: center; margin-left: 20px">
                      <div style="margin-bottom: 20px">{{val.name}}</div>
                      <div style="white-space:normal; word-break:break-all;overflow:hidden;">{{val.grade + " 级"}}</div>
                    </div>
                  </el-col>
                </el-row>
              </el-card>
            </el-col>
          </el-row>
        </el-tab-pane>
        <el-tab-pane label="毕业成员" name="alumni">
          <el-row :gutter="30">
            <el-col :span="6" v-for="(val, index) in infoList" :key="index">
              <el-card shadow="always"  @click.native="goToMemberPage(val.uid)" class="userCard">
                <el-row :gutter="40">
                  <el-col :span="12">
                    <img :src="val.avatarUrl" class="memberAvatar" v-if="val.avatarUrl !== ''">
                    <img src="../../../assets/defaultAvatar.png" class="memberAvatar" v-else>
                  </el-col>
                  <el-col :span="12">
                    <div style="padding: 14px;position: absolute;top: 30%; text-align: center;  margin-left: 20px">
                      <div style="margin-bottom: 20px">{{val.name}}</div>
                      <div style="white-space:normal; word-break:break-all;overflow:hidden;">{{val.grade + " 级"}}</div>
                    </div>
                  </el-col>
                </el-row>
              </el-card>
            </el-col>
          </el-row>
        </el-tab-pane>
      </el-tabs>
    </template>
  </el-card>

</template>

<script>
  export default {
    name: "Members",
    data() {
      return {
        activeName: 'faculty',
        infoList: [{
          uid: '',
          name: '',
          stid: '',
          grade: '',
          email: '',
          identity: '',
          avatarUrl: ''
        }],
      }
    },
    methods: {
      handleClick: async function(tab) {
        if (tab.name === "faculty") {
          this.getFacultyInfo()
        } else if (tab.name === "students") {
          this.getStudentsInfo()
        } else if (tab.name === "alumni") {
          this.getAlumniInfo()
        } else {
          this.infoList = []
        }
      },
      getFacultyInfo: async function () {
        const {data: res} = await this.$http.get('user_members_faculty_brief');
        if (res.code !== 200)
          return this.$message.error(res.msg);
        else {
          this.infoList = res.data;
        }
      },
      getStudentsInfo: async function () {
        const {data: res} = await this.$http.get('user_members_students_brief');
        if (res.code !== 200)
          return this.$message.error(res.msg);
        else {
          this.infoList = res.data;
        }
      },
      getAlumniInfo: async function () {
        const {data: res} = await this.$http.get('user_members_alumni_brief');
        if (res.code !== 200)
          return this.$message.error(res.msg);
        else {
          this.infoList = res.data;
        }
      },
      goToMemberPage: function (uid) {
        const { href } = this.$router.resolve(`member/`+uid);
        window.open(href, "_blank");
      }
    },
    created() {
      this.getFacultyInfo()
    }
  }
</script>

<style lang="less" scoped>
  .cardBox {
    height: auto;
  }
  .memberAvatar {
    width: 150px;
    height: 200px;
  }
  .userCard {
    cursor: pointer;
    margin: 10px 10px 20px 20px
  }
  .userCard:hover {
    border-style:solid;
    border-color: #03a9f4;
    box-shadow: 0 0 15px #03a9f4;
  }
</style>
