angular.module('KailsApp')
	.factory("Friends", function($resource) {
		return $resource("/friends/:user", null, null);
	})