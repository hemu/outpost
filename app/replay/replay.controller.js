'use strict'

angular.module('dominionReplayApp')
.config(['$stateProvider', '$urlRouterProvider', function($stateProvider, $urlRouterProvider) {
  // $urlRouterProvider.otherwise("/replay/1");
  $stateProvider
    .state('replay', {
      url: "/replay/:turn",
      templateUrl: "public/view/replay.html",
      controller: 'ReplayController'
    });
}])
.factory('mySocket', function (socketFactory) {
  return socketFactory({
    url: 'http://localhost:3000/echo'
  });
})
.controller('ReplayController', ['$scope', '$stateParams', 'mySocket', function($scope, $stateParams, mySocket) {
  $scope.message = 'Hello';
  mySocket.setHandler('open', function() {
    console.log("opened event");
    mySocket.send("gonna make a change")
  });
  mySocket.setHandler('message', function(msg) {
    console.log("server sent a message");
    console.log(msg);
    console.log(msg.data);
  });
  console.log("sent socket msg");
  $scope.prevTurn = function(){
    console.log("prev turn button pressed");
  };
  $scope.nextTurn = function(){
    console.log("next turn button pressed");
  };
}]);

// myApp.config(function($stateProvider, $urlRouterProvider) {
//   //
//   // For any unmatched url, redirect to /state1
//   $urlRouterProvider.otherwise("/state1");
//   //
//   // Now set up the states
//   $stateProvider
//     .state('state1', {
//       url: "/state1",
//       templateUrl: "partials/state1.html"
//     })
//     .state('state1.list', {
//       url: "/list",
//       templateUrl: "partials/state1.list.html",
//       controller: function($scope) {
//         $scope.items = ["A", "List", "Of", "Items"];
//       }
//     })
//     .state('state2', {
//       url: "/state2",
//       templateUrl: "partials/state2.html"
//     })
//     .state('state2.list', {
//       url: "/list",
//       templateUrl: "partials/state2.list.html",
//       controller: function($scope) {
//         $scope.things = ["A", "Set", "Of", "Things"];
//       }
//     });
// });