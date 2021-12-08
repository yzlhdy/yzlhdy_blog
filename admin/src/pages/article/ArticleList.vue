<template>
  <a-card>
    <a-table
      :columns="columns"
      :dataSource="articleList"
      :loading="loading"
      bordered
      :scroll="{ x: 'max-content' }"
      :rowKey="(record) => record.id"
      :pagination="{
        current: pageIndex,
        pageSize: pageSize,
        total: total,
        showSizeChanger: true,
        showLessItems: true,
        showQuickJumper: true,
        onChange: onPageChange,
        onShowSizeChange: onSizeChange,
      }"
    >
      <template slot="title">
        <a-button type="primary" @click="handleCreate">
          <a-icon type="plus"></a-icon>
          新增
        </a-button>
      </template>
      <template slot="image" slot-scope="item">
        <div style="width: 100px; height: 40px">
          <img :src="item" alt="" style="width: 100%; height: 100%" />
        </div>
      </template>
      <template slot="actions" slot-scope="item">
        <a-space size="middle">
          <a @click="handleUpdate(item)">
            <a-icon type="edit"></a-icon>
            修改
          </a>
          <a @click="handleDelete(item.id)">
            <a-icon type="delete"></a-icon>
            删除
          </a>
        </a-space>
      </template>
      <span slot="created_at" slot-scope="time">
        {{ time | formateDate }}
      </span>
    </a-table>
    <article-form
      :loading="confirmLoading"
      :status="isStatus"
      :visible="isVisible"
      ref="articleRef"
      :modules="modules"
      @cancel="handleCancel"
      @ok="handleOk"
    />
  </a-card>
</template>


<script>
import { articleService } from "@/services";
import ArticleForm from "./components/ArticleForm.vue";
export default {
  components: { ArticleForm },
  name: "ArticleList",
  data() {
    const columns = [
      {
        title: "#",
        dataIndex: "id",
        align: "center",
      },
      {
        title: "文章标题",
        dataIndex: "title",
        align: "center",
      },
      {
        title: "文章副标题",
        dataIndex: "sub_title",
        align: "center",
      },
      {
        title: "文章分类",
        dataIndex: "classification.name",
        align: "center",
      },
      {
        title: "图片",
        dataIndex: "image",
        align: "center",
        scopedSlots: {
          customRender: "image",
        },
      },
      {
        title: "内容",
        dataIndex: "content",
        align: "center",
        ellipsis: true,
        customRender: (text) => {
          return {
            children: (
              <a-tooltip placement="top">
                <template slot="title">{text}</template>
                <div class="title-text">{text}</div>
              </a-tooltip>
            ),
          };
        },
      },
      {
        title: "添加时间",
        dataIndex: "created_at",
        align: "center",
        scopedSlots: { customRender: "created_at" },
      },
      {
        title: "操作",
        align: "center",
        width: 200,
        fixed: "right",
        scopedSlots: {
          customRender: "actions",
        },
      },
    ];
    return {
      articleList: [],
      loading: false,
      columns,
      pageIndex: 1,
      pageSize: 10,
      total: 0,
      isStatus: 1,
      confirmLoading: false,
      isVisible: false,
      modules: null,
    };
  },
  created() {
    this._getArticleList();
  },
  methods: {
    onPageChange(page, pageSize) {
      this.pageIndex = page;
      this.pageSize = pageSize;
      this._getArticleList();
    },
    onSizeChange(current, size) {
      this.pageIndex = 1;
      this.pageSize = size;
      this._getArticleList();
    },
    _getArticleList() {
      this.loading = true;
      articleService
        .getArticleList({
          pageIndex: this.pageIndex,
          pageSize: this.pageSize,
        })
        .then((res) => {
          this.loading = false;
          const { status, message, data, total } = res.data;
          if (status !== 200) return this.$message.error(message);
          this.articleList = data;
          this.total = total;
        })
        .catch((err) => {
          this.loading = false;
          this.$message.error(err);
        });
    },
    handleUpdate(item) {
      const articleRef = this.$refs.articleRef;
      this.isStatus = 2;
      articleRef.initWangeEditor();
      this.modules = {
        id: item.id,
        cid: item.cid,
        title: item.title,
        sub_title: item.sub_title,
        image: item.image,
      };
      this.isVisible = true;
      articleRef.context = item.content; // 重新设置编辑器内容
    },

    handleDelete(id) {
      articleService.deleteArticle(id).then((res) => {
        const { status, message } = res.data;
        if (status !== 200) return this.$message.error(message);
        this.$message.success("删除成功");
        this._getArticleList();
      });
    },
    handleCreate() {
      this.isStatus = 1;
      this.isVisible = true;
      this.$refs.articleRef.initWangeEditor();
    },
    handleCancel() {
      this.isVisible = false;
    },
    handleOk() {
      this.confirmLoading = true;
      const form = this.$refs.articleRef.form;
      form.validateFields((error, values) => {
        if (!error) {
          if (this.isStatus === 1) {
            articleService
              .createArticle({
                ...values,
                content: this.$refs.articleRef.editor.txt.html(),
              })
              .then((res) => {
                this.confirmLoading = false;
                const { status, message } = res.data;
                if (status !== 200) return this.$message.error(message);
                this.$message.success("新增成功");
                form.resetFields();
                this.isVisible = false;
                this._getArticleList();
              })
              .catch((err) => {
                this.confirmLoading = false;
                this.$message.error(err);
              });
          } else {
            articleService
              .updateArticle(values?.id, {
                ...values,
                content: this.$refs.articleRef.editor.txt.html(),
              })
              .then((res) => {
                this.confirmLoading = false;
                const { status, message } = res.data;
                if (status !== 200) return this.$message.error(message);
                this.$message.success("修改成功");
                form.resetFields();
                this.isVisible = false;
                this._getArticleList();
              })
              .catch((err) => {
                this.confirmLoading = false;
                this.$message.error(err);
              });
          }
        } else {
          this.confirmLoading = false;
        }
      });
    },
  },
};
</script>

<style lang="less" scoped>
.title-text {
  max-width: 200px;
  text-align: center;
  // overflow: hidden;
  white-space: nowrap;
  text-overflow: ellipsis;
  cursor: pointer;
}
</style>