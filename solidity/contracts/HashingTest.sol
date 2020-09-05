pragma solidity ^0.6.6;

import "@nomiclabs/buidler/console.sol";
import "@openzeppelin/contracts/math/SafeMath.sol";

contract HashingTest {
	using SafeMath for uint256;

	bytes32 public lastCheckpoint;

	function IterativeHash(
		address[] memory _validators,
		uint256[] memory _powers,
		uint256 _valsetNonce,
		bytes32 _peggyId
	) public {
		// bytes32 encoding of the string "checkpoint"
		bytes32 methodName = 0x636865636b706f696e7400000000000000000000000000000000000000000000;

		bytes32 checkpoint = keccak256(abi.encode(_peggyId, methodName, _valsetNonce));

		// Iterative hashing of valset
		{
			for (uint256 i = 0; i < _validators.length; i = i.add(1)) {
				// Check that validator powers are decreasing or equal (this allows the next
				// caller to break out of signature evaluation ASAP to save more gas)
				if (i != 0) {
					require(
						!(_powers[i] > _powers[i - 1]),
						"Validator power must not be higher than previous validator in batch"
					);
				}
				checkpoint = keccak256(abi.encode(checkpoint, _validators[i], _powers[i]));
			}
		}

		lastCheckpoint = checkpoint;
	}

	function ConcatHash(
		address[] memory _validators,
		uint256[] memory _powers,
		uint256 _valsetNonce,
		bytes32 _peggyId
	) public {
		// bytes32 encoding of the string "checkpoint"
		bytes32 methodName = 0x636865636b706f696e7400000000000000000000000000000000000000000000;

		bytes32 idHash = keccak256(abi.encode(_peggyId, methodName, _valsetNonce));

		bytes32 validatorHash = keccak256(abi.encode(_validators));

		bytes32 powersHash = keccak256(abi.encode(_powers));

		bytes32 checkpoint = keccak256(abi.encode(idHash, validatorHash, powersHash));

		lastCheckpoint = checkpoint;
	}

	function ConcatHash2(
		address[] memory _validators,
		uint256[] memory _powers,
		uint256 _valsetNonce,
		bytes32 _peggyId
	) public {
		// bytes32 encoding of the string "checkpoint"
		bytes32 methodName = 0x636865636b706f696e7400000000000000000000000000000000000000000000;

		bytes32 checkpoint = keccak256(
			abi.encode(_peggyId, methodName, _valsetNonce, _validators, _powers)
		);

		lastCheckpoint = checkpoint;
	}
}
