package main

import (
	"fmt"
	"reflect"
	"strings"

	"github.com/Katkoin/katd/infrastructure/network/netadapter/server/grpcserver/protowire"
)

var commandTypes = []reflect.Type{
	reflect.TypeOf(protowire.KatdMessage_AddPeerRequest{}),
	reflect.TypeOf(protowire.KatdMessage_GetConnectedPeerInfoRequest{}),
	reflect.TypeOf(protowire.KatdMessage_GetPeerAddressesRequest{}),
	reflect.TypeOf(protowire.KatdMessage_GetCurrentNetworkRequest{}),
	reflect.TypeOf(protowire.KatdMessage_GetInfoRequest{}),

	reflect.TypeOf(protowire.KatdMessage_GetBlockRequest{}),
	reflect.TypeOf(protowire.KatdMessage_GetBlocksRequest{}),
	reflect.TypeOf(protowire.KatdMessage_GetHeadersRequest{}),
	reflect.TypeOf(protowire.KatdMessage_GetBlockCountRequest{}),
	reflect.TypeOf(protowire.KatdMessage_GetBlockDagInfoRequest{}),
	reflect.TypeOf(protowire.KatdMessage_GetSelectedTipHashRequest{}),
	reflect.TypeOf(protowire.KatdMessage_GetVirtualSelectedParentBlueScoreRequest{}),
	reflect.TypeOf(protowire.KatdMessage_GetVirtualSelectedParentChainFromBlockRequest{}),
	reflect.TypeOf(protowire.KatdMessage_ResolveFinalityConflictRequest{}),
	reflect.TypeOf(protowire.KatdMessage_EstimateNetworkHashesPerSecondRequest{}),

	reflect.TypeOf(protowire.KatdMessage_GetBlockTemplateRequest{}),
	reflect.TypeOf(protowire.KatdMessage_SubmitBlockRequest{}),

	reflect.TypeOf(protowire.KatdMessage_GetMempoolEntryRequest{}),
	reflect.TypeOf(protowire.KatdMessage_GetMempoolEntriesRequest{}),
	reflect.TypeOf(protowire.KatdMessage_GetMempoolEntriesByAddressesRequest{}),

	reflect.TypeOf(protowire.KatdMessage_SubmitTransactionRequest{}),

	reflect.TypeOf(protowire.KatdMessage_GetUtxosByAddressesRequest{}),
	reflect.TypeOf(protowire.KatdMessage_GetBalanceByAddressRequest{}),
	reflect.TypeOf(protowire.KatdMessage_GetCoinSupplyRequest{}),

	reflect.TypeOf(protowire.KatdMessage_BanRequest{}),
	reflect.TypeOf(protowire.KatdMessage_UnbanRequest{}),
}

type commandDescription struct {
	name       string
	parameters []*parameterDescription
	typeof     reflect.Type
}

type parameterDescription struct {
	name   string
	typeof reflect.Type
}

func commandDescriptions() []*commandDescription {
	commandDescriptions := make([]*commandDescription, len(commandTypes))

	for i, commandTypeWrapped := range commandTypes {
		commandType := unwrapCommandType(commandTypeWrapped)

		name := strings.TrimSuffix(commandType.Name(), "RequestMessage")
		numFields := commandType.NumField()

		var parameters []*parameterDescription
		for i := 0; i < numFields; i++ {
			field := commandType.Field(i)

			if !isFieldExported(field) {
				continue
			}

			parameters = append(parameters, &parameterDescription{
				name:   field.Name,
				typeof: field.Type,
			})
		}
		commandDescriptions[i] = &commandDescription{
			name:       name,
			parameters: parameters,
			typeof:     commandTypeWrapped,
		}
	}

	return commandDescriptions
}

func (cd *commandDescription) help() string {
	sb := &strings.Builder{}
	sb.WriteString(cd.name)
	for _, parameter := range cd.parameters {
		_, _ = fmt.Fprintf(sb, " [%s]", parameter.name)
	}
	return sb.String()
}
