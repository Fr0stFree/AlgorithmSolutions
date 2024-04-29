# Welcome. In this kata, you are asked to square every digit of a number and concatenate them.
# For example, if we run 9119 through the function, 811181 will come out, because 92 is 81 and 12 is 1. (81-1-1-81)

def solution [number: int] {
    $number | into string
            | split chars
            | each {|digit| ($digit | into int) ** 2}
            | str join
            | into int
}

def main [] {
    let result = solution 9119
    let expected = 811181
    if $result == $expected { print "PASSED" } else { print $"Expected '($expected)', got '($result)'" }

    let result = solution 765
    let expected = 493625
    if $result == $expected { print "PASSED" } else { print $"Expected '($expected)', got '($result)'" }
}