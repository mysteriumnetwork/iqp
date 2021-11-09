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

// RentalTokenMetaData contains all meta data concerning the RentalToken contract.
var RentalTokenMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"burner\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEnterprise\",\"outputs\":[{\"internalType\":\"contractIEnterprise\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"contractIEnterprise\",\"name\":\"enterprise\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractIEnterprise\",\"name\":\"enterprise\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50612254806100206000396000f3fe608060405234801561001057600080fd5b506004361061011d5760003560e01c806301ffc9a71461012257806306fdde031461014a578063077f224a1461015f578063081812fc14610174578063095ea7b31461019f57806318160ddd146101b257806323b872dd146101c45780632f745c59146101d757806342842e0e146101ea5780634cd88b76146101fd5780634f6ccce7146102105780636352211e146102235780636a6278421461023657806370a082311461024957806395d89b411461025c578063a22cb46514610264578063b88d4fde14610277578063c4d66de81461028a578063c87b56dd1461029d578063caa0f92a146102b0578063e298a2fd146102b8578063e985e9c5146102c0578063fcd3533c146102d3575b600080fd5b610135610130366004611ce1565b6102e6565b60405190151581526020015b60405180910390f35b610152610311565b604051610141919061203d565b61017261016d366004611deb565b6103a3565b005b610187610182366004611f22565b6103bb565b6040516001600160a01b039091168152602001610141565b6101726101ad366004611c9a565b610425565b6009545b604051908152602001610141565b6101726101d2366004611bb1565b6104e2565b6101b66101e5366004611c9a565b610531565b6101726101f8366004611bb1565b6105a1565b61017261020b366004611d8b565b6105bc565b6101b661021e366004611f22565b610629565b610187610231366004611f22565b6106a3565b6101b6610244366004611b41565b6106ff565b6101b6610257366004611b41565b610779565b6101526107d9565b610172610272366004611c6d565b6107e8565b610172610285366004611bf1565b610898565b610172610298366004611b41565b6108ee565b6101526102ab366004611f22565b610995565b6101b6610a36565b610187610a85565b6101356102ce366004611b79565b610a94565b6101726102e1366004611f3a565b610ac2565b60006001600160e01b0319821663780e9d6360e01b148061030b575061030b82610c54565b92915050565b60606001805461032090612116565b80601f016020809104026020016040519081016040528092919081815260200182805461034c90612116565b80156103995780601f1061036e57610100808354040283529160200191610399565b820191906000526020600020905b81548152906001019060200180831161037c57829003601f168201915b5050505050905090565b6103ac816108ee565b6103b683836105bc565b505050565b60006103c682610ca4565b604051806040016040528060028152602001610c8d60f21b815250906104085760405162461bcd60e51b81526004016103ff919061203d565b60405180910390fd5b50506000908152600560205260409020546001600160a01b031690565b6000610430826106a3565b9050806001600160a01b0316836001600160a01b0316141560405180604001604052806002815260200161191960f11b815250906104815760405162461bcd60e51b81526004016103ff919061203d565b50336001600160a01b038216148061049e575061049e8133610a94565b60405180604001604052806002815260200161323360f01b815250906104d75760405162461bcd60e51b81526004016103ff919061203d565b506103b68383610cc1565b6104ec3382610d2f565b60405180604001604052806002815260200161191b60f11b815250906105255760405162461bcd60e51b81526004016103ff919061203d565b506103b6838383610dd2565b600061053c83610779565b8210604051806040016040528060028152602001610ccd60f21b815250906105775760405162461bcd60e51b81526004016103ff919061203d565b50506001600160a01b03919091166000908152600760209081526040808320938352929052205490565b6103b683838360405180602001604052806000815250610898565b600180546105c990612116565b6040805180820190915260018152601960f91b60208201529150156106015760405162461bcd60e51b81526004016103ff919061203d565b508151610615906001906020850190611a0e565b5080516103b6906002906020840190611a0e565b600061063460095490565b821060405180604001604052806002815260200161333560f01b8152509061066f5760405162461bcd60e51b81526004016103ff919061203d565b506009828154811061069157634e487b7160e01b600052603260045260246000fd5b90600052602060002001549050919050565b60008181526003602090815260408083205481518083019092526002825261323160f01b928201929092526001600160a01b0390911690816106f85760405162461bcd60e51b81526004016103ff919061203d565b5092915050565b600080546040805180820190915260018152600d60fa1b6020820152906001600160a01b031633146107445760405162461bcd60e51b81526004016103ff919061203d565b50600061074f610a36565b905061075b8382610f26565b600b805490600061076b83612151565b90915550909150505b919050565b604080518082019091526002815261032360f41b60208201526000906001600160a01b0383166107bc5760405162461bcd60e51b81526004016103ff919061203d565b50506001600160a01b031660009081526004602052604090205490565b60606002805461032090612116565b604080518082019091526002815261323560f01b60208201526001600160a01b03831633141561082b5760405162461bcd60e51b81526004016103ff919061203d565b503360008181526006602090815260408083206001600160a01b03871680855290835292819020805460ff191686151590811790915590519081529192917f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a35050565b6108a23383610d2f565b60405180604001604052806002815260200161191b60f11b815250906108db5760405162461bcd60e51b81526004016103ff919061203d565b506108e884848484610f44565b50505050565b6000546040805180820190915260018152601960f91b6020820152906001600160a01b0316156109315760405162461bcd60e51b81526004016103ff919061203d565b50604080518082019091526002815261353960f01b60208201526001600160a01b0382166109725760405162461bcd60e51b81526004016103ff919061203d565b50600080546001600160a01b0319166001600160a01b0392909216919091179055565b60606109a082610ca4565b60405180604001604052806002815260200161333360f01b815250906109d95760405162461bcd60e51b81526004016103ff919061203d565b5060006109e4610f94565b90506000815111610a045760405180602001604052806000815250610a2f565b80610a0e8461105b565b604051602001610a1f929190611fa6565b6040516020818303038152906040525b9392505050565b600b54604051603960f91b60208201526001600160601b03193060601b16602182015260358101919091526000906055016040516020818303038152906040528051906020012060001c905090565b6000546001600160a01b031690565b6001600160a01b03918216600090815260066020908152604080832093909416825291909152205460ff1690565b6000546040805180820190915260018152600d60fa1b6020820152906001600160a01b03163314610b065760405162461bcd60e51b81526004016103ff919061203d565b506000610b11610a85565b6040516318a8a6cf60e01b8152600481018590529091506000906001600160a01b038316906318a8a6cf906024016101006040518083038186803b158015610b5857600080fd5b505afa158015610b6c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b909190611e5f565b60e081015160405163cdc2aebf60e01b815261ffff90911660048201529091506000906001600160a01b0384169063cdc2aebf9060240160206040518083038186803b158015610bdf57600080fd5b505afa158015610bf3573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c179190611b5d565b9050610c44848360c001516001600160701b0316836001600160a01b03166111749092919063ffffffff16565b610c4d856111c6565b5050505050565b60006001600160e01b031982166380ac58cd60e01b1480610c8557506001600160e01b03198216635b5e139f60e01b145b8061030b57506301ffc9a760e01b6001600160e01b031983161461030b565b6000908152600360205260409020546001600160a01b0316151590565b600081815260056020526040902080546001600160a01b0319166001600160a01b0384169081179091558190610cf6826106a3565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b6000610d3a82610ca4565b60405180604001604052806002815260200161064760f31b81525090610d735760405162461bcd60e51b81526004016103ff919061203d565b506000610d7f836106a3565b9050806001600160a01b0316846001600160a01b03161480610dba5750836001600160a01b0316610daf846103bb565b6001600160a01b0316145b80610dca5750610dca8185610a94565b949350505050565b826001600160a01b0316610de5826106a3565b6001600160a01b03161460405180604001604052806002815260200161333160f01b81525090610e285760405162461bcd60e51b81526004016103ff919061203d565b50604080518082019091526002815261199960f11b60208201526001600160a01b038316610e695760405162461bcd60e51b81526004016103ff919061203d565b50610e7583838361125b565b610e80600082610cc1565b6001600160a01b0383166000908152600460205260408120805460019290610ea99084906120d3565b90915550506001600160a01b0382166000908152600460205260408120805460019290610ed79084906120a7565b909155505060008181526003602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716916000805160206121ff83398151915291a4505050565b610f408282604051806020016040528060008152506112df565b5050565b610f4f848484610dd2565b610f5b8484848461132f565b60405180604001604052806002815260200161323760f01b81525090610c4d5760405162461bcd60e51b81526004016103ff919061203d565b60606000610fa0610a85565b6001600160a01b0316630cac36b26040518163ffffffff1660e01b815260040160006040518083038186803b158015610fd857600080fd5b505afa158015610fec573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526110149190810190611d19565b905060008151116110345760405180602001604052806000815250611055565b806040516020016110459190611fd5565b6040516020818303038152906040525b91505090565b60608161107f5750506040805180820190915260018152600360fc1b602082015290565b8160005b81156110a9578061109381612151565b91506110a29050600a836120bf565b9150611083565b6000816001600160401b038111156110d157634e487b7160e01b600052604160045260246000fd5b6040519080825280601f01601f1916602001820160405280156110fb576020820181803683370190505b5090505b8415610dca576111106001836120d3565b915061111d600a8661216c565b6111289060306120a7565b60f81b81838151811061114b57634e487b7160e01b600052603260045260246000fd5b60200101906001600160f81b031916908160001a90535061116d600a866120bf565b94506110ff565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b1790526103b6908490611454565b60006111d1826106a3565b90506111df8160008461125b565b6111ea600083610cc1565b6001600160a01b03811660009081526004602052604081208054600192906112139084906120d3565b909155505060008281526003602052604080822080546001600160a01b0319169055518391906001600160a01b038416906000805160206121ff833981519152908390a45050565b611266838383611526565b61126e610a85565b60405163960970c760e01b81526001600160a01b038581166004830152848116602483015260448201849052919091169063960970c790606401600060405180830381600087803b1580156112c257600080fd5b505af11580156112d6573d6000803e3d6000fd5b50505050505050565b6112e983836115de565b6112f6600084848461132f565b60405180604001604052806002815260200161323760f01b815250906108e85760405162461bcd60e51b81526004016103ff919061203d565b60006001600160a01b0384163b1561144957604051630a85bd0160e11b81526001600160a01b0385169063150b7a0290611373903390899088908890600401612000565b602060405180830381600087803b15801561138d57600080fd5b505af19250505080156113bd575060408051601f3d908101601f191682019092526113ba91810190611cfd565b60015b61142f573d8080156113eb576040519150601f19603f3d011682016040523d82523d6000602084013e6113f0565b606091505b508051611427576040805180820182526002815261323760f01b6020820152905162461bcd60e51b81526103ff919060040161203d565b805181602001fd5b6001600160e01b031916630a85bd0160e11b149050610dca565b506001949350505050565b60006114a9826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166116e49092919063ffffffff16565b8051909150156103b657808060200190518101906114c79190611cc5565b6103b65760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b60648201526084016103ff565b6001600160a01b0383166115815761157c81600980546000838152600a60205260408120829055600182018355919091527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af0155565b6115a4565b816001600160a01b0316836001600160a01b0316146115a4576115a483826116f3565b6001600160a01b0382166115bb576103b681611790565b826001600160a01b0316826001600160a01b0316146103b6576103b68282611869565b604080518082019091526002815261323960f01b60208201526001600160a01b03831661161e5760405162461bcd60e51b81526004016103ff919061203d565b5061162881610ca4565b1560405180604001604052806002815260200161033360f41b815250906116625760405162461bcd60e51b81526004016103ff919061203d565b5061166f6000838361125b565b6001600160a01b03821660009081526004602052604081208054600192906116989084906120a7565b909155505060008181526003602052604080822080546001600160a01b0319166001600160a01b03861690811790915590518392906000805160206121ff833981519152908290a45050565b6060610dca84846000856118ad565b6000600161170084610779565b61170a91906120d3565b60008381526008602052604090205490915080821461175d576001600160a01b03841660009081526007602090815260408083208584528252808320548484528184208190558352600890915290208190555b5060009182526008602090815260408084208490556001600160a01b039094168352600781528383209183525290812055565b6009546000906117a2906001906120d3565b6000838152600a6020526040812054600980549394509092849081106117d857634e487b7160e01b600052603260045260246000fd5b90600052602060002001549050806009838154811061180757634e487b7160e01b600052603260045260246000fd5b6000918252602080832090910192909255828152600a9091526040808220849055858252812055600980548061184d57634e487b7160e01b600052603160045260246000fd5b6001900381819060005260206000200160009055905550505050565b600061187483610779565b6001600160a01b039093166000908152600760209081526040808320868452825280832085905593825260089052919091209190915550565b60608247101561190e5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b60648201526084016103ff565b843b61195c5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016103ff565b600080866001600160a01b031685876040516119789190611f8a565b60006040518083038185875af1925050503d80600081146119b5576040519150601f19603f3d011682016040523d82523d6000602084013e6119ba565b606091505b50915091506119ca8282866119d5565b979650505050505050565b606083156119e4575081610a2f565b8251156119f45782518084602001fd5b8160405162461bcd60e51b81526004016103ff919061203d565b828054611a1a90612116565b90600052602060002090601f016020900481019282611a3c5760008555611a82565b82601f10611a5557805160ff1916838001178555611a82565b82800160010185558215611a82579182015b82811115611a82578251825591602001919060010190611a67565b50611a8e929150611a92565b5090565b5b80821115611a8e5760008155600101611a93565b6000611aba611ab584612080565b612050565b9050828152838383011115611ace57600080fd5b828260208301376000602084830101529392505050565b600082601f830112611af5578081fd5b610a2f83833560208501611aa7565b80516001600160701b038116811461077457600080fd5b805161ffff8116811461077457600080fd5b805163ffffffff8116811461077457600080fd5b600060208284031215611b52578081fd5b8135610a2f816121c2565b600060208284031215611b6e578081fd5b8151610a2f816121c2565b60008060408385031215611b8b578081fd5b8235611b96816121c2565b91506020830135611ba6816121c2565b809150509250929050565b600080600060608486031215611bc5578081fd5b8335611bd0816121c2565b92506020840135611be0816121c2565b929592945050506040919091013590565b60008060008060808587031215611c06578081fd5b8435611c11816121c2565b93506020850135611c21816121c2565b92506040850135915060608501356001600160401b03811115611c42578182fd5b8501601f81018713611c52578182fd5b611c6187823560208401611aa7565b91505092959194509250565b60008060408385031215611c7f578182fd5b8235611c8a816121c2565b91506020830135611ba6816121da565b60008060408385031215611cac578081fd5b8235611cb7816121c2565b946020939093013593505050565b600060208284031215611cd6578081fd5b8151610a2f816121da565b600060208284031215611cf2578081fd5b8135610a2f816121e8565b600060208284031215611d0e578081fd5b8151610a2f816121e8565b600060208284031215611d2a578081fd5b81516001600160401b03811115611d3f578182fd5b8201601f81018413611d4f578182fd5b8051611d5d611ab582612080565b818152856020838501011115611d71578384fd5b611d828260208301602086016120ea565b95945050505050565b60008060408385031215611d9d578182fd5b82356001600160401b0380821115611db3578384fd5b611dbf86838701611ae5565b93506020850135915080821115611dd4578283fd5b50611de185828601611ae5565b9150509250929050565b600080600060608486031215611dff578081fd5b83356001600160401b0380821115611e15578283fd5b611e2187838801611ae5565b94506020860135915080821115611e36578283fd5b50611e4386828701611ae5565b9250506040840135611e54816121c2565b809150509250925092565b6000610100808385031215611e72578182fd5b604051908101906001600160401b0382118183101715611e9457611e946121ac565b81604052611ea184611b04565b8152611eaf60208501611b1b565b6020820152611ec060408501611b2d565b6040820152611ed160608501611b2d565b6060820152611ee260808501611b2d565b6080820152611ef360a08501611b2d565b60a0820152611f0460c08501611b04565b60c0820152611f1560e08501611b1b565b60e0820152949350505050565b600060208284031215611f33578081fd5b5035919050565b60008060408385031215611f4c578182fd5b823591506020830135611ba6816121c2565b60008151808452611f768160208601602086016120ea565b601f01601f19169290920160200192915050565b60008251611f9c8184602087016120ea565b9190910192915050565b60008351611fb88184602088016120ea565b835190830190611fcc8183602088016120ea565b01949350505050565b60008251611fe78184602087016120ea565b6672656e74616c2f60c81b920191825250600701919050565b6001600160a01b038581168252841660208201526040810183905260806060820181905260009061203390830184611f5e565b9695505050505050565b602081526000610a2f6020830184611f5e565b604051601f8201601f191681016001600160401b0381118282101715612078576120786121ac565b604052919050565b60006001600160401b03821115612099576120996121ac565b50601f01601f191660200190565b600082198211156120ba576120ba612180565b500190565b6000826120ce576120ce612196565b500490565b6000828210156120e5576120e5612180565b500390565b60005b838110156121055781810151838201526020016120ed565b838111156108e85750506000910152565b600181811c9082168061212a57607f821691505b6020821081141561214b57634e487b7160e01b600052602260045260246000fd5b50919050565b600060001982141561216557612165612180565b5060010190565b60008261217b5761217b612196565b500690565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b03811681146121d757600080fd5b50565b80151581146121d757600080fd5b6001600160e01b0319811681146121d757600080fdfeddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa26469706673582212206098950c638aefba0b2b27e0a08a62e98373eacc58a5fc5b8fa40758a34a08bf64736f6c63430008040033",
}

// RentalTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use RentalTokenMetaData.ABI instead.
var RentalTokenABI = RentalTokenMetaData.ABI

// RentalTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use RentalTokenMetaData.Bin instead.
var RentalTokenBin = RentalTokenMetaData.Bin

// DeployRentalToken deploys a new Ethereum contract, binding an instance of RentalToken to it.
func DeployRentalToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *RentalToken, error) {
	parsed, err := RentalTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(RentalTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &RentalToken{RentalTokenCaller: RentalTokenCaller{contract: contract}, RentalTokenTransactor: RentalTokenTransactor{contract: contract}, RentalTokenFilterer: RentalTokenFilterer{contract: contract}}, nil
}

// RentalToken is an auto generated Go binding around an Ethereum contract.
type RentalToken struct {
	RentalTokenCaller     // Read-only binding to the contract
	RentalTokenTransactor // Write-only binding to the contract
	RentalTokenFilterer   // Log filterer for contract events
}

// RentalTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type RentalTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RentalTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type RentalTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RentalTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type RentalTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// RentalTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type RentalTokenSession struct {
	Contract     *RentalToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// RentalTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type RentalTokenCallerSession struct {
	Contract *RentalTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// RentalTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type RentalTokenTransactorSession struct {
	Contract     *RentalTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// RentalTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type RentalTokenRaw struct {
	Contract *RentalToken // Generic contract binding to access the raw methods on
}

// RentalTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type RentalTokenCallerRaw struct {
	Contract *RentalTokenCaller // Generic read-only contract binding to access the raw methods on
}

// RentalTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type RentalTokenTransactorRaw struct {
	Contract *RentalTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewRentalToken creates a new instance of RentalToken, bound to a specific deployed contract.
func NewRentalToken(address common.Address, backend bind.ContractBackend) (*RentalToken, error) {
	contract, err := bindRentalToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &RentalToken{RentalTokenCaller: RentalTokenCaller{contract: contract}, RentalTokenTransactor: RentalTokenTransactor{contract: contract}, RentalTokenFilterer: RentalTokenFilterer{contract: contract}}, nil
}

// NewRentalTokenCaller creates a new read-only instance of RentalToken, bound to a specific deployed contract.
func NewRentalTokenCaller(address common.Address, caller bind.ContractCaller) (*RentalTokenCaller, error) {
	contract, err := bindRentalToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &RentalTokenCaller{contract: contract}, nil
}

// NewRentalTokenTransactor creates a new write-only instance of RentalToken, bound to a specific deployed contract.
func NewRentalTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*RentalTokenTransactor, error) {
	contract, err := bindRentalToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &RentalTokenTransactor{contract: contract}, nil
}

// NewRentalTokenFilterer creates a new log filterer instance of RentalToken, bound to a specific deployed contract.
func NewRentalTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*RentalTokenFilterer, error) {
	contract, err := bindRentalToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &RentalTokenFilterer{contract: contract}, nil
}

// bindRentalToken binds a generic wrapper to an already deployed contract.
func bindRentalToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(RentalTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RentalToken *RentalTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RentalToken.Contract.RentalTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RentalToken *RentalTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RentalToken.Contract.RentalTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RentalToken *RentalTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RentalToken.Contract.RentalTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_RentalToken *RentalTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _RentalToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_RentalToken *RentalTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _RentalToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_RentalToken *RentalTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _RentalToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_RentalToken *RentalTokenCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _RentalToken.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_RentalToken *RentalTokenSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _RentalToken.Contract.BalanceOf(&_RentalToken.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_RentalToken *RentalTokenCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _RentalToken.Contract.BalanceOf(&_RentalToken.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_RentalToken *RentalTokenCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RentalToken.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_RentalToken *RentalTokenSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _RentalToken.Contract.GetApproved(&_RentalToken.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_RentalToken *RentalTokenCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _RentalToken.Contract.GetApproved(&_RentalToken.CallOpts, tokenId)
}

// GetEnterprise is a free data retrieval call binding the contract method 0xe298a2fd.
//
// Solidity: function getEnterprise() view returns(address)
func (_RentalToken *RentalTokenCaller) GetEnterprise(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _RentalToken.contract.Call(opts, &out, "getEnterprise")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEnterprise is a free data retrieval call binding the contract method 0xe298a2fd.
//
// Solidity: function getEnterprise() view returns(address)
func (_RentalToken *RentalTokenSession) GetEnterprise() (common.Address, error) {
	return _RentalToken.Contract.GetEnterprise(&_RentalToken.CallOpts)
}

// GetEnterprise is a free data retrieval call binding the contract method 0xe298a2fd.
//
// Solidity: function getEnterprise() view returns(address)
func (_RentalToken *RentalTokenCallerSession) GetEnterprise() (common.Address, error) {
	return _RentalToken.Contract.GetEnterprise(&_RentalToken.CallOpts)
}

// GetNextTokenId is a free data retrieval call binding the contract method 0xcaa0f92a.
//
// Solidity: function getNextTokenId() view returns(uint256)
func (_RentalToken *RentalTokenCaller) GetNextTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RentalToken.contract.Call(opts, &out, "getNextTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextTokenId is a free data retrieval call binding the contract method 0xcaa0f92a.
//
// Solidity: function getNextTokenId() view returns(uint256)
func (_RentalToken *RentalTokenSession) GetNextTokenId() (*big.Int, error) {
	return _RentalToken.Contract.GetNextTokenId(&_RentalToken.CallOpts)
}

// GetNextTokenId is a free data retrieval call binding the contract method 0xcaa0f92a.
//
// Solidity: function getNextTokenId() view returns(uint256)
func (_RentalToken *RentalTokenCallerSession) GetNextTokenId() (*big.Int, error) {
	return _RentalToken.Contract.GetNextTokenId(&_RentalToken.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_RentalToken *RentalTokenCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _RentalToken.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_RentalToken *RentalTokenSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _RentalToken.Contract.IsApprovedForAll(&_RentalToken.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_RentalToken *RentalTokenCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _RentalToken.Contract.IsApprovedForAll(&_RentalToken.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RentalToken *RentalTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RentalToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RentalToken *RentalTokenSession) Name() (string, error) {
	return _RentalToken.Contract.Name(&_RentalToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_RentalToken *RentalTokenCallerSession) Name() (string, error) {
	return _RentalToken.Contract.Name(&_RentalToken.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_RentalToken *RentalTokenCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _RentalToken.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_RentalToken *RentalTokenSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _RentalToken.Contract.OwnerOf(&_RentalToken.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_RentalToken *RentalTokenCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _RentalToken.Contract.OwnerOf(&_RentalToken.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RentalToken *RentalTokenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _RentalToken.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RentalToken *RentalTokenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RentalToken.Contract.SupportsInterface(&_RentalToken.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_RentalToken *RentalTokenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _RentalToken.Contract.SupportsInterface(&_RentalToken.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RentalToken *RentalTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _RentalToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RentalToken *RentalTokenSession) Symbol() (string, error) {
	return _RentalToken.Contract.Symbol(&_RentalToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_RentalToken *RentalTokenCallerSession) Symbol() (string, error) {
	return _RentalToken.Contract.Symbol(&_RentalToken.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_RentalToken *RentalTokenCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RentalToken.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_RentalToken *RentalTokenSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _RentalToken.Contract.TokenByIndex(&_RentalToken.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_RentalToken *RentalTokenCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _RentalToken.Contract.TokenByIndex(&_RentalToken.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_RentalToken *RentalTokenCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _RentalToken.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_RentalToken *RentalTokenSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _RentalToken.Contract.TokenOfOwnerByIndex(&_RentalToken.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_RentalToken *RentalTokenCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _RentalToken.Contract.TokenOfOwnerByIndex(&_RentalToken.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_RentalToken *RentalTokenCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _RentalToken.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_RentalToken *RentalTokenSession) TokenURI(tokenId *big.Int) (string, error) {
	return _RentalToken.Contract.TokenURI(&_RentalToken.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_RentalToken *RentalTokenCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _RentalToken.Contract.TokenURI(&_RentalToken.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RentalToken *RentalTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _RentalToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RentalToken *RentalTokenSession) TotalSupply() (*big.Int, error) {
	return _RentalToken.Contract.TotalSupply(&_RentalToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_RentalToken *RentalTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _RentalToken.Contract.TotalSupply(&_RentalToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_RentalToken *RentalTokenTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RentalToken.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_RentalToken *RentalTokenSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RentalToken.Contract.Approve(&_RentalToken.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_RentalToken *RentalTokenTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RentalToken.Contract.Approve(&_RentalToken.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0xfcd3533c.
//
// Solidity: function burn(uint256 tokenId, address burner) returns()
func (_RentalToken *RentalTokenTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int, burner common.Address) (*types.Transaction, error) {
	return _RentalToken.contract.Transact(opts, "burn", tokenId, burner)
}

// Burn is a paid mutator transaction binding the contract method 0xfcd3533c.
//
// Solidity: function burn(uint256 tokenId, address burner) returns()
func (_RentalToken *RentalTokenSession) Burn(tokenId *big.Int, burner common.Address) (*types.Transaction, error) {
	return _RentalToken.Contract.Burn(&_RentalToken.TransactOpts, tokenId, burner)
}

// Burn is a paid mutator transaction binding the contract method 0xfcd3533c.
//
// Solidity: function burn(uint256 tokenId, address burner) returns()
func (_RentalToken *RentalTokenTransactorSession) Burn(tokenId *big.Int, burner common.Address) (*types.Transaction, error) {
	return _RentalToken.Contract.Burn(&_RentalToken.TransactOpts, tokenId, burner)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string name, string symbol, address enterprise) returns()
func (_RentalToken *RentalTokenTransactor) Initialize(opts *bind.TransactOpts, name string, symbol string, enterprise common.Address) (*types.Transaction, error) {
	return _RentalToken.contract.Transact(opts, "initialize", name, symbol, enterprise)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string name, string symbol, address enterprise) returns()
func (_RentalToken *RentalTokenSession) Initialize(name string, symbol string, enterprise common.Address) (*types.Transaction, error) {
	return _RentalToken.Contract.Initialize(&_RentalToken.TransactOpts, name, symbol, enterprise)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string name, string symbol, address enterprise) returns()
func (_RentalToken *RentalTokenTransactorSession) Initialize(name string, symbol string, enterprise common.Address) (*types.Transaction, error) {
	return _RentalToken.Contract.Initialize(&_RentalToken.TransactOpts, name, symbol, enterprise)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string name_, string symbol_) returns()
func (_RentalToken *RentalTokenTransactor) Initialize0(opts *bind.TransactOpts, name_ string, symbol_ string) (*types.Transaction, error) {
	return _RentalToken.contract.Transact(opts, "initialize0", name_, symbol_)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string name_, string symbol_) returns()
func (_RentalToken *RentalTokenSession) Initialize0(name_ string, symbol_ string) (*types.Transaction, error) {
	return _RentalToken.Contract.Initialize0(&_RentalToken.TransactOpts, name_, symbol_)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string name_, string symbol_) returns()
func (_RentalToken *RentalTokenTransactorSession) Initialize0(name_ string, symbol_ string) (*types.Transaction, error) {
	return _RentalToken.Contract.Initialize0(&_RentalToken.TransactOpts, name_, symbol_)
}

// Initialize1 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address enterprise) returns()
func (_RentalToken *RentalTokenTransactor) Initialize1(opts *bind.TransactOpts, enterprise common.Address) (*types.Transaction, error) {
	return _RentalToken.contract.Transact(opts, "initialize1", enterprise)
}

// Initialize1 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address enterprise) returns()
func (_RentalToken *RentalTokenSession) Initialize1(enterprise common.Address) (*types.Transaction, error) {
	return _RentalToken.Contract.Initialize1(&_RentalToken.TransactOpts, enterprise)
}

// Initialize1 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address enterprise) returns()
func (_RentalToken *RentalTokenTransactorSession) Initialize1(enterprise common.Address) (*types.Transaction, error) {
	return _RentalToken.Contract.Initialize1(&_RentalToken.TransactOpts, enterprise)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256)
func (_RentalToken *RentalTokenTransactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _RentalToken.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256)
func (_RentalToken *RentalTokenSession) Mint(to common.Address) (*types.Transaction, error) {
	return _RentalToken.Contract.Mint(&_RentalToken.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256)
func (_RentalToken *RentalTokenTransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _RentalToken.Contract.Mint(&_RentalToken.TransactOpts, to)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_RentalToken *RentalTokenTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RentalToken.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_RentalToken *RentalTokenSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RentalToken.Contract.SafeTransferFrom(&_RentalToken.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_RentalToken *RentalTokenTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RentalToken.Contract.SafeTransferFrom(&_RentalToken.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_RentalToken *RentalTokenTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _RentalToken.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_RentalToken *RentalTokenSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _RentalToken.Contract.SafeTransferFrom0(&_RentalToken.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_RentalToken *RentalTokenTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _RentalToken.Contract.SafeTransferFrom0(&_RentalToken.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_RentalToken *RentalTokenTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _RentalToken.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_RentalToken *RentalTokenSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _RentalToken.Contract.SetApprovalForAll(&_RentalToken.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_RentalToken *RentalTokenTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _RentalToken.Contract.SetApprovalForAll(&_RentalToken.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_RentalToken *RentalTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RentalToken.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_RentalToken *RentalTokenSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RentalToken.Contract.TransferFrom(&_RentalToken.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_RentalToken *RentalTokenTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _RentalToken.Contract.TransferFrom(&_RentalToken.TransactOpts, from, to, tokenId)
}

// RentalTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the RentalToken contract.
type RentalTokenApprovalIterator struct {
	Event *RentalTokenApproval // Event containing the contract specifics and raw log

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
func (it *RentalTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RentalTokenApproval)
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
		it.Event = new(RentalTokenApproval)
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
func (it *RentalTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RentalTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RentalTokenApproval represents a Approval event raised by the RentalToken contract.
type RentalTokenApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_RentalToken *RentalTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*RentalTokenApprovalIterator, error) {

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

	logs, sub, err := _RentalToken.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &RentalTokenApprovalIterator{contract: _RentalToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_RentalToken *RentalTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *RentalTokenApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _RentalToken.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RentalTokenApproval)
				if err := _RentalToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_RentalToken *RentalTokenFilterer) ParseApproval(log types.Log) (*RentalTokenApproval, error) {
	event := new(RentalTokenApproval)
	if err := _RentalToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RentalTokenApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the RentalToken contract.
type RentalTokenApprovalForAllIterator struct {
	Event *RentalTokenApprovalForAll // Event containing the contract specifics and raw log

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
func (it *RentalTokenApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RentalTokenApprovalForAll)
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
		it.Event = new(RentalTokenApprovalForAll)
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
func (it *RentalTokenApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RentalTokenApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RentalTokenApprovalForAll represents a ApprovalForAll event raised by the RentalToken contract.
type RentalTokenApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_RentalToken *RentalTokenFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*RentalTokenApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _RentalToken.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &RentalTokenApprovalForAllIterator{contract: _RentalToken.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_RentalToken *RentalTokenFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *RentalTokenApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _RentalToken.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RentalTokenApprovalForAll)
				if err := _RentalToken.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_RentalToken *RentalTokenFilterer) ParseApprovalForAll(log types.Log) (*RentalTokenApprovalForAll, error) {
	event := new(RentalTokenApprovalForAll)
	if err := _RentalToken.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// RentalTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the RentalToken contract.
type RentalTokenTransferIterator struct {
	Event *RentalTokenTransfer // Event containing the contract specifics and raw log

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
func (it *RentalTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(RentalTokenTransfer)
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
		it.Event = new(RentalTokenTransfer)
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
func (it *RentalTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *RentalTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// RentalTokenTransfer represents a Transfer event raised by the RentalToken contract.
type RentalTokenTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_RentalToken *RentalTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*RentalTokenTransferIterator, error) {

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

	logs, sub, err := _RentalToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &RentalTokenTransferIterator{contract: _RentalToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_RentalToken *RentalTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *RentalTokenTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _RentalToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(RentalTokenTransfer)
				if err := _RentalToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_RentalToken *RentalTokenFilterer) ParseTransfer(log types.Log) (*RentalTokenTransfer, error) {
	event := new(RentalTokenTransfer)
	if err := _RentalToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
