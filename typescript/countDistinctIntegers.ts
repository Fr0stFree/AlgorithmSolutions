function countDistinctIntegers(nums: number[]): number {
    const uniqueNumbers = new Set<number>();
    nums.forEach((num) => {
        uniqueNumbers.add(num);
        const flippedNumber = Number(String(num).split("").reverse().join(""));
        uniqueNumbers.add(flippedNumber);
        nums.push(flippedNumber);
    });
    return uniqueNumbers.size;
}

let result: number;
let expectedResult: typeof result;

result = countDistinctIntegers([1,13,10,12,31]);
expectedResult = 6;
console.log(result === expectedResult);

result = countDistinctIntegers([2, 2, 2]);
expectedResult = 1;
console.log(result === expectedResult);
