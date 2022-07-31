<template>
  <div>
    <el-upload
      action=""
      class="upload-demo"
      :http-request="uploadFile"
      :on-preview="handlePreview"
      :on-remove="handleRemove"
      multiple
      :limit="3"
      :on-exceed="handleExceed"
      :file-list="fileList"
    >
      <el-button size="small" type="primary">点击上传</el-button>
      <div slot="tip" class="el-upload__tip">只能上传jpg/png文件，且不超过500kb</div>
    </el-upload>

    <img v-for="item in fileNames" :key="item" :src="`https://guide.time-traveler.cn/utils/getPic/${item}`" alt="">
  </div>
</template>

<script>
import api from '@/api'

export default {
  name: 'TextPage',
  data() {
    return {
      fileList: [],
      fileNames: []
    }
  },
  mounted() {
    api.getPic({ pic_name: '1659198494猫猫1.jpeg' }).then(res => {
      console.log(res)
    })
  },
  methods: {
    handlePreview(...params) {
      console.log(params)
    },
    handleRemove(fileObj) {
      this.fileNames = this.fileNames.filter(item => {
        const str = item.substring(10, item.length)
        return fileObj.name !== str
      })
    },
    handleExceed(...params) {
      console.log(params)
    },
    uploadFile(params) {
      console.log(params.file)
      const formData = new FormData()
      formData.append('file', params.file)
      formData.append('file_type', 'pic')
      api.uploadFile(formData).then(res => {
        this.fileNames.push(res.data.file_name)
        console.log(this.fileNames)
      })
    }
  }
}
</script>

<style scoped>

</style>
