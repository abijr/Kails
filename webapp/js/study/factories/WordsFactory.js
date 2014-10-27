angular.module('KailsApp')
	.factory('Words', function($http, $q) {
		var WrongWords = {};
		var GoodWords = {};
		var AllWords = {};
		var ReviewedWords = {};

		return {
			"Get": function() {
				var deferred = $q.defer(); //init promise
				$http.get("webapp/words/all")
					.success(function(data) {
						AllWords = data;
						deferred.resolve(data);
					})
					.error(function () {
						deferred.reject();
					});
				return deferred.promise;
			},
			"Save": function() {
				jsontxt = JSON.stringify(ReviewedWords);
				$http.post("/webapp/words", jsontxt)
					.success(function() {
						ReviewedWords = {};
					});
			},
			"AddGoodWord": function(word) {
				GoodWords[word.Word] = word;
				word.isGood = true;
				ReviewedWords[word.Word] = word;
			},
			"AddWrongWord": function(word) {
				WrongWords[word.Word] = word;
				word.isGood = false;
				ReviewedWords[word.Word] = word;
			},"AllWords": function() {
				return AllWords;
			},
			"WrongWords": function() {
				return WrongWords;
			},
			"GoodWords": function() {
				return GoodWords;
			},
		};
	});
