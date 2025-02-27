package core

import (
	"strconv"
	"strings"
	"testing"

	"github.com/gogo/protobuf/jsonpb"
	"github.com/regen-network/gocuke"
	"github.com/stretchr/testify/require"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type msgCreateClass struct {
	t   gocuke.TestingT
	msg *MsgCreateClass
	err error
}

func TestMsgCreateClass(t *testing.T) {
	gocuke.NewRunner(t, &msgCreateClass{}).Path("./features/msg_create_class.feature").Run()
}

func (s *msgCreateClass) Before(t gocuke.TestingT) {
	s.t = t

	// TODO: move to init function in the root directory of the module #1243
	cfg := sdk.GetConfig()
	cfg.SetBech32PrefixForAccount("regen", "regenpub")
}

func (s *msgCreateClass) TheMessage(a gocuke.DocString) {
	s.msg = &MsgCreateClass{}
	err := jsonpb.UnmarshalString(a.Content, s.msg)
	require.NoError(s.t, err)
}

func (s *msgCreateClass) MetadataWithLength(a string) {
	length, err := strconv.ParseInt(a, 10, 64)
	require.NoError(s.t, err)

	s.msg.Metadata = strings.Repeat("x", int(length))
}

func (s *msgCreateClass) TheMessageIsValidated() {
	s.err = s.msg.ValidateBasic()
}

func (s *msgCreateClass) ExpectTheError(a string) {
	require.EqualError(s.t, s.err, a)
}

func (s *msgCreateClass) ExpectNoError() {
	require.NoError(s.t, s.err)
}
