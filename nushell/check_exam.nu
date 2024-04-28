# The two arrays are not empty and are the same length. Return the score for this array of answers, giving +4 for each
# correct answer, -1 for each incorrect answer, and +0 for each blank answer, represented as an empty string
def solution [answers, correct_answers] {
    $answers | enumerate
             | each {|it| if $it.item == ($correct_answers | get $it.index) {4} else {-1} }
             | math sum
}

def main [] {
    let result = solution ["a", "a", "b", "b"] ["a", "c", "b", "d"]
    let expected = 6
    if $result == $expected { print "PASSED" } else { print $"Expected '($expected)', got '($result)'" }

}