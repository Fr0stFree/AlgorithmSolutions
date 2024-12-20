// Given an integer array nums, move all 0's to the end of it while maintaining
// the relative order of the non-zero elements.
// Note that you must do this in-place without making a copy of the array.


const moveZeroes = function(nums) {
    nums.sort((a, b) => {
        if (a === 0) {
            return 1;
        }
        if (b === 0) {
            return -1;
        }
        return 0;
    });
};

const testArray = [0, 1, 0, 3, 12]

moveZeroes(testArray)

console.assert(JSON.stringify(testArray) === JSON.stringify([1, 3, 12, 0, 0]))