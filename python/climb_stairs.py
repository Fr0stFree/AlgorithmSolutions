from functools import lru_cache


@lru_cache
def solve(n: int) -> int:
    if n <= 0:
        return 0
    if n == 1:
        return 1
    if n == 2:
        return 2
    
    amount = solve(n-1) + solve(n-2)
        
    return amount

print(solve(40))