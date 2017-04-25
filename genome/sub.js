const random = require('../modules/random.js');

module.exports = {
	name: 'SUB',

	create() {
		return `${this.name} # #`
	},

	getRunFunction(gen, individual) {
		const Gen = require('./gen.js');
		const firstArg = Gen.getRunFunction(individual[parseInt(gen[1])], individual);
		const secondArg = Gen.getRunFunction(individual[parseInt(gen[2])], individual);
		// console.log(individual, '>>>', individual[parseInt(gen[1])]);

		return function (args) {
			// console.log(individual, '>>>', args, '>>>', firstArg(args) - secondArg(args));
			return firstArg(args) - secondArg(args);
		}
	}
};
