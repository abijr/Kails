level_1 = {
	"Level": 1,
	"Topics": 1,
	"Friends": 5,
	"Features": ["chat"],
	"Time": 10 //in minutes
};

angular.module('KailsApp')
	.controller("PracticeController", function($scope, $routeParams, User) {	
		var add = true;	
		var topics = ["sports", "entertainment", "vehicles", "various"];

		$scope.topicsAvailable = [];
		$scope.pageID = "practicePage";

		User.get({name: 'user'}, function (info) {
			$scope.topics = info.Topics;	
			console.log($scope.topics);				
		});

		$scope.showTopics = function() {
			var numTopicsAllowed = level_1.Topics;
			var currentNumTopics = $scope.topics.length;
			
			if(numTopicsAllowed > currentNumTopics) {
				$scope.add = true;

				if(currentNumTopics > 0) {
					for(var i = 0; i < topics.length; i++) {
						for(var j = 0; j < currentNumTopics; j++) {
							if(topics[i] == $scope.topics[j]) {
								add = false;
							}
						}

						if(add) {
							$scope.topicsAvailable.push(topics[i]);
						}
						add = true;
					}
				}
				else {
					$scope.topicsAvailable = topics;
				}
			}
			else {
				$scope.notAllowed = true;
			}
		}

		$scope.addTopic = function(topic) {
			var jsontxt;

			$scope.topics.push(topic);
			jsontxt = JSON.stringify($scope.topics);
			User.save({name:'other'}, jsontxt);
			$scope.add = false;
		}

		$scope.deleteTopic = function(topic) {
			var jsontxt;
			var index;

			index = $scope.topics.indexOf(topic);
			$scope.topics = $scope.topics.splice(index, 1);
			jsontxt = JSON.stringify($scope.topics);
			User.save({name:'other'}, jsontxt);
		}
	});