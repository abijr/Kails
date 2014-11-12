angular.module('KailsApp')
	.factory("Status", function($resource) {
		return $resource("/friends/connected", null, null);
	});
