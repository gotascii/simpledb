package kvdb

import (
	"testing"

	amock "github.com/gotascii/simpledb/aggregate/mock"
	"github.com/gotascii/simpledb/data"
	dmock "github.com/gotascii/simpledb/data/mock"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type TransactionSuite struct {
	suite.Suite
	DB     *dmock.Treap
	Counts *amock.Counts
	Tact   *Transaction
	AnyInt mock.AnythingOfTypeArgument
	AnyKV  mock.AnythingOfTypeArgument
}

func (suite *TransactionSuite) SetupTest() {
	suite.DB = &dmock.Treap{}
	suite.Counts = &amock.Counts{}
	suite.Tact = &Transaction{DB: suite.DB, Counts: suite.Counts}
	suite.AnyInt = mock.AnythingOfType("int")
	suite.AnyKV = mock.AnythingOfType("*data.KV")
}

func (suite *TransactionSuite) TestCopy() {
	suite.Counts.On("Copy").Return(suite.Counts)
	t := suite.Tact.Copy()

	suite.Equal(suite.DB, t.DB)
	suite.Equal(suite.Counts, t.Counts)
}

func (suite *TransactionSuite) TestGetExistingKV() {
	suite.DB.On("Get", &data.KV{K: "A"}).Return(&data.KV{"A", "13"})

	suite.Equal("13", suite.Tact.Get("A"))
}

func (suite *TransactionSuite) TestGetNilKV() {
	suite.DB.On("Get", &data.KV{K: "A"}).Return(nil)

	suite.Nil(suite.Tact.Get("A"))
}

func (suite *TransactionSuite) TestUnsetNilKV() {
	suite.DB.On("Get", &data.KV{K: "A"}).Return(nil)

	suite.Tact.Unset("A")
}

func (suite *TransactionSuite) TestDeleteFromDBOnUnsetKV() {
	suite.Counts.On("Remove", "1", 1)
	suite.DB.On("Get", &data.KV{K: "A"}).Return(&data.KV{"A", "1"})

	expected := &dmock.Treap{}
	suite.DB.On("Delete", &data.KV{"A", "1"}).Return(expected)

	suite.Tact.Unset("A")

	suite.Equal(expected, suite.Tact.DB)
}

func (suite *TransactionSuite) TestRemoveFromCountsOnUnsetKV() {
	suite.DB.On("Delete", &data.KV{"A", "1"}).Return(suite.DB)
	suite.DB.On("Get", &data.KV{K: "A"}).Return(&data.KV{"A", "1"})

	suite.Counts.On("Remove", "1", 1)

	suite.Tact.Unset("A")

	suite.Counts.AssertCalled(suite.T(), "Remove", "1", 1)
}

func (suite *TransactionSuite) TestNoUpsertOnSetIfKVUnchanged() {
	suite.DB.On("Get", &data.KV{K: "A"}).Return(&data.KV{"A", "1"})

	suite.Tact.Set("A", "1")

	suite.Equal(suite.DB, suite.Tact.DB)
}

func (suite *TransactionSuite) TestUpsertDBOnSetIfKVIsNil() {
	// Stub Counts Append.
	suite.Counts.On("Append", "1", 1)

	suite.DB.On("Get", &data.KV{K: "A"}).Return(nil)
	expected := &dmock.Treap{}
	suite.DB.On("Upsert", &data.KV{"A", "1"}, suite.AnyInt).Return(expected)

	suite.Tact.Set("A", "1")

	suite.Equal(expected, suite.Tact.DB)
}

func (suite *TransactionSuite) TestUpsertDBOnSetIfKVIsUpdated() {
	// Stub Counts.
	suite.Counts.On("Append", "2", 1)
	suite.Counts.On("Remove", "1", 1)

	suite.DB.On("Get", &data.KV{K: "A"}).Return(&data.KV{"A", "1"})
	expected := &dmock.Treap{}
	suite.DB.On("Upsert", &data.KV{"A", "2"}, suite.AnyInt).Return(expected)

	suite.Tact.Set("A", "2")

	suite.Equal(expected, suite.Tact.DB)
}

func (suite *TransactionSuite) TestAppendCountsOnSetIfKVIsNil() {
	// Stub DB Upsert.
	suite.DB.On("Upsert", &data.KV{"A", "1"}, suite.AnyInt).Return(suite.DB)

	suite.DB.On("Get", &data.KV{K: "A"}).Return(nil)
	suite.Counts.On("Append", "1", 1)

	suite.Tact.Set("A", "1")
	suite.Counts.AssertCalled(suite.T(), "Append", "1", 1)
}

func (suite *TransactionSuite) TestAppendCountsOnSetIfKVIsUpdated() {
	// Stub DB Upsert.
	suite.DB.On("Upsert", &data.KV{"A", "2"}, suite.AnyInt).Return(suite.DB)
	// Stub Counts Remove
	suite.Counts.On("Remove", "1", 1)

	suite.DB.On("Get", &data.KV{K: "A"}).Return(&data.KV{"A", "1"})
	suite.Counts.On("Append", "2", 1)

	suite.Tact.Set("A", "2")
	suite.Counts.AssertCalled(suite.T(), "Append", "2", 1)
}

func (suite *TransactionSuite) TestRemoveOldCountsOnSetIfKVIsUpdated() {
	// Stub DB Upsert.
	suite.DB.On("Upsert", &data.KV{"A", "2"}, suite.AnyInt).Return(suite.DB)
	// Stub Counts Append
	suite.Counts.On("Append", "2", 1)

	suite.DB.On("Get", &data.KV{K: "A"}).Return(&data.KV{"A", "1"})
	suite.Counts.On("Remove", "1", 1)

	suite.Tact.Set("A", "2")
	suite.Counts.AssertCalled(suite.T(), "Remove", "1", 1)
}

func (suite *TransactionSuite) TestNumequaltoDelegatesToCounts() {
	suite.Counts.On("Compute", "1").Return(1)

	actual := suite.Tact.Numequalto("1")

	suite.Equal(1, actual)
}

func TestTransactionSuite(t *testing.T) {
	suite.Run(t, new(TransactionSuite))
}
