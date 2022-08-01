import request from '@/utils/request'

function getScenicList() {
  return request({
    url: '/scenic/getAll',
    method: 'get'
  })
}

function getScenic({ id }) {
  return request({
    url: `/scenic/get/${id}`,
    method: 'get'
  })
}

function deleteScenic({ id }) {
  return request({
    url: `/scenic/delete/${id}`,
    method: 'delete'
  })
}

function addScenic(data) {
  return request({
    url: '/scenic/add',
    method: 'post',
    data: data
  })
}

function editScenic(data) {
  return request({
    url: '/scenic/edit',
    method: 'post',
    data: data
  })
}

export default {
  getScenicList,
  addScenic,
  editScenic,
  getScenic,
  deleteScenic
}
