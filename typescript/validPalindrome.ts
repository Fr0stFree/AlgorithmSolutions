function isPalindrome(s: string): boolean {
    const letterArray = s.toLowerCase().replace(/[\W|_]/g, "").split("");
    for (let index = 0; index < letterArray.length / 2; index++) {
        if (!(letterArray[index] === letterArray[letterArray.length-index-1])) {
            return false;
        }
    }
    return true;
}

let result: boolean
let expectedResult: boolean

result = isPalindrome("A man, a plan, a canal: Panama");
expectedResult = true;
console.log(result === expectedResult);

result = isPalindrome("race a car");
expectedResult = false;
console.log(result === expectedResult);

result = isPalindrome("ab_a");
expectedResult = true;
console.log(result === expectedResult);