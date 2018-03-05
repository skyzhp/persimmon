webpackJsonp([11],{

/***/ 175:
/***/ (function(module, exports, __webpack_require__) {

// style-loader: Adds some css to the DOM by adding a <style> tag

// load the styles
var content = __webpack_require__(176);
if(typeof content === 'string') content = [[module.i, content, '']];
// add the styles to the DOM
var update = __webpack_require__(5)(content, {});
if(content.locals) module.exports = content.locals;
// Hot Module Replacement
if(false) {
	// When the styles change, update the <style> tags
	if(!content.locals) {
		module.hot.accept("!!../../../../../node_modules/css-loader/index.js!../../../../../node_modules/vue-loader/lib/style-compiler/index.js?{\"id\":\"data-v-3d1dfc8a\",\"scoped\":false,\"hasInlineConfig\":false}!../../../../../node_modules/less-loader/index.js!../../../../../node_modules/vue-loader/lib/selector.js?type=styles&index=0!./comments.vue", function() {
			var newContent = require("!!../../../../../node_modules/css-loader/index.js!../../../../../node_modules/vue-loader/lib/style-compiler/index.js?{\"id\":\"data-v-3d1dfc8a\",\"scoped\":false,\"hasInlineConfig\":false}!../../../../../node_modules/less-loader/index.js!../../../../../node_modules/vue-loader/lib/selector.js?type=styles&index=0!./comments.vue");
			if(typeof newContent === 'string') newContent = [[module.id, newContent, '']];
			update(newContent);
		});
	}
	// When the module is disposed, remove the <style> tags
	module.hot.dispose(function() { update(); });
}

/***/ }),

/***/ 176:
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__(4)();
// imports


// module
exports.push([module.i, "\n.links {\n  text-decoration: none;\n  color: #000;\n}\n.myForm {\n  width: 100% !important;\n}\n", ""]);

// exports


/***/ }),

/***/ 177:
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
            listData: [],
            categorys: [],
            currentPage: 1,
            total: 0,
            pageSize: 20,
            myForm: {
                id: 0,
                name: '',
                url: '',
                email: '',
                markdown: ''
            },
            myRules: {
                name: [{ required: true, type: "string", message: '请填写链接名称', trigger: 'blur' }],
                email: [{ required: true, type: "string", message: '请填写邮箱', trigger: 'blur' }],
                markdown: [{ required: true, type: "string", message: '请填写评论内容', trigger: 'blur' }]
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
                title: '名称',
                key: 'title',
                render: (h, params) => {
                    return h('div', [h('div', {
                        attrs: {
                            'class': 'comment-avatar'
                        }
                    }, [h('p', [h('img', {
                        attrs: {
                            'src': 'https://cn.gravatar.com/avatar/' + params.row.md5 + '?d=identicon&s=60'
                        }
                    })])]), h('div', {
                        attrs: {
                            'class': 'comment-item'
                        }
                    }, [h('p', {}, [h('a', {
                        attrs: {
                            'href': params.row.url
                        }
                    }, params.row.name)]), h('p', {}, [h('a', {
                        attrs: {
                            'href': 'mailto:' + params.row.email
                        }
                    }, params.row.email)]), h('p', params.row.ipaddress)])]);
                }
            }, {
                title: '评论内容',
                key: 'content',
                render: (h, params) => {
                    return h('div', {
                        domProps: {
                            innerHTML: params.row.content
                        }
                    });
                }
            }, {
                title: '发表于',
                key: 'posts',
                render: (h, params) => {
                    return h('span', params.row.post.title);
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
            sizeOpts: [10, 20, 30, 40, 50, 60, 70, 80, 90, 100]
        };
    },
    methods: {
        getData: function () {
            let that = this;
            that.listLoading = true;
            __WEBPACK_IMPORTED_MODULE_0__libs_util__["a" /* default */].ajax.get('/comments', {
                params: {
                    rows: that.pageSize,
                    page: this.currentPage
                }
            }).then(function (response) {
                let res = response.data;
                if (res != false) {
                    that.listData = res.data;
                    that.total = res.total;
                    that.currentPage = res.current_page;
                    that.listLoading = false;
                } else {
                    that.$Notice.warning({
                        message: '数据获取失败',
                        type: 'error',
                        duration: 3 * 1000
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
            __WEBPACK_IMPORTED_MODULE_0__libs_util__["a" /* default */].ajax.get('/comments/' + row.id).then(function (response) {
                let res = response.data;
                if (res != false) {
                    that.myForm = res;
                } else {
                    that.$Notice.warning({
                        message: '数据获取失败',
                        type: 'error'
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
                            message: '请选择需要删除的数据',
                            type: 'warning'
                        });
                        return false;
                    }
                    idsParam = { ids: [row.id] };
                    break;
                case 'multi':
                    let ids = __WEBPACK_IMPORTED_MODULE_0__libs_util__["a" /* default */].getIdByArr(that.checkedAll);
                    if (ids.length <= 0) {
                        that.$Notice.warning({
                            message: '请选择需要删除的数据',
                            type: 'warning'
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
                    __WEBPACK_IMPORTED_MODULE_0__libs_util__["a" /* default */].ajax.delete('/comments/destroy', { data: idsParam }).then(function (response) {
                        that.listLoading = false;
                        let res = response.data;
                        that.$Notice.warning({
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
                    __WEBPACK_IMPORTED_MODULE_0__libs_util__["a" /* default */].ajax.put('/comments/update', that.myForm).then(function (response) {
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
                    __WEBPACK_IMPORTED_MODULE_0__libs_util__["a" /* default */].ajax.post('/links', that.myForm).then(function (response) {
                        let res = response.data;
                        if (res.status == 'success') {
                            that.closeForm('myForm');
                            that.getData();
                        }
                        that.$Notice.warning({
                            title: res.status == 'success' ? '新增成功' : '新增失败',
                            desc: ''
                        });
                    }).catch(function (error) {
                        if (error.response) {
                            if (error.response.status == 422) {
                                for (let index in error.response.data) {
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

/***/ 178:
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
      attrs: {
        "rules": _vm.myRules,
        "label-width": 100,
        "model": _vm.myForm
      }
    }, [_c('FormItem', {
      attrs: {
        "label": "名称",
        "prop": "name"
      }
    }, [_c('Input', {
      attrs: {
        "auto-complete": "off"
      },
      model: {
        value: _vm.myForm.name,
        callback: function ($$v) {
          _vm.$set(_vm.myForm, "name", $$v);
        },
        expression: "myForm.name"
      }
    })], 1), _vm._v(" "), _c('FormItem', {
      attrs: {
        "label": "链接"
      }
    }, [_c('Input', {
      attrs: {
        "auto-complete": "off"
      },
      model: {
        value: _vm.myForm.url,
        callback: function ($$v) {
          _vm.$set(_vm.myForm, "url", $$v);
        },
        expression: "myForm.url"
      }
    })], 1), _vm._v(" "), _c('FormItem', {
      attrs: {
        "label": "E-Mail",
        "prop": "email"
      }
    }, [_c('Input', {
      attrs: {
        "auto-complete": "off"
      },
      model: {
        value: _vm.myForm.email,
        callback: function ($$v) {
          _vm.$set(_vm.myForm, "email", $$v);
        },
        expression: "myForm.email"
      }
    })], 1), _vm._v(" "), _c('FormItem', {
      attrs: {
        "label": "内容",
        "prop": "markdown"
      }
    }, [_c('Input', {
      attrs: {
        "type": "textarea",
        "autosize": "",
        "auto-complete": "off"
      },
      model: {
        value: _vm.myForm.markdown,
        callback: function ($$v) {
          _vm.$set(_vm.myForm, "markdown", $$v);
        },
        expression: "myForm.markdown"
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
    require("vue-loader/node_modules/vue-hot-reload-api").rerender("data-v-3d1dfc8a", module.exports);
  }
}

/***/ }),

/***/ 95:
/***/ (function(module, exports, __webpack_require__) {


/* styles */
__webpack_require__(175)

var Component = __webpack_require__(2)(
  /* script */
  __webpack_require__(177),
  /* template */
  __webpack_require__(178),
  /* scopeId */
  null,
  /* cssModules */
  null
)
Component.options.__file = "/Users/mrcong/GoLang/src/github.com/cong5/persimmon/assets/backend/src/views/plugins/comments.vue"
if (Component.esModule && Object.keys(Component.esModule).some(function (key) {return key !== "default" && key !== "__esModule"})) {console.error("named exports are not supported in *.vue files.")}
if (Component.options.functional) {console.error("[vue-loader] comments.vue: functional components are not supported with templates, they should use render functions.")}

/* hot reload */
if (false) {(function () {
  var hotAPI = require("vue-loader/node_modules/vue-hot-reload-api")
  hotAPI.install(require("vue"), false)
  if (!hotAPI.compatible) return
  module.hot.accept()
  if (!module.hot.data) {
    hotAPI.createRecord("data-v-3d1dfc8a", Component.options)
  } else {
    hotAPI.reload("data-v-3d1dfc8a", Component.options)
  }
})()}

module.exports = Component.exports


/***/ })

});
//# sourceMappingURL=11.chunk.js.map