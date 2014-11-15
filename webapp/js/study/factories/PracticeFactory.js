angular.module('KailsApp')
	.factory("User", function($resource) {
		return $resource('webapp/practice/:name', null, null);
	});
