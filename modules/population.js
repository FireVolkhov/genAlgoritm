const _ = require('lodash');
const random = require('./random.js');
const utils = require('./utils.js');
const genome = require('../genome');

const Individual = require('./individual.js');

module.exports = {
	/**
	 *
	 * @param {Object} EnterGen
	 * @param {Object} options
	 * @returns {Array}
	 */
	create(EnterGen, options) {
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

	selection(inputIndividuals, testFunction, options) {
		const ratingResult = inputIndividuals
			.map((individual) => {
				const run = Individual.getRunFunction(individual);

				return {
					individual: individual,
					rating: testFunction((...args) => run(args)),
					gensCount: individual.split('\n').length
				};
			})
			.sort((a, b) => {
				if (a.rating < b.rating) {
					return 1;
				} else if (a.rating > b.rating) {
					return -1;
				} else {
					return a.gensCount > b.gensCount ? 1 : -1;
				}
			})
		;

		const individuals = ratingResult
			.slice(0, Math.round(options.survivalPercent * inputIndividuals.length))
			.map((x) => x.individual)
		;

		return [individuals, ratingResult];
	},

	mutation(individuals, options) {
		while (individuals.length < options.count) {
			individuals.push(
					Individual.mutation(random.getItem(individuals))
			);
		}

		return individuals;
	}
};
