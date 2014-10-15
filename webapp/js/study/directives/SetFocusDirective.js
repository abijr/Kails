angular.module('KailsApp')
	.directive('getFocus', function() {
		return {
			restrict: 'A',
			controller: function($scope, $element, $attrs, $timeout) {
				var functionName = 'setFocus' + $attrs.getFocus;
				$scope[functionName] = function() {
					$timeout(function() {
						$element.focus();
					});
				};
			},
		};
	});
