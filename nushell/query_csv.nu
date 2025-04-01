def main [] {
    let first_three = open "./data/employees.csv" | first 3
    let sorted_by_salary = open "./data/employees.csv" | sort-by salary | reverse
    let avg_salary_by_department = open "./data/employees.csv" 
        | group-by department 
        | transpose
        | each { 
            |item| { 
                dep: $item.column0,
                salaries: ($item.column1.salary | math avg)
            } 
        }
    let max_salary_empl = open "./data/employees.csv" | sort-by salary | reverse | first
    $max_salary_empl 
}
