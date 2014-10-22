angular.module('KailsApp')
	.factory("Status", function($resource) {
		return $resource("/friends/:user/:topic", null, null);
	})