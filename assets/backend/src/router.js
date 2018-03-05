import Main from './views/Main.vue';

// 不作为Main组件的子页面展示的页面单独写，如下
export const loginRouter = {
    path: '/login',
    name: 'login',
    meta: {
        title: 'Login - 登录'
    },
    component: resolve => {
        require(['./views/login.vue'], resolve);
    }
};

// 作为Main组件的子页面展示但是不在左侧菜单显示的路由写在otherRouter里
export const otherRouter = {
    path: '/',
    name: 'otherRouter',
    redirect: '/home',
    component: Main,
    children: [
        {
            path: 'home',
            title: '首页',
            name: 'home_index',
            component: resolve => {
                require(['./views/home/home.vue'], resolve);
            }
        },
        {
            path: 'ownspace',
            title: '个人中心',
            name: 'ownspace_index',
            component: resolve => {
                require(['./views/own-space/own-space.vue'], resolve);
            }
        },
        {
            path: 'posts/edit/:id',
            title: '文章编辑',
            name: 'edit_post',
            component: resolve => {
                require(['./views/posts/post-edit.vue'], resolve);
            }
        }
    ]
};

// 作为Main组件的子页面展示并且在左侧菜单显示的路由写在appRouter里
export const appRouter = [
    {
        path: '/dashboard',
        icon: 'home',
        name: 'dashboard',
        title: '首页',
        component: Main,
        children: [
            {
                path: 'home',
                icon: 'home',
                title: '首页',
                name: 'home',
                component: resolve => {
                    require(['./views/home/home.vue'], resolve);
                }
            }
        ]
    },
    {
        path: '/posts',
        icon: 'document-text',
        name: 'posts',
        title: '文章管理',
        component: Main,
        children: [
            {
                path: 'post-list',
                title: '文章管理',
                name: 'artical-list',
                icon: 'document',
                component: resolve => {
                    require(['./views/posts/post-list.vue'], resolve);
                }
            },
            {
                path: 'post-add',
                title: '文章发布',
                name: 'post-add',
                icon: 'compose',
                component: resolve => {
                    require(['./views/posts/post-edit.vue'], resolve);
                }
            },
            {
                path: 'categorys',
                title: '文章分类',
                name: 'categorys',
                icon: 'navicon',
                component: resolve => {
                    require(['./views/posts/categorys.vue'], resolve);
                }
            },
            {
                path: 'post-trash',
                title: '回收站',
                name: 'post-trash',
                icon: 'ios-trash',
                component: resolve => {
                    require(['./views/posts/post-trash.vue'], resolve);
                }
            },
            {
                path: 'post-tags',
                title: '文章标签',
                name: 'post-tags',
                icon: 'pricetags',
                component: resolve => {
                    require(['./views/posts/post-tags.vue'], resolve);
                }
            }
        ]
    },
    {
        path: '/extends',
        icon: 'hammer',
        name: 'extends',
        title: '扩展',
        component: Main,
        children: [
            {
                path: 'navigations',
                title: '导航管理',
                name: 'navigations',
                icon: 'ios-navigate-outline',
                component: resolve => {
                    require(['./views/plugins/navigations.vue'], resolve);
                }
            },
            {
                path: 'links',
                title: '左邻右舍',
                name: 'links',
                icon: 'link',
                component: resolve => {
                    require(['./views/plugins/links.vue'], resolve);
                }
            },
            {
                path: 'comments',
                title: '评论管理',
                name: 'comments',
                icon: 'chatboxes',
                component: resolve => {
                    require(['./views/plugins/comments.vue'], resolve);
                }
            },
            {
                path: 'options',
                title: '配置项目',
                name: 'options',
                icon: 'ios-settings',
                component: resolve => {
                    require(['./views/plugins/options.vue'], resolve);
                }
            }
        ]
    },
    {
        path: '/settings',
        icon: 'settings',
        name: 'settings',
        title: '系统设置',
        component: Main,
        children: [
            {
                path: 'setting',
                title: '系统设置',
                name: 'setting',
                component: resolve => {
                    require(['./views/plugins/settings.vue'], resolve);
                }
            }
        ]
    }
];

// 所有上面定义的路由都要写在下面的routers里
export const routers = [
    loginRouter,
    otherRouter,
    ...appRouter,
];
