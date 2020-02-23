import BinarySearchTree from '../binary-search-tree/BinarySearchTree';
import BinarySearchTreeNode from '../binary-search-tree/BinarySearchTreeNode';

//
const RED_BLACK_TREE_COLORS = {
    red: 'red',
    black: 'black',
};

// color property
const COLOR_PROP_NAME = 'color';

export default class RedBlackTree extends BinarySearchTree {
    /**
     * 
     * @param {*} value
     * @return {BinarySearchTreeNode} 
     */
    insert(value) {
        const insertedNode = super.insert(value);

        //
        if (this.nodeComparator.equal(insertedNode, this.root)) {
            // root always black
            this.makeNodeBlack(insertedNode);
        } else {
            // make all newly inserted node to be red
            this.makeNodeRed(insertedNode);
        }
        // check all condition and balance
        this.balance(insertedNode);
        return insertedNode;
    }
    /**
     * 
     * @param {*} value
     * @return {boolean} 
     */
    remove(value) {
        //
    }
    /**
     * 
     * @param {BinarySearchTreeNode} node 
     */
    balance(node) {
        // root, do nothing
        if (this.nodeComparator.equal(node, this.root)) {
            return;
        }
        // parent black, do nothing
        if (this.isNodeBlack(node.parent)) {
            return;
        }
        //
        const grandParent = node.parent.parent;
        if (node.uncle && this.isNodeRed(node.uncle)) {
            // red uncle, recolor
            this.makeNodeBlack(node.uncle);
            this.makeNodeBlack(node.parent);
            if (!this.nodeComparator.equal(grandParent, this.root)) {
                //
                this.makeNodeRed(grandParent);
            } else {
                return;
            }
            // further balance grandParent
            this.balance(grandParent);

        } else if (!node.uncle || this.isNodeBlack(node.uncle)) {
            // black or absent uncle, do rotations
            if (grandParent) {
                // grand parent
                let newGrandParent;
                if (this.nodeComparator.equal(grandParent.left, node.parent)) {
                    //left
                    if (this.nodeComparator.equal(node.parent.left, node)) {
                        //
                        newGrandParent = this.leftLeftRotation(grandParent);
                    } else {
                        //
                        newGrandParent = this.leftRightRotation(grandParent);
                    }

                } else {
                    //right
                    if (this.nodeComparator.equal(node.parent.right, node)) {
                        //
                        newGrandParent = this.rightRightRotation(grandParent);
                    } else {
                        //
                        newGrandParent = this.rightLeftRotation(grandParent);
                    }
                } 
                // set new as root
                if (newGrandParent && newGrandParent.parent === null) {
                    this.root = newGrandParent;
                    //
                    this.makeNodeBlack(this.root);
                }
                //
                this.balance(newGrandParent);
            }
        }
    }
    leftLeftRotation(grandParentNode) {

    }
    leftRightRotation(grandParentNode) {

    }
    rightLeftRotation(grandParentNode) {

    }
    rightRightRotation(grandParentNode) {

    }
    makeNodeRed(node) {

    }
    makeNodeBlack(node) {

    }
    isNodeRed(node) {

    }
    isNodeBlack(node) {

    }
    isNodeColored(node) {

    }
    swapNodeColors(firstNode, secondNode) {
        
    }

}