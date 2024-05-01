# To be a senior, a member must be at least 55 years old and have a handicap greater than 7. In this croquet club,
# handicaps range from -2 to +26; the better the player the lower the handicap.
def solution [player_stats: list<list<int>>] {
    $player_stats | each {|stats| if $stats.0 > 55 and $stats.1 > 7 {"Senior"} else {"Open"}}
}

def main [] {
    let result = solution [[18, 20], [45, 2], [61, 12], [37, 6], [21, 21], [78, 9]]
    let expected = ["Open", "Open", "Senior", "Open", "Open", "Senior"]
    if $result == $expected { print "PASSED" } else { print $"Expected '($expected)', got '($result)'" }
}