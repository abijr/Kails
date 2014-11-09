angular.module('KailsApp')
	.controller('VideochatController', function($scope, $timeout, Communication, Websocket) {
		var enableAudio = true;
		var enableVideo = true;
		var userlist;
		var accepted;

		$scope.Section = "";

		var videochat = function() {
			var isAccepted = function(accepted, easyrtcid) {
				if (!accepted) {
					easyrtc.showError("REJECTED MESSAGE", "Your call to " + easyrtc.idToName(easyrtcid) + " has been rejected.");
				} else {
					console.log("Accepted");
				}
			};

			var successCB = function(easyrtcid) {
				console.log("Call to " + easyrtc.idToName(easyrtcid) + " succeded.");
				easyrtc.setVideoObjectSrc(document.getElementById("local"), easyrtc.getLocalStream());
			};

			var failureCB = function(errorCode, errMessage) {
				console.log("Call to " + easyrtc.idToName(remoteEasyRTCID) + " failed." + errMessage);
			};

			var convertListToButtons = function(roomName, data) {
				console.log("user list updated");
				userlist = data;
			};

			easyrtc.setStreamAcceptor(function(easyrtcid, stream) {
				$timeout(function() {
					$scope.Section = "Chat";
					easyrtc.setVideoObjectSrc(document.getElementById("local"), easyrtc.getLocalStream());
					easyrtc.setVideoObjectSrc(document.getElementById("remote"), stream);
					Websocket.Close();
					console.log("Connecting with video to " + easyrtcid);
				}, 0);
			});

			easyrtc.setAcceptChecker(function(easyrtcid, acceptor) {
				var message;
				acceptor(true);
			});

			easyrtc.setOnStreamClosed(function(easyrtcid) {
				easyrtc.setVideoObjectSrc(document.getElementById("remote"), "");
			});

			easyrtc.setCallCancelled(function(easyrtcid) {
				console.log("Call cancelled");
			});

			return {
				start: function() {
					easyrtc.enableVideo(enableVideo);
					easyrtc.enableAudio(enableAudio);
					easyrtc.setRoomOccupantListener(convertListToButtons);
					Communication.connect().then(function(data) {
						selfEasyrtcid = data;
						var message = {
							"Type": "videochat",
							"Data": data,
						};

						// Send data to server to join queue
						Websocket.Send(JSON.stringify(message));
					});
				},

				call: function(remoteEasyRTCID) {
					easyrtc.hangupAll();
					easyrtc.call(remoteEasyRTCID, successCB, failureCB, isAccepted);
				},

				hangUp: function() {
					easyrtc.hangupAll();
				}
			};
		};

		$scope.Videochat = videochat();

		// Start websocket connection
		Websocket.Connect();
		Websocket.OnMessageFunction(function(packet) {
			console.log(packet.data);
			var ActiveConversation = JSON.parse(packet.data).webrtc;
			$scope.Videochat.call(ActiveConversation);
		});

		$scope.Videochat.start();


	});
