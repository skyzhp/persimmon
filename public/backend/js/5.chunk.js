webpackJsonp([5],{

/***/ 247:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
Object.defineProperty(__webpack_exports__, "__esModule", { value: true });
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_post_tags_vue__ = __webpack_require__(267);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_post_tags_vue___default = __webpack_require__.n(__WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_post_tags_vue__);
/* harmony namespace reexport (unknown) */ for(var __WEBPACK_IMPORT_KEY__ in __WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_post_tags_vue__) if(__WEBPACK_IMPORT_KEY__ !== 'default') (function(key) { __webpack_require__.d(__webpack_exports__, key, function() { return __WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_post_tags_vue__[key]; }) }(__WEBPACK_IMPORT_KEY__));
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__babel_loader_node_modules_vue_loader_lib_template_compiler_index_id_data_v_16ea7d25_hasScoped_false_buble_transforms_node_modules_vue_loader_lib_selector_type_template_index_0_post_tags_vue__ = __webpack_require__(323);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1__babel_loader_node_modules_vue_loader_lib_template_compiler_index_id_data_v_16ea7d25_hasScoped_false_buble_transforms_node_modules_vue_loader_lib_selector_type_template_index_0_post_tags_vue___default = __webpack_require__.n(__WEBPACK_IMPORTED_MODULE_1__babel_loader_node_modules_vue_loader_lib_template_compiler_index_id_data_v_16ea7d25_hasScoped_false_buble_transforms_node_modules_vue_loader_lib_selector_type_template_index_0_post_tags_vue__);
var disposed = false
function injectStyle (ssrContext) {
  if (disposed) return
  __webpack_require__(321)
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
  __WEBPACK_IMPORTED_MODULE_0__babel_loader_node_modules_vue_loader_lib_selector_type_script_index_0_post_tags_vue___default.a,
  __WEBPACK_IMPORTED_MODULE_1__babel_loader_node_modules_vue_loader_lib_template_compiler_index_id_data_v_16ea7d25_hasScoped_false_buble_transforms_node_modules_vue_loader_lib_selector_type_template_index_0_post_tags_vue___default.a,
  __vue_template_functional__,
  __vue_styles__,
  __vue_scopeId__,
  __vue_module_identifier__
)
Component.options.__file = "assets/backend/src/views/posts/post-tags.vue"

/* hot reload */
if (false) {(function () {
  var hotAPI = require("vue-hot-reload-api")
  hotAPI.install(require("vue"), false)
  if (!hotAPI.compatible) return
  module.hot.accept()
  if (!module.hot.data) {
    hotAPI.createRecord("data-v-16ea7d25", Component.options)
  } else {
    hotAPI.reload("data-v-16ea7d25", Component.options)
  }
  module.hot.dispose(function (data) {
    disposed = true
  })
})()}

/* harmony default export */ __webpack_exports__["default"] = (Component.exports);


/***/ }),

/***/ 267:
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
            listData: [],
            categorys: [],
            currentPage: 1,
            total: 0,
            pageSize: 20,
            myForm: {
                id: 0,
                tags_name: '',
                tags_flag: ''
            },
            myRules: {
                tags_name: [{ required: true, type: "string", message: '请填写标签名称', trigger: 'blur' }],
                tags_flag: [{ required: true, type: "string", message: '请填写标签别名', trigger: 'blur' }]
            },
            editFormVisible: false,
            editFormLoading: false,
            listLoading: false,
            myFormTitle: '编辑',
            checkedAll: [],
            sizeOpts: [10, 20, 30, 40, 50, 60, 70, 80, 90, 100],
            tableColumns: [{
                type: 'selection',
                width: 60,
                align: 'center'
            }, {
                title: '标签名',
                key: 'tags_name'
            }, {
                title: 'uri',
                key: 'tags_flag'
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
                                _this.handleEdit(params.row);
                            }
                        }
                    }, '编辑'), h('Button', {
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
            }]
        };
    },

    methods: {
        formatterFlag: function formatterFlag(row, column) {
            if (row.tags_flag == '') {
                return '';
            }
            return decodeURI(row.tags_flag);
        },
        getData: function getData() {
            var that = this;
            that.listLoading = true;
            _util2.default.ajax.get('/tags', {
                params: {
                    rows: this.pageSize,
                    page: this.currentPage
                }
            }).then(function (response) {
                var res = response.data;
                if (res != false) {
                    that.listData = res.data;
                    that.total = res.total;
                    that.currentPage = res.current_page;
                    that.listLoading = false;
                } else {
                    that.$Notice.warning({
                        title: '数据获取失败',
                        desc: ''
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

            this.getData();
        },

        handleCreate: function handleCreate() {
            var that = this;
            that.myFormTitle = '新增';
            that.myForm.id = 0;
            that.editFormVisible = true;
            that.setTopCategorys();
        },
        handleEdit: function handleEdit(row) {
            var that = this;
            that.setTopCategorys();
            that.editFormLoading = true;
            that.myFormTitle = '编辑';
            that.editFormVisible = true;
            _util2.default.ajax.get('/tags/' + row.id).then(function (response) {
                var res = response.data;
                if (res != false) {
                    res.tags_flag = decodeURI(res.tags_flag);
                    that.myForm = res;
                } else {
                    that.$Notice.warning({
                        title: '数据获取失败',
                        desc: ''
                    });
                }
                that.editFormLoading = false;
            }).catch(function (error) {
                console.log(error);
                that.editFormLoading = false;
            });
        },
        handleDistory: function handleDistory(type, row) {
            var that = this,
                idsParam = {};
            switch (type) {
                case 'one':
                    if (parseInt(row.id) <= 0) {
                        that.$Notice.warning({
                            title: '请选择需要删除的数据',
                            desc: ''
                        });
                        return false;
                    }
                    idsParam = { ids: [row.id] };
                    break;
                case 'multi':
                    var ids = that.util.getIdByArr(that.checkedAll);
                    if (ids.length <= 0) {
                        that.$Notice.warning({
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
                    _util2.default.ajax.delete('/tags/destroy', { data: idsParam }).then(function (response) {
                        that.listLoading = false;
                        var res = response.data;
                        that.$Notice.open({
                            title: res.status == 'success' ? '删除成功' : '删除失败',
                            desc: ''
                        });
                        if (type == 'one') {
                            that.util.removeByValue(that.listData, row.id);
                        } else {
                            for (var index in that.checkedAll) {
                                that.util.removeByValue(that.listData, that.checkedAll[index].id);
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
        submitMyForm: function submitMyForm() {
            var that = this;
            that.$refs['myForm'].validate(function (valid) {
                if (!valid) {
                    console.log('myForm valid error.');
                    return false;
                }

                if (that.myForm.id > 0) {
                    _util2.default.ajax.put('/tags/update', that.myForm).then(function (response) {
                        var res = response.data;
                        that.$message({
                            message: res.status == 'success' ? '编辑成功' : '编辑失败',
                            type: res.status
                        });
                        if (res.status == 'success') {
                            that.closeForm('myForm');
                            that.getData();
                        }
                    }).catch(function (error) {
                        console.log(error);
                    });
                } else {
                    _util2.default.ajax.post('/tags', that.myForm).then(function (response) {
                        console.log(response);

                        var res = response.data;
                        if (res.status == 'success') {
                            that.closeForm('myForm');
                            that.getData();
                        }
                        that.$Notice.open({
                            title: res.status == 'success' ? '新增成功' : '新增失败',
                            desc: ''
                        });
                    }).catch(function (error) {
                        if (error.response) {
                            if (error.response.status == 422) {
                                for (var index in error.response.data) {
                                    that.$Notice.warning({
                                        title: '警告',
                                        desc: error.response.data[index][0]
                                    });
                                }
                            }
                        } else {
                            console.log(error);
                        }
                    });
                }
            });
        },
        closeForm: function closeForm() {
            this.editFormVisible = false;
            this.$refs['myForm'].resetFields();
            this.myForm = {
                id: 0,
                tags_name: '',
                tags_flag: ''
            };
            console.log('closeForm');
        },
        selectRow: function selectRow(row) {
            this.checkedAll = row;
        },
        selectAll: function selectAll(selection) {
            this.checkedAll = selection;
        },

        setTopCategorys: function setTopCategorys() {
            var categorys = this.listData.concat();
            categorys.splice(0, 0, { id: 0, category_name: '顶级分类', hidden: true, category_parent: 0 });
            this.categorys = categorys;
        },
        translateWords: function translateWords(event) {
            var that = this;
            var query = that.myForm.tags_name;
            if (query.match(/\w+/g) != null) {
                that.myForm.flag = query;
            }
            if (query == null || query == '') {
                return false;
            }
            _util2.default.ajax.get('/util', {
                params: {
                    action: 'translates',
                    q: query
                }
            }).then(function (response) {
                var res = response.data;
                if (res.status == 200 && res.trans_result) {
                    var flag = res.trans_result.toLowerCase();
                    that.myForm.tags_flag = flag.replaceAll(' ', '-', flag);
                }
            }).catch(function (error) {
                console.log(error);
            });
        }
    },
    watch: {
        'myForm.tags_name': {
            handler: function handler(curVal, oldVal) {
                this.myForm.tags_flag = curVal.replace(' ', '-');
            },

            deep: true
        }
    },
    mounted: function mounted() {
        this.getData();
    }
};

/***/ }),

/***/ 321:
/***/ (function(module, exports, __webpack_require__) {

// style-loader: Adds some css to the DOM by adding a <style> tag

// load the styles
var content = __webpack_require__(322);
if(typeof content === 'string') content = [[module.i, content, '']];
if(content.locals) module.exports = content.locals;
// add the styles to the DOM
var update = __webpack_require__(17)("b8053e4e", content, false, {});
// Hot Module Replacement
if(false) {
 // When the styles change, update the <style> tags
 if(!content.locals) {
   module.hot.accept("!!../../../../../node_modules/css-loader/index.js!../../../../../node_modules/vue-loader/lib/style-compiler/index.js?{\"vue\":true,\"id\":\"data-v-16ea7d25\",\"scoped\":false,\"hasInlineConfig\":false}!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/selector.js?type=styles&index=0!./post-tags.vue", function() {
     var newContent = require("!!../../../../../node_modules/css-loader/index.js!../../../../../node_modules/vue-loader/lib/style-compiler/index.js?{\"vue\":true,\"id\":\"data-v-16ea7d25\",\"scoped\":false,\"hasInlineConfig\":false}!../../../../../node_modules/less-loader/dist/cjs.js!../../../../../node_modules/vue-loader/lib/selector.js?type=styles&index=0!./post-tags.vue");
     if(typeof newContent === 'string') newContent = [[module.id, newContent, '']];
     update(newContent);
   });
 }
 // When the module is disposed, remove the <style> tags
 module.hot.dispose(function() { update(); });
}

/***/ }),

/***/ 322:
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__(16)(false);
// imports


// module
exports.push([module.i, "", ""]);

// exports


/***/ }),

/***/ 323:
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
    attrs: { type: "primary", icon: "plus" },
    on: { click: _vm.handleCreate }
  }, [_vm._v("新增")]), _vm._v(" "), _c("Button", {
    attrs: { type: "primary", icon: "delete" },
    on: {
      click: function click($event) {
        _vm.handleDistory("multi", {});
      }
    }
  }, [_vm._v("删除")])], 1), _vm._v(" "), _c("div", { staticClass: "data-list" }, [_c("Table", {
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
  })], 1)]), _vm._v(" "), _c("Modal", {
    attrs: { title: _vm.myFormTitle },
    on: { "on-ok": _vm.submitMyForm, "on-cancel": _vm.closeForm },
    model: {
      value: _vm.editFormVisible,
      callback: function callback($$v) {
        _vm.editFormVisible = $$v;
      },
      expression: "editFormVisible"
    }
  }, [_c("div", { staticClass: "pit-dialog-edit-form" }, [_c("Form", {
    ref: "myForm",
    staticClass: "myForm",
    attrs: {
      rules: _vm.myRules,
      "label-width": 100,
      model: _vm.myForm
    }
  }, [_c("FormItem", { attrs: { label: "标签名称", prop: "tags_name" } }, [_c("Input", {
    attrs: { "auto-complete": "off" },
    on: {
      "on-blur": _vm.translateWords,
      "on-enter": _vm.translateWords
    },
    model: {
      value: _vm.myForm.tags_name,
      callback: function callback($$v) {
        _vm.$set(_vm.myForm, "tags_name", $$v);
      },
      expression: "myForm.tags_name"
    }
  })], 1), _vm._v(" "), _c("FormItem", { attrs: { label: "标签别名", prop: "tags_flag" } }, [_c("Input", {
    attrs: { "auto-complete": "off" },
    model: {
      value: _vm.myForm.tags_flag,
      callback: function callback($$v) {
        _vm.$set(_vm.myForm, "tags_flag", $$v);
      },
      expression: "myForm.tags_flag"
    }
  })], 1), _vm._v(" "), _vm.myForm.id ? _c("FormItem", [_c("Input", {
    staticStyle: { display: "none" },
    model: {
      value: _vm.myForm.id,
      callback: function callback($$v) {
        _vm.$set(_vm.myForm, "id", $$v);
      },
      expression: "myForm.id"
    }
  })], 1) : _vm._e()], 1), _vm._v(" "), _vm.editFormLoading ? _c("Spin", { attrs: { size: "large", fix: "" } }) : _vm._e()], 1)])], 1);
};
var staticRenderFns = [];
render._withStripped = true;
var esExports = { render: render, staticRenderFns: staticRenderFns };
exports.default = esExports;

if (false) {
  module.hot.accept();
  if (module.hot.data) {
    require("vue-hot-reload-api").rerender("data-v-16ea7d25", esExports);
  }
}

/***/ })

});
//# sourceMappingURL=5.chunk.js.map