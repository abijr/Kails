/*
* Permite interceptar el evento de "Enter".
* Cuando el elemento recibe la tecla "<enter>"
* se ejecuta el atributo de la directiva. Ejemplo:

<input type="text" on-enter="miFuncion()"/>

* Ejecuta "miFuncion()" cuando se entra la tecla
* "<enter>"
*/
angular.module('KailsApp')
    .directive('onEnter', function() {
        return function(scope, element, attrs) {
            element.bind("keydown keypress", function(event) {
                if(event.which === 13) {
                    scope.$apply(function(){
                        scope.$eval(attrs.onEnter, {'event': event});
                    });

                    event.preventDefault();
                }
            });
        };
    });
