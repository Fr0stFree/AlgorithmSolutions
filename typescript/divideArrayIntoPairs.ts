function divideArray(nums: number[]): boolean {
    nums.sort();
    for (let index = 0; index < nums.length; index += 2) {
        if (nums[index] !== nums[index + 1]) return false;
    }
    return true;
}

let result: boolean;
let expectedResult: typeof result;

result = divideArray([3, 2, 3, 2, 2, 2]);
expectedResult = true;
console.log(result === expectedResult);

result = divideArray([1, 2, 3, 4]);
expectedResult = false;
console.log(result === expectedResult);
