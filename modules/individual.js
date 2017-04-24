const _ = require('lodash');
const random = require('./random.js');
const utils = require('./utils.js');
const genome = require('../genome');
const Gen = require('../genome/gen.js');
const genomeNames = genome.map((x) => x.name);

module.exports = {
	create() {
		return '';
	},

	/**
	 *
	 * @param {string} individual
	 * @param {string} gen
	 * @returns {string}
	 */
	addGen(individual, gen) {
		const array = this.stringToArray(individual);

		gen = gen.split(' ');
		gen = gen.map((x) => {
			if (x === '#') {
				return '' + random.int(0, array.length - 1);
			}

			return x;
		});

		array.push(gen);

		return this.arrayToString(array);
	},

	/**
	 *
	 * @param {string} individual
	 * @returns {Array}
	 */
	stringToArray(individual) {
		if (individual === '') {
			return [];
		}

		return individual
			.split('\n')
			.map((x) => x.split(' '));
	},

	/**
	 *
	 * @param {Array} individual
	 * @returns {string}
	 */
	arrayToString(individual) {
		return individual
			.map((x) => x.join(' '))
			.join('\n');
	},

	/**
	 *
	 * @param {string} individual
	 * @returns {Function}
	 */
	getRunFunction(individual) {
		const array = this.stringToArray(individual);

		return Gen.getRunFunction(_.last(array), array);
	},

	/**
	 *
	 * @param {string} individual
	 * @returns {string}
	 */
	mutation(individual) {
		const array = this.stringToArray(individual);
		const newArray = random.getItem([
			() => this.mutationAddGen(array),
			() => this.mutationModifyGen(array),
			() => this.mutationRemoveGen(array)
		])();

		return this.arrayToString(newArray);
	},

	/**
	 *
	 * @param {Array} individual
	 * @returns {Number}
	 */
	countEnterGens(individual) {
		return individual.filter((x) => x[0] === 'ENTER').length;
	},

	/**
	 *
	 * @param {Array} individual
	 * @returns {Array}
	 */
	mutationAddGen(individual) {
		let newGen = random.getItem(genome).create().split(' ');
		const countEnterGens = this.countEnterGens(individual);
		const position = random.int(countEnterGens, individual.length - 1);

		newGen = newGen.map((x) => {
			if (x === '#') {
				return '' + random.int(0, position - 1);
			}

			return x;
		});

		individual.splice(position, 0, newGen);

		const target = random.getItem(individual.slice(position + 1));
		const targetPos = random.int(1, target.length - 1);
		target[targetPos] = '' + position;

		return individual;
	},

	/**
	 * @param {Array} individual
	 * @returns {Array}
	 */
	mutationModifyGen(individual) {
		let changeType = !!random.getItem([0, 1]);
		const changeArgs = !!random.getItem([0, 1]);

		const countEnterGens = this.countEnterGens(individual);
		const position = random.int(countEnterGens, individual.length - 1);
		const target = individual[position];

		if (!changeType && !changeArgs) {
			changeType = true;
		}

		if (changeType) {
			const names = genomeNames.filter((x) => x !== target[0]);
			const newGen = random.getItem(names);

			target[0] = newGen;
		}

		if (changeArgs) {
			const targetPos = random.int(1, target.length - 1);
			const newArg = random.int(0, position - 1);

			target[targetPos] = '' + newArg;
		}

		return individual;
	},

	/**
	 * @param {Array} inputIndividual
	 * @returns {Array}
	 */
	mutationRemoveGen(inputIndividual) {
		let individual = _.cloneDeep(inputIndividual);
		const countEnterGens = this.countEnterGens(individual);
		const position = random.int(countEnterGens, individual.length - 1);

		if (individual.length === (countEnterGens + 1)) {
			return inputIndividual;
		}

		individual.splice(position, 1);

		individual = individual.map((gen, i) => {
			if (i >= position) {
				return gen.map((x, j) => {
					if (j > 0) {
						return '' + (Math.max(0, parseInt(x) - 1));
					} else {
						return x;
					}
				});
			} else {
				return gen;
			}
		});

		return individual;
	}
};
