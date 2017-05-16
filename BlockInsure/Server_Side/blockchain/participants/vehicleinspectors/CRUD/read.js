'use strict';

let tracing = require(__dirname+'/../../../../tools/traces/trace.js');
let participants = require(__dirname+'/../../participants_info.js');

let read = function(req, res)
{
    tracing.create('ENTER', 'GET blockchain/participants/vehicleinspectors', {});

    if(!participants.hasOwnProperty('vehicleinspectors'))
	{
        res.status(404);
        let error = {};
        error.message = 'Unable to retrieve vehicleinspectors';
        error.error = true;
        tracing.create('ERROR', 'GET blockchain/participants/vehicleinspectors', error);
        res.send(error);
    }
    else
	{
        tracing.create('EXIT', 'GET blockchain/participants/vehicleinspectors', {'result':participants.vehicleinspectors});
        res.send({'result':participants.vehicleinspectors});
    }
};
exports.read = read;
