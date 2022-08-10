<template>
  <div class="image-upload-container">
    <el-upload
      action=""
      class="upload-demo"
      :http-request="uploadFile"
      :on-preview="handlePreview"
      :on-remove="handleRemove"
      :on-change="handleChange"
      multiple
      :limit="limit"
      :on-exceed="handleExceed"
      :file-list="fileList"
      :list-type="fileType === 'pic' ? 'picture-card' : 'text'"
    >
      <el-button size="small" type="primary">点击上传</el-button>
      <div slot="tip" class="el-upload__tip">视频可以上传mp4，音频可以上传mp3，图片支持png，jpg格式</div>
    </el-upload>
    <!--    <img v-for="item in fileNames" :key="item" :src="`https://guide.time-traveler.cn/utils/getPic/${item}`" alt="">-->
  </div>
</template>

<script>
import api from '@/api'

export default {
  name: 'MediaUpload',
  props: {
    limit: {
      type: Number,
      default: 1
    },
    fileType: {
      type: String,
      default: 'pic'
    },
    fileList: {
      type: Array,
      default: () => ([])
    }
  },
  data() {
    return {
      // fileList: [],
      fileNames: []
    }
  },
  watch: {
    fileNames(val) {
      this.$emit('file-list-change', val)
    },
    fileList(val) {
      this.fileNames = val.map(item => item.name)
      console.log(this.fileNames)
    }
  },
  methods: {
    // 点击预览
    handlePreview(...params) {
      console.log(params)
    },
    handleRemove(fileObj) {
      console.log(fileObj)
      this.fileNames = this.fileNames.filter(item => fileObj.name !== item)
    },
    handleExceed(...params) {
      this.$message.info(`只能上传${this.limit}张图片哦！`)
    },
    handleChange(...params) {
      console.log(params)
    },
    uploadFile(params) {
      const formData = new FormData()
      formData.append('file', params.file)
      formData.append('file_type', this.fileType)
      api.uploadFile(formData).then(res => {
        this.fileNames.push(res.data.file_name)
      })
    }
  }
}
</script>

<style scoped>

</style>
