module.exports = function(grunt) {
	grunt.initConfig({
		pkg: grunt.file.readJSON('package.json'),

		// Esta operacion genera el css desde lo que se encuentra
		// en el archivo scss/app.scss
		sass: {
			options: {
				includePaths: ['bower_components/foundation/scss', 'bower_components/fontawesome/scss/']
			},
			dist: {
				options: {
					outputStyle: 'nested'
				},
				files: {
					'dist/css/app.css': 'scss/app.scss',
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
					'fastclick/lib/fastclick.js',
					'modernizr/modernizr.js',
					'angular/angular.min.js',
					'angular/angular.js',
					'angular-resource/angular-resource.min.js',
					'angular-resource/angular-resource.js',
					'angular-route/angular-route.min.js',
					'angular-route/angular-route.js',
					'angular-messages/angular-messages.min.js',
					'angular-messages/angular-messages.js',
				],
				dest: 'dist/js/',
				flatten: true,
				filter: 'isFile',
			},
			fonts: {
				expand: true,
				cwd: 'bower_components/fontawesome/fonts/',
				src: ["**"],
				dest: 'dist/fonts/'
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
				src: ['js/study/study.js', 'js/study/*/*.js', 'js/webrtc/*/*.js'],
				dest: 'dist/js/study_app.js'
			},
			com: {
				src: ['js/webrtc/*.js', 'js/webrtc/*/*.js'],
				dest: 'dist/js/com.js'
			}
		},

		nodemon: {
			easyrtc: {
				script: "server.js"
			}
		},

		// Observa los archivos para detectar cambios
		// ejecuta las tareas correspondientes.
		watch: {
			grunt: {
				files: ['Gruntfile.js']
			},

			sass: {
				files: 'scss/**/*.scss',
				tasks: ['sass']
			},

			study: {
				files: ['js/study/**/*.js'],
				tasks: ['concat:study']
			},
			webrtc: {
				files: "js/webrtc/**/*.js",
				tasks: ['concat:com']
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
