const mix = require('laravel-mix');

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
// mix.options({
//     hmrOptions: {
//         host: 'http://vfs.test',  // site's host name
//     }
// });

mix.webpackConfig({
    // add any webpack dev server config here
    devServer: {
        proxy: {
            target: 'http://zmdevops.com:8000',  // host machine ip

            changeOrigin: true
        },
        watchOptions:{
            aggregateTimeout:200,
            poll:5000
        },

    }
});

mix.js('resources/js/app.js', 'dist/js')
    .sass('resources/sass/app.scss', 'dist/css');
