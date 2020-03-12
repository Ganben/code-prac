/**
 * 
 * @param {number} n
 * @return {number} 
 */
export default function integerPartition(n) {
    // create patrition for solving
    const partitionMatrix = Array(n + 1).fill(null).map(() => {
        return Array(n + 1).fill(null);
    });
    // fill patrition matrix with initial value
    // fill first row 
    for (let numberIndex = 1; numberIndex <= n; numberIndex += 1) {
        partitionMatrix[0][numberIndex] = 0;
    }
    // fill 1st col
    for (let summandIndex = 0; summandIndex <= n; summandIndex += 1) {
        partitionMatrix[summandIndex][0] = 1;
    }
    // form m out of summands
    for (let summandIndex = 1; summandIndex <= n; summandIndex += 1) {
        for (let numberIndex = 1; numberIndex <= n; numberIndex += 1) {
            if (summandIndex > numberIndex) {
                // 
                partitionMatrix[summandIndex][numberIndex] = partitionMatrix[summandIndex - 1][numberIndex];
            } else {
                // number of ways to form 5 using 0, 1, 2 e.g.
                const combosWithoutSummand = partitionMatrix[summandIndex - 1][numberIndex];
                const combosWithSummand = partitionMatrix[summandIndex][numberIndex - summandIndex];

                partitionMatrix[summandIndex][numberIndex] = combosWithoutSummand + combosWithSummand;
            }
        }
    }
    return partitionMatrix[n][n];
}