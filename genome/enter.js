const Gen = require('./gen.js');

module.exports = {
	name: 'ENTER',
	needGens: 0,

	/**
	 *
	 * @param {int} index
	 * @returns {string}
	 */
	create(index) {
		return `${this.name} ${index}`;
	},

	getRunFunction(gen, individual) {
		const index = parseInt(gen[1]);

		return function (args) {
			return args[index];
		};
	}
};
