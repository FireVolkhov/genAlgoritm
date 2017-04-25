const Table = require('./table.js');

const indexDate = 0;
const indexOpen = 1;
const indexHigh = 2;
const indexLow = 3;
const indexClose = 4;
const indexVolume = 5;
const indexAdjClose = 6;
const indexDif = 6;

module.exports = {
	config: {
		count: 200,
		enterGens: 7,
		survivalPercent: 0.25,
		startGensCount: [1, 100]
	},
	test: (run) => {
		const iterator = Table.createIterator(8);

		let result = 1;

		while(!iterator.isFinished) {
			const slice = iterator.next();
			const futureDay = slice.splice(slice.length - 1, 1)[0];
			const futureDayDif = futureDay[indexClose] - futureDay[indexOpen];
			const futureDayPercent = Math.abs(futureDayDif) / futureDay[indexOpen];

			const args = slice.map((row) => row[indexDif]);
			const monkeyResult = run.apply(null, args);

			const marketBay = 0 < futureDayDif;
			const marketWait = futureDayDif === 0;
			const marketSell = futureDayDif < 0;

			const monkeyBay = 0 < monkeyResult;
			const monkeyWait = monkeyResult === 0;
			const monkeySell = monkeyResult < 0;

			if ((monkeyBay && marketBay) ||
				(monkeySell && marketSell)) {
				result = result * (1 + futureDayPercent);
			} else if (
				(monkeyBay && (marketSell || marketWait)) ||
				(monkeySell && (marketBay || marketWait))) {
				result = result * (1 - futureDayPercent);
			}
		}

		return result;
	}
};
