/**
 * 
 * @param {LinkedListNode} node 
 * @param {function} callback 
 */
function reverseTraversalRecursive(node, callback) {
    if (node) {
        reverseTraversalRecursive(node.next, callback);
        callback(node.value);
    }
}
/**
 * 
 * @param {LinkedList} linkedList 
 * @param {function} callback 
 */
export default function reverseTraversal(linkedList, callback) {
    reverseTraversalRecursive(linkedList.head, callback);
}