module.exports = function(grunt) {
	grunt.initConfig({
		pkg: grunt.file.readJSON('package.json'),

		// Esta operacion genera el css desde lo que se encuentra
		// en el archivo scss/app.scss
		sass: {
			options: {
				includePaths: ['bower_components/foundation/scss']
			},
			dist: {
				options: {
					outputStyle: 'nested'
				},
				files: {
					'dist/css/app.css': 'scss/app.scss',
					'dist/css/webrtc.css': 'scss/webrtc.scss'
				}
			}
		},

		// Copia librerias de javascript (manejadas por bower)
		// a el directorio dist/js
		copy: {
			js: {
				expand: true,
				cwd: 'bower_components/',
				src: [
					'jquery/dist/jquery.min.js',
					'foundation/js/foundation.min.js',
					'modernizr/modernizr.js',
					'angular/angular.min.js',
					'angular-resource/angular-resource.min.js',
					'angular-route/angular-route.min.js',
				],
				dest: 'dist/js/',
				flatten: true,
				filter: 'isFile',
			},
			webrtc: {
				expand: true,
				cwd: 'js/webrtc/',
				src: [
					'chat.js',
					'communication.js',
					'ui.js',
					'videochat.js'
				],
				dest: 'dist/js/',
				flatten: true,
				filter: 'isFile',
			},
			easyrtc: {
				expand: true,
				cwd: 'node_modules/easyrtc/api/',
				src: ["**"],
				dest: 'dist/easyrtc/'
			},
			socketio: {
				expand: true,
				cwd: 'node_modules/socket.io/node_modules/socket.io-client/dist/',
				src: ["**"],
				dest: 'dist/socket.io/'

			},
			express: {
				expand: true,
				cwd: 'node_modules/express/lib/',
				src:["**"],
				dest: 'dist/express/'
			}
		},

		// Genera un solo archivo conteniendo la aplicacion
		// de `study`
		concat: {
			study: {
				src: ['js/study/study.js', 'js/study/*/*.js'],
				dest: 'dist/js/study_app.js'
			}
		},

		nodemon: {
			easyrtc: {
				script: "server.js"
			}
		},

		watch: {
			grunt: {
				files: ['Gruntfile.js']
			},

			sass: {
				files: 'scss/**/*.scss',
				tasks: ['sass']
			},

			study: {
				files: 'js/study/**/*.js',
				tasks: ['concat:study']
			}
		},

		concurrent: {
			dev: {
				tasks: ["watch", "nodemon"],
				options: {
					logConcurrentOutput: true
				}
			}
		}
	});

	grunt.loadNpmTasks('grunt-sass');
	grunt.loadNpmTasks('grunt-nodemon');
	grunt.loadNpmTasks('grunt-concurrent');
	grunt.loadNpmTasks('grunt-contrib-watch');
	grunt.loadNpmTasks('grunt-contrib-copy');
	grunt.loadNpmTasks('grunt-contrib-concat');

	grunt.registerTask('build', ['sass', 'copy', 'concat']);
	grunt.registerTask('default', ['build', 'concurrent']);
};
