const Table = require('./table.js');

module.exports = {
	config: {
		count: 200,
		enterGens: 7,
		survivalPercent: 0.1,
		startGensCount: [1, 100]
	},
	test: (run) => {
		const iterator = Table.createIterator(8);

		let result = 0;

		while(!iterator.isFinished) {
			const args = iterator.next();
			const resultCheck = args.splice(args.length - 1, 1)[0];
			const resultIteration = run.apply(null, args);
			// console.log(args, resultIteration);
			// console.log(args, '>', resultCheck);

			if (resultCheck < 0 && resultIteration < 0) {
				// console.log('good');
				result++;
			} else if (resultCheck === 0 && resultIteration === 0) {
				// Do nothing
			} else if (resultCheck > 0 && resultIteration > 0) {
				result++;
			} else {
				result--;
			}
		}

		return result;
	}
};
