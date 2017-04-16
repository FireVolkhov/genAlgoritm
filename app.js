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

const Population = require('./modules/population.js');
//const genome = require('./genome');
const enterGen = require('./genome/enter.js');

const add = require('./test/add.js');

let population = Population.create(enterGen, add.population);
population = Population.selection(population, add.test, add.population);
population = Population.mutation(population, add.population);
population = Population.selection(population, add.test, add.population);

console.log(population[0]);
//population.selection(add.test);

