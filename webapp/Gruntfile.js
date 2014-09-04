module.exports = function(grunt) {
  grunt.initConfig({
    pkg: grunt.file.readJSON('package.json'),

    sass: {
      options: {
        includePaths: ['bower_components/foundation/scss']
      },
      dist: {
        options: {
          outputStyle: 'compressed'
        },
        files: {
          'dist/css/app.css': 'scss/app.scss'
        }
      }
    },

    copy: {
        js: {
            expand: true,
            cwd: 'bower_components/',
            src: [
                'jquery/dist/jquery.min.js',
                'foundation/js/foundation.min.js',
                'modernizr/modernizr.js',
                'angular/angular.min.js',
                ],
            dest: 'dist/js/',
            flatten: true,
            filter: 'isFile',
        }
    },

    concat: {
        study: {
            src: ['js/study/study.js','js/study/*/*.js'],
            dest: 'dist/js/study_app.js'
        }
    },

    watch: {
        grunt: { files: ['Gruntfile.js'] },

        sass: {
            files: 'scss/**/*.scss',
            tasks: ['sass']
      }
    }
  });

  grunt.loadNpmTasks('grunt-sass');
  grunt.loadNpmTasks('grunt-contrib-watch');
  grunt.loadNpmTasks('grunt-contrib-copy');
  grunt.loadNpmTasks('grunt-contrib-concat');

  grunt.registerTask('build', ['sass','copy','concat']);
  grunt.registerTask('default', ['build','watch']);
}
