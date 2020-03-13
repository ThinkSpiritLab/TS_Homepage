<template>
  <el-card class="cardBox">
    <template>
      <el-tabs v-model="activeName">
        <el-tab-pane label="教师与教练" name="faculty">
<!--          <el-row>-->
            <el-col :span="6" v-for="(o, index) in 5" :key="o" :offset="index%3 > 0 ? 3 : 0">
              <el-card :body-style="{ padding: '0px' }">
                <img src="https://shadow.elemecdn.com/app/element/hamburger.9cf7b091-55e9-11e9-a976-7f4d0b07eef6.png" class="image">
                <div style="padding: 14px;">
                  <span>好吃的汉堡</span>
                  <div class="bottom clearfix">
                    <el-button type="text" class="button">操作按钮</el-button>
                  </div>
                </div>
              </el-card>
            </el-col>
<!--          </el-row>-->
        </el-tab-pane>
        <el-tab-pane label="毕业成员" name="alumni">毕业成员</el-tab-pane>
        <el-tab-pane label="在校成员" name="students">在校成员</el-tab-pane>
      </el-tabs>
    </template>
  </el-card>

</template>

<script>
  export default {
    name: "Members",
    data() {
      return {
        activeName: 'faculty'
      }
    }
  }
</script>

<style scoped>
  .cardBox {
    height: auto;
  }
</style>
