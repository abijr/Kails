var enableAudio = true;
var enableVideo = true;

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

	var failureCB = function(errorCode, errMessage) {
		console.log("Call to " + easyrtc.idToName(remoteEasyRTCID) + " failed." + errMessage);
	}

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
					clearConnectList();
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
			easyrtc.enableVideo(enableVideo);
			easyrtc.enableAudio(enableAudio);
			easyrtc.setRoomOccupantListener(convertListToButtons);
			Communication.connect();
		}, 

		call: function(remoteEasyRTCID) {
			easyrtc.hangupAll();
			easyrtc.call(remoteEasyRTCID, successCB, failureCB, isAccepted);
		},

		hangUp: function() {
			easyrtc.hangupAll();
		},
	}
})();

var Controls = (function() {
	return {
		setMicro: function() {
			var img = document.getElementById("imageMicro");

			if(enableAudio) {
				enableAudio = false;
				easyrtc.enableMicrophone(enableAudio);
				img.src = "img/microCancel.png"
			} else {
				enableAudio = true;
				easyrtc.enableMicrophone(enableAudio);
				img.src = "img/micro.png"
			}
		},

		setCamera: function() {
			var img = document.getElementById("imageVideo");

			if(enableVideo) {	
				enableVideo = false;
				easyrtc.enableCamera(enableVideo);
				img.src = "img/cameraCancel.png"
			} else {
				enableVideo = true;
				easyrtc.enableCamera(enableVideo);
				img.src = "img/camera.png"
			}
		},

		setLocalWindow: function(position) {
			switch(position) {
				case 'left_top':
					$("#local").css({'left':'1%', 'top':'6%'});
					break;
				case 'right_top':
					$("#local").css({'left':'74%', 'top':'6%'});
					break;
				case 'left_bottom':
					$("#local").css({'left':'1%', 'top':'72%'});
					break;
				case 'right_bottom':
					$("#local").css({'left':'74%', 'top':'72%'});
					break;
			}
		}
	}
})();