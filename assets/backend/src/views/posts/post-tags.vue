<style lang="less">

</style>

<template>
    <div class="pit-content">
        <div class="pit-action-btn">
            <Button type="primary" @click="handleCreate" icon="plus">新增</Button>
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
                    <FormItem label="标签名称" prop="tags_name">
                        <Input v-model="myForm.tags_name" auto-complete="off" @on-blur="translateWords" @on-enter="translateWords"></Input>
                    </FormItem>
                    <FormItem label="标签别名" prop="tags_flag">
                        <Input v-model="myForm.tags_flag" auto-complete="off"></Input>
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

<script>
    import util from '../../libs/util';

    export default{
        data(){
            return {
                listData: [],
                categorys: [],
                currentPage: 1,
                total: 0,
                pageSize: 20,
                myForm: {
                    id: 0,
                    tags_name: '',
                    tags_flag: ''
                },
                myRules: {
                    tags_name: [
                        {required: true, type: "string", message: '请填写标签名称', trigger: 'blur'}
                    ],
                    tags_flag: [
                        {required: true, type: "string", message: '请填写标签别名', trigger: 'blur'}
                    ]
                },
                editFormVisible: false,
                editFormLoading: false,
                listLoading: false,
                myFormTitle: '编辑',
                checkedAll: [],
                sizeOpts: [10, 20, 30, 40, 50, 60, 70, 80, 90, 100],
                tableColumns: [
                    {
                        type: 'selection',
                        width: 60,
                        align: 'center'
                    },
                    {
                        title: '标签名',
                        key: 'tags_name'
                    },
                    {
                        title: 'uri',
                        key: 'tags_flag'
                    },
                    {
                        title: '日期',
                        key: 'created_at',
                        render: (h, params) => {
                            let time = util.timeFormat(params.row.created_at);
                            return h('span', time);
                        }
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
            }
        },
        methods: {
            formatterFlag: function (row, column) {
                if (row.tags_flag == '') {
                    return '';
                }
                return decodeURI(row.tags_flag);
            },
            getData: function () {
                let that = this;
                that.listLoading = true;
                util.ajax.get('/backend/tags', {
                    params: {
                        rows: this.pageSize,
                        page: this.currentPage
                    }
                }).then(function (response) {
                    let res = response.data;
                    if (res.status == 200) {
                        that.listData = res.list.data;
                        that.total = res.list.total;
                        that.currentPage = res.list.current_page;
                        that.listLoading = false;
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
            handleSizeChange(val) {
                //console.log(`每页 ${val} 条`);
                this.pageSize = val;
                this.getData();
            },
            handleCurrentChange(val) {
                this.currentPage = val;
                //console.log(`当前页: ${val}`);
                this.getData();
            },
            handleCreate: function () {
                let that = this;
                that.myFormTitle = '新增';
                that.myForm.id = 0;
                that.editFormVisible = true;
                that.setTopCategorys();
            },
            handleEdit: function (row) {
                let that = this;
                that.setTopCategorys();
                that.editFormLoading = true;
                that.myFormTitle = '编辑';
                that.editFormVisible = true;
                util.ajax.get('/backend/tags/' + row.id).then(function (response) {
                    let res = response.data;
                    if (res.status == 200) {
                        res.tags_flag = decodeURI(res.item.tags_flag);
                        that.myForm = res.item;
                    } else {
                        that.$Notice.warning({
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
            handleDistory: function (type, row) {
                let that = this, idsParam = {};
                switch (type) {
                    case 'one':
                        if (parseInt(row.id) <= 0) {
                            that.$Notice.warning({
                                title: '请选择需要删除的数据',
                                desc: ''
                            });
                            return false;
                        }
                        idsParam = {ids: [row.id]};
                        break;
                    case 'multi':
                        let ids = that.util.getIdByArr(that.checkedAll);
                        if (ids.length <= 0) {
                            that.$Notice.warning({
                                title: '请选择需要删除的数据',
                                desc: ''
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
                        util.ajax.post('/backend/tags/destroy', util.stringify(idsParam)).then(function (response) {
                            that.listLoading = false;
                            let res = response.data;
                            that.$Notice.open({
                                title: res.status == 200 ? '删除成功' : '删除失败',
                                desc: ''
                            });
                            if (type == 'one') {
                                that.util.removeByValue(that.listData, row.id);
                            } else {
                                for (let index in that.checkedAll) {
                                    that.util.removeByValue(that.listData, that.checkedAll[index].id);
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
                        util.ajax.put('/backend/tags/update', that.myForm).then(function (response) {
                            let res = response.data;
                            that.$message({
                                message: res.status == 200 ? '编辑成功' : '编辑失败',
                                type: res.status
                            });
                            if (res.status == 200) {
                                that.closeForm('myForm');
                                that.getData();
                            }
                        }).catch(function (error) {
                            console.log(error);
                        });
                    } else {
                        util.ajax.post('/backend/tags/store', that.myForm).then(function (response) {
                            //console.log(response);
                            let res = response.data;
                            if (res.status == 200) {
                                that.closeForm('myForm');
                                that.getData();
                            }
                            that.$Notice.open({
                                title: res.status == 200 ? '新增成功' : '新增失败',
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
                    tags_name: '',
                    tags_flag: ''
                };
                console.log('closeForm');
            },
            selectRow(row) {
                this.checkedAll = row;
            },
            selectAll(selection) {
                this.checkedAll = selection;
            },
            setTopCategorys: function () {
                let categorys = this.listData.concat();
                categorys.splice(0, 0, {id: 0, category_name: '顶级分类', hidden: true, category_parent: 0});
                this.categorys = categorys;
            },
            translateWords(event) {
                let that = this;
                let query = that.myForm.tags_name;
                if (query.match(/\w+/g) != null) {
                    that.myForm.flag = query;
                }
                if (query == null || query == '') {
                    return false;
                }
                util.ajax.get('/backend/utils/fanyi/' + query).then(function (response) {
                    let res = response.data;
                    if (res.status == 200) {
                        that.myForm.tags_flag = res.item
                    }
                }).catch(function (error) {
                    console.log(error);
                });
            },
        },
        watch: {
            'myForm.tags_name': {//监听路由改变
                handler(curVal, oldVal){
                    this.myForm.tags_flag = curVal.replace(' ', '-');
                },
                deep: true
            }
        },
        mounted() {
            this.getData();
        }
    }
</script>
