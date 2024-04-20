package main

import (
	"fmt"
	"reflect"
	"strings"
)

var commandTypes = []reflect.Type{
	reflect.TypeOf(protowire.katdMessage_AddPeerRequest{}),
	reflect.TypeOf(protowire.katdMessage_GetConnectedPeerInfoRequest{}),
	reflect.TypeOf(protowire.katdMessage_GetPeerAddressesRequest{}),
	reflect.TypeOf(protowire.katdMessage_GetCurrentNetworkRequest{}),
	reflect.TypeOf(protowire.katdMessage_GetInfoRequest{}),

	reflect.TypeOf(protowire.katdMessage_GetBlockRequest{}),
	reflect.TypeOf(protowire.katdMessage_GetBlocksRequest{}),
	reflect.TypeOf(protowire.katdMessage_GetHeadersRequest{}),
	reflect.TypeOf(protowire.katdMessage_GetBlockCountRequest{}),
	reflect.TypeOf(protowire.katdMessage_GetBlockDagInfoRequest{}),
	reflect.TypeOf(protowire.katdMessage_GetSelectedTipHashRequest{}),
	reflect.TypeOf(protowire.katdMessage_GetVirtualSelectedParentBlueScoreRequest{}),
	reflect.TypeOf(protowire.katdMessage_GetVirtualSelectedParentChainFromBlockRequest{}),
	reflect.TypeOf(protowire.katdMessage_ResolveFinalityConflictRequest{}),
	reflect.TypeOf(protowire.katdMessage_EstimateNetworkHashesPerSecondRequest{}),

	reflect.TypeOf(protowire.katdMessage_GetBlockTemplateRequest{}),
	reflect.TypeOf(protowire.katdMessage_SubmitBlockRequest{}),

	reflect.TypeOf(protowire.katdMessage_GetMempoolEntryRequest{}),
	reflect.TypeOf(protowire.katdMessage_GetMempoolEntriesRequest{}),
	reflect.TypeOf(protowire.katdMessage_GetMempoolEntriesByAddressesRequest{}),

	reflect.TypeOf(protowire.katdMessage_SubmitTransactionRequest{}),

	reflect.TypeOf(protowire.katdMessage_GetUtxosByAddressesRequest{}),
	reflect.TypeOf(protowire.katdMessage_GetBalanceByAddressRequest{}),
	reflect.TypeOf(protowire.katdMessage_GetCoinSupplyRequest{}),

	reflect.TypeOf(protowire.katdMessage_BanRequest{}),
	reflect.TypeOf(protowire.katdMessage_UnbanRequest{}),
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
