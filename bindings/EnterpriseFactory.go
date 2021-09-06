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

// EnterpriseFactoryMetaData contains all meta data concerning the EnterpriseFactory contract.
var EnterpriseFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"enterpriseImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"powerTokenImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"interestTokenImpl\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"borrowTokenImpl\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"liquidityToken\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"baseUri\",\"type\":\"string\"},{\"indexed\":false,\"internalType\":\"address\",\"name\":\"deployed\",\"type\":\"address\"}],\"name\":\"EnterpriseDeployed\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"},{\"internalType\":\"contractIERC20Metadata\",\"name\":\"liquidityToken\",\"type\":\"address\"},{\"internalType\":\"string\",\"name\":\"baseUri\",\"type\":\"string\"},{\"internalType\":\"uint16\",\"name\":\"gcFeePercent\",\"type\":\"uint16\"},{\"internalType\":\"contractIConverter\",\"name\":\"converter\",\"type\":\"address\"}],\"name\":\"deploy\",\"outputs\":[{\"internalType\":\"contractEnterprise\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"contractProxyAdmin\",\"name\":\"admin\",\"type\":\"address\"}],\"name\":\"deployService\",\"outputs\":[{\"internalType\":\"contractPowerToken\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getBorrowTokenImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEnterpriseImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getInterestTokenImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getPowerTokenImpl\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x6101006040523480156200001257600080fd5b50604051620024893803806200248983398101604081905262000035916200019b565b604080518082019091526002815261353560f01b60208201526001600160a01b038516620000815760405162461bcd60e51b8152600401620000789190620001f7565b60405180910390fd5b506040805180820190915260028152611a9b60f11b60208201526001600160a01b038416620000c55760405162461bcd60e51b8152600401620000789190620001f7565b50604080518082019091526002815261353760f01b60208201526001600160a01b038316620001095760405162461bcd60e51b8152600401620000789190620001f7565b5060408051808201909152600281526106a760f31b60208201526001600160a01b0382166200014d5760405162461bcd60e51b8152600401620000789190620001f7565b506001600160601b0319606094851b811660805292841b831660a05290831b821660c05290911b1660e0526200024d565b80516001600160a01b03811681146200019657600080fd5b919050565b60008060008060808587031215620001b1578384fd5b620001bc856200017e565b9350620001cc602086016200017e565b9250620001dc604086016200017e565b9150620001ec606086016200017e565b905092959194509250565b6000602080835283518082850152825b81811015620002255785810183015185820160400152820162000207565b81811115620002375783604083870101525b50601f01601f1916929092016040019392505050565b60805160601c60a05160601c60c05160601c60e05160601c6121dd620002ac6000396000818160a0015261063f01526000818161012c015261055001526000818160ee015261047201526000818160c7015261018a01526121dd6000f3fe60806040523480156200001157600080fd5b50600436106200006a5760003560e01c806335ff0aef146200006f57806362832955146200009e57806382f9340914620000c557806383b9a00914620000ec578063c31011cc1462000113578063d2c598b5146200012a575b600080fd5b6200008662000080366004620006f0565b62000151565b60405162000095919062000969565b60405180910390f35b7f000000000000000000000000000000000000000000000000000000000000000062000086565b7f000000000000000000000000000000000000000000000000000000000000000062000086565b7f000000000000000000000000000000000000000000000000000000000000000062000086565b6200008662000124366004620006ca565b6200046a565b7f000000000000000000000000000000000000000000000000000000000000000062000086565b600080604051620001629062000665565b604051809103906000f0801580156200017f573d6000803e3d6000fd5b5090506000620001b07f0000000000000000000000000000000000000000000000000000000000000000836200049e565b60405163f2fde38b60e01b81529091506001600160a01b0383169063f2fde38b90620001e190849060040162000969565b600060405180830381600087803b158015620001fc57600080fd5b505af115801562000211573d6000803e3d6000fd5b5050604051636815f33760e01b81526001600160a01b0384169250636815f337915062000251908d908d908c908c908c908c908b903390600401620009c3565b600060405180830381600087803b1580156200026c57600080fd5b505af115801562000281573d6000803e3d6000fd5b5050505060006200030c896001600160a01b03166395d89b416040518163ffffffff1660e01b815260040160006040518083038186803b158015620002c557600080fd5b505afa158015620002da573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f19168201604052620003049190810190620007a9565b8385620004fb565b90506000620003958a6001600160a01b03166395d89b416040518163ffffffff1660e01b815260040160006040518083038186803b1580156200034e57600080fd5b505afa15801562000363573d6000803e3d6000fd5b505050506040513d6000823e601f3d908101601f191682016040526200038d9190810190620007a9565b8486620005ea565b60405163ef1f9f3960e01b81526001600160a01b038c81166004830152848116602483015280831660448301529192509084169063ef1f9f3990606401600060405180830381600087803b158015620003ed57600080fd5b505af115801562000402573d6000803e3d6000fd5b505050505050876001600160a01b0316336001600160a01b03167faf94cde051ff1210197fe31203f2210a396ffd110ab74cf8014d6bfd600842d58c8c8b8b87604051620004559594939291906200097d565b60405180910390a39998505050505050505050565b6000620004987f0000000000000000000000000000000000000000000000000000000000000000836200049e565b92915050565b60008282604051620004b09062000673565b6001600160a01b03928316815291166020820152606060408201819052600090820152608001604051809103906000f080158015620004f3573d6000803e3d6000fd5b509392505050565b60008084604051602001620005119190620008b4565b60405160208183030381529060405290506000856040516020016200053791906200094b565b60405160208183030381529060405290506000620005767f0000000000000000000000000000000000000000000000000000000000000000866200049e565b6040516303bf912560e11b81529091506001600160a01b0382169063077f224a90620005ab90869086908b9060040162000a25565b600060405180830381600087803b158015620005c657600080fd5b505af1158015620005db573d6000803e3d6000fd5b50929998505050505050505050565b60008084604051602001620006009190620008ef565b604051602081830303815290604052905060008560405160200162000626919062000920565b60405160208183030381529060405290506000620005767f0000000000000000000000000000000000000000000000000000000000000000866200049e565b6107d28062000aca83390190565b610f0c806200129c83390190565b60008083601f84011262000693578182fd5b5081356001600160401b03811115620006aa578182fd5b602083019150836020828501011115620006c357600080fd5b9250929050565b600060208284031215620006dc578081fd5b8135620006e98162000ab0565b9392505050565b600080600080600080600060a0888a0312156200070b578283fd5b87356001600160401b038082111562000722578485fd5b620007308b838c0162000681565b909950975060208a01359150620007478262000ab0565b909550604089013590808211156200075d578485fd5b506200076c8a828b0162000681565b909550935050606088013561ffff8116811462000787578283fd5b91506080880135620007998162000ab0565b8091505092959891949750929550565b600060208284031215620007bb578081fd5b81516001600160401b0380821115620007d2578283fd5b818401915084601f830112620007e6578283fd5b815181811115620007fb57620007fb62000a9a565b604051601f8201601f19908116603f0116810190838211818310171562000826576200082662000a9a565b816040528281528760208487010111156200083f578586fd5b6200085283602083016020880162000a67565b979650505050505050565b81835281816020850137506000828201602090810191909152601f909101601f19169091010190565b60008151808452620008a081602086016020860162000a67565b601f01601f19169290920160200192915050565b70024b73a32b932b9ba102132b0b934b7339607d1b815260008251620008e281601185016020870162000a67565b9190910160110192915050565b6602137b93937bb960cd1b8152600082516200091381600785016020870162000a67565b9190910160070192915050565b603160f91b8152600082516200093e81600185016020870162000a67565b9190910160010192915050565b606960f81b8152600082516200093e81600185016020870162000a67565b6001600160a01b0391909116815260200190565b606081526000620009936060830187896200085d565b8281036020840152620009a88186886200085d565b91505060018060a01b03831660408301529695505050505050565b60c081526000620009d960c083018a8c6200085d565b8281036020840152620009ee81898b6200085d565b61ffff97909716604084015250506001600160a01b039384166060820152918316608083015290911660a090910152949350505050565b60608152600062000a3a606083018662000886565b828103602084015262000a4e818662000886565b91505060018060a01b0383166040830152949350505050565b60005b8381101562000a8457818101518382015260200162000a6a565b8381111562000a94576000848401525b50505050565b634e487b7160e01b600052604160045260246000fd5b6001600160a01b038116811462000ac657600080fd5b5056fe608060405234801561001057600080fd5b50600080546001600160a01b031916339081178255604051909182917f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0908290a350610771806100616000396000f3fe60806040526004361061006b5760003560e01c8063204e1c7a14610070578063715018a6146100a65780637eff275e146100bd5780638da5cb5b146100dd5780639623609d146100f257806399a88ec414610105578063f2fde38b14610125578063f3b7dead14610145575b600080fd5b34801561007c57600080fd5b5061009061008b3660046104fa565b610165565b60405161009d9190610641565b60405180910390f35b3480156100b257600080fd5b506100bb6101f6565b005b3480156100c957600080fd5b506100bb6100d8366004610539565b610266565b3480156100e957600080fd5b506100906102f7565b6100bb610100366004610571565b610306565b34801561011157600080fd5b506100bb610120366004610539565b61039c565b34801561013157600080fd5b506100bb6101403660046104fa565b6103f7565b34801561015157600080fd5b506100906101603660046104fa565b6104d4565b6000806000836001600160a01b031660405161018b90635c60da1b60e01b815260040190565b600060405180830381855afa9150503d80600081146101c6576040519150601f19603f3d011682016040523d82523d6000602084013e6101cb565b606091505b5091509150816101da57600080fd5b808060200190518101906101ee919061051d565b949350505050565b336101ff6102f7565b6001600160a01b03161461022e5760405162461bcd60e51b8152600401610225906106b8565b60405180910390fd5b600080546040516001600160a01b039091169060008051602061071c833981519152908390a3600080546001600160a01b0319169055565b3361026f6102f7565b6001600160a01b0316146102955760405162461bcd60e51b8152600401610225906106b8565b6040516308f2839760e41b81526001600160a01b03831690638f283970906102c1908490600401610641565b600060405180830381600087803b1580156102db57600080fd5b505af11580156102ef573d6000803e3d6000fd5b505050505050565b6000546001600160a01b031690565b3361030f6102f7565b6001600160a01b0316146103355760405162461bcd60e51b8152600401610225906106b8565b60405163278f794360e11b81526001600160a01b03841690634f1ef2869034906103659086908690600401610655565b6000604051808303818588803b15801561037e57600080fd5b505af1158015610392573d6000803e3d6000fd5b5050505050505050565b336103a56102f7565b6001600160a01b0316146103cb5760405162461bcd60e51b8152600401610225906106b8565b604051631b2ce7f360e11b81526001600160a01b03831690633659cfe6906102c1908490600401610641565b336104006102f7565b6001600160a01b0316146104265760405162461bcd60e51b8152600401610225906106b8565b6001600160a01b03811661048b5760405162461bcd60e51b815260206004820152602660248201527f4f776e61626c653a206e6577206f776e657220697320746865207a65726f206160448201526564647265737360d01b6064820152608401610225565b600080546040516001600160a01b038085169392169160008051602061071c83398151915291a3600080546001600160a01b0319166001600160a01b0392909216919091179055565b6000806000836001600160a01b031660405161018b906303e1469160e61b815260040190565b60006020828403121561050b578081fd5b813561051681610703565b9392505050565b60006020828403121561052e578081fd5b815161051681610703565b6000806040838503121561054b578081fd5b823561055681610703565b9150602083013561056681610703565b809150509250929050565b600080600060608486031215610585578081fd5b833561059081610703565b925060208401356105a081610703565b915060408401356001600160401b03808211156105bb578283fd5b818601915086601f8301126105ce578283fd5b8135818111156105e0576105e06106ed565b604051601f8201601f19908116603f01168101908382118183101715610608576106086106ed565b81604052828152896020848701011115610620578586fd5b82602086016020830137856020848301015280955050505050509250925092565b6001600160a01b0391909116815260200190565b60018060a01b0383168152600060206040818401528351806040850152825b8181101561069057858101830151858201606001528201610674565b818111156106a15783606083870101525b50601f01601f191692909201606001949350505050565b6020808252818101527f4f776e61626c653a2063616c6c6572206973206e6f7420746865206f776e6572604082015260600190565b634e487b7160e01b600052604160045260246000fd5b6001600160a01b038116811461071857600080fd5b5056fe8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0a26469706673582212209f58ddf2e62e932313ef03a43663de644f281755cb24cf59c7579808f4a16aa264736f6c63430008040033608060405260405162000f0c38038062000f0c8339810160408190526200002691620004cc565b82816200005560017f360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbd620005fb565b60008051602062000ec5833981519152146200008157634e487b7160e01b600052600160045260246000fd5b6200008f82826000620000ff565b50620000bf905060017fb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6104620005fb565b60008051602062000ea583398151915214620000eb57634e487b7160e01b600052600160045260246000fd5b620000f68262000170565b50505062000664565b6200010a83620001cb565b6040516001600160a01b038416907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a26000825111806200014c5750805b156200016b576200016983836200029360201b620002601760201c565b505b505050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f6200019b620002c2565b604080516001600160a01b03928316815291841660208301520160405180910390a1620001c881620002fb565b50565b620001e1816200038b60201b6200028c1760201c565b620002495760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b60648201526084015b60405180910390fd5b806200027260008051602062000ec583398151915260001b6200039160201b620002081760201c565b80546001600160a01b0319166001600160a01b039290921691909117905550565b6060620002bb838360405180606001604052806027815260200162000ee56027913962000394565b9392505050565b6000620002ec60008051602062000ea583398151915260001b6200039160201b620002081760201c565b546001600160a01b0316919050565b6001600160a01b038116620003625760405162461bcd60e51b815260206004820152602660248201527f455243313936373a206e65772061646d696e20697320746865207a65726f206160448201526564647265737360d01b606482015260840162000240565b806200027260008051602062000ea583398151915260001b6200039160201b620002081760201c565b3b151590565b90565b6060833b620003f55760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b606482015260840162000240565b600080856001600160a01b031685604051620004129190620005a8565b600060405180830381855af49150503d80600081146200044f576040519150601f19603f3d011682016040523d82523d6000602084013e62000454565b606091505b5090925090506200046782828662000471565b9695505050505050565b6060831562000482575081620002bb565b825115620004935782518084602001fd5b8160405162461bcd60e51b8152600401620002409190620005c6565b80516001600160a01b0381168114620004c757600080fd5b919050565b600080600060608486031215620004e1578283fd5b620004ec84620004af565b9250620004fc60208501620004af565b60408501519092506001600160401b038082111562000519578283fd5b818601915086601f8301126200052d578283fd5b8151818111156200054257620005426200064e565b604051601f8201601f19908116603f011681019083821181831017156200056d576200056d6200064e565b8160405282815289602084870101111562000586578586fd5b620005998360208301602088016200061f565b80955050505050509250925092565b60008251620005bc8184602087016200061f565b9190910192915050565b6020815260008251806020840152620005e78160408501602087016200061f565b601f01601f19169190910160400192915050565b6000828210156200061a57634e487b7160e01b81526011600452602481fd5b500390565b60005b838110156200063c57818101518382015260200162000622565b83811115620001695750506000910152565b634e487b7160e01b600052604160045260246000fd5b61083180620006746000396000f3fe60806040526004361061004e5760003560e01c80633659cfe6146100655780634f1ef286146100855780635c60da1b146100985780638f283970146100c9578063f851a440146100e95761005d565b3661005d5761005b6100fe565b005b61005b6100fe565b34801561007157600080fd5b5061005b610080366004610682565b610118565b61005b61009336600461069c565b61015f565b3480156100a457600080fd5b506100ad6101d0565b6040516001600160a01b03909116815260200160405180910390f35b3480156100d557600080fd5b5061005b6100e4366004610682565b61020b565b3480156100f557600080fd5b506100ad610235565b610106610292565b610116610111610331565b61033b565b565b61012061035f565b6001600160a01b0316336001600160a01b031614156101575761015481604051806020016040528060008152506000610380565b50565b6101546100fe565b61016761035f565b6001600160a01b0316336001600160a01b031614156101c8576101c38383838080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525060019250610380915050565b505050565b6101c36100fe565b60006101da61035f565b6001600160a01b0316336001600160a01b03161415610200576101fb610331565b905090565b6102086100fe565b90565b61021361035f565b6001600160a01b0316336001600160a01b0316141561015757610154816103df565b600061023f61035f565b6001600160a01b0316336001600160a01b03161415610200576101fb61035f565b606061028583836040518060600160405280602781526020016107d560279139610433565b9392505050565b3b151590565b61029a61035f565b6001600160a01b0316336001600160a01b031614156101165760405162461bcd60e51b815260206004820152604260248201527f5472616e73706172656e745570677261646561626c6550726f78793a2061646d60448201527f696e2063616e6e6f742066616c6c6261636b20746f2070726f78792074617267606482015261195d60f21b608482015260a4015b60405180910390fd5b60006101fb610507565b3660008037600080366000845af43d6000803e80801561035a573d6000f35b3d6000fd5b60006000805160206107958339815191525b546001600160a01b0316919050565b6103898361051d565b6040516001600160a01b038416907fbc7cd75a20ee27fd9adebab32041f755214dbc6bffa90cc0225b39da2e5c2d3b90600090a26000825111806103ca5750805b156101c3576103d98383610260565b50505050565b7f7e644d79422f17c01e4894b5f4f588d331ebfa28653d42ae832dc59e38c9798f61040861035f565b604080516001600160a01b03928316815291841660208301520160405180910390a1610154816105b3565b6060833b6104925760405162461bcd60e51b815260206004820152602660248201527f416464726573733a2064656c65676174652063616c6c20746f206e6f6e2d636f6044820152651b9d1c9858dd60d21b6064820152608401610328565b600080856001600160a01b0316856040516104ad9190610719565b600060405180830381855af49150503d80600081146104e8576040519150601f19603f3d011682016040523d82523d6000602084013e6104ed565b606091505b50915091506104fd82828661062d565b9695505050505050565b60006000805160206107b5833981519152610371565b803b6105815760405162461bcd60e51b815260206004820152602d60248201527f455243313936373a206e657720696d706c656d656e746174696f6e206973206e60448201526c1bdd08184818dbdb9d1c9858dd609a1b6064820152608401610328565b806000805160206107b58339815191525b80546001600160a01b0319166001600160a01b039290921691909117905550565b6001600160a01b0381166106185760405162461bcd60e51b815260206004820152602660248201527f455243313936373a206e65772061646d696e20697320746865207a65726f206160448201526564647265737360d01b6064820152608401610328565b80600080516020610795833981519152610592565b6060831561063c575081610285565b82511561064c5782518084602001fd5b8160405162461bcd60e51b81526004016103289190610735565b80356001600160a01b038116811461067d57600080fd5b919050565b600060208284031215610693578081fd5b61028582610666565b6000806000604084860312156106b0578182fd5b6106b984610666565b925060208401356001600160401b03808211156106d4578384fd5b818601915086601f8301126106e7578384fd5b8135818111156106f5578485fd5b876020828501011115610706578485fd5b6020830194508093505050509250925092565b6000825161072b818460208701610768565b9190910192915050565b6020815260008251806020840152610754816040850160208701610768565b601f01601f19169190910160400192915050565b60005b8381101561078357818101518382015260200161076b565b838111156103d9575050600091015256feb53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220834feeba798ffeacd61d277e2e2e474f2b7d3b5c3c506bf659be3c5ef6046bdb64736f6c63430008040033b53127684a568b3173ae13b9f8a6016e243e63b6e8ee1178d6a717850b5d6103360894a13ba1a3210667c828492db98dca3e2076cc3735a920a3ca505d382bbc416464726573733a206c6f772d6c6576656c2064656c65676174652063616c6c206661696c6564a2646970667358221220d90371cbcbe207956a2f47854a9e7fe553fedd0e613cc871658c7fc21f3f9df464736f6c63430008040033",
}

// EnterpriseFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use EnterpriseFactoryMetaData.ABI instead.
var EnterpriseFactoryABI = EnterpriseFactoryMetaData.ABI

// EnterpriseFactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use EnterpriseFactoryMetaData.Bin instead.
var EnterpriseFactoryBin = EnterpriseFactoryMetaData.Bin

// DeployEnterpriseFactory deploys a new Ethereum contract, binding an instance of EnterpriseFactory to it.
func DeployEnterpriseFactory(auth *bind.TransactOpts, backend bind.ContractBackend, enterpriseImpl common.Address, powerTokenImpl common.Address, interestTokenImpl common.Address, borrowTokenImpl common.Address) (common.Address, *types.Transaction, *EnterpriseFactory, error) {
	parsed, err := EnterpriseFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(EnterpriseFactoryBin), backend, enterpriseImpl, powerTokenImpl, interestTokenImpl, borrowTokenImpl)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &EnterpriseFactory{EnterpriseFactoryCaller: EnterpriseFactoryCaller{contract: contract}, EnterpriseFactoryTransactor: EnterpriseFactoryTransactor{contract: contract}, EnterpriseFactoryFilterer: EnterpriseFactoryFilterer{contract: contract}}, nil
}

// EnterpriseFactory is an auto generated Go binding around an Ethereum contract.
type EnterpriseFactory struct {
	EnterpriseFactoryCaller     // Read-only binding to the contract
	EnterpriseFactoryTransactor // Write-only binding to the contract
	EnterpriseFactoryFilterer   // Log filterer for contract events
}

// EnterpriseFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type EnterpriseFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnterpriseFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type EnterpriseFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnterpriseFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type EnterpriseFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// EnterpriseFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type EnterpriseFactorySession struct {
	Contract     *EnterpriseFactory // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// EnterpriseFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type EnterpriseFactoryCallerSession struct {
	Contract *EnterpriseFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// EnterpriseFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type EnterpriseFactoryTransactorSession struct {
	Contract     *EnterpriseFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// EnterpriseFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type EnterpriseFactoryRaw struct {
	Contract *EnterpriseFactory // Generic contract binding to access the raw methods on
}

// EnterpriseFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type EnterpriseFactoryCallerRaw struct {
	Contract *EnterpriseFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// EnterpriseFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type EnterpriseFactoryTransactorRaw struct {
	Contract *EnterpriseFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewEnterpriseFactory creates a new instance of EnterpriseFactory, bound to a specific deployed contract.
func NewEnterpriseFactory(address common.Address, backend bind.ContractBackend) (*EnterpriseFactory, error) {
	contract, err := bindEnterpriseFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &EnterpriseFactory{EnterpriseFactoryCaller: EnterpriseFactoryCaller{contract: contract}, EnterpriseFactoryTransactor: EnterpriseFactoryTransactor{contract: contract}, EnterpriseFactoryFilterer: EnterpriseFactoryFilterer{contract: contract}}, nil
}

// NewEnterpriseFactoryCaller creates a new read-only instance of EnterpriseFactory, bound to a specific deployed contract.
func NewEnterpriseFactoryCaller(address common.Address, caller bind.ContractCaller) (*EnterpriseFactoryCaller, error) {
	contract, err := bindEnterpriseFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &EnterpriseFactoryCaller{contract: contract}, nil
}

// NewEnterpriseFactoryTransactor creates a new write-only instance of EnterpriseFactory, bound to a specific deployed contract.
func NewEnterpriseFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*EnterpriseFactoryTransactor, error) {
	contract, err := bindEnterpriseFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &EnterpriseFactoryTransactor{contract: contract}, nil
}

// NewEnterpriseFactoryFilterer creates a new log filterer instance of EnterpriseFactory, bound to a specific deployed contract.
func NewEnterpriseFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*EnterpriseFactoryFilterer, error) {
	contract, err := bindEnterpriseFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &EnterpriseFactoryFilterer{contract: contract}, nil
}

// bindEnterpriseFactory binds a generic wrapper to an already deployed contract.
func bindEnterpriseFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(EnterpriseFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnterpriseFactory *EnterpriseFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnterpriseFactory.Contract.EnterpriseFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnterpriseFactory *EnterpriseFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnterpriseFactory.Contract.EnterpriseFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnterpriseFactory *EnterpriseFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnterpriseFactory.Contract.EnterpriseFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_EnterpriseFactory *EnterpriseFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _EnterpriseFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_EnterpriseFactory *EnterpriseFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _EnterpriseFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_EnterpriseFactory *EnterpriseFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _EnterpriseFactory.Contract.contract.Transact(opts, method, params...)
}

// GetBorrowTokenImpl is a free data retrieval call binding the contract method 0x62832955.
//
// Solidity: function getBorrowTokenImpl() view returns(address)
func (_EnterpriseFactory *EnterpriseFactoryCaller) GetBorrowTokenImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EnterpriseFactory.contract.Call(opts, &out, "getBorrowTokenImpl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetBorrowTokenImpl is a free data retrieval call binding the contract method 0x62832955.
//
// Solidity: function getBorrowTokenImpl() view returns(address)
func (_EnterpriseFactory *EnterpriseFactorySession) GetBorrowTokenImpl() (common.Address, error) {
	return _EnterpriseFactory.Contract.GetBorrowTokenImpl(&_EnterpriseFactory.CallOpts)
}

// GetBorrowTokenImpl is a free data retrieval call binding the contract method 0x62832955.
//
// Solidity: function getBorrowTokenImpl() view returns(address)
func (_EnterpriseFactory *EnterpriseFactoryCallerSession) GetBorrowTokenImpl() (common.Address, error) {
	return _EnterpriseFactory.Contract.GetBorrowTokenImpl(&_EnterpriseFactory.CallOpts)
}

// GetEnterpriseImpl is a free data retrieval call binding the contract method 0x82f93409.
//
// Solidity: function getEnterpriseImpl() view returns(address)
func (_EnterpriseFactory *EnterpriseFactoryCaller) GetEnterpriseImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EnterpriseFactory.contract.Call(opts, &out, "getEnterpriseImpl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetEnterpriseImpl is a free data retrieval call binding the contract method 0x82f93409.
//
// Solidity: function getEnterpriseImpl() view returns(address)
func (_EnterpriseFactory *EnterpriseFactorySession) GetEnterpriseImpl() (common.Address, error) {
	return _EnterpriseFactory.Contract.GetEnterpriseImpl(&_EnterpriseFactory.CallOpts)
}

// GetEnterpriseImpl is a free data retrieval call binding the contract method 0x82f93409.
//
// Solidity: function getEnterpriseImpl() view returns(address)
func (_EnterpriseFactory *EnterpriseFactoryCallerSession) GetEnterpriseImpl() (common.Address, error) {
	return _EnterpriseFactory.Contract.GetEnterpriseImpl(&_EnterpriseFactory.CallOpts)
}

// GetInterestTokenImpl is a free data retrieval call binding the contract method 0xd2c598b5.
//
// Solidity: function getInterestTokenImpl() view returns(address)
func (_EnterpriseFactory *EnterpriseFactoryCaller) GetInterestTokenImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EnterpriseFactory.contract.Call(opts, &out, "getInterestTokenImpl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetInterestTokenImpl is a free data retrieval call binding the contract method 0xd2c598b5.
//
// Solidity: function getInterestTokenImpl() view returns(address)
func (_EnterpriseFactory *EnterpriseFactorySession) GetInterestTokenImpl() (common.Address, error) {
	return _EnterpriseFactory.Contract.GetInterestTokenImpl(&_EnterpriseFactory.CallOpts)
}

// GetInterestTokenImpl is a free data retrieval call binding the contract method 0xd2c598b5.
//
// Solidity: function getInterestTokenImpl() view returns(address)
func (_EnterpriseFactory *EnterpriseFactoryCallerSession) GetInterestTokenImpl() (common.Address, error) {
	return _EnterpriseFactory.Contract.GetInterestTokenImpl(&_EnterpriseFactory.CallOpts)
}

// GetPowerTokenImpl is a free data retrieval call binding the contract method 0x83b9a009.
//
// Solidity: function getPowerTokenImpl() view returns(address)
func (_EnterpriseFactory *EnterpriseFactoryCaller) GetPowerTokenImpl(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _EnterpriseFactory.contract.Call(opts, &out, "getPowerTokenImpl")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// GetPowerTokenImpl is a free data retrieval call binding the contract method 0x83b9a009.
//
// Solidity: function getPowerTokenImpl() view returns(address)
func (_EnterpriseFactory *EnterpriseFactorySession) GetPowerTokenImpl() (common.Address, error) {
	return _EnterpriseFactory.Contract.GetPowerTokenImpl(&_EnterpriseFactory.CallOpts)
}

// GetPowerTokenImpl is a free data retrieval call binding the contract method 0x83b9a009.
//
// Solidity: function getPowerTokenImpl() view returns(address)
func (_EnterpriseFactory *EnterpriseFactoryCallerSession) GetPowerTokenImpl() (common.Address, error) {
	return _EnterpriseFactory.Contract.GetPowerTokenImpl(&_EnterpriseFactory.CallOpts)
}

// Deploy is a paid mutator transaction binding the contract method 0x35ff0aef.
//
// Solidity: function deploy(string name, address liquidityToken, string baseUri, uint16 gcFeePercent, address converter) returns(address)
func (_EnterpriseFactory *EnterpriseFactoryTransactor) Deploy(opts *bind.TransactOpts, name string, liquidityToken common.Address, baseUri string, gcFeePercent uint16, converter common.Address) (*types.Transaction, error) {
	return _EnterpriseFactory.contract.Transact(opts, "deploy", name, liquidityToken, baseUri, gcFeePercent, converter)
}

// Deploy is a paid mutator transaction binding the contract method 0x35ff0aef.
//
// Solidity: function deploy(string name, address liquidityToken, string baseUri, uint16 gcFeePercent, address converter) returns(address)
func (_EnterpriseFactory *EnterpriseFactorySession) Deploy(name string, liquidityToken common.Address, baseUri string, gcFeePercent uint16, converter common.Address) (*types.Transaction, error) {
	return _EnterpriseFactory.Contract.Deploy(&_EnterpriseFactory.TransactOpts, name, liquidityToken, baseUri, gcFeePercent, converter)
}

// Deploy is a paid mutator transaction binding the contract method 0x35ff0aef.
//
// Solidity: function deploy(string name, address liquidityToken, string baseUri, uint16 gcFeePercent, address converter) returns(address)
func (_EnterpriseFactory *EnterpriseFactoryTransactorSession) Deploy(name string, liquidityToken common.Address, baseUri string, gcFeePercent uint16, converter common.Address) (*types.Transaction, error) {
	return _EnterpriseFactory.Contract.Deploy(&_EnterpriseFactory.TransactOpts, name, liquidityToken, baseUri, gcFeePercent, converter)
}

// DeployService is a paid mutator transaction binding the contract method 0xc31011cc.
//
// Solidity: function deployService(address admin) returns(address)
func (_EnterpriseFactory *EnterpriseFactoryTransactor) DeployService(opts *bind.TransactOpts, admin common.Address) (*types.Transaction, error) {
	return _EnterpriseFactory.contract.Transact(opts, "deployService", admin)
}

// DeployService is a paid mutator transaction binding the contract method 0xc31011cc.
//
// Solidity: function deployService(address admin) returns(address)
func (_EnterpriseFactory *EnterpriseFactorySession) DeployService(admin common.Address) (*types.Transaction, error) {
	return _EnterpriseFactory.Contract.DeployService(&_EnterpriseFactory.TransactOpts, admin)
}

// DeployService is a paid mutator transaction binding the contract method 0xc31011cc.
//
// Solidity: function deployService(address admin) returns(address)
func (_EnterpriseFactory *EnterpriseFactoryTransactorSession) DeployService(admin common.Address) (*types.Transaction, error) {
	return _EnterpriseFactory.Contract.DeployService(&_EnterpriseFactory.TransactOpts, admin)
}

// EnterpriseFactoryEnterpriseDeployedIterator is returned from FilterEnterpriseDeployed and is used to iterate over the raw logs and unpacked data for EnterpriseDeployed events raised by the EnterpriseFactory contract.
type EnterpriseFactoryEnterpriseDeployedIterator struct {
	Event *EnterpriseFactoryEnterpriseDeployed // Event containing the contract specifics and raw log

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
func (it *EnterpriseFactoryEnterpriseDeployedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(EnterpriseFactoryEnterpriseDeployed)
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
		it.Event = new(EnterpriseFactoryEnterpriseDeployed)
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
func (it *EnterpriseFactoryEnterpriseDeployedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *EnterpriseFactoryEnterpriseDeployedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// EnterpriseFactoryEnterpriseDeployed represents a EnterpriseDeployed event raised by the EnterpriseFactory contract.
type EnterpriseFactoryEnterpriseDeployed struct {
	Creator        common.Address
	LiquidityToken common.Address
	Name           string
	BaseUri        string
	Deployed       common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterEnterpriseDeployed is a free log retrieval operation binding the contract event 0xaf94cde051ff1210197fe31203f2210a396ffd110ab74cf8014d6bfd600842d5.
//
// Solidity: event EnterpriseDeployed(address indexed creator, address indexed liquidityToken, string name, string baseUri, address deployed)
func (_EnterpriseFactory *EnterpriseFactoryFilterer) FilterEnterpriseDeployed(opts *bind.FilterOpts, creator []common.Address, liquidityToken []common.Address) (*EnterpriseFactoryEnterpriseDeployedIterator, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var liquidityTokenRule []interface{}
	for _, liquidityTokenItem := range liquidityToken {
		liquidityTokenRule = append(liquidityTokenRule, liquidityTokenItem)
	}

	logs, sub, err := _EnterpriseFactory.contract.FilterLogs(opts, "EnterpriseDeployed", creatorRule, liquidityTokenRule)
	if err != nil {
		return nil, err
	}
	return &EnterpriseFactoryEnterpriseDeployedIterator{contract: _EnterpriseFactory.contract, event: "EnterpriseDeployed", logs: logs, sub: sub}, nil
}

// WatchEnterpriseDeployed is a free log subscription operation binding the contract event 0xaf94cde051ff1210197fe31203f2210a396ffd110ab74cf8014d6bfd600842d5.
//
// Solidity: event EnterpriseDeployed(address indexed creator, address indexed liquidityToken, string name, string baseUri, address deployed)
func (_EnterpriseFactory *EnterpriseFactoryFilterer) WatchEnterpriseDeployed(opts *bind.WatchOpts, sink chan<- *EnterpriseFactoryEnterpriseDeployed, creator []common.Address, liquidityToken []common.Address) (event.Subscription, error) {

	var creatorRule []interface{}
	for _, creatorItem := range creator {
		creatorRule = append(creatorRule, creatorItem)
	}
	var liquidityTokenRule []interface{}
	for _, liquidityTokenItem := range liquidityToken {
		liquidityTokenRule = append(liquidityTokenRule, liquidityTokenItem)
	}

	logs, sub, err := _EnterpriseFactory.contract.WatchLogs(opts, "EnterpriseDeployed", creatorRule, liquidityTokenRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(EnterpriseFactoryEnterpriseDeployed)
				if err := _EnterpriseFactory.contract.UnpackLog(event, "EnterpriseDeployed", log); err != nil {
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

// ParseEnterpriseDeployed is a log parse operation binding the contract event 0xaf94cde051ff1210197fe31203f2210a396ffd110ab74cf8014d6bfd600842d5.
//
// Solidity: event EnterpriseDeployed(address indexed creator, address indexed liquidityToken, string name, string baseUri, address deployed)
func (_EnterpriseFactory *EnterpriseFactoryFilterer) ParseEnterpriseDeployed(log types.Log) (*EnterpriseFactoryEnterpriseDeployed, error) {
	event := new(EnterpriseFactoryEnterpriseDeployed)
	if err := _EnterpriseFactory.contract.UnpackLog(event, "EnterpriseDeployed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
