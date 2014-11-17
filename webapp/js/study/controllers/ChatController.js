angular.module('KailsApp')
	.controller('ChatController', function($scope, $timeout, Communication, Websocket) {


		var selfEasyrtcid;
		var conversation = [];
		var ActiveConversation = "";

		$scope.section = "";
		$scope.glue = true; // For sticky scrolling
		$scope.textinput = {};
		$scope.Messages = [];
		$scope.Message = {};
		$scope.ChatPartners = [];
		$scope.Data = {};

		var UpdateChatPartners = function(roomName, data) {
			if (ActiveConversation === "") return;
			for (var id in data) {
				if (id == ActiveConversation) return;
			}

			$timeout(function function_name(argument) {
				Communication.disconnect();
				ActiveConversation = "";
				$scope.textinput.disabled = true;
				addMessageToConversation(selfEasyrtcid, "server", "Connection with peer has failed :/");
			}, 1500);
		};

		function addMessageToConversation(who, messageType, message) {
			$timeout(function() {
				$scope.Messages.push({
					"isServer": messageType == "server",
					"isUser": who === selfEasyrtcid,
					"Message": message,
				});
			}, 0);
		}

		// Start websocket connection
		Websocket.Connect();
		Websocket.OnMessageFunction(function(packet) {
			console.log(packet.data);
			$scope.Data = JSON.parse(packet.data);
			ActiveConversation = $scope.Data.webrtc;
			addMessageToConversation(selfEasyrtcid, "server", "Connected!");
			Websocket.Close();
			$scope.Section = "Chat";
			$scope.$apply();
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

		$scope.UserInfo = function() {
			return "/userinfo/" + $scope.Data.name;
		};
		
		$scope.$on('$destroy', function() {
			// Make sure that the interval is destroyed
			$('#acceptModal').foundation('reveal', 'close');
			$interval.cancel(stop);
			Websocket.Close();
		});
	});
