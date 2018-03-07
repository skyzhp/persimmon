<style lang="less">
    @import './post-edit.less';

    .ivu-form-item-error .ivu-input-group-prepend {
        background: #f0f0f0;
    }
</style>

<template>
    <div class="pit-post-form">
        <Form :model="myForm" ref="myForm" :rules="myRules" label-position="right" :label-width="100"
              style="width: 90%;">
            <FormItem label="标题" prop="title">
                <Input v-model="myForm.title" @on-blur="titleBlur"></Input>
            </FormItem>
            <FormItem label="Uri" prop="flag">
                <Input v-model="myForm.flag">
                <span slot="prepend">{{domain}}</span>
                </Input>
            </FormItem>
            <FormItem label="标签" prop="tags">
                <Tag v-for="tag in myForm.tags" color="blue" :key="tag" :name="tag" closable @on-close="closeTags">
                    {{ tag }}
                </Tag>
                <Input v-model="newTag" @on-enter="addTag" @on-blur="addTag" size="small" placeholder="+"
                       style="width: 50px;"></Input>
            </FormItem>
            <FormItem label="分类" prop="category_id">
                <Select v-model="myForm.category_id">
                    <Option v-for="item in categorys" :value="item.id" :key="item.id">{{ item.category_name }}</Option>
                </Select>
            </FormItem>
            <FormItem label="封面图">
                <div class="upload-list" v-for="item in imageList">
                    <template v-if="item.status === 'finished'">
                        <img :src="item.url">
                        <div class="upload-list-cover">
                            <Icon type="ios-eye-outline" @click.native="thumbView(item.url)"></Icon>
                            <Icon type="ios-trash-outline" @click.native="thumbRemove(item)"></Icon>
                        </div>
                    </template>
                    <template v-else>
                        <Progress v-if="item.showProgress" :percent="item.percentage" hide-info></Progress>
                    </template>
                </div>
                <Upload
                        ref="upload"
                        :show-upload-list="false"
                        :default-file-list="imageList"
                        :on-success="uploadsSuccess"
                        :format="['jpg','jpeg','png']"
                        :max-size="2048"
                        :on-format-error="uploadsFormatError"
                        :on-exceeded-size="uploadsMaxSize"
                        :before-upload="uploadsBeforeUpload"
                        type="drag"
                        :action="uploadsUrl"
                        class="my-thumbnail">
                    <div class="inner-camera">
                        <Icon type="camera" size="20"></Icon>
                    </div>
                </Upload>
            </FormItem>

            <FormItem label="">
                <textarea id="markdown_editor" v-model="myForm.markdown"></textarea>
            </FormItem>
            <FormItem label="">
                <Button type="ghost" @click="clearCache()">清缓存</Button>
                <Button type="ghost" @click="closeForm('myForm')">取 消</Button>
                <Button type="primary" @click="submitMyForm('myForm')">确 定</Button>
            </FormItem>
        </Form>
    </div>
</template>

<script>
    import Util from '../../libs/util';
    import inlineAttachment from '../../libs/inline-attachment';
    import localforage from 'localforage';
    import SimpleMDE from 'simplemde';

    let simpleEditor;

    export default {
        data() {
            return {
                editFormLoading: false,
                loading: false,
                myFormTitle: '编辑',
                domain: 'http://',
                newTag: '',
                textareaRow: 15,
                categorys: [],
                imageList: [],
                markdownPreviews: '',
                uploadsUrl: persimmon.baseUrl + '/uploads',
                previewVisible: false,
                showPreview: '',
                myForm: {
                    id: 0,
                    title: '',
                    flag: '',
                    thumb: '',
                    tags: [],
                    category_id: 0,
                    markdown: ''
                },
                headers: {'X-CSRF-TOKEN': persimmon.csrf},
                myRules: {
                    title: [
                        {required: true, type: "string", message: '请填写文章标题', trigger: 'blur'}
                    ],
                    flag: [
                        {required: true, type: "string", message: '请填写文章别名', trigger: 'blur'},
                        {pattern: /^[a-zA-Z0-9_-]+$/, message: '只允许英文或者拼音,横杠(-),下划线(_)', trigger: 'blur'}
                    ],
                    category_id: [
                        {required: true, type: "integer", message: '请选择分类', trigger: 'blur'}
                    ],
                    markdown: [
                        {required: true, type: "string", message: '文章内容不能为空', trigger: 'blur'}
                    ]
                },
            }
        },
        created() {
            let postId = this.$route.params.id;
            if (postId != undefined) {
                this.getPost(postId);
            }
        },
        methods: {
            thumbView(name) {
                /*
                this.imgName = name;
                this.visible = true;
                */
            },
            thumbRemove(file) {
                const fileList = this.$refs.upload.imageList;
                this.$refs.upload.imageList.splice(fileList.indexOf(file), 1);
            },
            uploadsSuccess(res, file) {
                if (res.status == 200) {
                    this.myForm.thumb = res.filename;
                } else {
                    this.myForm.thumb = '';
                }
                this.imageList.push(res);
            },
            uploadsFormatError(file) {
                this.$Notice.warning({
                    title: '文件格式错误',
                    desc: '文件 ' + file.name + ' 格式不对, 请选择 jpg 或者 png.'
                });
            },
            uploadsMaxSize(file) {
                this.$Notice.warning({
                    title: '文件过大',
                    desc: '文件  ' + file.name + ' 太大了，请选择 2M 以内的图片.'
                });
            },
            uploadsBeforeUpload() {

            },
            clearCache() {
                this.removeItem();
                this.myForm.markdown = '';
                console.log('Key is cleared!');
            },
            closeTags(tag) {
                this.myForm.tags.splice(this.myForm.tags.indexOf(tag), 1);
            },
            addTag() {
                let inputValue = this.newTag;
                if (inputValue) {
                    this.myForm.tags.push(inputValue);
                }
                this.newTag = '';
            },
            titleBlur(event) {
                let that = this;
                let query = that.myForm.title;
                if (query.match(/\w+/g) != null) {
                    that.myForm.flag = query;
                }
                if (query == null || query == '') {
                    return false;
                }
                Util.ajax.get('/util', {
                    params: {
                        action: 'translates',
                        q: query
                    }
                }).then(function (response) {
                    let res = response.data;
                    if (res.status == 200 && res.trans_result) {
                        let flag = res.trans_result.toLowerCase();
                        that.myForm.flag = flag.replaceAll(' ', '-', flag);
                    }
                }).catch(function (error) {
                    console.log(error);
                });
            },
            getPost(id) {
                if (parseInt(id) <= 0) {
                    return false;
                }
                let that = this;
                that.editFormLoading = true;
                that.myFormTitle = '编辑';
                Util.ajax.get('/posts/' + id).then(function (response) {
                    let res = response.data;
                    console.log(res);
                    if (res != false) {
                        let tags = [];
                        that.myForm = res;
                        //tags
                        if (res.tags.length > 0) {
                            for (let index in res.tags) {
                                tags.push(res.tags[index].tags_name);
                            }
                            that.myForm.tags = tags;
                        }
                        //thumb
                        if (res.thumb != '') {
                            that.imageList = [{url: res.thumb}]
                        }
                        simpleEditor.value(res.markdown);
                    } else {
                        that.$Notice.error({
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
            submitMyForm() {
                let that = this;
                that.$refs['myForm'].validate((valid) => {
                    if (!valid) {
                        console.log('myForm valid error.');
                        return false;
                    }

                    if (that.myForm.id > 0) {
                        Util.ajax.put('/posts/update', that.myForm).then(function (response) {
                            let res = response.data;
                            that.$Notice.open({
                                title: res.status == 'success' ? '编辑成功' : '编辑失败',
                                desc: ''
                            });
                            if (res.status == 'success') {
                                that.closeForm('myForm');
                            }
                        }).catch(function (error) {
                            console.log(error);
                        });
                    } else {
                        Util.ajax.post('/posts', that.myForm).then(function (response) {
                            let res = response.data;
                            if (res.status == 'success') {
                                that.closeForm('myForm');
                            }
                            that.$Notice.open({
                                title: res.status == 'success' ? '新增成功' : '新增失败',
                                desc: ''
                            });
                        }).catch(function (error) {
                            if (error.response) {
                                if (error.response.status == 422) {
                                    for (let index in error.response.data) {
                                        that.$Notice.warning({
                                            title: '警告',
                                            desc: error.response.data[index][0]
                                        });
                                    }
                                }
                            } else {
                                console.log(error);
                            }
                        });
                    }
                });
            },
            closeForm() {
                this.removeItem();
                this.$refs['myForm'].resetFields();
                this.$router.replace('/posts');
                console.log('closeForm');
            },
            removeItem(){
                localStorage.removeItem('smde_MyPersimmonMarkdown');
            },
            getCategorys() {
                let that = this;
                Util.ajax.get('/categorys', {
                    params: {
                        rows: 999
                    }
                }).then(function (response) {
                    let res = response.data;
                    if (res != false) {
                        res.data.splice(0, 0, {id: 0, category_name: '顶级分类', hidden: true, category_parent: 0});
                        that.categorys = res.data;
                    } else {
                        that.$Notice.warning({
                            title: '数据获取失败',
                            desc: ''
                        });
                    }
                }).catch(function (error) {
                    console.log(error);
                });
            },
            setDomain() {
                let location = window.location;
                this.domain = location.protocol + '//' + location.host + '/';
            },
            simpleMDE() {
                let that = this;
                simpleEditor = new SimpleMDE({
                    autofocus: true,
                    element: document.getElementById('markdown_editor'),
                    autosave: {
                        enabled: true,
                        uniqueId: "MyPersimmonMarkdown",
                        delay: 1000,
                    },
                    styleSelectedText: false,
                    renderingConfig: {
                        singleLineBreaks: true,
                        codeSyntaxHighlighting: true
                    },
                    toolbar: ['bold', 'italic', 'strikethrough', 'heading', 'heading-smaller', 'heading-bigger', 'heading-1', 'heading-2', 'heading-3', '|', 'code', 'quote', 'unordered-list', 'clean-block', '|', 'link', 'image', 'table', 'horizontal-rule', '|', 'preview', 'guide']
                });
                simpleEditor.codemirror.on("change", function () {
                    that.myForm.markdown = simpleEditor.value();
                });
            },
            inlineAtta() {
                let that = this;
                let el = that.$refs['myForm'].$el;
                let textarea = el.querySelector('.CodeMirror textarea');
                inlineAttachment.editors.input.attachToInput(textarea, {
                    uploadUrl: that.uploadsUrl,
                    uploadFieldName: 'file',
                    allowedTypes: ['image/jpeg', 'image/png', 'image/jpg', 'image/gif'],
                    progressText: '![文件上传中...]()',
                    errorText: '文件上传失败，请重试.',
                    extraParams: {},
                    extraHeaders: {
                        'X-CSRF-TOKEN': persimmon.csrf
                    },
                    onFileUploaded: function (response) {
                    }
                });
            },
        },
        watch: {
            '$route'(to, from) {
                //监听路由改变
                let postId = this.$route.params.id;
                if (postId !== undefined) {
                    this.getPost(postId);
                } else {
                    this.$refs['myForm'].resetFields();
                    this.imageList = [];
                    this.removeItem();
                    simpleEditor.value('');
                }
            }
        },
        mounted() {
            this.getCategorys();
            this.setDomain();
            this.simpleMDE();
            this.inlineAtta();
        }
    }
</script>
