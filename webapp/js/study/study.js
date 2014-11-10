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
		.when('/videochat', {
			templateUrl: '/webapp/videochat',
			controller: 'VideochatController'
		})
		.when('/settings', {
			templateUrl: '/webapp/settings',
			// controller: 'VideochatController'
		})
		.when('/practice', {
			templateUrl: '/webapp/practice',
			controller: 'PracticeController'
		})
		.when('/words', {
			templateUrl: '/webapp/words',
			controller: 'WordsController'
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
