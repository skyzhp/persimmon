webpackJsonp([6],{

/***/ 147:
/***/ (function(module, exports, __webpack_require__) {

// style-loader: Adds some css to the DOM by adding a <style> tag

// load the styles
var content = __webpack_require__(148);
if(typeof content === 'string') content = [[module.i, content, '']];
// add the styles to the DOM
var update = __webpack_require__(5)(content, {});
if(content.locals) module.exports = content.locals;
// Hot Module Replacement
if(false) {
	// When the styles change, update the <style> tags
	if(!content.locals) {
		module.hot.accept("!!../../../../../node_modules/css-loader/index.js!../../../../../node_modules/vue-loader/lib/style-compiler/index.js?{\"id\":\"data-v-0fc6c467\",\"scoped\":false,\"hasInlineConfig\":false}!../../../../../node_modules/less-loader/index.js!../../../../../node_modules/vue-loader/lib/selector.js?type=styles&index=0!./post-list.vue", function() {
			var newContent = require("!!../../../../../node_modules/css-loader/index.js!../../../../../node_modules/vue-loader/lib/style-compiler/index.js?{\"id\":\"data-v-0fc6c467\",\"scoped\":false,\"hasInlineConfig\":false}!../../../../../node_modules/less-loader/index.js!../../../../../node_modules/vue-loader/lib/selector.js?type=styles&index=0!./post-list.vue");
			if(typeof newContent === 'string') newContent = [[module.id, newContent, '']];
			update(newContent);
		});
	}
	// When the module is disposed, remove the <style> tags
	module.hot.dispose(function() { update(); });
}

/***/ }),

/***/ 148:
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__(4)();
// imports


// module
exports.push([module.i, "\n.pit-content {\n  margin-bottom: 50px;\n}\n.pit-action-btn,\n.myp-page {\n  margin-top: 10px;\n}\n", ""]);

// exports


/***/ }),

/***/ 149:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
Object.defineProperty(__webpack_exports__, "__esModule", { value: true });
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__libs_util__ = __webpack_require__(6);
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//
//



/* harmony default export */ __webpack_exports__["default"] = ({
    data() {
        return {
            tableColumns: [{
                type: 'selection',
                width: 60,
                align: 'center'
            }, {
                title: '标题',
                key: 'title'
            }, {
                title: '分类',
                key: 'categories',
                render: (h, params) => {
                    return h('span', params.row.categories.category_name);
                }
            }, {
                title: '标签',
                key: 'tags',
                render: (h, params) => {
                    let tagsArr = [];
                    let tags = params.row.tags;
                    for (let i = 0, len = tags.length; i < len; i++) {
                        tagsArr[i] = h('Tag', {}, tags[i].tags_name);
                    }
                    return h('div', tagsArr);
                }
            }, {
                title: '日期',
                key: 'created_at'
            }, {
                title: '操作',
                key: 'action',
                width: 150,
                align: 'center',
                render: (h, params) => {
                    return h('div', [h('Button', {
                        props: {
                            type: 'primary',
                            size: 'small'
                        },
                        style: {
                            marginRight: '5px'
                        },
                        on: {
                            click: () => {
                                this.postEditor(params.row);
                            }
                        }
                    }, '编辑'), h('Button', {
                        props: {
                            type: 'error',
                            size: 'small'
                        },
                        on: {
                            click: () => {
                                this.handleDistory('one', params.row);
                            }
                        }
                    }, '删除')]);
                }
            }],
            listData: [],
            category_id: '',
            categorys: [],
            currentPage: 1,
            total: 0,
            pageSize: 20,
            listLoading: false,
            checkedAll: [],
            sizeOpts: [10, 20, 30, 40, 50, 60, 70, 80, 90, 100],
            q: ''
        };
    },
    methods: {
        filterCategory(value) {
            this.getData();
        },
        searchBtn(event) {
            this.getData();
        },
        getData() {
            let that = this;
            that.listLoading = true;
            let query = {
                rows: that.pageSize,
                category_id: that.category_id,
                q: that.q,
                page: that.currentPage
            };

            __WEBPACK_IMPORTED_MODULE_0__libs_util__["a" /* default */].ajax.get('/posts', { params: query }).then(function (response) {
                let res = response.data;
                if (res != false) {
                    that.listData = res.data;
                    that.total = res.total;
                    that.currentPage = res.current_page;
                    that.listLoading = false;
                } else {
                    this.$Notice.warning({
                        title: '数据获取失败',
                        desc: ''
                    });
                }
            }).catch(function (error) {
                console.log(error);
            });
        },
        handleSizeChange(val) {
            //console.log(`每页 ${val} 条`);
            this.pageSize = val;
            this.getData();
        },
        handleCurrentChange(val) {
            this.currentPage = val;
            this.getData();
            //console.log(`当前页: ${val}`);
        },
        handleDistory(type, row) {
            let that = this,
                idsParam = {};
            switch (type) {
                case 'one':
                    if (parseInt(row.id) <= 0) {
                        this.$Notice.warning({
                            title: '请选择需要删除的数据',
                            desc: ''
                        });
                        return false;
                    }
                    idsParam = { ids: [row.id] };
                    break;
                case 'multi':
                    var ids = __WEBPACK_IMPORTED_MODULE_0__libs_util__["a" /* default */].getIdByArr(that.checkedAll);
                    if (ids.length <= 0) {
                        this.$Notice.warning({
                            title: '请选择需要删除的数据',
                            desc: ''
                        });
                        return false;
                    }
                    idsParam = { ids: ids };
                    break;
                default:
                    break;
            }

            that.$Modal.confirm({
                title: '确认删除选中的记录吗?',
                content: '<p>您确认删除选中的记录吗?</p>',
                onOk: () => {
                    that.listLoading = true;
                    __WEBPACK_IMPORTED_MODULE_0__libs_util__["a" /* default */].ajax.delete('/posts/destroy', { data: idsParam }).then(function (response) {
                        that.listLoading = false;
                        let res = response.data;
                        that.$Notice.open({
                            title: res.status == 'success' ? '删除成功' : '删除失败',
                            desc: '',
                            duration: 2
                        });

                        if (type == 'one') {
                            __WEBPACK_IMPORTED_MODULE_0__libs_util__["a" /* default */].removeByValue(that.listData, row.id);
                        } else {
                            for (let index in that.checkedAll) {
                                __WEBPACK_IMPORTED_MODULE_0__libs_util__["a" /* default */].removeByValue(that.listData, that.checkedAll[index].id);
                            }
                        }
                    }).catch(function (error) {
                        console.log(error);
                    });
                },
                onCancel: () => {
                    that.listLoading = false;
                }
            });
        },
        selectRow(row) {
            this.checkedAll = row;
        },
        selectAll(selection) {
            this.checkedAll = selection;
        },
        getCategorys() {
            let that = this;
            __WEBPACK_IMPORTED_MODULE_0__libs_util__["a" /* default */].ajax.get('/categorys', {
                params: {
                    rows: 999
                }
            }).then(function (response) {
                let res = response.data;
                if (res != false) {
                    res.data.splice(0, 0, { id: 0, category_name: '顶级分类', hidden: true, category_parent: 0 });
                    that.categorys = res.data;
                } else {
                    this.$Notice.error({
                        title: '数据获取失败',
                        desc: ''
                    });
                }
            }).catch(function (error) {
                console.log(error);
            });
        },
        setTopCategorys() {
            var categorys = this.listData.concat();
            categorys.splice(0, 0, { id: 0, category_name: '顶级分类', hidden: true, category_parent: 0 });
            this.categorys = categorys;
        },
        postEditor(row) {
            let path = '/posts/edit/' + row.id;
            this.$router.push({ path: path });
        },
        addPost() {
            let path = "post-add";
            this.$router.push({ path: path });
        }
    },
    mounted() {
        this.getCategorys();
        this.getData();
    }
});

/***/ }),

/***/ 150:
/***/ (function(module, exports, __webpack_require__) {

module.exports = { render: function () {
    var _vm = this;var _h = _vm.$createElement;var _c = _vm._self._c || _h;
    return _c('div', {
      staticClass: "pit-content"
    }, [_c('div', {
      staticClass: "pit-action-btn"
    }, [_c('Button', {
      staticClass: "myp-search-item",
      attrs: {
        "type": "primary"
      },
      on: {
        "click": function ($event) {
          _vm.handleDistory('multi', {});
        }
      }
    }, [_vm._v("删除")]), _vm._v(" "), _c('Button', {
      staticClass: "myp-search-item",
      attrs: {
        "type": "primary",
        "icon": "android-add"
      },
      on: {
        "click": function ($event) {
          _vm.addPost();
        }
      }
    }, [_vm._v("新增")]), _vm._v(" "), _c('Select', {
      staticClass: "myp-search-item",
      staticStyle: {
        "width": "200px"
      },
      on: {
        "change": _vm.filterCategory
      },
      model: {
        value: _vm.category_id,
        callback: function ($$v) {
          _vm.category_id = $$v;
        },
        expression: "category_id"
      }
    }, _vm._l(_vm.categorys, function (item) {
      return _c('Option', {
        key: item.id,
        attrs: {
          "value": item.id
        }
      }, [_vm._v(_vm._s(item.category_name))]);
    })), _vm._v(" "), _c('Input', {
      staticClass: "myp-search-item",
      staticStyle: {
        "width": "200px"
      },
      attrs: {
        "placeholder": "请输入内容"
      },
      model: {
        value: _vm.q,
        callback: function ($$v) {
          _vm.q = $$v;
        },
        expression: "q"
      }
    }, [_c('Button', {
      attrs: {
        "slot": "append",
        "icon": "ios-search"
      },
      on: {
        "click": _vm.searchBtn
      },
      slot: "append"
    })], 1)], 1), _vm._v(" "), _c('div', {
      staticClass: "data-list"
    }, [_c('Table', {
      attrs: {
        "loading": _vm.listLoading,
        "columns": _vm.tableColumns,
        "data": _vm.listData,
        "stripe": ""
      },
      on: {
        "on-select": _vm.selectRow,
        "on-select-all": _vm.selectAll
      }
    })], 1), _vm._v(" "), _c('div', {
      staticStyle: {
        "margin": "10px",
        "overflow": "hidden"
      }
    }, [_c('div', {
      staticStyle: {
        "float": "right"
      }
    }, [_c('Page', {
      staticClass: "myp-page",
      attrs: {
        "total": _vm.total,
        "page-size": _vm.pageSize,
        "current": _vm.currentPage,
        "page-size-opts": _vm.sizeOpts,
        "show-elevator": true,
        "show-sizer": true
      },
      on: {
        "on-page-size-change": _vm.handleSizeChange,
        "on-change": _vm.handleCurrentChange
      }
    })], 1)])]);
  }, staticRenderFns: [] };
module.exports.render._withStripped = true;
if (false) {
  module.hot.accept();
  if (module.hot.data) {
    require("vue-loader/node_modules/vue-hot-reload-api").rerender("data-v-0fc6c467", module.exports);
  }
}

/***/ }),

/***/ 89:
/***/ (function(module, exports, __webpack_require__) {


/* styles */
__webpack_require__(147)

var Component = __webpack_require__(2)(
  /* script */
  __webpack_require__(149),
  /* template */
  __webpack_require__(150),
  /* scopeId */
  null,
  /* cssModules */
  null
)
Component.options.__file = "/Users/mrcong/GoLang/src/github.com/cong5/persimmon/assets/backend/src/views/posts/post-list.vue"
if (Component.esModule && Object.keys(Component.esModule).some(function (key) {return key !== "default" && key !== "__esModule"})) {console.error("named exports are not supported in *.vue files.")}
if (Component.options.functional) {console.error("[vue-loader] post-list.vue: functional components are not supported with templates, they should use render functions.")}

/* hot reload */
if (false) {(function () {
  var hotAPI = require("vue-loader/node_modules/vue-hot-reload-api")
  hotAPI.install(require("vue"), false)
  if (!hotAPI.compatible) return
  module.hot.accept()
  if (!module.hot.data) {
    hotAPI.createRecord("data-v-0fc6c467", Component.options)
  } else {
    hotAPI.reload("data-v-0fc6c467", Component.options)
  }
})()}

module.exports = Component.exports


/***/ })

});
//# sourceMappingURL=6.chunk.js.map