package protowire

import (
	"github.com/Katkoin/katd/app/appmessage"
	"github.com/pkg/errors"
)

func (x *katdMessage_Ready) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "katdMessage_Ready is nil")
	}
	return &appmessage.MsgReady{}, nil
}

func (x *katdMessage_Ready) fromAppMessage(_ *appmessage.MsgReady) error {
	return nil
}
