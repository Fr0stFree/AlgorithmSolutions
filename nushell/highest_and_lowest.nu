# In this little assignment you are given a string of space separated numbers, and have to return the highest and lowest number.

def solution [digits: string] {
    let array = $digits | split row " "
                        | each {|digit| $digit | into int}
                        | sort

    [($array | last) ($array | first)] | str join " "
}

def main [] {
    let result = solution "1 2 3 4 5"
    let expected = "5 1"
    if $result == $expected { print "PASSED" } else { print $"Expected '($expected)', got '($result)'" }

    let result = solution "1 9 3 4 -5"
    let expected = "9 -5"
    if $result == $expected { print "PASSED" } else { print $"Expected '($expected)', got '($result)'" }
}
