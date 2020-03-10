<template>
  <div class="tinymce-box">
    <editor v-model="myValue"
            :init="init"
            :disabled="disabled"
            @onClick="onClick">
    </editor>
  </div>
</template>

<script>
  import tinymce from 'tinymce/tinymce' //tinymce默认hidden，不引入不显示
  import Editor from '@tinymce/tinymce-vue'
  import 'tinymce/themes/silver'
  // 编辑器插件plugins
  // 更多插件参考：https://www.tiny.cloud/docs/plugins/
  import 'tinymce/plugins/image'// 插入上传图片插件
  import 'tinymce/plugins/media'// 插入视频插件
  import 'tinymce/plugins/table'// 插入表格插件
  import 'tinymce/plugins/lists'// 列表插件
  import 'tinymce/plugins/wordcount'// 字数统计插件
  export default {
    components:{
      Editor
    },
    name:'tinymce',
    props: {
      value: {
        type: String,
        default: ''
      },
      disabled: {
        type: Boolean,
        default: false
      },
      plugins: {
        type: [String, Array],
        default: 'lists image media table wordcount'
      },
      toolbar: {
        type: [String, Array],
        default: 'undo redo |  formatselect | fontsizeselect bold italic forecolor backcolor underline strikethrough | alignleft aligncenter alignright alignjustify | bullist numlist outdent indent | lists image media table | removeformat blockquote'
      }
    },
    data(){
      return{
        init: {
          language_url: '/tinymce/langs/zh_CN.js',
          language: 'zh_CN',
          skin_url: '/tinymce/skins/ui/oxide',
          // skin_url: 'tinymce/skins/ui/oxide-dark',//暗色系
          height: 400,
          plugins: this.plugins,
          toolbar: this.toolbar,
          branding: false,
          menubar: false,

          images_upload_handler: async (blobInfo, success, failure) => {
            if (blobInfo.blob().size > 5242880) {
              failure('图片大小不得超过 5 MB')
            }
            let formData = new FormData();
            // 服务端接收文件的参数名，文件数据，文件名
            formData.append('upfile', blobInfo.blob(), blobInfo.filename());
            const {data: res} = await this.$http.post('upload_image', formData)
            if (res.code === 200) {
              // 这里返回的是你图片的地址
              success(res.data)
            } else {
              failure(res.msg)
            }
          }
        },
        myValue: this.value
      }
    },
    mounted () {
      tinymce.init({})
    },
    methods: {
      // 添加相关的事件，可用的事件参照文档=> https://github.com/tinymce/tinymce-vue => All available events
      // 需要什么事件可以自己增加
      onClick (e) {
        this.$emit('onClick', e, tinymce)
      },
      // 可以添加一些自己的自定义事件，如清空内容
      clear () {
        this.myValue = ''
      }
    },
    watch: {
      value (newValue) {
        this.myValue = newValue
      },
      myValue (newValue) {
        this.$emit('input', newValue)
      }
    }
  }

</script>
<style scoped>

</style>
