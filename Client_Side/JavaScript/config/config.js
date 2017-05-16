$(document).ready(function(){
    loadLogo(pgNm);
});



let config = {};

/******* Images *******/

// Variable Setup
config.logo = {};
config.logo.main = {};
config.logo.regulator = {};
config.logo.manufacturer = {};
config.logo.dealership = {};
config.logo.lease_company = {};
config.logo.leasee = {};
config.logo.scrap_merchant = {};

// Logo size
config.logo.width = '8em';
config.logo.height = '3.6em';

//Main Logo
config.logo.main.src = 'Icons/new-logo.png';

// Regulator Logo
config.logo.regulator.src = 'Icons/Regulator/new-logo.png';

// Manufacturer Logo
config.logo.manufacturer.src = 'Icons/Manufacturer/new-logo.png';

// Dealership Logo
config.logo.dealership.src = 'Icons/Dealership/new-logo.png';

// Lease Company Logo
config.logo.lease_company.src = 'Icons/Lease_Company/new-logo.png';

// Leasee Logo
config.logo.leasee.src = 'Icons/Leasee/new-logo.png';

// Scrap Merchant Logo
config.logo.scrap_merchant.src = 'Icons/Scrap_Merchant/new-logo.png';

/******* Participants *******/
//This is where we define the details of the users for each company (name and password)

// Variable Setup
config.participants = {};
config.participants.users = {};
config.participants.users.regulators = [];
config.participants.users.manufacturers = [];
config.participants.users.dealerships = [];
config.participants.users.lease_companies = [];
config.participants.users.leasees = [];
config.participants.users.scrap_merchants = [];
config.participants.regulator = {};
config.participants.manufacturer = {};
config.participants.dealership = {};
config.participants.lease_company = {};
config.participants.leasee = {};
config.participants.scrap_merchant = {};

// Regulators
config.participants.users.regulators[0]= {};
config.participants.users.regulators[0].company = 'DVLA';
config.participants.users.regulators[0].type = 'Regulator';
config.participants.users.regulators[0].user = 'Ronald';
config.participants.users.regulators[0].label = 'Regulator';

// Manufacturers
config.participants.users.manufacturers[0] = {};
config.participants.users.manufacturers[0].company = 'Geico';
config.participants.users.manufacturers[0].type = 'Manufacturer';
config.participants.users.manufacturers[0].user = 'Martin';
config.participants.users.manufacturers[0].label = 'Insurance Company';

config.participants.users.manufacturers[1] = {};
config.participants.users.manufacturers[1].company = 'Mercury';
config.participants.users.manufacturers[1].type = 'Manufacturer';
config.participants.users.manufacturers[1].user = 'Maria';
config.participants.users.manufacturers[1].label = 'Insurance Company';

config.participants.users.manufacturers[2] = {};
config.participants.users.manufacturers[2].company = 'Allstate';
config.participants.users.manufacturers[2].type = 'Manufacturer';
config.participants.users.manufacturers[2].user = 'Mandy';
config.participants.users.manufacturers[2].label = 'Insurance Company';

// Dealerships
config.participants.users.dealerships[0] = {};
config.participants.users.dealerships[0].company = 'AutoVerify';
config.participants.users.dealerships[0].type = 'Dealership';
config.participants.users.dealerships[0].user = 'Deborah';
config.participants.users.dealerships[0].label = 'Identity Verification';

config.participants.users.dealerships[1] = {};
config.participants.users.dealerships[1].company = 'Justify';
config.participants.users.dealerships[1].type = 'Dealership';
config.participants.users.dealerships[1].user = 'Dennis';
config.participants.users.dealerships[1].label = 'Identity Verification';

config.participants.users.dealerships[2] = {};
config.participants.users.dealerships[2].company = 'Autosure';
config.participants.users.dealerships[2].type = 'Dealership';
config.participants.users.dealerships[2].user = 'Delia';
config.participants.users.dealerships[2].label = 'Identity Verification';

// Lease Companies
config.participants.users.lease_companies[0] = {};
config.participants.users.lease_companies[0].company = 'AutoInspect';
config.participants.users.lease_companies[0].type = 'Lease Company';
config.participants.users.lease_companies[0].user = 'Lesley';
config.participants.users.lease_companies[0].label = 'Vehicle Inspection';

config.participants.users.lease_companies[1] = {};
config.participants.users.lease_companies[1].company = 'InspectSure';
config.participants.users.lease_companies[1].type = 'Lease Company';
config.participants.users.lease_companies[1].user = 'Larry';
config.participants.users.lease_companies[1].label = 'Vehicle Inspection';

config.participants.users.lease_companies[2] = {};
config.participants.users.lease_companies[2].company = 'AllstateVerify';
config.participants.users.lease_companies[2].type = 'Lease Company';
config.participants.users.lease_companies[2].user = 'Luke';
config.participants.users.lease_companies[2].label = 'Vehicle Inspection';

// Leasees
config.participants.users.leasees[0] = {};
config.participants.users.leasees[0].company = 'Sudoku';
config.participants.users.leasees[0].type = 'Leasee';
config.participants.users.leasees[0].user = 'Joe';
config.participants.users.leasees[0].label = 'Claim Verification';

config.participants.users.leasees[1] = {};
config.participants.users.leasees[1].company = 'QuickClaim';
config.participants.users.leasees[1].type = 'Leasee';
config.participants.users.leasees[1].user = 'Andrew';
config.participants.users.leasees[1].label = 'Claim Verification';

config.participants.users.leasees[2] = {};
config.participants.users.leasees[2].company = 'HHM';
config.participants.users.leasees[2].type = 'Leasee';
config.participants.users.leasees[2].user = 'Anthony';
config.participants.users.leasees[2].label = 'Claim Verification';

// Scrap Merchants
config.participants.users.scrap_merchants[0] = {};
config.participants.users.scrap_merchants[0].company = 'SamSort';
config.participants.users.scrap_merchants[0].type = 'Scrap Merchant';
config.participants.users.scrap_merchants[0].user = 'Sandy';
config.participants.users.scrap_merchants[0].label = 'Settlement Authority';

config.participants.users.scrap_merchants[1] = {};
config.participants.users.scrap_merchants[1].company = 'SetSurance';
config.participants.users.scrap_merchants[1].type = 'Scrap Merchant';
config.participants.users.scrap_merchants[1].user = 'Scott';
config.participants.users.scrap_merchants[1].label = 'Settlement Authority';

config.participants.users.scrap_merchants[2] = {};
config.participants.users.scrap_merchants[2].company = 'InsuranceIt';
config.participants.users.scrap_merchants[2].type = 'Scrap Merchant';
config.participants.users.scrap_merchants[2].user = 'Sid';
config.participants.users.scrap_merchants[2].label = 'Settlement Authority';


/******* Used Particpants *******/
//This is where we select which participants will be used for the pages

// Regulator
config.participants.regulator = config.participants.users.regulators[0];

// Manufacturer
config.participants.manufacturer = config.participants.users.manufacturers[0];

// Dealership
config.participants.dealership = config.participants.users.dealerships[0];

// Lease Company
config.participants.lease_company = config.participants.users.lease_companies[0];

// Leasee
config.participants.leasee = config.participants.users.leasees[0];

// Scrap Merchant
config.participants.scrap_merchant = config.participants.users.scrap_merchants[0];

function loadLogo(pageType)
{
    $('#logo').attr('src', config.logo[pageType.toLowerCase()].src);
    $('#logo').css('width', config.logo.width);
    $('#logo').css('height', config.logo.height);
}

function loadParticipant(pageType)
{
    $('#username').html(config.participants[pageType].user);
    $('#company').html(config.participants[pageType].company);
}
