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
					openNewConversation(easyrtcid);
					conversation[selfEasyrtcid] = easyrtcid;
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
		var conversationArea = document.getElementById('conversationArea');
		var messageWrapper = document.createElement('div');
		var imageWrapper = document.createElement('div');
		var image = document.createElement('img');
		var messageArea = document.createElement('div');
		var p = document.createElement('p');

		messageWrapper.className = "messageWrapper";
		imageWrapper.className = "friendImage";
		messageArea.className = "friendName";
		image.src = "img/not_available.jpg";

		p.innerHTML = message;

		messageArea.appendChild(p);
		imageWrapper.appendChild(image);
		messageWrapper.appendChild(imageWrapper);
		messageWrapper.appendChild(messageArea);
		conversationArea.appendChild(messageWrapper);

		//document.getElementById("conversationArea").innerHTML += who + " : " + message + "</br>"; 
	}

	var openNewConversation = function(peer) {
		var chatArea = document.createElement('div');
		var status = document.createElement('div');
		var header = document.createElement('button');
		var conversation = document.createElement('div');
		var textarea = document.createElement('textarea');

		chatArea.id = "chatArea";
		chatArea.className = "genericBox";
		status.className = "status";
		header.className = "header";
		conversation.id = "conversationArea";
		textarea.id = "sendText";

		textarea.onkeydown = function(event) {
			if(event.keyCode === 13) {
				Chat.sendMessage();
			}				
		}

		header.innerHTML = peer;

		chatArea.appendChild(status);
		chatArea.appendChild(header);
		chatArea.appendChild(conversation);
		chatArea.appendChild(textarea);
		document.body.appendChild(chatArea);

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
