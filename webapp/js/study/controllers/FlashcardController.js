angular.module('KailsApp')
	.controller('FlashcardController', function($scope, Words) {
		console.log(Words.WrongWords());
	});
