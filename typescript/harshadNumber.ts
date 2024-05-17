function sumOfTheDigitsOfHarshadNumber(x: number): number {
    const numberSum = [...String(x)].reduce((acc, strNumber) => acc + Number(strNumber), 0)
    return x % numberSum === 0 ? numberSum : -1;
}

let result: number;
let expectedResult: number;

result = sumOfTheDigitsOfHarshadNumber(18);
expectedResult = 9;
console.log(result === expectedResult);

result = sumOfTheDigitsOfHarshadNumber(23);
expectedResult = -1;
console.log(result === expectedResult)