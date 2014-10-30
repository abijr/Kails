angular.module('KailsApp')
	.factory('Words', function($http, $q) {
		var WrongWords = {};
		var GoodWords = {};
		var AllWords = {};
		var ReviewedWords = {};
		var got = false;

		function Get() {
			var deferred = $q.defer(); //init promise
			$http.get("webapp/words/all")
				.success(function(data) {
					AllWords = data;
					deferred.resolve(data);
					got = true;
				})
				.error(function() {
					deferred.reject();
				});
			return deferred.promise;
		}

		var setWorstWords = function(n) {
			if (n === undefined) {
				n = 5;
			}
			var arrWords = [];

			if (n > Object.keys(WrongWords).length) {

				$.each(AllWords, function(key, value) {
					var word = value;
					word.Word = key;
					arrWords.push(word);
				});

				arrWords.sort(function(a, b) {
					return a.LastReview < b.LastReview;
				});
			}

			var i = 0;
			for (var j in arrWords) {
				if (i >= n) break;
				WrongWords[arrWords[j].Word] = arrWords[j];
				i++;
			}
		};

		return {
			"Get": Get,
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
			},
			"AllWords": function() {
				return AllWords;
			},
			"WrongWords": function() {
				var deferred = $q.defer(); //init promise
				if (!got) {
					Get().then(function() {
						setWorstWords();
						deferred.resolve(WrongWords);
					});
				} else {
					setWorstWords();
					deferred.resolve(WrongWords);
				}

				return deferred.promise;
			},
			"GoodWords": function() {
				return GoodWords;
			},
		};
	});
