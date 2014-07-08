var Videochat = (function() {
	var isAccepted = function(accepted, easyrtcid) {
		if(!accepted) {
			easyrtc.showError("REJECTED MESSAGE", "Your call to " + easyrtc.idToName(easyrtcid) + " has been rejected.");
		}
	}

	var successCB = function(easyrtcid) {
		console.log("Call to " + easyrtc.idToName(easyrtcid) + " succeded.");
		easyrtc.setVideoObjectSrc(document.getElementById("local"), easyrtc.getLocalStream());
	}

	var failureCB = function(errorCode, errMessage) {console.log("Call to " + easyrtc.idToName(remoteEasyRTCID) + " failed." + errMessage)}

	var clearConnectList = function() {
		var otherClientDiv = document.getElementById("otherClients");
		while (otherClientDiv.hasChildNodes()) {
			otherClientDiv.removeChild(otherClientDiv.lastChild);
		}
	}

	var convertListToButtons = function(roomName, data) {
		clearConnectList();
		var otherClientDiv = document.getElementById("otherClients");

		for(var easyrtcid in data) {
			var button = document.createElement("button");
			button.id = "otherPeer";
			button.onclick = function(easyrtcid) {
				return function() {
					Videochat.call(easyrtcid);
				}
			}(easyrtcid);
			 
				var label = document.createTextNode(easyrtc.idToName(easyrtcid));
				button.appendChild(label);
				otherClientDiv.appendChild(button);
		}		
	}

	easyrtc.setStreamAcceptor(function(easyrtcid, stream) {
		easyrtc.setVideoObjectSrc(document.getElementById("local"), easyrtc.getLocalStream());
		easyrtc.setVideoObjectSrc(document.getElementById("remote"), stream);
		console.log("saw video from " + easyrtcid);
	});

	easyrtc.setAcceptChecker(function(easyrtcid, acceptor) {
		var accepted = confirm("Do you accept the call from " +	 easyrtc.idToName(easyrtcid) + " user?");

		if(accepted) {
			acceptor(true);
		} else {
			easyrtc.hangupAll();
			acceptor(false);
		}
	});

	easyrtc.setOnStreamClosed(function(easyrtcid) {
		easyrtc.setVideoObjectSrc(document.getElementById("remote"), "");
	});
	 
	easyrtc.setCallCancelled(function(easyrtcid){
		console.log("Call cancelled");
	});
	
	return {
		start: function() {
			easyrtc.enableVideo(true);
			easyrtc.enableAudio(true);
			easyrtc.setRoomOccupantListener(convertListToButtons);
			Communication.connect();
		}, 

		call: function(remoteEasyRTCID) {
			easyrtc.hangupAll();
			easyrtc.call(remoteEasyRTCID, successCB, failureCB, isAccepted);
		},

		hangUp: function() {
			easyrtc.hangupAll();
		}
	}
})();