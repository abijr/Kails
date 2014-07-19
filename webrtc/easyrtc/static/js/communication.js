var Communication = (function() {
	var id = "";

	var loginSucces = function(easyrtcid) {
		id = easyrtcid;
		document.getElementById("user").innerHTML = "User: " + easyrtc.cleanId(easyrtcid);
	}

	var loginFailure = function(errorCode, message) {
		easyrtc.showError(errorCode, message);
	}

	return {
		connect: function() {
			easyrtc.connect("kails", loginSucces, loginFailure);	
		},

		disconnect: function() {
			easyrtc.disconnect();
		},

		getID: function() {
			return id;
		}
	}
})();
