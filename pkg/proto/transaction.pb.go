// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.21.4
// source: transaction.proto

package proto

import (
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type ACTION int32

const (
	ACTION_EMPTY                 ACTION = 0
	ACTION_STAKE                 ACTION = 1
	ACTION_UNSTAKE               ACTION = 2
	ACTION_DEPLOY_SMART_CONTRACT ACTION = 3
	ACTION_CALL_SMART_CONTRACT   ACTION = 4
	ACTION_REWARD                ACTION = 5
	ACTION_PUNISH                ACTION = 6
	ACTION_MINE                  ACTION = 7
)

// Enum value maps for ACTION.
var (
	ACTION_name = map[int32]string{
		0: "EMPTY",
		1: "STAKE",
		2: "UNSTAKE",
		3: "DEPLOY_SMART_CONTRACT",
		4: "CALL_SMART_CONTRACT",
		5: "REWARD",
		6: "PUNISH",
		7: "MINE",
	}
	ACTION_value = map[string]int32{
		"EMPTY":                 0,
		"STAKE":                 1,
		"UNSTAKE":               2,
		"DEPLOY_SMART_CONTRACT": 3,
		"CALL_SMART_CONTRACT":   4,
		"REWARD":                5,
		"PUNISH":                6,
		"MINE":                  7,
	}
)

func (x ACTION) Enum() *ACTION {
	p := new(ACTION)
	*p = x
	return p
}

func (x ACTION) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (ACTION) Descriptor() protoreflect.EnumDescriptor {
	return file_transaction_proto_enumTypes[0].Descriptor()
}

func (ACTION) Type() protoreflect.EnumType {
	return &file_transaction_proto_enumTypes[0]
}

func (x ACTION) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use ACTION.Descriptor instead.
func (ACTION) EnumDescriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{0}
}

type FEE_TYPE int32

const (
	FEE_TYPE_USER_CHARGE_FEE           FEE_TYPE = 0
	FEE_TYPE_SMART_CONTRACT_CHARGE_FEE FEE_TYPE = 1
)

// Enum value maps for FEE_TYPE.
var (
	FEE_TYPE_name = map[int32]string{
		0: "USER_CHARGE_FEE",
		1: "SMART_CONTRACT_CHARGE_FEE",
	}
	FEE_TYPE_value = map[string]int32{
		"USER_CHARGE_FEE":           0,
		"SMART_CONTRACT_CHARGE_FEE": 1,
	}
)

func (x FEE_TYPE) Enum() *FEE_TYPE {
	p := new(FEE_TYPE)
	*p = x
	return p
}

func (x FEE_TYPE) String() string {
	return protoimpl.X.EnumStringOf(x.Descriptor(), protoreflect.EnumNumber(x))
}

func (FEE_TYPE) Descriptor() protoreflect.EnumDescriptor {
	return file_transaction_proto_enumTypes[1].Descriptor()
}

func (FEE_TYPE) Type() protoreflect.EnumType {
	return &file_transaction_proto_enumTypes[1]
}

func (x FEE_TYPE) Number() protoreflect.EnumNumber {
	return protoreflect.EnumNumber(x)
}

// Deprecated: Use FEE_TYPE.Descriptor instead.
func (FEE_TYPE) EnumDescriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{1}
}

type Transaction struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash             []byte   `protobuf:"bytes,1,opt,name=Hash,proto3" json:"Hash,omitempty"`
	LastHash         []byte   `protobuf:"bytes,2,opt,name=LastHash,proto3" json:"LastHash,omitempty"`
	PublicKey        []byte   `protobuf:"bytes,3,opt,name=PublicKey,proto3" json:"PublicKey,omitempty"` // public key of sender
	ToAddress        []byte   `protobuf:"bytes,4,opt,name=ToAddress,proto3" json:"ToAddress,omitempty"`
	PendingUse       []byte   `protobuf:"bytes,5,opt,name=PendingUse,proto3" json:"PendingUse,omitempty"`
	Amount           []byte   `protobuf:"bytes,6,opt,name=Amount,proto3" json:"Amount,omitempty"` //
	MaxGas           uint64   `protobuf:"varint,7,opt,name=MaxGas,proto3" json:"MaxGas,omitempty"`
	MaxGasPrice      uint64   `protobuf:"varint,8,opt,name=MaxGasPrice,proto3" json:"MaxGasPrice,omitempty"`
	MaxTimeUse       uint64   `protobuf:"varint,9,opt,name=MaxTimeUse,proto3" json:"MaxTimeUse,omitempty"` // millisecond
	Action           ACTION   `protobuf:"varint,10,opt,name=Action,proto3,enum=transaction.ACTION" json:"Action,omitempty"`
	Data             []byte   `protobuf:"bytes,11,opt,name=Data,proto3" json:"Data,omitempty"` //
	RelatedAddresses [][]byte `protobuf:"bytes,12,rep,name=RelatedAddresses,proto3" json:"RelatedAddresses,omitempty"`
	LastDeviceKey    []byte   `protobuf:"bytes,13,opt,name=LastDeviceKey,proto3" json:"LastDeviceKey,omitempty"`   // hash last transaction deviceKey
	NewDeviceKey     []byte   `protobuf:"bytes,14,opt,name=NewDeviceKey,proto3" json:"NewDeviceKey,omitempty"`     // hash of hash for new deviceKey
	Sign             []byte   `protobuf:"bytes,15,opt,name=Sign,proto3" json:"Sign,omitempty"`                     //
	CommissionSign   []byte   `protobuf:"bytes,16,opt,name=CommissionSign,proto3" json:"CommissionSign,omitempty"` // sign of contract creator
}

func (x *Transaction) Reset() {
	*x = Transaction{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transaction) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transaction) ProtoMessage() {}

func (x *Transaction) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transaction.ProtoReflect.Descriptor instead.
func (*Transaction) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{0}
}

func (x *Transaction) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *Transaction) GetLastHash() []byte {
	if x != nil {
		return x.LastHash
	}
	return nil
}

func (x *Transaction) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *Transaction) GetToAddress() []byte {
	if x != nil {
		return x.ToAddress
	}
	return nil
}

func (x *Transaction) GetPendingUse() []byte {
	if x != nil {
		return x.PendingUse
	}
	return nil
}

func (x *Transaction) GetAmount() []byte {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *Transaction) GetMaxGas() uint64 {
	if x != nil {
		return x.MaxGas
	}
	return 0
}

func (x *Transaction) GetMaxGasPrice() uint64 {
	if x != nil {
		return x.MaxGasPrice
	}
	return 0
}

func (x *Transaction) GetMaxTimeUse() uint64 {
	if x != nil {
		return x.MaxTimeUse
	}
	return 0
}

func (x *Transaction) GetAction() ACTION {
	if x != nil {
		return x.Action
	}
	return ACTION_EMPTY
}

func (x *Transaction) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *Transaction) GetRelatedAddresses() [][]byte {
	if x != nil {
		return x.RelatedAddresses
	}
	return nil
}

func (x *Transaction) GetLastDeviceKey() []byte {
	if x != nil {
		return x.LastDeviceKey
	}
	return nil
}

func (x *Transaction) GetNewDeviceKey() []byte {
	if x != nil {
		return x.NewDeviceKey
	}
	return nil
}

func (x *Transaction) GetSign() []byte {
	if x != nil {
		return x.Sign
	}
	return nil
}

func (x *Transaction) GetCommissionSign() []byte {
	if x != nil {
		return x.CommissionSign
	}
	return nil
}

type TransactionHashData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	LastHash         []byte   `protobuf:"bytes,1,opt,name=LastHash,proto3" json:"LastHash,omitempty"`
	PublicKey        []byte   `protobuf:"bytes,2,opt,name=PublicKey,proto3" json:"PublicKey,omitempty"`
	ToAddress        []byte   `protobuf:"bytes,3,opt,name=ToAddress,proto3" json:"ToAddress,omitempty"`
	PendingUse       []byte   `protobuf:"bytes,4,opt,name=PendingUse,proto3" json:"PendingUse,omitempty"`
	Amount           []byte   `protobuf:"bytes,5,opt,name=Amount,proto3" json:"Amount,omitempty"`
	MaxGas           uint64   `protobuf:"varint,6,opt,name=MaxGas,proto3" json:"MaxGas,omitempty"`
	MaxGasPrice      uint64   `protobuf:"varint,7,opt,name=MaxGasPrice,proto3" json:"MaxGasPrice,omitempty"`
	MaxTimeUse       uint64   `protobuf:"varint,8,opt,name=MaxTimeUse,proto3" json:"MaxTimeUse,omitempty"`
	Action           ACTION   `protobuf:"varint,9,opt,name=Action,proto3,enum=transaction.ACTION" json:"Action,omitempty"`
	Data             []byte   `protobuf:"bytes,10,opt,name=Data,proto3" json:"Data,omitempty"`
	RelatedAddresses [][]byte `protobuf:"bytes,11,rep,name=RelatedAddresses,proto3" json:"RelatedAddresses,omitempty"`
	LastDeviceKey    []byte   `protobuf:"bytes,12,opt,name=LastDeviceKey,proto3" json:"LastDeviceKey,omitempty"` // hash last transaction deviceKey
	NewDeviceKey     []byte   `protobuf:"bytes,13,opt,name=NewDeviceKey,proto3" json:"NewDeviceKey,omitempty"`   // hash of hash for new deviceKey
}

func (x *TransactionHashData) Reset() {
	*x = TransactionHashData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionHashData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionHashData) ProtoMessage() {}

func (x *TransactionHashData) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionHashData.ProtoReflect.Descriptor instead.
func (*TransactionHashData) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{1}
}

func (x *TransactionHashData) GetLastHash() []byte {
	if x != nil {
		return x.LastHash
	}
	return nil
}

func (x *TransactionHashData) GetPublicKey() []byte {
	if x != nil {
		return x.PublicKey
	}
	return nil
}

func (x *TransactionHashData) GetToAddress() []byte {
	if x != nil {
		return x.ToAddress
	}
	return nil
}

func (x *TransactionHashData) GetPendingUse() []byte {
	if x != nil {
		return x.PendingUse
	}
	return nil
}

func (x *TransactionHashData) GetAmount() []byte {
	if x != nil {
		return x.Amount
	}
	return nil
}

func (x *TransactionHashData) GetMaxGas() uint64 {
	if x != nil {
		return x.MaxGas
	}
	return 0
}

func (x *TransactionHashData) GetMaxGasPrice() uint64 {
	if x != nil {
		return x.MaxGasPrice
	}
	return 0
}

func (x *TransactionHashData) GetMaxTimeUse() uint64 {
	if x != nil {
		return x.MaxTimeUse
	}
	return 0
}

func (x *TransactionHashData) GetAction() ACTION {
	if x != nil {
		return x.Action
	}
	return ACTION_EMPTY
}

func (x *TransactionHashData) GetData() []byte {
	if x != nil {
		return x.Data
	}
	return nil
}

func (x *TransactionHashData) GetRelatedAddresses() [][]byte {
	if x != nil {
		return x.RelatedAddresses
	}
	return nil
}

func (x *TransactionHashData) GetLastDeviceKey() []byte {
	if x != nil {
		return x.LastDeviceKey
	}
	return nil
}

func (x *TransactionHashData) GetNewDeviceKey() []byte {
	if x != nil {
		return x.NewDeviceKey
	}
	return nil
}

type DeployData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Code           []byte `protobuf:"bytes,1,opt,name=Code,proto3" json:"Code,omitempty"`
	StorageHost    string `protobuf:"bytes,2,opt,name=StorageHost,proto3" json:"StorageHost,omitempty"`
	StorageAddress []byte `protobuf:"bytes,3,opt,name=StorageAddress,proto3" json:"StorageAddress,omitempty"`
}

func (x *DeployData) Reset() {
	*x = DeployData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeployData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeployData) ProtoMessage() {}

func (x *DeployData) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeployData.ProtoReflect.Descriptor instead.
func (*DeployData) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{2}
}

func (x *DeployData) GetCode() []byte {
	if x != nil {
		return x.Code
	}
	return nil
}

func (x *DeployData) GetStorageHost() string {
	if x != nil {
		return x.StorageHost
	}
	return ""
}

func (x *DeployData) GetStorageAddress() []byte {
	if x != nil {
		return x.StorageAddress
	}
	return nil
}

type CallData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Input []byte `protobuf:"bytes,1,opt,name=Input,proto3" json:"Input,omitempty"`
}

func (x *CallData) Reset() {
	*x = CallData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CallData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CallData) ProtoMessage() {}

func (x *CallData) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CallData.ProtoReflect.Descriptor instead.
func (*CallData) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{3}
}

func (x *CallData) GetInput() []byte {
	if x != nil {
		return x.Input
	}
	return nil
}

type MassTransferData struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	MapAddressAmount map[string][]byte `protobuf:"bytes,1,rep,name=MapAddressAmount,proto3" json:"MapAddressAmount,omitempty" protobuf_key:"bytes,1,opt,name=key,proto3" protobuf_val:"bytes,2,opt,name=value,proto3"`
}

func (x *MassTransferData) Reset() {
	*x = MassTransferData{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *MassTransferData) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*MassTransferData) ProtoMessage() {}

func (x *MassTransferData) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use MassTransferData.ProtoReflect.Descriptor instead.
func (*MassTransferData) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{4}
}

func (x *MassTransferData) GetMapAddressAmount() map[string][]byte {
	if x != nil {
		return x.MapAddressAmount
	}
	return nil
}

type Transactions struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Transactions []*Transaction `protobuf:"bytes,1,rep,name=Transactions,proto3" json:"Transactions,omitempty"`
}

func (x *Transactions) Reset() {
	*x = Transactions{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Transactions) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Transactions) ProtoMessage() {}

func (x *Transactions) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Transactions.ProtoReflect.Descriptor instead.
func (*Transactions) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{5}
}

func (x *Transactions) GetTransactions() []*Transaction {
	if x != nil {
		return x.Transactions
	}
	return nil
}

type VerifyTransactionSignRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash   []byte `protobuf:"bytes,1,opt,name=Hash,proto3" json:"Hash,omitempty"`
	Pubkey []byte `protobuf:"bytes,2,opt,name=Pubkey,proto3" json:"Pubkey,omitempty"`
	Sign   []byte `protobuf:"bytes,3,opt,name=Sign,proto3" json:"Sign,omitempty"`
}

func (x *VerifyTransactionSignRequest) Reset() {
	*x = VerifyTransactionSignRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyTransactionSignRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyTransactionSignRequest) ProtoMessage() {}

func (x *VerifyTransactionSignRequest) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyTransactionSignRequest.ProtoReflect.Descriptor instead.
func (*VerifyTransactionSignRequest) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{6}
}

func (x *VerifyTransactionSignRequest) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *VerifyTransactionSignRequest) GetPubkey() []byte {
	if x != nil {
		return x.Pubkey
	}
	return nil
}

func (x *VerifyTransactionSignRequest) GetSign() []byte {
	if x != nil {
		return x.Sign
	}
	return nil
}

type VerifyTransactionSignResult struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Hash  []byte `protobuf:"bytes,1,opt,name=Hash,proto3" json:"Hash,omitempty"`
	Valid bool   `protobuf:"varint,2,opt,name=Valid,proto3" json:"Valid,omitempty"`
}

func (x *VerifyTransactionSignResult) Reset() {
	*x = VerifyTransactionSignResult{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *VerifyTransactionSignResult) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*VerifyTransactionSignResult) ProtoMessage() {}

func (x *VerifyTransactionSignResult) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use VerifyTransactionSignResult.ProtoReflect.Descriptor instead.
func (*VerifyTransactionSignResult) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{7}
}

func (x *VerifyTransactionSignResult) GetHash() []byte {
	if x != nil {
		return x.Hash
	}
	return nil
}

func (x *VerifyTransactionSignResult) GetValid() bool {
	if x != nil {
		return x.Valid
	}
	return false
}

type TransactionError struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	TransactionHash []byte `protobuf:"bytes,1,opt,name=TransactionHash,proto3" json:"TransactionHash,omitempty"`
	Code            int64  `protobuf:"varint,2,opt,name=Code,proto3" json:"Code,omitempty"`
	Description     string `protobuf:"bytes,3,opt,name=Description,proto3" json:"Description,omitempty"`
}

func (x *TransactionError) Reset() {
	*x = TransactionError{}
	if protoimpl.UnsafeEnabled {
		mi := &file_transaction_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *TransactionError) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*TransactionError) ProtoMessage() {}

func (x *TransactionError) ProtoReflect() protoreflect.Message {
	mi := &file_transaction_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use TransactionError.ProtoReflect.Descriptor instead.
func (*TransactionError) Descriptor() ([]byte, []int) {
	return file_transaction_proto_rawDescGZIP(), []int{8}
}

func (x *TransactionError) GetTransactionHash() []byte {
	if x != nil {
		return x.TransactionHash
	}
	return nil
}

func (x *TransactionError) GetCode() int64 {
	if x != nil {
		return x.Code
	}
	return 0
}

func (x *TransactionError) GetDescription() string {
	if x != nil {
		return x.Description
	}
	return ""
}

var File_transaction_proto protoreflect.FileDescriptor

var file_transaction_proto_rawDesc = []byte{
	0x0a, 0x11, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x0b, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x22, 0xfe, 0x03, 0x0a, 0x0b, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x48, 0x61, 0x73, 0x68,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x4c, 0x61, 0x73, 0x74, 0x48, 0x61, 0x73, 0x68,
	0x12, 0x1c, 0x0a, 0x09, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x18, 0x03, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x09, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b, 0x65, 0x79, 0x12, 0x1c,
	0x0a, 0x09, 0x54, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x18, 0x04, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x09, 0x54, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x12, 0x1e, 0x0a, 0x0a,
	0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x55, 0x73, 0x65, 0x18, 0x05, 0x20, 0x01, 0x28, 0x0c,
	0x52, 0x0a, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x55, 0x73, 0x65, 0x12, 0x16, 0x0a, 0x06,
	0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x06, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x41, 0x6d,
	0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x61, 0x78, 0x47, 0x61, 0x73, 0x18, 0x07,
	0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x4d, 0x61, 0x78, 0x47, 0x61, 0x73, 0x12, 0x20, 0x0a, 0x0b,
	0x4d, 0x61, 0x78, 0x47, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28,
	0x04, 0x52, 0x0b, 0x4d, 0x61, 0x78, 0x47, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65, 0x12, 0x1e,
	0x0a, 0x0a, 0x4d, 0x61, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x73, 0x65, 0x18, 0x09, 0x20, 0x01,
	0x28, 0x04, 0x52, 0x0a, 0x4d, 0x61, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x73, 0x65, 0x12, 0x2b,
	0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0e, 0x32, 0x13,
	0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x41, 0x43, 0x54,
	0x49, 0x4f, 0x4e, 0x52, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x12, 0x12, 0x0a, 0x04, 0x44,
	0x61, 0x74, 0x61, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x44, 0x61, 0x74, 0x61, 0x12,
	0x2a, 0x0a, 0x10, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x65, 0x73, 0x18, 0x0c, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x10, 0x52, 0x65, 0x6c, 0x61, 0x74,
	0x65, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x12, 0x24, 0x0a, 0x0d, 0x4c,
	0x61, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x0d, 0x20, 0x01,
	0x28, 0x0c, 0x52, 0x0d, 0x4c, 0x61, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4b, 0x65,
	0x79, 0x12, 0x22, 0x0a, 0x0c, 0x4e, 0x65, 0x77, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4b, 0x65,
	0x79, 0x18, 0x0e, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x4e, 0x65, 0x77, 0x44, 0x65, 0x76, 0x69,
	0x63, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x12, 0x0a, 0x04, 0x53, 0x69, 0x67, 0x6e, 0x18, 0x0f, 0x20,
	0x01, 0x28, 0x0c, 0x52, 0x04, 0x53, 0x69, 0x67, 0x6e, 0x12, 0x26, 0x0a, 0x0e, 0x43, 0x6f, 0x6d,
	0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x18, 0x10, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x69, 0x73, 0x73, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x67,
	0x6e, 0x22, 0xb6, 0x03, 0x0a, 0x13, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x48, 0x61, 0x73, 0x68, 0x44, 0x61, 0x74, 0x61, 0x12, 0x1a, 0x0a, 0x08, 0x4c, 0x61, 0x73,
	0x74, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x08, 0x4c, 0x61, 0x73,
	0x74, 0x48, 0x61, 0x73, 0x68, 0x12, 0x1c, 0x0a, 0x09, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63, 0x4b,
	0x65, 0x79, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x50, 0x75, 0x62, 0x6c, 0x69, 0x63,
	0x4b, 0x65, 0x79, 0x12, 0x1c, 0x0a, 0x09, 0x54, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x09, 0x54, 0x6f, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x12, 0x1e, 0x0a, 0x0a, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x55, 0x73, 0x65, 0x18,
	0x04, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0a, 0x50, 0x65, 0x6e, 0x64, 0x69, 0x6e, 0x67, 0x55, 0x73,
	0x65, 0x12, 0x16, 0x0a, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x05, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x06, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x16, 0x0a, 0x06, 0x4d, 0x61, 0x78,
	0x47, 0x61, 0x73, 0x18, 0x06, 0x20, 0x01, 0x28, 0x04, 0x52, 0x06, 0x4d, 0x61, 0x78, 0x47, 0x61,
	0x73, 0x12, 0x20, 0x0a, 0x0b, 0x4d, 0x61, 0x78, 0x47, 0x61, 0x73, 0x50, 0x72, 0x69, 0x63, 0x65,
	0x18, 0x07, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0b, 0x4d, 0x61, 0x78, 0x47, 0x61, 0x73, 0x50, 0x72,
	0x69, 0x63, 0x65, 0x12, 0x1e, 0x0a, 0x0a, 0x4d, 0x61, 0x78, 0x54, 0x69, 0x6d, 0x65, 0x55, 0x73,
	0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x04, 0x52, 0x0a, 0x4d, 0x61, 0x78, 0x54, 0x69, 0x6d, 0x65,
	0x55, 0x73, 0x65, 0x12, 0x2b, 0x0a, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x09, 0x20,
	0x01, 0x28, 0x0e, 0x32, 0x13, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f,
	0x6e, 0x2e, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x52, 0x06, 0x41, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x12, 0x12, 0x0a, 0x04, 0x44, 0x61, 0x74, 0x61, 0x18, 0x0a, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x44, 0x61, 0x74, 0x61, 0x12, 0x2a, 0x0a, 0x10, 0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73, 0x18, 0x0b, 0x20, 0x03, 0x28, 0x0c, 0x52, 0x10,
	0x52, 0x65, 0x6c, 0x61, 0x74, 0x65, 0x64, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x65, 0x73,
	0x12, 0x24, 0x0a, 0x0d, 0x4c, 0x61, 0x73, 0x74, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4b, 0x65,
	0x79, 0x18, 0x0c, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0d, 0x4c, 0x61, 0x73, 0x74, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x12, 0x22, 0x0a, 0x0c, 0x4e, 0x65, 0x77, 0x44, 0x65, 0x76,
	0x69, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x18, 0x0d, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0c, 0x4e, 0x65,
	0x77, 0x44, 0x65, 0x76, 0x69, 0x63, 0x65, 0x4b, 0x65, 0x79, 0x22, 0x6a, 0x0a, 0x0a, 0x44, 0x65,
	0x70, 0x6c, 0x6f, 0x79, 0x44, 0x61, 0x74, 0x61, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x48, 0x6f, 0x73, 0x74, 0x12, 0x26,
	0x0a, 0x0e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0e, 0x53, 0x74, 0x6f, 0x72, 0x61, 0x67, 0x65, 0x41,
	0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x22, 0x20, 0x0a, 0x08, 0x43, 0x61, 0x6c, 0x6c, 0x44, 0x61,
	0x74, 0x61, 0x12, 0x14, 0x0a, 0x05, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x18, 0x01, 0x20, 0x01, 0x28,
	0x0c, 0x52, 0x05, 0x49, 0x6e, 0x70, 0x75, 0x74, 0x22, 0xb8, 0x01, 0x0a, 0x10, 0x4d, 0x61, 0x73,
	0x73, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66, 0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x12, 0x5f, 0x0a,
	0x10, 0x4d, 0x61, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x41, 0x6d, 0x6f, 0x75, 0x6e,
	0x74, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x33, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61,
	0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x4d, 0x61, 0x73, 0x73, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x66,
	0x65, 0x72, 0x44, 0x61, 0x74, 0x61, 0x2e, 0x4d, 0x61, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73,
	0x73, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x52, 0x10, 0x4d, 0x61,
	0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x41, 0x6d, 0x6f, 0x75, 0x6e, 0x74, 0x1a, 0x43,
	0x0a, 0x15, 0x4d, 0x61, 0x70, 0x41, 0x64, 0x64, 0x72, 0x65, 0x73, 0x73, 0x41, 0x6d, 0x6f, 0x75,
	0x6e, 0x74, 0x45, 0x6e, 0x74, 0x72, 0x79, 0x12, 0x10, 0x0a, 0x03, 0x6b, 0x65, 0x79, 0x18, 0x01,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x03, 0x6b, 0x65, 0x79, 0x12, 0x14, 0x0a, 0x05, 0x76, 0x61, 0x6c,
	0x75, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x05, 0x76, 0x61, 0x6c, 0x75, 0x65, 0x3a,
	0x02, 0x38, 0x01, 0x22, 0x4c, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x12, 0x3c, 0x0a, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69,
	0x6f, 0x6e, 0x73, 0x18, 0x01, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x18, 0x2e, 0x74, 0x72, 0x61, 0x6e,
	0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x2e, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74,
	0x69, 0x6f, 0x6e, 0x52, 0x0c, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e,
	0x73, 0x22, 0x5e, 0x0a, 0x1c, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x71, 0x75, 0x65, 0x73,
	0x74, 0x12, 0x12, 0x0a, 0x04, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52,
	0x04, 0x48, 0x61, 0x73, 0x68, 0x12, 0x16, 0x0a, 0x06, 0x50, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x18,
	0x02, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x06, 0x50, 0x75, 0x62, 0x6b, 0x65, 0x79, 0x12, 0x12, 0x0a,
	0x04, 0x53, 0x69, 0x67, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04, 0x53, 0x69, 0x67,
	0x6e, 0x22, 0x47, 0x0a, 0x1b, 0x56, 0x65, 0x72, 0x69, 0x66, 0x79, 0x54, 0x72, 0x61, 0x6e, 0x73,
	0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x53, 0x69, 0x67, 0x6e, 0x52, 0x65, 0x73, 0x75, 0x6c, 0x74,
	0x12, 0x12, 0x0a, 0x04, 0x48, 0x61, 0x73, 0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x04,
	0x48, 0x61, 0x73, 0x68, 0x12, 0x14, 0x0a, 0x05, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x18, 0x02, 0x20,
	0x01, 0x28, 0x08, 0x52, 0x05, 0x56, 0x61, 0x6c, 0x69, 0x64, 0x22, 0x72, 0x0a, 0x10, 0x54, 0x72,
	0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x45, 0x72, 0x72, 0x6f, 0x72, 0x12, 0x28,
	0x0a, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73,
	0x68, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0c, 0x52, 0x0f, 0x54, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63,
	0x74, 0x69, 0x6f, 0x6e, 0x48, 0x61, 0x73, 0x68, 0x12, 0x12, 0x0a, 0x04, 0x43, 0x6f, 0x64, 0x65,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x04, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x20, 0x0a, 0x0b,
	0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x44, 0x65, 0x73, 0x63, 0x72, 0x69, 0x70, 0x74, 0x69, 0x6f, 0x6e, 0x2a, 0x81,
	0x01, 0x0a, 0x06, 0x41, 0x43, 0x54, 0x49, 0x4f, 0x4e, 0x12, 0x09, 0x0a, 0x05, 0x45, 0x4d, 0x50,
	0x54, 0x59, 0x10, 0x00, 0x12, 0x09, 0x0a, 0x05, 0x53, 0x54, 0x41, 0x4b, 0x45, 0x10, 0x01, 0x12,
	0x0b, 0x0a, 0x07, 0x55, 0x4e, 0x53, 0x54, 0x41, 0x4b, 0x45, 0x10, 0x02, 0x12, 0x19, 0x0a, 0x15,
	0x44, 0x45, 0x50, 0x4c, 0x4f, 0x59, 0x5f, 0x53, 0x4d, 0x41, 0x52, 0x54, 0x5f, 0x43, 0x4f, 0x4e,
	0x54, 0x52, 0x41, 0x43, 0x54, 0x10, 0x03, 0x12, 0x17, 0x0a, 0x13, 0x43, 0x41, 0x4c, 0x4c, 0x5f,
	0x53, 0x4d, 0x41, 0x52, 0x54, 0x5f, 0x43, 0x4f, 0x4e, 0x54, 0x52, 0x41, 0x43, 0x54, 0x10, 0x04,
	0x12, 0x0a, 0x0a, 0x06, 0x52, 0x45, 0x57, 0x41, 0x52, 0x44, 0x10, 0x05, 0x12, 0x0a, 0x0a, 0x06,
	0x50, 0x55, 0x4e, 0x49, 0x53, 0x48, 0x10, 0x06, 0x12, 0x08, 0x0a, 0x04, 0x4d, 0x49, 0x4e, 0x45,
	0x10, 0x07, 0x2a, 0x3e, 0x0a, 0x08, 0x46, 0x45, 0x45, 0x5f, 0x54, 0x59, 0x50, 0x45, 0x12, 0x13,
	0x0a, 0x0f, 0x55, 0x53, 0x45, 0x52, 0x5f, 0x43, 0x48, 0x41, 0x52, 0x47, 0x45, 0x5f, 0x46, 0x45,
	0x45, 0x10, 0x00, 0x12, 0x1d, 0x0a, 0x19, 0x53, 0x4d, 0x41, 0x52, 0x54, 0x5f, 0x43, 0x4f, 0x4e,
	0x54, 0x52, 0x41, 0x43, 0x54, 0x5f, 0x43, 0x48, 0x41, 0x52, 0x47, 0x45, 0x5f, 0x46, 0x45, 0x45,
	0x10, 0x01, 0x42, 0x33, 0x0a, 0x29, 0x63, 0x6f, 0x6d, 0x2e, 0x6d, 0x65, 0x74, 0x61, 0x5f, 0x6e,
	0x6f, 0x64, 0x65, 0x2e, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x73, 0x2e, 0x63, 0x6f, 0x6d, 0x70, 0x69,
	0x6c, 0x65, 0x64, 0x2e, 0x74, 0x72, 0x61, 0x6e, 0x73, 0x61, 0x63, 0x74, 0x69, 0x6f, 0x6e, 0x5a,
	0x06, 0x2f, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x62, 0x06, 0x70, 0x72, 0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_transaction_proto_rawDescOnce sync.Once
	file_transaction_proto_rawDescData = file_transaction_proto_rawDesc
)

func file_transaction_proto_rawDescGZIP() []byte {
	file_transaction_proto_rawDescOnce.Do(func() {
		file_transaction_proto_rawDescData = protoimpl.X.CompressGZIP(file_transaction_proto_rawDescData)
	})
	return file_transaction_proto_rawDescData
}

var file_transaction_proto_enumTypes = make([]protoimpl.EnumInfo, 2)
var file_transaction_proto_msgTypes = make([]protoimpl.MessageInfo, 10)
var file_transaction_proto_goTypes = []interface{}{
	(ACTION)(0),                          // 0: transaction.ACTION
	(FEE_TYPE)(0),                        // 1: transaction.FEE_TYPE
	(*Transaction)(nil),                  // 2: transaction.Transaction
	(*TransactionHashData)(nil),          // 3: transaction.TransactionHashData
	(*DeployData)(nil),                   // 4: transaction.DeployData
	(*CallData)(nil),                     // 5: transaction.CallData
	(*MassTransferData)(nil),             // 6: transaction.MassTransferData
	(*Transactions)(nil),                 // 7: transaction.Transactions
	(*VerifyTransactionSignRequest)(nil), // 8: transaction.VerifyTransactionSignRequest
	(*VerifyTransactionSignResult)(nil),  // 9: transaction.VerifyTransactionSignResult
	(*TransactionError)(nil),             // 10: transaction.TransactionError
	nil,                                  // 11: transaction.MassTransferData.MapAddressAmountEntry
}
var file_transaction_proto_depIdxs = []int32{
	0,  // 0: transaction.Transaction.Action:type_name -> transaction.ACTION
	0,  // 1: transaction.TransactionHashData.Action:type_name -> transaction.ACTION
	11, // 2: transaction.MassTransferData.MapAddressAmount:type_name -> transaction.MassTransferData.MapAddressAmountEntry
	2,  // 3: transaction.Transactions.Transactions:type_name -> transaction.Transaction
	4,  // [4:4] is the sub-list for method output_type
	4,  // [4:4] is the sub-list for method input_type
	4,  // [4:4] is the sub-list for extension type_name
	4,  // [4:4] is the sub-list for extension extendee
	0,  // [0:4] is the sub-list for field type_name
}

func init() { file_transaction_proto_init() }
func file_transaction_proto_init() {
	if File_transaction_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_transaction_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transaction); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_transaction_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionHashData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_transaction_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeployData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_transaction_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CallData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_transaction_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*MassTransferData); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_transaction_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Transactions); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_transaction_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyTransactionSignRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_transaction_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*VerifyTransactionSignResult); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_transaction_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*TransactionError); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_transaction_proto_rawDesc,
			NumEnums:      2,
			NumMessages:   10,
			NumExtensions: 0,
			NumServices:   0,
		},
		GoTypes:           file_transaction_proto_goTypes,
		DependencyIndexes: file_transaction_proto_depIdxs,
		EnumInfos:         file_transaction_proto_enumTypes,
		MessageInfos:      file_transaction_proto_msgTypes,
	}.Build()
	File_transaction_proto = out.File
	file_transaction_proto_rawDesc = nil
	file_transaction_proto_goTypes = nil
	file_transaction_proto_depIdxs = nil
}
