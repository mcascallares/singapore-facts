echo "Processing recycling-bins.kml - Generating json for importing into MongoDB"
togeojson recycling-bins.kml > recycling-bins.geojson

echo "var pois = " > recycling-bins.js
jq '[{layer:"Recycling Bin", loc: .features | .[] | .geometry}]' recycling-bins.geojson >> recycling-bins.js
echo ";" >> recycling-bins.js

echo "Importing pois into MongoDB"
mongo recycling-bins.js insert-pois.js

rm recycling-bins.geojson recycling-bins.js