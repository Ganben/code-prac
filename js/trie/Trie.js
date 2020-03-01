import TrieNode from './TrieNode';

//
const HEAD_CHARACTER = '*';

export default class Trie {
    constructor() {
        this.head = new TrieNode(HEAD_CHARACTER);
    }
    /**
     * 
     * @param {string} word
     * @returns {Trie} 
     */
    addWord(word) {
        const characters = Array.from(word);
        let currentNode = this.head;

        for (let charIndex = 0; charIndex < characters.length; charIndex += 1) {
            const isComplete = charIndex === characters.length - 1;
            currentNode = currentNode.addChild(characters[charIndex], isComplete);
        }
        return this;
    }
    /**
     * 
     * @param {string} word
     * @returns {Trie} 
     */
    deleteWord(word) {
        const depthFirstDelete = (currentNode, charIndex = 0) => {
            if (charIndex >= word.length) {
                //
                return;
            }
            const character = word[charIndex];
            const nextNode = currentNode.getChild(character);

            if (nextNode == null) {
                //
                return;
            }
            //
            depthFirstDelete(nextNode, charIndex + 1);

            if (charIndex === (word.length - 1)) {
                nextNode.isCompleteWord = false;
            }
            // childnode delete if 
            // - nochildren
            // not complete
            currentNode.removeChild(character);
        };
        depthFirstDelete(this.head);
        return this;
    }
    /**
     * @param {string} word
     * @returns {string[]}
     */
    suggestNextCharacters(word) {
        const lastChar = this.getLastCharacterNode(word);

        if (!lastChar) {
            return null;
        }
        return lastChar.suggestChildren();
    }
    /**
     * 
     * @param {string} word
     * @returns {boolean} 
     */
    doesWordExist(word) {
        const lastChar = this.getLastCharacterNode(word);
        return !!lastChar && lastChar.isCompleteWord;
    }
    /**
     * 
     * @param {string} word
     * @returns {TrieNode} 
     */
    getLastCharacterNode(word) {
        const chars = Array.from(word);
        let currentNode = this.head;

        for (let charIndex = 0; charIndex < chars.length; charIndex += 1 ) {
            if (!currentNode.hasChild(chars[charIndex])) {
                return null;
            }
            currentNode = currentNode.getChild(chars[charIndex]);
        }
        return currentNode;
    }

}