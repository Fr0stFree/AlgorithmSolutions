// Given an integer n, return an array ans of length n + 1 such that for each i (0 <= i <= n),
// ans[i] is the number of 1's in the binary representation of i.

const countBits = n => {
    const result = [];
    for (let number = 0; number <= n; number++) {
        result.push(number.toString(2)
                          .split('')
                          .filter(value => value === '1')
                          .length);
    }
    return result
};

console.assert(JSON.stringify(countBits(5)) === JSON.stringify([0, 1, 1, 2, 1, 2]))