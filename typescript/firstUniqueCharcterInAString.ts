function firstUniqChar(s: string): number {
    const map = new Map<string, {index: number; isUnique: boolean}>();
    let result = -1;

    s.split("").forEach((letter, index) => {
        const item = map.get(letter);
        if (!item) {
            map.set(letter, {index, isUnique: true});
        } else {
            item.isUnique = false
        }
    })

    for (let [key, {index, isUnique}] of map.entries()) {
        if (isUnique) {
            result = index;
            break
        }
    }
    return result
}

let result: number
let expectedResult: number

result = firstUniqChar("leetcode");
expectedResult = 0;
console.log(result === expectedResult);

result = firstUniqChar("loveleetcode");
expectedResult = 2;
console.log(result === expectedResult);

result = firstUniqChar("aabb");
expectedResult = -1;
console.log(result === expectedResult);
