import Heap from './Heap';

export default class MaxHeap extends Heap {
    /**
     * impl. this 
     * @param {*} first 
     * @param {*} second
     * @returns {boolean} 
     */
    pairIsInCorrectOrder(first, second) {
        return this.compare.greaterThanOrEqual(first, second);
    }
}