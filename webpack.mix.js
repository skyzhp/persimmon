let mix = require('laravel-mix');

/*
 |--------------------------------------------------------------------------
 | Mix Asset Management
 |--------------------------------------------------------------------------
 |
 | Mix provides a clean, fluent API for defining some Webpack build steps
 | for your Laravel application. By default, we are compiling the Sass
 | file for the application as well as bundling up all the JS files.
 |
 */

/**
 * App
 */
/*
mix.js('assets/app/src/app.js', 'public/js/app.js')
    .sass('assets/app/src/styles/app.scss', 'public/css/app.css')
    .options({processCssUrls: false})
    .copyDirectory('node_modules/font-awesome/fonts', 'public/fonts')
    .copyDirectory('assets/app/src/images', 'public/images');
*/


/**
 * Backend
 */

//mix.setPublicPath('public/backend/').setResourceRoot('public/backend/');

mix.webpackConfig({
    module: {
        rules: [
            {
                test: /\.(woff2?|ttf|eot|svg|otf)$/,
                loader: 'file-loader',
                options: {
                    name: path => {
                        return Config.fileLoaderDirs.fonts + '/[name].[ext]?[hash]';
                    },
                    publicPath: Config.resourceRoot
                }
            },
            {
                test: /\.js$/,
                loader: 'webpack-replace-loader',
                options: {
                    arr: [
                        {
                            search: 'maxcdn.bootstrapcdn.com/font-awesome/latest/css',
                            replace: 'cdn.bootcss.com/font-awesome/4.7.0/css',
                            attr: 'g'
                        },
                        {
                            search: 'maxcdn.bootstrapcdn.com/font-awesome',
                            replace: 'cdn.bootcss.com/font-awesome',
                            attr: 'g'
                        }
                    ]
                }
            }
        ]
    },
    output: {
        chunkFilename: mix.inProduction() ? 'backend/js/[name].[chunkhash].js' : 'backend/js/[name].js'
    }
});

mix.js('assets/backend/src/main.js', 'public/backend/js/main.js')
    .styles([
        'assets/backend/src/styles/common.css',
        'node_modules/iview/dist/styles/iview.css',
        'assets/backend/src/styles/simplemde.min.css'
    ], 'public/backend/css/main.css')
    .copyDirectory('assets/backend/src/styles/fonts', 'public/backend/css/fonts')
    .copyDirectory('assets/backend/src/images', 'public/backend/images');


if (mix.inProduction()) {
    mix.version();
}