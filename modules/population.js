const _ = require('lodash');
const random = require('./random.js');
const utils = require('./utils.js');
const genome = require('../genome');

const Individual = require('./individual.js');

const defaultOptions = {
	count: NaN,
	enterGens: NaN,

	survivalPercent: 0.5,
	startGensCount: [1, 10]
};

module.exports = {
	/**
	 *
	 * @param {Object} EnterGen
	 * @param {Object} options
	 * @returns {Array}
	 */
	create(EnterGen, options) {
		options = _.extend(_.clone(defaultOptions), _.clone(options));

		return utils.array(options.count).map(() => {
			let individual = Individual.create();

			utils.array(options.enterGens).map((x, i) => {
				individual = Individual.addGen(individual, EnterGen.create(i))
			});

			const countGens = random.int(options.startGensCount[0], options.startGensCount[1]);

			utils.array(countGens).map(() => {
				individual = Individual.addGen(individual, random.getItem(genome).create());
			});

			return individual;
		});
	},

	selection(individuals, testFunction, options) {
		options = _.extend(_.clone(defaultOptions), _.clone(options));

		return individuals
				.map((individual) => {
					const run = Individual.getRunFunction(individual);

					return [
						individual,
						testFunction((...args) => run(args)),
						individual.split('\n').length];
				})
				.sort((a, b) => {
					if (a[1] < b[1]) {
						return 1;
					} else if (a[1] > b[1]) {
						return -1;
					} else {
						return a[2] > b[2];
					}
				})
				.slice(0, Math.round(options.survivalPercent * individuals.length))
				.map((x) => x[0])
		;
	},

	mutation(individuals, options) {
		options = _.extend(_.clone(defaultOptions), _.clone(options));

		while (individuals.length < options.count) {
			individuals.push(
					Individual.mutation(random.getItem(individuals))
			);
		}

		return individuals;
	}
};
