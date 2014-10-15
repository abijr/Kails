angular.module('KailsApp')
	.controller('LessonController', function($scope, $routeParams, Lesson) {
		var Data;
		var CurrentCard;
		var CorrectCount = 0;
		var Counter = 0;

		// Initialize lesson results variable
		var LessonResults = {
			"pass": false,
			"sentences": []
		};

		// Initailize visibility variables
		// (what's supposed to be seen)
		$scope.InStart = true;
		$scope.InLesson = false;
		$scope.InValidation = false;
		$scope.InAfterLesson = false;


		// Get data from server
		Data = Lesson.get({id: $routeParams.LessonId});
		// Old way doing it...
		// $http.get('study/1').success(function(data) {
		// 	// console.log(data)
		// 	Data = data;
		// });

		var NextCard = function() {

			// If card where the 'NextCard'
			// was issued is the last on, trigger
			// 'SendData' function
			if (Counter == Data.length) {
				$scope.InLesson = false;
				$scope.InAfterLesson = true;
				SendData();
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
				CorrectCount++;
			} else {
				$scope.Correct = false;
			}

			$scope.InValidation = true;

			$scope.Answer = "";
			$scope.Next = NextCard;
			$scope.setFocus();

		};


		// SenData is the function triggered when
		// all cards have been answered. It sends,
		// the results of the study session.
		var SendData = function () {
			var pass;
			if (CorrectCount/Data.length >= 0.7) {
				pass = true;
			} else {
				pass = false;
			}

			LessonResults.pass = pass;

			jsontxt = JSON.stringify(LessonResults);

			// Debugging stuff:
			// console.log("sending data: ");
			// console.log(data);

			var results = Lesson.save({id:1}, jsontxt,
				// Success function
				function (){
					console.log(results);
					$scope.ExperienceGained = results.ExperienceGained;
				}
			);
			// Old way of doing it.
			// $http.post('study/1', JSON.stringify(data)).success(function(data) {
			// 	// console.log("success!");
			// });
		};

		$scope.startLesson = function() {
			$scope.InLesson = true;
			$scope.InStart = false;

			// Display first card
			NextCard();

		};
	});
