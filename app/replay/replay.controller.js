'use strict'

angular.module('dominionReplayApp')
.factory('mySocket', function (socketFactory) {
  return socketFactory({
    url: 'http://localhost:3000/echo'
  });
})
.config(['$stateProvider', '$urlRouterProvider', function($stateProvider, $urlRouterProvider) {
  $urlRouterProvider.otherwise("/replay/1");
  $stateProvider
    .state('replay', {
      url: "/replay/:turn",
      templateUrl: "public/view/replay.html",
      controller: 'ReplayController'
    });
}])
.controller('ReplayController', ['$scope', '$state', '$stateParams', 'mySocket', 
                                  function($scope, $state, $stateParams, mySocket) {
  $scope.message = 'Hello';
  console.log("!!!!!!!!!!!!!!!");
  console.log(parseInt($stateParams.turn));
  console.log(mySocket.socket.readyState == SockJS.OPEN)
  mySocket.send('testtttt');
  if(mySocket.socket.readyState == SockJS.OPEN){
    // mySocket.send(JSON.stringify({turn: $stateParams.turn}));
    console.log("trying to send new json data, sockjs.open true");
    // mySocket.send(JSON.stringify({
    //   mtype: "turn", 
    //   mdata: {num: 7, pnum: 0}
    // }));
    mySocket.send('testtttt');
  }
  console.log("--- Showing Turn " + $stateParams.turn + " ---");
  mySocket.setHandler('open', function() {
    console.log("opened event");
    mySocket.send(JSON.stringify({
      mtype: "turn", 
      mdata: {num: parseInt($stateParams.turn), pnum: 0}
    }));
    // mySocket.send(JSON.stringify({mtype: "turn"}));
  });
  mySocket.setHandler('message', function(msg) {
    console.log("server sent a message:");
    console.log(msg);
  });
  $scope.prevTurn = function(){
    var curTurn = parseInt($stateParams.turn)
    var prevTurn = curTurn - 1
    if(prevTurn > 0){
      $state.transitionTo('replay', {turn: prevTurn});
    }
  };
  $scope.nextTurn = function(){
    var curTurn = parseInt($stateParams.turn)
    var nextTurn = curTurn + 1
    $state.transitionTo('replay', {turn: nextTurn});
    // $stateProvider.go('replay')
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