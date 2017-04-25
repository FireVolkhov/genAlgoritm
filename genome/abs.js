const random = require('../modules/random.js');

module.exports = {
	name: 'ABS',
	needGens: 1,

	create() {
		return `${this.name} #`;
	},

	getRunFunction(gen, individual) {
		const Gen = require('./gen.js');
		const firstArg = Gen.getRunFunction(individual[parseInt(gen[1])], individual);

		return function (args) {
			return Math.abs(firstArg(args));
		}
	}
};
