const _ = require('lodash');

const LifeEvents = {
	fireEvents: (history, inputConfig) => {
		let config = inputConfig;

		LifeEvents.events.forEach((event) => {
			if (event.isTrigger(history)) {
				config = event.action(history, config);
			}
		});

		return config;
	},

	events: [

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

		// Завершения поиска решения
		{
			name: 'finish',
			isTrigger: (history) => {
				return true;
			},
			action: (history, config) => {
				if (50 < history.length) {
					console.log('--- FINISH ---');
					console.log(history[0][0].individual);
					process.exit(0);
				} else {
					return config;
				}
			}
		}
	]
};

module.exports = LifeEvents;
