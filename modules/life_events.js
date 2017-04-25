
const stepsCount = 10166;
const toPercent = function (count) {
	return Math.round(count * 100 * 100) / 100 + '%';
};

const getPercent = function (count) {
	return toPercent(count / stepsCount);
};

const Open = 107.459999;
const Close = 2348.689941;

const timeStart = new Date().getTime();
let timeTick = timeStart;

const events = [

	// Начало поиска решения
	{
		name: 'start',
		isTrigger: (history) => {
			return history.length < 2;
		},
		action: (history, config) => {
			console.log('--- START ---');
			return config;
		}
	},

	{
		name: 'tick',
		isTrigger: (history) => {
			return true;
		},
		action: (history, config) => {
			const tick = new Date().getTime();
			console.log(`${history.length} STEP`, toPercent(history[0][0].rating), `${tick - timeTick} ms`);
			timeTick = tick;
			return config;
		}
	},

	{
		name: 'mutation shock',
		isActive: false,
		isTrigger: (history) => {
			const slice = history.slice(0, 5);
			const ratings = slice.map((step) => step[0].rating);
			const topRatings = ratings.filter((rating) => rating === ratings[0]);

			return 5 < history.length && ((topRatings.length === ratings.length) || this.isActive);
		},
		action: (history, config) => {
			if (this.isActive) {
				this.isActive = false;
				config.count = config.count / 10;
				console.log(`MUTATION SHOCK OFF set count: ${config.count}`);
			} else {
				this.isActive = true;
				config.count = config.count * 10;
				console.log(`MUTATION SHOCK set count: ${config.count}`);
			}

			return config;
		}
	},

	// Завершения поиска решения
	{
		name: 'finish',
		isTrigger: (history) => {
			return true;
		},
		action: (history, config) => {
			if (500 < history.length) {
				const timeAll = new Date().getTime();
				console.log('--- FINISH ---');
				console.log(`Time: ${timeAll - timeStart} ms`);
				console.log(history[0][0].individual);
				console.log(`Result: ${toPercent(history[0][0].rating)}`);
				console.log(`Normal result: ${toPercent((Close - Open) / Open)}`);
				// console.log(getPercent(history[0][0].rating));
				process.exit(0);
			} else {
				return config;
			}
		}
	}
];

module.exports = {
	fireEvents: (history, inputConfig) => {
		let config = inputConfig;

		events.forEach((event) => {
			if (event.isTrigger(history)) {
				config = event.action(history, config);
			}
		});

		return config;
	}
};
