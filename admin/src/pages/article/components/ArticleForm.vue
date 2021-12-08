<template>
  <a-modal
    :title="status === 1 ? '新增文章' : '修改文章'"
    :visible="visible"
    :confirmLoading="loading"
    width="25%"
    @ok="
      () => {
        $emit('ok');
      }
    "
    @cancel="
      () => {
        $emit('cancel');
      }
    "
  >
    <a-spin :spinning="loading">
      <a-form :form="form" v-bind="formLayout">
        <a-form-item v-show="false" label="ID">
          <a-input v-decorator="['id', { initialValue: 0 }]" disabled />
        </a-form-item>
        <a-form-item label="分类名称">
          <a-select
            placeholder="请输入分类名称"
            v-decorator="[
              'cid',
              {
                rules: [
                  {
                    required: true,
                    message: '请选择分类',
                  },
                ],
              },
            ]"
          >
            <a-select-option
              :key="item.id"
              :value="item.id"
              v-for="item in articleTypeList"
              >{{ item.name }}</a-select-option
            >
          </a-select>
        </a-form-item>
        <a-form-item label="标题">
          <a-input
            placeholder="请输入标题"
            v-decorator="[
              'title',
              {
                rules: [
                  {
                    required: true,
                    message: '请输入标题',
                  },
                ],
              },
            ]"
          />
        </a-form-item>
        <a-form-item label="副标题">
          <a-input
            placeholder="请输入标题"
            v-decorator="[
              'sub_title',
              {
                rules: [
                  {
                    required: true,
                    message: '请输入标题',
                  },
                ],
              },
            ]"
          />
        </a-form-item>
        <a-form-item label="图片地址">
          <a-input
            placeholder="请输入图片地址"
            v-decorator="[
              'image',
              {
                rules: [
                  {
                    required: true,
                    message: '请输入图片地址',
                  },
                ],
              },
            ]"
          />
        </a-form-item>
        <a-form-item label="文章内容">
          <div ref="editor"></div>
        </a-form-item>
      </a-form>
    </a-spin>
  </a-modal>
</template>

<script>
import { articleService } from "@/services";
import E from "wangeditor";
// 表单字段
const fields = ["id", "title", "cid", "sub_title", "image"];
export default {
  name: "StaffForm",
  props: {
    visible: {
      type: Boolean,
      required: true,
    },
    loading: {
      type: Boolean,
      default: () => false,
    },
    modules: {
      type: Object,
      default: () => null,
    },

    status: {
      type: Number,
      default: 1,
    },
  },
  data() {
    this.formLayout = {
      labelCol: {
        span: 5,
      },
      wrapperCol: {
        span: 19,
      },
    };
    return {
      form: this.$form.createForm(this),
      articleTypeList: [],
      editor: null,
      context: "",
    };
  },
  created() {
    // 防止表单未注册
    this._articleTypeList();
    fields.forEach((v) => this.form.getFieldDecorator(v));
    // 当 model 发生改变时，为表单设置值
    this.$watch("modules", () => {
      this.modules && this.form.setFieldsValue(this.modules);
    });
  },
  methods: {
    _articleTypeList() {
      articleService
        .getArticleListType({
          page: 1,
          limit: 100,
        })
        .then((res) => {
          const { status, message, data } = res.data;
          if (status !== 200) this.$message.error(message);
          this.articleTypeList = data;
        });
    },
    initWangeEditor() {
      this.$nextTick(() => {
        const editor = new E(this.$refs.editor);
        console.log("==================>new成功");
        // editor.config.uploadImgHeaders = {
        //   token: this.headers.token,
        // };
        editor.config.height = 40;
        // 自定义 fileName
        // 上传函数
        editor.config.uploadImgAccept = [
          "jpg",
          "jpeg",
          "png",
          "gif",
          "bmp",
          "webp",
        ];

        // 图片上传回调
        // editor.config.customUploadImg = function (files, insertImgFn) {
        //   let formData = new FormData();
        //   formData.append("picture", files[0]);
        //   formData.append("type", "help");
        //   //   let data = await this.upload(formData);
        //   commodityService
        //     .addImage(formData)
        //     .then((res) => {
        //       const { code, msg, data } = res.data;
        //       if (code !== 200) {
        //         this.$message.error(msg);
        //       } else {
        //         this.$message.success("上传成功");
        //         const url = `${process.env.VUE_APP_IMAGE_URL}/${data.img_path}`;
        //         insertImgFn(url);
        //       }
        //     })
        //     .catch((err) => {
        //       this.$message.error(err);
        //     });

        //   // 上传图片，返回结果，将图片插入到编辑器中
        // };
        editor.config.showLinkImg = false;
        this.editor = editor;
        editor.create();
        this.editor.txt.html(this.context); // 重新设置编辑器内容
      });
    },
  },
};
</script>