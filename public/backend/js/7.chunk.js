webpackJsonp([7],{

/***/ 151:
/***/ (function(module, exports, __webpack_require__) {

// style-loader: Adds some css to the DOM by adding a <style> tag

// load the styles
var content = __webpack_require__(152);
if(typeof content === 'string') content = [[module.i, content, '']];
// add the styles to the DOM
var update = __webpack_require__(5)(content, {});
if(content.locals) module.exports = content.locals;
// Hot Module Replacement
if(false) {
	// When the styles change, update the <style> tags
	if(!content.locals) {
		module.hot.accept("!!../../../../../node_modules/css-loader/index.js!../../../../../node_modules/vue-loader/lib/style-compiler/index.js?{\"id\":\"data-v-fd29961e\",\"scoped\":false,\"hasInlineConfig\":false}!../../../../../node_modules/vue-loader/lib/selector.js?type=styles&index=0!./categorys.vue", function() {
			var newContent = require("!!../../../../../node_modules/css-loader/index.js!../../../../../node_modules/vue-loader/lib/style-compiler/index.js?{\"id\":\"data-v-fd29961e\",\"scoped\":false,\"hasInlineConfig\":false}!../../../../../node_modules/vue-loader/lib/selector.js?type=styles&index=0!./categorys.vue");
			if(typeof newContent === 'string') newContent = [[module.id, newContent, '']];
			update(newContent);
		});
	}
	// When the module is disposed, remove the <style> tags
	module.hot.dispose(function() { update(); });
}

/***/ }),

/***/ 152:
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__(4)();
// imports


// module
exports.push([module.i, "\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n", ""]);

// exports


/***/ }),

/***/ 153:
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
//
//



/* harmony default export */ __webpack_exports__["default"] = ({
    data() {
        return {
            listData: [],
            categorys: [],
            myForm: {
                id: 0,
                category_name: '',
                category_flag: '',
                category_description: '',
                category_parent: 0
            },
            myRules: {
                category_name: [{ required: true, type: "string", title: '请填写分类名称', trigger: 'blur' }],
                category_flag: [{ required: true, type: "string", title: '请填写分类别名', trigger: 'blur' }, { pattern: /^[a-zA-Z0-9_-]+$/, title: '只允许英文或者拼音,横杠(-),下划线(_)', trigger: 'blur' }]
            },
            editFormVisible: false,
            editFormLoading: false,
            listLoading: false,
            myFormTitle: '编辑',
            checkedAll: [],
            tableColumns: [{
                type: 'selection',
                width: 60,
                align: 'center'
            }, {
                title: '分类名称',
                key: 'category_name'
            }, {
                title: '分类别名',
                key: 'category_flag'
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
                                this.handleEdit(params.row);
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
            }]
        };
    },
    methods: {
        getData: function () {
            let that = this;
            that.listLoading = true;
            __WEBPACK_IMPORTED_MODULE_0__libs_util__["a" /* default */].ajax.get('/categorys', {
                params: {}
            }).then(function (response) {
                let res = response.data;
                if (res != false) {
                    that.listData = res.data;
                    that.listLoading = false;
                } else {
                    that.$Notice.warning({
                        title: '数据获取失败',
                        duration: 3 * 1000
                    });
                }
            }).catch(function (error) {
                console.log(error);
            });
        },
        handleCreate: function () {
            let that = this;
            that.myFormTitle = '新增';
            that.myForm.id = 0;
            that.editFormVisible = true;
            that.setTopCategorys();
        },
        handleEdit: function (row) {
            let that = this;
            that.setTopCategorys();
            that.editFormLoading = true;
            that.myFormTitle = '编辑';
            that.editFormVisible = true;
            __WEBPACK_IMPORTED_MODULE_0__libs_util__["a" /* default */].ajax.get('/categorys/' + row.id).then(function (response) {
                let res = response.data;
                if (res != false) {
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
        handleDistory: function (type, row) {
            let that = this,
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
                    var ids = __WEBPACK_IMPORTED_MODULE_0__libs_util__["a" /* default */].getIdByArr(that.checkedAll);
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
                onOk: () => {
                    that.listLoading = true;
                    __WEBPACK_IMPORTED_MODULE_0__libs_util__["a" /* default */].ajax.delete('/categorys/destroy', { data: idsParam }).then(function (response) {
                        that.listLoading = false;
                        let res = response.data;
                        that.$Notice.warning({
                            title: res.status == 'success' ? '删除成功' : '删除失败',
                            desc: ''
                        });
                        if (type == 'one') {
                            __WEBPACK_IMPORTED_MODULE_0__libs_util__["a" /* default */].removeByValue(that.listData, row.id);
                        } else {
                            for (var index in that.checkedAll) {
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
        submitMyForm: function () {
            let that = this;
            that.$refs['myForm'].validate(valid => {
                if (!valid) {
                    console.log('myForm valid error.');
                    return false;
                }

                if (that.myForm.id > 0) {
                    __WEBPACK_IMPORTED_MODULE_0__libs_util__["a" /* default */].ajax.put('/categorys/update', that.myForm).then(function (response) {
                        let res = response.data;
                        that.$Notice.open({
                            title: res.status == 'success' ? '编辑成功' : '编辑失败',
                            desc: ''
                        });
                        if (res.status == 'success') {
                            that.closeForm('myForm');
                            that.getData();
                        }
                    }).catch(function (error) {
                        console.log(error);
                    });
                } else {
                    __WEBPACK_IMPORTED_MODULE_0__libs_util__["a" /* default */].ajax.post('/categorys', that.myForm).then(function (response) {
                        console.log(response);

                        let res = response.data;
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
        closeForm: function () {
            this.editFormVisible = false;
            this.$refs['myForm'].resetFields();
            this.myForm = {
                id: 0,
                category_name: '',
                category_flag: '',
                category_description: '',
                category_parent: 0
            };
            console.log('closeForm');
        },
        setTopCategorys: function () {
            let categorys = this.listData.concat();
            categorys.splice(0, 0, { id: 0, category_name: '顶级分类', hidden: true, category_parent: 0 });
            this.categorys = categorys;
        },
        selectRow(row) {
            this.checkedAll = row;
        },
        selectAll(selection) {
            this.checkedAll = selection;
        },
        translateWords(event) {
            let that = this;
            let query = that.myForm.category_name;
            if (query.match(/\w+/g) != null) {
                that.myForm.flag = query;
            }
            if (query == null || query == '') {
                return false;
            }
            __WEBPACK_IMPORTED_MODULE_0__libs_util__["a" /* default */].ajax.get('/util', {
                params: {
                    action: 'translates',
                    q: query
                }
            }).then(function (response) {
                let res = response.data;
                if (res.status == 200 && res.trans_result) {
                    let flag = res.trans_result.toLowerCase();
                    that.myForm.category_flag = flag.replaceAll(' ', '-', flag);
                }
            }).catch(function (error) {
                console.log(error);
            });
        }
    },
    mounted() {
        this.getData();
    }
});

/***/ }),

/***/ 154:
/***/ (function(module, exports, __webpack_require__) {

module.exports = { render: function () {
    var _vm = this;var _h = _vm.$createElement;var _c = _vm._self._c || _h;
    return _c('div', {
      staticClass: "pit-content"
    }, [_c('div', {
      staticClass: "pit-action-btn"
    }, [_c('Button', {
      attrs: {
        "type": "primary",
        "icon": "plus"
      },
      on: {
        "click": _vm.handleCreate
      }
    }, [_vm._v("新增")]), _vm._v(" "), _c('Button', {
      attrs: {
        "type": "primary",
        "icon": "delete"
      },
      on: {
        "click": function ($event) {
          _vm.handleDistory('multi', {});
        }
      }
    }, [_vm._v("删除")])], 1), _vm._v(" "), _c('div', {
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
    })], 1), _vm._v(" "), _c('Modal', {
      attrs: {
        "title": _vm.myFormTitle
      },
      on: {
        "on-ok": _vm.submitMyForm,
        "on-cancel": _vm.closeForm
      },
      model: {
        value: _vm.editFormVisible,
        callback: function ($$v) {
          _vm.editFormVisible = $$v;
        },
        expression: "editFormVisible"
      }
    }, [_c('div', {
      staticClass: "pit-dialog-edit-form"
    }, [_c('Form', {
      ref: "myForm",
      staticClass: "myForm",
      staticStyle: {
        "width": "90%"
      },
      attrs: {
        "rules": _vm.myRules,
        "model": _vm.myForm,
        "label-position": "right",
        "label-width": 100
      }
    }, [_c('FormItem', {
      attrs: {
        "label": "分类名称",
        "prop": "category_name"
      }
    }, [_c('Input', {
      attrs: {
        "auto-complete": "off"
      },
      on: {
        "on-blur": _vm.translateWords,
        "on-enter": _vm.translateWords
      },
      model: {
        value: _vm.myForm.category_name,
        callback: function ($$v) {
          _vm.$set(_vm.myForm, "category_name", $$v);
        },
        expression: "myForm.category_name"
      }
    })], 1), _vm._v(" "), _c('FormItem', {
      attrs: {
        "label": "分类别名",
        "prop": "category_flag"
      }
    }, [_c('Input', {
      attrs: {
        "auto-complete": "off"
      },
      model: {
        value: _vm.myForm.category_flag,
        callback: function ($$v) {
          _vm.$set(_vm.myForm, "category_flag", $$v);
        },
        expression: "myForm.category_flag"
      }
    })], 1), _vm._v(" "), _c('FormItem', {
      attrs: {
        "label": "分类描述"
      }
    }, [_c('Input', {
      attrs: {
        "type": "textarea"
      },
      model: {
        value: _vm.myForm.category_description,
        callback: function ($$v) {
          _vm.$set(_vm.myForm, "category_description", $$v);
        },
        expression: "myForm.category_description"
      }
    })], 1), _vm._v(" "), _c('FormItem', {
      attrs: {
        "label": "父分类"
      }
    }, [_c('Select', {
      attrs: {
        "placeholder": "请选择父分类"
      },
      model: {
        value: _vm.myForm.category_parent,
        callback: function ($$v) {
          _vm.$set(_vm.myForm, "category_parent", $$v);
        },
        expression: "myForm.category_parent"
      }
    }, _vm._l(_vm.categorys, function (item) {
      return _c('Option', {
        key: item.id,
        attrs: {
          "value": item.id
        }
      }, [_vm._v(_vm._s(item.category_name))]);
    }))], 1), _vm._v(" "), _vm.myForm.id ? _c('FormItem', [_c('Input', {
      staticStyle: {
        "display": "none"
      },
      model: {
        value: _vm.myForm.id,
        callback: function ($$v) {
          _vm.$set(_vm.myForm, "id", $$v);
        },
        expression: "myForm.id"
      }
    })], 1) : _vm._e()], 1), _vm._v(" "), _vm.editFormLoading ? _c('Spin', {
      attrs: {
        "size": "large",
        "fix": ""
      }
    }) : _vm._e()], 1)])], 1);
  }, staticRenderFns: [] };
module.exports.render._withStripped = true;
if (false) {
  module.hot.accept();
  if (module.hot.data) {
    require("vue-loader/node_modules/vue-hot-reload-api").rerender("data-v-fd29961e", module.exports);
  }
}

/***/ }),

/***/ 90:
/***/ (function(module, exports, __webpack_require__) {


/* styles */
__webpack_require__(151)

var Component = __webpack_require__(2)(
  /* script */
  __webpack_require__(153),
  /* template */
  __webpack_require__(154),
  /* scopeId */
  null,
  /* cssModules */
  null
)
Component.options.__file = "/Users/mrcong/GoLang/src/github.com/cong5/persimmon/assets/backend/src/views/posts/categorys.vue"
if (Component.esModule && Object.keys(Component.esModule).some(function (key) {return key !== "default" && key !== "__esModule"})) {console.error("named exports are not supported in *.vue files.")}
if (Component.options.functional) {console.error("[vue-loader] categorys.vue: functional components are not supported with templates, they should use render functions.")}

/* hot reload */
if (false) {(function () {
  var hotAPI = require("vue-loader/node_modules/vue-hot-reload-api")
  hotAPI.install(require("vue"), false)
  if (!hotAPI.compatible) return
  module.hot.accept()
  if (!module.hot.data) {
    hotAPI.createRecord("data-v-fd29961e", Component.options)
  } else {
    hotAPI.reload("data-v-fd29961e", Component.options)
  }
})()}

module.exports = Component.exports


/***/ })

});
//# sourceMappingURL=7.chunk.js.map