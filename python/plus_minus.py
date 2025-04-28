def plusMinus(arr: list[int]) -> None:
    # Write your code here
    positive, negative = 0, 0
    for number in arr:
        if number > 0:
            positive += 1
        elif number < 0:
            negative += 1

    zeros = len(arr) - positive - negative
    
    for value in (positive, negative, zeros):
        ratio = round(value/len(arr), 5)
        print(f"{ratio:.6f}")


if __name__ == '__main__':
    n = int(input().strip())

    arr = list(map(int, input().rstrip().split()))

    plusMinus(arr)
