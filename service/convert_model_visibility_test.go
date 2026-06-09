package service

import (
	"testing"

	"github.com/QuantumNous/new-api/dto"
	relaycommon "github.com/QuantumNous/new-api/relay/common"
	"github.com/stretchr/testify/require"
)

func TestResponseOpenAI2ClaudeUsesOriginModelForMappedResponse(t *testing.T) {
	info := &relaycommon.RelayInfo{
		OriginModelName: "gpt-5.5",
		ChannelMeta: &relaycommon.ChannelMeta{
			UpstreamModelName: "gpt-5.5-free",
			IsModelMapped:     true,
		},
	}
	openAIResponse := &dto.OpenAITextResponse{
		Id:    "chatcmpl_1",
		Model: "gpt-5.5-free",
		Usage: dto.Usage{
			PromptTokens:     11,
			CompletionTokens: 7,
		},
	}

	claudeResponse := ResponseOpenAI2Claude(openAIResponse, info)

	require.Equal(t, "gpt-5.5", claudeResponse.Model)
	require.NotNil(t, claudeResponse.Usage)
	require.EqualValues(t, 11, claudeResponse.Usage.InputTokens)
	require.EqualValues(t, 7, claudeResponse.Usage.OutputTokens)
	require.Equal(t, "gpt-5.5-free", info.ChannelMeta.UpstreamModelName)
}

func TestResponseOpenAI2ClaudeKeepsFallbackModelForUnmappedResponse(t *testing.T) {
	info := &relaycommon.RelayInfo{
		OriginModelName: "gpt-5.5",
		ChannelMeta: &relaycommon.ChannelMeta{
			UpstreamModelName: "gpt-5.5",
			IsModelMapped:     false,
		},
	}
	openAIResponse := &dto.OpenAITextResponse{
		Id:    "chatcmpl_1",
		Model: "provider-returned-model",
	}

	claudeResponse := ResponseOpenAI2Claude(openAIResponse, info)

	require.Equal(t, "provider-returned-model", claudeResponse.Model)
}

func TestStreamResponseOpenAI2ClaudeUsesOriginModelForMappedMessageStart(t *testing.T) {
	info := &relaycommon.RelayInfo{
		OriginModelName:   "gpt-5.5",
		SendResponseCount: 1,
		ClaudeConvertInfo: &relaycommon.ClaudeConvertInfo{},
		ChannelMeta: &relaycommon.ChannelMeta{
			UpstreamModelName: "gpt-5.5-free",
			IsModelMapped:     true,
		},
	}
	openAIResponse := &dto.ChatCompletionsStreamResponse{
		Id:    "chatcmpl_1",
		Model: "gpt-5.5-free",
	}

	claudeResponses := StreamResponseOpenAI2Claude(openAIResponse, info)

	require.NotEmpty(t, claudeResponses)
	require.Equal(t, "message_start", claudeResponses[0].Type)
	require.NotNil(t, claudeResponses[0].Message)
	require.Equal(t, "gpt-5.5", claudeResponses[0].Message.Model)
	require.Equal(t, "gpt-5.5-free", info.ChannelMeta.UpstreamModelName)
}

func TestStreamResponseOpenAI2ClaudeKeepsFallbackModelForUnmappedMessageStart(t *testing.T) {
	info := &relaycommon.RelayInfo{
		OriginModelName:   "gpt-5.5",
		SendResponseCount: 1,
		ClaudeConvertInfo: &relaycommon.ClaudeConvertInfo{},
		ChannelMeta: &relaycommon.ChannelMeta{
			UpstreamModelName: "gpt-5.5",
			IsModelMapped:     false,
		},
	}
	openAIResponse := &dto.ChatCompletionsStreamResponse{
		Id:    "chatcmpl_1",
		Model: "provider-returned-model",
	}

	claudeResponses := StreamResponseOpenAI2Claude(openAIResponse, info)

	require.NotEmpty(t, claudeResponses)
	require.NotNil(t, claudeResponses[0].Message)
	require.Equal(t, "provider-returned-model", claudeResponses[0].Message.Model)
}
