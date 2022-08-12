// scenic.js
// 获取应用实例
const app = getApp()

Page({
  data: {
    msg: '你好',
    defaultData: {
      title: '地图'
    },
    markers: [],
    customCalloutMarkerIds: [4],
  },
  // latitude="23.560987"
  // longitude="116.829499"
  // 事件处理函数
  onLoad() {
    const markers = [
      {
        id: 4,
        latitude: 23.561987,
        longitude: 116.830499,
        iconPath: '/assets/image/location.png',
        name: '地点1',
        customCallout: {
          anchorY: 0,
          anchorX: 20,
          display: 'ALWAYS',
        }
      }
    ]
    this.setData({
      markers
    })
  },
  onReady() {
    this.mapCtx = wx.createMapContext('easyMap')
  }
  ,
  switchTab(e) {
    wx.switchTab({
      url: e.detail.url
    })
  },
  clickMarkers(e) {
    console.log(e)
  },
  clickCallOut(e){
    console.log(e)
  },
  clickLabel (e) {
    console.log(e)
  }
})
