# Given a string, return if all occurrences of a given letter are always immediately followed by the other given letter.

def solution [letters: list<string>, sentence: string] {
    return (
        $sentence | split words
                  | where {|word| $letters.0 in $word}
                  | each {|word| $word | str index-of ($letters | str join)}
                  | all {|element| $element >= 0}
    )
}

solution ["h" "e"] "he headed to the store"