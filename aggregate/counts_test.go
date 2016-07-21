package aggregate

import (
	"testing"

	"github.com/gotascii/simpledb/data"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/suite"
)

type CountsSuite struct {
	suite.Suite
	DB     *data.MockTreap
	Counts *Counts
	AnyInt mock.AnythingOfTypeArgument
	AnyKV  mock.AnythingOfTypeArgument
}

func (suite *CountsSuite) SetupTest() {
	suite.DB = &data.MockTreap{}
	suite.Counts = &Counts{DB: suite.DB}
	suite.AnyInt = mock.AnythingOfType("int")
	suite.AnyKV = mock.AnythingOfType("*data.KV")
}

func (suite *CountsSuite) TestInitCountOnAppendIfCountIsNil() {
	suite.DB.On("Get", &data.KV{K: "1"}).Return(nil)

	expected := &data.MockTreap{}
	suite.DB.On("Upsert", &data.KV{"1", 1}, suite.AnyInt).Return(expected)

	suite.Counts.Append("1", 1)

	suite.Equal(expected, suite.Counts.DB)
}

func (suite *CountsSuite) TestIncrementCountOnAppendIfCountIsNotNil() {
	suite.DB.On("Get", &data.KV{K: "1"}).Return(&data.KV{"1", 1})

	expected := &data.MockTreap{}
	suite.DB.On("Upsert", &data.KV{"1", 2}, suite.AnyInt).Return(expected)

	suite.Counts.Append("1", 1)

	suite.Equal(expected, suite.Counts.DB)
}

func (suite *CountsSuite) TestDeleteCountOnRemoveIfCountWouldBeReducedToZero() {
	suite.DB.On("Get", &data.KV{K: "1"}).Return(&data.KV{"1", 1})
	expected := &data.MockTreap{}
	suite.DB.On("Delete", &data.KV{"1", 1}).Return(expected)

	suite.Counts.Remove("1", 1)

	suite.Equal(expected, suite.Counts.DB)
}

func (suite *CountsSuite) TestDecrementCountOnRemoveIfCountWouldNotBeReducedToZero() {
	suite.DB.On("Get", &data.KV{K: "1"}).Return(&data.KV{"1", 2})
	expected := &data.MockTreap{}
	suite.DB.On("Upsert", &data.KV{"1", 1}, suite.AnyInt).Return(expected)

	suite.Counts.Remove("1", 1)

	suite.Equal(expected, suite.Counts.DB)
}

func (suite *CountsSuite) TestReturnZeroIfCountIsNil() {
	suite.DB.On("Get", &data.KV{K: "1"}).Return(nil)

	suite.Equal(0, suite.Counts.Compute("1"))
}

func (suite *CountsSuite) TestReturnCount() {
	suite.DB.On("Get", &data.KV{K: "1"}).Return(&data.KV{"1", 1})

	suite.Equal(1, suite.Counts.Compute("1"))
}

func TestCountsSuite(t *testing.T) {
	suite.Run(t, new(CountsSuite))
}
