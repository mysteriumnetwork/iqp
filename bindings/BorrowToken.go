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

// BorrowTokenMetaData contains all meta data concerning the BorrowToken contract.
var BorrowTokenMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"approved\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"ApprovalForAll\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"burner\",\"type\":\"address\"}],\"name\":\"burn\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"getApproved\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEnterprise\",\"outputs\":[{\"internalType\":\"contractEnterprise\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getNextTokenId\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol\",\"type\":\"string\"},{\"internalType\":\"contractEnterprise\",\"name\":\"enterprise\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name_\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"symbol_\",\"type\":\"string\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractEnterprise\",\"name\":\"enterprise\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"}],\"name\":\"isApprovedForAll\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"}],\"name\":\"mint\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"ownerOf\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"},{\"internalType\":\"bytes\",\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"safeTransferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"operator\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"approved\",\"type\":\"bool\"}],\"name\":\"setApprovalForAll\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes4\",\"name\":\"interfaceId\",\"type\":\"bytes4\"}],\"name\":\"supportsInterface\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"owner\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"tokenOfOwnerByIndex\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"tokenURI\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"from\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"to\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"tokenId\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50612253806100206000396000f3fe608060405234801561001057600080fd5b506004361061011d5760003560e01c806301ffc9a71461012257806306fdde031461014a578063077f224a1461015f578063081812fc14610174578063095ea7b31461019f57806318160ddd146101b257806323b872dd146101c45780632f745c59146101d757806342842e0e146101ea5780634cd88b76146101fd5780634f6ccce7146102105780636352211e146102235780636a6278421461023657806370a082311461024957806395d89b411461025c578063a22cb46514610264578063b88d4fde14610277578063c4d66de81461028a578063c87b56dd1461029d578063caa0f92a146102b0578063e298a2fd146102b8578063e985e9c5146102c0578063fcd3533c146102d3575b600080fd5b610135610130366004611ce0565b6102e6565b60405190151581526020015b60405180910390f35b610152610311565b604051610141919061203c565b61017261016d366004611dea565b6103a3565b005b610187610182366004611f21565b6103bb565b6040516001600160a01b039091168152602001610141565b6101726101ad366004611c99565b610425565b6009545b604051908152602001610141565b6101726101d2366004611bb0565b6104e2565b6101b66101e5366004611c99565b610531565b6101726101f8366004611bb0565b6105a1565b61017261020b366004611d8a565b6105bc565b6101b661021e366004611f21565b610629565b610187610231366004611f21565b6106a3565b6101b6610244366004611b40565b6106ff565b6101b6610257366004611b40565b610779565b6101526107d9565b610172610272366004611c6c565b6107e8565b610172610285366004611bf0565b610898565b610172610298366004611b40565b6108ee565b6101526102ab366004611f21565b610995565b6101b6610a36565b610187610a85565b6101356102ce366004611b78565b610a94565b6101726102e1366004611f39565b610ac2565b60006001600160e01b0319821663780e9d6360e01b148061030b575061030b82610c53565b92915050565b60606001805461032090612115565b80601f016020809104026020016040519081016040528092919081815260200182805461034c90612115565b80156103995780601f1061036e57610100808354040283529160200191610399565b820191906000526020600020905b81548152906001019060200180831161037c57829003601f168201915b5050505050905090565b6103ac816108ee565b6103b683836105bc565b505050565b60006103c682610ca3565b604051806040016040528060028152602001610c8d60f21b815250906104085760405162461bcd60e51b81526004016103ff919061203c565b60405180910390fd5b50506000908152600560205260409020546001600160a01b031690565b6000610430826106a3565b9050806001600160a01b0316836001600160a01b0316141560405180604001604052806002815260200161191960f11b815250906104815760405162461bcd60e51b81526004016103ff919061203c565b50336001600160a01b038216148061049e575061049e8133610a94565b60405180604001604052806002815260200161323360f01b815250906104d75760405162461bcd60e51b81526004016103ff919061203c565b506103b68383610cc0565b6104ec3382610d2e565b60405180604001604052806002815260200161191b60f11b815250906105255760405162461bcd60e51b81526004016103ff919061203c565b506103b6838383610dd1565b600061053c83610779565b8210604051806040016040528060028152602001610ccd60f21b815250906105775760405162461bcd60e51b81526004016103ff919061203c565b50506001600160a01b03919091166000908152600760209081526040808320938352929052205490565b6103b683838360405180602001604052806000815250610898565b600180546105c990612115565b6040805180820190915260018152601960f91b60208201529150156106015760405162461bcd60e51b81526004016103ff919061203c565b508151610615906001906020850190611a0d565b5080516103b6906002906020840190611a0d565b600061063460095490565b821060405180604001604052806002815260200161333560f01b8152509061066f5760405162461bcd60e51b81526004016103ff919061203c565b506009828154811061069157634e487b7160e01b600052603260045260246000fd5b90600052602060002001549050919050565b60008181526003602090815260408083205481518083019092526002825261323160f01b928201929092526001600160a01b0390911690816106f85760405162461bcd60e51b81526004016103ff919061203c565b5092915050565b600080546040805180820190915260018152600d60fa1b6020820152906001600160a01b031633146107445760405162461bcd60e51b81526004016103ff919061203c565b50600061074f610a36565b905061075b8382610f25565b600b805490600061076b83612150565b90915550909150505b919050565b604080518082019091526002815261032360f41b60208201526000906001600160a01b0383166107bc5760405162461bcd60e51b81526004016103ff919061203c565b50506001600160a01b031660009081526004602052604090205490565b60606002805461032090612115565b604080518082019091526002815261323560f01b60208201526001600160a01b03831633141561082b5760405162461bcd60e51b81526004016103ff919061203c565b503360008181526006602090815260408083206001600160a01b03871680855290835292819020805460ff191686151590811790915590519081529192917f17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31910160405180910390a35050565b6108a23383610d2e565b60405180604001604052806002815260200161191b60f11b815250906108db5760405162461bcd60e51b81526004016103ff919061203c565b506108e884848484610f43565b50505050565b6000546040805180820190915260018152601960f91b6020820152906001600160a01b0316156109315760405162461bcd60e51b81526004016103ff919061203c565b50604080518082019091526002815261353960f01b60208201526001600160a01b0382166109725760405162461bcd60e51b81526004016103ff919061203c565b50600080546001600160a01b0319166001600160a01b0392909216919091179055565b60606109a082610ca3565b60405180604001604052806002815260200161333360f01b815250906109d95760405162461bcd60e51b81526004016103ff919061203c565b5060006109e4610f93565b90506000815111610a045760405180602001604052806000815250610a2f565b80610a0e8461105a565b604051602001610a1f929190611fa5565b6040516020818303038152906040525b9392505050565b600b54604051603160f91b60208201526001600160601b03193060601b16602182015260358101919091526000906055016040516020818303038152906040528051906020012060001c905090565b6000546001600160a01b031690565b6001600160a01b03918216600090815260066020908152604080832093909416825291909152205460ff1690565b6000546040805180820190915260018152600d60fa1b6020820152906001600160a01b03163314610b065760405162461bcd60e51b81526004016103ff919061203c565b506000610b11610a85565b6040516301c5004b60e21b8152600481018590529091506000906001600160a01b03831690630714012c906024016101006040518083038186803b158015610b5857600080fd5b505afa158015610b6c573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610b909190611e5e565b60e0810151604051627080db60e51b815261ffff90911660048201529091506000906001600160a01b03841690630e101b609060240160206040518083038186803b158015610bde57600080fd5b505afa158015610bf2573d6000803e3d6000fd5b505050506040513d601f19601f82011682018060405250810190610c169190611b5c565b9050610c43848360c001516001600160701b0316836001600160a01b03166111739092919063ffffffff16565b610c4c856111c5565b5050505050565b60006001600160e01b031982166380ac58cd60e01b1480610c8457506001600160e01b03198216635b5e139f60e01b145b8061030b57506301ffc9a760e01b6001600160e01b031983161461030b565b6000908152600360205260409020546001600160a01b0316151590565b600081815260056020526040902080546001600160a01b0319166001600160a01b0384169081179091558190610cf5826106a3565b6001600160a01b03167f8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b92560405160405180910390a45050565b6000610d3982610ca3565b60405180604001604052806002815260200161064760f31b81525090610d725760405162461bcd60e51b81526004016103ff919061203c565b506000610d7e836106a3565b9050806001600160a01b0316846001600160a01b03161480610db95750836001600160a01b0316610dae846103bb565b6001600160a01b0316145b80610dc95750610dc98185610a94565b949350505050565b826001600160a01b0316610de4826106a3565b6001600160a01b03161460405180604001604052806002815260200161333160f01b81525090610e275760405162461bcd60e51b81526004016103ff919061203c565b50604080518082019091526002815261199960f11b60208201526001600160a01b038316610e685760405162461bcd60e51b81526004016103ff919061203c565b50610e7483838361125a565b610e7f600082610cc0565b6001600160a01b0383166000908152600460205260408120805460019290610ea89084906120d2565b90915550506001600160a01b0382166000908152600460205260408120805460019290610ed69084906120a6565b909155505060008181526003602052604080822080546001600160a01b0319166001600160a01b0386811691821790925591518493918716916000805160206121fe83398151915291a4505050565b610f3f8282604051806020016040528060008152506112de565b5050565b610f4e848484610dd1565b610f5a8484848461132e565b60405180604001604052806002815260200161323760f01b81525090610c4c5760405162461bcd60e51b81526004016103ff919061203c565b60606000610f9f610a85565b6001600160a01b0316630cac36b26040518163ffffffff1660e01b815260040160006040518083038186803b158015610fd757600080fd5b505afa158015610feb573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526110139190810190611d18565b905060008151116110335760405180602001604052806000815250611054565b806040516020016110449190611fd4565b6040516020818303038152906040525b91505090565b60608161107e5750506040805180820190915260018152600360fc1b602082015290565b8160005b81156110a8578061109281612150565b91506110a19050600a836120be565b9150611082565b6000816001600160401b038111156110d057634e487b7160e01b600052604160045260246000fd5b6040519080825280601f01601f1916602001820160405280156110fa576020820181803683370190505b5090505b8415610dc95761110f6001836120d2565b915061111c600a8661216b565b6111279060306120a6565b60f81b81838151811061114a57634e487b7160e01b600052603260045260246000fd5b60200101906001600160f81b031916908160001a90535061116c600a866120be565b94506110fe565b604080516001600160a01b038416602482015260448082018490528251808303909101815260649091019091526020810180516001600160e01b031663a9059cbb60e01b1790526103b6908490611453565b60006111d0826106a3565b90506111de8160008461125a565b6111e9600083610cc0565b6001600160a01b03811660009081526004602052604081208054600192906112129084906120d2565b909155505060008281526003602052604080822080546001600160a01b0319169055518391906001600160a01b038416906000805160206121fe833981519152908390a45050565b611265838383611525565b61126d610a85565b60405163144a94ef60e21b81526001600160a01b038581166004830152848116602483015260448201849052919091169063512a53bc90606401600060405180830381600087803b1580156112c157600080fd5b505af11580156112d5573d6000803e3d6000fd5b50505050505050565b6112e883836115dd565b6112f5600084848461132e565b60405180604001604052806002815260200161323760f01b815250906108e85760405162461bcd60e51b81526004016103ff919061203c565b60006001600160a01b0384163b1561144857604051630a85bd0160e11b81526001600160a01b0385169063150b7a0290611372903390899088908890600401611fff565b602060405180830381600087803b15801561138c57600080fd5b505af19250505080156113bc575060408051601f3d908101601f191682019092526113b991810190611cfc565b60015b61142e573d8080156113ea576040519150601f19603f3d011682016040523d82523d6000602084013e6113ef565b606091505b508051611426576040805180820182526002815261323760f01b6020820152905162461bcd60e51b81526103ff919060040161203c565b805181602001fd5b6001600160e01b031916630a85bd0160e11b149050610dc9565b506001949350505050565b60006114a8826040518060400160405280602081526020017f5361666545524332303a206c6f772d6c6576656c2063616c6c206661696c6564815250856001600160a01b03166116e39092919063ffffffff16565b8051909150156103b657808060200190518101906114c69190611cc4565b6103b65760405162461bcd60e51b815260206004820152602a60248201527f5361666545524332303a204552433230206f7065726174696f6e20646964206e6044820152691bdd081cdd58d8d9595960b21b60648201526084016103ff565b6001600160a01b0383166115805761157b81600980546000838152600a60205260408120829055600182018355919091527f6e1540171b6c0c960b71a7020d9f60077f6af931a8bbf590da0223dacf75c7af0155565b6115a3565b816001600160a01b0316836001600160a01b0316146115a3576115a383826116f2565b6001600160a01b0382166115ba576103b68161178f565b826001600160a01b0316826001600160a01b0316146103b6576103b68282611868565b604080518082019091526002815261323960f01b60208201526001600160a01b03831661161d5760405162461bcd60e51b81526004016103ff919061203c565b5061162781610ca3565b1560405180604001604052806002815260200161033360f41b815250906116615760405162461bcd60e51b81526004016103ff919061203c565b5061166e6000838361125a565b6001600160a01b03821660009081526004602052604081208054600192906116979084906120a6565b909155505060008181526003602052604080822080546001600160a01b0319166001600160a01b03861690811790915590518392906000805160206121fe833981519152908290a45050565b6060610dc984846000856118ac565b600060016116ff84610779565b61170991906120d2565b60008381526008602052604090205490915080821461175c576001600160a01b03841660009081526007602090815260408083208584528252808320548484528184208190558352600890915290208190555b5060009182526008602090815260408084208490556001600160a01b039094168352600781528383209183525290812055565b6009546000906117a1906001906120d2565b6000838152600a6020526040812054600980549394509092849081106117d757634e487b7160e01b600052603260045260246000fd5b90600052602060002001549050806009838154811061180657634e487b7160e01b600052603260045260246000fd5b6000918252602080832090910192909255828152600a9091526040808220849055858252812055600980548061184c57634e487b7160e01b600052603160045260246000fd5b6001900381819060005260206000200160009055905550505050565b600061187383610779565b6001600160a01b039093166000908152600760209081526040808320868452825280832085905593825260089052919091209190915550565b60608247101561190d5760405162461bcd60e51b815260206004820152602660248201527f416464726573733a20696e73756666696369656e742062616c616e636520666f6044820152651c8818d85b1b60d21b60648201526084016103ff565b843b61195b5760405162461bcd60e51b815260206004820152601d60248201527f416464726573733a2063616c6c20746f206e6f6e2d636f6e747261637400000060448201526064016103ff565b600080866001600160a01b031685876040516119779190611f89565b60006040518083038185875af1925050503d80600081146119b4576040519150601f19603f3d011682016040523d82523d6000602084013e6119b9565b606091505b50915091506119c98282866119d4565b979650505050505050565b606083156119e3575081610a2f565b8251156119f35782518084602001fd5b8160405162461bcd60e51b81526004016103ff919061203c565b828054611a1990612115565b90600052602060002090601f016020900481019282611a3b5760008555611a81565b82601f10611a5457805160ff1916838001178555611a81565b82800160010185558215611a81579182015b82811115611a81578251825591602001919060010190611a66565b50611a8d929150611a91565b5090565b5b80821115611a8d5760008155600101611a92565b6000611ab9611ab48461207f565b61204f565b9050828152838383011115611acd57600080fd5b828260208301376000602084830101529392505050565b600082601f830112611af4578081fd5b610a2f83833560208501611aa6565b80516001600160701b038116811461077457600080fd5b805161ffff8116811461077457600080fd5b805163ffffffff8116811461077457600080fd5b600060208284031215611b51578081fd5b8135610a2f816121c1565b600060208284031215611b6d578081fd5b8151610a2f816121c1565b60008060408385031215611b8a578081fd5b8235611b95816121c1565b91506020830135611ba5816121c1565b809150509250929050565b600080600060608486031215611bc4578081fd5b8335611bcf816121c1565b92506020840135611bdf816121c1565b929592945050506040919091013590565b60008060008060808587031215611c05578081fd5b8435611c10816121c1565b93506020850135611c20816121c1565b92506040850135915060608501356001600160401b03811115611c41578182fd5b8501601f81018713611c51578182fd5b611c6087823560208401611aa6565b91505092959194509250565b60008060408385031215611c7e578182fd5b8235611c89816121c1565b91506020830135611ba5816121d9565b60008060408385031215611cab578081fd5b8235611cb6816121c1565b946020939093013593505050565b600060208284031215611cd5578081fd5b8151610a2f816121d9565b600060208284031215611cf1578081fd5b8135610a2f816121e7565b600060208284031215611d0d578081fd5b8151610a2f816121e7565b600060208284031215611d29578081fd5b81516001600160401b03811115611d3e578182fd5b8201601f81018413611d4e578182fd5b8051611d5c611ab48261207f565b818152856020838501011115611d70578384fd5b611d818260208301602086016120e9565b95945050505050565b60008060408385031215611d9c578182fd5b82356001600160401b0380821115611db2578384fd5b611dbe86838701611ae4565b93506020850135915080821115611dd3578283fd5b50611de085828601611ae4565b9150509250929050565b600080600060608486031215611dfe578081fd5b83356001600160401b0380821115611e14578283fd5b611e2087838801611ae4565b94506020860135915080821115611e35578283fd5b50611e4286828701611ae4565b9250506040840135611e53816121c1565b809150509250925092565b6000610100808385031215611e71578182fd5b604051908101906001600160401b0382118183101715611e9357611e936121ab565b81604052611ea084611b03565b8152611eae60208501611b1a565b6020820152611ebf60408501611b2c565b6040820152611ed060608501611b2c565b6060820152611ee160808501611b2c565b6080820152611ef260a08501611b2c565b60a0820152611f0360c08501611b03565b60c0820152611f1460e08501611b1a565b60e0820152949350505050565b600060208284031215611f32578081fd5b5035919050565b60008060408385031215611f4b578182fd5b823591506020830135611ba5816121c1565b60008151808452611f758160208601602086016120e9565b601f01601f19169290920160200192915050565b60008251611f9b8184602087016120e9565b9190910192915050565b60008351611fb78184602088016120e9565b835190830190611fcb8183602088016120e9565b01949350505050565b60008251611fe68184602087016120e9565b66626f72726f772f60c81b920191825250600701919050565b6001600160a01b038581168252841660208201526040810183905260806060820181905260009061203290830184611f5d565b9695505050505050565b602081526000610a2f6020830184611f5d565b604051601f8201601f191681016001600160401b0381118282101715612077576120776121ab565b604052919050565b60006001600160401b03821115612098576120986121ab565b50601f01601f191660200190565b600082198211156120b9576120b961217f565b500190565b6000826120cd576120cd612195565b500490565b6000828210156120e4576120e461217f565b500390565b60005b838110156121045781810151838201526020016120ec565b838111156108e85750506000910152565b600181811c9082168061212957607f821691505b6020821081141561214a57634e487b7160e01b600052602260045260246000fd5b50919050565b60006000198214156121645761216461217f565b5060010190565b60008261217a5761217a612195565b500690565b634e487b7160e01b600052601160045260246000fd5b634e487b7160e01b600052601260045260246000fd5b634e487b7160e01b600052604160045260246000fd5b6001600160a01b03811681146121d657600080fd5b50565b80151581146121d657600080fd5b6001600160e01b0319811681146121d657600080fdfeddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3efa2646970667358221220cc2be07caf24a5b36647a2250a2a7ceefb7cf14856ae7f172e688f5e3ef4a2e364736f6c63430008040033",
}

// BorrowTokenABI is the input ABI used to generate the binding from.
// Deprecated: Use BorrowTokenMetaData.ABI instead.
var BorrowTokenABI = BorrowTokenMetaData.ABI

// BorrowTokenBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BorrowTokenMetaData.Bin instead.
var BorrowTokenBin = BorrowTokenMetaData.Bin

// DeployBorrowToken deploys a new Ethereum contract, binding an instance of BorrowToken to it.
func DeployBorrowToken(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *BorrowToken, error) {
	parsed, err := BorrowTokenMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BorrowTokenBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BorrowToken{BorrowTokenCaller: BorrowTokenCaller{contract: contract}, BorrowTokenTransactor: BorrowTokenTransactor{contract: contract}, BorrowTokenFilterer: BorrowTokenFilterer{contract: contract}}, nil
}

// BorrowToken is an auto generated Go binding around an Ethereum contract.
type BorrowToken struct {
	BorrowTokenCaller     // Read-only binding to the contract
	BorrowTokenTransactor // Write-only binding to the contract
	BorrowTokenFilterer   // Log filterer for contract events
}

// BorrowTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type BorrowTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BorrowTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BorrowTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BorrowTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BorrowTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BorrowTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BorrowTokenSession struct {
	Contract     *BorrowToken      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BorrowTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BorrowTokenCallerSession struct {
	Contract *BorrowTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// BorrowTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BorrowTokenTransactorSession struct {
	Contract     *BorrowTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// BorrowTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type BorrowTokenRaw struct {
	Contract *BorrowToken // Generic contract binding to access the raw methods on
}

// BorrowTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BorrowTokenCallerRaw struct {
	Contract *BorrowTokenCaller // Generic read-only contract binding to access the raw methods on
}

// BorrowTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BorrowTokenTransactorRaw struct {
	Contract *BorrowTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBorrowToken creates a new instance of BorrowToken, bound to a specific deployed contract.
func NewBorrowToken(address common.Address, backend bind.ContractBackend) (*BorrowToken, error) {
	contract, err := bindBorrowToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BorrowToken{BorrowTokenCaller: BorrowTokenCaller{contract: contract}, BorrowTokenTransactor: BorrowTokenTransactor{contract: contract}, BorrowTokenFilterer: BorrowTokenFilterer{contract: contract}}, nil
}

// NewBorrowTokenCaller creates a new read-only instance of BorrowToken, bound to a specific deployed contract.
func NewBorrowTokenCaller(address common.Address, caller bind.ContractCaller) (*BorrowTokenCaller, error) {
	contract, err := bindBorrowToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BorrowTokenCaller{contract: contract}, nil
}

// NewBorrowTokenTransactor creates a new write-only instance of BorrowToken, bound to a specific deployed contract.
func NewBorrowTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*BorrowTokenTransactor, error) {
	contract, err := bindBorrowToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BorrowTokenTransactor{contract: contract}, nil
}

// NewBorrowTokenFilterer creates a new log filterer instance of BorrowToken, bound to a specific deployed contract.
func NewBorrowTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*BorrowTokenFilterer, error) {
	contract, err := bindBorrowToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BorrowTokenFilterer{contract: contract}, nil
}

// bindBorrowToken binds a generic wrapper to an already deployed contract.
func bindBorrowToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BorrowTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BorrowToken *BorrowTokenRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BorrowToken.Contract.BorrowTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BorrowToken *BorrowTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BorrowToken.Contract.BorrowTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BorrowToken *BorrowTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BorrowToken.Contract.BorrowTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BorrowToken *BorrowTokenCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BorrowToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BorrowToken *BorrowTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BorrowToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BorrowToken *BorrowTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BorrowToken.Contract.contract.Transact(opts, method, params...)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_BorrowToken *BorrowTokenCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var out []interface{}
	err := _BorrowToken.contract.Call(opts, &out, "balanceOf", owner)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_BorrowToken *BorrowTokenSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _BorrowToken.Contract.BalanceOf(&_BorrowToken.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(address owner) view returns(uint256)
func (_BorrowToken *BorrowTokenCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _BorrowToken.Contract.BalanceOf(&_BorrowToken.CallOpts, owner)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_BorrowToken *BorrowTokenCaller) GetApproved(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BorrowToken.contract.Call(opts, &out, "getApproved", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_BorrowToken *BorrowTokenSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _BorrowToken.Contract.GetApproved(&_BorrowToken.CallOpts, tokenId)
}

// GetApproved is a free data retrieval call binding the contract method 0x081812fc.
//
// Solidity: function getApproved(uint256 tokenId) view returns(address)
func (_BorrowToken *BorrowTokenCallerSession) GetApproved(tokenId *big.Int) (common.Address, error) {
	return _BorrowToken.Contract.GetApproved(&_BorrowToken.CallOpts, tokenId)
}

// GetEnterprise is a free data retrieval call binding the contract method 0xe298a2fd.
//
// Solidity: function getEnterprise() view returns(address)
func (_BorrowToken *BorrowTokenCaller) GetEnterprise(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BorrowToken.contract.Call(opts, &out, "getEnterprise")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEnterprise is a free data retrieval call binding the contract method 0xe298a2fd.
//
// Solidity: function getEnterprise() view returns(address)
func (_BorrowToken *BorrowTokenSession) GetEnterprise() (common.Address, error) {
	return _BorrowToken.Contract.GetEnterprise(&_BorrowToken.CallOpts)
}

// GetEnterprise is a free data retrieval call binding the contract method 0xe298a2fd.
//
// Solidity: function getEnterprise() view returns(address)
func (_BorrowToken *BorrowTokenCallerSession) GetEnterprise() (common.Address, error) {
	return _BorrowToken.Contract.GetEnterprise(&_BorrowToken.CallOpts)
}

// GetNextTokenId is a free data retrieval call binding the contract method 0xcaa0f92a.
//
// Solidity: function getNextTokenId() view returns(uint256)
func (_BorrowToken *BorrowTokenCaller) GetNextTokenId(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BorrowToken.contract.Call(opts, &out, "getNextTokenId")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetNextTokenId is a free data retrieval call binding the contract method 0xcaa0f92a.
//
// Solidity: function getNextTokenId() view returns(uint256)
func (_BorrowToken *BorrowTokenSession) GetNextTokenId() (*big.Int, error) {
	return _BorrowToken.Contract.GetNextTokenId(&_BorrowToken.CallOpts)
}

// GetNextTokenId is a free data retrieval call binding the contract method 0xcaa0f92a.
//
// Solidity: function getNextTokenId() view returns(uint256)
func (_BorrowToken *BorrowTokenCallerSession) GetNextTokenId() (*big.Int, error) {
	return _BorrowToken.Contract.GetNextTokenId(&_BorrowToken.CallOpts)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_BorrowToken *BorrowTokenCaller) IsApprovedForAll(opts *bind.CallOpts, owner common.Address, operator common.Address) (bool, error) {
	var out []interface{}
	err := _BorrowToken.contract.Call(opts, &out, "isApprovedForAll", owner, operator)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_BorrowToken *BorrowTokenSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _BorrowToken.Contract.IsApprovedForAll(&_BorrowToken.CallOpts, owner, operator)
}

// IsApprovedForAll is a free data retrieval call binding the contract method 0xe985e9c5.
//
// Solidity: function isApprovedForAll(address owner, address operator) view returns(bool)
func (_BorrowToken *BorrowTokenCallerSession) IsApprovedForAll(owner common.Address, operator common.Address) (bool, error) {
	return _BorrowToken.Contract.IsApprovedForAll(&_BorrowToken.CallOpts, owner, operator)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BorrowToken *BorrowTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BorrowToken.contract.Call(opts, &out, "name")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BorrowToken *BorrowTokenSession) Name() (string, error) {
	return _BorrowToken.Contract.Name(&_BorrowToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() view returns(string)
func (_BorrowToken *BorrowTokenCallerSession) Name() (string, error) {
	return _BorrowToken.Contract.Name(&_BorrowToken.CallOpts)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_BorrowToken *BorrowTokenCaller) OwnerOf(opts *bind.CallOpts, tokenId *big.Int) (common.Address, error) {
	var out []interface{}
	err := _BorrowToken.contract.Call(opts, &out, "ownerOf", tokenId)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_BorrowToken *BorrowTokenSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _BorrowToken.Contract.OwnerOf(&_BorrowToken.CallOpts, tokenId)
}

// OwnerOf is a free data retrieval call binding the contract method 0x6352211e.
//
// Solidity: function ownerOf(uint256 tokenId) view returns(address)
func (_BorrowToken *BorrowTokenCallerSession) OwnerOf(tokenId *big.Int) (common.Address, error) {
	return _BorrowToken.Contract.OwnerOf(&_BorrowToken.CallOpts, tokenId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BorrowToken *BorrowTokenCaller) SupportsInterface(opts *bind.CallOpts, interfaceId [4]byte) (bool, error) {
	var out []interface{}
	err := _BorrowToken.contract.Call(opts, &out, "supportsInterface", interfaceId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BorrowToken *BorrowTokenSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BorrowToken.Contract.SupportsInterface(&_BorrowToken.CallOpts, interfaceId)
}

// SupportsInterface is a free data retrieval call binding the contract method 0x01ffc9a7.
//
// Solidity: function supportsInterface(bytes4 interfaceId) view returns(bool)
func (_BorrowToken *BorrowTokenCallerSession) SupportsInterface(interfaceId [4]byte) (bool, error) {
	return _BorrowToken.Contract.SupportsInterface(&_BorrowToken.CallOpts, interfaceId)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BorrowToken *BorrowTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _BorrowToken.contract.Call(opts, &out, "symbol")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BorrowToken *BorrowTokenSession) Symbol() (string, error) {
	return _BorrowToken.Contract.Symbol(&_BorrowToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() view returns(string)
func (_BorrowToken *BorrowTokenCallerSession) Symbol() (string, error) {
	return _BorrowToken.Contract.Symbol(&_BorrowToken.CallOpts)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_BorrowToken *BorrowTokenCaller) TokenByIndex(opts *bind.CallOpts, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BorrowToken.contract.Call(opts, &out, "tokenByIndex", index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_BorrowToken *BorrowTokenSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _BorrowToken.Contract.TokenByIndex(&_BorrowToken.CallOpts, index)
}

// TokenByIndex is a free data retrieval call binding the contract method 0x4f6ccce7.
//
// Solidity: function tokenByIndex(uint256 index) view returns(uint256)
func (_BorrowToken *BorrowTokenCallerSession) TokenByIndex(index *big.Int) (*big.Int, error) {
	return _BorrowToken.Contract.TokenByIndex(&_BorrowToken.CallOpts, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_BorrowToken *BorrowTokenCaller) TokenOfOwnerByIndex(opts *bind.CallOpts, owner common.Address, index *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _BorrowToken.contract.Call(opts, &out, "tokenOfOwnerByIndex", owner, index)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_BorrowToken *BorrowTokenSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _BorrowToken.Contract.TokenOfOwnerByIndex(&_BorrowToken.CallOpts, owner, index)
}

// TokenOfOwnerByIndex is a free data retrieval call binding the contract method 0x2f745c59.
//
// Solidity: function tokenOfOwnerByIndex(address owner, uint256 index) view returns(uint256)
func (_BorrowToken *BorrowTokenCallerSession) TokenOfOwnerByIndex(owner common.Address, index *big.Int) (*big.Int, error) {
	return _BorrowToken.Contract.TokenOfOwnerByIndex(&_BorrowToken.CallOpts, owner, index)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_BorrowToken *BorrowTokenCaller) TokenURI(opts *bind.CallOpts, tokenId *big.Int) (string, error) {
	var out []interface{}
	err := _BorrowToken.contract.Call(opts, &out, "tokenURI", tokenId)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_BorrowToken *BorrowTokenSession) TokenURI(tokenId *big.Int) (string, error) {
	return _BorrowToken.Contract.TokenURI(&_BorrowToken.CallOpts, tokenId)
}

// TokenURI is a free data retrieval call binding the contract method 0xc87b56dd.
//
// Solidity: function tokenURI(uint256 tokenId) view returns(string)
func (_BorrowToken *BorrowTokenCallerSession) TokenURI(tokenId *big.Int) (string, error) {
	return _BorrowToken.Contract.TokenURI(&_BorrowToken.CallOpts, tokenId)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BorrowToken *BorrowTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BorrowToken.contract.Call(opts, &out, "totalSupply")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BorrowToken *BorrowTokenSession) TotalSupply() (*big.Int, error) {
	return _BorrowToken.Contract.TotalSupply(&_BorrowToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() view returns(uint256)
func (_BorrowToken *BorrowTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _BorrowToken.Contract.TotalSupply(&_BorrowToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_BorrowToken *BorrowTokenTransactor) Approve(opts *bind.TransactOpts, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BorrowToken.contract.Transact(opts, "approve", to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_BorrowToken *BorrowTokenSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BorrowToken.Contract.Approve(&_BorrowToken.TransactOpts, to, tokenId)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(address to, uint256 tokenId) returns()
func (_BorrowToken *BorrowTokenTransactorSession) Approve(to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BorrowToken.Contract.Approve(&_BorrowToken.TransactOpts, to, tokenId)
}

// Burn is a paid mutator transaction binding the contract method 0xfcd3533c.
//
// Solidity: function burn(uint256 tokenId, address burner) returns()
func (_BorrowToken *BorrowTokenTransactor) Burn(opts *bind.TransactOpts, tokenId *big.Int, burner common.Address) (*types.Transaction, error) {
	return _BorrowToken.contract.Transact(opts, "burn", tokenId, burner)
}

// Burn is a paid mutator transaction binding the contract method 0xfcd3533c.
//
// Solidity: function burn(uint256 tokenId, address burner) returns()
func (_BorrowToken *BorrowTokenSession) Burn(tokenId *big.Int, burner common.Address) (*types.Transaction, error) {
	return _BorrowToken.Contract.Burn(&_BorrowToken.TransactOpts, tokenId, burner)
}

// Burn is a paid mutator transaction binding the contract method 0xfcd3533c.
//
// Solidity: function burn(uint256 tokenId, address burner) returns()
func (_BorrowToken *BorrowTokenTransactorSession) Burn(tokenId *big.Int, burner common.Address) (*types.Transaction, error) {
	return _BorrowToken.Contract.Burn(&_BorrowToken.TransactOpts, tokenId, burner)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string name, string symbol, address enterprise) returns()
func (_BorrowToken *BorrowTokenTransactor) Initialize(opts *bind.TransactOpts, name string, symbol string, enterprise common.Address) (*types.Transaction, error) {
	return _BorrowToken.contract.Transact(opts, "initialize", name, symbol, enterprise)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string name, string symbol, address enterprise) returns()
func (_BorrowToken *BorrowTokenSession) Initialize(name string, symbol string, enterprise common.Address) (*types.Transaction, error) {
	return _BorrowToken.Contract.Initialize(&_BorrowToken.TransactOpts, name, symbol, enterprise)
}

// Initialize is a paid mutator transaction binding the contract method 0x077f224a.
//
// Solidity: function initialize(string name, string symbol, address enterprise) returns()
func (_BorrowToken *BorrowTokenTransactorSession) Initialize(name string, symbol string, enterprise common.Address) (*types.Transaction, error) {
	return _BorrowToken.Contract.Initialize(&_BorrowToken.TransactOpts, name, symbol, enterprise)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string name_, string symbol_) returns()
func (_BorrowToken *BorrowTokenTransactor) Initialize0(opts *bind.TransactOpts, name_ string, symbol_ string) (*types.Transaction, error) {
	return _BorrowToken.contract.Transact(opts, "initialize0", name_, symbol_)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string name_, string symbol_) returns()
func (_BorrowToken *BorrowTokenSession) Initialize0(name_ string, symbol_ string) (*types.Transaction, error) {
	return _BorrowToken.Contract.Initialize0(&_BorrowToken.TransactOpts, name_, symbol_)
}

// Initialize0 is a paid mutator transaction binding the contract method 0x4cd88b76.
//
// Solidity: function initialize(string name_, string symbol_) returns()
func (_BorrowToken *BorrowTokenTransactorSession) Initialize0(name_ string, symbol_ string) (*types.Transaction, error) {
	return _BorrowToken.Contract.Initialize0(&_BorrowToken.TransactOpts, name_, symbol_)
}

// Initialize1 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address enterprise) returns()
func (_BorrowToken *BorrowTokenTransactor) Initialize1(opts *bind.TransactOpts, enterprise common.Address) (*types.Transaction, error) {
	return _BorrowToken.contract.Transact(opts, "initialize1", enterprise)
}

// Initialize1 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address enterprise) returns()
func (_BorrowToken *BorrowTokenSession) Initialize1(enterprise common.Address) (*types.Transaction, error) {
	return _BorrowToken.Contract.Initialize1(&_BorrowToken.TransactOpts, enterprise)
}

// Initialize1 is a paid mutator transaction binding the contract method 0xc4d66de8.
//
// Solidity: function initialize(address enterprise) returns()
func (_BorrowToken *BorrowTokenTransactorSession) Initialize1(enterprise common.Address) (*types.Transaction, error) {
	return _BorrowToken.Contract.Initialize1(&_BorrowToken.TransactOpts, enterprise)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256)
func (_BorrowToken *BorrowTokenTransactor) Mint(opts *bind.TransactOpts, to common.Address) (*types.Transaction, error) {
	return _BorrowToken.contract.Transact(opts, "mint", to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256)
func (_BorrowToken *BorrowTokenSession) Mint(to common.Address) (*types.Transaction, error) {
	return _BorrowToken.Contract.Mint(&_BorrowToken.TransactOpts, to)
}

// Mint is a paid mutator transaction binding the contract method 0x6a627842.
//
// Solidity: function mint(address to) returns(uint256)
func (_BorrowToken *BorrowTokenTransactorSession) Mint(to common.Address) (*types.Transaction, error) {
	return _BorrowToken.Contract.Mint(&_BorrowToken.TransactOpts, to)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_BorrowToken *BorrowTokenTransactor) SafeTransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BorrowToken.contract.Transact(opts, "safeTransferFrom", from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_BorrowToken *BorrowTokenSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BorrowToken.Contract.SafeTransferFrom(&_BorrowToken.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom is a paid mutator transaction binding the contract method 0x42842e0e.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId) returns()
func (_BorrowToken *BorrowTokenTransactorSession) SafeTransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BorrowToken.Contract.SafeTransferFrom(&_BorrowToken.TransactOpts, from, to, tokenId)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_BorrowToken *BorrowTokenTransactor) SafeTransferFrom0(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _BorrowToken.contract.Transact(opts, "safeTransferFrom0", from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_BorrowToken *BorrowTokenSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _BorrowToken.Contract.SafeTransferFrom0(&_BorrowToken.TransactOpts, from, to, tokenId, _data)
}

// SafeTransferFrom0 is a paid mutator transaction binding the contract method 0xb88d4fde.
//
// Solidity: function safeTransferFrom(address from, address to, uint256 tokenId, bytes _data) returns()
func (_BorrowToken *BorrowTokenTransactorSession) SafeTransferFrom0(from common.Address, to common.Address, tokenId *big.Int, _data []byte) (*types.Transaction, error) {
	return _BorrowToken.Contract.SafeTransferFrom0(&_BorrowToken.TransactOpts, from, to, tokenId, _data)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_BorrowToken *BorrowTokenTransactor) SetApprovalForAll(opts *bind.TransactOpts, operator common.Address, approved bool) (*types.Transaction, error) {
	return _BorrowToken.contract.Transact(opts, "setApprovalForAll", operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_BorrowToken *BorrowTokenSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _BorrowToken.Contract.SetApprovalForAll(&_BorrowToken.TransactOpts, operator, approved)
}

// SetApprovalForAll is a paid mutator transaction binding the contract method 0xa22cb465.
//
// Solidity: function setApprovalForAll(address operator, bool approved) returns()
func (_BorrowToken *BorrowTokenTransactorSession) SetApprovalForAll(operator common.Address, approved bool) (*types.Transaction, error) {
	return _BorrowToken.Contract.SetApprovalForAll(&_BorrowToken.TransactOpts, operator, approved)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_BorrowToken *BorrowTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BorrowToken.contract.Transact(opts, "transferFrom", from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_BorrowToken *BorrowTokenSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BorrowToken.Contract.TransferFrom(&_BorrowToken.TransactOpts, from, to, tokenId)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(address from, address to, uint256 tokenId) returns()
func (_BorrowToken *BorrowTokenTransactorSession) TransferFrom(from common.Address, to common.Address, tokenId *big.Int) (*types.Transaction, error) {
	return _BorrowToken.Contract.TransferFrom(&_BorrowToken.TransactOpts, from, to, tokenId)
}

// BorrowTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the BorrowToken contract.
type BorrowTokenApprovalIterator struct {
	Event *BorrowTokenApproval // Event containing the contract specifics and raw log

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
func (it *BorrowTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BorrowTokenApproval)
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
		it.Event = new(BorrowTokenApproval)
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
func (it *BorrowTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BorrowTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BorrowTokenApproval represents a Approval event raised by the BorrowToken contract.
type BorrowTokenApproval struct {
	Owner    common.Address
	Approved common.Address
	TokenId  *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_BorrowToken *BorrowTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, approved []common.Address, tokenId []*big.Int) (*BorrowTokenApprovalIterator, error) {

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

	logs, sub, err := _BorrowToken.contract.FilterLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BorrowTokenApprovalIterator{contract: _BorrowToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: event Approval(address indexed owner, address indexed approved, uint256 indexed tokenId)
func (_BorrowToken *BorrowTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *BorrowTokenApproval, owner []common.Address, approved []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _BorrowToken.contract.WatchLogs(opts, "Approval", ownerRule, approvedRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BorrowTokenApproval)
				if err := _BorrowToken.contract.UnpackLog(event, "Approval", log); err != nil {
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
func (_BorrowToken *BorrowTokenFilterer) ParseApproval(log types.Log) (*BorrowTokenApproval, error) {
	event := new(BorrowTokenApproval)
	if err := _BorrowToken.contract.UnpackLog(event, "Approval", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BorrowTokenApprovalForAllIterator is returned from FilterApprovalForAll and is used to iterate over the raw logs and unpacked data for ApprovalForAll events raised by the BorrowToken contract.
type BorrowTokenApprovalForAllIterator struct {
	Event *BorrowTokenApprovalForAll // Event containing the contract specifics and raw log

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
func (it *BorrowTokenApprovalForAllIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BorrowTokenApprovalForAll)
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
		it.Event = new(BorrowTokenApprovalForAll)
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
func (it *BorrowTokenApprovalForAllIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BorrowTokenApprovalForAllIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BorrowTokenApprovalForAll represents a ApprovalForAll event raised by the BorrowToken contract.
type BorrowTokenApprovalForAll struct {
	Owner    common.Address
	Operator common.Address
	Approved bool
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalForAll is a free log retrieval operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_BorrowToken *BorrowTokenFilterer) FilterApprovalForAll(opts *bind.FilterOpts, owner []common.Address, operator []common.Address) (*BorrowTokenApprovalForAllIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BorrowToken.contract.FilterLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return &BorrowTokenApprovalForAllIterator{contract: _BorrowToken.contract, event: "ApprovalForAll", logs: logs, sub: sub}, nil
}

// WatchApprovalForAll is a free log subscription operation binding the contract event 0x17307eab39ab6107e8899845ad3d59bd9653f200f220920489ca2b5937696c31.
//
// Solidity: event ApprovalForAll(address indexed owner, address indexed operator, bool approved)
func (_BorrowToken *BorrowTokenFilterer) WatchApprovalForAll(opts *bind.WatchOpts, sink chan<- *BorrowTokenApprovalForAll, owner []common.Address, operator []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var operatorRule []interface{}
	for _, operatorItem := range operator {
		operatorRule = append(operatorRule, operatorItem)
	}

	logs, sub, err := _BorrowToken.contract.WatchLogs(opts, "ApprovalForAll", ownerRule, operatorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BorrowTokenApprovalForAll)
				if err := _BorrowToken.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
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
func (_BorrowToken *BorrowTokenFilterer) ParseApprovalForAll(log types.Log) (*BorrowTokenApprovalForAll, error) {
	event := new(BorrowTokenApprovalForAll)
	if err := _BorrowToken.contract.UnpackLog(event, "ApprovalForAll", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// BorrowTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the BorrowToken contract.
type BorrowTokenTransferIterator struct {
	Event *BorrowTokenTransfer // Event containing the contract specifics and raw log

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
func (it *BorrowTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(BorrowTokenTransfer)
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
		it.Event = new(BorrowTokenTransfer)
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
func (it *BorrowTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *BorrowTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// BorrowTokenTransfer represents a Transfer event raised by the BorrowToken contract.
type BorrowTokenTransfer struct {
	From    common.Address
	To      common.Address
	TokenId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_BorrowToken *BorrowTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address, tokenId []*big.Int) (*BorrowTokenTransferIterator, error) {

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

	logs, sub, err := _BorrowToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return &BorrowTokenTransferIterator{contract: _BorrowToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: event Transfer(address indexed from, address indexed to, uint256 indexed tokenId)
func (_BorrowToken *BorrowTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *BorrowTokenTransfer, from []common.Address, to []common.Address, tokenId []*big.Int) (event.Subscription, error) {

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

	logs, sub, err := _BorrowToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule, tokenIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(BorrowTokenTransfer)
				if err := _BorrowToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
func (_BorrowToken *BorrowTokenFilterer) ParseTransfer(log types.Log) (*BorrowTokenTransfer, error) {
	event := new(BorrowTokenTransfer)
	if err := _BorrowToken.contract.UnpackLog(event, "Transfer", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
