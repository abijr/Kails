angular.module('KailsApp')
	.controller('FriendsController', function($scope, Friends, Status) {
		$scope.pageID = "practicePage";
		$scope.friends = [];
		$scope.statusStyle = {color: "gray"};

		var topPosition = 10;
		var leftPosition = 10;
		var positions = {};
		var fiendInfo = [];
		var len = 0;


		Friends.get({user: 'user'}, function(info) {
			console.log(info);
			getFriendInfo(info);
			checkStatus();
		});

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
			var len = friendInfo.length;

			for(var i = 0; i < len; i++) {
				$scope.friends[i].Username = friendInfo[i].Username;
				$scope.friends[i].StudyLanguage = friendInfo[i].StudyLanguage;
				positions = getPositions();
				$scope.friends[i].top = positions.top;
				$scope.friends[i].left = positions.left;
			}
		}

		checkStatus = function() {
			Status.get({topic: 'user/sports'}, function(friend) {
				updateStatus(friend);
				checkStatus();
			});
		}

		updateStatus = function(friend) {
			var isFriend = false;
			var shareTopic = false;
			var len = friendInfo.length;
			var numTopics = friend.Topics.length;

			if(friend.Username != undefined) {
				for(var i = 0; i < len; i++) {
					if($scope.friends[i].Username === friend.Username) {
						isFriend = true;
					}
				}

				for(var j = 0; j < numTopics; j++) {
					if(topic === friend.Topics[j]) {
						shareTopic = true;
					}
				}

				if(isFriend && shareTopic) {
					//update status
					console.log("updating status")
				}
			}
		}
	});