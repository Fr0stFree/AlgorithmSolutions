def solution [number: int] {$number | into string | split chars | $in == ($in | reverse)}

def main [] {
    let result = solution 121
    let expected = true
    if $result == $expected { print "PASSED" } else { print $"Expected '($expected)', got '($result)'" }

    let result = solution 10
    let expected = false
    if $result == $expected { print "PASSED" } else { print $"Expected '($expected)', got '($result)'" }
}