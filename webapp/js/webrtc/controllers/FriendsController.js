angular.module('KailsApp')
	.controller('FriendsController', function($scope, Friends) {
		$scope.pageID = "practicePage";
		$scope.friends = ["other", "another"]

		/*Friends.get({user: 'user'}, function(friends) {
			//console.log($scope.friends);
		});*/

		$scope.$on("$viewContentLoaded", function() {
			checkStatus();
		});

		checkStatus = function() {
			console.log("enter function");
			Friends.get({topic: 'sports'}, function(info) {
				updateStatus(info.Username);
				checkStatus();
			});
		}

		updateStatus = function(name) {
			console.log(name);
		}
	});