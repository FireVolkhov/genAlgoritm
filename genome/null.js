const random = require('../modules/random.js');

module.exports = {
	name: 'NULL',
	needGens: 0,

	create() {
		return `${this.name}`
	},

	getRunFunction(gen, individual) {
		return function (args) {
			return 0;
		}
	}
};
