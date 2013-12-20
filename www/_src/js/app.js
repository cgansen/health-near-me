---
---

// h/t http://stackoverflow.com/a/12262820/1247272
angular.module('analytics', ['ng']).service('analytics', [
      '$rootScope', '$window', '$location', function($rootScope, $window, $location) {
        var track = function() {
            ga('send', 'pageview', {'page': $location.path()});
        };
      $rootScope.$on('$routeChangeSuccess', track);
    }
]);

var healthApp = angular.module("healthApp", ['ngRoute', 'analytics'])
.config(function($routeProvider, $locationProvider){
    $routeProvider.when("/help", {
        templateUrl: "help.html",
        controller: "IndexCtrl",        
    });
    
    $routeProvider.when("/about", {
        templateUrl: "about.html",
        controller: "IndexCtrl",        
    });

    $routeProvider.when("/sms", {
        templateUrl: "sms.html",
        controller: "IndexCtrl",        
    });    
    
    $routeProvider.when("/", {
        templateUrl: "search.html",
        controller: "SearchCtrl",        
    });    
});

healthApp.filter('escape', function() {
  return window.encodeURIComponent;
});

healthApp.controller('IndexCtrl', ['$scope', '$http', 'analytics', function($scope, $http, analytics){
    
}]);

healthApp.controller('SearchCtrl', ['$scope', '$http', '$location', 'analytics', function($scope, $http, $location, analytics){
    $scope.searchLocation = ($location.search()).l || "";
    $scope.geocodeError = false;
    $scope.noResultsError = false;
    $scope.lat = ($location.search()).lat;
    $scope.lon = ($location.search()).lon;
    $scope.dist = ($location.search()).d || 1609;
    $scope.searchType = ($location.search()).typ || 0;
    
    // $scope.map = new google.maps.Map(document.getElementById("map-canvas"), {zoom: 12, mapTypeId: google.maps.MapTypeId.ROADMAP});
    
    $scope.doGeocode = function(){
        // hit google for lat/lng
        var geocoder = new google.maps.Geocoder();
        
        geocoder.geocode( { 'address': $scope.searchLocation + " chicago, il"}, function(results, status) {
            if (status == google.maps.GeocoderStatus.OK) {
                $scope.lat = results[0].geometry.location.lat();
                $scope.lon = results[0].geometry.location.lng();
                           
                $scope.doSearch();
                
                // $scope.map.setCenter(results[0].geometry.location);
                //                     $scope.map.setZoom(14)
                //                     var marker = new google.maps.Marker({
                //                         map: $scope.map,
                //                         position: results[0].geometry.location
                //                     });
                $scope.geocodeError = false
            } else {
                $scope.geocodeError = true
                console.log("Geocode was not successful for the following reason: " + status);
            }
            
            $scope.updatePath();
        });        
    };
    
    if($location.search().s == "1"){ $scope.doGeocode() };
    
    $scope.updatePath = function(){
        $location.path("/").search({
            l: $scope.searchLocation, 
            typ: $scope.searchType, 
            lat: $scope.lat, 
            lon: $scope.lon,
            d: $scope.dist,
            s: 1
        });
    }
    
    $scope.doSearch = function(){        
        
        // hit API for nearby health providers
        $http.jsonp("{{ site.api_url }}/search",
                    { params: { 
                        lat: $scope.lat, 
                        lon: $scope.lon, 
                        dist: $scope.dist, 
                        searchType: $scope.searchType,
                        callback: "JSON_CALLBACK" } }
        )
        .success(function(data){
            $scope.noResultsError = false;
            $scope.searchResults = angular.copy(data);
        })
        .error(function(data, status, headers, config){
            $scope.noResultsError = true;
            $scope.searchResults = []
        });
    };
}]);