export default class Graph {
    /**
     * @param {boolean} isDirected
     */
    constructor(isDirected = false) {
        this.vertices = {};
        this.edges = {};
        this.isDirected = isDirected;
    }
    /**
     * 
     * @param {GraphVertex} new vert
     * @returns {Graph} 
     */
    addVertex(vert) {
        this.vertices[vert.getKey()] = vert;
        return this;
    }
    /**
     * 
     * @param {string} vertKey
     * @returns {GraphVertex} 
     */
    getVertexByKey(vertKey) {
        return this.vertices[vertKey];

    }
    /**
     * @param {GraphVertex} vert
     * @returns {GraphVertex[]}
     */
    getNeighbors(vert) {
        return this.addVertex.getNeighbors();
    }
    /**
     * @returns {GraphVertex[]}
     */
    getAllVertices() {
        //
        return Object.values(this.vertices);

    }
    /**
     * @returns {GraphEdge[]}
     */
    getAllEdges() {
        return Object.values(this.edges);

    }
    /**
     * 
     * @param {GraphEdge} edge
     * @returns {Graph} 
     */
    addEdge(edge) {


    }
    /**
     * 
     * @param {GraphEdge} edge
     *  
     */
    deleteEdge(edge) {

    }
    /**
     * @param {GraphVertex} startVert
     * @param {GraphVertex} endVert
     * @returns {(GraphEdge|null)}
     */
    findEdge(startVert, endVert) {

    }
    /**
     * @returns {number}
     */
    getWeight() {

    }
    /**
     * @returns {Graph}
     */
    reverse() {

    }
    /**
     * @returns {object}
     */
    getVerticesIndices() {

    }
    /**
     * @returns {*[][]}
     */
    getAdjacencyMatrix() {

    }
    /**
     * @returns {string}
     */
    toString() {

    }
}