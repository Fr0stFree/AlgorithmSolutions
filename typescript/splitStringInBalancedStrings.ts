function balancedStringSplit(s: string): number {
    let result: number = 0;
    let borderIndex: number = 0
    const counter: Record<string, number> = {L: 0, R:0}

    const addSubstring = (index: number) => {
        counter.R = 0;
        counter.L = 0;
        result++;
        borderIndex = index
    }

    [...s].forEach((letter, index) => {
        counter[letter]++;
        counter.L > 0 && counter.R === counter.L && addSubstring(index);
    });
    return result;
}

let result: number;
let expectedResult: number;

result = balancedStringSplit("RLRRLLRLRL");
expectedResult = 4;
console.log(result === expectedResult);

result = balancedStringSplit("RLRRRLLRLL");
expectedResult = 2;
console.log(result === expectedResult);