angular.module('KailsApp')
	.controller('HomeController', function($scope, Users) {
		$scope.Users = [];
		$scope.Search = function () {
			$scope.Users = [];
			Users.get({name: $scope.Query}, function (data) {
				if (data.Error != "") {
					$scope.Users = ["No users found."];
				} else {
					$scope.Users = data.Data;
				}
			});
		};
		$scope.Close = function () {
			$scope.Query = "";
			$scope.Users = [];
		};
    });
