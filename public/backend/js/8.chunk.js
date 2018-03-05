webpackJsonp([8],{

/***/ 183:
/***/ (function(module, exports, __webpack_require__) {

// style-loader: Adds some css to the DOM by adding a <style> tag

// load the styles
var content = __webpack_require__(184);
if(typeof content === 'string') content = [[module.i, content, '']];
// add the styles to the DOM
var update = __webpack_require__(5)(content, {});
if(content.locals) module.exports = content.locals;
// Hot Module Replacement
if(false) {
	// When the styles change, update the <style> tags
	if(!content.locals) {
		module.hot.accept("!!../../../../../node_modules/css-loader/index.js!../../../../../node_modules/vue-loader/lib/style-compiler/index.js?{\"id\":\"data-v-5e599f2c\",\"scoped\":false,\"hasInlineConfig\":false}!../../../../../node_modules/vue-loader/lib/selector.js?type=styles&index=0!./settings.vue", function() {
			var newContent = require("!!../../../../../node_modules/css-loader/index.js!../../../../../node_modules/vue-loader/lib/style-compiler/index.js?{\"id\":\"data-v-5e599f2c\",\"scoped\":false,\"hasInlineConfig\":false}!../../../../../node_modules/vue-loader/lib/selector.js?type=styles&index=0!./settings.vue");
			if(typeof newContent === 'string') newContent = [[module.id, newContent, '']];
			update(newContent);
		});
	}
	// When the module is disposed, remove the <style> tags
	module.hot.dispose(function() { update(); });
}

/***/ }),

/***/ 184:
/***/ (function(module, exports, __webpack_require__) {

exports = module.exports = __webpack_require__(4)();
// imports


// module
exports.push([module.i, "\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n\n", ""]);

// exports


/***/ }),

/***/ 185:
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



/* harmony default export */ __webpack_exports__["default"] = ({
    data() {
        return {
            editFormLoading: false,
            options: [],
            myForm: {}
        };
    },
    created() {
        this.getData();
    },
    methods: {
        getData: function () {
            let that = this;
            that.editFormLoading = true;
            __WEBPACK_IMPORTED_MODULE_0__libs_util__["a" /* default */].ajax.get('/settings').then(function (response) {
                let res = response.data;
                if (res != false) {
                    if (res.length > 0) {
                        for (var index in res) {
                            that.myForm[res[index].option_name] = res[index].option_value;
                        }
                        that.options = res;
                    }
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

        submitMyForm: function () {
            let that = this;
            __WEBPACK_IMPORTED_MODULE_0__libs_util__["a" /* default */].ajax.put('/settings/update', that.myForm).then(function (response) {
                let res = response.data;
                that.$Notice.warning({
                    title: res.status == 'success' ? '更新成功' : '更新失败',
                    desc: ''
                });
            }).catch(function (error) {
                console.log(error);
            });
        }
    },
    watch: {},
    mounted() {}
});

/***/ }),

/***/ 186:
/***/ (function(module, exports, __webpack_require__) {

module.exports = { render: function () {
    var _vm = this;var _h = _vm.$createElement;var _c = _vm._self._c || _h;
    return _c('div', {
      staticClass: "pit-post-form",
      staticStyle: {
        "width": "60%"
      }
    }, [_c('Form', {
      ref: "myForm",
      staticClass: "pit-common",
      attrs: {
        "model": _vm.myForm,
        "label-width": 150
      }
    }, [_vm._l(_vm.options, function (option) {
      return _c('FormItem', {
        key: option.id,
        attrs: {
          "value": option.id,
          "label": option.option_title
        }
      }, [_c('Input', {
        attrs: {
          "type": option.data_type
        },
        model: {
          value: _vm.myForm[option.option_name],
          callback: function ($$v) {
            _vm.$set(_vm.myForm, option.option_name, $$v);
          },
          expression: "myForm[option.option_name]"
        }
      })], 1);
    }), _vm._v(" "), _c('FormItem', [_c('Button', {
      on: {
        "click": function ($event) {
          _vm.closeForm('myForm');
        }
      }
    }, [_vm._v("取 消")]), _vm._v(" "), _c('Button', {
      attrs: {
        "type": "primary"
      },
      on: {
        "click": function ($event) {
          _vm.submitMyForm('myForm');
        }
      }
    }, [_vm._v("确 定")])], 1)], 2), _vm._v(" "), _vm.editFormLoading ? _c('Spin', {
      attrs: {
        "size": "large",
        "fix": ""
      }
    }) : _vm._e()], 1);
  }, staticRenderFns: [] };
module.exports.render._withStripped = true;
if (false) {
  module.hot.accept();
  if (module.hot.data) {
    require("vue-loader/node_modules/vue-hot-reload-api").rerender("data-v-5e599f2c", module.exports);
  }
}

/***/ }),

/***/ 97:
/***/ (function(module, exports, __webpack_require__) {


/* styles */
__webpack_require__(183)

var Component = __webpack_require__(2)(
  /* script */
  __webpack_require__(185),
  /* template */
  __webpack_require__(186),
  /* scopeId */
  null,
  /* cssModules */
  null
)
Component.options.__file = "/Users/mrcong/GoLang/src/github.com/cong5/persimmon/assets/backend/src/views/plugins/settings.vue"
if (Component.esModule && Object.keys(Component.esModule).some(function (key) {return key !== "default" && key !== "__esModule"})) {console.error("named exports are not supported in *.vue files.")}
if (Component.options.functional) {console.error("[vue-loader] settings.vue: functional components are not supported with templates, they should use render functions.")}

/* hot reload */
if (false) {(function () {
  var hotAPI = require("vue-loader/node_modules/vue-hot-reload-api")
  hotAPI.install(require("vue"), false)
  if (!hotAPI.compatible) return
  module.hot.accept()
  if (!module.hot.data) {
    hotAPI.createRecord("data-v-5e599f2c", Component.options)
  } else {
    hotAPI.reload("data-v-5e599f2c", Component.options)
  }
})()}

module.exports = Component.exports


/***/ })

});
//# sourceMappingURL=8.chunk.js.map