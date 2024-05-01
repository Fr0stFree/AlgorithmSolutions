# Check to see if a string has the same amount of 'x's and 'o's.
# The method must return a boolean and be case insensitive. The string can contain any char.
def solution [card_number: string] {
    $card_number | split chars
                 | reverse
                 | enumerate
                 | each {|it| if $it.index < 4 {$it.item} else {"#"}}
                 | reverse
                 | str join
}

def main [] {
    let result = solution "4556364607935616"
    let expected = "############5616"
    if $result == $expected { print "PASSED" } else { print $"Expected '($expected)', got '($result)'" }

    let result = solution "Nananananananananananananananana Batman!"
    let expected = "####################################man!"
    if $result == $expected { print "PASSED" } else { print $"Expected '($expected)', got '($result)'" }

    let result = solution "1"
    let expected = "1"
    if $result == $expected { print "PASSED" } else { print $"Expected '($expected)', got '($result)'" }
}