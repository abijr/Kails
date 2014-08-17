var Chat = (function() {
	var conversation = new Array();
	var selfEasyrtcid = Communication.getID();

	var clearConnectedUsersList = function() {
		var friends = document.getElementById("friends");
		while (friends.hasChildNodes()) {
			friends.removeChild(friends.lastChild);
		}
	}

	var convertListToButtons = function(roomName, data) {
		clearConnectedUsersList();

		for(var easyrtcid in data) {
			var friends = document.getElementById("friends");
			var button = document.createElement('button');
			var divImage = document.createElement('div');
			var divName = document.createElement('div');
			var divStatus = document.createElement('div');
			var image = document.createElement('img');
			var name = document.createElement('p');

			button.className = 'friend';
			divImage.className = 'friendImage';
			divName.className = 'friendName';
			divStatus.className = 'friendStatus';
			image.src = 'img/not_available.jpg';

			button.onclick = function(easyrtcid) {
				return function() {
					conversation[selfEasyrtcid] = easyrtcid;
					clearConnectedUsersList();
				}
			}(easyrtcid);

			name.innerHTML = easyrtc.idToName(easyrtcid);

			divName.appendChild(name);
			divImage.appendChild(image);
			button.appendChild(divImage);
			button.appendChild(divName);
			button.appendChild(divStatus);			
			friends.appendChild(button);
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
