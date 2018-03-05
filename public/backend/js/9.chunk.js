webpackJsonp([9],{

/***/ 179:
/***/ (function(module, exports, __webpack_require__) {

// style-loader: Adds some css to the DOM by adding a <style> tag

// load the styles
var content = __webpack_require__(180);
if(typeof content === 'string') content = [[module.i, content, '']];
// add the styles to the DOM
var update = __webpack_require__(5)(content, {});
if(content.locals) module.exports = content.locals;
// Hot Module Replacement
if(false) {
	// When the styles change, update the <style> tags
	if(!content.locals) {
		module.hot.accept("!!../../../../../node_modules/css-loader/index.js!../../../../../node_modules/vue-loader/lib/style-compiler/index.js?{\"id\":\"data-v-4ee6e5b2\",\"scoped\":false,\"hasInlineConfig\":false}!../../../../../node_modules/vue-loader/lib/selector.js?type=styles&index=0!./options.vue", function() {
			var newContent = require("!!../../../../../node_modules/css-loader/index.js!../../../../../node_modules/vue-loader/lib/style-compiler/index.js?{\"id\":\"data-v-4ee6e5b2\",\"scoped\":false,\"hasInlineConfig\":false}!../../../../../node_modules/vue-loader/lib/selector.js?type=styles&index=0!./options.vue");
			if(typeof newContent === 'string') newContent = [[module.id, newContent, '']];
			update(newContent);
		});
	}
	// When the module is disposed, remove the <style> tags
	module.hot.dispose(function() { update(); });
}

/***/ }),

/***/ 180:
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__(4)();
// imports


// module
exports.push([module.i, "\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n", ""]);

// exports


/***/ }),

/***/ 181:
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
                title: '配置项说明',
                key: 'option_title'
            }, {
                title: '配置项名称',
                key: 'option_name'
            }, {
                title: '配置分组',
                key: 'option_group'
            }, {
                title: '配置项备注',
                key: 'option_remark'
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
            }],
            listData: [],
            currentPage: 1,
            total: 0,
            pageSize: 20,
            myForm: {
                id: 0,
                option_title: '',
                option_name: '',
                option_value: '',
                option_group: '',
                option_remark: ''
            },
            myRules: {
                option_title: [{ required: true, type: "string", message: '请填写配置项说明', trigger: 'blur' }],
                option_name: [{ required: true, type: "string", message: '请填写配置项名称', trigger: 'blur' }, { pattern: /^[a-zA-Z0-9_]+$/, message: '只允许英文或者拼音,横杠(-)', trigger: 'blur' }]
            },
            editFormVisible: false,
            editFormLoading: false,
            listLoading: false,
            myFormTitle: '编辑',
            checkedAll: [],
            sizeOpts: [10, 20, 30, 40, 50, 60, 70, 80, 90, 100]
        };
    },
    methods: {
        getData: function () {
            let that = this;
            that.listLoading = true;
            __WEBPACK_IMPORTED_MODULE_0__libs_util__["a" /* default */].ajax.get('/options', {
                params: {
                    rows: this.pageSize
                }
            }).then(function (response) {
                let res = response.data;
                if (res != false) {
                    that.listData = res.data;
                    that.total = res.total;
                    that.currentPage = res.current_page;
                    that.listLoading = false;
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
        handleSizeChange(val) {
            //console.log(`每页 ${val} 条`);
            this.pageSize = val;
            this.getData();
        },
        handleCurrentChange(val) {
            this.currentPage = val;
            //console.log(`当前页: ${val}`);
        },
        handleCreate: function () {
            let that = this;
            that.myFormTitle = '新增';
            that.myForm.id = 0;
            that.editFormVisible = true;
        },
        handleEdit: function (row) {
            let that = this;
            that.editFormLoading = true;
            that.myFormTitle = '编辑';
            that.editFormVisible = true;
            __WEBPACK_IMPORTED_MODULE_0__libs_util__["a" /* default */].ajax.get('/options/' + row.id).then(function (response) {
                let res = response.data;
                if (res != false) {
                    that.myForm = res;
                } else {
                    that.$Notice.open({
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
                        that.$Notice.open({
                            title: '请选择需要删除的数据',
                            desc: ''
                        });
                        return false;
                    }
                    idsParam = { ids: [row.id] };
                    break;
                case 'multi':
                    let ids = __WEBPACK_IMPORTED_MODULE_0__libs_util__["a" /* default */].getIdByArr(that.checkedAll);
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
                onOk: () => {
                    that.listLoading = true;
                    __WEBPACK_IMPORTED_MODULE_0__libs_util__["a" /* default */].ajax.delete('/options/destroy', { data: idsParam }).then(function (response) {
                        that.listLoading = false;
                        let res = response.data;
                        that.$Notice.open({
                            title: res.status == 'success' ? '删除成功' : '删除失败',
                            desc: ''
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
        submitMyForm: function () {
            let that = this;
            that.$refs['myForm'].validate(valid => {
                if (!valid) {
                    console.log('myForm valid error.');
                    return false;
                }

                if (that.myForm.id > 0) {
                    __WEBPACK_IMPORTED_MODULE_0__libs_util__["a" /* default */].ajax.put('/options/update', that.myForm).then(function (response) {
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
                    __WEBPACK_IMPORTED_MODULE_0__libs_util__["a" /* default */].ajax.post('/options', that.myForm).then(function (response) {

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
                                for (let index in error.response.data) {
                                    that.$notify({
                                        title: '警告',
                                        message: error.response.data[index][0]
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
                name: '',
                url: '',
                logo: '',
                group: ''
            };
            console.log('closeForm');
        },
        selectRow(row) {
            this.checkedAll = row;
        },
        selectAll(selection) {
            this.checkedAll = selection;
        }
    },
    watch: {},
    mounted() {
        this.getData();
    }
});

/***/ }),

/***/ 182:
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
    })], 1)]), _vm._v(" "), _c('Modal', {
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
        "width": "80%"
      },
      attrs: {
        "rules": _vm.myRules,
        "label-width": 100,
        "model": _vm.myForm
      }
    }, [_c('FormItem', {
      attrs: {
        "label": "配置项说明",
        "prop": "option_title"
      }
    }, [_c('Input', {
      attrs: {
        "auto-complete": "off"
      },
      model: {
        value: _vm.myForm.option_title,
        callback: function ($$v) {
          _vm.$set(_vm.myForm, "option_title", $$v);
        },
        expression: "myForm.option_title"
      }
    })], 1), _vm._v(" "), _c('FormItem', {
      attrs: {
        "label": "配置项名称",
        "prop": "option_name"
      }
    }, [_c('Input', {
      attrs: {
        "auto-complete": "off"
      },
      model: {
        value: _vm.myForm.option_name,
        callback: function ($$v) {
          _vm.$set(_vm.myForm, "option_name", $$v);
        },
        expression: "myForm.option_name"
      }
    })], 1), _vm._v(" "), _c('FormItem', {
      attrs: {
        "label": "配置项值"
      }
    }, [_c('Input', {
      attrs: {
        "type": _vm.myForm.data_type,
        "autosize": _vm.myForm.data_type == 'textarea',
        "auto-complete": "off"
      },
      model: {
        value: _vm.myForm.option_value,
        callback: function ($$v) {
          _vm.$set(_vm.myForm, "option_value", $$v);
        },
        expression: "myForm.option_value"
      }
    })], 1), _vm._v(" "), _c('FormItem', {
      attrs: {
        "label": "配置分组"
      }
    }, [_c('Input', {
      attrs: {
        "auto-complete": "off"
      },
      model: {
        value: _vm.myForm.option_group,
        callback: function ($$v) {
          _vm.$set(_vm.myForm, "option_group", $$v);
        },
        expression: "myForm.option_group"
      }
    })], 1), _vm._v(" "), _c('FormItem', {
      attrs: {
        "label": "配置项备注"
      }
    }, [_c('Input', {
      attrs: {
        "type": "textarea",
        "autosize": {
          minRows: 3,
          maxRows: 5
        },
        "auto-complete": "off"
      },
      model: {
        value: _vm.myForm.option_remark,
        callback: function ($$v) {
          _vm.$set(_vm.myForm, "option_remark", $$v);
        },
        expression: "myForm.option_remark"
      }
    })], 1), _vm._v(" "), _vm.myForm.id ? _c('FormItem', [_c('Input', {
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
    require("vue-loader/node_modules/vue-hot-reload-api").rerender("data-v-4ee6e5b2", module.exports);
  }
}

/***/ }),

/***/ 96:
/***/ (function(module, exports, __webpack_require__) {


/* styles */
__webpack_require__(179)

var Component = __webpack_require__(2)(
  /* script */
  __webpack_require__(181),
  /* template */
  __webpack_require__(182),
  /* scopeId */
  null,
  /* cssModules */
  null
)
Component.options.__file = "/Users/mrcong/GoLang/src/github.com/cong5/persimmon/assets/backend/src/views/plugins/options.vue"
if (Component.esModule && Object.keys(Component.esModule).some(function (key) {return key !== "default" && key !== "__esModule"})) {console.error("named exports are not supported in *.vue files.")}
if (Component.options.functional) {console.error("[vue-loader] options.vue: functional components are not supported with templates, they should use render functions.")}

/* hot reload */
if (false) {(function () {
  var hotAPI = require("vue-loader/node_modules/vue-hot-reload-api")
  hotAPI.install(require("vue"), false)
  if (!hotAPI.compatible) return
  module.hot.accept()
  if (!module.hot.data) {
    hotAPI.createRecord("data-v-4ee6e5b2", Component.options)
  } else {
    hotAPI.reload("data-v-4ee6e5b2", Component.options)
  }
})()}

module.exports = Component.exports


/***/ })

});
//# sourceMappingURL=9.chunk.js.map