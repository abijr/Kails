angular.module('KailsApp')
	.controller('FriendsController', function($scope, Friends) {
		$scope.pageID = "practicePage";
		$scope.friends = []

		var topPosition = 10;
		var leftPosition = 10;
		var positions = {};
		var fiendInfo = [];
		var len = 0;


		Friends.get({user: 'user'}, function(info) {
			console.log(info);
			getFriendInfo(info);
		});

		/*$scope.$on("$viewContentLoaded", function() {
			checkStatus();
		});*/

		getPositions = function() {
			var topValue = topPosition;
			var leftValue = leftValue;

			if(topPosition >= 70) {
				topPosition += 15;
				leftPosition = 10;
			} else {
				leftPosition += 30;
			}

			return {
				top: topValue,
				left: leftValue
			}
		}

		getFriendInfo = function(friendInfo) {
			len = friendInfo.length;

			for(var i = 0; i < len; i++) {
				$scope.friends[i].Username = friendInfo[i].Username;
				$scope.friends[i].StudyLanguage = friendInfo[i].StudyLanguage;
				positions = getPositions();
				$scope.friends[i].top = positions.top;
				$scope.friends[i].left = positions.left;
			}
		}

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