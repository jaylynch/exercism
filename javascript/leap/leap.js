//
// This is only a SKELETON file for the "Leap" exercise. It's been provided as a
// convenience to get you started writing code faster.
//

var Year = function (input) {
    this.year = input;
};

Year.prototype.isLeap = function () {
    const y = this.year;

    return (
        y % 4 === 0 &&
        (
            y % 100 !== 0 ||
            y % 400 === 0
        )
    );
};

module.exports = Year;
