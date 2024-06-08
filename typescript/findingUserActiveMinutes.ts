function findingUsersActiveMinutes(logs: number[][], k: number): number[] {
    const statistics: Record<number, Set<number>> = {};
    logs.forEach((log) => {
        const [userId, minute] = log;
        statistics[userId] ? statistics[userId].add(minute) : statistics[userId] = new Set([minute]);
    })
    const result = new Array(k).fill(0);
    Object.entries(statistics).forEach(([id, minutes]) => result[minutes.size-1]++)
    return result;
}


let result: number[];
let expectedResult: typeof result;

result = findingUsersActiveMinutes([[0,5],[1,2],[0,2],[0,5],[1,3]], 5);
expectedResult = [0,2,0,0,0];
console.log(JSON.stringify(result) === JSON.stringify(expectedResult));

result = findingUsersActiveMinutes([[1,1],[2,2],[2,3]], 4);
expectedResult = [1,1,0,0];
console.log(JSON.stringify(result) === JSON.stringify(expectedResult));