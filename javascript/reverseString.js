// Write a function that reverses a string. The input string is given as an array of characters s.
// You must do this by modifying the input array in-place with O(1) extra memory.


const reverseString = function(s) {
    let index = 0;
    s.sort((a, b) => {
        index++;
        return -index;
    })
    return s;
};

console.assert(JSON.stringify(reverseString(["h","e","l","l","o"])) === JSON.stringify(["o","l","l","e","h"]))