package protowire

import (
	"github.com/Katkoin/katd/app/appmessage"
	"github.com/pkg/errors"
)

func (x *KatdMessage_Verack) toAppMessage() (appmessage.Message, error) {
	if x == nil {
		return nil, errors.Wrapf(errorNil, "katdMessage_Verack is nil")
	}
	return &appmessage.MsgVerAck{}, nil
}

func (x *KatdMessage_Verack) fromAppMessage(_ *appmessage.MsgVerAck) error {
	return nil
}
