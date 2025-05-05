def is_palindrome(number: int) -> bool:
    digits = []

    while number > 9:
        part = number % 10
        number = number // 10
        digits.append(part)
    digits.append(number)

    return digits == digits[::-1]


if __name__ == "__main__":
    print(is_palindrome(123))
    print(is_palindrome(121))
