pragma solidity ^0.6.6;
import "@openzeppelin/contracts/token/ERC20/ERC20.sol";
import "@uniswap/v2-periphery/contracts/interfaces/IUniswapV2Router02.sol";
import "@uniswap/v2-core/contracts/interfaces/IUniswapV2Factory.sol";

contract wFUND is ERC20 {

	address public uniswapV2Pair;

	constructor(
		address _tokensReceiver,
		string memory _name,
		string memory _symbol,
		uint8 _decimals,
		uint256 _supply,
		address _router
	) public ERC20(_name, _symbol) {
		_setupDecimals(_decimals);
		_mint(_tokensReceiver, _supply);

		IUniswapV2Router02 _uniswapV2Router = IUniswapV2Router02(_router);

		uniswapV2Pair = IUniswapV2Factory(_uniswapV2Router.factory())
            .createPair(address(this), _uniswapV2Router.WETH());
	}

	
}
