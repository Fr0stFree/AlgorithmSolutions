function majorityElement(nums: number[]): number {
    const visitedElements: Map<number, number> = new Map;
    const threshold = nums.length / 2;
    for (let num of nums) {
        const amount = visitedElements.get(num);
        let newAmount;
        amount ? newAmount = amount + 1 : newAmount = 1;
        visitedElements.set(num, newAmount)
        if (newAmount >= threshold) return num;
    }
    return -1;
}

const result = majorityElement([1])
const expected = 1;
console.log(result === expected);