'use strict';

let tracing = require(__dirname+'/../../../../tools/traces/trace.js');
let participants = require(__dirname+'/../../participants_info.js');

let read = function(req, res)
{
    tracing.create('ENTER', 'GET blockchain/participants/settlementofficer', {});

    if(!participants.hasOwnProperty('settlementofficer'))
	{
        res.status(404);
        let error = {};
        error.message = 'Unable to retrieve settlementofficer';
        error.error = true;
        tracing.create('ERROR', 'GET blockchain/participants/settlementofficer', error);
        res.send(error);
    }
    else
	{
        tracing.create('EXIT', 'GET blockchain/participants/settlementofficer', {'result':participants.settlementofficer});
        res.send({'result':participants.settlementofficer});
    }
};
exports.read = read;
