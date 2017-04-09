module.exports = {
	population: {
		count: 50,
		enterGens: 2
	},
	test: (run) => {
		const results = [
			run(5, 10) === 15,
			run(0, 0) === 0,
			run(1, 0) === 1,
			run(0, 1) === 1,
			run(1, 1) === 2,
			run(1, 2) === 3,
			run(-1, 0) === -1,
			run(0, -1) === -1,
			run(-3, 3) === 0,
			run(-10, 3) === -7
		];

		return results.reduce(function (a, b) {
			return a + b;
		}) / results.length;
	}
};
