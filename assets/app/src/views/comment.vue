<template>
    <div class="comments-area">
        <div id="respond" class="comment-respond">
            <h3 id="reply-title" class="comment-reply-title">发表评论</h3>
            <div class="comment-form">
                <form :model="myForm" id="commentform" class="comment-form" novalidate>
                    <p class="comment-notes">
                        <span id="email-notes">电子邮件地址不会被公开。</span> 必填项已用<span class="required">*</span>标注
                    </p>
                    <p class="comment-form-author">
                        <label for="name">姓名 <span class="required">*</span></label>
                        <input id="name" v-model="myForm.name" ref="name" type="text" value="" size="30" maxlength="245"
                               aria-required='true' required='required'/>
                    </p>
                    <p class="comment-form-email">
                        <label for="email">电子邮件 <span class="required">*</span></label>
                        <input id="email" v-model="myForm.email" ref="email" type="text" value="" size="30"
                               maxlength="100" aria-describedby="email-notes" aria-required='true' required='required'/>
                    </p>
                    <p class="comment-form-url">
                        <label for="url">站点</label>
                        <input id="url" v-model="myForm.url" ref="url" type="text" value="" size="30" maxlength="200"/>
                    </p>
                    <p class="mk-tips">Tips：支持markdown 语法 <a name="comment"></a></p>
                    <p class="comment-form-comment">
                        <label for="markdown">评论</label>
                        <textarea id="markdown" ref="markdown" v-model="myForm.markdown" cols="45" rows="8"
                                  maxlength="65525" aria-required="true" required="required"></textarea>
                    </p>
                    <p class="form-submit">
                        <button @click="comment" type="button" id="submit" class="submit" v-html="sublimtText"></button>
                    </p>
                </form>
                <div class="commentPreview" v-html="commentPreview"></div>
            </div>
        </div><!-- #respond -->
        <div class="commentshow">
            <!-- .comment-list -->
            <ol class="commentlist">

                <li class="comment" v-for="comment in comments" itemprop="reviews" itemscope
                    itemtype="https://schema.org/Review">
                    <article class="comment-body comment-body-parent">
                        <a :name="comment.id"></a>
                        <div class="comment-author">
                            <a :name="comment.md5"></a>
                            <img :src="'https://cn.gravatar.com/avatar/'+comment.md5+'?d=identicon&s=60'"
                                 class="avatar avatar-96" height="96" width="96">
                        </div>
                        <div class="comment-content">
                            <div class="comment-entry">
                                <span class="name author" itemprop="author">
                                    <a :href="comment.url" target="_blank">{{ comment.name }}：</a>
                                </span>
                                <section itemprop="reviewBody" v-html="comment.content"></section>
                            </div>
                            <div class="comment-head">
                                <span class="date">
                                    <time :datetime="comment.created_at" itemprop="datePublished">{{ comment.created_at }}</time></span>
                                <a rel='nofollow' class='comment-reply-link' href="#comment"
                                   :aria-label="'回复'+comment.name" @click="reply(comment.id,comment.name)">回复</a>
                                <!--<a class="comment-edit-link" href="javascript:void(0)">删除</a>-->
                            </div>
                        </div>
                    </article>
                </li><!-- #comment-## -->

            </ol>
        </div>
    </div>
</template>
<style lang="less">
</style>
<script>
    import marked from 'marked';
    import util from '../libs/util';

    export default {
        data: function () {
            return {
                loading: false,
                sublimtText: '发表评论',
                sublimtLoading: false,
                comments: [],
                myForm: {
                    posts_id: this.slug,
                    parent_id: 0,
                    name: '',
                    email: '',
                    url: '',
                    markdown: ''
                }
            };
        },
        props: ['slug', 'page'],
        computed: {
            commentPreview () {
                return marked(this.myForm.markdown, {sanitize: true});
            }
        },
        methods: {
            comment (event) {
                let that = this;
                for (let key in that.myForm) {
                    if(key === 'url'){
                        continue;
                    }
                    if (that.myForm[key] === '') {
                        that.$refs[key].focus();
                        that.$layer.toast({
                            icon: 'icon-check',
                            content: "必填信息不能为空",
                            time: 2000
                        });
                        return false;
                        break;
                    }
                }
                that.sublimtText = '<img class="sublimt-loading" src="data:image/svg+xml;base64,PHN2ZyB3aWR0aD0nNDBweCcgaGVpZ2h0PSc0MHB4JyB4bWxucz0iaHR0cDovL3d3dy53My5vcmcvMjAwMC9zdmciIHZpZXdCb3g9IjAgMCAxMDAgMTAwIiBwcmVzZXJ2ZUFzcGVjdFJhdGlvPSJ4TWlkWU1pZCIgY2xhc3M9InVpbC1kZWZhdWx0Ij48cmVjdCB4PSIwIiB5PSIwIiB3aWR0aD0iMTAwIiBoZWlnaHQ9IjEwMCIgZmlsbD0ibm9uZSIgY2xhc3M9ImJrIj48L3JlY3Q+PHJlY3QgIHg9JzQ5JyB5PSc0MCcgd2lkdGg9JzInIGhlaWdodD0nMjAnIHJ4PSc1JyByeT0nNScgZmlsbD0nI2ZmZmZmZicgdHJhbnNmb3JtPSdyb3RhdGUoMCA1MCA1MCkgdHJhbnNsYXRlKDAgLTMwKSc+ICA8YW5pbWF0ZSBhdHRyaWJ1dGVOYW1lPSdvcGFjaXR5JyBmcm9tPScxJyB0bz0nMCcgZHVyPScxcycgYmVnaW49JzBzJyByZXBlYXRDb3VudD0naW5kZWZpbml0ZScvPjwvcmVjdD48cmVjdCAgeD0nNDknIHk9JzQwJyB3aWR0aD0nMicgaGVpZ2h0PScyMCcgcng9JzUnIHJ5PSc1JyBmaWxsPScjZmZmZmZmJyB0cmFuc2Zvcm09J3JvdGF0ZSgzMCA1MCA1MCkgdHJhbnNsYXRlKDAgLTMwKSc+ICA8YW5pbWF0ZSBhdHRyaWJ1dGVOYW1lPSdvcGFjaXR5JyBmcm9tPScxJyB0bz0nMCcgZHVyPScxcycgYmVnaW49JzAuMDgzMzMzMzMzMzMzMzMzMzNzJyByZXBlYXRDb3VudD0naW5kZWZpbml0ZScvPjwvcmVjdD48cmVjdCAgeD0nNDknIHk9JzQwJyB3aWR0aD0nMicgaGVpZ2h0PScyMCcgcng9JzUnIHJ5PSc1JyBmaWxsPScjZmZmZmZmJyB0cmFuc2Zvcm09J3JvdGF0ZSg2MCA1MCA1MCkgdHJhbnNsYXRlKDAgLTMwKSc+ICA8YW5pbWF0ZSBhdHRyaWJ1dGVOYW1lPSdvcGFjaXR5JyBmcm9tPScxJyB0bz0nMCcgZHVyPScxcycgYmVnaW49JzAuMTY2NjY2NjY2NjY2NjY2NjZzJyByZXBlYXRDb3VudD0naW5kZWZpbml0ZScvPjwvcmVjdD48cmVjdCAgeD0nNDknIHk9JzQwJyB3aWR0aD0nMicgaGVpZ2h0PScyMCcgcng9JzUnIHJ5PSc1JyBmaWxsPScjZmZmZmZmJyB0cmFuc2Zvcm09J3JvdGF0ZSg5MCA1MCA1MCkgdHJhbnNsYXRlKDAgLTMwKSc+ICA8YW5pbWF0ZSBhdHRyaWJ1dGVOYW1lPSdvcGFjaXR5JyBmcm9tPScxJyB0bz0nMCcgZHVyPScxcycgYmVnaW49JzAuMjVzJyByZXBlYXRDb3VudD0naW5kZWZpbml0ZScvPjwvcmVjdD48cmVjdCAgeD0nNDknIHk9JzQwJyB3aWR0aD0nMicgaGVpZ2h0PScyMCcgcng9JzUnIHJ5PSc1JyBmaWxsPScjZmZmZmZmJyB0cmFuc2Zvcm09J3JvdGF0ZSgxMjAgNTAgNTApIHRyYW5zbGF0ZSgwIC0zMCknPiAgPGFuaW1hdGUgYXR0cmlidXRlTmFtZT0nb3BhY2l0eScgZnJvbT0nMScgdG89JzAnIGR1cj0nMXMnIGJlZ2luPScwLjMzMzMzMzMzMzMzMzMzMzNzJyByZXBlYXRDb3VudD0naW5kZWZpbml0ZScvPjwvcmVjdD48cmVjdCAgeD0nNDknIHk9JzQwJyB3aWR0aD0nMicgaGVpZ2h0PScyMCcgcng9JzUnIHJ5PSc1JyBmaWxsPScjZmZmZmZmJyB0cmFuc2Zvcm09J3JvdGF0ZSgxNTAgNTAgNTApIHRyYW5zbGF0ZSgwIC0zMCknPiAgPGFuaW1hdGUgYXR0cmlidXRlTmFtZT0nb3BhY2l0eScgZnJvbT0nMScgdG89JzAnIGR1cj0nMXMnIGJlZ2luPScwLjQxNjY2NjY2NjY2NjY2NjdzJyByZXBlYXRDb3VudD0naW5kZWZpbml0ZScvPjwvcmVjdD48cmVjdCAgeD0nNDknIHk9JzQwJyB3aWR0aD0nMicgaGVpZ2h0PScyMCcgcng9JzUnIHJ5PSc1JyBmaWxsPScjZmZmZmZmJyB0cmFuc2Zvcm09J3JvdGF0ZSgxODAgNTAgNTApIHRyYW5zbGF0ZSgwIC0zMCknPiAgPGFuaW1hdGUgYXR0cmlidXRlTmFtZT0nb3BhY2l0eScgZnJvbT0nMScgdG89JzAnIGR1cj0nMXMnIGJlZ2luPScwLjVzJyByZXBlYXRDb3VudD0naW5kZWZpbml0ZScvPjwvcmVjdD48cmVjdCAgeD0nNDknIHk9JzQwJyB3aWR0aD0nMicgaGVpZ2h0PScyMCcgcng9JzUnIHJ5PSc1JyBmaWxsPScjZmZmZmZmJyB0cmFuc2Zvcm09J3JvdGF0ZSgyMTAgNTAgNTApIHRyYW5zbGF0ZSgwIC0zMCknPiAgPGFuaW1hdGUgYXR0cmlidXRlTmFtZT0nb3BhY2l0eScgZnJvbT0nMScgdG89JzAnIGR1cj0nMXMnIGJlZ2luPScwLjU4MzMzMzMzMzMzMzMzMzRzJyByZXBlYXRDb3VudD0naW5kZWZpbml0ZScvPjwvcmVjdD48cmVjdCAgeD0nNDknIHk9JzQwJyB3aWR0aD0nMicgaGVpZ2h0PScyMCcgcng9JzUnIHJ5PSc1JyBmaWxsPScjZmZmZmZmJyB0cmFuc2Zvcm09J3JvdGF0ZSgyNDAgNTAgNTApIHRyYW5zbGF0ZSgwIC0zMCknPiAgPGFuaW1hdGUgYXR0cmlidXRlTmFtZT0nb3BhY2l0eScgZnJvbT0nMScgdG89JzAnIGR1cj0nMXMnIGJlZ2luPScwLjY2NjY2NjY2NjY2NjY2NjZzJyByZXBlYXRDb3VudD0naW5kZWZpbml0ZScvPjwvcmVjdD48cmVjdCAgeD0nNDknIHk9JzQwJyB3aWR0aD0nMicgaGVpZ2h0PScyMCcgcng9JzUnIHJ5PSc1JyBmaWxsPScjZmZmZmZmJyB0cmFuc2Zvcm09J3JvdGF0ZSgyNzAgNTAgNTApIHRyYW5zbGF0ZSgwIC0zMCknPiAgPGFuaW1hdGUgYXR0cmlidXRlTmFtZT0nb3BhY2l0eScgZnJvbT0nMScgdG89JzAnIGR1cj0nMXMnIGJlZ2luPScwLjc1cycgcmVwZWF0Q291bnQ9J2luZGVmaW5pdGUnLz48L3JlY3Q+PHJlY3QgIHg9JzQ5JyB5PSc0MCcgd2lkdGg9JzInIGhlaWdodD0nMjAnIHJ4PSc1JyByeT0nNScgZmlsbD0nI2ZmZmZmZicgdHJhbnNmb3JtPSdyb3RhdGUoMzAwIDUwIDUwKSB0cmFuc2xhdGUoMCAtMzApJz4gIDxhbmltYXRlIGF0dHJpYnV0ZU5hbWU9J29wYWNpdHknIGZyb209JzEnIHRvPScwJyBkdXI9JzFzJyBiZWdpbj0nMC44MzMzMzMzMzMzMzMzMzM0cycgcmVwZWF0Q291bnQ9J2luZGVmaW5pdGUnLz48L3JlY3Q+PHJlY3QgIHg9JzQ5JyB5PSc0MCcgd2lkdGg9JzInIGhlaWdodD0nMjAnIHJ4PSc1JyByeT0nNScgZmlsbD0nI2ZmZmZmZicgdHJhbnNmb3JtPSdyb3RhdGUoMzMwIDUwIDUwKSB0cmFuc2xhdGUoMCAtMzApJz4gIDxhbmltYXRlIGF0dHJpYnV0ZU5hbWU9J29wYWNpdHknIGZyb209JzEnIHRvPScwJyBkdXI9JzFzJyBiZWdpbj0nMC45MTY2NjY2NjY2NjY2NjY2cycgcmVwZWF0Q291bnQ9J2luZGVmaW5pdGUnLz48L3JlY3Q+PC9zdmc+" alt=""> 提交中...';
                that.myForm.posts_id = parseInt(that.myForm.posts_id);
                util.ajax.post('/comment', that.myForm).then(function (response) {
                    let data = response.data;
                    if (data.status === 200) {
                        that.myForm.markdown = '';
                        that.myForm.parent_id = 0;
                        util.remenberMe(that.myForm);
                        that.getData();
                    } else {
                        that.$layer.toast({
                            icon: 'icon-check',
                            content: data.info,
                            time: 2000
                        });
                    }
                    that.sublimtText = '发表评论';
                }).catch(function (error) {
                    console.log(error);
                    that.sublimtText = '发表评论';
                });
                return false;
            },
            getData () {
                let that = this;
                util.ajax.get('/comment/' + that.slug + '/' + that.page).then(function (response) {
                    that.comments = response.data.data;
                }).catch(function (error) {
                    console.log(error);
                });
            },
            reply (id, name) {
                let at = '@' + name + ' ';
                this.myForm.markdown += at;
                this.myForm.parent_id = id;
                return false;
            }
        },
        mounted () {
            let viewer = util.getViewer();
            this.myForm.posts_id = this.slug;
            this.myForm.name = viewer.name;
            this.myForm.email = viewer.email;
            this.myForm.url = viewer.url;
            this.getData();
        }
    };
</script>
