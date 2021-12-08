//跨域代理前缀
// const API_PROXY_PREFIX = "/api";
// const BASE_URL =
//   process.env.NODE_ENV === "production"
//     ? process.env.VUE_APP_API_BASE_URL
//     : API_PROXY_PREFIX;
const BASE_URL = process.env.VUE_APP_API_BASE_URL;
module.exports = {
  LOGIN: `${BASE_URL}/auth/login`,
  ROUTES: `${BASE_URL}/routes`,
  // 文章
  ARTICLE_TYPE: `${BASE_URL}/classification`,
  ARTICLE_TYPE_LIST: `${BASE_URL}/classification/list`,
  ARTICLE: `${BASE_URL}/article`,
};
