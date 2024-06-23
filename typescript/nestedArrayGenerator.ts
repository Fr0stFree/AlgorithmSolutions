type MultidimensionalArray = (MultidimensionalArray | number)[]

function* inorderTraversal(arr: MultidimensionalArray): Generator<number, void, unknown> {
    while (arr.length) {
        const element = arr.shift()
        if (typeof element == 'object') {
            arr.unshift(...element);
            continue;
        }
        if (typeof element == 'number') yield element;
    }
}


const generator = inorderTraversal([[[6]],[1,3],[]]);

console.assert(generator.next().value == 6)
console.assert(generator.next().value == 1)
console.assert(generator.next().value == 3)
