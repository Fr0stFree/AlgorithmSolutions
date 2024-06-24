def solution [divisible: int, divider: int] {
    mut current: int = $divisible
    mut arr = []
    while $current != 0 {
        $arr = ($arr | append $current)
        if $current mod $divider == 0 {$current = $current // $divider} else {$current = $current - 1}
    }
    $arr | str join " -> "
}

def main [] {
    let result = solution 29 3
    let expected = '29 -> 28 -> 27 -> 9 -> 3 -> 1'
    if $result == $expected { print "PASSED" } else { print $"Expected '($expected)', got '($result)'" }
}