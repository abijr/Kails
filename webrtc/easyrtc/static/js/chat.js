var conversation = new Array();
var selfEasyrtcid = "";

var connect = function() {
	easyrtc.setPeerListener(addMessageToConversation);
	easyrtc.setRoomOccupantListener(convertListToButton);
	easyrtc.connect("kails.chat", loginSucces, loginFailure);
}

var loginSucces = function(easyrtcid) {
	selfEasyrtcid = easyrtcid;
	document.getElementById("user").innerHTML = "User: " + easyrtc.idToName(easyrtcid) + "<br>";
}

var loginFailure = function(errorCode, message) {
	easyrtc.showError(errorCode, message);
}

var sendMessage = function(myEasyrtcid) {
	var message = document.getElementById("sendText").value;

	if(message.replace(/\s/g, "").length === 0) {
		console.log("There is no message to send");
		return;
	} else {
		console.log("Message sent");
	}

	console.log(conversation[selfEasyrtcid]);
	easyrtc.sendDataWS(conversation[selfEasyrtcid], "message", message);
	addMessageToConversation("Me", "message", message);
	document.getElementById("sendText").value = "";
}

var addMessageToConversation = function(who, msgType,message) {
	message = message.replace(/&/g,"&amp;").replace(/</g,"&lt;").replace(/>/g,"&gt;");
	message = message.replace(/\n/g, "<br/>");
	document.getElementById("conversationArea").innerHTML += who + " : " + message + "</br>"; 
}

var clearConnectedUsersList = function() {
	var users = document.getElementById("usersConnected");
	while(users.hasChildNodes()) {
		users.removeChild(users.lastChild);
	}
}

var convertListToButton = function(roomName, data, myEasyrtcid) {
	clearConnectedUsersList();

	var usersConnected = document.getElementById("usersConnected");

	for(var otherPeerEasyrtcid in data) {
		var button = document.createElement("button");
		button.onclick = function(otherPeerEasyrtcid) {
			return function() {
				console.log(selfEasyrtcid);
				conversation[selfEasyrtcid] = otherPeerEasyrtcid;
				clearConnectedUsersList();
			}
		}(otherPeerEasyrtcid)
		
		var label = document.createTextNode(easyrtc.idToName(otherPeerEasyrtcid));
		var space = document.createElement("br");
		var space2 = document.createElement("br");

		button.appendChild(label);
		usersConnected.appendChild(button);
		usersConnected.appendChild(space);
		usersConnected.appendChild(space2);
	}
}