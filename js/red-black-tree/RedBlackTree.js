import BinarySearchTree from '../binary-search-tree/BinarySearchTree';
import BinarySearchTreeNode from '../binary-search-tree/BinarySearchTreeNode';
import BinaryTreeNode from '../tree/BinaryTreeNode';

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
    /**
     * 
     * @param {BinarySearchTreeNode|BinaryTreeNode} grandParentNode 
     * @return {BinarySearchTree}
     */
    leftLeftRotation(grandParentNode) {
        // memorize parent of grand parent;
        const grandGrandPa = grandParentNode.parent;
        // 
        let grandParentNodeIsLeft;
        if (grandGrandPa) {
            grandParentNodeIsLeft = this.nodeComparator.equal(grandGrandPa.left, grandParentNode);
        }
        // memo grandpa's left
        const parentNode = grandParentNode.left;
        // memo parent's right, transfer it to grandpa's left sub
        const parentRightNode = parentNode.right;
        // make grandpa to right child of parent
        parentNode.setRight(grandParentNode);
        // move child's right sub to grandpa's left sub
        grandParentNode.setLeft(parentRightNode);
        // put parentNode in place of grandParentNode
        if (grandGrandPa) {
            if (grandParentNodeIsLeft) {
                grandGrandPa.setLeft(parentNode);
            } else {
                grandGrandPa.setRight(parentNode);
            }
        } else {
            //make it root
            parentNode.parent = null;
        }
        //swap colors 
        this.swapNodeColors(parentNode, grandParentNode);
        //return new root
        return parentNode;
    }
    /**
     * 
     * @param {BinarySearchTreeNode|BinaryTreeNode} grandParentNode
     * @return {BinarySearchTreeNode} 
     */
    leftRightRotation(grandParentNode) {
        // memo left and left-right
        const parentNode = grandParentNode.left;
        const childNode = parentNode.right;
        // memo child left 
        // reassign to parent's right
        const childLeftNode = childNode.left;
        // make parentnode be left child of child node
        childNode.setLeft(parentNode);
        // move child's left sub to parent's right sub
        parentNode.setRight(childLeftNode);
        // put left-right node in place of left
        grandParentNode.setLeft(childNode);
        // then do left-left rotation
        return this.leftLeftRotation(grandParentNode);
    }
    /**
     * 
     * @param {BinarySearchTreeNode|BinaryTreeNode} grandParentNode 
     * @return {BinarySearchTree}
     */
    rightLeftRotation(grandParentNode) {
        // memo right and right-left
        const newParent = grandParentNode.right;
        const childNode = newParent.left;
        // memo child right and sub
        // re-assign to parent's left sub
        const childRightNode = childNode.right;
        // make parent right child of child
        childNode.setRight(newParent);
        // move child's right sub to parent left sub
        newParent.setLeft(childRightNode);
        // do right-right rotation
        return this.rightRightRotation(grandParentNode);
    }
    /**
     * 
     * @param {BinarySearchTreeNode|BinaryTreeNode} grandParentNode 
     * @return {BinarySearchTree}
     */
    rightRightRotation(grandParentNode) {
        // memo parent of input node
        const oldParent = grandParentNode.parent;
        // check sibling
        let grandParentNodeIsLeft;
        if (oldParent) {
            grandParentNodeIsLeft = this.nodeComparator.equal(oldParent.left, grandParentNode);
        }
        // memo input's right to be new root
        const newParent = grandParentNode.right;
        // memo parent's left then transfer to grand parent's right sub
        const leftNode = newParent.left;
        // make input to left child of new parent
        newParent.setLeft(grandParentNode);
        // transfer all left to right sub of input
        grandParentNode.setRight(leftNode);
        // put new parent in place of input
        if (oldParent) {
            if (grandParentNodeIsLeft) {
                oldParent.setLeft(newParent);
            } else {
                oldParent.setRight(newParent);
            }
        } else {
            // make new parent root
            newParent.parent = null;
        }
        // swap color to input
        this.swapNodeColors(newParent, grandParentNode);
        return newParent;
    }
    /**
     * 
     * @param {BinarySearchTreeNode|BinaryTreeNode} node 
     * @return {BinarySearchTreeNode}
     */
    makeNodeRed(node) {
        node.meta.set(COLOR_PROP_NAME, RED_BLACK_TREE_COLORS.red);
        return node;

    }
    /**
     * 
     * @param {BinarySearchTreeNode|BinaryTreeNode} node 
     * @return {BinarySearchTreeNode}
     */
    makeNodeBlack(node) {
        node.meta.set(COLOR_PROP_NAME, RED_BLACK_TREE_COLORS.black);
        return node;
    }
    /**
     * 
     * @param {BinarySearchTreeNode|BinaryTreeNode} node
     * @return {boolean} 
     */
    isNodeRed(node) {
        return node.meta.get(COLOR_PROP_NAME) === RED_BLACK_TREE_COLORS.red;
    }
    /**
     * 
     * @param {BinarySearchTreeNode|BinaryTreeNode} node 
     * @return {boolean}
     */
    isNodeBlack(node) {
        return node.meta.get(COLOR_PROP_NAME) === RED_BLACK_TREE_COLORS.black;
    }/**
     * 
     * @param {BinarySearchTreeNode|BinaryTreeNode} node
     * @return {boolean} 
     */
    isNodeColored(node) {
        return this.isNodeBlack(node) || this.isNodeRed(node);
    }
    /**
     * 
     * @param {BinarySearchTreeNode|BinaryTreeNode} firstNode 
     * @param {BinarySearchTreeNode|BinaryTreeNode} secondNode 
     */
    swapNodeColors(firstNode, secondNode) {
        const firstColor = firstNode.meta.get(COLOR_PROP_NAME);
        const secondColor = secondColor.meta.get(COLOR_PROP_NAME);

        firstNode.meta.set(COLOR_PROP_NAME, secondColor);
        secondNode.meta.set(COLOR_PROP_NAME, firstColor);
    }

}