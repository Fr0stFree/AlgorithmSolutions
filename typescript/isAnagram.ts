function isAnagram(s: string, t: string): boolean {
    if (s.length !== t.length) return false;

    const map = new Map<string, number>()

    for (let letter of s.split("")) {
        const amount = map.get(letter);
        amount !== undefined ? map.set(letter, amount + 1) : map.set(letter, 1)
    }
    for (let letter of t.split("")) {
        const amount = map.get(letter);
        amount !== undefined ? map.set(letter, amount - 1) : map.set(letter, -1)
    }
    return Array.from(map.entries()).every(([_, amount]) => amount === 0);
}

let result: boolean
let expectedResult: boolean

result = isAnagram("anagram", "nagaram");
expectedResult = true;
console.log(result === expectedResult);

result = isAnagram("rat", "car");
expectedResult = false;
console.log(result === expectedResult);