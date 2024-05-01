# Usually when you buy something, you're asked whether your credit card number, phone number or answer to your most
# secret question is still correct. However, since someone could look over your shoulder, you don't want that shown on
# your screen. Instead, we mask it. Your task is to write a function maskify, which changes all but the last four characters into '#'.

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