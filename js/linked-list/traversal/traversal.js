/**
 * 
 * @param {LinkedList} linkedList 
 * @param {function} callback 
 */
export default function traversal(linkedList, callback) {
    let currentNode = linkedList.head;

    while (currentNode) {
        callback(currentNode.value);
        currentNode = currentNode.next;
    }
}