# Capitalise the first letter of the first word.
# Add a period (.) to the end of the sentence.
# Join the words into a complete string, with spaces.
# Do no other manipulation on the words.
def solution [input: list<string>] {$input | str join " " | append "." | str join}

def main [] {
    let result = solution ["FIELDS","of","CORN","are","to","be","sown"]
    let expected = "FIELDS of CORN are to be sown."
    if $result == $expected { print "PASSED" } else { print $"Expected '($expected)', got '($result)'" }
}