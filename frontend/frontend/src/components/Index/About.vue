<template>
  <div>
    <el-card class="aboutCard"header="About Us">
      暂缺的介绍
    </el-card>
    <el-card class="JoinCard" header="加入我们">
      <div>首先感谢您有兴趣选择 Think Spirit Lab 作为大学期间的学习平台，为帮助您更加确信您的选择是否明智，请浏览本网站以便知晓我们的活动与工作，并耐心阅读如下说明：</div>
      <div>重要：您是否了解我们的学习方向？在大致了解的基础上，您是否确实对这些方向感兴趣？这条路可能充满挫败感，如果没有足够的兴趣支撑，可能会感到痛苦。</div>
      <div>重要：您是否接受我们互助友善的价值观？无论是算法竞赛还是开发，缺乏互助友善是件非常糟糕的事情。尽管我们会进行价值观培养，但如果您不能赞同我们的价值观，Think Spirit Lab 可能不是您的最佳选择。</div>
      <div>对我们来说，理想的申请者大致是这样的：愿意学习编程，乐于研究算法或是参与开发，同时逻辑思维清晰、表达能力强，最好还能乐观积极、坚韧不拔、亲和力强。但上述要求并不是绝对的。</div>
      <div>若您对 Think Spirit Lab 有足够的兴趣，可与大三或大四的在校学长学姐取得联系，获得更多信息。</div>
      <div>需要注意的是，即使与某位学长或学姐相谈甚欢，您仍需通过 Think Spirit Lab 的正式面谈，才能获得加入资格，二者之间没有直接联系。</div>
      <div>我们将于每年 9~10 月份组织招新面谈，详情请静候公告。面谈的目的是给我们双方一个面对面了解对方的机会，以此来帮助您判断是否适合到 Think Spirit Lab 工作和学习。因此，被接受并不意味着成功，没被接受当然也不是失败。</div>
    </el-card>
  </div>
</template>

<script>
  export default {
    name: "About"
  }
</script>

<style lang="less" scoped>
  .aboutCard {
    /*height: 400px;*/
    position: relative;
    width: 96%;
    left: 2%;
  }
  .JoinCard {
    margin-top: 20px;
    position: relative;
    width: 96%;
    left: 2%;
    div {
      margin-bottom: 20px;
    }
  }
</style>
