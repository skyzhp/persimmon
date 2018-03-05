<template>
    <div class="pit-content">
        <div class="pit-action-btn">
            <!--<Button type="primary" @click="handleCreate" icon="plus">新增</Button>-->
            <Button type="primary" @click="handleDistory('multi',{})" icon="delete">删除</Button>
        </div>

        <div class="data-list">
            <Table :loading="listLoading" :columns="tableColumns" :data="listData" stripe @on-select="selectRow" @on-select-all="selectAll"></Table>
        </div>
        <div style="margin: 10px;overflow: hidden">
            <div style="float: right;">
                <Page :total="total"
                      @on-page-size-change="handleSizeChange"
                      @on-change="handleCurrentChange"
                      :page-size="pageSize"
                      :current="currentPage"
                      :page-size-opts="sizeOpts"
                      :show-elevator="true"
                      :show-sizer="true"
                      class="myp-page">
                </Page>
            </div>
        </div>

        <Modal :title="myFormTitle" v-model="editFormVisible" @on-ok="submitMyForm" @on-cancel="closeForm">
            <div class="pit-dialog-edit-form">
                <Form ref="myForm" :rules="myRules" class="myForm" :label-width="100" :model="myForm">
                    <FormItem label="名称" prop="name">
                        <Input v-model="myForm.name" auto-complete="off"></Input>
                    </FormItem>
                    <FormItem label="链接">
                        <Input v-model="myForm.url" auto-complete="off"></Input>
                    </FormItem>
                    <FormItem label="E-Mail" prop="email">
                        <Input v-model="myForm.email" auto-complete="off"></Input>
                    </FormItem>
                    <FormItem label="内容" prop="markdown">
                        <Input type="textarea" autosize v-model="myForm.markdown" auto-complete="off"></Input>
                    </FormItem>
                    <FormItem v-if="myForm.id">
                        <Input v-model="myForm.id" style="display: none;"></Input>
                    </FormItem>
                </Form>
                <Spin size="large" fix v-if="editFormLoading"></Spin>
            </div>
        </Modal>
    </div>
</template>
<style lang="less">
    .links {
        text-decoration: none;
        color: #000
    }

    .myForm {
        width: 100% !important;
    }
</style>
<script>
    import Util from '../../libs/util';

    export default {
        data() {
            return {
                listData: [],
                categorys: [],
                currentPage: 1,
                total: 0,
                pageSize: 20,
                myForm: {
                    id: 0,
                    name: '',
                    url: '',
                    email: '',
                    markdown: ''
                },
                myRules: {
                    name: [
                        {required: true, type: "string", message: '请填写链接名称', trigger: 'blur'}
                    ],
                    email: [
                        {required: true, type: "string", message: '请填写邮箱', trigger: 'blur'}
                    ],
                    markdown: [
                        {required: true, type: "string", message: '请填写评论内容', trigger: 'blur'}
                    ]
                },
                editFormVisible: false,
                editFormLoading: false,
                listLoading: false,
                myFormTitle: '编辑',
                checkedAll: [],
                tableColumns: [
                    {
                        type: 'selection',
                        width: 60,
                        align: 'center'
                    },
                    {
                        title: '名称',
                        key: 'title',
                        render: (h, params) => {
                            return h('div', [
                                h('div', {
                                    attrs: {
                                        'class': 'comment-avatar'
                                    }
                                }, [
                                    h('p', [
                                        h('img', {
                                            attrs: {
                                                'src': 'https://cn.gravatar.com/avatar/' + params.row.md5 + '?d=identicon&s=60'
                                            }
                                        })
                                    ])
                                ]),
                                h('div', {
                                    attrs: {
                                        'class': 'comment-item'
                                    }
                                }, [
                                    h('p',{}, [
                                        h('a', {
                                            attrs: {
                                                'href': params.row.url,
                                            }
                                        },params.row.name)
                                    ]),
                                    h('p', {},[
                                        h('a', {
                                            attrs: {
                                                'href': 'mailto:' + params.row.email,
                                            }
                                        },params.row.email),
                                    ]),
                                    h('p', params.row.ipaddress)
                                ])
                            ]);
                        }
                    },
                    {
                        title: '评论内容',
                        key: 'content',
                        render: (h, params) => {
                            return h('div', {
                                domProps: {
                                    innerHTML: params.row.content
                                }
                            });
                        }
                    },
                    {
                        title: '发表于',
                        key: 'posts',
                        render: (h, params) => {
                            return h('span', params.row.post.title);
                        }
                    },
                    {
                        title: '日期',
                        key: 'created_at'
                    },
                    {
                        title: '操作',
                        key: 'action',
                        width: 150,
                        align: 'center',
                        render: (h, params) => {
                            return h('div', [
                                h('Button', {
                                    props: {
                                        type: 'primary',
                                        size: 'small'
                                    },
                                    style: {
                                        marginRight: '5px'
                                    },
                                    on: {
                                        click: () => {
                                            this.handleEdit(params.row)
                                        }
                                    }
                                }, '编辑'),
                                h('Button', {
                                    props: {
                                        type: 'error',
                                        size: 'small'
                                    },
                                    on: {
                                        click: () => {
                                            this.handleDistory('one', params.row);
                                        }
                                    }
                                }, '删除')
                            ]);
                        }
                    }
                ],
                sizeOpts: [10, 20, 30, 40, 50, 60, 70, 80, 90, 100],
            }
        },
        methods: {
            getData: function () {
                let that = this;
                that.listLoading = true;
                Util.ajax.get('/comments', {
                    params: {
                        rows: that.pageSize,
                        page: this.currentPage
                    }
                }).then(function (response) {
                    let res = response.data;
                    if (res != false) {
                        that.listData = res.data;
                        that.total = res.total;
                        that.currentPage = res.current_page;
                        that.listLoading = false;
                    } else {
                        that.$Notice.warning({
                            message: '数据获取失败',
                            type: 'error',
                            duration: 3 * 1000
                        });
                    }
                }).catch(function (error) {
                    console.log(error);
                });
            },
            handleSizeChange(val) {
                //console.log(`每页 ${val} 条`);
                this.pageSize = val;
                this.getData();
            },
            handleCurrentChange(val) {
                this.currentPage = val;
                this.getData();
                //console.log(`当前页: ${val}`);
            },
            handleCreate: function () {
                let that = this;
                that.myFormTitle = '新增';
                that.myForm.id = 0;
                that.editFormVisible = true;
            },
            handleEdit: function (row) {
                let that = this;
                that.editFormLoading = true;
                that.myFormTitle = '编辑';
                that.editFormVisible = true;
                Util.ajax.get('/comments/' + row.id).then(function (response) {
                    let res = response.data;
                    if (res != false) {
                        that.myForm = res;
                    } else {
                        that.$Notice.warning({
                            message: '数据获取失败',
                            type: 'error'
                        });
                    }
                    that.editFormLoading = false;
                }).catch(function (error) {
                    console.log(error);
                    that.editFormLoading = false;
                });
            },
            handleDistory: function (type, row) {
                let that = this, idsParam = {};
                switch (type) {
                    case 'one':
                        if (parseInt(row.id) <= 0) {
                            that.$Notice.warning({
                                message: '请选择需要删除的数据',
                                type: 'warning'
                            });
                            return false;
                        }
                        idsParam = {ids: [row.id]};
                        break;
                    case 'multi':
                        let ids = Util.getIdByArr(that.checkedAll);
                        if (ids.length <= 0) {
                            that.$Notice.warning({
                                message: '请选择需要删除的数据',
                                type: 'warning'
                            });
                            return false;
                        }
                        idsParam = {ids: ids};
                        break;
                    default:
                        break;
                }

                that.$Modal.confirm({
                    title: '确认删除选中的记录吗?',
                    content: '<p>您确认删除选中的记录吗?</p>',
                    onOk: () => {
                        that.listLoading = true;
                        Util.ajax.delete('/comments/destroy', {data: idsParam}).then(function (response) {
                            that.listLoading = false;
                            let res = response.data;
                            that.$Notice.warning({
                                title: res.status == 'success' ? '删除成功' : '删除失败',
                                desc: ''
                            });
                            if (type == 'one') {
                                Util.removeByValue(that.listData, row.id);
                            } else {
                                for (let index in that.checkedAll) {
                                    Util.removeByValue(that.listData, that.checkedAll[index].id);
                                }
                            }

                        }).catch(function (error) {
                            console.log(error);
                        });
                    },
                    onCancel: () => {
                        that.listLoading = false;
                    }
                });
            },
            submitMyForm: function () {
                let that = this;
                that.$refs['myForm'].validate((valid) => {
                    if (!valid) {
                        console.log('myForm valid error.');
                        return false;
                    }

                    if (that.myForm.id > 0) {
                        Util.ajax.put('/comments/update', that.myForm).then(function (response) {
                            let res = response.data;
                            that.$Notice.open({
                                title: res.status == 'success' ? '编辑成功' : '编辑失败',
                                desc: ''
                            });
                            if (res.status == 'success') {
                                that.closeForm('myForm');
                                that.getData();
                            }
                        }).catch(function (error) {
                            console.log(error);
                        });
                    } else {
                        Util.ajax.post('/links', that.myForm).then(function (response) {
                            let res = response.data;
                            if (res.status == 'success') {
                                that.closeForm('myForm');
                                that.getData();
                            }
                            that.$Notice.warning({
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
            closeForm: function () {
                this.editFormVisible = false;
                this.$refs['myForm'].resetFields();
                this.myForm = {
                    id: 0,
                    name: '',
                    url: '',
                    logo: '',
                    group: '',
                };
                console.log('closeForm');
            },
            selectRow(row) {
                this.checkedAll = row;
            },
            selectAll(selection) {
                this.checkedAll = selection;
            },
        },
        watch: {},
        mounted() {
            this.getData();
        }
    }
</script>