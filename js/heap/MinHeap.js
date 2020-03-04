import Heap from './Heap';

export default class MinHeap extends Heap {
    /**
     * 
     * @param {*} first 
     * @param {*} second
     * @returns {boolean} 
     */
    pairIsInCorrectOrder(first, second) {
        return this.compare.lessThanOrEqual(first, second);
    }
}