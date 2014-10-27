angular.module('KailsApp')
	.controller('WordsController', function($scope, Words) {
		Words.Get()
		.then(function (words) {
			$scope.Words = words;
			console.log(Words.AllWords());
		});
	});
