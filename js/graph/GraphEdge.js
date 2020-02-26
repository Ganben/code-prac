export default class GraphEdge {
    /**
     * @param {GraphVertex} startVertex
     * @param {GraphVertex} endVertex
     * @param {number} [weight=1]
     */
    constructor(startVertex, endVertex, weight=1) {
        this.startVertex = startVertex;
        this.endVertex = endVertex;
        this.weight = weight;
    }
    /**
     * @returns {string}
     */
    getKey() {
        const startVertexKey = this.startVertex.getKey();
        const endVertexKey = this.endVertex.getKey();
        return `${startVertexKey}_${endVertexKey}`;
    }
    /**
     * @returns {GraphEdge}
     */
    reverse() {
        const tmp = this.startVertex;
        this.startVertex = this.endVertex;
        this.endVertex = tmp;
        return this;
    }
    /**
     * @returns {string}
     */
    toString() {
        return this.getKey();
    }
}