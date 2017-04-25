const random = require('../modules/random.js');

module.exports = {
	name: 'MULT',

	create() {
		return `${this.name} # #`
	},

	getRunFunction(gen, individual) {
		const Gen = require('./gen.js');
		const firstArg = Gen.getRunFunction(individual[parseInt(gen[1])], individual);
		const secondArg = Gen.getRunFunction(individual[parseInt(gen[2])], individual);

		return function (args) {
			return firstArg(args) * secondArg(args);
		}
	}
};
