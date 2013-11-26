---
---

var healthApp = angular.module("healthApp", ['ngRoute'])
.config(function($routeProvider, $locationProvider){
    $routeProvider.when("/help", {
        templateUrl: "help.html",
        controller: "IndexCtrl",        
    });
    
    $routeProvider.when("/about", {
        templateUrl: "about.html",
        controller: "IndexCtrl",        
    });
    
    $routeProvider.when("/", {
        templateUrl: "search.html",
        controller: "SearchCtrl",        
    });    
});


healthApp.controller('IndexCtrl', ['$scope', '$http', function($scope, $http){
    
}]);

healthApp.controller('SearchCtrl', ['$scope', '$http', '$location', function($scope, $http, $location){
    $scope.searchLocation = ($location.search()).l;
    $scope.geocodeError = false;
    $scope.noResultsError = false;
    $scope.lat = ($location.search()).lat;
    $scope.lon = ($location.search()).lon;
    $scope.dist = ($location.search()).d || 1609;
    $scope.searchType = ($location.search()).typ;
    
    // $scope.map = new google.maps.Map(document.getElementById("map-canvas"), {zoom: 12, mapTypeId: google.maps.MapTypeId.ROADMAP});
    
    $scope.doGeocode = function(){
        // hit google for lat/lng
        var geocoder = new google.maps.Geocoder();
        
        geocoder.geocode( { 'address': $scope.searchLocation + " chicago, il"}, function(results, status) {
            if (status == google.maps.GeocoderStatus.OK) {
                $scope.lat = results[0].geometry.location.ob;
                $scope.lon = results[0].geometry.location.pb;
                           
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