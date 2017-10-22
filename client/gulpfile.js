var gulp  = require('gulp')
var path = require('path');
var less = require('gulp-less');

gulp.task('default', ['less', 'watch']);

gulp.task('less', function () {
    return gulp.src('./less/**/*.less')
        .pipe(less({
            paths: [ './less/**/*.less' ]
        }))
        .pipe(gulp.dest('../public/css/'));
})

gulp.task('watch', function () {
  gulp.watch('./less/**/*.less', ['less'])
})