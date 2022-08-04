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
    categoryList: [
      {
        type: 'wharf',
        name: '码头',
        icon: '../../assets/image/categories/1-1.png',
        iconActive: '../../assets/image/categories/1-2.png'
      },
      {
        type: 'exhibition', // 展览馆
        name: '展览馆',
        icon: '../../assets/image/categories/2-1.png',
        iconActive: '../../assets/image/categories/2-2.png'
      },
      {
        type: 'ancestral',
        name: '祠堂',
        icon: '../../assets/image/categories/3-1.png',
        iconActive: '../../assets/image/categories/3-2.png'
      }
    ],
    activeCategory: null
  },
  attached: function () {
  },
  methods: {
    clickCategoryItem ({currentTarget}) {
      console.log(currentTarget.dataset.type)
      this.setData({
        activeCategory: currentTarget.dataset.type
      })
    },
    clickCategory() {
      if (this.data.active === 1) {
        this.setData({
          active: null,
          button1Bg: '../../assets/image/point.png'
        })
      } else {
        this.setData({
          active: 1,
          button1Bg: '../../assets/image/point-active.png',
          button2Bg: '../../assets/image/list.png',
        })
      }
    },
    clickList() {
      if (this.data.active === 2) {
        this.setData({
          active: null,
          button2Bg: '../../assets/image/list.png',
          button1Bg: '../../assets/image/point.png'
        })
      } else {
        this.setData({
          active: 2,
          button2Bg: '../../assets/image/list-active.png',
          button1Bg: '../../assets/image/point.png',
        })
      }
    }
  },
  lifetimes: {},
})
