angular.module('KailsApp')
	.controller('LessonController', function($scope, $routeParams, Lesson) {
		var Data;
		var CurrentCard;
		var CurrentWord;
		var CorrectCount = 0;
		var Counter = 0;
		var LessonId = $routeParams.LessonId;

		// Initialize lesson results variable
		var LessonResults = {
			"Pass": false,
			"WrongWords": []
		};

		// Initailize visibility variables
		// (what's supposed to be seen)
		$scope.InStart = true;
		$scope.InLesson = false;
		$scope.InValidation = false;
		$scope.InAfterLesson = false;


		// Get data from server
		Data = Lesson.get({
			id: LessonId
		});
		/* Data is of the form:
		Data = [{
				"Sentence": {
					"Native": "sentence",
					"Translation": "translation"
				},
				"Word": "Word",
				"Definition": "of word"
			},{
				"Sentence": {
					"Native": "sentence",
					"Translation": "translation"
				},
				"Word": "Word",
				"Definition": "of word"
			}]
		*/

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

			// Expose sentence to scope
			$scope.Card = Data[Counter].Sentence;

			// Save current word
			CurrentWord = Data[Counter].Word;

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
			$scope.setFocusTextInput();
		};

		// ValidateAnswer is the function
		// that validates currently entered
		// answer.
		var ValidateAnswer = function() {
			var c = $scope.Card;
			if (c.Translation == c.Answer) {
				$scope.Correct = true;
				CorrectCount++;
			} else {
				$scope.Correct = false;
				LessonResults.WrongWords.push(CurrentWord);
			}

			$scope.InValidation = true;

			// Clean textbox
			$scope.Answer = "";
			// Set button action.
			$scope.Next = NextCard;
			$scope.setFocusButton();

		};


		// SenData is the function triggered when
		// all cards have been answered. It sends,
		// the results of the study session.
		var SendData = function() {
			// If more than 70% of answers correct,
			// user has "passed" the lesson.
			var pass;
			if (CorrectCount / Data.length >= 0.7) {
				pass = true;
			} else {
				pass = false;
			}

			LessonResults.Pass = pass;

			jsontxt = JSON.stringify(LessonResults);

			var results = Lesson.save({ id: LessonId }, jsontxt,
				function() { // Success function
					$scope.ExperienceGained = results.ExperienceGained;
				}
			);
		};

		$scope.startLesson = function() {
			$scope.InLesson = true;
			$scope.InStart = false;

			// Display first card
			NextCard();

		};
	});
