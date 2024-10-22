"""
Дано выражение в обратной польской записи. Определите его значение
"""

def solve(data: str) -> int:
    stack: list[int] = []
    for obj in data.split(" "):
        if obj.isdigit():
            stack.append(int(obj))
        else:
            first, second = stack.pop(), stack.pop()
            statement = f"{second}{obj}{first}"
            result = eval(statement)
            stack.append(result)
    return stack[0]


if __name__ == "__main__":
    data = input()
    solution = solve(data)
    print(solution)