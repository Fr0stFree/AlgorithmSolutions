function minimumAverage(nums: number[]): number {
    let currentMin: number = Infinity;

    nums.sort((a, b) => a - b)
    for (let index = 0; index < nums.length / 2; index++) {
        const min = (nums[index] + nums[nums.length-1-index]) / 2
        currentMin = Math.min(currentMin, min);
    }
    return currentMin
}

let result: number;
let expectedResult: typeof result;

result = minimumAverage([7,8,3,4,15,13,4,1]);
expectedResult = 5.5;
console.log(result === expectedResult);

result = minimumAverage([1,9,8,3,10,5]);
expectedResult = 5.5;
console.log(result === expectedResult);
