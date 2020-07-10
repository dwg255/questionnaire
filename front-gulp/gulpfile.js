const gulp = require("gulp");
const htmlmin = require('gulp-htmlmin'); //html压缩
const watch = require('gulp-watch');
const less = require('gulp-less'); //处理less文件
const autoprefixer = require('gulp-autoprefixer')
const postcss = require('gulp-postcss')
const csso = require('gulp-csso'); //css压缩
const pump = require('pump');

const babel = require('gulp-babel'); //转换ES6语法
const uglify = require('gulp-uglify'); //js压缩

//处理html
gulp.task("htmlmin", async() => {
    gulp.src("./src/*.html")
        .pipe(htmlmin({ collapseWhitespace: true }))
        .pipe(gulp.dest('dist'));
})

// css任务
gulp.task('cssmin', async() => {
    // 选择css目录下的所有less文件以及css文件
    gulp.src(['./src/css/*.less', './src/css/*.css'])
        // 将less语法转换为css语法
        .pipe(less())
        //加入浏览器前缀
        .pipe(autoprefixer())
        // 将css代码进行压缩
        .pipe(csso())
        // 将处理的结果进行输出
        .pipe(gulp.dest('dist/css'));
    gulp.src(['./node_modules/normalize.css/normalize.css'])
        //加入浏览器前缀
        .pipe(autoprefixer())
        // 将css代码进行压缩
        .pipe(csso())
        // 将处理的结果进行输出
        .pipe(gulp.dest('dist/css'))
});

// js任务
// 1.es6代码转换
// 2.代码压缩
gulp.task('jsmin', async() => {
    gulp.src('./src/js/*.js')
        // .pipe(babel({
        //     // 它可以判断当前代码的运行环境 将代码转换为当前运行环境所支持的代码
        //     presets: ['@babel/env']
        // }))
        // .pipe(uglify())
        .pipe(gulp.dest('dist/js'))
});


// 复制文件夹
gulp.task('copy', async() => {
    gulp.src('./src/images/*')
        .pipe(gulp.dest('dist/images'));
});

// 构建任务
gulp.task('run', gulp.series(['htmlmin', 'cssmin', 'jsmin', 'copy']));

gulp.task('default', function() {
    gulp.watch('src/*.html', gulp.series(['htmlmin']));
    gulp.watch('src/js/*.js', gulp.series(['jsmin']));
    gulp.watch('src/css/*.less', gulp.series(['cssmin']));
});