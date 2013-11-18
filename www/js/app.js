var healthApp = angular.module("healthApp", []);

healthApp.controller('SearchCtrl', ['$scope', '$http', function($scope, $http){
    $scope.searchLocation = "";
    // $scope.searchResults = [];
    $scope.lat = "";
    $scope.lon = "";
    
    $scope.map = new google.maps.Map(document.getElementById("map-canvas"), {zoom: 12, mapTypeId: google.maps.MapTypeId.ROADMAP});
    
    $scope.doGeocode = function(){
        // hit google for lat/lng
        var geocoder = new google.maps.Geocoder();
        
        geocoder.geocode( { 'address': $scope.searchLocation}, function(results, status) {
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
            } else {
                console.log("Geocode was not successful for the following reason: " + status);
            }
        });        
    };
    
    $scope.doSearch = function(){
        // hit API for nearby health providers
        $http.jsonp("http://localhost:8080/search",
                    { params: { lat: $scope.lat, lon: $scope.lon, dist: 1500, callback: "JSON_CALLBACK" } }
        )
        .success(function(data){
            console.log(data);
            $scope.searchResults = angular.copy(data);
        })
        .error(function(data, status, headers, config){
            console.log(data);
            console.log(status);
            console.log(headers);
            console.log(config);
        });
    };
}]);