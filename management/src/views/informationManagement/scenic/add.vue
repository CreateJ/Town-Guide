<template>
  <div class="scenic-add-container">
    <h3 class="form-name">添加景区</h3>
    <el-form ref="from" :model="form" label-width="120px">
      <el-form-item label="景区名称">
        <el-input v-model="form.name"/>
      </el-form-item>
      <el-form-item label="地点">
        <el-input v-model="form.local_desc"/>
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
        <el-input v-model="form.banner"/>
      </el-form-item>
      <el-form-item label="分类id">
        <el-input v-model="form.category_id"/>
      </el-form-item>
      <el-form-item>
        <el-button type="primary" @click="onSubmit">立即创建</el-button>
        <el-button>取消</el-button>
      </el-form-item>
    </el-form>
  </div>
</template>

<script>
export default {
  name: 'ScenicAdd',
  data() {
    return {
      form: {
        name: '景区名字',
        local_desc: '地点地点',
        description: '地点描述唱的描述',
        intro: '段简介',
        pic_name: '图片名称',
        icon: 'icon图片名称',
        video_name: '视频名称',
        tag: 'tag标签',
        open_time: '开放时间',
        check_num: 0,
        category_id: 1,
        banner: 'banner图片名称,多个通过|分割'
      }
    }
  },
  methods: {
    onSubmit() {
      console.log(this.form)
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
