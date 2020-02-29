import HashTable from '../hashtable/HashTable';

export default class TrieNode {
    /**
     * @param {string} character
     * @param {boolean} isCompleteWord
     */
    constructor(character, isCompleteWord = false) {
        this.character = character;
        this.isCompleteWord = isCompleteWord;
        this.children = new HashTable();
    }
    /**
     * 
     * @param {string} character
     * @returns {TrieNode} 
     */

    getChild(character) {
        return this.children.get(character);
    }
    /**
     * 
     * @param {string} character 
     * @param {boolean} isCompleteWord 
     */
    addChild(character, isCompleteWord = false) {
        if (!this.children.has(character)) {
            this.children.set(character, new TrieNode(character, isCompleteWord));
        }
        const childNode = this.children.get(character);
        //
        childNode.isCompleteWord = childNode.isCompleteWord || isCompleteWord;
        return childNode;
    }
    /**
     * 
     * @param {string} character
     * @returns {boolean} 
     */
    removeChild(character) {
        const childNode = this.getChild(character);
        //
        if (
            childNode
            && !childNode.isCompleteWord
            && !childNode.hasChildren()
            ) {
                this.children.delete(character);
        }
        return this;
    }
    /**
     * 
     * @param {string} character
     * @returns {boolean} 
     */
    hasChild(character) {
        return this.children.has(character);
    }
    /**
     * @returns {boolean}
     */
    hasChildren() {
        return this.children.getKeys().length !== 0;
    }
    /**
     * @returns {string[]}
     */
    suggestChildren() {
        return [...this.children.getKeys()];
    }
    /**
     * @returns {string}
     */
    toString() {
        let childrenAsString = this.suggestChildren().toString();
        childrenAsString = childrenAsString ? `${childrenAsString}` : '';
        const isCompleteString = this.isCompleteWord ? '*' : '';
        return `${this.character}${isCompleteString}${childrenAsString}`;
    }
}