def main [] {
    open './data/app.log' | parse '{datetime} [{level}] {message}' 
                          | insert request_type {|row| $row.message | parse -r '(?P<method>\w+) \/(?P<path>[^\s\)]+)' | first }
                          | flatten
                          | where level == 'ERROR'
}