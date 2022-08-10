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
      <el-form-item label="坐标">
        <el-input v-model="form.location"/>
      </el-form-item>
      <el-form-item label="描述">
        <el-input v-model="form.description" type="textarea" :autosize="{ minRows: 4, maxRows: 8}"/>
      </el-form-item>
      <el-form-item label="简介">
        <el-input v-model="form.intro" type="textarea" :autosize="{ minRows: 4, maxRows: 8}"/>
      </el-form-item>
      <el-form-item label="详情图片">
        <media-upload :file-list="picList" :limit="1" @file-list-change="picListChange"/>
      </el-form-item>
      <el-form-item label="封面图">
        <media-upload :file-list="iconList" :limit="1" @file-list-change="iconListChange"/>
      </el-form-item>
      <el-form-item label="印章">
        <media-upload :file-list="clockIconList" :limit="1" @file-list-change="clockIconListChange"/>
      </el-form-item>
      <el-form-item label="视频">
        <media-upload
          :file-list="videoList"
          file-type="media"
          :limit="5"
          @file-list-change="videoListChange"
        />
      </el-form-item>
      <el-form-item label="音频">
        <media-upload
          :file-list="audioList"
          file-type="media"
          :limit="5"
          @file-list-change="audioListChange"
        />
      </el-form-item>
      <el-form-item label="标签">
        <el-input v-model="form.tag"/>
      </el-form-item>
      <el-form-item label="开放时间">
        <el-input v-model="form.open_time"/>
      </el-form-item>
      <el-form-item label="轮播图">
        <media-upload :file-list="bannerList" :limit="5" @file-list-change="bannerListChange"/>
      </el-form-item>
      <el-form-item label="分类id">
        <el-select v-model="form.category_id">
          <el-option
            v-for="item in categoryOptions"
            :key="item.id"
            :label="item.name"
            :value="item.id"
          />
        </el-select>
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
        location: '',
        location_desc: '',
        description: '',
        intro: '',
        pic_name: '',
        icon: '',
        video_name: '',
        tag: '',
        open_time: '',
        clock_num: 0, // 打卡数量
        category_id: 0,
        banner: '',
        clock_icon: ''
      },
      bannerList: [],
      categoryOptions: [],
      picList: [],
      iconList: [],
      videoList: [],
      audioList: [],
      clockIconList: []
    }
  },
  mounted() {
    if (this.$route.query.id) {
      this.form = this.$route.query
      if (this.form.banner.length) {
        this.bannerList = this.form.banner.split('|').map(item => {
          return {
            name: item,
            url: `https://guide.time-traveler.cn:4443/utils/getPic/${item}`
          }
        })
      }
      if (this.form.pic_name.length) {
        this.picList = this.form.pic_name.split('|').map(item => {
          return {
            name: item,
            url: `https://guide.time-traveler.cn:4443/utils/getPic/${item}`
          }
        })
      }
      if (this.form.icon.length) {
        this.iconList = this.form.icon.split('|').map(item => {
          return {
            name: item,
            url: `https://guide.time-traveler.cn:4443/utils/getPic/${item}`
          }
        })
      }

      if (this.form.clock_icon.length) {
        this.clockIconList = this.form.clock_icon.split('|').map(item => {
          return {
            name: item,
            url: `https://guide.time-traveler.cn:4443/utils/getPic/${item}`
          }
        })
        console.log(this.clockIconList)
      }
    }
    api.category.getCategoryList().then(res => {
      console.log(res)
      this.categoryOptions = res.data
      this.categoryOptions.unshift({ id: 0, name: '其他' })
    })
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
    },
    picListChange(list) {
      this.form.pic_name = list.join('|')
    },
    iconListChange(list) {
      this.form.icon = list.join('|')
    },
    videoListChange(list) {
      console.log(list)
    },
    audioListChange(list) {
      console.log(list)
      this.form.audio_name = list.join('|')
    },
    clockIconListChange(list) {
      console.log(list)
      this.form.clock_icon = list.join('|')
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
