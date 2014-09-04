angular.module('KailsApp')
    .controller('LessonController', function($scope, $http){
        var Data;
        var Counter = 0;

        var NextCard = function () {
            $scope.Card = Data[Counter].Sentence;
            width = (100 * (Counter + 1) / Data.length).toString() + "%";
            $scope.Width = {"width": width}
            Counter += 1;
        }

        $http.get('study/1').success(function(data){
            // console.log(data)
            Data = data;
            NextCard();
        });

        $scope.checkAnswer = function() {
            NextCard();
        }
    }
);
