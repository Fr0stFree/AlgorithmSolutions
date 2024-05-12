declare global {
    interface Array<T> {
        groupBy(fn: (item: T) => string): Record<string, T[]>
    }
}

Array.prototype.groupBy = function<T>(fn: (item: T) => string): Record<string, T[]> {
    const result: Record<string, T[]> = {}
    for (let argument of this) {
        const key = fn(argument);
        if (result[key] === undefined) {
            result[key] = [argument]
        } else {
            result[key].push(argument)
        }
    }
    return result
}

let result;
let expectedResult;

result = [1,2,3].groupBy(String)
expectedResult = {"1":[1],"2":[2],"3":[3]}
console.log(JSON.stringify(result) === JSON.stringify(expectedResult))

result = [{"id":"1"}, {"id":"1"}, {"id":"2"}].groupBy(function (item) {return item.id})
expectedResult = {"1": [{"id": "1"}, {"id": "1"}], "2": [{"id": "2"}]}
console.log(JSON.stringify(result) === JSON.stringify(expectedResult))