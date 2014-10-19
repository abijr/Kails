angular.module('KailsApp', ['ngRoute', 'ngResource', 'ngMessages'])
	.config(function($routeProvider, $locationProvider, $interpolateProvider) {
		$routeProvider
		.when('/', {
			templateUrl: '/webapp/program',
			controller: 'ProgramController'
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
		.when('/flashcard', {
			templateUrl: '/webapp/flashcard',
			controller: 'FlashcardController'
		})
		.when('/practice', {
			templateUrl: '/webapp/practice',
			controller: 'PracticeController'
		});
		$locationProvider.html5Mode(true);
	    $locationProvider.hashPrefix('!');
		$interpolateProvider.startSymbol('[{');
		$interpolateProvider.endSymbol('}]');
	});
