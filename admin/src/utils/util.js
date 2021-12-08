import enquireJs from "enquire.js";
import moment from "moment";
export function isDef(v) {
  return v !== undefined && v !== null;
}

/**
 * Remove an item from an array.
 */
export function remove(arr, item) {
  if (arr.length) {
    const index = arr.indexOf(item);
    if (index > -1) {
      return arr.splice(index, 1);
    }
  }
}

export function isRegExp(v) {
  return _toString.call(v) === "[object RegExp]";
}

export function enquireScreen(call) {
  const handler = {
    match: function() {
      call && call(true);
    },
    unmatch: function() {
      call && call(false);
    },
  };
  enquireJs.register("only screen and (max-width: 767.99px)", handler);
}
export const formateDate = (str) => {
  return moment(str).format("YYYY-MM-DD HH:mm:ss");
};
const _toString = Object.prototype.toString;
