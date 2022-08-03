<template>
  <div class="scenic-add-container">
    <h3 class="form-name">主页地图信息</h3>
    <el-form ref="from" :model="form" label-width="150px">
      <el-form-item label="地图底图">
        <media-upload :limit="1" file-type="pic" :file-list="mapImageList" />
      </el-form-item>
      <el-form-item label="图片左上角经纬度">
        <el-input v-model="form.NE"/>
      </el-form-item>
      <el-form-item label="图片右下角经纬度">
        <el-input v-model="form.SW"/>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">保存</el-button>
        <el-button @click="$router.go(-1)">取消</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
import MediaUpload from '@/components/MediaUpload'

export default {
  name: 'MapEdit',
  components: { MediaUpload },
  data () {
    return {
      form: {
        mapImage: '',
        NE: '',
        SW: ''
      },
      bannerList: [],
      categoryOptions: [],
      mapImageList: []
    }
  },
  mounted () {
    // 获取地图信息
    // if (this.$route.query.id) {
    //   this.form = this.$route.query
    //   if (this.form.banner.length) {
    //     this.bannerList = this.form.banner.split('|').map(item => {
    //       return {
    //         name: item,
    //         url: `https://guide.time-traveler.cn/utils/getPic/${item}`
    //       }
    //     })
    //   }
    // }
  },
  methods: {
    onSubmit () {
      // const submitApi = this.form.id ? api.scenic.editScenic : api.scenic.addScenic
      // console.log(this.form)
      // submitApi(this.form).then(res => {
      //   this.$message.success('保存成功！')
      //   this.$router.go(-1)
      // })
    },
    bannerListChange (list) {
      this.form.banner = list.join('|')
    }
  }
}
</script>

<style scoped lang="scss">
.scenic-add-container {
  padding: 0 15px;
  max-width: 800px;
  margin: 0 auto;

  .form-name {
    text-align: center;
  }
}
</style>
