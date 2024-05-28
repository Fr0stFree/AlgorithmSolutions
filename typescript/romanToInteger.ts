function romanToInt(s: string): number {
    const romanNumbers: Record<string, number> = {I:1, V:5, X:10, L:50, C:100, D:500, M:1000};
    return [...s].reduce<number>((acc, letter, index, array) => {
        if (index === array.length - 1 || romanNumbers[letter] >= romanNumbers[array[index+1]]) {
            return acc + romanNumbers[letter];
        }
        return acc - romanNumbers[letter];
    }, 0);
}

let result: number;
let expectedResult: typeof result;

result = romanToInt("III");
expectedResult = 3;
console.log(result === expectedResult);

result = romanToInt("LVIII");
expectedResult = 58;
console.log(result === expectedResult);

result = romanToInt("MCMXCIV");
expectedResult = 1994;
console.log(result === expectedResult);