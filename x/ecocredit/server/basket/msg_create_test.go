package basket_test

import (
	"strconv"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
	bank "github.com/cosmos/cosmos-sdk/x/bank/types"

	api "github.com/regen-network/regen-ledger/api/regen/ecocredit/basket/v1"
	coreapi "github.com/regen-network/regen-ledger/api/regen/ecocredit/v1"
	"github.com/regen-network/regen-ledger/x/ecocredit/basket"
	"github.com/regen-network/regen-ledger/x/ecocredit/core"
)

type createSuite struct {
	*baseSuite
	alice               sdk.AccAddress
	aliceBalance        sdk.Coin
	params              core.Params
	basketName          string
	creditTypeAbbrev    string
	creditTypePrecision uint32
	res                 *basket.MsgCreateResponse
	err                 error
}

func TestCreate(t *testing.T) {
	gocuke.NewRunner(t, &createSuite{}).Path("./features/msg_create.feature").Run()
}

func (s *createSuite) Before(t gocuke.TestingT) {
	s.baseSuite = setupBase(t)
	s.alice = s.addrs[0]
	s.basketName = "NCT"
	s.creditTypeAbbrev = "C"
	s.creditTypePrecision = 6
}

func (s *createSuite) AllowedBasketFee(a string) {
	basketFee, err := sdk.ParseCoinsNormalized(a)
	require.NoError(s.t, err)

	s.params.BasketFee = basketFee
}

func (s *createSuite) ACreditType() {
	err := s.coreStore.CreditTypeTable().Insert(s.ctx, &coreapi.CreditType{
		Abbreviation: s.creditTypeAbbrev,
		Precision:    s.creditTypePrecision,
	})
	require.NoError(s.t, err)
}

func (s *createSuite) ACreditTypeWithAbbreviation(a string) {
	err := s.coreStore.CreditTypeTable().Insert(s.ctx, &coreapi.CreditType{
		Abbreviation: a,
		Precision:    s.creditTypePrecision,
	})
	require.NoError(s.t, err)
}

func (s *createSuite) ACreditTypeWithPrecision(b string) {
	precision, err := strconv.ParseUint(b, 10, 32)
	require.NoError(s.t, err)

	s.creditTypePrecision = uint32(precision)

	err = s.coreStore.CreditTypeTable().Insert(s.ctx, &coreapi.CreditType{
		Abbreviation: s.creditTypeAbbrev,
		Precision:    s.creditTypePrecision,
	})
	require.NoError(s.t, err)
}

func (s *createSuite) ACreditTypeWithAbbreviationAndPrecision(a string, b string) {
	precision, err := strconv.ParseUint(b, 10, 32)
	require.NoError(s.t, err)

	s.creditTypePrecision = uint32(precision)

	err = s.coreStore.CreditTypeTable().Insert(s.ctx, &coreapi.CreditType{
		Abbreviation: a,
		Precision:    s.creditTypePrecision,
	})
	require.NoError(s.t, err)
}

func (s *createSuite) ACreditClassWithId(a string) {
	creditTypeAbbrev := core.GetCreditTypeAbbrevFromClassId(a)

	err := s.coreStore.ClassTable().Insert(s.ctx, &coreapi.Class{
		Id:               a,
		CreditTypeAbbrev: creditTypeAbbrev,
	})
	require.NoError(s.t, err)
}

func (s *createSuite) ABasketWithName(a string) {
	err := s.stateStore.BasketTable().Insert(s.ctx, &api.Basket{
		Name: a,
	})
	require.NoError(s.t, err)
}

func (s *createSuite) AliceHasATokenBalance(a string) {
	coin, err := sdk.ParseCoinNormalized(a)
	require.NoError(s.t, err)

	s.aliceBalance = coin
}

func (s *createSuite) AliceAttemptsToCreateABasketWithFee(a string) {
	coin, err := sdk.ParseCoinNormalized(a)
	require.NoError(s.t, err)

	basketFee := sdk.NewCoins(coin)

	s.createExpectCalls()

	s.res, s.err = s.k.Create(s.ctx, &basket.MsgCreate{
		Curator:          s.alice.String(),
		Name:             s.basketName,
		Fee:              basketFee,
		CreditTypeAbbrev: s.creditTypeAbbrev,
	})
}

func (s *createSuite) AliceAttemptsToCreateABasketWithNoFee() {
	s.createExpectCalls()

	s.res, s.err = s.k.Create(s.ctx, &basket.MsgCreate{
		Curator:          s.alice.String(),
		Name:             s.basketName,
		CreditTypeAbbrev: s.creditTypeAbbrev,
	})
}

func (s *createSuite) AliceAttemptsToCreateABasketWithCreditType(a string) {
	s.createExpectCalls()

	s.res, s.err = s.k.Create(s.ctx, &basket.MsgCreate{
		Curator:          s.alice.String(),
		Name:             s.basketName,
		CreditTypeAbbrev: a,
	})
}

func (s *createSuite) AliceAttemptsToCreateABasketWithCreditTypeAndAllowedClass(a string, b string) {
	s.createExpectCalls()

	s.res, s.err = s.k.Create(s.ctx, &basket.MsgCreate{
		Curator:          s.alice.String(),
		Name:             s.basketName,
		CreditTypeAbbrev: a,
		AllowedClasses:   []string{b},
	})
}

func (s *createSuite) AliceAttemptsToCreateABasketWithAllowedClass(a string) {
	s.createExpectCalls()

	s.res, s.err = s.k.Create(s.ctx, &basket.MsgCreate{
		Curator:          s.alice.String(),
		Name:             s.basketName,
		CreditTypeAbbrev: s.creditTypeAbbrev,
		AllowedClasses:   []string{a},
	})
}

func (s *createSuite) AliceAttemptsToCreateABasketWithName(a string) {
	s.createExpectCalls()

	s.res, s.err = s.k.Create(s.ctx, &basket.MsgCreate{
		Curator:          s.alice.String(),
		Name:             a,
		CreditTypeAbbrev: s.creditTypeAbbrev,
	})
}

func (s *createSuite) AliceAttemptsToCreateABasketWithNameAndCreditType(a string, b string) {
	s.createExpectCalls()

	s.res, s.err = s.k.Create(s.ctx, &basket.MsgCreate{
		Curator:          s.alice.String(),
		Name:             a,
		CreditTypeAbbrev: b,
	})
}

func (s *createSuite) ExpectNoError() {
	require.NoError(s.t, s.err)
}

func (s *createSuite) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *createSuite) ExpectAliceTokenBalance(a string) {
	coin, err := sdk.ParseCoinNormalized(a)
	require.NoError(s.t, err)

	require.Equal(s.t, coin, s.aliceBalance)
}

func (s *createSuite) ExpectTheResponse(a gocuke.DocString) {
	res := &basket.MsgCreateResponse{}
	err := jsonpb.UnmarshalString(a.Content, res)
	require.NoError(s.t, err)

	require.Equal(s.t, res, s.res)
}

func (s *createSuite) createExpectCalls() {
	var basketFee sdk.Coins

	s.paramsKeeper.EXPECT().
		Get(s.sdkCtx, core.KeyBasketFee, &basketFee).
		Do(func(ctx sdk.Context, key []byte, basketFee *sdk.Coins) {
			*basketFee = s.params.BasketFee
		}).
		AnyTimes() // not expected on failed attempt

	var expectedFee sdk.Coin
	var expectedFees sdk.Coins

	if len(s.params.BasketFee) == 1 {
		expectedFee = s.params.BasketFee[0]
		expectedFees = sdk.Coins{expectedFee}
	}

	if len(s.params.BasketFee) == 2 {
		expectedFee = s.params.BasketFee[1]
		expectedFees = sdk.Coins{expectedFee}
	}

	s.bankKeeper.EXPECT().
		GetBalance(s.sdkCtx, s.alice, expectedFee.Denom).
		Return(s.aliceBalance).
		AnyTimes() // not expected on failed attempt

	s.bankKeeper.EXPECT().
		SendCoinsFromAccountToModule(s.sdkCtx, s.alice, basket.BasketSubModuleName, expectedFees).
		Do(func(sdk.Context, sdk.AccAddress, string, sdk.Coins) {
			if s.params.BasketFee != nil {
				// simulate token balance update unavailable with mocks
				s.aliceBalance = s.aliceBalance.Sub(expectedFee)
			}
		}).
		Return(nil).
		AnyTimes() // not expected on failed attempt

	s.bankKeeper.EXPECT().
		BurnCoins(s.sdkCtx, basket.BasketSubModuleName, expectedFees).
		Return(nil).
		AnyTimes() // not expected on failed attempt

	s.bankKeeper.EXPECT().
		SetDenomMetaData(s.sdkCtx, s.getDenomMetadata()).
		AnyTimes() // not expected on failed attempt
}

func (s *createSuite) getDenomMetadata() bank.Metadata {
	denom, displayDenom, err := basket.FormatBasketDenom(s.basketName, s.creditTypeAbbrev, s.creditTypePrecision)
	require.NoError(s.t, err)

	denomUnits := make([]*bank.DenomUnit, 0)
	if denom != displayDenom {
		denomUnits = append(denomUnits, &bank.DenomUnit{
			Denom: denom,
		})
	}
	denomUnits = append(denomUnits, &bank.DenomUnit{
		Denom:    displayDenom,
		Exponent: s.creditTypePrecision,
	})

	return bank.Metadata{
		Name:       s.basketName,
		Display:    displayDenom,
		Base:       denom,
		Symbol:     s.basketName,
		DenomUnits: denomUnits,
	}
}
