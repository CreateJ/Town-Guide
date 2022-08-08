import request from '@/utils/request'

function getCategoryList() {
  return request({
    url: '/category/getAll',
    method: 'get'
  })
}

function addCategory(data) {
  return request({
    url: '/category/add',
    method: 'post',
    data
  })
}

function editCategory(data) {
  return request({
    url: '/category/edit',
    method: 'post',
    data
  })
}

function deleteCategory(data) {
  return request({
    url: '/category/delete/' + data.id,
    method: 'delete'
  })
}

export default {
  getCategoryList,
  addCategory,
  deleteCategory,
  editCategory
}
