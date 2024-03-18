# Task 1: Write an expression that outputs the "name" element of the shopping list..
jq '.name' shopping-list.json

# Task 2: Only count the "required" ingredients. Don't include the optional ingredients.
jq '.ingredients | length' shopping-list.json

# Task 3: Write an expression that outputs the amount of sugar. Just the numeric part is wanted.
jq '.ingredients[] | select(.item == "sugar") | .amount.quantity' shopping-list.json

# Task 3: Some of the ingredients can be substituted (if you don't have ingredient X you can use Y).
# Output a JSON object mapping the recommended ingredient to its substitution.
# Include the optional ingredients in the mapping. Check both ingredients and optional ingredients.
jq '.ingredients + .["optional ingredients"] | map(select(.substitute != null)) | map({(.item): .substitute}) | add' shopping-list.json
