
const stepsCount = 10166;
const getPercent = function (count) {
	return (Math.round((count * 100 / stepsCount) * 100) / 100) + '%'
};

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
			console.log(`${history.length} STEP`, history[0][0].rating, getPercent(history[0][0].rating), `${tick - timeTick} ms`);
			timeTick = tick;
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
				console.log(history[0][0].rating);
				console.log(getPercent(history[0][0].rating));
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
