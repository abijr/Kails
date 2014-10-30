angular.module('KailsApp')
	.controller('FlashcardController', function($scope, Words) {
		$scope.isFlipped = true;
		$scope.Done = false;
		// TODO: Using "AllWords" temporarily
		// need to start using WrongWords
		var words;
		var keys;
		var Counter = 0;

		// Edge case
		// if (words.length === 0) {
		// }

		$scope.NextCard = function(isGood) {

			if (Counter == keys.length) {
				$scope.Done = true;
				return;
			}

			// If card where the 'NextCard'
			$scope.isFlipped = !$scope.isFlipped;


			// Expose sentence to scope
			$scope.Front = keys[Counter];
			$scope.Back = words[$scope.Front].Definition;


			Counter += 1;
		};

		// SenData is the function triggered when
		// all cards have been answered. It sends,
		// the results of the study session.
		var SendData = function() {
		};

		Words.WrongWords().then(function (data) {
			words = data;
			keys = Object.keys(words);
			$scope.NextCard();
		});


	});
