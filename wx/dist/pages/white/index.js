/******/ (() => { // webpackBootstrap
var __webpack_exports__ = {};
// scenic.js
// 获取应用实例
const app = getApp();
Page({
  data: {
    navBarHeight: app.globalData.navBarHeight
  },

  goBack() {
    wx.navigateBack();
  }

});
/******/ })()
;