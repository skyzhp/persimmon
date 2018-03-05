<template>
    <div class="pit-post-form" style="width: 60%">
        <Form ref="myForm" :model="myForm" :label-width="150"
                 class="pit-common">
            <FormItem v-for="option in options" :value="option.id" :key="option.id" :label="option.option_title">
                <Input :type="option.data_type" v-model="myForm[option.option_name]"></Input>
            </FormItem>
            <FormItem>
                <Button @click="closeForm('myForm')">取 消</Button>
                <Button type="primary" @click="submitMyForm('myForm')">确 定</Button>
            </FormItem>
        </Form>
        <Spin size="large" fix v-if="editFormLoading"></Spin>
    </div>
</template>
<style type="text/css">

</style>
<script>
    import Util from '../../libs/util';

    export default{
        data(){
            return {
                editFormLoading: false,
                options: [],
                myForm: {}
            }
        },
        created () {
            this.getData();
        },
        methods: {
            getData: function () {
                let that = this;
                that.editFormLoading = true;
                Util.ajax.get('/settings').then(function (response) {
                    let res = response.data;
                    if (res != false) {
                        if (res.length > 0) {
                            for (var index in res) {
                                that.myForm[res[index].option_name] = res[index].option_value;
                            }
                            that.options = res;
                        }
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

            submitMyForm: function () {
                let that = this;
                Util.ajax.put('/settings/update', that.myForm).then(function (response) {
                    let res = response.data;
                    that.$Notice.warning({
                        title: res.status == 'success' ? '更新成功' : '更新失败',
                        desc: ''
                    });
                }).catch(function (error) {
                    console.log(error);
                });
            },
        },
        watch: {},
        mounted() {
        }
    }
</script>
