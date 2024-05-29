function splitWordsBySeparator(words: string[], separator: string): string[] {
    return words.reduce<string[]>((acc, word) => (
        acc.concat(word.split(separator).filter((value) => value))
    ), [])
}

let result: string[];
let expectedResult: typeof result;

result = splitWordsBySeparator(["one.two.three","four.five","six"], ".");
expectedResult = ["one","two","three","four","five","six"];
console.log(JSON.stringify(result) === JSON.stringify(expectedResult));

result = splitWordsBySeparator(["$easy$","$problem$"], "$");
expectedResult = ["easy","problem"];
console.log(JSON.stringify(result) === JSON.stringify(expectedResult));
