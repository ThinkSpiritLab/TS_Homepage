<template>
  <div>
    <el-carousel indicator-position="outside" height="547px">
      <el-carousel-item v-for="item in imgList" :key="item.id" style="text-align: center">
        <img :src="item.url" class="imgTop">
      </el-carousel-item>
    </el-carousel>
    <el-row :gutter="20">
      <el-col :span="16">
        <el-card class="moments" body-style="padding-top: 0">
          <div slot="header" class="clearfix">
            <span>动态</span>
            <span style="float: right; padding: 3px 0; color: #999999"
            >更多动态请查看 竞赛历史 及 新闻与活动 栏目</span>
          </div>
          <el-table :data="moments" style="width: 100%; margin-top: 0" :show-header="false">
            <el-table-column prop="date" label="日期" width="180px"></el-table-column>
            <el-table-column prop="title" label="事件">
              <template slot-scope="scope">
                <span @click="goToDetail(scope.row)" class="momentsItem">{{scope.row.title}}</span>
              </template>
            </el-table-column>
          </el-table>
        </el-card>
      </el-col>
      <el-col :span="8">
        <div style="text-align: center; margin-bottom: 40px">
          <el-button class="web-font-main" type="primary" @click="goToAbout">关于我们 & 加入我们</el-button>
        </div>
        <el-card class="bulletinIndex" body-style="padding-top: 0">
          <div slot="header" class="clearfix">
            <span>公告</span>
            <el-button style="float: right; padding: 3px 0" type="text">查看全部公告</el-button>
          </div>
          <div class="bulletin">
            <el-table :data="bulletin" style="width: 100%; margin-top: 0" :show-header="false">
              <el-table-column prop="date" label="日期" width="180px"></el-table-column>
              <el-table-column prop="title" label="标题">
                <template slot-scope="scope">
                  <span @click="goToBulletinDetail(scope.row.bid)" class="momentsItem">{{scope.row.title}}</span>
                </template>
              </el-table-column>
            </el-table>
          </div>
        </el-card>

      </el-col>
    </el-row>
  </div>
</template>

<script>
  export default {
    name: "Main",
    data() {
      return {
        moments: [],
        bulletin: [],
        imgList: []
      }
    },
    methods: {
      getMoments: async function () {
        const {data: res} = await this.$http.get('moments');
        if (res.code !== 200)
          return this.$message.error(res.msg);
        this.moments = res.data;
      },
      getBulletinIndex: async function() {
        const {data: res} = await this.$http.get('bulletinIndex');
        if (res.code !== 200)
          return this.$message.error(res.msg);
        this.bulletin = res.data;
      },
      goToDetail: function (item) {
        if (item.type === "c") {
          const { href } = this.$router.resolve(`contest/`+item.id);
          window.open(href, "_blank");
        } else if (item.type === "n") {
          const { href } = this.$router.resolve(`news/`+item.id);
          window.open(href, "_blank");
        }
      },
      goToBulletinDetail: function(bid) {
        const { href } = this.$router.resolve(`announcement/`+bid);
        window.open(href, "_blank");
      },
      goToAbout: function () {
        this.$router.push('about')
      },
      getIndexImgList: async function () {
        const {data: res} = await this.$http.get('indexImgList');
        if (res.code !== 200)
          return this.$message.error(res.msg);
        this.imgList = res.data;
      },
    },
    created() {
      this.getMoments();
      this.getBulletinIndex();
      this.getIndexImgList()
    }
  }
</script>

<style lang="less" scoped>
  .imgTop {
    width: 1360px;
    height: 547px;
  }
  .moments {
    height: 400px;
  }
  .bulletinIndex {
    height: 300px;
  }
  @font-face {
    font-family: 'webfont-main';
    font-display: swap;
    src: url('//at.alicdn.com/t/webfont_9h7p6t9mmm.eot'); /* IE9*/
    src: url('//at.alicdn.com/t/webfont_9h7p6t9mmm.eot?#iefix') format('embedded-opentype'), /* IE6-IE8 */
    url('//at.alicdn.com/t/webfont_9h7p6t9mmm.woff2') format('woff2'),
    url('//at.alicdn.com/t/webfont_9h7p6t9mmm.woff') format('woff'), /* chrome、firefox */
    url('//at.alicdn.com/t/webfont_9h7p6t9mmm.ttf') format('truetype'), /* chrome、firefox、opera、Safari, Android, iOS 4.2+*/
    url('//at.alicdn.com/t/webfont_9h7p6t9mmm.svg#杨任东竹石体-Bold') format('svg'); /* iOS 4.1- */
  }
  .web-font-main{
    font-family:"webfont-main" !important;
    font-size:30px;font-style:normal;
    -webkit-font-smoothing: antialiased;
    -webkit-text-stroke-width: 0.2px;
    -moz-osx-font-smoothing: grayscale;
  }
  .momentsItem:hover {
    text-decoration: underline;
    cursor: pointer;
  }
</style>
