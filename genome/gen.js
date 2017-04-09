module.exports = class Gen {
	constructor () {
		this.name = 'No name';
		this.number = -1;
		this.needGens = 0;
		this.meUsed = null;
	}

	clone() {
		return new Gen();
	}
};
