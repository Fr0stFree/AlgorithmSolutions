def solution [numbers: list<int>] {
    $numbers | uniq
             | enumerate
             | insert amount {|record| $numbers | where $it == $record.item | length}
             | sort-by -r amount
             | get item
}

def main [] {
    let result = solution [4, 1, 5, 5, 0, 0, 3, 1, 6, 2, 2, 1, 1, 2, 2, 3, 3, 4]
    let expected = [2, 1, 3, 0, 5, 4, 6]
    if $result == $expected { print "PASSED" } else { print $"Expected '($expected)', got '($result)'" }
}