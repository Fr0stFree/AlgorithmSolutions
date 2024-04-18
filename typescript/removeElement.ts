function removeElement(nums: number[], val: number): number {
    let counter = nums.length;
    nums.sort((a, b) => a === val ? 1 : -1).forEach((element, index) => element === val && delete nums[index] && counter--)
    return counter;
}

const result = removeElement([0,1,2,2,3,0,4,2], 2);
const expectedResult = 5;
console.log(result === expectedResult);