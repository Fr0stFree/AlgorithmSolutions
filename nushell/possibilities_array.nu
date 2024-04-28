# A non-empty array a of length n is called an array of all possibilities if it contains all numbers
# between 0 and a.length - 1 (both inclusive). Write a function that accepts an integer array and
# returns true if the array is an array of all possibilities, else false.
def solution [array: list<int>] {($array | sort | uniq) == (0..(($array | length) - 1) | each {|it| $it })}

def main [] {
    let result = solution [1,2,0,3]
    let expected = true
    if $result == $expected { print "PASSED" } else { print $"Expected '($expected)', got '($result)'" }

    let result = solution [0,1,2,2,3]
    let expected = false
    if $result == $expected { print "PASSED" } else { print $"Expected '($expected)', got '($result)'" }
}