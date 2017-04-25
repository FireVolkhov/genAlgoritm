const Table = require('./table.js');

module.exports = {
	config: {
		count: 200,
		enterGens: 7,
		survivalPercent: 0.25,
		startGensCount: [1, 100]
	},
	test: (run) => {
		const iterator = Table.createIterator(8);

		let result = 0;

		while(!iterator.isFinished) {
			const args = iterator.next();
			const resultCheck = args.splice(args.length - 1, 1)[0];
			const resultIteration = run.apply(null, args);

			const marketBay = 0 < resultCheck;
			const marketWait = resultCheck === 0;
			const marketSell = resultCheck < 0;

			const monkeyBay = 0 < resultIteration;
			const monkeyWait = resultIteration === 0;
			const monkeySell = resultIteration < 0;

			if ((monkeyBay && marketBay) ||
				(monkeySell && marketSell)) {
				result++;
			} else if (
				(monkeyBay && (marketSell || marketWait)) ||
				(monkeySell && (marketBay || marketWait))) {
				result--;
			}

			//if (resultCheck < 0 && resultIteration < 0) {
			//	// console.log('good');
			//	result++;
			//} else if (resultCheck === 0 && resultIteration === 0) {
			//	// Do nothing
			//} else if (resultCheck > 0 && resultIteration > 0) {
			//	result++;
			//} else if (resultCheck > 0 && resultIteration) {
			//	result--;
			//}
		}

		return result;
	}
};
