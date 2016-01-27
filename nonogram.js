// var input = '    *\n' +
// '   **\n' +
// '  * *\n' +
// ' *  *\n' +
// '*****';

var input = '    ** *  \n' +
'   *****  \n' +
'  ******  \n' +
' ******** \n' +
'**********\n' +
' *      * \n' +
' * ** * * \n' +
' * ** * * \n' +
' * **   * \n' +
' ******** ';

// var input = '     ***       \n' +
// '  **** **      \n' +
// ' ****** ****** \n' +
// ' * **** **    *\n' +
// ' ****** ***  **\n' +
// ' ****** *******\n' +
// '****** ********\n' +
// ' *   **********\n' +
// ' *   **********\n' +
// ' *   **********\n' +
// ' * * ****  ****\n' +
// ' *** ****  ****\n' +
// '     ****  ****\n' +
// '     ****  ****\n' +
// '     ****  ****';

var grid = [];
var inputrows = input.split('\n');
for (var i = 0; i < inputrows.length; i++) {
    grid.push(inputrows[i].split(''));
}

// Build row description
var rows = [];
for (var r = 0; r < grid.length; r++) {
    var count = 0;
    var counts = [];
    for (var c = 0; c < grid[r].length; c++) {
        if (grid[r][c] === '*') {
            count++;
        } else if (grid[r][c] === ' ' && count > 0) {
            counts.push(count);
            count = 0;
        }
    }
    if (count > 0) {
        counts.push(count);
    }
    rows.push(counts);
}

// Build column description
var columns = [];
for (var c = 0; c < grid[0].length; c++) {
    var count = 0;
    var counts = [];
    for (var r = 0; r < grid.length; r++) {
        if (grid[r][c] === '*') {
            count++;
        } else if (grid[r][c] === ' ' && count > 0) {
            counts.push(count);
            count = 0;
        }
    }
    if (count > 0) {
        counts.push(count);
    }
    columns.push(counts);
}

// We need these to properly pad the column strings later
var maxrow = 0;
for (var i = 0; i < rows.length; i++) {
    if (rows[i].length > maxrow) {
        maxrow = rows[i].length;
    }
}

// Print the values so they all line up
function pprintrows(values, inner_sep, line_sep) {
    var to_print = [];
    for (var i = 0; i < values.length; i++) {
        var printable = [];
        if (values[i].length < maxrow) {
            for (var j = 0; j < maxrow - values[i].length; j++) {
                printable.push(' ');
            }
        }
        printable = printable.concat(values[i]);
        to_print.push(printable.join(inner_sep) + ' ' + inputrows[i].split('').join(' '));
    }
    return to_print.join(line_sep);
}

// Print the values so they all line up
function pprintcols(values, inner_sep, line_sep) {
    var maxcol = 0;
    for (var i = 0; i < values.length; i++) {
        if (values[i].length > maxcol) {
            maxcol = values[i].length;
        }
    }
    var to_print = [];
    for (var i = 0; i < values.length; i++) {
        var printable = [];
        if (values[i].length < maxcol) {
            for (var j = 0; j < maxcol - values[i].length; j++) {
                printable.push(' ');
            }
        }
        printable = printable.concat(values[i]);
        to_print.push(printable);
    }

    var lines = [];
    for (var l = 0; l < to_print[0].length; l++) {
        var line = [];
        for (var i = 0; i < maxrow; i++) {
            line.push(' ');
        }
        for (var c = 0; c < to_print.length; c++) {
            line.push(to_print[c][l]);
        }
        lines.push(line.join(inner_sep));
    }
    return lines.join(line_sep);
}

console.log(pprintcols(columns, ' ', '\n'));
console.log(pprintrows(rows, ' ', '\n'));
