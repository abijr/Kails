var dialog = function(message, option1, option2, answer) {
	var buttonAccept = document.getElementById("option1");
	var buttonReject = document.getElementById("option2");

	document.getElementById("dialog-box").style.visibility = "visible";
	buttonAccept.innerHTML = option1;
	buttonReject.innerHTML = option2;

	buttonAccept.onclick = function(evt) {
		answer(true);
	}

	buttonReject.onclick = function(evt) {
		answer(false);
	}
}

var confirmation = function(answer) {
	document.getElementById("confirm").style.visibility = "visible";

	document.getElementById("accept").onclick = function(evt) {
		answer(true);
	}

	document.getElementById("reject").onclick = function(evt) {
		answer(false);
	}
}