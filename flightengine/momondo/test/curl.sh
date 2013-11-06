curl --verbose \
    "http://www.momondo.com/Momondo.asmx/WhereToGoGetUpdated?origCode=BRU&destCode=LIM&year=2014&month=03"

curl -H "Content-Type: application/json; charset=UTF-8" \
    -H "Accept: application/json" \
    --data '{"origCode":"BRU", "destCode":"LIM", "year": 2014, "month":3}' \
    http://www.momondo.com/Momondo.asmx/WhereToGoGetUpdated

curl -H "Content-Type: application/json; charset=UTF-8" \
    -H "Accept: application/json" \
    --data @WhereToGoGetUpdated.json \
    http://www.momondo.com/Momondo.asmx/WhereToGoGetUpdated

curl --verbose \
    -H "Content-Type: application/json; charset=UTF-8" \
    -H "Accept: */*" \
    --data @StartFlightSearch.json \
    http://www.momondo.com/FlightWS.asmx/StartFlightSearch
