# In this kata you will create a function that takes a list of non-negative integers and strings and returns a new list
# with the strings filtered out.

def solution [elements: list<any>] {$elements | where ($it | describe) == int}

def main [] {
    let result = solution [1,2,'aasf','1','123',123]
    let expected = [1,2,123]
    if $result == $expected { print "PASSED" } else { print $"Expected '($expected)', got '($result)'" }
}