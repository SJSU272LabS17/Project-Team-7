'use strict';

let tracing = require(__dirname+'/../../../../tools/traces/trace.js');
let participants = require(__dirname+'/../../participants_info.js');

let read = function(req, res)
{
    tracing.create('ENTER', 'GET blockchain/participants/claiminspectors', {});

    if(!participants.hasOwnProperty('claiminspectors'))
	{
        res.status(404);
        let error = {};
        error.message = 'Unable to retrieve claim inspectors';
        error.error = true;
        tracing.create('ERROR', 'GET blockchain/participants/claiminspectors', error);
        res.send(error);
    }
    else
	{
        tracing.create('EXIT', 'GET blockchain/participants/claiminspectors', {'result':participants.claiminspectors});
        res.send({'result':participants.claiminspectors});
    }
};
exports.read = read;
