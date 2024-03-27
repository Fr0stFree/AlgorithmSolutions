// Given two strings needle and haystack, return the index of the first occurrence of needle in haystack,
// or -1 if needle is not part of haystack.

var strStr = function(haystack, needle) {
    for (let index = 0; index < haystack.length; index++) {
        if (haystack[index] === needle[0] && haystack.substring(index, index + needle.length) === needle) {
            return index
        }
    }
    return -1
};

const res = strStr("a", "a");
console.log(res);