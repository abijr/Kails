angular.module('KailsApp')
	.controller('FriendsController', function($scope, Friends, Status, Connected) {
		$scope.pageID = "practicePage";
		$scope.friends = [];

		var topPosition = 10;
		var leftPosition = 10;
		var positions = {};
		var fiendInfo = [];
		var len = 0;

		myFriend = {
			Username: "",
			StudyLanguage: "",
			top: '',
			left:'',
			color: ''
		}

		getPositions = function() {
			var topValue = topPosition;
			var leftValue = leftPosition;

			if(topPosition >= 70) {
				topPosition += 15;
				leftPosition = 10;
			} else {
				leftPosition += 30;
			}

			return {
				top: "'" + topValue.toString() + "%'",
				left: "'" + leftValue.toString() + "%'"
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
				myFriend.color = "'#808080'";
				$scope.friends.push(myFriend);
			}
		}

		checkUsersConnected = function() {
			Connected.get(function(friends) {
				console.log(friends);
				if(friends.length > 0) {
					for(var i = 0; i < friends.length; i++) {
						if(isFriend(friends[i].User) && friends[i].isLogged) {
							var id = "#" + friends[i].User;
							$(id).css({color: '#0F0'});
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
			console.log("checking...");
			Status.get({topic: 'user/sports'}, function(friend) {
				console.log(friend.User);
				console.log(friend.Topics);
				friend.isLogged = true;
				console.log(friend.isLogged);
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

			if(friend.User != undefined) {
				for(var i = 0; i < len; i++) {
					if($scope.friends[i].Username === friend.User) {
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
					var id = "#" + friend.User;
					console.log(id);
					console.log("updating status");
					if(friend.isLogged) {
						//$(id).css({'color': '#0F0'});
						$scope.friends[0].color = "'#0F0'";
					} else {
						//$(id).css({'color': '#808080'});
						$scope.friends[0].color = "'#808080'";
					}
				
				}
			}
		}

		Friends.get({user: 'user'}, function(info) {
			console.log(info);
			getFriendInfo(info);
			checkUsersConnected();
			checkStatus();
		});
	});
