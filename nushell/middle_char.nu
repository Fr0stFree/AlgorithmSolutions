def solution [word: string] {
    $word | split chars | get (($word | str length) / 2 | into int) 
}

def main [] {
    let result = solution "wow"
    let expected = "o"
    if $result == $expected { print "PASSED" } else { print $"Expected '($expected)', got '($result)'" }
}