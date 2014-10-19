angular.module('KailsApp')
	.factory("Friends", function($resource) {
		return $resource("/friends/:topic", null, {
			'get': {
				method: "GET",
				isArray: true
			}
		});
	})