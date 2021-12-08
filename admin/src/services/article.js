import { ARTICLE_TYPE, ARTICLE_TYPE_LIST, ARTICLE } from "./api";
import { request, METHOD } from "@/utils/request";

// 获取文章列表
export const getArticleListType = (params) => {
  return request(ARTICLE_TYPE_LIST, METHOD.GET, params);
};

// 获取文章详情
export const createArticleType = (data) => {
  return request(ARTICLE_TYPE, METHOD.POST, data);
};

export const updateArticleType = (id, data) => {
  return request(`${ARTICLE_TYPE}/${id}`, METHOD.PUT, data);
};

export const deleteArticleType = (id) => {
  return request(`${ARTICLE_TYPE}/${id}`, METHOD.DELETE);
};

export const getArticleList = (params) => {
  return request(ARTICLE, METHOD.GET, params);
};
export const createArticle = (data) => {
  return request(ARTICLE, METHOD.POST, data);
};
export const updateArticle = (id, data) => {
  return request(`${ARTICLE}/${id}`, METHOD.PUT, data);
};
export const deleteArticle = (id) => {
  return request(`${ARTICLE}/${id}`, METHOD.DELETE);
};
