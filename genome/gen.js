const _ = require('lodash');
const enterGen = require('./enter.js');
const genome = require('./index.js');
const genomeNames = {
	[enterGen.name]: enterGen
};

genome.map((x) => genomeNames[x.name] = x);

module.exports = {
	/**
	 *
	 * @param {Array} gen
	 * @param {Array} individual
	 * @returns {function}
	 */
	getRunFunction(gen, individual) {
		const name = gen[0];
		const genObj = genomeNames[name];

		return genObj.getRunFunction(gen, individual);
	}
};
