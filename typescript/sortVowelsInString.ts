const VOWELS: Readonly<Set<string>> = new Set(["a", "A", "u", "U", "e", "E", "o", "O", "i", "I"])

function sortVowels(s: string): string {
    const vowels: string[] = [...s].filter((letter) => VOWELS.has(letter)).sort().reverse()
    return [...s].map((letter) => VOWELS.has(letter) ? vowels.pop() : letter).join("");
}


let result: string;
let expectedResult: typeof result;

result = sortVowels("lEetcOde");
expectedResult = "lEOtcede";
console.log(result === expectedResult);

result = sortVowels("lYmpH");
expectedResult = "lYmpH";
console.log(result === expectedResult);
