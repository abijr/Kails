angular.module('KailsApp')
	.factory('Users', function($resource) {
        return $resource('/webapp/search/:name', null, null);
    });
