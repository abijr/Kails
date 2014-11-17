angular.module('KailsApp')
	.controller('FriendsController', function($scope, $interval, $location, Status, Websocket, Communication, Pending) {
		$scope.LoggedIn = {};
		$scope.request = {};
		$scope.acceptFunction = function() {};

		function acceptVideo() {
			console.log("Accepting video");
			Pending.pendingRequest = true;
			Pending.webrtc = $scope.request.webrtc;
			Pending.user = $scope.request.user;
			$location.url("/videochat");
		}

		function acceptChat() {
			console.log("Accepting chat");
			Pending.pendingRequest = true;
			Pending.webrtc = $scope.request.webrtc;
			Pending.user = $scope.request.user;
			$location.url("/chat");
		}

		Websocket.Connect("/ws2");
		Websocket.OnMessageFunction(function(data) {
			$scope.request = JSON.parse(data.data);
			if ($scope.request.type == "videochat") {
				$scope.acceptFunction = acceptVideo;
			} else if ($scope.request.type == "chat") {
				$scope.acceptFunction = acceptChat;
			}

			$('#acceptModal').foundation('reveal', 'open');
			console.log($scope.request);
			$scope.$apply();
		});

		Communication.connect().then(function(data) {
			var message = {
				"Data": {
					"id": data
				},
			};
			// Send data to server to join queue
			Websocket.Send(JSON.stringify(message));
		});

		var stop = $interval(function() {
			Status.get({}, function(data) {
				$scope.LoggedIn = data;
			});
		}, 1000);

		$scope.$on('$destroy', function() {
			// Make sure that the interval is destroyed
			$('#acceptModal').foundation('reveal', 'close');
			$interval.cancel(stop);
			Websocket.Close();
		});

		$scope.requestVideo = function(user) {
			var message = {
				"Type": "request",
				"Data": {
					"request": "videochat",
					"peer": user,
				},
			};

			Pending.pendingAccept = true;
			Pending.user = $scope.request.user;
			// Send data to server to join queue
			Websocket.Send(JSON.stringify(message));
			$location.url("/videochat");
		};

		$scope.requestChat = function(user) {
			var message = {
				"Type": "request",
				"Data": {
					"request": "chat",
					"peer": user,
				},
			};

			Pending.pendingAccept = true;
			Pending.user = $scope.request.user;
			// Send data to server to join queue
			Websocket.Send(JSON.stringify(message));
			$location.url("/chat");
		};

	});
