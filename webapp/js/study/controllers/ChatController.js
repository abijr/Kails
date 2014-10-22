angular.module('KailsApp')
	.controller('ChatController', function($scope, $timeout, Communication) {

    var selfEasyrtcid;
	var conversation    = [];
    $scope.glue = true; // For sticky scrolling
    $scope.Messages     = [];
    $scope.ChatPartners = [];

	var UpdateChatPartners = function(roomName, data) {
        $timeout(function() {
            var i = 0;
            for (var peer in data) {
                conversation[i] = peer;
                $scope.ChatPartners[i] = {
                    id: i,
                    name: easyrtc.idToName(peer)
                };
                i++;

            }
        }, 0);
	};

	var addMessageToConversation = function(who, messageType, message) {

        $timeout(function () {
            $scope.Messages.push({
                "isPeer": who === selfEasyrtcid,
                "Message": message,
            });
        }, 0);
	};


	easyrtc.setPeerListener(addMessageToConversation);
	easyrtc.setRoomOccupantListener(UpdateChatPartners);
	Communication.connect();

	selfEasyrtcid = Communication.getID();

    $scope.setActive = function (id) {
        $scope.ActiveConversation = id;
    };

	$scope.sendMessage = function(myEasyrtcid) {

		if($scope.Message.replace(/\s/g, "").length === 0) {
			return;
		}
		easyrtc.sendDataWS(
            conversation[$scope.ActiveConversation], "message", $scope.Message);

		addMessageToConversation(selfEasyrtcid, "message", $scope.Message);
        $scope.Message = "";
	};
});
