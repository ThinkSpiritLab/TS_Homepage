<template>
  <el-card class="cardBox">
    <div slot="header" class="clearfix" style="text-align: left">
      <span>竞赛信息</span>
    </div>
    <div>
      <el-row>
        <el-col :span="12">
          <span class="labelInfo">竞赛名称</span>
          <span>{{contestInfo.c_name_zh}}</span>
        </el-col>
        <el-col :span="12">
          <span class="labelInfo">Contest Name</span>
          <span>{{contestInfo.c_name_en}}</span>
        </el-col>
      </el-row>
      <el-row class="rowInfo">
        <el-col :span="12">
          <span class="labelInfo">竞赛地点</span>
          <span>{{contestInfo.c_location_zh}}</span>
        </el-col>
        <el-col :span="12">
          <span class="labelInfo">Venue of Competition</span>
          <span>{{contestInfo.c_location_en}}</span>
        </el-col>
      </el-row>
      <el-row class="rowInfo">
        <el-col :span="12">
          <span class="labelInfo">竞赛日期</span>
          <span>{{contestInfo.c_time}}</span>
        </el-col>
        <el-col :span="12">
          <span class="labelInfo">队伍总数</span>
          <span>{{contestInfo.total_team_num}}</span>
        </el-col>
      </el-row>
    </div>
    <el-divider content-position="left" class="dividerInfo">竞赛记忆</el-divider>
    <div>
      <div v-html="contestInfo.c_extras" v-if="contestInfo.c_extras !== ''"></div>
      <el-alert title="无记录" type="info" center show-icon :closable="false"  v-else
                style="margin-bottom: 10px"></el-alert>
    </div>
    <el-divider content-position="left" class="dividerInfo">我校赛况</el-divider>
    <div>
      <div v-for="(val, index) in teams" :key="index">
        <el-divider v-if="index!==0"><i class="el-icon-help"></i></el-divider>
        <el-row>
          <el-col :span="10">
            <span class="labelInfo">队伍名称</span>
            <span>{{val.t_name_zh}}</span>
          </el-col>
          <el-col :span="9">
            <span class="labelInfo">Team Name</span>
            <span>{{val.t_name_en}}</span>
          </el-col>
          <el-col :span="3">
            <span class="labelInfo">奖项</span>
              <span v-if="contestInfo.r_type==='1'">
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
              </span>
              <span v-if="contestInfo.r_type==='2'">
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
              </span>
              <span v-if="val.t_star==='2'">
                <span>
                  <svg class="icon" aria-hidden="true"><use xlink:href="#icon-asterisk"></use></svg>
                  <span style="margin-left: 10px">Star</span>
                </span>
              </span>
          </el-col>
          <el-col :span="2">
            <span class="labelInfo">Rank</span>
              {{val.t_rank}}
          </el-col>
        </el-row>
        <el-row class="rowInfo">
          <el-col :span="8">
            <el-col :span="4">
              <span>队员 1</span>
            </el-col>
            <el-col :span="7">
              <span>{{val.t_mem1_stid}}</span>
            </el-col>
            <el-col :span="4">
              <span>{{val.t_mem1_name}}</span>
            </el-col>
            <el-col :span="9">
              <el-tooltip effect="dark" content="查看个人主页" placement="top" :enterable="false">
                <el-button type="info" icon="el-icon-user" size="mini" @click="viewUserDetail(val.t_mem1_uid)"></el-button>
              </el-tooltip>
            </el-col>
          </el-col>
          <el-col :span="8">
            <el-col :span="4">
              <span>队员 2</span>
            </el-col>
            <el-col :span="7">
              <span>{{val.t_mem2_stid}}</span>
            </el-col>
            <el-col :span="4">
              <span>{{val.t_mem2_name}}</span>
            </el-col>
            <el-col :span="9">
              <el-tooltip effect="dark" content="查看个人主页" placement="top" :enterable="false">
                <el-button type="info" icon="el-icon-user" size="mini" @click="viewUserDetail(val.t_mem2_uid)"></el-button>
              </el-tooltip>
            </el-col>
          </el-col>
          <el-col :span="8">
            <el-col :span="4">
              <span>队员 3</span>
            </el-col>
            <el-col :span="7">
              <span>{{val.t_mem3_stid}}</span>
            </el-col>
            <el-col :span="4">
              <span>{{val.t_mem3_name}}</span>
            </el-col>
            <el-col :span="9">
              <el-tooltip effect="dark" content="查看个人主页" placement="top" :enterable="false">
                <el-button type="info" icon="el-icon-user" size="mini" @click="viewUserDetail(val.t_mem3_uid)"></el-button>
              </el-tooltip>
            </el-col>
          </el-col>
        </el-row>
      </div>
    </div>
  </el-card>
</template>

<script>
  export default {
    name: "ContestView",
    data() {
      return {
        contestInfo: {
          cid: 0,
          c_name_zh: '',
          c_name_en: '',
          c_location_zh: '',
          c_location_en: '',
          c_time: '',
          r_type: '',
          total_team_num: 0,
          c_extras: '',
        },
        teams: [{
          t_name_zh: '',
          t_name_en: '',
          t_mem1_stid: '',
          t_mem1_name: '',
          t_mem1_uid: 0,
          t_mem2_stid: '',
          t_mem2_name: '',
          t_mem2_uid: 0,
          t_mem3_stid: '',
          t_mem3_name: '',
          t_mem3_uid: 0,
          t_star: '',
          t_rank: 0,
          t_awards: ''
        }]
      }
    },
    methods: {
      fetchContestInfo: async function () {
        const {data: res} = await this.$http.get('contest_detail',{params: {cid: this.$route.params.cid}});
        if (res.code !== 200)
          return this.$message.error(res.msg);
        else {
          this.contestInfo = res.data.contest;
          this.teams = res.data.teams
        }
      },
      viewUserDetail: function (uid) {
        const { href } = this.$router.resolve(`/member/`+uid);
        window.open(href, "_blank");
      }
    },
    created() {
      this.contestInfo.cid = Number(this.$route.params.cid)
      this.fetchContestInfo();
    }
  }
</script>

<style scoped>
  .cardBox {
    position: relative;
    top: 5%;
  }
  .rowInfo {
    margin-top: 40px;
  }
  .labelInfo {
    margin-right: 20px;
    color: #606266;
  }
  .dividerInfo {
    margin: 40px 0;
    height: 3px;
  }
  .icon {
    width: 1em;
    height: 1em;
    vertical-align: -0.15em;
    fill: currentColor;
    overflow: hidden;
  }
</style>
