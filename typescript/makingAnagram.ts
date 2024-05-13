function minSteps(s: string, t: string): number {
    const dictionary = new Map<string, number>();
    const insertStringToDict = (word: string, amountModifier: (arg: number) => number, initialAmount: number) => {
        [...word].forEach((letter) => {
            const amount = dictionary.get(letter);
            amount !== undefined ? dictionary.set(letter, amountModifier(amount)) : dictionary.set(letter, initialAmount);
        });
    };
    insertStringToDict(s, (amount) => amount + 1, 1);
    insertStringToDict(t, (amount) => amount - 1, -1);
    return Array.from(dictionary.entries()).reduce((acc, [_, number]) => Math.abs(number) + acc, 0) / 2
}

let result: number;
let expectedResult: number;

result = minSteps("bab", "aba");
expectedResult = 1;
console.log(result === expectedResult);

result = minSteps("leetcode", "practice");
expectedResult = 5;
console.log(result === expectedResult);

result = minSteps("anagram", "mangaar");
expectedResult = 0;
console.log(result === expectedResult);