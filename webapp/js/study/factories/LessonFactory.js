angular.module('KailsApp')
	.factory('Lesson', function($resource) {
        return $resource('/webapp/study/:id', null, {
            'get': {
                method: "GET",
                isArray: true
            }
        });
    });
