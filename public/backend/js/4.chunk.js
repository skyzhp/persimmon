webpackJsonp([4],{

/***/ 246:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
Object.defineProperty(__webpack_exports__, "__esModule", { value: true });
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_post_trash_vue__ = __webpack_require__(266);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_post_trash_vue___default = __webpack_require__.n(__WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_post_trash_vue__);
/* harmony namespace reexport (unknown) */ for(var __WEBPACK_IMPORT_KEY__ in __WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_post_trash_vue__) if(__WEBPACK_IMPORT_KEY__ !== 'default') (function(key) { __webpack_require__.d(__webpack_exports__, key, function() { return __WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_post_trash_vue__[key]; }) }(__WEBPACK_IMPORT_KEY__));
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__babel_loader_node_modules_vue_loader_lib_template_compiler_index_id_data_v_72045d5c_hasScoped_false_buble_transforms_node_modules_vue_loader_lib_selector_type_template_index_0_post_trash_vue__ = __webpack_require__(320);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__babel_loader_node_modules_vue_loader_lib_template_compiler_index_id_data_v_72045d5c_hasScoped_false_buble_transforms_node_modules_vue_loader_lib_selector_type_template_index_0_post_trash_vue___default = __webpack_require__.n(__WEBPACK_IMPORTED_MODULE_1__babel_loader_node_modules_vue_loader_lib_template_compiler_index_id_data_v_72045d5c_hasScoped_false_buble_transforms_node_modules_vue_loader_lib_selector_type_template_index_0_post_trash_vue__);
var disposed = false
function injectStyle (ssrContext) {
  if (disposed) return
  __webpack_require__(318)
}
var normalizeComponent = __webpack_require__(1)
/* script */


/* template */

/* template functional */
var __vue_template_functional__ = false
/* styles */
var __vue_styles__ = injectStyle
/* scopeId */
var __vue_scopeId__ = null
/* moduleIdentifier (server only) */
var __vue_module_identifier__ = null
var Component = normalizeComponent(
  __WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_post_trash_vue___default.a,
  __WEBPACK_IMPORTED_MODULE_1__babel_loader_node_modules_vue_loader_lib_template_compiler_index_id_data_v_72045d5c_hasScoped_false_buble_transforms_node_modules_vue_loader_lib_selector_type_template_index_0_post_trash_vue___default.a,
  __vue_template_functional__,
  __vue_styles__,
  __vue_scopeId__,
  __vue_module_identifier__
)
Component.options.__file = "assets/backend/src/views/posts/post-trash.vue"

/* hot reload */
if (false) {(function () {
  var hotAPI = require("vue-hot-reload-api")
  hotAPI.install(require("vue"), false)
  if (!hotAPI.compatible) return
  module.hot.accept()
  if (!module.hot.data) {
    hotAPI.createRecord("data-v-72045d5c", Component.options)
  } else {
    hotAPI.reload("data-v-72045d5c", Component.options)
  }
  module.hot.dispose(function (data) {
    disposed = true
  })
})()}

/* harmony default export */ __webpack_exports__["default"] = (Component.exports);


/***/ }),

/***/ 266:
/***/ (function(module, exports, __webpack_require__) {

"use strict";


Object.defineProperty(exports, "__esModule", {
    value: true
});

var _util = __webpack_require__(18);

var _util2 = _interopRequireDefault(_util);

function _interopRequireDefault(obj) { return obj && obj.__esModule ? obj : { default: obj }; }

exports.default = {
    data: function data() {
        var _this = this;

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
                render: function render(h, params) {
                    return h('span', params.row.categories.category_name);
                }
            }, {
                title: '标签',
                key: 'tags',
                render: function render(h, params) {
                    var tagsArr = [];
                    var tags = params.row.tags;
                    for (var i = 0, len = tags.length; i < len; i++) {
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
                render: function render(h, params) {
                    return h('div', [h('Button', {
                        props: {
                            type: 'primary',
                            size: 'small'
                        },
                        style: {
                            marginRight: '5px'
                        },
                        on: {
                            click: function click() {
                                _this.handleRestore(params.row);
                            }
                        }
                    }, '还原'), h('Button', {
                        props: {
                            type: 'error',
                            size: 'small'
                        },
                        on: {
                            click: function click() {
                                _this.handleDistory('one', params.row);
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
        filterCategory: function filterCategory(value) {
            this.getData();
        },
        searchBtn: function searchBtn(event) {
            this.getData();
        },
        getData: function getData() {
            var that = this;
            that.listLoading = true;
            var query = {
                rows: that.pageSize,
                category_id: that.category_id,
                q: that.q
            };

            _util2.default.ajax.get('/trash', { params: query }).then(function (response) {
                var res = response.data;
                if (res != false) {
                    that.listData = res.data;
                    that.total = res.total;
                    that.currentPage = res.current_page;
                    that.listLoading = false;
                } else {
                    that.$message({
                        message: '数据获取失败',
                        type: 'error',
                        duration: 3 * 1000
                    });
                }
            }).catch(function (error) {
                console.log(error);
            });
        },
        handleSizeChange: function handleSizeChange(val) {
            this.pageSize = val;
            this.getData();
        },
        handleCurrentChange: function handleCurrentChange(val) {
            this.currentPage = val;
        },

        handleRestore: function handleRestore(row) {
            var that = this;

            that.$Modal.confirm({
                title: '确认恢复选中的文章吗?',
                content: '<p>确认恢复选中的文章吗?</p>',
                onOk: function onOk() {
                    that.listLoading = true;
                    _util2.default.ajax.put('/trash/update', { 'ids': [row.id] }).then(function (response) {
                        that.listLoading = false;
                        var res = response.data;
                        that.$Notice.open({
                            title: res.status == 'success' ? '恢复成功' : '恢复失败',
                            desc: ''
                        });
                        _util2.default.removeByValue(that.listData, row.id);
                    }).catch(function (error) {
                        console.log(error);
                    });
                },
                onCancel: function onCancel() {
                    that.listLoading = false;
                }
            });
        },
        handleDistory: function handleDistory(type, row) {
            var that = this,
                idsParam = {};
            switch (type) {
                case 'one':
                    if (parseInt(row.id) <= 0) {
                        that.$Notice.open({
                            title: '请选择需要删除的数据',
                            desc: ''
                        });
                        return false;
                    }
                    idsParam = { ids: [row.id] };
                    break;
                case 'multi':
                    var ids = _util2.default.getIdByArr(that.checkedAll);
                    if (ids.length <= 0) {
                        that.$Notice.open({
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
                onOk: function onOk() {
                    that.listLoading = true;
                    _util2.default.ajax.delete('/trash/destroy', { data: idsParam }).then(function (response) {
                        that.listLoading = false;
                        var res = response.data;
                        that.$Notice.open({
                            title: res.status == 'success' ? '删除成功' : '删除失败',
                            desc: ''
                        });
                        if (type == 'one') {
                            _util2.default.removeByValue(that.listData, row.id);
                        } else {
                            for (var index in that.checkedAll) {
                                _util2.default.removeByValue(that.listData, that.checkedAll[index].id);
                            }
                        }
                    }).catch(function (error) {
                        console.log(error);
                    });
                },
                onCancel: function onCancel() {
                    that.listLoading = false;
                }
            });
        },
        selectRow: function selectRow(row) {
            this.checkedAll = row;
        },
        selectAll: function selectAll(selection) {
            this.checkedAll = selection;
        },

        getCategorys: function getCategorys() {
            var that = this;
            _util2.default.ajax.get('/categorys', {
                params: {
                    rows: 999
                }
            }).then(function (response) {
                var res = response.data;
                if (res != false) {
                    res.data.splice(0, 0, { id: 0, category_name: '顶级分类', hidden: true, category_parent: 0 });
                    that.categorys = res.data;
                } else {
                    that.$Notice.open({
                        title: '数据获取失败',
                        desc: ''
                    });
                }
            }).catch(function (error) {
                console.log(error);
            });
        },
        setTopCategorys: function setTopCategorys() {
            var categorys = this.listData.concat();
            categorys.splice(0, 0, { id: 0, category_name: '顶级分类', hidden: true, category_parent: 0 });
            this.categorys = categorys;
        }
    },
    mounted: function mounted() {
        this.getCategorys();
        this.getData();
    }
};

/***/ }),

/***/ 318:
/***/ (function(module, exports, __webpack_require__) {

// style-loader: Adds some css to the DOM by adding a <style> tag

// load the styles
var content = __webpack_require__(319);
if(typeof content === 'string') content = [[module.i, content, '']];
if(content.locals) module.exports = content.locals;
// add the styles to the DOM
var update = __webpack_require__(17)("c8ff3ff8", content, false, {});
// Hot Module Replacement
if(false) {
 // When the styles change, update the <style> tags
 if(!content.locals) {
   module.hot.accept("!!../../../../../node_modules/css-loader/index.js!../../../../../node_modules/vue-loader/lib/style-compiler/index.js?{\"vue\":true,\"id\":\"data-v-72045d5c\",\"scoped\":false,\"hasInlineConfig\":false}!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/selector.js?type=styles&index=0!./post-trash.vue", function() {
     var newContent = require("!!../../../../../node_modules/css-loader/index.js!../../../../../node_modules/vue-loader/lib/style-compiler/index.js?{\"vue\":true,\"id\":\"data-v-72045d5c\",\"scoped\":false,\"hasInlineConfig\":false}!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/selector.js?type=styles&index=0!./post-trash.vue");
     if(typeof newContent === 'string') newContent = [[module.id, newContent, '']];
     update(newContent);
   });
 }
 // When the module is disposed, remove the <style> tags
 module.hot.dispose(function() { update(); });
}

/***/ }),

/***/ 319:
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__(16)(false);
// imports


// module
exports.push([module.i, "", ""]);

// exports


/***/ }),

/***/ 320:
/***/ (function(module, exports, __webpack_require__) {

"use strict";


Object.defineProperty(exports, "__esModule", {
  value: true
});
var render = function render() {
  var _vm = this;
  var _h = _vm.$createElement;
  var _c = _vm._self._c || _h;
  return _c("div", { staticClass: "pit-content" }, [_c("div", { staticClass: "pit-action-btn" }, [_c("Button", {
    staticClass: "myp-search-item",
    attrs: { type: "primary", icon: "delete" },
    on: {
      click: function click($event) {
        _vm.handleDistory("multi", {});
      }
    }
  }, [_vm._v("删除")]), _vm._v(" "), _c("Select", {
    staticClass: "myp-search-item",
    staticStyle: { width: "200px" },
    attrs: { clearable: "", placeholder: "请选择" },
    on: { change: _vm.filterCategory },
    model: {
      value: _vm.category_id,
      callback: function callback($$v) {
        _vm.category_id = $$v;
      },
      expression: "category_id"
    }
  }, _vm._l(_vm.categorys, function (item) {
    return _c("Option", { key: item.id, attrs: { value: item.id } }, [_vm._v(_vm._s(item.category_name))]);
  })), _vm._v(" "), _c("Input", {
    staticClass: "myp-search-item",
    staticStyle: { width: "200px" },
    attrs: { placeholder: "请输入内容", icon: "search" },
    model: {
      value: _vm.q,
      callback: function callback($$v) {
        _vm.q = $$v;
      },
      expression: "q"
    }
  }, [_c("Button", {
    attrs: { slot: "append", icon: "ios-search" },
    on: { click: _vm.searchBtn },
    slot: "append"
  })], 1)], 1), _vm._v(" "), _c("div", { staticClass: "data-list" }, [_c("Table", {
    attrs: {
      loading: _vm.listLoading,
      columns: _vm.tableColumns,
      data: _vm.listData,
      stripe: ""
    },
    on: { "on-select": _vm.selectRow, "on-select-all": _vm.selectAll }
  })], 1), _vm._v(" "), _c("div", { staticStyle: { margin: "10px", overflow: "hidden" } }, [_c("div", { staticStyle: { float: "right" } }, [_c("Page", {
    staticClass: "myp-page",
    attrs: {
      total: _vm.total,
      "page-size": _vm.pageSize,
      current: _vm.currentPage,
      "page-size-opts": _vm.sizeOpts,
      "show-elevator": true,
      "show-sizer": true
    },
    on: {
      "on-page-size-change": _vm.handleSizeChange,
      "on-change": _vm.handleCurrentChange
    }
  })], 1)])]);
};
var staticRenderFns = [];
render._withStripped = true;
var esExports = { render: render, staticRenderFns: staticRenderFns };
exports.default = esExports;

if (false) {
  module.hot.accept();
  if (module.hot.data) {
    require("vue-hot-reload-api").rerender("data-v-72045d5c", esExports);
  }
}

/***/ })

});
//# sourceMappingURL=4.chunk.js.map