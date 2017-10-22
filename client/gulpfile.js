const gulp = require('gulp')
const less = require('gulp-less');
const babel = require('gulp-babel');

gulp.task('default', ['less', 'babel', 'watch']);

gulp.task('less', function () {
    return gulp.src('./less/**/*.less')
        .pipe(less({
            paths: ['./less/**/*.less']
        }))
        .pipe(gulp.dest('../public/css/'));
})

gulp.task('watch', function () {
    gulp.watch('./less/**/*.less', ['less'])
    gulp.watch('./js/**/*.js', ['babel'])
})

gulp.task('babel', function () {
    gulp.src('js/**/*.js')
        .pipe(babel({
            presets: ['env']
        }))
        .pipe(gulp.dest('../public/js/'))
})

