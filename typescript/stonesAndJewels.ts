function numJewelsInStones(jewels: string, stones: string): number {
    const jewelsSet = new Set(jewels);
    return [...stones].filter((element) => jewelsSet.has(element)).length
}

let result: number;
let expectedResult: number;

result = numJewelsInStones("aA", "aAAbbbb");
expectedResult = 3;
console.log(result === expectedResult);

result = numJewelsInStones("z", "ZZ");
expectedResult = 0;
console.log(result === expectedResult);