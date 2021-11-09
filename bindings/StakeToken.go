// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// StakeTokenMetaData contains all meta data concerning the StakeToken contract.
var StakeTokenMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEnterprise\",\"outputs\":[{\"internalType\":\"contractIEnterprise\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"contractIEnterprise\",\"name\":\"enterprise\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIEnterprise\",\"name\":\"enterprise\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50611c7a806100206000396000f3fe608060405234801561001057600080fd5b506004361061011d5760003560e01c806301ffc9a71461012257806306fdde031461014a578063077f224a1461015f578063081812fc14610174578063095ea7b31461019f57806318160ddd146101b257806323b872dd146101c45780632f745c59146101d757806342842e0e146101ea57806342966c68146101fd5780634cd88b76146102105780634f6ccce7146102235780636352211e146102365780636a6278421461024957806370a082311461025c57806395d89b411461026f578063a22cb46514610277578063b88d4fde1461028a578063c4d66de81461029d578063c87b56dd146102b0578063caa0f92a146102c3578063e298a2fd146102cb578063e985e9c5146102d3575b600080fd5b61013561013036600461181c565b6102e6565b60405190151581526020015b60405180910390f35b610152610311565b6040516101419190611a74565b61017261016d366004611926565b6103a3565b005b61018761018236600461199a565b6103bb565b6040516001600160a01b039091168152602001610141565b6101726101ad3660046117f1565b610425565b6009545b604051908152602001610141565b6101726101d2366004611704565b6104e2565b6101b66101e53660046117f1565b610531565b6101726101f8366004611704565b6105a1565b61017261020b36600461199a565b6105bc565b61017261021e3660046118c6565b61060d565b6101b661023136600461199a565b61067a565b61018761024436600461199a565b6106f4565b6101b66102573660046116b0565b61074f565b6101b661026a3660046116b0565b6107c8565b610152610828565b6101726102853660046117c0565b610837565b610172610298366004611744565b6108e7565b6101726102ab3660046116b0565b61093d565b6101526102be36600461199a565b6109e4565b6101b6610a85565b610187610ad4565b6101356102e13660046116cc565b610ae3565b60006001600160e01b0319821663780e9d6360e01b148061030b575061030b82610b11565b92915050565b60606000805461032090611b4d565b80601f016020809104026020016040519081016040528092919081815260200182805461034c90611b4d565b80156103995780601f1061036e57610100808354040283529160200191610399565b820191906000526020600020905b81548152906001019060200180831161037c57829003601f168201915b5050505050905090565b6103ac8161093d565b6103b6838361060d565b505050565b60006103c682610b61565b604051806040016040528060028152602001610c8d60f21b815250906104085760405162461bcd60e51b81526004016103ff9190611a74565b60405180910390fd5b50506000908152600460205260409020546001600160a01b031690565b6000610430826106f4565b9050806001600160a01b0316836001600160a01b0316141560405180604001604052806002815260200161191960f11b815250906104815760405162461bcd60e51b81526004016103ff9190611a74565b50336001600160a01b038216148061049e575061049e8133610ae3565b60405180604001604052806002815260200161323360f01b815250906104d75760405162461bcd60e51b81526004016103ff9190611a74565b506103b68383610b7e565b6104ec3382610bec565b60405180604001604052806002815260200161191b60f11b815250906105255760405162461bcd60e51b81526004016103ff9190611a74565b506103b6838383610c8f565b600061053c836107c8565b8210604051806040016040528060028152602001610ccd60f21b815250906105775760405162461bcd60e51b81526004016103ff9190611a74565b50506001600160a01b03919091166000908152600760209081526040808320938352929052205490565b6103b6838383604051806020016040528060008152506108e7565b6006546040805180820190915260018152600d60fa1b6020820152906001600160a01b031633146106005760405162461bcd60e51b81526004016103ff9190611a74565b5061060a81610de3565b50565b6000805461061a90611b4d565b6040805180820190915260018152601960f91b60208201529150156106525760405162461bcd60e51b81526004016103ff9190611a74565b5081516106669060009060208501906115ba565b5080516103b69060019060208401906115ba565b600061068560095490565b821060405180604001604052806002815260200161333560f01b815250906106c05760405162461bcd60e51b81526004016103ff9190611a74565b50600982815481106106e257634e487b7160e01b600052603260045260246000fd5b90600052602060002001549050919050565b600081815260026020818152604080842054815180830190925292815261323160f01b918101919091526001600160a01b0390911690816107485760405162461bcd60e51b81526004016103ff9190611a74565b5092915050565b6006546040805180820190915260018152600d60fa1b60208201526000916001600160a01b031633146107955760405162461bcd60e51b81526004016103ff9190611a74565b5060006107a0610a85565b90506107ac8382610e78565b600b80549060006107bc83611b88565b90915550909392505050565b604080518082019091526002815261032360f41b60208201526000906001600160a01b03831661080b5760405162461bcd60e51b81526004016103ff9190611a74565b50506001600160a01b031660009081526003602052604090205490565b60606001805461032090611b4d565b604080518082019091526002815261323560f01b60208201526001600160a01b03831633141561087a5760405162461bcd60e51b81526004016103ff9190611a74565b503360008181526005602090815260408083206001600160a01b03871680855290835292819020805460ff191686151590811790915590519081529192917f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a35050565b6108f13383610bec565b60405180604001604052806002815260200161191b60f11b8152509061092a5760405162461bcd60e51b81526004016103ff9190611a74565b5061093784848484610e96565b50505050565b6006546040805180820190915260018152601960f91b6020820152906001600160a01b0316156109805760405162461bcd60e51b81526004016103ff9190611a74565b50604080518082019091526002815261353960f01b60208201526001600160a01b0382166109c15760405162461bcd60e51b81526004016103ff9190611a74565b50600680546001600160a01b0319166001600160a01b0392909216919091179055565b60606109ef82610b61565b60405180604001604052806002815260200161333360f01b81525090610a285760405162461bcd60e51b81526004016103ff9190611a74565b506000610a33610eed565b90506000815111610a535760405180602001604052806000815250610a7e565b80610a5d84610fb4565b604051602001610a6e9291906119de565b6040516020818303038152906040525b9392505050565b600b54604051607360f81b60208201526001600160601b03193060601b16602182015260358101919091526000906055016040516020818303038152906040528051906020012060001c905090565b6006546001600160a01b031690565b6001600160a01b03918216600090815260056020908152604080832093909416825291909152205460ff1690565b60006001600160e01b031982166380ac58cd60e01b1480610b4257506001600160e01b03198216635b5e139f60e01b145b8061030b57506301ffc9a760e01b6001600160e01b031983161461030b565b6000908152600260205260409020546001600160a01b0316151590565b600081815260046020526040902080546001600160a01b0319166001600160a01b0384169081179091558190610bb3826106f4565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b6000610bf782610b61565b60405180604001604052806002815260200161064760f31b81525090610c305760405162461bcd60e51b81526004016103ff9190611a74565b506000610c3c836106f4565b9050806001600160a01b0316846001600160a01b03161480610c775750836001600160a01b0316610c6c846103bb565b6001600160a01b0316145b80610c875750610c878185610ae3565b949350505050565b826001600160a01b0316610ca2826106f4565b6001600160a01b03161460405180604001604052806002815260200161333160f01b81525090610ce55760405162461bcd60e51b81526004016103ff9190611a74565b50604080518082019091526002815261199960f11b60208201526001600160a01b038316610d265760405162461bcd60e51b81526004016103ff9190611a74565b50610d328383836110cd565b610d3d600082610b7e565b6001600160a01b0383166000908152600360205260408120805460019290610d66908490611b0a565b90915550506001600160a01b0382166000908152600360205260408120805460019290610d94908490611ade565b909155505060008181526002602052604080822080546001600160a01b0319166001600160a01b038681169182179092559151849391871691600080516020611c2583398151915291a4505050565b6000610dee826106f4565b9050610dfc816000846110cd565b610e07600083610b7e565b6001600160a01b0381166000908152600360205260408120805460019290610e30908490611b0a565b909155505060008281526002602052604080822080546001600160a01b0319169055518391906001600160a01b03841690600080516020611c25833981519152908390a45050565b610e92828260405180602001604052806000815250611185565b5050565b610ea1848484610c8f565b610ead848484846111d5565b60405180604001604052806002815260200161323760f01b81525090610ee65760405162461bcd60e51b81526004016103ff9190611a74565b5050505050565b60606000610ef9610ad4565b6001600160a01b0316630cac36b26040518163ffffffff1660e01b815260040160006040518083038186803b158015610f3157600080fd5b505afa158015610f45573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052610f6d9190810190611854565b90506000815111610f8d5760405180602001604052806000815250610fae565b80604051602001610f9e9190611a0d565b6040516020818303038152906040525b91505090565b606081610fd85750506040805180820190915260018152600360fc1b602082015290565b8160005b81156110025780610fec81611b88565b9150610ffb9050600a83611af6565b9150610fdc565b6000816001600160401b0381111561102a57634e487b7160e01b600052604160045260246000fd5b6040519080825280601f01601f191660200182016040528015611054576020820181803683370190505b5090505b8415610c8757611069600183611b0a565b9150611076600a86611ba3565b611081906030611ade565b60f81b8183815181106110a457634e487b7160e01b600052603260045260246000fd5b60200101906001600160f81b031916908160001a9053506110c6600a86611af6565b9450611058565b6001600160a01b0383166111285761112381600980546000838152600a60205260408120829055600182018355919091527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af0155565b61114b565b816001600160a01b0316836001600160a01b03161461114b5761114b83826112fa565b6001600160a01b038216611162576103b681611397565b826001600160a01b0316826001600160a01b0316146103b6576103b68282611470565b61118f83836114b4565b61119c60008484846111d5565b60405180604001604052806002815260200161323760f01b815250906109375760405162461bcd60e51b81526004016103ff9190611a74565b60006001600160a01b0384163b156112ef57604051630a85bd0160e11b81526001600160a01b0385169063150b7a0290611219903390899088908890600401611a37565b602060405180830381600087803b15801561123357600080fd5b505af1925050508015611263575060408051601f3d908101601f1916820190925261126091810190611838565b60015b6112d5573d808015611291576040519150601f19603f3d011682016040523d82523d6000602084013e611296565b606091505b5080516112cd576040805180820182526002815261323760f01b6020820152905162461bcd60e51b81526103ff9190600401611a74565b805181602001fd5b6001600160e01b031916630a85bd0160e11b149050610c87565b506001949350505050565b60006001611307846107c8565b6113119190611b0a565b600083815260086020526040902054909150808214611364576001600160a01b03841660009081526007602090815260408083208584528252808320548484528184208190558352600890915290208190555b5060009182526008602090815260408084208490556001600160a01b039094168352600781528383209183525290812055565b6009546000906113a990600190611b0a565b6000838152600a6020526040812054600980549394509092849081106113df57634e487b7160e01b600052603260045260246000fd5b90600052602060002001549050806009838154811061140e57634e487b7160e01b600052603260045260246000fd5b6000918252602080832090910192909255828152600a9091526040808220849055858252812055600980548061145457634e487b7160e01b600052603160045260246000fd5b6001900381819060005260206000200160009055905550505050565b600061147b836107c8565b6001600160a01b039093166000908152600760209081526040808320868452825280832085905593825260089052919091209190915550565b604080518082019091526002815261323960f01b60208201526001600160a01b0383166114f45760405162461bcd60e51b81526004016103ff9190611a74565b506114fe81610b61565b1560405180604001604052806002815260200161033360f41b815250906115385760405162461bcd60e51b81526004016103ff9190611a74565b50611545600083836110cd565b6001600160a01b038216600090815260036020526040812080546001929061156e908490611ade565b909155505060008181526002602052604080822080546001600160a01b0319166001600160a01b0386169081179091559051839290600080516020611c25833981519152908290a45050565b8280546115c690611b4d565b90600052602060002090601f0160209004810192826115e8576000855561162e565b82601f1061160157805160ff191683800117855561162e565b8280016001018555821561162e579182015b8281111561162e578251825591602001919060010190611613565b5061163a92915061163e565b5090565b5b8082111561163a576000815560010161163f565b600061166661166184611ab7565b611a87565b905082815283838301111561167a57600080fd5b828260208301376000602084830101529392505050565b600082601f8301126116a1578081fd5b610a7e83833560208501611653565b6000602082840312156116c1578081fd5b8135610a7e81611bf9565b600080604083850312156116de578081fd5b82356116e981611bf9565b915060208301356116f981611bf9565b809150509250929050565b600080600060608486031215611718578081fd5b833561172381611bf9565b9250602084013561173381611bf9565b929592945050506040919091013590565b60008060008060808587031215611759578081fd5b843561176481611bf9565b9350602085013561177481611bf9565b92506040850135915060608501356001600160401b03811115611795578182fd5b8501601f810187136117a5578182fd5b6117b487823560208401611653565b91505092959194509250565b600080604083850312156117d2578182fd5b82356117dd81611bf9565b9150602083013580151581146116f9578182fd5b60008060408385031215611803578182fd5b823561180e81611bf9565b946020939093013593505050565b60006020828403121561182d578081fd5b8135610a7e81611c0e565b600060208284031215611849578081fd5b8151610a7e81611c0e565b600060208284031215611865578081fd5b81516001600160401b0381111561187a578182fd5b8201601f8101841361188a578182fd5b805161189861166182611ab7565b8181528560208385010111156118ac578384fd5b6118bd826020830160208601611b21565b95945050505050565b600080604083850312156118d8578182fd5b82356001600160401b03808211156118ee578384fd5b6118fa86838701611691565b9350602085013591508082111561190f578283fd5b5061191c85828601611691565b9150509250929050565b60008060006060848603121561193a578283fd5b83356001600160401b0380821115611950578485fd5b61195c87838801611691565b94506020860135915080821115611971578384fd5b5061197e86828701611691565b925050604084013561198f81611bf9565b809150509250925092565b6000602082840312156119ab578081fd5b5035919050565b600081518084526119ca816020860160208601611b21565b601f01601f19169290920160200192915050565b600083516119f0818460208801611b21565b835190830190611a04818360208801611b21565b01949350505050565b60008251611a1f818460208701611b21565b657374616b652f60d01b920191825250600601919050565b6001600160a01b0385811682528416602082015260408101839052608060608201819052600090611a6a908301846119b2565b9695505050505050565b602081526000610a7e60208301846119b2565b604051601f8201601f191681016001600160401b0381118282101715611aaf57611aaf611be3565b604052919050565b60006001600160401b03821115611ad057611ad0611be3565b50601f01601f191660200190565b60008219821115611af157611af1611bb7565b500190565b600082611b0557611b05611bcd565b500490565b600082821015611b1c57611b1c611bb7565b500390565b60005b83811015611b3c578181015183820152602001611b24565b838111156109375750506000910152565b600181811c90821680611b6157607f821691505b60208210811415611b8257634e487b7160e01b600052602260045260246000fd5b50919050565b6000600019821415611b9c57611b9c611bb7565b5060010190565b600082611bb257611bb2611bcd565b500690565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b038116811461060a57600080fd5b6001600160e01b03198116811461060a57600080fdfeddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa264697066735822122009a492bb13b933c623b1efd9c4210637f15e48b1e8b9a05bd8f737dc22d6cf6764736f6c63430008040033",
}

// StakeTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use StakeTokenMetaData.ABI instead.
var StakeTokenABI = StakeTokenMetaData.ABI

// StakeTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StakeTokenMetaData.Bin instead.
var StakeTokenBin = StakeTokenMetaData.Bin

// DeployStakeToken deploys a new Ethereum contract, binding an instance of StakeToken to it.
func DeployStakeToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *StakeToken, error) {
	parsed, err := StakeTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StakeTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &StakeToken{StakeTokenCaller: StakeTokenCaller{contract: contract}, StakeTokenTransactor: StakeTokenTransactor{contract: contract}, StakeTokenFilterer: StakeTokenFilterer{contract: contract}}, nil
}

// StakeToken is an auto generated Go binding around an Ethereum contract.
type StakeToken struct {
	StakeTokenCaller     // Read-only binding to the contract
	StakeTokenTransactor // Write-only binding to the contract
	StakeTokenFilterer   // Log filterer for contract events
}

// StakeTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type StakeTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StakeTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StakeTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StakeTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StakeTokenSession struct {
	Contract     *StakeToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StakeTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StakeTokenCallerSession struct {
	Contract *StakeTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// StakeTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StakeTokenTransactorSession struct {
	Contract     *StakeTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// StakeTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type StakeTokenRaw struct {
	Contract *StakeToken // Generic contract binding to access the raw methods on
}

// StakeTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StakeTokenCallerRaw struct {
	Contract *StakeTokenCaller // Generic read-only contract binding to access the raw methods on
}

// StakeTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StakeTokenTransactorRaw struct {
	Contract *StakeTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewStakeToken creates a new instance of StakeToken, bound to a specific deployed contract.
func NewStakeToken(address common.Address, backend bind.ContractBackend) (*StakeToken, error) {
	contract, err := bindStakeToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &StakeToken{StakeTokenCaller: StakeTokenCaller{contract: contract}, StakeTokenTransactor: StakeTokenTransactor{contract: contract}, StakeTokenFilterer: StakeTokenFilterer{contract: contract}}, nil
}

// NewStakeTokenCaller creates a new read-only instance of StakeToken, bound to a specific deployed contract.
func NewStakeTokenCaller(address common.Address, caller bind.ContractCaller) (*StakeTokenCaller, error) {
	contract, err := bindStakeToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StakeTokenCaller{contract: contract}, nil
}

// NewStakeTokenTransactor creates a new write-only instance of StakeToken, bound to a specific deployed contract.
func NewStakeTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*StakeTokenTransactor, error) {
	contract, err := bindStakeToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StakeTokenTransactor{contract: contract}, nil
}

// NewStakeTokenFilterer creates a new log filterer instance of StakeToken, bound to a specific deployed contract.
func NewStakeTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*StakeTokenFilterer, error) {
	contract, err := bindStakeToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StakeTokenFilterer{contract: contract}, nil
}

// bindStakeToken binds a generic wrapper to an already deployed contract.
func bindStakeToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StakeTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakeToken *StakeTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakeToken.Contract.StakeTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakeToken *StakeTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeToken.Contract.StakeTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakeToken *StakeTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakeToken.Contract.StakeTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_StakeToken *StakeTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _StakeToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_StakeToken *StakeTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _StakeToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_StakeToken *StakeTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _StakeToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_StakeToken *StakeTokenCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _StakeToken.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_StakeToken *StakeTokenSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _StakeToken.Contract.BalanceOf(&_StakeToken.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_StakeToken *StakeTokenCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _StakeToken.Contract.BalanceOf(&_StakeToken.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_StakeToken *StakeTokenCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StakeToken.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_StakeToken *StakeTokenSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _StakeToken.Contract.GetApproved(&_StakeToken.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_StakeToken *StakeTokenCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _StakeToken.Contract.GetApproved(&_StakeToken.CallOpts, tokenId)
}

// GetEnterprise is a free data retrieval call binding the contract method 0xe298a2fd.
//
// Solidity: function getEnterprise() view returns(address)
func (_StakeToken *StakeTokenCaller) GetEnterprise(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _StakeToken.contract.Call(opts, &out, "getEnterprise")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEnterprise is a free data retrieval call binding the contract method 0xe298a2fd.
//
// Solidity: function getEnterprise() view returns(address)
func (_StakeToken *StakeTokenSession) GetEnterprise() (common.Address, error) {
	return _StakeToken.Contract.GetEnterprise(&_StakeToken.CallOpts)
}

// GetEnterprise is a free data retrieval call binding the contract method 0xe298a2fd.
//
// Solidity: function getEnterprise() view returns(address)
func (_StakeToken *StakeTokenCallerSession) GetEnterprise() (common.Address, error) {
	return _StakeToken.Contract.GetEnterprise(&_StakeToken.CallOpts)
}

// GetNextTokenId is a free data retrieval call binding the contract method 0xcaa0f92a.
//
// Solidity: function getNextTokenId() view returns(uint256)
func (_StakeToken *StakeTokenCaller) GetNextTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeToken.contract.Call(opts, &out, "getNextTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextTokenId is a free data retrieval call binding the contract method 0xcaa0f92a.
//
// Solidity: function getNextTokenId() view returns(uint256)
func (_StakeToken *StakeTokenSession) GetNextTokenId() (*big.Int, error) {
	return _StakeToken.Contract.GetNextTokenId(&_StakeToken.CallOpts)
}

// GetNextTokenId is a free data retrieval call binding the contract method 0xcaa0f92a.
//
// Solidity: function getNextTokenId() view returns(uint256)
func (_StakeToken *StakeTokenCallerSession) GetNextTokenId() (*big.Int, error) {
	return _StakeToken.Contract.GetNextTokenId(&_StakeToken.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_StakeToken *StakeTokenCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _StakeToken.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_StakeToken *StakeTokenSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _StakeToken.Contract.IsApprovedForAll(&_StakeToken.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_StakeToken *StakeTokenCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _StakeToken.Contract.IsApprovedForAll(&_StakeToken.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StakeToken *StakeTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StakeToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StakeToken *StakeTokenSession) Name() (string, error) {
	return _StakeToken.Contract.Name(&_StakeToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_StakeToken *StakeTokenCallerSession) Name() (string, error) {
	return _StakeToken.Contract.Name(&_StakeToken.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_StakeToken *StakeTokenCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _StakeToken.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_StakeToken *StakeTokenSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _StakeToken.Contract.OwnerOf(&_StakeToken.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_StakeToken *StakeTokenCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _StakeToken.Contract.OwnerOf(&_StakeToken.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakeToken *StakeTokenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _StakeToken.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakeToken *StakeTokenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StakeToken.Contract.SupportsInterface(&_StakeToken.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_StakeToken *StakeTokenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _StakeToken.Contract.SupportsInterface(&_StakeToken.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StakeToken *StakeTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _StakeToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StakeToken *StakeTokenSession) Symbol() (string, error) {
	return _StakeToken.Contract.Symbol(&_StakeToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_StakeToken *StakeTokenCallerSession) Symbol() (string, error) {
	return _StakeToken.Contract.Symbol(&_StakeToken.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_StakeToken *StakeTokenCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakeToken.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_StakeToken *StakeTokenSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _StakeToken.Contract.TokenByIndex(&_StakeToken.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_StakeToken *StakeTokenCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _StakeToken.Contract.TokenByIndex(&_StakeToken.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_StakeToken *StakeTokenCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _StakeToken.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_StakeToken *StakeTokenSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _StakeToken.Contract.TokenOfOwnerByIndex(&_StakeToken.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_StakeToken *StakeTokenCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _StakeToken.Contract.TokenOfOwnerByIndex(&_StakeToken.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_StakeToken *StakeTokenCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _StakeToken.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_StakeToken *StakeTokenSession) TokenURI(tokenId *big.Int) (string, error) {
	return _StakeToken.Contract.TokenURI(&_StakeToken.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_StakeToken *StakeTokenCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _StakeToken.Contract.TokenURI(&_StakeToken.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StakeToken *StakeTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _StakeToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StakeToken *StakeTokenSession) TotalSupply() (*big.Int, error) {
	return _StakeToken.Contract.TotalSupply(&_StakeToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_StakeToken *StakeTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _StakeToken.Contract.TotalSupply(&_StakeToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_StakeToken *StakeTokenTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _StakeToken.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_StakeToken *StakeTokenSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _StakeToken.Contract.Approve(&_StakeToken.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_StakeToken *StakeTokenTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _StakeToken.Contract.Approve(&_StakeToken.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_StakeToken *StakeTokenTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int) (*types.Transaction, error) {
	return _StakeToken.contract.Transact(opts, "burn", tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_StakeToken *StakeTokenSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _StakeToken.Contract.Burn(&_StakeToken.TransactOpts, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0x42966c68.
//
// Solidity: function burn(uint256 tokenId) returns()
func (_StakeToken *StakeTokenTransactorSession) Burn(tokenId *big.Int) (*types.Transaction, error) {
	return _StakeToken.Contract.Burn(&_StakeToken.TransactOpts, tokenId)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string name, string symbol, address enterprise) returns()
func (_StakeToken *StakeTokenTransactor) Initialize(opts *bind.TransactOpts, name string, symbol string, enterprise common.Address) (*types.Transaction, error) {
	return _StakeToken.contract.Transact(opts, "initialize", name, symbol, enterprise)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string name, string symbol, address enterprise) returns()
func (_StakeToken *StakeTokenSession) Initialize(name string, symbol string, enterprise common.Address) (*types.Transaction, error) {
	return _StakeToken.Contract.Initialize(&_StakeToken.TransactOpts, name, symbol, enterprise)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string name, string symbol, address enterprise) returns()
func (_StakeToken *StakeTokenTransactorSession) Initialize(name string, symbol string, enterprise common.Address) (*types.Transaction, error) {
	return _StakeToken.Contract.Initialize(&_StakeToken.TransactOpts, name, symbol, enterprise)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string name_, string symbol_) returns()
func (_StakeToken *StakeTokenTransactor) Initialize0(opts *bind.TransactOpts, name_ string, symbol_ string) (*types.Transaction, error) {
	return _StakeToken.contract.Transact(opts, "initialize0", name_, symbol_)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string name_, string symbol_) returns()
func (_StakeToken *StakeTokenSession) Initialize0(name_ string, symbol_ string) (*types.Transaction, error) {
	return _StakeToken.Contract.Initialize0(&_StakeToken.TransactOpts, name_, symbol_)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string name_, string symbol_) returns()
func (_StakeToken *StakeTokenTransactorSession) Initialize0(name_ string, symbol_ string) (*types.Transaction, error) {
	return _StakeToken.Contract.Initialize0(&_StakeToken.TransactOpts, name_, symbol_)
}

// Initialize1 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address enterprise) returns()
func (_StakeToken *StakeTokenTransactor) Initialize1(opts *bind.TransactOpts, enterprise common.Address) (*types.Transaction, error) {
	return _StakeToken.contract.Transact(opts, "initialize1", enterprise)
}

// Initialize1 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address enterprise) returns()
func (_StakeToken *StakeTokenSession) Initialize1(enterprise common.Address) (*types.Transaction, error) {
	return _StakeToken.Contract.Initialize1(&_StakeToken.TransactOpts, enterprise)
}

// Initialize1 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address enterprise) returns()
func (_StakeToken *StakeTokenTransactorSession) Initialize1(enterprise common.Address) (*types.Transaction, error) {
	return _StakeToken.Contract.Initialize1(&_StakeToken.TransactOpts, enterprise)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256)
func (_StakeToken *StakeTokenTransactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _StakeToken.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256)
func (_StakeToken *StakeTokenSession) Mint(to common.Address) (*types.Transaction, error) {
	return _StakeToken.Contract.Mint(&_StakeToken.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256)
func (_StakeToken *StakeTokenTransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _StakeToken.Contract.Mint(&_StakeToken.TransactOpts, to)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_StakeToken *StakeTokenTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _StakeToken.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_StakeToken *StakeTokenSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _StakeToken.Contract.SafeTransferFrom(&_StakeToken.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_StakeToken *StakeTokenTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _StakeToken.Contract.SafeTransferFrom(&_StakeToken.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_StakeToken *StakeTokenTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _StakeToken.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_StakeToken *StakeTokenSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _StakeToken.Contract.SafeTransferFrom0(&_StakeToken.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_StakeToken *StakeTokenTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _StakeToken.Contract.SafeTransferFrom0(&_StakeToken.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_StakeToken *StakeTokenTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _StakeToken.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_StakeToken *StakeTokenSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _StakeToken.Contract.SetApprovalForAll(&_StakeToken.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_StakeToken *StakeTokenTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _StakeToken.Contract.SetApprovalForAll(&_StakeToken.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_StakeToken *StakeTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _StakeToken.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_StakeToken *StakeTokenSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _StakeToken.Contract.TransferFrom(&_StakeToken.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_StakeToken *StakeTokenTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _StakeToken.Contract.TransferFrom(&_StakeToken.TransactOpts, from, to, tokenId)
}

// StakeTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the StakeToken contract.
type StakeTokenApprovalIterator struct {
	Event *StakeTokenApproval // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakeTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeTokenApproval)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakeTokenApproval)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakeTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeTokenApproval represents a Approval event raised by the StakeToken contract.
type StakeTokenApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_StakeToken *StakeTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*StakeTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _StakeToken.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &StakeTokenApprovalIterator{contract: _StakeToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_StakeToken *StakeTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *StakeTokenApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var approvedRule []interface{}
	for _, approvedItem := range approved {
		approvedRule = append(approvedRule, approvedItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _StakeToken.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeTokenApproval)
				if err := _StakeToken.contract.UnpackLog(event, "Approval", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApproval is a log parse operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_StakeToken *StakeTokenFilterer) ParseApproval(log types.Log) (*StakeTokenApproval, error) {
	event := new(StakeTokenApproval)
	if err := _StakeToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeTokenApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the StakeToken contract.
type StakeTokenApprovalForAllIterator struct {
	Event *StakeTokenApprovalForAll // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakeTokenApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeTokenApprovalForAll)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakeTokenApprovalForAll)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakeTokenApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeTokenApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeTokenApprovalForAll represents a ApprovalForAll event raised by the StakeToken contract.
type StakeTokenApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_StakeToken *StakeTokenFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*StakeTokenApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _StakeToken.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &StakeTokenApprovalForAllIterator{contract: _StakeToken.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_StakeToken *StakeTokenFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *StakeTokenApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _StakeToken.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeTokenApprovalForAll)
				if err := _StakeToken.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseApprovalForAll is a log parse operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_StakeToken *StakeTokenFilterer) ParseApprovalForAll(log types.Log) (*StakeTokenApprovalForAll, error) {
	event := new(StakeTokenApprovalForAll)
	if err := _StakeToken.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// StakeTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the StakeToken contract.
type StakeTokenTransferIterator struct {
	Event *StakeTokenTransfer // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *StakeTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(StakeTokenTransfer)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(StakeTokenTransfer)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *StakeTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *StakeTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// StakeTokenTransfer represents a Transfer event raised by the StakeToken contract.
type StakeTokenTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_StakeToken *StakeTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*StakeTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _StakeToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &StakeTokenTransferIterator{contract: _StakeToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_StakeToken *StakeTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *StakeTokenTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var tokenIdRule []interface{}
	for _, tokenIdItem := range tokenId {
		tokenIdRule = append(tokenIdRule, tokenIdItem)
	}

	logs, sub, err := _StakeToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(StakeTokenTransfer)
				if err := _StakeToken.contract.UnpackLog(event, "Transfer", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// ParseTransfer is a log parse operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_StakeToken *StakeTokenFilterer) ParseTransfer(log types.Log) (*StakeTokenTransfer, error) {
	event := new(StakeTokenTransfer)
	if err := _StakeToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
