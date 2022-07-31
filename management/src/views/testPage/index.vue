<template>
  <div>
  </div>
</template>

<script>
import api from '@/api'

export default {
  name: 'TextPage',
  data() {
    return {
    }
  },
  mounted() {
    api.scenic.getScenicList().then(res => {
      console.log(res)
    })
  },
  methods: {}
}
</script>

<style scoped>

</style>
