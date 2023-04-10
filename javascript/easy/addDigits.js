// Given an integer num, repeatedly add all its digits until the result has only one digit, and return it.

const addDigits = function(num) {
    if (num < 10) return num;
    return addDigits(num.toString()
                        .split('')
                        .reduce((accumulator, current) => accumulator + Number(current), 0));
};

console.assert(addDigits(38) === 2)