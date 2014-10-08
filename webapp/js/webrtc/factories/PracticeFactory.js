angular.module('KailsApp')
	.factory("User", function($resource) {
		return $resource('/practice/:name', null, null)
	});