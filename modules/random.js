module.exports = {
	int: (start, end) => {
		return Math.round(Math.random() * (end - start)) + start;
	},

	getItem(array) {
		return array[this.int(0, array.length - 1)];
	}
};
