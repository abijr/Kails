angular.module('KailsApp')
	.factory('WrongWords', function() {
		var Words = [];

		return {
			"Save": function() {
				jsontxt = JSON.stringify(Words);
				$http.post("/words", jsontxt);
			},
			"AddWord": function(word) {
				Words.push(word);
			},
			"Get": function() {
				return Words;
			}
		};
	});
