angular.module('KailsApp')
	.controller('HomeController', function($scope, Users) {
		$scope.Users = [];

		// Sends user search query to server
		// and updates ui with results.
		$scope.Search = function () {
			$scope.Users = [];
			Users.get({name: $scope.Query}, function (data) {
				if (data.Error !== "") {
					$scope.Users = ["No users found."];
				} else {
					$scope.Users = data.Data;
				}
			});
		};

		// Reset variables and remove
		// answers' view from ui.
		$scope.Close = function () {
			$scope.Query = "";
			$scope.Users = [];
		};
    });
