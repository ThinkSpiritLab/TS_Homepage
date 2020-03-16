<template>
  <el-card>
    <el-upload
            class="upload-demo"
            action="default" :http-request = "customUpload"  :before-upload="beforeAvatarUpload"
            :auto-upload="true" :file-list="fileList" list-type="picture" :on-remove="handleRemove">
      <el-button size="small" type="primary">点击上传</el-button>
      <div slot="tip" class="el-upload__tip">
        <div>只能上传jpg/png文件，且不超过 1Mb</div>
        <div>为保证显示效果，图片大小应为 1360*547 或同比例缩放</div>
      </div>
    </el-upload>
    <el-button @click="submitIndexImg">确定</el-button>
  </el-card>
</template>

<script>
  export default {
    name: "CarouseSetting",
    data() {
      return {
        fileList: []
      };
    },
    methods: {
      beforeAvatarUpload(file) {
        const isJPG = file.type === 'image/jpg';
        const isPNG = file.type === 'image/png';
        const isJPEG = file.type === 'image/jpeg';
        const isLt2M = file.size / 1024 / 1024 < 5;
        if (!isJPG && !isPNG && !isJPEG) {
          this.$message.error('上传头像图片只能是 JPG 格式!');
        }
        if (!isLt2M) {
          this.$message.error('上传头像图片大小不能超过 5MB!');
        }
        return (isJPG || isPNG || isJPEG) && isLt2M;
      },
      async customUpload(file) {
        let formData = new FormData();
        formData.append('upfile', file.file, file.file.name);
        const {data: res} = await this.$http.post('upload_image', formData);
        if (res.code === 200) {
          this.fileList.push({name: this.fileList.length, url: res.data})
        } else {
          this.$message(res.msg)
        }
      },
      getIndexImgList: async function () {
        const {data: res} = await this.$http.get('indexImgList');
        if (res.code !== 200)
          return this.$message.error(res.msg);
        this.fileList = res.data;
      },
      handleRemove: function(file) {
        const index = this.fileList.indexOf(file);
        if (index !== -1) {
          this.fileList.splice(index, 1)
        }
      },
      submitIndexImg: async function () {
        const { data:result } = await this.$http.post('/console/setIndexImg', this.fileList);
        if (result.code !== 200) {
          this.$message.error(result.msg);
        } else {
          this.$message.success('设置首页照片墙成功');
        }
      }
    },
    created() {
      this.getIndexImgList()
    }
  }
</script>

<style scoped>

</style>
