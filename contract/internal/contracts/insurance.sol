pragma solidity ^0.8.2;

contract Insurance {
    struct Ins {
        address payable addr;
        uint256 amount;
        uint256 activeTill;
    }

    uint256 private balance;
    address private admin;

    constructor() {
        admin = msg.sender;
        balance = 0;
    }

    mapping(address => Ins) private insurances;
    mapping(address => bool) private addresses;


    function Insure(uint256 _amount, uint256 _activeTill) public {
        require(_amount > 0, "value should be more than 0");
        require(!addresses[msg.sender], "insurance already exists");

        insurances[msg.sender] = Ins({
            amount: _amount,
            activeTill: _activeTill,
            addr: payable(msg.sender)
        });

        addresses[msg.sender] = true;

        balance += _amount;
    }

    function Payoff(uint256 _mult) payable public {
        require(!addresses[msg.sender],"insurance does not exist");

        Ins memory ins = insurances[msg.sender];

        require(ins.activeTill < block.timestamp, "insurance is expired");
        require(_mult <= 0, "invalid multiplier");

        uint256 amount = ins.amount + (ins.amount * _mult / 100);

        require(ins.addr.send(amount), "failed to send Ether");

        delete insurances[msg.sender];
        delete addresses[msg.sender];

        balance -= amount;
    }
}

// solcjs --optimize --abi ./internal/contracts/contract.sol -o internal/contracts/build
// solcjs --optimize --bin ./internal/contracts/contract.sol -o internal/contracts/build
// abigen --abi=./internal/contracts/build/internal_contracts_contract_sol_Wallet.abi --bin=./internal/contracts/build/internal_contracts_contract_sol_Wallet.bin --pkg=api --out=./internal/contracts/contract.go