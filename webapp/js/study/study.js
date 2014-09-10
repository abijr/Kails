angular.module('KailsApp', ['ngRoute', 'ngResource'])
	.config(function($routeProvider, $locationProvider) {
		$routeProvider.
		when('/', {
			templateUrl: '/study',
		});
	});
	
