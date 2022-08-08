/******/ (() => { // webpackBootstrap
var __webpack_exports__ = {};
// scenic.js
// 获取应用实例
const app = getApp();
Page({
  data: {
    msg: '你好',
    defaultData: {
      title: '地图'
    }
  },

  // 事件处理函数
  onLoad() {},

  onReady() {},

  switchTab(e) {
    wx.switchTab({
      url: e.detail.url
    });
  }

});
/******/ })()
;