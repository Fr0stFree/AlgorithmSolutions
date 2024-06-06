def solution [haystack: string needle: string] {
    $haystack | str index-of $needle
}

def main [] {
    let result = solution "sadbutsad" "sad"
    let expected = 0
    if $result == $expected { print "PASSED" } else { print $"Expected '($expected)', got '($result)'" }

    let result = solution "leetcode" "leeto"
    let expected = -1
    if $result == $expected { print "PASSED" } else { print $"Expected '($expected)', got '($result)'" }
}