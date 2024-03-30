// Given two arrays of strings a1 and a2 return a sorted array r in lexicographical order of the strings of a1 which
// are substrings of strings of a2

function inArray(array1, array2){
    const result = [];

    for (word of array2) {
        array1.forEach((subword, index) => word.includes(subword) && array1.splice(index, 1) && result.push(subword));
        if (!array1.length) {
            break;
        }
    }
    return result.sort();
}

const result = inArray(["arp", "live", "strong"], ["lively", "alive", "harp", "sharp", "armstrong", "aa", "bb", "cc"])
console.assert(JSON.stringify(result) === JSON.stringify(["arp", "live", "strong"]))