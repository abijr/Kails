angular.module('KailsApp')
	.factory('Websocket', function($timeout) {

		var connection;

		var ws = {};
		ws.OnMessageFunction = function(func) {
			connection.onmessage = func;
		};

		ws.Connect = function() {
			if (window.WebSocket) {
				connection = new WebSocket("ws://localhost:3000/ws");
			} else {
                console.log("Unable to connect to websocket");
            }
		};

        ws.Send = function (message) {
			console.log("Sending message: " + message);
            $timeout(function () {
                if (connection.readyState === 1) {
                    connection.send(message);
                } else {
                    ws.Send(message);
                }
            }, 5);
        };

		ws.Close = function () {
			connection.close();
		};

        return ws;

	});
