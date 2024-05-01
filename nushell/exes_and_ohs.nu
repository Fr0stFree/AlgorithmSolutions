# Check to see if a string has the same amount of 'x's and 'o's.
# The method must return a boolean and be case insensitive. The string can contain any char.
def solution [word: string] {
    let splitted = $word | str downcase
                         | split chars
                         | where $it in ["x" "o"]
                         | group-by
    ($splitted.o | length) == ($splitted.x | length)
}

def main [] {
    let result = solution "ooxXm"
    let expected = true
    if $result == $expected { print "PASSED" } else { print $"Expected '($expected)', got '($result)'" }

    let result = solution "xooxx"
    let expected = false
    if $result == $expected { print "PASSED" } else { print $"Expected '($expected)', got '($result)'" }
}