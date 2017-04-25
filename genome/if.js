const random = require('../modules/random.js');

module.exports = {
	name: 'IF',
	needGens: 4,

	create() {
		return `${this.name} # # # #`;
	},

	getRunFunction(gen, individual) {
		const Gen = require('./gen.js');
		const ifArg = Gen.getRunFunction(individual[parseInt(gen[1])], individual);
		const negativeArg = Gen.getRunFunction(individual[parseInt(gen[2])], individual);
		const nullArg = Gen.getRunFunction(individual[parseInt(gen[3])], individual);
		const positiveArg = Gen.getRunFunction(individual[parseInt(gen[4])], individual);

		return function (args) {
			const result = ifArg(args);

			if (result < 0) {
				return negativeArg(args);
			} else if (result === 0) {
				return nullArg(args);
			} else if (0 < result) {
				return positiveArg(args)
			}
		}
	}
};
