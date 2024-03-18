# Task: get most common film rating
jq '[[.[] | {rating, movie}] | group_by(.rating) | .[] | {"rating": .[0].rating, "count": . | length}] | sort_by(.count) | reverse | limit(1;.[]) | .rating' movies.json
