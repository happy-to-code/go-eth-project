pragma solidity ^0.4.24;

contract SimpleStorage {
    string message;

    constructor(string src) public {
        message = src;
    }

    function setMessage(string _msg)public{
        message = _msg;
    }

    function getMessage()public view returns(string){
        return message ;
    }
}