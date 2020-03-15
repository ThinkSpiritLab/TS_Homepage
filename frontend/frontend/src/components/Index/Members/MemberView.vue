<template>
  <el-card class="box-card">
    <div slot="header" class="clearfix" style="text-align: left">
      <span>成员信息</span>
    </div>
    <el-form ref="infoForm" :model="infoForm" label-width="80px">
      <el-row>
        <el-col :span="8" style="text-align: center">
          <div class="block">
            <el-image :src="infoForm.avatarUrl" class="userAvatar" v-if="infoForm.avatarUrl !== ''">
              <div slot="placeholder" class="image-slot">
                加载中<span class="dot">...</span>
              </div>
            </el-image>
            <img src="../../../assets/defaultAvatar.png" class="userAvatar" v-else>
          </div>
        </el-col>
        <el-col :span="16">
          <el-row>
            <el-col :span="12" style="text-align: left">
              <el-form-item label="学号/工号">
                <div>{{infoForm.stid}}</div>
              </el-form-item>
              <el-form-item label="姓名">
                <div>{{infoForm.name}}</div>
              </el-form-item>
              <el-form-item label="身份">
                <div class="banner">
                  <el-alert title="教师/教练" type="error" effect="dark" :closable="false" center
                            v-if="infoForm.identity==='1'"></el-alert>
                  <el-alert title="正式成员" type="success" effect="dark" :closable="false" center
                            v-else-if="infoForm.identity==='2'"></el-alert>
                  <el-alert title="预备成员" type="info" effect="dark" :closable="false" center
                            v-else-if="infoForm.identity==='3'"></el-alert>
                </div>
              </el-form-item>
            </el-col>
            <el-col :span="12" style="text-align: left">
              <el-form-item label="Email">
                {{infoForm.email}}
              </el-form-item>
              <el-form-item label="QQ">
                {{infoForm.QQ}}
              </el-form-item>
              <el-form-item label="URL">
                {{infoForm.URL}}
              </el-form-item>
            </el-col>
          </el-row>
          <el-row style="text-align: left">
            <el-form-item label="地址">
              {{infoForm.address}}
            </el-form-item>
          </el-row>
        </el-col>
      </el-row>
      <el-divider content-position="left">硕士及以上教育经历</el-divider>
      <div>
        <el-row v-for="(val, index) in infoForm.education" :key='index' class="experienceTip">
          <el-col :span="5" :offset="3">
            <el-form-item label="时间" label-width="50px">
              {{val.begin}} ~ {{val.end}}
            </el-form-item>
          </el-col>
          <el-col :span="5">
            <el-form-item label="学校" label-width="50px">
              {{val.university}}
            </el-form-item>
          </el-col>
          <el-col :span="5" label-width="50px">
            <el-form-item label="学院" label-width="50px">
              {{val.college}}
            </el-form-item>
          </el-col>
          <el-col :span="5" label-width="50px">
            <el-form-item label="学历" label-width="50px">
              <div v-if="val.level==='1'">硕士研究生</div>
              <div v-else-if="val.level==='2'">博士研究生</div>
              <div v-else-if="val.level==='3'">博士以上</div>
            </el-form-item>
          </el-col>
        </el-row>
        <div>
          <el-alert title="无硕士及以上教育记录" type="info" center show-icon :closable="false"
                    v-if="!infoForm.education||infoForm.education.length===0" style="margin-bottom: 10px"></el-alert>
        </div>
      </div>
      <el-divider content-position="left">职业经历</el-divider>
      <div>
        <el-row v-for="(val, index) in infoForm.career" :key='index' class="experienceTip">
          <el-col :span="5" :offset="3">
            <el-form-item label="时间" label-width="50px">
              {{val.begin}} ~ {{val.end}}
            </el-form-item>
          </el-col>
          <el-col :span="5">
            <el-form-item label="公司" label-width="50px">
              {{val.company}}
            </el-form-item>
          </el-col>
          <el-col :span="5" label-width="50px">
            <el-form-item label="部门" label-width="50px">
              {{val.department}}
            </el-form-item>
          </el-col>
        </el-row>
        <div>
          <el-alert title="无职业记录" type="info" center show-icon :closable="false"
                    v-if="!infoForm.career||infoForm.career.length===0" style="margin-bottom: 10px"></el-alert>
        </div>
      </div>
      <el-divider content-position="left">CPC 经历</el-divider>
      <div>
        <div v-for="(val, index) in contestForm" :key='index'>
        <el-row :class="{'experienceTip': true, 'NotFirstLine': index!==0}">
          <el-col :span="9" :offset="0">
            <el-form-item label="竞赛名" label-width="60px">
              {{val.c_name_zh}}
            </el-form-item>
          </el-col>
          <el-col :span="10" :offset="0">
            <el-form-item label="Contest Name" label-width="110px">
              {{val.c_name_en}}
            </el-form-item>
          </el-col>
          <el-col :span="3" :offset="0">
            <el-form-item label="竞赛时间" label-width="70px">
              {{val.c_time}}
            </el-form-item>
          </el-col>
          <el-col :span="2" :offset="0">
            <el-form-item label="详情" label-width="40px">
              <el-button type="info" icon="el-icon-files" size="mini" @click="goToContestDetail(val.c_id)"></el-button>
            </el-form-item>
          </el-col>
        </el-row>
        <el-row class="experienceTip">
          <el-col :span="9" :offset="0">
            <el-form-item label="队伍名" label-width="60px">
              {{val.t_name_zh}}
            </el-form-item>
          </el-col>
          <el-col :span="10" :offset="0">
            <el-form-item label="Team Name" label-width="110px">
              {{val.t_name_en}}
            </el-form-item>
          </el-col>
          <el-col :span="3" :offset="0">
            <el-form-item label="奖项" label-width="70px">
              <div v-if="val.r_type==='1'">
                <span v-if="val.t_awards==='1'">
                  <svg class="icon" aria-hidden="true"><use xlink:href="#icon-jinpai"></use></svg>
                  <span style="margin-left: 10px">金牌</span>
                </span>
                <span v-if="val.t_awards==='2'">
                  <svg class="icon" aria-hidden="true"><use xlink:href="#icon-yinpai"></use></svg>
                  <span style="margin-left: 10px">银牌</span>
                </span>
                <span v-if="val.t_awards==='3'">
                  <svg class="icon" aria-hidden="true"><use xlink:href="#icon-tongpai"></use></svg>
                  <span style="margin-left: 10px">铜牌</span>
                </span>
                <span v-if="val.t_awards==='4'">
                  <svg class="icon" aria-hidden="true"><use xlink:href="#icon-tiepai"></use></svg>
                  <span style="margin-left: 10px">铁牌</span>
                </span>
              </div>
              <div v-if="val.r_type==='2'">
                <span v-if="val.t_awards==='0'">
                  <svg class="icon" aria-hidden="true"><use xlink:href="#icon-tedengjiang"></use></svg>
                  <span style="margin-left: 10px">特等奖</span>
                </span>
                <span v-if="val.t_awards==='1'">
                  <svg class="icon" aria-hidden="true"><use xlink:href="#icon-jinpai"></use></svg>
                  <span style="margin-left: 10px">一等奖</span>
                </span>
                <span v-if="val.t_awards==='2'">
                  <svg class="icon" aria-hidden="true"><use xlink:href="#icon-yinpai"></use></svg>
                  <span style="margin-left: 10px">二等奖</span>
                </span>
                <span v-if="val.t_awards==='3'">
                  <svg class="icon" aria-hidden="true"><use xlink:href="#icon-tongpai"></use></svg>
                  <span style="margin-left: 10px">三等奖</span>
                </span>
                <span v-if="val.t_awards==='4'">
                  <svg class="icon" aria-hidden="true"><use xlink:href="#icon-tiepai"></use></svg>
                  <span style="margin-left: 10px">铁</span>
                </span>
              </div>
              <div v-if="val.t_star==='2'">
                <span>
                  <svg class="icon" aria-hidden="true"><use xlink:href="#icon-asterisk"></use></svg>
                  <span style="margin-left: 10px">Star</span>
                </span>
              </div>
            </el-form-item>
          </el-col>
          <el-col :span="2" :offset="0">
            <el-form-item label="Rank" label-width="40px">
              {{val.t_rank}}
            </el-form-item>
          </el-col>
        </el-row>
      </div>
        <div>
          <el-alert title="无 CPC 记录" type="info" center show-icon :closable="false"
                    v-if="!contestForm" style="margin-bottom: 10px"></el-alert>
        </div>
      </div>
    </el-form>
    <el-divider content-position="left">其余信息</el-divider>
    <div v-html="infoForm.introduction"></div>
  </el-card>
</template>

<script>
  export default {
    name: "MemberView",
    data() {
      return {
        infoForm: {
          uid: 0,
          stid: '',
          name: '',
          identity: '0',
          privilege: '0',
          email: '',
          phone: '',
          QQ: '',
          URL: '',
          introduction: '',
          avatarUrl: '',
          education: [],
          career: []
        },
        contestForm: [],
      }
    },
    methods: {
      feachInfoForm: async function () {
        const {data: res} = await this.$http.get('user_all_info',{params: {uid: this.$route.params.uid}});
        if (res.code !== 200)
          return this.$message.error(res.msg);
        else {
          this.infoForm = res.data.userAllInfo;
          this.infoForm.identity  = this.infoForm.identity + "";
          this.infoForm.privilege  = this.infoForm.privilege + "";
          if (this.infoForm.education) {
            this.infoForm.education = JSON.parse(this.infoForm.education);
          } else {
            this.infoForm.education = []
          }
          if (this.infoForm.career) {
            this.infoForm.career = JSON.parse(this.infoForm.career);
          } else {
            this.infoForm.career = []
          }
          this.contestForm = res.data.userContest;
        }
      },
      goToContestDetail: function (cid) {
        const { href } = this.$router.resolve(`/contest/`+cid);
        window.open(href, "_blank");
      }
    },
    created() {
      this.infoForm.uid = Number(this.$route.params.uid);
      this.feachInfoForm()
    }
  }
</script>

<style lang="less" scoped>
  .box-card {
    text-align: center;
    position: relative;
    width: 96%;
    left: 2%;
    top: 5%;
  }
  .banner {
    width: 150px;
    .el-alert {
      text-align: center;
      padding: 0;
    }
  }
  .userAvatar {
    width: 150px;
    height: 200px;
  }
  .experienceTip {
    text-align: left;
    height: 40px;
    .el-form-item {
      margin-bottom: 0 !important;
    }
  }
  .icon {
    width: 1em;
    height: 1em;
    vertical-align: -0.15em;
    fill: currentColor;
    overflow: hidden;
  }
  .NotFirstLine {
    margin-top: 20px;
  }
</style>
