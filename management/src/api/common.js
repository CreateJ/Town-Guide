import request from '@/utils/request'

function uploadFile(data) {
  return request({
    url: '/utils/uploadFile',
    method: 'post',
    data: data
  })
}

function getPic(data) {
  return request({
    url: `/utils/getPic/${data.pic_name}`,
    method: 'get'
    // params: data
  })
}

function getVideo(data) {
  return request({
    url: '/utils/getVideo',
    method: 'get',
    params: data
  })
}

export default {
  uploadFile,
  getPic,
  getVideo
}
