angular.module('KailsApp', ['ngRoute', 'ngResource'])
	.config(function($routeProvider, $locationProvider) {
		$routeProvider.
		when('/', {
			templateUrl: 'partial/program',
		})
		.when('/study/:LessonId', {
			templateUrl: 'partial/study',
			controller: 'LessonController'
		});
	});
