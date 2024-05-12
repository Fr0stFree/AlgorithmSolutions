function wordPattern(pattern: string, s: string): boolean {
    const tokenToWord = new Map<string, string>();
    const wordToToken = new Map<string, string>();
    const words = s.split(" ")

    if (words.length !== pattern.length) {
        return false
    }

    for (let [index , token] of pattern.split("").entries()) {
        const mappedWord = tokenToWord.get(token);
        const actualWord = words[index];
        if (mappedWord) {
            if (mappedWord !== actualWord) {
                return false
            }
        } else {
            if (wordToToken.get(actualWord)) {
                return false
            }
            tokenToWord.set(token, actualWord);
            wordToToken.set(actualWord, token);
        }
    }
    return true;
}

let result: boolean
let expectedResult: boolean

result = wordPattern("abba", "dog cat cat dog");
expectedResult = true;
console.log(result === expectedResult);

result = wordPattern("abba", "dog cat cat fish");
expectedResult = false;
console.log(result === expectedResult);

result = wordPattern("abba", "dog dog dog dog");
expectedResult = false;
console.log(result === expectedResult);

result = wordPattern("aaa", "aa aa aa aa");
expectedResult = false;
console.log(result === expectedResult);