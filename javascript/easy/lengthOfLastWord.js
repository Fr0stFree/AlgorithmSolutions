// iven a string s consisting of words and spaces, return the length of the last word in the string.
// A word is a maximal substring consisting of non-space characters only.


const lengthOfLastWord = s => s.replace(/\s+/g, " ").trim().split(" ").pop().length;

console.assert(lengthOfLastWord("     Hello      World  ") === 5);