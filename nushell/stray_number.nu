def solution [numbers: list<int>] {
    $numbers | group-by 
             | transpose key values 
             | insert amount {|x| $x.values | length } 
             | sort-by amount 
             | first 
             | get key 
             | into int
}

def main [] {
    let result = solution [17 17 3 17 ]
    let expected = 3
    if $result == $expected { print "PASSED" } else { print $"Expected '($expected)', got '($result)'" }
}