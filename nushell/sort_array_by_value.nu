# Your task is to sort an array of integer numbers by the product of the value and the index of the positions.

def solution [array: list<number>] {
    $array | enumerate
           | insert value {|row| ($row.index + 1) * $row.item}
           | sort-by value
           | get item
    }

def main [] {
    let result = solution [23, 2, 3, 4, 5]
    let expected = [2, 3, 4, 23, 5]
    if $result == $expected { print "PASSED" } else { print $"Expected '($expected)', got '($result)'" }
}