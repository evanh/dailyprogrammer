// npm install async
var async = require('async');

var https = require('https');
var querystring = require('querystring');

var hostname = 'challenge.opencare.com';

// Helper function for reading information from a GET request
var request = function(options, callback) {
  options['hostname'] = hostname;
  var req = https.request(options, function(res) {
    res.setEncoding('utf8');
    var body = []
    res.on('data', function (chunk) {
      body.push(chunk);
    });
    res.on('end', function() {
      ret = JSON.parse(body);
      callback(ret)
    });
  });
  req.on('error', function(e) {
    console.log('ERROR: ' + e.message);
  });
  req.end()
}

// Check to see if the nucleotide is at the specified location
var get_nucleotide = function(callback, location, project_info, nucleotide) {
  var query_params = {
    location: location,
    nucleotide: nucleotide,
    uniq: project_info.uniq,
  };
  request({path: "/location?" + querystring.stringify(query_params)}, function(ret){
    if (ret.expressed == true) {
      return callback(null, nucleotide);
    } else {
      return callback();
    }
  })
}

// Need this because JS evaluates variables in a different order than Python.
// I feel like there's a cleaner way to do this but I wasn't able to figure out
// the right design pattern.
var gn_builder = function(location, project_info, nucleotide) {
  return function(callback) {
    get_nucleotide(callback, location, project_info, nucleotide);
  };
}

var NUCLEOTIDES = ['A', 'C', 'G', 'T'];

// Find all the nucleotides at a specific location
var check_location = function(return_result, location, project_info) {
  var func_list = [];
  for (var i = 0; i < NUCLEOTIDES.length; i++) {
    func_list.push(gn_builder(location, project_info, NUCLEOTIDES[i]));
  };
  async.parallel(func_list, function(err, results) {
    var output = [];
    for (var i = 0; i < results.length; i++) {
      if (results[i] != undefined) {
        output.push(results[i]);
      }
    }
    return return_result(null, output);
  });
}

// Ditto the gn_builder
var cl_builder = function(i, project_info) {
  return function(callback) {
    check_location(callback, i, project_info);
  };
}

// Fetch all the possible nucleotides at each location, and store it
// in a 'tree', or an array of arrays
var build_tree = function(project_info) {
  var func_list = [];
  for (var i = 0; i < project_info.resultLength; i++) {
    func_list.push(cl_builder(i, project_info));
  };
  async.parallel(func_list, function(err, results) {
    get_matching_genomes(results, project_info);
  })
}

// Recursive function on the tree. Does a depth first search through all the possible
// genomes, and adds them to the matching genomes if the resulting string contains
// the partial we're interested in.
var get_genome = function(string_so_far, tree, project_info, matching_genomes) {
  if (tree.length == 1) {
    for (var i = 0; i < tree[0].length; i++) {
      var genome = string_so_far + tree[0][i];
      if (genome.indexOf(project_info.partialSequence) != -1) {
        matching_genomes.push(genome);
      };
    };
    return
  }

  for (var i = 0; i < tree[0].length; i++) {
    var x = get_genome(string_so_far + tree[0][i], tree.slice(1), project_info, matching_genomes);
  }
}

// Get all the matching genomes, and check them against the test endpoint.
var get_matching_genomes = function(genome_tree, project_info) {
  var matching_genomes = [];
  get_genome("", genome_tree, project_info, matching_genomes);
  console.log(project_info.uniq);
  console.log(matching_genomes);

  var body = JSON.stringify({
    sequences: matching_genomes,
    uniq: project_info.uniq,
  });

  var check_options = {
    hostname: hostname,
    path: '/check',
    method: 'POST',
    headers: {
      'Content-Type': 'application/json',
      'Content-Length': body.length,
    }
  };
  var req = https.request(check_options, function(res) {
    res.setEncoding('utf8');
    var body = "";
    res.on('data', function (chunk) {
          body += chunk;
      });
    res.on('end', function () {
      if (body != "200") {
        console.log("SO MANY ERRORS");
      }
    });
  });
  req.on('error', function(e) {
    console.log('ERROR: ' + e.message);
  });
  req.write(body);
  req.end()
}

// Start the callback chain by fetching the project data
request({path: '/start'}, function(ret) {
  var project_info = ret;
  build_tree(project_info);
});