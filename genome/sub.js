const random = require('../modules/random.js');
const Gen = require('./gen.js');

module.exports = class SubGen extends Gen {
	constructor (a, b) {
		super();
		this.name = 'SUB';
		this.needGens = 2;

		this.a = null;
		this.b = null;
	}

	setGens(gens) {
		this.a = gens[0];
		this.b = gens[1];
		this.a.meUsed = this;
		this.b.meUsed = this;
		return this;
	}

	run() {
		return this.a.run() - this.b.run();
	}

	replaceGen(oldGen, newGen) {
		if (this.a === oldGen) {
			this.a = newGen;
		}

		if (this.b === oldGen) {
			this.b = newGen;
		}

		return this;
	}

	clone(genome) {
		console.log(genome.length, this.a.number, this.b.number);
		const clone = new SubGen();
		clone.a = genome[this.a.number];
		clone.b = genome[this.b.number];
		clone.a.meUsed = clone;
		clone.b.meUsed = clone;
		return clone;
	}

	restoreLink() {
		if (this.meUsed) {
			this.meUsed.replaceGen(
				this,
				random.getItem([this.a, this.b])
			);
		}

		return this;
	}

	randomSetGen(gen) {
		if (random.int([0, 1])) {
			this.a = gen;
		} else {
			this.b = gen;
		}

		return this;
	}

	toString() {
		return `#${this.number} - ${this.name} #${this.a.number} #${this.b.number}`;
	}
};
