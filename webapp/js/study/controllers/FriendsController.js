angular.module('KailsApp')
	.controller('FriendsController', function($scope, $interval, Status, Websocket, Communication) {
		$scope.LoggedIn = {};

		Websocket.Connect("/ws2");
		Communication.connect().then(function(data) {
			var message = {
				"Data": data,
			};
			// Send data to server to join queue
			Websocket.Send(JSON.stringify(message));
		});

		$interval(function () {
			Status.get({}, function (data) {
				$scope.LoggedIn = data;
			});
		}, 500);
	});
