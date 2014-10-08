angular.module('KailsApp')
	.controller("InfoController", function($scope, $routeParams, User) {
		$scope.pageID = "practicePage";

		User.get({name: 'user'}, function (info) {			
			$scope.topics = info.Topics;	
			console.log($scope.topics);					
		});
	});