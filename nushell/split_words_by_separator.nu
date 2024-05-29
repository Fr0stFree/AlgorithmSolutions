def solution [array: list<string> separator: string] {
    $array | each {|item| $item | split row $separator} | flatten | where $it != ""
}

def main [] {
    let result = solution ["one.two.three","four.five","six"] "."
    let expected = ["one","two","three","four","five","six"]
    if $result == $expected { print "PASSED" } else { print $"Expected '($expected)', got '($result)'" }

    let result = solution ["$easy$","$problem$"] "$"
    let expected = ["easy","problem"]
    if $result == $expected { print "PASSED" } else { print $"Expected '($expected)', got '($result)'" }
}