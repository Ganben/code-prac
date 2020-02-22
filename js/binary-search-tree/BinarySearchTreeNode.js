// import HashTable from '../hashtable/HashTable';
import Comparator from '../comparator/Comparator';
import BinaryTreeNode from '../tree/BinaryTreeNode';

export default class BinarySearchTreeNode extends BinaryTreeNode {
    /**
     * 
     * @param {*} value of node
     * @param {function} compareFunction 
     */
    constructor(value = null, compareFunction = undefined) {
        super(value);
        //
        this.compareFunction = compareFunction;
        this.nodeValueComparator = new Comparator(compareFunction);
    }
    /**
     * 
     * @param {*} value 
     * @return {BinarySearchTreeNode}
     */
    insert(value) {
        if (this.nodeValueComparator.equal(this.value, null)) {
            this.value = value;
            return this;
        }

        if (this.nodeValueComparator.lessThan(value, this.value)) {
            if (this.left) {
                return this.left.insert(value);
            }
            const newNode = new BinarySearchTreeNode(value, this.compareFunction);
            this.setLeft(newNode);

            return newNode;
        }

        if (this.nodeValueComparator.greaterThan(value, this.value)) {
            if (this.right) {
                return this.right.insert(value);
            }

            const newNode = new BinarySearchTreeNode(value, this.compareFunction);
            this.setRight(newNode);

            return newNode;
        }
        return this;
    }
    /**
     * 
     * @param {*} value 
     * @return {boolean}
     */
    contains(value) {
        return !!this.find(value);
    }
    /**
     * 
     * @param {*} value 
     * @return {boolean}
     */
    remove(value) {
        const nodeToRemove = this.find(value);

        if (!nodeToRemove) {
            throw new Error('item not found in tree');

        }
        const { parent } = nodeToRemove;

        if (!nodeToRemove.left && !nodeToRemove.right) {
            //leaf, no child
            if ( parent ) {
                parent.removeChild(nodeToRemove);
            } else {
                nodeToRemove.setValue(undefined);
            }
        } else if (nodeToRemove.left && nodeToRemove.right) {
            //replace node's value by find next bigger in children
            const nextBiggerNode = nodeToRemove.right.findMin();
            if (!this.nodeComparator.equal(nextBiggerNode, nodeToRemove.right)) {
                this.remove(nextBiggerNode.value);
                nodeToRemove.setValue(nextBiggerNode.value);
            } else {
                // next right without left child
                nodeToRemove.setValue(nodeToRemove.right.value);
                nodeToRemove.setRight(nodeToRemove.right.right);
            }

        } else {
            // only one child
            /** @var BinarySearchTreeNode */
            const childNode = nodeToRemove.left || nodeToRemove.right;
            if (parent) {
                parent.replaceChild(nodeToRemove, childNode);
            } else {
                BinaryTreeNode.copyNode(childNode, nodeToRemove);
            }
        }
        // clear parent
        nodeToRemove.parent = null;
        return true;

    }
    /**
     * @return {BinarySearchTreeNode}
     */
    findMin() {
        if (!this.left) {
            return this;
        }
        return this.left.findMin();

    }
}


