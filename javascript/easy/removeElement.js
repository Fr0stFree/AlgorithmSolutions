// Given an integer array nums and an integer val, remove all occurrences of val in nums in-place.
// The order of the elements may be changed. Then return the number of elements in nums which are not equal to val.
//
// Consider the number of elements in nums which are not equal to val be k, to get accepted,
// you need to do the following things:
//
//Change the array nums such that the first k elements of nums contain the elements which are not equal to val.
//The remaining elements of nums are not important as well as the size of nums.
//Return k.

const removeElement = function(nums, val) {
    for (let index = 0; index < nums.length; index++) {
        if (nums[index] === val) {
            nums.splice(index, 1);
            index--;
        }
    }
};

const test_array = [1, 2, 3, 5, 4, 2, 1, 2]

removeElement(test_array, 2)

console.assert(JSON.stringify(test_array) === JSON.stringify([1, 3, 5, 4, 1]))
