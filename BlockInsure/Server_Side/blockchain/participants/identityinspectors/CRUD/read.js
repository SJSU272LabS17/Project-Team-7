'use strict';

let tracing = require(__dirname+'/../../../../tools/traces/trace.js');
let participants = require(__dirname+'/../../participants_info.js');

let read = function(req, res)
{
    tracing.create('ENTER', 'GET blockchain/participants/identityinspectors', {});

    if(!participants.hasOwnProperty('identityinspectors'))
    {
        res.status(404);
        let error = {};
        error.message = 'Unable to retrieve identityinspectors';
        error.error = true;
        tracing.create('ERROR', 'GET blockchain/participants/identityinspectors', error);
        res.send(error);
    }
    else
    {
        tracing.create('EXIT', 'GET blockchain/participants/identityinspectors', {'result':participants.identityinspectors});
        res.send({'result':participants.identityinspectors});
    }

};
exports.read = read;
