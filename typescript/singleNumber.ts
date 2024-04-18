function singleNumber(nums: number[]): number {
    const set = new Set<number>
    for (let num of nums) {
        set.has(num) ? set.delete(num) : set.add(num);
    }
    return set.values().next().value
}

const result = singleNumber([4,1,2,1,2]);
const expected = 4;
console.log(expected === result);