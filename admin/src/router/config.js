import TabsView from "@/layouts/tabs/TabsView";
// import BlankView from "@/layouts/BlankView";
import PageView from "@/layouts/PageView";

// 路由配置
const options = {
  routes: [
    {
      path: "/login",
      name: "登录页",
      component: () => import("@/pages/login"),
    },
    {
      path: "*",
      name: "404",
      component: () => import("@/pages/exception/404"),
    },
    {
      path: "/403",
      name: "403",
      component: () => import("@/pages/exception/403"),
    },
    {
      path: "/",
      name: "首页",
      component: TabsView,
      redirect: "/login",
      children: [
        {
          path: "demo",
          name: "控制台",
          meta: {
            icon: "dashboard",
          },
          component: () => import("@/pages/demo"),
        },
        {
          path: "blog",
          name: "博客管理",
          meta: {
            icon: "user",
          },
          component: PageView,
          children: [
            {
              path: "article-type",
              name: "文章类型",
              component: () => import("@/pages/article/ArticleType"),
            },
            {
              path: "article",
              name: "文章列表",
              component: () => import("@/pages/article/ArticleList"),
            },
          ],
        },
        {
          name: "验权页面",
          path: "auth/demo",
          meta: {
            icon: "file-ppt",
            authority: {
              permission: "form",
              role: "manager",
            },
            component: () => import("@/pages/demo"),
          },
        },
      ],
    },
  ],
};

export default options;
