/**
 * Binet's formula
 * @param {number} position 
 * @returns {number}
 */
export default function fibonacciClosedForm(position) {
    const topMaxValidPosition = 70;
    //
    if (position < 1|| position > topMaxValidPosition) {
        throw new Error(`invalid positon < 1 or > ${topMaxValidPosition}`);
    }
    //
    const sqrt5 = Math.sqrt(5);
    //
    const phi = ( 1 + sqrt5 ) / 2;

    return Math.floor((phi ** position) / sqrt5 + 0.5);
}