function rotateString(s: string, goal: string): boolean {
    if (s.length !== goal.length) return false;
    if (s.length === 0) return true;

    const letters = s.split("");

    const indexes: number[] = [...s.matchAll(new RegExp(goal[0], 'gi'))].map(letter => letter.index)
    for (let index of indexes) {
        const shiftedS = letters.slice(index).concat(letters.slice(0, index)).join("")
        if (shiftedS === goal) return true;
    }
    return false
}

let result: boolean
let expectedResult: boolean

result = rotateString("abcde", "cdeab");
expectedResult = true;
console.log(result === expectedResult);

result = rotateString("abcde", "abced");
expectedResult = false;
console.log(result === expectedResult);

result = rotateString("bbbacddceeb", "ceebbbbacdd");
expectedResult = true;
console.log(result === expectedResult);