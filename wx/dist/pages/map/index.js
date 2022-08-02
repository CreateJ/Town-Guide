/******/ (() => { // webpackBootstrap
var __webpack_exports__ = {};
// index.js
// 获取应用实例
const app = getApp();
Page({
  data: {
    msg: '你好'
  },

  // 事件处理函数
  onLoad() {// wx.request({
    //   url: 'https://guide.time-traveler.cn/user/1',
    //   method: 'GET',
    //   success(res) {
    //     console.log(res)
    //   }
    // })
  },

  onReady() {
    this.mapCtx = wx.createMapContext('easyMap');
    console.log(this.mapCtx);
    this.mapCtx.addGroundOverlay({
      id: 1,
      src: 'https://guide.time-traveler.cn/utils/getPic/1659283570猫猫3.jpeg',
      bounds: {
        southwest: {
          longitude: 116.802075,
          latitude: 23.541833
        },
        northeast: {
          longitude: 116.863016,
          latitude: 23.587412
        }
      },
      visible: true,
      zIndex: 1000,
      success: res => {
        console.log(res, 1111);
        console.log(res, 1111);
      },
      fail: err => {
        console.log(err);
      }
    });
  }

});
/******/ })()
;