angular.module('KailsApp')
    .controller('ProgramController', function($scope, $timeout) {
        $timeout(function () {
            $scope.Width = {width: PercentDone};
        }, 50);
    });
