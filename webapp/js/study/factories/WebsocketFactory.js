angular.module('KailsApp')
	.factory('Websocket', function($timeout, $http, $q) {

		var connection;



		var ws = {};
		ws.OnMessageFunction = function(func) {
			connection.onmessage = func;
		};

		ws.Connect = function(ws) {
			if (ws === undefined)  ws = "/ws";

			if (window.WebSocket) {
				console.log("Host is: " + Host);
				connection = new WebSocket("ws://" + Host + ws);
			} else {
				console.log("Unable to connect to websocket");
			}
		};

		ws.Send = function(message) {
			$timeout(function() {
				if (connection.readyState === 1) {
					console.log("Sending message: " + message);
					connection.send(message);
				} else {
					ws.Send(message);
				}
			}, 5);
		};

		ws.Close = function() {
			connection.close();
		};

		return ws;

	});
