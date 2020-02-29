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
        let startVert = this.getVertexByKey(edge.startVert.getKey());
        let endVert = this.getVertexByKey(edge.endVert.getKey());
        //
        if (!startVert) {
            this.addVertex(edge.startVert);
            startVert = this.getVertexByKey(edge.endVert.getKey());
        }
        //
        if (this.edges[edge.getKey()]) {
            throw new Error('Edge has already beed added before');
        } else {
            this.edges[edge.getKey()] = edge;
        }
        // add edge to vert
        if (this.isDirected) {
            //
            startVert.addEdge(edge);
        } else {
            //
            startVert.addEdge(edge);
            endVert.addEdge(edge);
        }
        return this;
    }
    /**
     * 
     * @param {GraphEdge} edge
     *  
     */
    deleteEdge(edge) {
        if (this.edges[edge.getKey()]) {
            delete this.edges[edge.getKey()];
        } else {
            throw new Error('edge not found in graph');
        }
        //
        const startVert = this.getVertexByKey(edge.startVert.getKey());
        const endVert = this.getVertexByKey(edge.endVert.getKey());
        startVert.deleteEdge(edge);
        endVert.deleteEdge(edge);

    }
    /**
     * @param {GraphVertex} startVert
     * @param {GraphVertex} endVert
     * @returns {(GraphEdge|null)}
     */
    findEdge(startVert, endVert) {
        const vertex = this.getVertexByKey(startVert.getKey());
        if (!vertex) {
            return null;
        }
        return vertex.findEdge(endVert);
    }
    /**
     * @returns {number}
     */
    getWeight() {
        return this.getAllEdges().reduce((weight, graphEdge) => {
            return weight + graphEdge.weight;
        }, 0)
    }
    /**
     * @returns {Graph}
     */
    reverse() {
        /**@param {GraphEdge} edge */
        this.getAllEdges().forEach((edge) => {
            //
            this.deleteEdge(edge);
            //
            edge.reverse();
            //
            this.addEdge(edge);
        });
        return this;

    }
    /**
     * @returns {object}
     */
    getVerticesIndices() {
        const verticesIndices = [];
        this.getAllVertices().forEach((vertiex, index) => {
            verticesIndices[vertex.getKey()] = index;
        });
        return verticesIndices;
    }
    /**
     * @returns {*[][]}
     */
    getAdjacencyMatrix() {
        const vertices = this.getAllVertices();
        const verticesIndices = this.getVerticesIndices();
        //
        //
        const adjacentMatrix = Array(vertices.length).fill(null).map(() => {
            return Array(vertices.length).fill(Infinity);
        });
        //
        vertices.forEach((vertex, vertexIndex) => {
            vertex.getNeighbors().forEach((neighbor) => {
                const neighborIndex = verticesIndices[neighbor.getKey()];
                adjacentMatrix[vertexIndex][neighborIndex] = this.findEdge(vertex, neighbor).weight;
            });
        });
        return adjacentMatrix;

    }
    /**
     * @returns {string}
     */
    toString() {
        return Object.keys(this.vertices).toString();

    }
}