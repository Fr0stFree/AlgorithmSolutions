# sum up all numbers between 1...n with step equals 1

def solution [limit: number] {1..$limit | reduce {|item, accumulator| $item + $accumulator}}

def main [] {
    let result = solution 5
    let expected = 15
    if $result == $expected { print "PASSED" } else { print $"Expected '($expected)', got '($result)'" }
}