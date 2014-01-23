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

var healthApp = angular.module("healthApp", ['ngRoute', 'analytics', 'ngCookies'])
.config(function($routeProvider, $locationProvider){
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

healthApp.controller('SearchCtrl', ['$scope', '$http', '$location', 'analytics', '$cookies', function($scope, $http, $location, analytics, $cookies){
    $scope.searchLocation = ($location.search()).l || "";
    $scope.searched = ($location.search()).s || false;
    $scope.geocodeError = false;
    $scope.noResultsError = false;
    $scope.lat = ($location.search()).lat;
    $scope.lon = ($location.search()).lon;
    $scope.dist = ($location.search()).d || 1609;
    $scope.searchType = ($location.search()).typ || 0;
    $scope.showSMS = ($cookies.showSMS != "false");
    
    // $scope.map = new google.maps.Map(document.getElementById("map-canvas"), {zoom: 12, mapTypeId: google.maps.MapTypeId.ROADMAP});
    
    $scope.doGeocode = function(){
        // hit google for lat/lng
        var geocoder = new google.maps.Geocoder();
        var preferredBounds = new google.maps.LatLngBounds(new google.maps.LatLng(41.771041, -87.794856), new google.maps.LatLng( 42.030819, -87.577265));

        geocoder.geocode( { 'address': $scope.searchLocation, 'bounds': preferredBounds }, function(results, status) {
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
    
    $scope.hideSMS = function(){
        $scope.showSMS = false;
        $cookies.showSMS = 'false';
    };
    
}]);