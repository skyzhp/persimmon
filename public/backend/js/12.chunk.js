webpackJsonp([12],{

/***/ 121:
/***/ (function(module, exports, __webpack_require__) {

// style-loader: Adds some css to the DOM by adding a <style> tag

// load the styles
var content = __webpack_require__(122);
if(typeof content === 'string') content = [[module.i, content, '']];
// add the styles to the DOM
var update = __webpack_require__(5)(content, {});
if(content.locals) module.exports = content.locals;
// Hot Module Replacement
if(false) {
	// When the styles change, update the <style> tags
	if(!content.locals) {
		module.hot.accept("!!../../../../../node_modules/css-loader/index.js!../../../../../node_modules/vue-loader/lib/style-compiler/index.js?{\"id\":\"data-v-781823e7\",\"scoped\":false,\"hasInlineConfig\":false}!../../../../../node_modules/less-loader/index.js!../../../../../node_modules/vue-loader/lib/selector.js?type=styles&index=0!./own-space.vue", function() {
			var newContent = require("!!../../../../../node_modules/css-loader/index.js!../../../../../node_modules/vue-loader/lib/style-compiler/index.js?{\"id\":\"data-v-781823e7\",\"scoped\":false,\"hasInlineConfig\":false}!../../../../../node_modules/less-loader/index.js!../../../../../node_modules/vue-loader/lib/selector.js?type=styles&index=0!./own-space.vue");
			if(typeof newContent === 'string') newContent = [[module.id, newContent, '']];
			update(newContent);
		});
	}
	// When the module is disposed, remove the <style> tags
	module.hot.dispose(function() { update(); });
}

/***/ }),

/***/ 122:
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__(4)();
// imports


// module
exports.push([module.i, "\n.own-space-btn-box {\n  margin-bottom: 10px;\n}\n.own-space-btn-box button {\n  padding-left: 0;\n}\n.own-space-btn-box button span {\n  color: #2D8CF0;\n  transition: all .2s;\n}\n.own-space-btn-box button span:hover {\n  color: #0C25F1;\n  transition: all .2s;\n}\n.own-space-tra {\n  width: 10px;\n  height: 10px;\n  transform: rotate(45deg);\n  position: absolute;\n  top: 50%;\n  margin-top: -6px;\n  left: -3px;\n  box-shadow: 0 0 2px 3px rgba(0, 0, 0, 0.1);\n  background-color: white;\n  z-index: 100;\n}\n.own-space-input-identifycode-con {\n  position: absolute;\n  width: 200px;\n  height: 100px;\n  right: -220px;\n  top: 50%;\n  margin-top: -50px;\n  border-radius: 4px;\n  box-shadow: 0 0 2px 3px rgba(0, 0, 0, 0.1);\n}\n", ""]);

// exports


/***/ }),

/***/ 123:
/***/ (function(module, __webpack_exports__, __webpack_require__) {

"use strict";
Object.defineProperty(__webpack_exports__, "__esModule", { value: true });
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_0__libs_util__ = __webpack_require__(6);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1_js_cookie__ = __webpack_require__(9);
/* harmony import */ var __WEBPACK_IMPORTED_MODULE_1_js_cookie___default = __webpack_require__.n(__WEBPACK_IMPORTED_MODULE_1_js_cookie__);
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
            userName: '',
            saveLoading: false,
            myForm: {
                old_password: '',
                password: '',
                password_confirmation: ''
            },
            myRules: {
                old_password: [{ required: true, type: "string", message: '请填写旧密码', trigger: 'blur' }],
                password: [{ required: true, type: "string", message: '请填写新密码', trigger: 'blur' }, { min: 8, max: 64, message: '长度在 8 到 64 个字符', trigger: 'blur' }],
                password_confirmation: [{ required: true, type: "string", message: '请再填写一次密码', trigger: 'blur' }, { min: 8, max: 64, message: '长度在 8 到 64 个字符', trigger: 'blur' }]
            }
        };
    },
    created() {},
    methods: {
        submitMyForm() {
            let that = this;
            that.$refs['myForm'].validate(valid => {
                if (!valid) {
                    console.log('myForm valid error.');
                    return false;
                }
                __WEBPACK_IMPORTED_MODULE_0__libs_util__["a" /* default */].ajax.put('/user', that.myForm).then(function (response) {
                    let res = response.data;
                    that.$Notice.open({
                        title: res.status == 'success' ? '信息更新成功' : '信息更新失败',
                        desc: ''
                    });
                    if (res.status == 'success') {
                        that.closeForm('myForm');
                    }
                    setTimeout(function () {
                        __WEBPACK_IMPORTED_MODULE_1_js_cookie___default.a.remove('user');
                        that.$router.push({ path: '/login' });
                    }, 2 * 1000);
                }).catch(function (error) {
                    if (error.response) {
                        if (error.response.status == 422) {
                            for (var index in error.response.data) {
                                that.$Notice.open({
                                    title: '警告',
                                    message: error.response.data[index][0]
                                });
                            }
                        }
                    } else {
                        console.log(error);
                    }
                });
            });
        },
        closeForm() {
            this.saveLoading = false;
            this.$refs['myForm'].resetFields();
            this.myForm = {
                old_password: '',
                password: '',
                password_confirmation: ''
            };
            console.log('closeForm');
        }
    },
    mounted() {
        let user = __WEBPACK_IMPORTED_MODULE_1_js_cookie___default.a.get('user');
        let userInfo = JSON.parse(user);
        this.userName = userInfo.name;
    }
});

/***/ }),

/***/ 124:
/***/ (function(module, exports, __webpack_require__) {

module.exports = { render: function () {
    var _vm = this;var _h = _vm.$createElement;var _c = _vm._self._c || _h;
    return _c('div', [_c('Card', [_c('p', {
      attrs: {
        "slot": "title"
      },
      slot: "title"
    }, [_c('Icon', {
      attrs: {
        "type": "person"
      }
    }), _vm._v("\n            个人信息\n        ")], 1), _vm._v(" "), _c('div', [_c('Form', {
      ref: "myForm",
      attrs: {
        "model": _vm.myForm,
        "label-width": 100,
        "label-position": "right",
        "rules": _vm.myRules
      }
    }, [_c('FormItem', {
      attrs: {
        "label": "用户名："
      }
    }, [_c('span', [_vm._v(_vm._s(_vm.userName))])]), _vm._v(" "), _c('FormItem', {
      attrs: {
        "label": "旧密码：",
        "prop": "old_password"
      }
    }, [_c('div', {
      staticStyle: {
        "display": "inline-block",
        "width": "300px"
      }
    }, [_c('Input', {
      attrs: {
        "type": "password"
      },
      model: {
        value: _vm.myForm.old_password,
        callback: function ($$v) {
          _vm.$set(_vm.myForm, "old_password", $$v);
        },
        expression: "myForm.old_password"
      }
    }, [_c('Icon', {
      attrs: {
        "slot": "prepend",
        "type": "ios-locked-outline"
      },
      slot: "prepend"
    })], 1)], 1)]), _vm._v(" "), _c('FormItem', {
      attrs: {
        "label": "新密码：",
        "prop": "password"
      }
    }, [_c('div', {
      staticStyle: {
        "display": "inline-block",
        "width": "300px"
      }
    }, [_c('Input', {
      attrs: {
        "type": "password"
      },
      model: {
        value: _vm.myForm.password,
        callback: function ($$v) {
          _vm.$set(_vm.myForm, "password", $$v);
        },
        expression: "myForm.password"
      }
    }, [_c('Icon', {
      attrs: {
        "slot": "prepend",
        "type": "ios-locked-outline"
      },
      slot: "prepend"
    })], 1)], 1)]), _vm._v(" "), _c('FormItem', {
      attrs: {
        "label": "确认密码：",
        "prop": "password_confirmation"
      }
    }, [_c('div', {
      staticStyle: {
        "display": "inline-block",
        "width": "300px"
      }
    }, [_c('Input', {
      attrs: {
        "type": "password"
      },
      model: {
        value: _vm.myForm.password_confirmation,
        callback: function ($$v) {
          _vm.$set(_vm.myForm, "password_confirmation", $$v);
        },
        expression: "myForm.password_confirmation"
      }
    }, [_c('Icon', {
      attrs: {
        "slot": "prepend",
        "type": "ios-locked-outline"
      },
      slot: "prepend"
    })], 1)], 1)]), _vm._v(" "), _c('div', [_c('Button', {
      staticStyle: {
        "width": "100px"
      },
      attrs: {
        "type": "text"
      },
      on: {
        "click": _vm.closeForm
      }
    }, [_vm._v("取消")]), _vm._v(" "), _c('Button', {
      staticStyle: {
        "width": "100px"
      },
      attrs: {
        "type": "primary",
        "loading": _vm.saveLoading
      },
      on: {
        "click": _vm.submitMyForm
      }
    }, [_vm._v("保存\n                    ")])], 1)], 1)], 1)])], 1);
  }, staticRenderFns: [] };
module.exports.render._withStripped = true;
if (false) {
  module.hot.accept();
  if (module.hot.data) {
    require("vue-loader/node_modules/vue-hot-reload-api").rerender("data-v-781823e7", module.exports);
  }
}

/***/ }),

/***/ 88:
/***/ (function(module, exports, __webpack_require__) {


/* styles */
__webpack_require__(121)

var Component = __webpack_require__(2)(
  /* script */
  __webpack_require__(123),
  /* template */
  __webpack_require__(124),
  /* scopeId */
  null,
  /* cssModules */
  null
)
Component.options.__file = "/Users/mrcong/GoLang/src/github.com/cong5/persimmon/assets/backend/src/views/own-space/own-space.vue"
if (Component.esModule && Object.keys(Component.esModule).some(function (key) {return key !== "default" && key !== "__esModule"})) {console.error("named exports are not supported in *.vue files.")}
if (Component.options.functional) {console.error("[vue-loader] own-space.vue: functional components are not supported with templates, they should use render functions.")}

/* hot reload */
if (false) {(function () {
  var hotAPI = require("vue-loader/node_modules/vue-hot-reload-api")
  hotAPI.install(require("vue"), false)
  if (!hotAPI.compatible) return
  module.hot.accept()
  if (!module.hot.data) {
    hotAPI.createRecord("data-v-781823e7", Component.options)
  } else {
    hotAPI.reload("data-v-781823e7", Component.options)
  }
})()}

module.exports = Component.exports


/***/ })

});
//# sourceMappingURL=12.chunk.js.map