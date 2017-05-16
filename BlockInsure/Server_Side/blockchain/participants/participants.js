/*eslint-env node */
var create = require(__dirname+'/CRUD/create.js');
exports.create = create.create;

var read = require(__dirname+'/CRUD/read.js');
exports.read = read.read;

var regulatorsFile = require(__dirname+'/regulators/regulators.js');
var regulators = {};
regulators.read = regulatorsFile.read;
exports.regulators = regulators;

var identityinspectorsFile = require(__dirname+'/identityinspectors/identityinspectors.js');
var identityinspectors = {};
identityinspectors.read = identityinspectorsFile.read;
exports.identityinspectors = identityinspectors;

var testsFile = require(__dirname+'/tests/tests.js');
var tests = {};
tests.read = testsFile.read;
exports.tests = tests;

var vehicleinspectorsFile = require(__dirname+'/vehicleinspectors/vehicleinspectors.js');
var vehicleinspectors = {};
vehicleinspectors.read = vehicleinspectorsFile.read;
exports.vehicleinspectors = vehicleinspectors;

var claiminspectorsFile = require(__dirname+'/claiminspectors/claiminspectors.js');
var claiminspectors = {};
claiminspectors.read = claiminspectorsFile.read;
exports.claiminspectors = claiminspectors;

var settlementofficerFile = require(__dirname+'/settlementofficer/settlementofficer.js');
var settlementofficer = {};
settlementofficer.read = settlementofficerFile.read;
exports.settlementofficer = settlementofficer;