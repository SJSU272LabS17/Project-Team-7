'use strict';

const Util = require('./util.js');
const hfc = require('hfc');

class Claim {

    constructor(usersToSecurityContext) {
        this.usersToSecurityContext = usersToSecurityContext;
        this.chain = hfc.getChain('myChain'); //TODO: Make this a config param?
    }

    create(userId) {
        let securityContext = this.usersToSecurityContext[userId];
        let chainID = Claim.newclaimID();

        return this.doesclaimIDExist(userId, chainID)
        .then(function() {
            return Util.invokeChaincode(securityContext, 'create_claim', [ chainID ])
            .then(function() {
                return chainID;
            });
        });
    }

    transfer(userId, buyer, functionName, chainID) {
        return this.updateAttribute(userId, functionName , buyer, chainID);
    }

    updateAttribute(userId, functionName, value, chainID) {
        let securityContext = this.usersToSecurityContext[userId];
        return Util.invokeChaincode(securityContext, functionName, [ value, chainID ]);
    }

    doesV5cIDExist(userId, chainID) {
        let securityContext = this.usersToSecurityContext[userId];
        return Util.queryChaincode(securityContext, 'check_unique_v5c', [ chainID ]);
    }

    static newclaimID() {
        let numbers = '1234567890';
        let characters = 'ABCDEFGHIJKLMNOPQRSTUVWXYZ';
        let claimID = '';
        for(let i = 0; i < 7; i++)
            {
            v5cID += numbers.charAt(Math.floor(Math.random() * numbers.length));
        }
        claimID = characters.charAt(Math.floor(Math.random() * characters.length)) + claimID;
        claimID = characters.charAt(Math.floor(Math.random() * characters.length)) + claimID;
        return claimID;
    }
}

module.exports = Claim;
