angular.module('KailsApp')
	.factory("Friends", function($resource) {
        return $resource('/friends/:user', null, {
            'get': {
                method: "GET",
                isArray: true
            }
        });
	})
