
const stepsCount = 10166;
const getPercent = function (count) {
	return (Math.round((count * 100 / stepsCount) * 100) / 100) + '%'
};

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
			console.log(`${history.length} STEP`, history[0][0].rating, getPercent(history[0][0].rating));
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
				console.log('--- FINISH ---');
				console.log('--- 1 ---');
				console.log(history[0][0].individual);
				console.log(history[0][0].rating);
				console.log(getPercent(history[0][0].rating));
				console.log('--- 2 ---');
				console.log(history[0][1].individual);
				console.log(history[0][1].rating);
				console.log(getPercent(history[0][1].rating));
				console.log('--- 3 ---');
				console.log(history[0][2].individual);
				console.log(history[0][2].rating);
				console.log(getPercent(history[0][2].rating));
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
