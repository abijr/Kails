angular.module('KailsApp')
	.controller('FriendsController', function($scope, Friends, Status, Connected) {
		$scope.pageID = "practicePage";
		
		$scope.statusStyle = {color: "gray"};

		var topPosition = 10;
		var leftPosition = 10;
		var positions = {};
		var fiendInfo = [];
		var len = 0;

		myFriend = {
			Username: "",
			StudyLanguage: "",
			top: 0,
			left: 0
		}

		Friends.get({user: 'user'}, function(info) {
			console.log(info);
			getFriendInfo(info);
			checkUsersConnected();
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
			console.log("len: " + len)
			for(var i = 0; i < len; i++) {
				myFriend.Username = friendInfo[i].Username;
				myFriend.StudyLanguage = friendInfo[i].StudyLanguage;
				positions = getPositions();
				myFriend.top = positions.top;
				myFriend.left = positions.left;
				$scope.friends.push(myFriend);
			}
		}

		checkUsersConnected = function() {
			Connected.get(function(Data) {
				if(Data.length > 0) {
					for(var i = 0; i < Data.length; i++) {
						if(isFriend(Data[i])) {
							$scope.statusColor = "#0F0";
						}
					}
				}
			});
		}

		isFriend = function(user) {
			var len = $scope.friends.length;

			for(var i = 0; i < len; i++) {
				if($scope.friends[i].Username === user) {
					return true;
				}
			}

			return false;
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
			var len = $scope.friends.length;
			var numTopics = friend.Topics.length;
			var topic = "sports";

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
					console.log("updating status");
					$scope.statusColor = "#0F0";
				}
			}
		}
	});