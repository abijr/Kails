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
		.when('/chat', {
			templateUrl: '/webapp/chat',
			controller: 'ChatController'
		})
		.when('/practice', {
			templateUrl: '/webapp/practice',
			controller: 'PracticeController'
		})
		.when('/friends', {
			templateUrl: '/webapp/friends',
			controller: 'FriendsController'
		});
		$locationProvider.html5Mode(true);
	    $locationProvider.hashPrefix('!');
		$interpolateProvider.startSymbol('[{');
		$interpolateProvider.endSymbol('}]');
	});
