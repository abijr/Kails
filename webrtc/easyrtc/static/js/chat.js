var Chat = (function() {
	var conversation = new Array();
	var selfEasyrtcid = Communication.getID();

	var clearConnectedUsersList = function() {
		var otherClientDiv = document.getElementById("usersConnected");
		while (otherClientDiv.hasChildNodes()) {
			otherClientDiv.removeChild(otherClientDiv.lastChild);
		}
	}

	var convertListToButtons = function(roomName, data) {
		clearConnectedUsersList();
		var otherClientDiv = document.getElementById("usersConnected");

		for(var easyrtcid in data) {
			var button = document.createElement("button");
			button.id = "otherPeer";
			button.onclick = function(easyrtcid) {
				return function() {
					conversation[selfEasyrtcid] = easyrtcid;
					clearConnectedUsersList();
				}
			}(easyrtcid);
			 
			var label = document.createTextNode(easyrtc.idToName(easyrtcid));
			button.appendChild(label);
			otherClientDiv.appendChild(button);
		}		
	}

	var addMessageToConversation = function(who, messageType, message) {
		message = message.replace(/&/g,"&amp;").replace(/</g,"&lt;").replace(/>/g,"&gt;");
		message = message.replace(/\n/g, "<br/>");
		document.getElementById("conversationArea").innerHTML += who + " : " + message + "</br>"; 
	}

	return {
		start: function() {
			easyrtc.setPeerListener(addMessageToConversation);
			easyrtc.setRoomOccupantListener(convertListToButtons);
			Communication.connect();
		},

		sendMessage: function(myEasyrtcid) {
			var message = document.getElementById("sendText").value;

			if(message.replace(/\s/g, "").length === 0) {
				console.log("There is no message to send");
				return;
			} else {
				console.log("Message sent");
			}

			easyrtc.sendDataWS(conversation[selfEasyrtcid], "message", message);
			addMessageToConversation("Me", "message", message);
			document.getElementById("sendText").value = "";
		}
	}
})();