fs = require('fs');

const file = fs.readFileSync('./test/data/table.csv');
const lines = file.toString().split('\n');
// Убираем заголовки таблицы
lines.splice(0, 1);
// Убираем последнюю пустую строку
lines.splice(lines.length - 1, 1);

const tableRaw = lines.map((line) => line.split(','));
const tableNormalize = tableRaw.reverse();

const indexDate = 0;
const indexOpen = 1;
const indexHigh = 2;
const indexLow = 3;
const indexClose = 4;
const indexVolume = 5;
const indexAdjClose = 6;
const indexDif = 6;

const table = tableNormalize.map((row, indexRow) => {
	row[indexDate] = indexRow;
	row[indexOpen] = parseFloat(row[indexOpen]);
	row[indexHigh] = parseFloat(row[indexHigh]);
	row[indexLow] = parseFloat(row[indexLow]);
	row[indexClose] = parseFloat(row[indexClose]);
	row[indexVolume] = parseInt(row[indexVolume]);
	row[indexAdjClose] = parseFloat(row[indexAdjClose]);

	if (indexRow === 0) {
		row[indexDif] = 0;
	} else {
		row[indexDif] = row[indexClose] - tableNormalize[indexRow - 1][indexClose];
	}

	return row;
});

// const tableResult = table.map((row) => {
// 	if (row[indexDif] < 0) {
// 		return -1;
// 	} else if (row[indexDif] === 0) {
// 		return 0;
// 	} else {
// 		return 1;
// 	}
// });

// Убираем первую строку с 0
table.splice(0, 1);

// const getNextBlock = function(position, size) {
//
// };

module.exports = {
	createIterator: (size) => {
		return {
			position: 0,
			isFinished: false,

			next: function() {
				const position = this.position;
				const end = (position - 1) + size;
				const result = table.slice(position, end);

				this.position++;
				this.isFinished = (table.length < (end + 1));
				return result;
			}
		};
	}
};
