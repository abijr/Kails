angular.module('KailsApp')
	.directive('animatedEllipses', ['$interval', function($interval) {

		function link(scope, element, attrs) {
			var count = 1,
				timeoutId;

			function updateDots() {
                var dots = new Array(count).join('.');
				element.text(dots);
                count++;
                if (count > 4) {
                    count = 1;
                }
			}

			element.on('$destroy', function() {
				$interval.cancel(timeoutId);
			});

			// start the UI update process; save the timeoutId for canceling
			timeoutId = $interval(function() {
				updateDots(); // update DOM
			}, 700);
		}

		return {
			link: link
		};
	}]);
