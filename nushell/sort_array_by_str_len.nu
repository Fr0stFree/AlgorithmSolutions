def solution [words: list<string>] {
    $words | sort-by -c {|a, b| ($a | str length) < ($b | str length)}
}

def main [] {
    let result = solution ["Telescopes" "Glasses" "Eyes" "Monocles"]
    let expected = ["Eyes" "Glasses" "Monocles" "Telescopes"]
    if $result == $expected { print "PASSED" } else { print $"Expected '($expected)', got '($result)'" }
}