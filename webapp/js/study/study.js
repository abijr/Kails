angular.module('KailsApp', ['ngRoute', 'ngResource', 'ngMessages'])
	.config(function($routeProvider, $locationProvider, $interpolateProvider) {
		$routeProvider
		.when('/', {
			templateUrl: '/webapp/program',
		})
		.when('/user/:name', {
			templateUrl: function (params) {
				return "/webapp/user/" + params.name;
			}
		})
		.when('/study/:LessonId', {
			templateUrl: '/webapp/study',
			controller: 'LessonController'
		})
		.when('/practice', {
			templateUrl: '/practice',
			controller: 'InfoController'
		});
		$locationProvider.html5Mode(true);
	    $locationProvider.hashPrefix('!');
		$interpolateProvider.startSymbol('[{');
		$interpolateProvider.endSymbol('}]');
	});
