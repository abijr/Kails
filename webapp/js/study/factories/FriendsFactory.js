angular.module('KailsApp')
	.factory("Friends", function($resource) {
        return $resource('/webapp/friends/:user', null, {
            'get': {
                method: "GET",
                isArray: true
            }
        });
	})
