angular.module('KailsApp', ['ngRoute', 'ngResource'])
	.config(function($routeProvider, $locationProvider) {
		$routeProvider
		.when('/', {
			templateUrl: '/webapp/program',
		})
		.when('/study/:LessonId', {
			templateUrl: '/webapp/study',
			controller: 'LessonController'
		});
		$locationProvider.html5Mode(true);
	    $locationProvider.hashPrefix('!');
	});
