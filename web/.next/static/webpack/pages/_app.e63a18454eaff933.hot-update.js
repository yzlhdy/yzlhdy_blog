"use strict";
/*
 * ATTENTION: An "eval-source-map" devtool has been used.
 * This devtool is neither made for production nor for readable output files.
 * It uses "eval()" calls to create a separate source file with attached SourceMaps in the browser devtools.
 * If you are trying to read the output file, select a different devtool (https://webpack.js.org/configuration/devtool/)
 * or disable the default devtool with "devtool: false".
 * If you are looking for production-ready output files, see mode: "production" (https://webpack.js.org/configuration/mode/).
 */
self["webpackHotUpdate_N_E"]("pages/_app",{

/***/ "./src/components/Layout/index.tsx":
/*!*****************************************!*\
  !*** ./src/components/Layout/index.tsx ***!
  \*****************************************/
/***/ (function(module, __webpack_exports__, __webpack_require__) {

eval("__webpack_require__.r(__webpack_exports__);\n/* harmony import */ var react_jsx_runtime__WEBPACK_IMPORTED_MODULE_0__ = __webpack_require__(/*! react/jsx-runtime */ \"./node_modules/react/jsx-runtime.js\");\n/* harmony import */ var react_jsx_runtime__WEBPACK_IMPORTED_MODULE_0___default = /*#__PURE__*/__webpack_require__.n(react_jsx_runtime__WEBPACK_IMPORTED_MODULE_0__);\n/* harmony import */ var _chakra_ui_layout__WEBPACK_IMPORTED_MODULE_4__ = __webpack_require__(/*! @chakra-ui/layout */ \"./node_modules/@chakra-ui/layout/dist/chakra-ui-layout.esm.js\");\n/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_1__ = __webpack_require__(/*! react */ \"./node_modules/react/index.js\");\n/* harmony import */ var react__WEBPACK_IMPORTED_MODULE_1___default = /*#__PURE__*/__webpack_require__.n(react__WEBPACK_IMPORTED_MODULE_1__);\n/* harmony import */ var _Footer__WEBPACK_IMPORTED_MODULE_2__ = __webpack_require__(/*! ./Footer */ \"./src/components/Layout/Footer.tsx\");\n/* harmony import */ var _Header__WEBPACK_IMPORTED_MODULE_3__ = __webpack_require__(/*! ./Header */ \"./src/components/Layout/Header.tsx\");\n/* module decorator */ module = __webpack_require__.hmd(module);\n\n\n\n\n\nvar _this = undefined;\nvar Layout = function(param) {\n    var children = param.children;\n    return(/*#__PURE__*/ (0,react_jsx_runtime__WEBPACK_IMPORTED_MODULE_0__.jsx)(_chakra_ui_layout__WEBPACK_IMPORTED_MODULE_4__.Container, {\n        maxW: \"container.xl\",\n        __source: {\n            fileName: \"D:\\\\2-Exercise\\\\6-Go\\\\src\\\\yzlhdy_blog\\\\web\\\\src\\\\components\\\\Layout\\\\index.tsx\",\n            lineNumber: 10,\n            columnNumber: 3\n        },\n        __self: _this,\n        children: /*#__PURE__*/ (0,react_jsx_runtime__WEBPACK_IMPORTED_MODULE_0__.jsxs)(_chakra_ui_layout__WEBPACK_IMPORTED_MODULE_4__.VStack, {\n            spacing: 0,\n            __source: {\n                fileName: \"D:\\\\2-Exercise\\\\6-Go\\\\src\\\\yzlhdy_blog\\\\web\\\\src\\\\components\\\\Layout\\\\index.tsx\",\n                lineNumber: 11,\n                columnNumber: 4\n            },\n            __self: _this,\n            children: [\n                /*#__PURE__*/ (0,react_jsx_runtime__WEBPACK_IMPORTED_MODULE_0__.jsx)(_Header__WEBPACK_IMPORTED_MODULE_3__[\"default\"], {\n                    __source: {\n                        fileName: \"D:\\\\2-Exercise\\\\6-Go\\\\src\\\\yzlhdy_blog\\\\web\\\\src\\\\components\\\\Layout\\\\index.tsx\",\n                        lineNumber: 12,\n                        columnNumber: 5\n                    },\n                    __self: _this\n                }),\n                children,\n                /*#__PURE__*/ (0,react_jsx_runtime__WEBPACK_IMPORTED_MODULE_0__.jsx)(_Footer__WEBPACK_IMPORTED_MODULE_2__[\"default\"], {\n                    __source: {\n                        fileName: \"D:\\\\2-Exercise\\\\6-Go\\\\src\\\\yzlhdy_blog\\\\web\\\\src\\\\components\\\\Layout\\\\index.tsx\",\n                        lineNumber: 16,\n                        columnNumber: 5\n                    },\n                    __self: _this\n                })\n            ]\n        })\n    }));\n};\n_c = Layout;\n/* harmony default export */ __webpack_exports__[\"default\"] = (Layout);\nvar _c;\n$RefreshReg$(_c, \"Layout\");\n\n\n;\n    var _a, _b;\n    // Legacy CSS implementations will `eval` browser code in a Node.js context\n    // to extract CSS. For backwards compatibility, we need to check we're in a\n    // browser context before continuing.\n    if (typeof self !== 'undefined' &&\n        // AMP / No-JS mode does not inject these helpers:\n        '$RefreshHelpers$' in self) {\n        var currentExports = module.__proto__.exports;\n        var prevExports = (_b = (_a = module.hot.data) === null || _a === void 0 ? void 0 : _a.prevExports) !== null && _b !== void 0 ? _b : null;\n        // This cannot happen in MainTemplate because the exports mismatch between\n        // templating and execution.\n        self.$RefreshHelpers$.registerExportsForReactRefresh(currentExports, module.id);\n        // A module can be accepted automatically based on its exports, e.g. when\n        // it is a Refresh Boundary.\n        if (self.$RefreshHelpers$.isReactRefreshBoundary(currentExports)) {\n            // Save the previous exports on update so we can compare the boundary\n            // signatures.\n            module.hot.dispose(function (data) {\n                data.prevExports = currentExports;\n            });\n            // Unconditionally accept an update to this module, we'll check if it's\n            // still a Refresh Boundary later.\n            module.hot.accept();\n            // This field is set when the previous version of this module was a\n            // Refresh Boundary, letting us know we need to check for invalidation or\n            // enqueue an update.\n            if (prevExports !== null) {\n                // A boundary can become ineligible if its exports are incompatible\n                // with the previous exports.\n                //\n                // For example, if you add/remove/change exports, we'll want to\n                // re-execute the importing modules, and force those components to\n                // re-render. Similarly, if you convert a class component to a\n                // function, we want to invalidate the boundary.\n                if (self.$RefreshHelpers$.shouldInvalidateReactRefreshBoundary(prevExports, currentExports)) {\n                    module.hot.invalidate();\n                }\n                else {\n                    self.$RefreshHelpers$.scheduleUpdate();\n                }\n            }\n        }\n        else {\n            // Since we just executed the code for the module, it's possible that the\n            // new exports made it ineligible for being a boundary.\n            // We only care about the case when we were _previously_ a boundary,\n            // because we already accepted this update (accidental side effect).\n            var isNoLongerABoundary = prevExports !== null;\n            if (isNoLongerABoundary) {\n                module.hot.invalidate();\n            }\n        }\n    }\n//# sourceURL=[module]\n//# sourceMappingURL=data:application/json;charset=utf-8;base64,eyJ2ZXJzaW9uIjozLCJmaWxlIjoiLi9zcmMvY29tcG9uZW50cy9MYXlvdXQvaW5kZXgudHN4LmpzIiwibWFwcGluZ3MiOiI7Ozs7Ozs7Ozs7QUFBcUQ7QUFFTDtBQUNuQjtBQUNBOztBQUc3QixHQUFLLENBQUNLLE1BQU0sR0FBb0IsUUFBUSxRQUFVLENBQUM7UUFBaEJDLFFBQVEsU0FBUkEsUUFBUTtJQUMxQyxNQUFNLHNFQUNKTix3REFBUztRQUFDTyxJQUFJLEVBQUMsQ0FBYzs7Ozs7Ozt3RkFDNUJOLHFEQUFNO1lBQUNPLE9BQU8sRUFBRSxDQUFDOzs7Ozs7OztxRkFDaEJKLCtDQUFNOzs7Ozs7OztnQkFFTkUsUUFBUTtxRkFFUkgsK0NBQU07Ozs7Ozs7Ozs7O0FBTVgsQ0FBQztLQWRLRSxNQUFNO0FBZ0JaLCtEQUFlQSxNQUFNLEVBQUMiLCJzb3VyY2VzIjpbIndlYnBhY2s6Ly9fTl9FLy4vc3JjL2NvbXBvbmVudHMvTGF5b3V0L2luZGV4LnRzeD84ZjdkIl0sInNvdXJjZXNDb250ZW50IjpbImltcG9ydCB7IENvbnRhaW5lciwgVlN0YWNrIH0gZnJvbSAnQGNoYWtyYS11aS9sYXlvdXQnO1xyXG5cclxuaW1wb3J0IFJlYWN0LCB7IFByb3BzV2l0aENoaWxkcmVuIH0gZnJvbSAncmVhY3QnO1xyXG5pbXBvcnQgRm9vdGVyIGZyb20gJy4vRm9vdGVyJztcclxuaW1wb3J0IEhlYWRlciBmcm9tICcuL0hlYWRlcic7XHJcbnR5cGUgUHJvcHMgPSBQcm9wc1dpdGhDaGlsZHJlbjx7fT47XHJcblxyXG5jb25zdCBMYXlvdXQ6IFJlYWN0LkZDPFByb3BzPiA9ICh7IGNoaWxkcmVuIH0pID0+IHtcclxuXHRyZXR1cm4gKFxyXG5cdFx0PENvbnRhaW5lciBtYXhXPVwiY29udGFpbmVyLnhsXCIgPlxyXG5cdFx0XHQ8VlN0YWNrIHNwYWNpbmc9ezB9ID5cclxuXHRcdFx0XHQ8SGVhZGVyIC8+XHJcblx0XHRcdFx0e1xyXG5cdFx0XHRcdFx0Y2hpbGRyZW5cclxuXHRcdFx0XHR9XHJcblx0XHRcdFx0PEZvb3RlciAvPlxyXG5cclxuXHRcdFx0PC9WU3RhY2s+XHJcblx0XHQ8L0NvbnRhaW5lcj5cclxuXHJcblx0KVxyXG59XHJcblxyXG5leHBvcnQgZGVmYXVsdCBMYXlvdXQ7Il0sIm5hbWVzIjpbIkNvbnRhaW5lciIsIlZTdGFjayIsIlJlYWN0IiwiRm9vdGVyIiwiSGVhZGVyIiwiTGF5b3V0IiwiY2hpbGRyZW4iLCJtYXhXIiwic3BhY2luZyJdLCJzb3VyY2VSb290IjoiIn0=\n//# sourceURL=webpack-internal:///./src/components/Layout/index.tsx\n");

/***/ })

});