angular.module('KailsApp')
	.controller('LessonController', function($scope, $http) {
		var Data;
		var Counter = 0;

		$scope.InStart = true;
		$scope.InLesson = false;
		$scope.InAfterLesson = false;

		var NextCard = function() {

			if (Counter == Data.length) {
				$scope.InLesson = false;
				$scope.InAfterLesson = true;
			}

			$scope.Card = Data[Counter].Sentence;

			// width in the form of percentage.
			width = (100 * (Counter + 1) / Data.length).toString() + "%";

			// Width is the style binding for the progress bar.
			$scope.Width = {
				"width": width
			}

			Counter += 1;
		}

		$http.get('study/1').success(function(data) {
			// console.log(data)
			Data = data;
		});

		$scope.checkAnswer = function() {
			NextCard();
		}

		$scope.startLesson = function() {
			$scope.InLesson = true;
			$scope.InStart = false;
			NextCard();
		}
	});