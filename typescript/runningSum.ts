function runningSum(nums: number[]): number[] {
    return nums.reduce<number[]>((acc, item) => {
        acc.push(item + (acc[acc.length - 1] || 0))
        return acc
    }, [])
}

let result: number[];
let expectedResult: number[];

result = runningSum([1,2,3,4]);
expectedResult = [1,3,6,10];
console.log(JSON.stringify(result) === JSON.stringify(expectedResult));

result = runningSum([1,1,1,1,1]);
expectedResult = [1,2,3,4,5];
console.log(JSON.stringify(result) === JSON.stringify(expectedResult));