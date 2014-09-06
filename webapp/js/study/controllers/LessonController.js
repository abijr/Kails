angular.module('KailsApp')
	.controller('LessonController', function($scope, $http) {
		var Data;
		var CurrentCard;
		var Counter = 0;

		var LessonResults = {
			"pass": false,
			"sentences": []
		};

		$scope.InStart = true;
		$scope.InLesson = false;
		$scope.InValidation = false;
		$scope.InAfterLesson = false;

		var NextCard = function() {

			if (Counter == Data.length) {
				$scope.InLesson = false;
				$scope.InAfterLesson = true;
				return;
			}

			$scope.Card = Data[Counter].Sentence;

			// this is here for testing.
			// won't be needed later.
			$scope.Word = Data[Counter].Word;

			// width in the form of percentage.
			width = (100 * (Counter + 1) / Data.length).toString() + "%";

			// Width is the style binding for the progress bar.
			$scope.Width = {
				"width": width
			};

			Counter += 1;

			// Set button action.
			$scope.Next = ValidateAnswer;
			$scope.InValidation = false;
		};

		var ValidateAnswer = function () {
			var c = $scope.Card;
			if (c.Translation == c.Answer) {
				$scope.Correct = true;
			} else {
				$scope.Correct = false;
			}

			$scope.InValidation = true;

			$scope.Answer = "";
			$scope.Next = NextCard;
		};

		$http.get('study/1').success(function(data) {
			// console.log(data)
			Data = data;
		});


		$scope.startLesson = function() {
			$scope.InLesson = true;
			$scope.InStart = false;

			// Display first card
			NextCard();

		};
	});
