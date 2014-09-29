var Chat = (function() {
	var conversation = new Array();
	var selfEasyrtcid = Communication.getID();
	var isBlinking = false;
	var blinkingFunc;

	var clearConnectedUsersList = function() {
		var friends = document.getElementById("usersList");
		while (friends.hasChildNodes()) {
			friends.removeChild(friends.lastChild);
		}
	};

	var convertListToButtons = function(roomName, data) {
		clearConnectedUsersList();

		for(var easyrtcid in data) {
			var friends = document.getElementById("usersList");
			var li = document.createElement('li');
			var a = document.createElement('a');
			var divImage = document.createElement('div');
			var divName = document.createElement('div');
			var divStatus = document.createElement('div');
			var image = document.createElement('img');
			var name = document.createElement('p');

			li.id = easyrtcid;
			divImage.className = 'friendImage';
			divName.className = 'friendName';
			divStatus.className = 'friendStatus';
			image.src = 'img/not_available.jpg';

			a.onclick = function(easyrtcid) {
				return function() {
					if(isBlinking) {
						clearTimeout(blinkingFunc);
						li.style.background = "lightseagreen";
						isBlinking = false;
					} else {
						conversation[selfEasyrtcid] = easyrtcid;
					}
				}
			}(easyrtcid);

			name.innerHTML = easyrtc.idToName(easyrtcid);

			divName.appendChild(name);
			divImage.appendChild(image);
			a.appendChild(divImage);
			a.appendChild(divName);
			a.appendChild(divStatus);
			li.appendChild(a);
			friends.appendChild(li);
		}
	}

	var addMessageToConversation = function(who, messageType, message) {
		var conversationArea = document.getElementById('displayText');
		var messageWrapper = document.createElement('div');
		var imageWrapper = document.createElement('div');
		var image = document.createElement('img');
		var messageArea = document.createElement('div');
		var p = document.createElement('p');
		var otherUser;

		messageWrapper.className = "messageWrapper";
		imageWrapper.className = "image";
		messageArea.className = "message";
		image.src = "img/not_available.jpg";

		p.innerHTML = message;

		if(who === selfEasyrtcid) {
			messageWrapper.style.background = "blue";
		} else {
			otherUser = document.getElementById(who);
			messageWrapper.style.background = "red";
			messageWrapper.style.left = "50%";
			blinkingFunc = setInterval(function() {
				blink(otherUser, "lightseagreen", "green");
			}, 500);
			isBlinking = true;
		}

		messageArea.appendChild(p);
		imageWrapper.appendChild(image);
		messageWrapper.appendChild(imageWrapper);
		messageWrapper.appendChild(messageArea);
		conversationArea.appendChild(messageWrapper);
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
			addMessageToConversation(selfEasyrtcid, "message", message);
			document.getElementById("sendText").value = "";
		},

		stopBlink: function() {
			clearTimeout(blinking);
		}

	}
})();
