<template>
  <div class="scenic-add-container">
    <h3 class="form-name">景区信息</h3>
    <el-form ref="from" :model="form" label-width="120px">
      <el-form-item label="景区名称">
        <el-input v-model="form.name"/>
      </el-form-item>
      <el-form-item label="地点">
        <el-input v-model="form.location_desc"/>
      </el-form-item>
      <el-form-item label="描述">
        <el-input v-model="form.description"/>
      </el-form-item>
      <el-form-item label="简介">
        <el-input v-model="form.intro"/>
      </el-form-item>
      <el-form-item label="封面图">
        <el-input v-model="form.pic_name"/>
      </el-form-item>
      <el-form-item label="图标">
        <el-input v-model="form.icon"/>
      </el-form-item>
      <el-form-item label="视频">
        <el-input v-model="form.video_name"/>
      </el-form-item>
      <el-form-item label="标签">
        <el-input v-model="form.tag"/>
      </el-form-item>
      <el-form-item label="开放时间">
        <el-input v-model="form.open_time"/>
      </el-form-item>
      <el-form-item label="轮播图">
        <!--        <el-input v-model="form.banner"/>-->
        <media-upload :file-list="bannerList" @file-list-change="bannerListChange" :limit="5"></media-upload>
      </el-form-item>
      <el-form-item label="分类id">
        <el-input v-model="form.category_id"/>
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
import api from '@/api'

export default {
  name: 'ScenicAdd',
  components: { MediaUpload },
  data() {
    return {
      form: {
        name: '',
        location_desc: '',
        description: '',
        intro: '',
        pic_name: '',
        icon: '',
        video_name: '',
        tag: '',
        open_time: '',
        check_num: 0, // 打卡数量
        category_id: 1,
        banner: ''
      },
      bannerList: []
    }
  },
  mounted() {
    if (this.$route.query.id) {
      this.form = this.$route.query
      if (this.form.banner.length) {
        this.bannerList = this.form.banner.split('|').map(item => {
          return {
            name: item,
            url: `https://guide.time-traveler.cn/utils/getPic/${item}`
          }
        })
      }
    }
  },
  methods: {
    onSubmit() {
      const submitApi = this.form.id ? api.scenic.editScenic : api.scenic.addScenic
      console.log(this.form)
      submitApi(this.form).then(res => {
        this.$message.success('保存成功！')
        this.$router.go(-1)
      })
    },
    bannerListChange(list) {
      this.form.banner = list.join('|')
      console.log(this.form.banner)
    }
  }
}
</script>

<style scoped lang="scss">
.scenic-add-container {
  padding: 0 15px;

  .form-name {
    text-align: center;
  }
}
</style>
