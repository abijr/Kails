angular.module('KailsApp')
	.factory("Status", function($resource) {
		return $resource("/friends/:user/:topic", null, null);
	})
	.factory("Connected", function($resource) {
		return $resource("/friends/connected", null, {
			'get': {
				method: "GET",
				isArray: true
			}
		});
	});
