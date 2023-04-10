// Given an integer array nums, return true if any value appears at least twice
// in the array, and return false if every element is distinct.

const containsDuplicate = function(nums) {
    return new Set(nums).size !== nums.length;
};

console.assert(containsDuplicate([1,2,3,1]) === true)
console.assert(containsDuplicate([1,2,3,4]) === false)