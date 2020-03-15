<template>
  <el-card class="box-card">
    <div class="clearfix" style="text-align: left">
      <span>竞赛历史</span>
    </div>
    <el-table :data="lists" style="width: 100%; padding-top: 0" stripe>
      <el-table-column prop="c_time" label="日期" width="100"></el-table-column>
      <el-table-column prop="c_name_zh" label="名称" width="500">
        <template slot-scope="scope">
          <div>{{scope.row.c_name_zh}}</div>
          <div>{{scope.row.c_name_en}}</div>
        </template>
      </el-table-column>
      <el-table-column prop="c_location_zh" label="地点" width="400">
        <template slot-scope="scope">
          <div>{{scope.row.c_location_zh}}</div>
          <div>{{scope.row.c_location_en}}</div>
        </template>
      </el-table-column>
      <el-table-column prop="achievements" label="成绩" width="350">
        <template slot-scope="scope">
          <el-row>
            <el-col :span="5">
              <el-tooltip effect="dark" content="特等奖" placement="top" :enterable="false">
                <span :class="{'hiddenspan': scope.row['r_type']==='1'}">
                  <svg class="icon" aria-hidden="true"><use xlink:href="#icon-tedengjiang"></use></svg>
                  <span style="margin-left: 5px">{{" " + scope.row["num_0"]}}</span>
                </span>
              </el-tooltip>
            </el-col>
            <el-col :span="5">
              <el-tooltip effect="dark" content="金牌" v-if="scope.row['r_type']==='1'"
                          placement="top" :enterable="false">
                <span>
                  <svg class="icon" aria-hidden="true"><use xlink:href="#icon-jinpai"></use></svg>
                  <span style="margin-left: 5px">{{" " + scope.row["num_1"]}}</span>
                </span>
              </el-tooltip>
              <el-tooltip effect="dark" content="一等奖" v-else
                          placement="top" :enterable="false">
                <span>
                  <svg class="icon" aria-hidden="true"><use xlink:href="#icon-jinpai"></use></svg>
                  <span style="margin-left: 5px">{{" " + scope.row["num_1"]}}</span>
                </span>
              </el-tooltip>
            </el-col>
            <el-col :span="5">
              <el-tooltip effect="dark" content="银牌" v-if="scope.row['r_type']==='1'"
                          placement="top" :enterable="false">
                <span>
                  <svg class="icon" aria-hidden="true"><use xlink:href="#icon-yinpai"></use></svg>
                  <span style="margin-left: 5px">{{" " + scope.row["num_2"]}}</span>
                </span>
              </el-tooltip>
              <el-tooltip effect="dark" content="二等奖" v-else
                          placement="top" :enterable="false">
                <span>
                  <svg class="icon" aria-hidden="true"><use xlink:href="#icon-yinpai"></use></svg>
                  <span style="margin-left: 5px">{{" " + scope.row["num_2"]}}</span>
                </span>
              </el-tooltip>
            </el-col>
            <el-col :span="5">
              <el-tooltip effect="dark" content="铜牌" v-if="scope.row['r_type']==='1'"
                          placement="top" :enterable="false">
                <span>
                  <svg class="icon" aria-hidden="true"><use xlink:href="#icon-tongpai"></use></svg>
                  <span style="margin-left: 5px">{{" " + scope.row["num_3"]}}</span>
                </span>
              </el-tooltip>
              <el-tooltip effect="dark" content="三等奖" v-else
                          placement="top" :enterable="false">
                <span>
                  <svg class="icon" aria-hidden="true"><use xlink:href="#icon-tongpai"></use></svg>
                  <span style="margin-left: 5px">{{" " + scope.row["num_3"]}}</span>
                </span>
              </el-tooltip>
            </el-col>
            <el-col :span="4">
              <el-tooltip effect="dark" content="铁牌" placement="top" :enterable="false">
                <span>
                  <svg class="icon" aria-hidden="true"><use xlink:href="#icon-tiepai"></use></svg>
                  <span style="margin-left: 5px">{{" " + scope.row["num_4"]}}</span>
                </span>
              </el-tooltip>
            </el-col>
          </el-row>
        </template>
      </el-table-column>
      <el-table-column prop="options" label="详情">
        <template slot-scope="scope">
          <el-tooltip effect="dark" content="查看详细信息" placement="top" :enterable="false">
            <el-button type="info" icon="el-icon-files" size="mini" @click="goToContestDetail(scope.row.c_id)"></el-button>
          </el-tooltip>
        </template>
      </el-table-column>
    </el-table>
  </el-card>
</template>

<script>
  export default {
    name: "Contests",
    data() {
      return {
        lists: []
      }
    },
    methods: {
      getContestList: async function () {
        const {data: res} = await this.$http.get('contest_list');
        if (res.code !== 200)
          return this.$message.error(res.msg);
        this.lists = res.data;
      },
      goToContestDetail(cid) {
        const { href } = this.$router.resolve(`/contest/`+cid);
        window.open(href, "_blank");
      }
    },
    created() {
      this.getContestList()
    },
  }
</script>

<style scoped>
  .icon {
    width: 1em;
    height: 1em;
    vertical-align: -0.15em;
    fill: currentColor;
    overflow: hidden;
  }
  .hiddenspan {
    visibility: hidden;
  }
</style>
