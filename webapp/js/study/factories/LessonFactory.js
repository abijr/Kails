angular.module('KailsApp')
	.factory('Lesson', function($resource) {
        return $resource('/study/:id', null, {
            'get': {
                method: "GET",
                isArray: true
            }
        });
    });
