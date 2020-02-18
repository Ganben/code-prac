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
    find() {

    }
    /**
     * @return {LinkedListNode}
     */
    deleteTail() {

    }
    /**
     * @return {LinkedListNode}
     */
    deleteHead() {

    }

    /**
     * @param {*[]} values - array of values
     * @return {LinkedList} 
     */
    fromArray(values) {

    }
    /**
     * @return {LinkedListNode[]}
     */
    toArray() {

    }
    /**
     * @param {function} [callback]
     * @return {string}
     */
    toString(callback) {

    }
    /**
     * @return {LinkedList}
     */
    reverse() {

    }
}
