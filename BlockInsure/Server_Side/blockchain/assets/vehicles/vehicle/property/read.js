'use strict';

let tracing = require(__dirname+'/../../../../../tools/traces/trace.js');
let Util = require(__dirname+'/../../../../../tools/utils/util');
let map_ID = require(__dirname+'/../../../../../tools/map_ID/map_ID');
let user_id;
let user;
let securityContext;

function read(req,res,next,usersToSecurityContext,property) {
    let claimID = req.params.claimID;
    tracing.create('ENTER', 'GET blockchain/assets/vehicles/vehicle/'+claimID+'/' + property, {});

    if(typeof req.cookies.user !== 'undefined')
    {	
        req.session.user = req.cookies.user;
        req.session.identity = map_ID.user_to_id(req.cookies.user);
    }
    user_id = req.session.identity;
    securityContext = usersToSecurityContext[user_id];
    user = securityContext.getEnrolledMember();

    return Util.queryChaincode(securityContext, 'get_claim_details', [ claimID ]).
    then(function(data) {
        let vehicle = JSON.parse(data.toString());
        let result = {};
        result.message = vehicle[property];
        tracing.create('EXIT', 'GET blockchain/assets/vehicles/vehicle/'+claimID+'/' + property, result);
        res.send(result);
    })
    .catch(function(err) {
        res.status(400);
        let error = {};
        error.message = err;
        error.v5cID = v5cID;
        error.error = true;
        tracing.create('ERROR', 'GET blockchain/assets/vehicles/vehicle/'+claimID+'/' + property, error);
        res.send(error);
    });
}

module.exports = read;
