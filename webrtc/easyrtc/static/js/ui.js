var dialog = function(message, option1, option2, answer) {
	var buttonAccept = document.getElementById("option1");
	var buttonReject = document.getElementById("option2");

	document.getElementById("dialog-box").style.visibility = "visible";
	document.getElementById("message").innerHTML = message;

	buttonAccept.onclick = function(evt) {
		answer(true);
	}

	buttonReject.onclick = function(evt) {
		answer(false);
	}
}