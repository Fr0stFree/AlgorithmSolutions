
def solution [word: string] {
    $word | split chars
          | enumerate
          | each {|item| 1..($item.index + 1) | reduce -f "" {|it, acc| $acc + ($item.item | str downcase )}
                                              | str capitalize}
          | str join "-"
}

def main [] {
    let result = solution "cwAt"
    let expected = "C-Ww-Aaa-Tttt"
    if $result == $expected { print "PASSED" } else { print $"Expected '($expected)', got '($result)'" }
}