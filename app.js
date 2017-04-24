/**
 * Created by User on 09.04.2017.
 */

//var fs = require('fs');
//var util = require('util');
//var log_file = fs.createWriteStream(__dirname + '/debug.log', {flags : 'w', autoClose: true});
//var log_stdout = process.stdout;

//console.log = function(d) { //
//	fs.writeFile("./log.log", "b");
//};

const _ = require('lodash');
const Population = require('./modules/population.js');
const LifeEvents = require('./modules/life_events');
//const genome = require('./genome');
const enterGen = require('./genome/enter.js');

const add = require('./test/add.js');

const history = [];
const defaultConfig = {
	count: NaN,
	enterGens: NaN,

	survivalPercent: 0.5,
	startGensCount: [1, 10]
};

let config = _.extend(_.cloneDeep(defaultConfig), _.cloneDeep(add.config));
let population = Population.create(enterGen, config);

while(true) {
	const result = Population.selection(population, add.test, config);
	population = result[0];
	history.unshift(result[1]);

	config = LifeEvents.fireEvents(history, config);

	population = Population.mutation(population, config);
}


// population = Population.selection(population, add.test, add.population);

// console.log(population[0]);
//population.selection(add.test);

