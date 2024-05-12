function wordPattern(pattern: string, s: string): boolean {
    const words = s.split(" ");
    const tokens = pattern.split("");

    if (words.length !== pattern.length) return false;
    if (new Set(words).size !== new Set(tokens).size) return false;

    const tokenToWord = new Map<string, string>();
    const wordToToken = new Map<string, string>();

    for (let [index , token] of tokens.entries()) {
        const mappedWord = tokenToWord.get(token);
        const actualWord = words[index];
        if (!mappedWord) {
            if (wordToToken.has(actualWord)) return false;

            tokenToWord.set(token, actualWord);
            wordToToken.set(actualWord, token);
            continue
        }

        if (mappedWord !== actualWord) return false;
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