import Comparator from '../comparator/Comparator';
import LinkedListNode from './LinkedListNode';

export default class LinkedList {
    /**
     * @param {Function} [comparatorFunction]
     */
    constructor(comparatorFunction) {
        /** @var LinkedListNode */
        this.head = null;
        
        /** @var LinkedListNode */
        this.tail = null;

        this.compare = new Comparator(comparatorFunction);
    }
    /**
     * @param {*} value
     * @return {LinkedList}
     */
    prepend(value) {
        const newNode = new LinkedListNode(value, this.head);
        this.head = newNode;

        // tail ?
        if (!this.tail) {
            this.tail = newNode;
        }
        return this;
    }

    /**
     * @param {*} value
     * @return {LinkedList} 
     */
    append(value) {
        const newNode = new LinkedListNode(value);

        // head ?
        if (!this.head) {
            this.head = newNode;
            this.tail = newNode;
            return this;
        }

        // attach
        this.tail.next = newNode;
        this.tail = newNode;
        return this;
    }

    /**
     * @param {*} value
     * @return {LinkedListNode}
     */
    delete(value) {
        if (!this.head) {
            return null;
        }

        let deletedNode = null;
    }

    /**
     * @param {Object} findparams
     * @param {*} findparams.value
     * @param {function} [findparams.callback]
     * @return {LinkedListNode}
     */
    find({ value = undefined, callback = undefined}) {
        if (!this.head) {
            return null;
        }
        let currentNode = this.head;

        while (currentNode) {
            //if callback
            if (callback && callback(currentNode.value)) {
                return currentNode;
            }
            // if value
            if (value !== undefined && this.compare.equal(currentNode.value, value)) {
                return currentNode;
            }
        }
    }
    /**
     * @return {LinkedListNode}
     */
    deleteTail() {
        const deletedTail = this.tail;

        if (this.head === this.tail) {
            //
            this.head = null;
            this.tail = null;
            return deletedTail;
        }

        //
        let currentNode = this.head;
        while (currentNode.next) {
            if (!currentNode.next.next) {
                currentNode.next = null;
            } else {
                currentNode = currentNode.next;
            }
        }
        this.tail = currentNode;
        return deletedTail;

    }
    /**
     * @return {LinkedListNode}
     */
    deleteHead() {
        if (!this.head) {
            return null;
        }

        const deletedHead = this.head;

        if (this.head.next) {
            this.head = this.head.next;
        } else {
            this.head = null;
            this.tail = null;
        }

        return deletedHead;
    }

    /**
     * @param {*[]} values - array of values
     * @return {LinkedList} 
     */
    fromArray(values) {
        values.forEach(value => this.append(value));
        return this;

    }
    /**
     * @return {LinkedListNode[]}
     */
    toArray() {
        const nodes = [];

        let currentNode = this.head;
        while (currentNode) {
            nodes.push(currentNode);
            currentNode = currentNode.next;
        }
        return nodes;
    }
    /**
     * @param {function} [callback]
     * @return {string}
     */
    toString(callback) {
        return this.toArray().map(node => node.toString(callback)).toString();
    }
    /**
     * @return {LinkedList}
     */
    reverse() {
        let currNode = this.head;
        let prevNode = null;
        let nextNode = null;

        while (currNode) {
            //store next
            nextNode = currNode.next;

            // change 
            currNode.next = prevNode;

            // move prev and curr
            prevNode = currNode;
            currNode = nextNode;
        }

        this.tail = this.head;
        this.head = prevNode;
        return this;

    }
}
