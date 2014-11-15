angular.module('KailsApp')
	.controller('FriendsController', function($scope, $interval, $location, Status, Websocket, Communication, Pending) {
		$scope.LoggedIn = {};
		$scope.request = {};

		Websocket.Connect("/ws2");
		Websocket.OnMessageFunction(function(data) {
			$scope.request = JSON.parse(data.data);
			$scope.$apply();
			$('#myModal').foundation('reveal', 'open');
			console.log($scope.request);
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
			$interval.cancel(stop);
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

		$scope.acceptVideo = function () {
			Pending.pendingRequest = true;
			Pending.webrtc = $scope.request.webrtc;
			Pending.user = $scope.request.user;
			$('#myModal').foundation('reveal', 'close');
			$location.url("/videochat");
		};
	});
