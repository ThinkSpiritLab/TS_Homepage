<template>
  <el-card class="box-card">
    <el-page-header @back="goToHome" content="添加竞赛记录"></el-page-header>
    <el-divider></el-divider>
    <div>
      <div class="form1" v-show="activeStep===0">
        <el-form :model="addForm.baseInfo" :rules="addFormRules" :disabled="loading"
                 ref="addForm1" label-width="110px">
          <el-form-item label="竞赛名称(中)" prop="c_name_zh">
            <el-input v-model="addForm.baseInfo.c_name_zh"></el-input>
          </el-form-item>
          <el-form-item label="竞赛名称(英)" prop="c_name_en">
            <el-input v-model="addForm.baseInfo.c_name_en"></el-input>
          </el-form-item>
          <el-form-item label="竞赛地点(中)" prop="c_location_zh">
            <el-input v-model="addForm.baseInfo.c_location_zh"></el-input>
          </el-form-item>
          <el-form-item label="竞赛地点(英)" prop="c_location_en">
            <el-input v-model="addForm.baseInfo.c_location_en"></el-input>
          </el-form-item>
          <el-form-item label="比赛日期" prop="c_time">
            <el-date-picker v-model="addForm.baseInfo.c_time" type="date" placeholder="选择日期" value-format="yyyy-MM-dd">
            </el-date-picker>
          </el-form-item>
        </el-form>
      </div>
      <div class="form2" v-show="activeStep===1">
        <el-form :model="addForm.resultInfo" :disabled="loading"
                 ref="addForm2" label-width="110px">
          <el-row class="rowClass">
            <el-col :span="10">
              <el-form-item label="奖项类型" prop="c_type">
                <el-switch v-model="addForm.resultInfo.r_type"
                           inactive-color="#13ce66" active-color="#0081ff" active-value="2" inactive-value="1"
                           inactive-text="金牌 银牌 铜牌 铁牌" active-text="特等 一等 二等 三等 铁">
                </el-switch>
              </el-form-item>
            </el-col>
            <el-col :span="7">
              <el-form-item label="全场队伍总数" prop="t_all_sum">
                <el-input-number v-model="addForm.resultInfo.total_team_num" size="small" :min="1"></el-input-number>
              </el-form-item>
            </el-col>
            <el-col :span="7">
              <el-form-item>
                <el-button @click="addTeam" type="primary">新增本校队伍</el-button>
              </el-form-item>
            </el-col>
          </el-row>
          <div v-for="(result, index) in addForm.resultInfo.results">
            <el-divider class="mydivider"><i class="el-icon-help"></i></el-divider>
            <el-row>
              <el-col :span="2"><el-form-item :label="'队伍 ' + (index+1)"></el-form-item></el-col>
              <el-col :span="11">
                <el-form-item label="队伍名(中)" prop="t_name_zh">
                  <el-input v-model="result.t_name_zh"></el-input>
                </el-form-item>
              </el-col>
              <el-col :span="11">
                <el-form-item label="队伍名(英)"  prop="t_name_zh">
                  <el-input v-model="result.t_name_en"></el-input>
                </el-form-item>
              </el-col>
            </el-row>
            <el-row>
              <el-col :span="3"><el-form-item></el-form-item></el-col>
              <el-col :span="7">
                <el-form-item label="队员 1" prop="t_mem1_name" label-width="60px">
                  <el-autocomplete popper-class="my-autocomplete" v-model="result.t_mem1_stid"
                                   :fetch-suggestions="queryStidNameAsync" placeholder="学号">
                    <template slot-scope="{ item }">
                      <div class="stid">{{item.value+" "}}<span class="name">{{item.name}}</span></div>
                    </template>
                  </el-autocomplete>
                </el-form-item>
              </el-col>
              <el-col :span="7">
                <el-form-item label="队员 2" prop="t_mem2_stid" label-width="60px">
                  <el-autocomplete popper-class="my-autocomplete" v-model="result.t_mem2_stid"
                                   :fetch-suggestions="queryStidNameAsync" placeholder="学号">
                    <template slot-scope="{ item }">
                      <div class="stid">{{item.value+" "}}<span class="name">{{ item.name}}</span></div>
                    </template>
                  </el-autocomplete>
                </el-form-item>
              </el-col>
              <el-col :span="7">
                <el-form-item label="队员 3" prop="t_mem3_stid" label-width="60px">
                  <el-autocomplete popper-class="my-autocomplete" v-model="result.t_mem3_stid"
                                   :fetch-suggestions="queryStidNameAsync" placeholder="学号">
                    <template slot-scope="{ item }">
                      <div class="stid">{{item.value+" "}}<span class="name">{{item.name}}</span></div>
                    </template>
                  </el-autocomplete>
                </el-form-item>
              </el-col>
            </el-row>
            <el-row class="rowClass">
              <el-col :span="2"><el-form-item></el-form-item></el-col>
              <el-col :span="3">
                <el-form-item label="star" prop="t_star">
                  <el-switch v-model="result.t_star" active-color="#13ce66" inactive-color="#ff4949"
                             active-value="2" inactive-value="1"></el-switch>
                </el-form-item>
              </el-col>
              <el-col :span="5">
                <el-form-item label="名次" prop="t_rank">
                  <el-input-number v-model="result.t_rank" size="small"></el-input-number>
                </el-form-item>
              </el-col>
              <el-col :span="5">
                <el-form-item label="奖项" prop="t_awards" v-if="result.t_star==='1'">
                  <el-select v-model="result.t_awards" placeholder="请选择">
                    <el-option v-for="item in award_options.type1" v-if="addForm.resultInfo.r_type==='1'"
                               :key="item.value" :label="item.label" :value="item.value">
                      <svg class="icon" aria-hidden="true"><use :xlink:href="item.icon"></use></svg>{{item.label}}
                    </el-option>
                    <el-option v-for="item in award_options.type2" v-if="addForm.resultInfo.r_type==='2'"
                               :key="item.value" :label="item.label" :value="item.value">
                      <svg class="icon" aria-hidden="true"><use :xlink:href="item.icon"></use></svg>{{item.label}}
                    </el-option>
                  </el-select>
                </el-form-item>
              </el-col>
              <el-col :span="8">
                <el-button style="margin-left: 100px" @click.prevent="removeTeam(result)"
                 :disabled="addForm.resultInfo.results.length<=1" type="danger">删除该队伍</el-button>
              </el-col>
            </el-row>
          </div>
        </el-form>
      </div>
      <div class="myeditor" v-show="activeStep===2">
        <div class="home">
          <tinymce ref="editor" v-model="addForm.extras"/>
        </div>
      </div>
    </div>
    <el-divider class="mydivider"></el-divider>
    <div>
      <el-row>
        <el-col :span="16">
          <el-steps :active="activeStep" process-status="finish" finish-status="process" align-center>
            <el-step title="基本信息"></el-step>
            <el-step title="成绩信息"></el-step>
            <el-step title="额外信息"></el-step>
          </el-steps>
        </el-col>
        <el-col :span="8">
          <el-button :disabled="activeStep<=0" @click="activeStep--" type="primary" plain>上一步</el-button>
          <el-button :disabled="activeStep>=2" style="margin-left: 30px" @click="activeStep++" type="primary" plain>下一步</el-button>
          <el-button v-show="activeStep===2" style="margin-left: 30px" type="success" @click="submitContest">完 成</el-button>
        </el-col>
      </el-row>
    </div>
  </el-card>
</template>

<script>
  import tinymce from "../../tinymce";
  import icon from "../../../assets/font_awards/iconfont.js"
  export default {
    name: "AddContest",
    components: {
      tinymce
    },
    data() {
      return {
        activeStep: 0,
        addForm: {
          baseInfo: {
            c_name_zh: '',
            c_name_en: '',
            c_location_zh: '',
            c_location_en: '',
            c_time: '',
          },
          resultInfo: {
            r_type: '1',
            total_team_num: 1,
            results: [{
              t_name_zh: '',
              t_name_en: '',
              t_mem1_stid: '',
              t_mem2_stid: '',
              t_mem3_stid: '',
              t_star: '1',
              t_rank: 1,
              t_awards: '4'
            }]
          },
          extras: ''
        },
        loading: false,
        award_options: {
          type1: [{value: '1', label: '金牌', icon: "#icon-jinpai"}, {value: '2', label: '银牌', icon: "#icon-yinpai"},
            {value: '3', label: '铜牌', icon: "#icon-tongpai"}, {value: '4', label: '铁牌', icon: "#icon-tiepai"}],
          type2: [{value: '0', label: '特等奖', icon: "#icon-tedengjiang"}, {value: '1', label: '一等奖', icon: "#icon-jinpai"},
            {value: '2', label: '二等奖', icon: "#icon-yinpai"}, {value: '3', label: '三等奖', icon: "#icon-tongpai"},
            {value: '4', label: '铁', icon: "#icon-tiepai"}],
        },
        addFormRules: {
          c_name_zh: [{required: true, message: '请输本次竞赛中文名称', trigger: 'blur'}],
          c_name_en: [{required: true, message: '请输本次竞赛英文名称', trigger: 'blur'}],
          c_location_zh: [{required: true, message: '请输竞赛地点中文名称', trigger: 'blur'}],
          c_location_en: [{required: true, message: '请输竞赛地点英文名称', trigger: 'blur'}],
          c_time: [{required: true, message: '请输竞赛时间', trigger: 'blur'}],
        },
      }
    },
    methods: {
      goToHome: function () {
        this.$router.push('/console_home');
      },
      addTeam: function () {
        this.addForm.resultInfo.results.push({
          t_name_zh: '',
          t_name_en: '',
          t_mem1_stid: '',
          t_mem2_stid: '',
          t_mem3_stid: '',
          t_star: 0,
          t_rank: 1,
          t_awards: '4'
        })
      },
      removeTeam: function (item) {
        const index = this.addForm.resultInfo.results.indexOf(item);
        if (index !== -1) {
          this.addForm.resultInfo.results.splice(index, 1)
        }
      },
      queryStidNameAsync: async function (queryString, cb) {
        if (queryString==="")
          cb([]);
        else {
          const {data: res} = await this.$http.get('user_name_by_dim_stid',
              {params: {stid: queryString+""}});
          if (res.code === 200) {
            if (res.data)
              cb(res.data);
            else
              cb([]);
          } else {
            cb([]);
          }
        }
      },
      submitContest: async function() {
        this.loading = true;
        const { data:result } = await this.$http.post('/contest/contest_add', this.addForm);
        if (result.code !== 200) {
          this.$message.error(result.msg);
        } else {
          this.$message.success('添加竞赛记录成功');
          this.formReset();
        }
        this.loading = false;
      },
      formReset: function () {
        this.addForm.baseInfo.c_name_zh = this.addForm.baseInfo.c_name_en = this.addForm.baseInfo.c_location_zh = this.addForm.baseInfo.c_location_en = this.addForm.baseInfo.c_time = '';
        this.addForm.resultInfo.total_team_num = 1;
        this.addForm.resultInfo.results = [{t_name_zh: '', t_name_en: '', t_mem1_stid: '', t_mem2_stid: '', t_mem3_stid: '', t_star: '1', t_rank: 1, t_awards: '1'}];
        this.addForm.extras = '';
        this.activeStep = 0
      }
    }
  }
</script>

<style lang="less" scoped>
  .rowClass {
    height: 45px;
    margin-bottom: 0;
    .el-col {
      height: 45px;
    }
    .el-form-item {
      margin-bottom: 0;
    }
  }
  .el-divider {
    margin: 15px;
  }
  .my-autocomplete {
    li {
      line-height: normal;
      padding: 7px;

      .stid {
        text-overflow: ellipsis;
        overflow: hidden;
      }
      .name {
        font-size: 12px;
        color: #b4b4b4;
      }

      .highlighted .addr {
        color: #ddd;
      }
    }
  }
  .mydivider {
    height: 2px;
    /*background-color: #409EFF;*/
    background-color: #D0D0D0;
  }
  .icon {
     width: 1em;
     height: 1em;
     vertical-align: -0.15em;
     fill: currentColor;
     overflow: hidden;
   }
</style>
