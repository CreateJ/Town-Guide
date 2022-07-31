import request from '@/utils/request'

function getScenicList() {
  return request({
    url: '/scenic/getAll',
    method: 'get'
  })
}

export default {
  getScenicList
}
