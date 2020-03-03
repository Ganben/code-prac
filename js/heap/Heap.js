import Comparator from '../comparator/Comparator';

/**
 * 
 */
export default class Heap {
    /**
     * @constructs Heap
     * @param {Function} compareFunc 
     */
    constructor(compareFunc) {
        if (new.target === Heap) {
            throw new TypeError('cannot construct Heap instance directly');
        }
        //
        this.heapContainer = [];
        this.compare = new Comparator(compareFunc);
    }
    /**
     * 
     * @param {number} parentIndex
     * @returns {number} 
     */
    getLeftChildIndex(parentIndex) {
        return ( 2 * parentIndex ) + 1;
    }
    /**
     * 
     * @param {number} parentIndex
     * @returns {number} 
     */
    getRightChildIndex(parentIndex) {
        return ( 2 * parentIndex ) + 2;
    }
    /**
     * 
     * @param {number} childIndex
     * @returns {boolean} 
     */
    hasParent(childIndex) {
        return this.getParentIndex(childIndex) >= 0;
    }
    /**
     * 
     * @param {number} parentIndex
     * @returns {boolean} 
     */
    hasLeftChild(parentIndex) {
        return this.getLeftChildIndex(parentIndex) < this.heapContainer.length;
    }
    /**
     * 
     * @param {number} parentIndex
     * @returns {boolean} 
     */
    hasRightChild(parentIndex) {
        return this.getRightChildIndex(parentIndex) < this.heapContainer.length;
    }
    /**
     * 
     * @param {number} parentIndex 
     * @returns {*}
     */
    leftChild(parentIndex) {
        return this.heapContainer[this.getLeftChildIndex(parentIndex)];
    }
    /**
     * 
     * @param {number} parentIndex
     * @returns {*} 
     */
    rightChild(parentIndex) {
        return this.heapContainer[this.getRightChildIndex(parentIndex)];
    }
    /**
     * 
     * @param {number} childIndex
     * @returns {*} 
     */
    parent(childIndex) {
        return this.heapContainer[this.getParentIndex(childIndex)];
    }
    /**
     * 
     * @param {number} indexOne 
     * @param {number} indexTwo 
     */
    swap(indexOne, indexTwo) {
        const tmp = this.heapContainer[indexTwo];
        this.heapContainer[indexTwo] = this.heapContainer[indexOne];
        this.heapContainer[indexOne] = tmp;
    }
    /**
     * @returns {*}
     */
    peek() {
        if (this.heapContainer.length === 0) {
            return null;
        }
        return this.heapContainer[0];
    }
    /**
     * @returns {*}
     */
    poll() {
        if (this.heapContainer.length === 0) {
            return null;
        }
        if (this.heapContainer.length === 1) {
            return this.heapContainer.pop();
        }
        const item = this.heapContainer[0];
        //
        this.heapContainer[0] = this.heapContainer.pop();
        this.heapifyDown();
        return item;
    }
    /**
     * 
     * @param {*} item
     * @returns {Heap} 
     */
    add(item) {
        this.heapContainer.push(item);
        this.heapifyDown();
        return this;
    }
    /**
     * 
     * @param {*} item 
     * @param {Comparator} comparator 
     * @returns {Heap}
     */
    remove(item, comparator=this.compare) {
        //
        const numberOfItemsToRemove = this.getLeftChildIndex(item, comparator).length;

        for (let iteration = 0; iteration < numberOfItemsToRemove; iteration += 1) {
            // we find index to remove
            const indexToRemove = this.getLeftChildIndex(item, comparator).pop();
            //  need to remove last child
            if (indexToRemove === (this.heapContainer.length - 1)) {
                this.heapContainer.pop();
            } else {
                // move last ele to vacant
                this.heapContainer[indexToRemove] = this.heapContainer.pop();
                // get parent
                const parentItem = this.parent(indexToRemove);
                // if there is no parent
                if (
                    this.hasLeftChild(indexToRemove)
                    && (
                        !parentItem
                        || this.pairIsInCorrectOrder(parentItem, this.heapContainer[indexToRemove])
                    )
                ) {
                    this.heapifyDown(indexToRemove);
                } else {
                    this.heapifyUp(indexToRemove);
                }
            }

        }
        return this;
    }
    /**
     * 
     * @param {*} item 
     * @param {Comparator} {comparator}
     * @returns {Number[]} 
     */
    find(item, comparator = this.compare) {
        const foundItemIndices = [];
        for (let itemIndex = 0; itemIndex < this.heapContainer.length; itemIndex += 1) {
            if (comparator.equal(item, this.heapContainer[itemIndex])) {
                foundItemIndices.push(itemIndex);
            }
        }
        return foundItemIndices;
    }
    /**
     * @returns {boolean}
     */
    isEmpty() {
        return !this.heapContainer.length;
    }
    /**
     * @returns {string}
     */
    toString() {
        return this.heapContainer.toString();
    }
    /**
     * 
     * @param {number} customStartIndex 
     */
    heapifyUp(customStartIndex) {
        // take last, lift it
        let currentIndex = customStartIndex || this.heapContainer.length - 1;
        while (
            this.hasParent(currentIndex)
            && !this.pairIsInCorrectOrder(this.parent(currentIndex), this.heapContainer[currentIndex])
        ) {
            this.swap(currentIndex, this.getParentIndex(currentIndex));
            currentIndex = this.getParentIndex(currentIndex);
        }
    }
    /**
     * 
     * @param {number} customStartIndex 
     */
    heapifyDown(customStartIndex = 0) {
        // compare the parent to child and swap
        let currentIndex = customStartIndex;
        let nextIndex = null;
        while (this.hasLeftChild(currentIndex)) {
            if (
                this.hasRightChild(currentIndex)
                && this.pairIsInCorrectOrder(this.rightChild(currentIndex), this.leftChild(currentIndex))
            ) {
                nextIndex = this.getRightChildIndex(currentIndex);
            } else {
                nextIndex = this.getLeftChildIndex(currentIndex);
            }
            this.swap(currentIndex, nextIndex);
            currentIndex = nextIndex;
        }
        
    }
    /**
     * not yet implt ???
     * @param {*} first 
     * @param {*} second 
     * @returns {boolean}
     */
    pairIsInCorrectOrder(first, second) {
        throw new Error(`
        You have to impl compare method for ${first} and ${second}
        `);
    }
}