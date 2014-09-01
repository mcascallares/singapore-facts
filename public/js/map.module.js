var MapModule = (function() {

    var map = null;

    var initMap = function(selector) {
        map = $(selector).initMap({
            center : 'Singapore',
            type : 'roadmap',
            options : {
                zoom: 11,
                scrollwheel: false
            }
        });
    };

    var renderMarkers = function() {
        var jqxhr = $.get("/pois/get?layers[]=Recycling%20Bin")
        .done(function(data) {
            var markers = {};
            _.each(data, function(m) {
                markers[m.id] = {
                    position: [m.loc.coordinates[1], m.loc.coordinates[0]]
                };
            });
            map.markers.add(markers);
        })
        .fail(function() {
            alert("error");
        });
    };

    return {
        initMap: initMap,
        renderMarkers: renderMarkers
    }
})();