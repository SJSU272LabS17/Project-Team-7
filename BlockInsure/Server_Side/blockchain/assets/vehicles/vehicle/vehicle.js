
var remove = require(__dirname+'/CRUD/delete.js');
exports.delete = remove.delete;

var read = require(__dirname+'/CRUD/read.js');
exports.read = read.read;


var amountFile = require(__dirname+'/amount/amount.js');
var amount = {};
amount.update = amountFile.update;
amount.read = amountFile.read;
exports.amount = amount;

var emailFile = require(__dirname+'/email/email.js');
var email = {};
email.update = emailFile.update;
email.read = emailFile.read;
exports.email = email;

var incidentdateFile = require(__dirname+'/incidentdate/incidentdate.js');
var incidentdate = {};
incidentdate.update = incidentdateFile.update;
incidentdate.read = incidentdateFile.read;
exports.incidentdate = incidentdate;

var licenceplatenumberFile = require(__dirname+'/licenceplatenumber/licenceplatenumber.js');
var licenceplatenumber = {};
licenceplatenumber.update = licenceplatenumberFile.update;
licenceplatenumber.read = licenceplatenumberFile.read;
exports.licenceplatenumber = licenceplatenumber;

var settledFile = require(__dirname+'/settled/settled.js');
var settled = {};
settled.read = settledFile.read;
settled.scrapped = settled;

var VINFile = require(__dirname+'/VIN/vin.js');
var VIN = {};
VIN.update = VINFile.update;
VIN.read = VINFile.read;
exports.VIN = VIN;

var ownerFile = require(__dirname+'/owner/owner.js');
var owner = {};
owner.update = ownerFile.update;
owner.read = ownerFile.read;
exports.owner = owner;
