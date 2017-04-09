const _ = require('lodash');
const random = require('./random.js');
const utils = require('./utils.js');
const genome = require('../genome');

module.exports = class Individual {
	constructor () {
		this.id = _.uniqueId('Individual_');
		this.gens = [];
	}

	get entersGens() {
		return this.gens.filter((g) => g.name === 'ENTER');
	}

	addEnterGen(gen) {
		gen.number = this.gens.push(gen) - 1;
		return this;
	}

	addGen(gen) {
		gen.number = this.gens.push(
			gen.setGens(
				utils.array(gen.needGens).map(() =>
					random.getItem(this.gens)
				)
			)
		) - 1;

		return this;
	}

	run() {
		return this.gens[this.gens.length - 1].run();
	}

	mutation() {
		return random.getItem([
			//() => this.mutationAddGen(),
			() => this.mutationModifyGen(),
			//() => this.mutationRemoveGen()
		])();
	}

	mutationAddGen() {
		const individual = this.clone();
		const newGen = new (random.getItem(genome))();
		const position = random.int(individual.entersGens.length, individual.gens.length - 1);

		individual.gens.splice(position, 0, newGen);

		newGen.setGens(
			utils.array(newGen.needGens).map(() =>
				random.getItem(individual.gens.slice(0, position))
			)
		);

		//console.log(individual.toString());

		random.getItem(individual.gens.slice(position + 1)).randomSetGen(newGen);

		//console.log(individual.toString());

		individual.updateNumbers();

		//console.log(this.toString(), individual.toString(), position, individual.gens.slice(position + 1).length);
		console.log('add', this.toString(), individual.toString());

		return individual;
	}

	mutationModifyGen() {
		const individual = this.clone();
		const newGen = new (random.getItem(genome))();
		const position = random.int(individual.entersGens.length, individual.gens.length - 1);
		const gen = individual.gens[position];

		newGen.setGens(
			utils.array(newGen.needGens).map(() =>
				random.getItem(individual.gens.slice(0, position))
			)
		);

		individual.gens.splice(position, 1, newGen);
		gen.meUsed && gen.meUsed.replaceGen(gen, newGen);
		individual.updateNumbers();

		console.log('modify', this.toString(), individual.toString());

		return individual;
	}

	mutationRemoveGen() {
		const individual = this.clone();
		const removedGen = individual.gens.splice(
			random.int(individual.entersGens.length, individual.gens.length - 1),
			1
		)[0];

		removedGen.restoreLink();
		individual.updateNumbers();

		console.log('remove', this.toString(), individual.toString());

		return individual;
	}

	updateNumbers() {
		this.gens.forEach((g, i) => {
			g.number = i;
		});
	}

	clone() {
		const clone = new Individual();

		this.gens.forEach((g) => {

				//console.log(this.toString());
				//console.log(clone.toString());
			//try {
				const gClone = g.clone(clone.gens);
			//} catch(e) {
			//	console.log(e);
			//	process.exit(0);
			//}
			clone.gens.push(gClone);
		});

		return clone;
	}

	toString() {
		return '' +
			`Individual #${this.id}\n` +
			`Gens ${this.gens.length}\n` +
			this.gens.map((g) => g.toString()).join('\n') + '\n'
		;
	}
};
