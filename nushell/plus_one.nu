def solution [digits: list<int>] {
    $digits | str join | into int | $in + 1 | into string | split chars | each {|digit| $digit | into int}
}

def main [] {
    let result = solution [1,2,3]
    let expected = [1,2,4]
    if $result == $expected { print "PASSED" } else { print $"Expected '($expected)', got '($result)'" }

    let result = solution [4,3,2,1]
    let expected = [4,3,2,2]
    if $result == $expected { print "PASSED" } else { print $"Expected '($expected)', got '($result)'" }
}