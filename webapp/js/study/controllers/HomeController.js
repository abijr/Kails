angular.module('KailsApp')
	.controller('HomeController', function($scope, Users) {
		$scope.result = Users.get({name: "us"});
    });
