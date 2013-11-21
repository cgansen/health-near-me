var healthApp = angular.module("healthApp", []);

healthApp.controller('SearchCtrl', ['$scope', '$http', function($scope, $http){
    $scope.searchLocation = "";
    $scope.geocodeError = false;
    $scope.noResultsError = false;
    $scope.lat = "";
    $scope.lon = "";
    $scope.searchType = "all";
    
    $scope.map = new google.maps.Map(document.getElementById("map-canvas"), {zoom: 12, mapTypeId: google.maps.MapTypeId.ROADMAP});
    
    $scope.doGeocode = function(){
        // hit google for lat/lng
        var geocoder = new google.maps.Geocoder();
        
        geocoder.geocode( { 'address': $scope.searchLocation + " chicago, il"}, function(results, status) {
            if (status == google.maps.GeocoderStatus.OK) {
                $scope.lat = results[0].geometry.location.ob;
                $scope.lon = results[0].geometry.location.pb;
                
                $scope.doSearch();
                
                $scope.map.setCenter(results[0].geometry.location);
                $scope.map.setZoom(14)
                var marker = new google.maps.Marker({
                    map: $scope.map,
                    position: results[0].geometry.location
                });
                $scope.geocodeError = false
            } else {
                $scope.geocodeError = true
                console.log("Geocode was not successful for the following reason: " + status);
            }
        });        
    };
    
    $scope.doSearch = function(){
        // hit API for nearby health providers
        $http.jsonp("http://localhost:8080/search",
                    { params: { 
                        lat: $scope.lat, 
                        lon: $scope.lon, 
                        dist: 1500, 
                        searchType: $scope.searchType,
                        callback: "JSON_CALLBACK" } }
        )
        .success(function(data){
            $scope.noResultsError = false;
            console.log(data);
            $scope.searchResults = angular.copy(data);
        })
        .error(function(data, status, headers, config){
            $scope.noResultsError = true;
            $scope.searchResults = []
            console.log(data);
            console.log(status);
            console.log(headers);
            console.log(config);
        });
    };
}]);