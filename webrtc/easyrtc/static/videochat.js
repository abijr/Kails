var connect = function() {
	console.log("Initializating....");
	easyrtc.enableAudio(true);
	easyrtc.enableVideo(true);
	easyrtc.setRoomOccupantListener(convertListToButtons);
	easyrtc.easyApp("kails", "local", ["remote"], loginSuccess, loginFailure);
}

function clearConnectList() {
	var otherClientDiv = document.getElementById("otherClients");
	while (otherClientDiv.hasChildNodes()) {
		otherClientDiv.removeChild(otherClientDiv.lastChild);
	}
}
 
 
function convertListToButtons (roomName, data) {
	clearConnectList();
	var otherClientDiv = document.getElementById("otherClients");
	for(var easyrtcid in data) {
		var button = document.createElement("button");
		button.id = "otherPeer";
		button.onclick = function(easyrtcid) {
			return function() {
				call(easyrtcid);
			};
		}(easyrtcid);
		 
		var label = document.createTextNode(easyrtc.idToName(easyrtcid));
		button.appendChild(label);
		otherClientDiv.appendChild(button);
	}
}
 
 
var call = function(remoteEasyRTCID) {
	easyrtc.hangupAll();
 
	var successCB = function() {};
	var failureCB = function() {};
	easyrtc.call(remoteEasyRTCID, successCB, failureCB);
}

var hangUp = function() {
	easyrtc.hangupAll();
}

function loginSuccess(easyrtcid) {
	selfEasyrtcid = easyrtcid;
	document.getElementById("user").innerHTML = "User: " + easyrtc.cleanId(easyrtcid);
}
 
 
function loginFailure(errorCode, message) {
	easyrtc.showError(errorCode, message);
}

