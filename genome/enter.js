const Gen = require('./gen.js');

module.exports = class EnterGen extends Gen {
	constructor () {
		super();
		this.name = 'ENTER';
		this.needGens = 1;

		this.value = Function.prototype;
	}

	setGens(gens) {
		this.value = gens[0];
		return this;
	}

	run() {
		return this.value();
	}

	clone() {
		const clone = new EnterGen();
		clone.value = this.value;
		return clone;
	}

	toString() {
		return `#${this.number} - ${this.name}`
	}
};
