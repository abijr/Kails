angular.module('KailsApp')
	.controller('ChatController', function($scope, $timeout, Communication, Websocket) {


		var selfEasyrtcid;
		var conversation = [];
		var ActiveConversation;

		$scope.section = "";
		$scope.glue = true; // For sticky scrolling
		$scope.Messages = [];
		$scope.Message = {};
		$scope.ChatPartners = [];

		var UpdateChatPartners = function(roomName, data) {
			$timeout(function() {
				var i = 0;
				for (var peer in data) {
					conversation[i] = peer;
					$scope.ChatPartners[i] = {
						id: i,
						name: easyrtc.idToName(peer)
					};
					i++;

				}
			}, 0);
		};

		var addMessageToConversation = function(who, messageType, message) {
			$timeout(function() {
				$scope.Messages.push({
					"isServer": messageType == "server",
					"isUser": who === selfEasyrtcid,
					"Message": message,
				});
			}, 0);
		};

		// Start websocket connection
		Websocket.Connect();
		Websocket.OnMessageFunction(function(packet) {
			console.log(packet.data);
			ActiveConversation = JSON.parse(packet.data).webrtc;
			addMessageToConversation(selfEasyrtcid, "server", "Connected!");
			Websocket.Close();
			$timeout(function() {
				$scope.Section = "Chat";
			}, 0);
		});

		easyrtc.setPeerListener(addMessageToConversation);
		easyrtc.setRoomOccupantListener(UpdateChatPartners);

		Communication.connect().then(function(data) {
			selfEasyrtcid = data;
			var message = {
				"Type": "chat",
				"Data": data,
			};
			Websocket.Send(JSON.stringify(message));
		});


		$scope.setActive = function(id) {
			$scope.ActiveConversation = id;
		};

		$scope.sendMessage = function(myEasyrtcid) {

			if ($scope.Message.text.replace(/\s/g, "").length === 0) {
				console.log("nope! now returning.");
				return;
			}
			console.log(ActiveConversation);
			easyrtc.sendDataWS(ActiveConversation, "message", $scope.Message.text);

			addMessageToConversation(selfEasyrtcid, "message", $scope.Message.text);
			$scope.Message.text = "";
		};

		$scope.MessageClass = function(isUser, isServer) {

		}
	});
