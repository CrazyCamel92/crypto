/**
 * Created by mechanic on 13/02/17.
 */
var app = angular.module("app",[]);
app.controller("mainCtrl",function ($scope,$http) {
   $scope.result = "";
   $scope.input="";
   $scope.encrypt = function () {
       //var data = {value:$scope.input }
       $http({method:"GET",url:'http://localhost:8080?input='+$scope.input}).then(function (res) {
          $scope.result=res.data;
       });

   }
});