const app = getApp()
Component({
  properties: {
    // defaultData（父页面传递的数据-就是引用组件的页面）
    defaultData: {
      type: Object,
      value: {
        title: "我是默认标题"
      },
      observer: function (newVal, oldVal) {
      }
    }
  },
  data: {
    navBarHeight: app.globalData.navBarHeight,
    menuRight: app.globalData.menuRight,
    menuTop: app.globalData.menuTop,
    menuHeight: app.globalData.menuHeight,
    active: null,
    button1Bg: '../../assets/image/point.png',
    button2Bg: '../../assets/image/list.png',
    animationData: {},
    myAnimation: null,
    categoryList: [],
    activeCategory: null
  },

  methods: {
    clickCategoryItem({ currentTarget }) {
      this.setData({
        activeCategory: currentTarget.dataset.type
      })
    },
    clickCategory() {
      if (this.data.active === 1) {
        this.myAnimation.height(app.globalData.navBarHeight + 'px').step()
        this.setData({
          animationData: this.myAnimation.export(),
          active: null,
          button1Bg: '../../assets/image/point.png'
        })
      } else {
        this.changeTo1()
        this.triggerEvent('switch', { url: '/pages/map/index' })
      }
    },
    clickList() {
      if (this.data.active !== 2) {
        this.changeTo2()
        this.triggerEvent('switch', { url: '/pages/visits/index' })
      }
    },
    changeTo1() {
      this.myAnimation.height((app.globalData.navBarHeight + 55) + 'px').step()
      this.setData({
        active: 1,
        animationData: this.myAnimation.export(),
        button1Bg: '../../assets/image/point-active.png',
        button2Bg: '../../assets/image/list.png'
      })
    },
    changeTo2() {
      this.myAnimation.height(app.globalData.navBarHeight + 'px').step()
      this.setData({
        active: 2,
        animationData: this.myAnimation.export(),
        button2Bg: '../../assets/image/list-active.png',
        button1Bg: '../../assets/image/point.png',
      })
    }
  },
  lifetimes: {
    attached: function () {
      this.myAnimation = wx.createAnimation({
        duration: 500,
        timingFunction: 'ease'
      })

      wx.request({
        url: 'https://guide.time-traveler.cn/category/getAll',
        method: 'GET',
        success: (res) => {
          let list = res.data.data.map(item => {
            return {
              name: item.name,
              icon: 'https://guide.time-traveler.cn/utils/getPic/'+ item.icon,
              iconActive: 'https://guide.time-traveler.cn/utils/getPic/'+ item.icon_active,
              id: item.id
            }
          })
          this.setData({
            categoryList: list
          })
        }
      })
    },
  },
  pageLifetimes: {
    // 组件所在页面的生命周期函数
    show: function () {
      const active = getCurrentPages()[0].route.includes('map') ? 1 : 2
      this['changeTo' + active]()
      this.setData({
        active
      })
    },
    hide: function () {
    }
  },
})
