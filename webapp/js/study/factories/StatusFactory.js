angular.module('KailsApp')
	.factory("Status", function($resource) {
		return $resource("/webapp/friends/:user/:topic", null, null);
	})
	.factory("Connected", function($resource) {
		return $resource("/webapp/friends/connected", null, {
			'get': {
				method: "GET",
				isArray: true
			}
		});
	});
