const create = function(name, value) {
	return {
		name: name,
		needGens: 0,

		create() {
			return `${this.name}`
		},

		getRunFunction(gen, individual) {
			return function (args) {
				return value;
			}
		}
	};
};

module.exports = [
	create('MINUS_ONE', -1),
	create('NULL', 0),
	create('ONE', 1),
	create('PI', 3.1415),
	create('NEPER', 2.718),
	create('ELER', 0.5772),
	create('GOLD_MEMBER', 1.6180),
	create('ONE', 1),
];
