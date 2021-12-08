<template>
  <a-card>
    <a-table
      :columns="columns"
      :dataSource="articleTypeList"
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

      <span slot="updated_at" slot-scope="time">
        {{ time | formateDate }}
      </span>
    </a-table>
    <article-type-form
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
import ArticleTypeForm from "./components/ArticleTypeForm.vue";
const columns = [
  {
    title: "id",
    dataIndex: "id",
    align: "center",
  },

  {
    title: "创建时间",
    dataIndex: "created_at",
    align: "center",
    scopedSlots: { customRender: "created_at" },
  },
  {
    title: "名称",
    dataIndex: "name",
    align: "center",
  },
  {
    title: "修改时间",
    dataIndex: "updated_at",
    align: "center",
    scopedSlots: { customRender: "updated_at" },
  },

  {
    title: "操作",
    align: "center",
    scopedSlots: {
      customRender: "actions",
    },
  },
];
export default {
  components: { ArticleTypeForm },
  name: "ArticleType",
  data() {
    return {
      articleTypeList: [],
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
    this._getArticleType();
  },
  methods: {
    _getArticleType() {
      this.loading = true;
      articleService
        .getArticleListType({
          page: this.pageIndex,
          limit: this.pageSize,
        })
        .then((res) => {
          this.loading = false;
          const { status, message, data, total } = res.data;
          if (status !== 200) return this.$message.error(message);
          this.total = total;
          this.articleTypeList = data;
        })
        .catch((err) => {
          this.loading = false;
          this.$message.error(err);
        });
    },
    onPageChange(page, pageSize) {
      this.pageIndex = page;
      this.pageSize = pageSize;
      this._getArticleType();
    },
    onSizeChange(current, size) {
      this.pageIndex = 1;
      this.pageSize = size;
      this._getArticleType();
    },
    handleCreate() {
      this.isStatus = 1;

      this.isVisible = true;
    },
    handleUpdate(item) {
      this.isStatus = 2;
      this.modules = {
        id: item.id,
        name: item.name,
      };
      this.isVisible = true;
    },
    handleDelete(id) {
      articleService.deleteArticleType(id).then((res) => {
        const { status, message } = res.data;
        if (status !== 200) return this.$message.error(message);
        this.$message.success("删除成功");
        this._getArticleType();
      });
    },
    handleCancel() {
      this.isVisible = false;
    },
    handleOk() {
      this.confirmLoading = true;
      const form = this.$refs.articleRef.form;
      form.validateFields((error, values) => {
        console.log("==================>values", values);
        if (!error) {
          if (this.isStatus === 1) {
            articleService
              .createArticleType({ ...values })
              .then((res) => {
                this.confirmLoading = false;
                const { status, message } = res.data;
                if (status !== 200) return this.$message.error(message);
                this.$message.success("新增成功");
                form.resetFields();
                this.isVisible = false;
                this._getArticleType();
              })
              .catch((err) => {
                this.confirmLoading = false;
                this.$message.error(err);
              });
          } else {
            articleService
              .updateArticleType(values?.id, { ...values })
              .then((res) => {
                this.confirmLoading = false;
                const { status, message } = res.data;
                if (status !== 200) return this.$message.error(message);
                this.$message.success("修改成功");
                form.resetFields();
                this.isVisible = false;
                this._getArticleType();
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

<style lang="scss" scoped>
</style>