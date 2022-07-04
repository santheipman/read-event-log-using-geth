// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package market

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

// MarketMetaData contains all meta data concerning the Market contract.
var MarketMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"_USDAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"_NFTAddress\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"buyer\",\"type\":\"address\"}],\"name\":\"Buy\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"}],\"name\":\"CancelSell\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"uint256\",\"name\":\"_orderId\",\"type\":\"uint256\"},{\"indexed\":true,\"internalType\":\"address\",\"name\":\"_seller\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_tokenId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"_price\",\"type\":\"uint256\"}],\"name\":\"CreateSellOrder\",\"type\":\"event\"}]",
}

// MarketABI is the input ABI used to generate the binding from.
// Deprecated: Use MarketMetaData.ABI instead.
var MarketABI = MarketMetaData.ABI

// Market is an auto generated Go binding around an Ethereum contract.
type Market struct {
	MarketCaller     // Read-only binding to the contract
	MarketTransactor // Write-only binding to the contract
	MarketFilterer   // Log filterer for contract events
}

// MarketCaller is an auto generated read-only Go binding around an Ethereum contract.
type MarketCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MarketTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MarketFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MarketSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MarketSession struct {
	Contract     *Market           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MarketCallerSession struct {
	Contract *MarketCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MarketTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MarketTransactorSession struct {
	Contract     *MarketTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MarketRaw is an auto generated low-level Go binding around an Ethereum contract.
type MarketRaw struct {
	Contract *Market // Generic contract binding to access the raw methods on
}

// MarketCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MarketCallerRaw struct {
	Contract *MarketCaller // Generic read-only contract binding to access the raw methods on
}

// MarketTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MarketTransactorRaw struct {
	Contract *MarketTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMarket creates a new instance of Market, bound to a specific deployed contract.
func NewMarket(address common.Address, backend bind.ContractBackend) (*Market, error) {
	contract, err := bindMarket(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Market{MarketCaller: MarketCaller{contract: contract}, MarketTransactor: MarketTransactor{contract: contract}, MarketFilterer: MarketFilterer{contract: contract}}, nil
}

// NewMarketCaller creates a new read-only instance of Market, bound to a specific deployed contract.
func NewMarketCaller(address common.Address, caller bind.ContractCaller) (*MarketCaller, error) {
	contract, err := bindMarket(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MarketCaller{contract: contract}, nil
}

// NewMarketTransactor creates a new write-only instance of Market, bound to a specific deployed contract.
func NewMarketTransactor(address common.Address, transactor bind.ContractTransactor) (*MarketTransactor, error) {
	contract, err := bindMarket(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MarketTransactor{contract: contract}, nil
}

// NewMarketFilterer creates a new log filterer instance of Market, bound to a specific deployed contract.
func NewMarketFilterer(address common.Address, filterer bind.ContractFilterer) (*MarketFilterer, error) {
	contract, err := bindMarket(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MarketFilterer{contract: contract}, nil
}

// bindMarket binds a generic wrapper to an already deployed contract.
func bindMarket(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MarketABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Market *MarketRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Market.Contract.MarketCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Market *MarketRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.Contract.MarketTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Market *MarketRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Market.Contract.MarketTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Market *MarketCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Market.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Market *MarketTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Market.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Market *MarketTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Market.Contract.contract.Transact(opts, method, params...)
}

// MarketBuyIterator is returned from FilterBuy and is used to iterate over the raw logs and unpacked data for Buy events raised by the Market contract.
type MarketBuyIterator struct {
	Event *MarketBuy // Event containing the contract specifics and raw log

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
func (it *MarketBuyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketBuy)
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
		it.Event = new(MarketBuy)
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
func (it *MarketBuyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketBuyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketBuy represents a Buy event raised by the Market contract.
type MarketBuy struct {
	OrderId *big.Int
	Buyer   common.Address
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterBuy is a free log retrieval operation binding the contract event 0x7534e9c9396820aad756ba27f82162930f80bef863aabdf311e7dce26a79bccd.
//
// Solidity: event Buy(uint256 indexed _orderId, address indexed buyer)
func (_Market *MarketFilterer) FilterBuy(opts *bind.FilterOpts, _orderId []*big.Int, buyer []common.Address) (*MarketBuyIterator, error) {

	var _orderIdRule []interface{}
	for _, _orderIdItem := range _orderId {
		_orderIdRule = append(_orderIdRule, _orderIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "Buy", _orderIdRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return &MarketBuyIterator{contract: _Market.contract, event: "Buy", logs: logs, sub: sub}, nil
}

// WatchBuy is a free log subscription operation binding the contract event 0x7534e9c9396820aad756ba27f82162930f80bef863aabdf311e7dce26a79bccd.
//
// Solidity: event Buy(uint256 indexed _orderId, address indexed buyer)
func (_Market *MarketFilterer) WatchBuy(opts *bind.WatchOpts, sink chan<- *MarketBuy, _orderId []*big.Int, buyer []common.Address) (event.Subscription, error) {

	var _orderIdRule []interface{}
	for _, _orderIdItem := range _orderId {
		_orderIdRule = append(_orderIdRule, _orderIdItem)
	}
	var buyerRule []interface{}
	for _, buyerItem := range buyer {
		buyerRule = append(buyerRule, buyerItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "Buy", _orderIdRule, buyerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketBuy)
				if err := _Market.contract.UnpackLog(event, "Buy", log); err != nil {
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

// ParseBuy is a log parse operation binding the contract event 0x7534e9c9396820aad756ba27f82162930f80bef863aabdf311e7dce26a79bccd.
//
// Solidity: event Buy(uint256 indexed _orderId, address indexed buyer)
func (_Market *MarketFilterer) ParseBuy(log types.Log) (*MarketBuy, error) {
	event := new(MarketBuy)
	if err := _Market.contract.UnpackLog(event, "Buy", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketCancelSellIterator is returned from FilterCancelSell and is used to iterate over the raw logs and unpacked data for CancelSell events raised by the Market contract.
type MarketCancelSellIterator struct {
	Event *MarketCancelSell // Event containing the contract specifics and raw log

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
func (it *MarketCancelSellIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketCancelSell)
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
		it.Event = new(MarketCancelSell)
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
func (it *MarketCancelSellIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketCancelSellIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketCancelSell represents a CancelSell event raised by the Market contract.
type MarketCancelSell struct {
	OrderId *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCancelSell is a free log retrieval operation binding the contract event 0x87df48382ad11b84926580278f9ad32c3c1391b3abe3e28308695fab47e5c0fd.
//
// Solidity: event CancelSell(uint256 indexed _orderId)
func (_Market *MarketFilterer) FilterCancelSell(opts *bind.FilterOpts, _orderId []*big.Int) (*MarketCancelSellIterator, error) {

	var _orderIdRule []interface{}
	for _, _orderIdItem := range _orderId {
		_orderIdRule = append(_orderIdRule, _orderIdItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "CancelSell", _orderIdRule)
	if err != nil {
		return nil, err
	}
	return &MarketCancelSellIterator{contract: _Market.contract, event: "CancelSell", logs: logs, sub: sub}, nil
}

// WatchCancelSell is a free log subscription operation binding the contract event 0x87df48382ad11b84926580278f9ad32c3c1391b3abe3e28308695fab47e5c0fd.
//
// Solidity: event CancelSell(uint256 indexed _orderId)
func (_Market *MarketFilterer) WatchCancelSell(opts *bind.WatchOpts, sink chan<- *MarketCancelSell, _orderId []*big.Int) (event.Subscription, error) {

	var _orderIdRule []interface{}
	for _, _orderIdItem := range _orderId {
		_orderIdRule = append(_orderIdRule, _orderIdItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "CancelSell", _orderIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketCancelSell)
				if err := _Market.contract.UnpackLog(event, "CancelSell", log); err != nil {
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

// ParseCancelSell is a log parse operation binding the contract event 0x87df48382ad11b84926580278f9ad32c3c1391b3abe3e28308695fab47e5c0fd.
//
// Solidity: event CancelSell(uint256 indexed _orderId)
func (_Market *MarketFilterer) ParseCancelSell(log types.Log) (*MarketCancelSell, error) {
	event := new(MarketCancelSell)
	if err := _Market.contract.UnpackLog(event, "CancelSell", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// MarketCreateSellOrderIterator is returned from FilterCreateSellOrder and is used to iterate over the raw logs and unpacked data for CreateSellOrder events raised by the Market contract.
type MarketCreateSellOrderIterator struct {
	Event *MarketCreateSellOrder // Event containing the contract specifics and raw log

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
func (it *MarketCreateSellOrderIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(MarketCreateSellOrder)
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
		it.Event = new(MarketCreateSellOrder)
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
func (it *MarketCreateSellOrderIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *MarketCreateSellOrderIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// MarketCreateSellOrder represents a CreateSellOrder event raised by the Market contract.
type MarketCreateSellOrder struct {
	OrderId *big.Int
	Seller  common.Address
	TokenId *big.Int
	Price   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterCreateSellOrder is a free log retrieval operation binding the contract event 0x080493e207b9f00f98a17e07dddc2ea02575db77c2769f8ef3160198b10a6910.
//
// Solidity: event CreateSellOrder(uint256 indexed _orderId, address indexed _seller, uint256 _tokenId, uint256 _price)
func (_Market *MarketFilterer) FilterCreateSellOrder(opts *bind.FilterOpts, _orderId []*big.Int, _seller []common.Address) (*MarketCreateSellOrderIterator, error) {

	var _orderIdRule []interface{}
	for _, _orderIdItem := range _orderId {
		_orderIdRule = append(_orderIdRule, _orderIdItem)
	}
	var _sellerRule []interface{}
	for _, _sellerItem := range _seller {
		_sellerRule = append(_sellerRule, _sellerItem)
	}

	logs, sub, err := _Market.contract.FilterLogs(opts, "CreateSellOrder", _orderIdRule, _sellerRule)
	if err != nil {
		return nil, err
	}
	return &MarketCreateSellOrderIterator{contract: _Market.contract, event: "CreateSellOrder", logs: logs, sub: sub}, nil
}

// WatchCreateSellOrder is a free log subscription operation binding the contract event 0x080493e207b9f00f98a17e07dddc2ea02575db77c2769f8ef3160198b10a6910.
//
// Solidity: event CreateSellOrder(uint256 indexed _orderId, address indexed _seller, uint256 _tokenId, uint256 _price)
func (_Market *MarketFilterer) WatchCreateSellOrder(opts *bind.WatchOpts, sink chan<- *MarketCreateSellOrder, _orderId []*big.Int, _seller []common.Address) (event.Subscription, error) {

	var _orderIdRule []interface{}
	for _, _orderIdItem := range _orderId {
		_orderIdRule = append(_orderIdRule, _orderIdItem)
	}
	var _sellerRule []interface{}
	for _, _sellerItem := range _seller {
		_sellerRule = append(_sellerRule, _sellerItem)
	}

	logs, sub, err := _Market.contract.WatchLogs(opts, "CreateSellOrder", _orderIdRule, _sellerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(MarketCreateSellOrder)
				if err := _Market.contract.UnpackLog(event, "CreateSellOrder", log); err != nil {
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

// ParseCreateSellOrder is a log parse operation binding the contract event 0x080493e207b9f00f98a17e07dddc2ea02575db77c2769f8ef3160198b10a6910.
//
// Solidity: event CreateSellOrder(uint256 indexed _orderId, address indexed _seller, uint256 _tokenId, uint256 _price)
func (_Market *MarketFilterer) ParseCreateSellOrder(log types.Log) (*MarketCreateSellOrder, error) {
	event := new(MarketCreateSellOrder)
	if err := _Market.contract.UnpackLog(event, "CreateSellOrder", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
