const _ = require('lodash');
const random = require('./random.js');
const utils = require('./utils.js');

const Individual = require('./individual.js');

const defaultOptions = {
	count: NaN,
	enterGens: NaN,

	survivalPercent: 0.5,
	startGensCount: [1, 10]
};

module.exports = class Population {
	constructor (enterGen, genome, options) {
		this.enterGen = enterGen;
		this.genome = genome;
		this.options = _.extend(_.clone(defaultOptions), _.clone(options));

		this.individuals = [];
		this.habitat = [];

		this.create();
	}

	create() {
		this.individuals = utils.array(this.options.count).map(() => {
			const individual = new Individual();

			utils.array(this.options.enterGens).map((x, i) => {
				individual.addEnterGen(
					new this.enterGen()
						.setGens([() => this.habitat[i]])
				);
			});

			const countGens = random.int(this.options.startGensCount[0], this.options.startGensCount[1]);

			utils.array(countGens).map(() => {
				individual.addGen(new (random.getItem(this.genome))());
			});

			return individual;
		});
	}

	selection(testFunction) {
		this.individuals =
			this.individuals.map((individual) => {
				return [
					individual,
					testFunction((...args) => {
						this.habitat = args;
						return individual.run();
					}),
					individual.gens.length];
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
				.slice(0, Math.round(this.options.survivalPercent * this.individuals.length))
				.map((x) => x[0]);

		//this.individuals.forEach((x) => {
		//	console.log(x.toString());
		//});
		//console.log(this.individuals[0].toString());

		return this;
	}

	mutation() {
		while (this.individuals.length < this.options.count) {
			this.individuals.push(
				random.getItem(this.individuals).mutation()
			);
		}
	}
};
