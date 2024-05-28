function reverse(x: number): number {
    const threshold = 2147483648;
    const reversedStrNumber = String(x).split("")
                                              .filter((value) => value !== "-")
                                              .reverse()
                                              .join("")
                                              .replace(/^0*/, "");

    const isFit = (inputNumber: string) => {
        const reversedInputNumberArray = [...inputNumber];
        const thresholdArray = [...String(threshold)];
        for (let index = 0; index < inputNumber.length; index++) {
            const inputDigit = Number(reversedInputNumberArray[index]);
            const thresholdDigit = Number(thresholdArray[index]);
            if (thresholdDigit === inputDigit) continue;
            return inputDigit < thresholdDigit;
        }
    }

    if (String(threshold).length < reversedStrNumber.length) return 0;
    if (String(threshold).length > reversedStrNumber.length) return x < 0 ? -Number(reversedStrNumber) : Number(reversedStrNumber);
    if (isFit(reversedStrNumber)) return x < 0 ? -Number(reversedStrNumber) : Number(reversedStrNumber);
    return 0;
}

let result: number;
let expectedResult: typeof result;

result = reverse(123);
expectedResult = 321;
console.log(result === expectedResult);

result = reverse(-123);
expectedResult = -321;
console.log(result === expectedResult);

result = reverse(120);
expectedResult = 21;
console.log(result === expectedResult);

result = reverse(1534236469);
expectedResult = 0;
console.log(result === expectedResult);

result = reverse(-2147483412);
expectedResult = -2143847412;
console.log(result === expectedResult);