/**
 * (C) Copyright IBM Corp. 2022.
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

/*
 * IBM OpenAPI SDK Code Generator Version: 3.46.0-a4e29da0-20220224-210428
 */

// Package texttospeechv1 : Operations and models for the TextToSpeechV1 service
package texttospeechv1

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"reflect"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	common "github.com/watson-developer-cloud/go-sdk/v3/common"
)

// TextToSpeechV1 : The IBM Watson&trade; Text to Speech service provides APIs that use IBM's speech-synthesis
// capabilities to synthesize text into natural-sounding speech in a variety of languages, dialects, and voices. The
// service supports at least one male or female voice, sometimes both, for each language. The audio is streamed back to
// the client with minimal delay.
//
// For speech synthesis, the service supports a synchronous HTTP Representational State Transfer (REST) interface and a
// WebSocket interface. Both interfaces support plain text and SSML input. SSML is an XML-based markup language that
// provides text annotation for speech-synthesis applications. The WebSocket interface also supports the SSML
// <code>&lt;mark&gt;</code> element and word timings.
//
// The service offers a customization interface that you can use to define sounds-like or phonetic translations for
// words. A sounds-like translation consists of one or more words that, when combined, sound like the word. A phonetic
// translation is based on the SSML phoneme format for representing a word. You can specify a phonetic translation in
// standard International Phonetic Alphabet (IPA) representation or in the proprietary IBM Symbolic Phonetic
// Representation (SPR). For phonetic translation, the Arabic, Chinese, Dutch, Australian English, Korean, and Swedish
// voices support only IPA, not SPR.
//
// The service also offers a Tune by Example feature that lets you define custom prompts. You can also define speaker
// models to improve the quality of your custom prompts. The service support custom prompts only for US English custom
// models and voices.
//
// API Version: 1.0.0
// See: https://cloud.ibm.com/docs/text-to-speech
type TextToSpeechV1 struct {
	Service *core.BaseService
}

// DefaultServiceURL is the default URL to make service requests to.
const DefaultServiceURL = "https://api.us-south.text-to-speech.watson.cloud.ibm.com"

// DefaultServiceName is the default key used to find external configuration information.
const DefaultServiceName = "text_to_speech"

// TextToSpeechV1Options : Service options
type TextToSpeechV1Options struct {
	ServiceName   string
	URL           string
	Authenticator core.Authenticator
}

// NewTextToSpeechV1 : constructs an instance of TextToSpeechV1 with passed in options.
func NewTextToSpeechV1(options *TextToSpeechV1Options) (service *TextToSpeechV1, err error) {
	if options.ServiceName == "" {
		options.ServiceName = DefaultServiceName
	}

	serviceOptions := &core.ServiceOptions{
		URL:           DefaultServiceURL,
		Authenticator: options.Authenticator,
	}

	if serviceOptions.Authenticator == nil {
		serviceOptions.Authenticator, err = core.GetAuthenticatorFromEnvironment(options.ServiceName)
		if err != nil {
			return
		}
	}

	baseService, err := core.NewBaseService(serviceOptions)
	if err != nil {
		return
	}

	err = baseService.ConfigureService(options.ServiceName)
	if err != nil {
		return
	}

	if options.URL != "" {
		err = baseService.SetServiceURL(options.URL)
		if err != nil {
			return
		}
	}

	service = &TextToSpeechV1{
		Service: baseService,
	}

	return
}

// GetServiceURLForRegion returns the service URL to be used for the specified region
func GetServiceURLForRegion(region string) (string, error) {
	return "", fmt.Errorf("service does not support regional URLs")
}

// Clone makes a copy of "textToSpeech" suitable for processing requests.
func (textToSpeech *TextToSpeechV1) Clone() *TextToSpeechV1 {
	if core.IsNil(textToSpeech) {
		return nil
	}
	clone := *textToSpeech
	clone.Service = textToSpeech.Service.Clone()
	return &clone
}

// SetServiceURL sets the service URL
func (textToSpeech *TextToSpeechV1) SetServiceURL(url string) error {
	return textToSpeech.Service.SetServiceURL(url)
}

// GetServiceURL returns the service URL
func (textToSpeech *TextToSpeechV1) GetServiceURL() string {
	return textToSpeech.Service.GetServiceURL()
}

// SetDefaultHeaders sets HTTP headers to be sent in every request
func (textToSpeech *TextToSpeechV1) SetDefaultHeaders(headers http.Header) {
	textToSpeech.Service.SetDefaultHeaders(headers)
}

// SetEnableGzipCompression sets the service's EnableGzipCompression field
func (textToSpeech *TextToSpeechV1) SetEnableGzipCompression(enableGzip bool) {
	textToSpeech.Service.SetEnableGzipCompression(enableGzip)
}

// GetEnableGzipCompression returns the service's EnableGzipCompression field
func (textToSpeech *TextToSpeechV1) GetEnableGzipCompression() bool {
	return textToSpeech.Service.GetEnableGzipCompression()
}

// EnableRetries enables automatic retries for requests invoked for this service instance.
// If either parameter is specified as 0, then a default value is used instead.
func (textToSpeech *TextToSpeechV1) EnableRetries(maxRetries int, maxRetryInterval time.Duration) {
	textToSpeech.Service.EnableRetries(maxRetries, maxRetryInterval)
}

// DisableRetries disables automatic retries for requests invoked for this service instance.
func (textToSpeech *TextToSpeechV1) DisableRetries() {
	textToSpeech.Service.DisableRetries()
}

// DisableSSLVerification bypasses verification of the server's SSL certificate
func (textToSpeech *TextToSpeechV1) DisableSSLVerification() {
	textToSpeech.Service.DisableSSLVerification()
}

// ListVoices : List voices
// Lists all voices available for use with the service. The information includes the name, language, gender, and other
// details about the voice. The ordering of the list of voices can change from call to call; do not rely on an
// alphabetized or static list of voices. To see information about a specific voice, use the [Get a voice](#getvoice).
//
// **See also:** [Listing all available
// voices](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-voices#listVoices).
func (textToSpeech *TextToSpeechV1) ListVoices(listVoicesOptions *ListVoicesOptions) (result *Voices, response *core.DetailedResponse, err error) {
	return textToSpeech.ListVoicesWithContext(context.Background(), listVoicesOptions)
}

// ListVoicesWithContext is an alternate form of the ListVoices method which supports a Context parameter
func (textToSpeech *TextToSpeechV1) ListVoicesWithContext(ctx context.Context, listVoicesOptions *ListVoicesOptions) (result *Voices, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listVoicesOptions, "listVoicesOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = textToSpeech.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(textToSpeech.Service.Options.URL, `/v1/voices`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listVoicesOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "ListVoices")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = textToSpeech.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalVoices)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetVoice : Get a voice
// Gets information about the specified voice. The information includes the name, language, gender, and other details
// about the voice. Specify a customization ID to obtain information for a custom model that is defined for the language
// of the specified voice. To list information about all available voices, use the [List voices](#listvoices) method.
//
// **See also:** [Listing a specific
// voice](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-voices#listVoice).
//
// **Note:** The Arabic, Chinese, Czech, Dutch (Belgian and Netherlands), Australian English, Korean, and Swedish
// languages and voices are supported only for IBM Cloud; they are deprecated for IBM Cloud Pak for Data. Also, the
// `ar-AR_OmarVoice` voice is deprecated; use the `ar-MS_OmarVoice` voice instead.
func (textToSpeech *TextToSpeechV1) GetVoice(getVoiceOptions *GetVoiceOptions) (result *Voice, response *core.DetailedResponse, err error) {
	return textToSpeech.GetVoiceWithContext(context.Background(), getVoiceOptions)
}

// GetVoiceWithContext is an alternate form of the GetVoice method which supports a Context parameter
func (textToSpeech *TextToSpeechV1) GetVoiceWithContext(ctx context.Context, getVoiceOptions *GetVoiceOptions) (result *Voice, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getVoiceOptions, "getVoiceOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getVoiceOptions, "getVoiceOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"voice": *getVoiceOptions.Voice,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = textToSpeech.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(textToSpeech.Service.Options.URL, `/v1/voices/{voice}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getVoiceOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "GetVoice")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if getVoiceOptions.CustomizationID != nil {
		builder.AddQuery("customization_id", fmt.Sprint(*getVoiceOptions.CustomizationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = textToSpeech.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalVoice)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// Synthesize : Synthesize audio
// Synthesizes text to audio that is spoken in the specified voice. The service bases its understanding of the language
// for the input text on the specified voice. Use a voice that matches the language of the input text.
//
// The method accepts a maximum of 5 KB of input text in the body of the request, and 8 KB for the URL and headers. The
// 5 KB limit includes any SSML tags that you specify. The service returns the synthesized audio stream as an array of
// bytes.
//
// **See also:** [The HTTP
// interface](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-usingHTTP#usingHTTP).
//
// **Note:** The Arabic, Chinese, Czech, Dutch (Belgian and Netherlands), Australian English, Korean, and Swedish
// languages and voices are supported only for IBM Cloud; they are deprecated for IBM Cloud Pak for Data. Also, the
// `ar-AR_OmarVoice` voice is deprecated; use the `ar-MS_OmarVoice` voice instead.
//
// ### Audio formats (accept types)
//
//  The service can return audio in the following formats (MIME types).
// * Where indicated, you can optionally specify the sampling rate (`rate`) of the audio. You must specify a sampling
// rate for the `audio/l16` and `audio/mulaw` formats. A specified sampling rate must lie in the range of 8 kHz to 192
// kHz. Some formats restrict the sampling rate to certain values, as noted.
// * For the `audio/l16` format, you can optionally specify the endianness (`endianness`) of the audio:
// `endianness=big-endian` or `endianness=little-endian`.
//
// Use the `Accept` header or the `accept` parameter to specify the requested format of the response audio. If you omit
// an audio format altogether, the service returns the audio in Ogg format with the Opus codec
// (`audio/ogg;codecs=opus`). The service always returns single-channel audio.
// * `audio/basic` - The service returns audio with a sampling rate of 8000 Hz.
// * `audio/flac` - You can optionally specify the `rate` of the audio. The default sampling rate is 22,050 Hz.
// * `audio/l16` - You must specify the `rate` of the audio. You can optionally specify the `endianness` of the audio.
// The default endianness is `little-endian`.
// * `audio/mp3` - You can optionally specify the `rate` of the audio. The default sampling rate is 22,050 Hz.
// * `audio/mpeg` - You can optionally specify the `rate` of the audio. The default sampling rate is 22,050 Hz.
// * `audio/mulaw` - You must specify the `rate` of the audio.
// * `audio/ogg` - The service returns the audio in the `vorbis` codec. You can optionally specify the `rate` of the
// audio. The default sampling rate is 22,050 Hz.
// * `audio/ogg;codecs=opus` - You can optionally specify the `rate` of the audio. Only the following values are valid
// sampling rates: `48000`, `24000`, `16000`, `12000`, or `8000`. If you specify a value other than one of these, the
// service returns an error. The default sampling rate is 48,000 Hz.
// * `audio/ogg;codecs=vorbis` - You can optionally specify the `rate` of the audio. The default sampling rate is 22,050
// Hz.
// * `audio/wav` - You can optionally specify the `rate` of the audio. The default sampling rate is 22,050 Hz.
// * `audio/webm` - The service returns the audio in the `opus` codec. The service returns audio with a sampling rate of
// 48,000 Hz.
// * `audio/webm;codecs=opus` - The service returns audio with a sampling rate of 48,000 Hz.
// * `audio/webm;codecs=vorbis` - You can optionally specify the `rate` of the audio. The default sampling rate is
// 22,050 Hz.
//
// For more information about specifying an audio format, including additional details about some of the formats, see
// [Using audio formats](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-audio-formats).
//
// ### Warning messages
//
//  If a request includes invalid query parameters, the service returns a `Warnings` response header that provides
// messages about the invalid parameters. The warning includes a descriptive message and a list of invalid argument
// strings. For example, a message such as `"Unknown arguments:"` or `"Unknown url query arguments:"` followed by a list
// of the form `"{invalid_arg_1}, {invalid_arg_2}."` The request succeeds despite the warnings.
func (textToSpeech *TextToSpeechV1) Synthesize(synthesizeOptions *SynthesizeOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	return textToSpeech.SynthesizeWithContext(context.Background(), synthesizeOptions)
}

// SynthesizeWithContext is an alternate form of the Synthesize method which supports a Context parameter
func (textToSpeech *TextToSpeechV1) SynthesizeWithContext(ctx context.Context, synthesizeOptions *SynthesizeOptions) (result io.ReadCloser, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(synthesizeOptions, "synthesizeOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(synthesizeOptions, "synthesizeOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = textToSpeech.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(textToSpeech.Service.Options.URL, `/v1/synthesize`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range synthesizeOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "Synthesize")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "audio/basic")
	builder.AddHeader("Content-Type", "application/json")
	if synthesizeOptions.Accept != nil {
		builder.AddHeader("Accept", fmt.Sprint(*synthesizeOptions.Accept))
	}

	if synthesizeOptions.Voice != nil {
		builder.AddQuery("voice", fmt.Sprint(*synthesizeOptions.Voice))
	}
	if synthesizeOptions.CustomizationID != nil {
		builder.AddQuery("customization_id", fmt.Sprint(*synthesizeOptions.CustomizationID))
	}

	body := make(map[string]interface{})
	if synthesizeOptions.Text != nil {
		body["text"] = synthesizeOptions.Text
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = textToSpeech.Service.Request(request, &result)

	return
}

// GetPronunciation : Get pronunciation
// Gets the phonetic pronunciation for the specified word. You can request the pronunciation for a specific format. You
// can also request the pronunciation for a specific voice to see the default translation for the language of that voice
// or for a specific custom model to see the translation for that model.
//
// **See also:** [Querying a word from a
// language](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customWords#cuWordsQueryLanguage).
//
// **Note:** The Arabic, Chinese, Czech, Dutch (Belgian and Netherlands), Australian English, Korean, and Swedish
// languages and voices are supported only for IBM Cloud; they are deprecated for IBM Cloud Pak for Data. Also, the
// `ar-AR_OmarVoice` voice is deprecated; use the `ar-MS_OmarVoice` voice instead.
func (textToSpeech *TextToSpeechV1) GetPronunciation(getPronunciationOptions *GetPronunciationOptions) (result *Pronunciation, response *core.DetailedResponse, err error) {
	return textToSpeech.GetPronunciationWithContext(context.Background(), getPronunciationOptions)
}

// GetPronunciationWithContext is an alternate form of the GetPronunciation method which supports a Context parameter
func (textToSpeech *TextToSpeechV1) GetPronunciationWithContext(ctx context.Context, getPronunciationOptions *GetPronunciationOptions) (result *Pronunciation, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getPronunciationOptions, "getPronunciationOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getPronunciationOptions, "getPronunciationOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = textToSpeech.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(textToSpeech.Service.Options.URL, `/v1/pronunciation`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range getPronunciationOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "GetPronunciation")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddQuery("text", fmt.Sprint(*getPronunciationOptions.Text))
	if getPronunciationOptions.Voice != nil {
		builder.AddQuery("voice", fmt.Sprint(*getPronunciationOptions.Voice))
	}
	if getPronunciationOptions.Format != nil {
		builder.AddQuery("format", fmt.Sprint(*getPronunciationOptions.Format))
	}
	if getPronunciationOptions.CustomizationID != nil {
		builder.AddQuery("customization_id", fmt.Sprint(*getPronunciationOptions.CustomizationID))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = textToSpeech.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPronunciation)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateCustomModel : Create a custom model
// Creates a new empty custom model. You must specify a name for the new custom model. You can optionally specify the
// language and a description for the new model. The model is owned by the instance of the service whose credentials are
// used to create it.
//
// **See also:** [Creating a custom
// model](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customModels#cuModelsCreate).
//
// **Note:** The Arabic, Chinese, Czech, Dutch (Belgian and Netherlands), Australian English, Korean, and Swedish
// languages and voices are supported only for IBM Cloud; they are deprecated for IBM Cloud Pak for Data. Also, the
// `ar-AR` language identifier cannot be used to create a custom model; use the `ar-MS` identifier instead.
func (textToSpeech *TextToSpeechV1) CreateCustomModel(createCustomModelOptions *CreateCustomModelOptions) (result *CustomModel, response *core.DetailedResponse, err error) {
	return textToSpeech.CreateCustomModelWithContext(context.Background(), createCustomModelOptions)
}

// CreateCustomModelWithContext is an alternate form of the CreateCustomModel method which supports a Context parameter
func (textToSpeech *TextToSpeechV1) CreateCustomModelWithContext(ctx context.Context, createCustomModelOptions *CreateCustomModelOptions) (result *CustomModel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createCustomModelOptions, "createCustomModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createCustomModelOptions, "createCustomModelOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = textToSpeech.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(textToSpeech.Service.Options.URL, `/v1/customizations`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createCustomModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "CreateCustomModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if createCustomModelOptions.Name != nil {
		body["name"] = createCustomModelOptions.Name
	}
	if createCustomModelOptions.Language != nil {
		body["language"] = createCustomModelOptions.Language
	}
	if createCustomModelOptions.Description != nil {
		body["description"] = createCustomModelOptions.Description
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = textToSpeech.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCustomModel)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// ListCustomModels : List custom models
// Lists metadata such as the name and description for all custom models that are owned by an instance of the service.
// Specify a language to list the custom models for that language only. To see the words and prompts in addition to the
// metadata for a specific custom model, use the [Get a custom model](#getcustommodel) method. You must use credentials
// for the instance of the service that owns a model to list information about it.
//
// **See also:** [Querying all custom
// models](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customModels#cuModelsQueryAll).
func (textToSpeech *TextToSpeechV1) ListCustomModels(listCustomModelsOptions *ListCustomModelsOptions) (result *CustomModels, response *core.DetailedResponse, err error) {
	return textToSpeech.ListCustomModelsWithContext(context.Background(), listCustomModelsOptions)
}

// ListCustomModelsWithContext is an alternate form of the ListCustomModels method which supports a Context parameter
func (textToSpeech *TextToSpeechV1) ListCustomModelsWithContext(ctx context.Context, listCustomModelsOptions *ListCustomModelsOptions) (result *CustomModels, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listCustomModelsOptions, "listCustomModelsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = textToSpeech.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(textToSpeech.Service.Options.URL, `/v1/customizations`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listCustomModelsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "ListCustomModels")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	if listCustomModelsOptions.Language != nil {
		builder.AddQuery("language", fmt.Sprint(*listCustomModelsOptions.Language))
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = textToSpeech.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCustomModels)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// UpdateCustomModel : Update a custom model
// Updates information for the specified custom model. You can update metadata such as the name and description of the
// model. You can also update the words in the model and their translations. Adding a new translation for a word that
// already exists in a custom model overwrites the word's existing translation. A custom model can contain no more than
// 20,000 entries. You must use credentials for the instance of the service that owns a model to update it.
//
// You can define sounds-like or phonetic translations for words. A sounds-like translation consists of one or more
// words that, when combined, sound like the word. Phonetic translations are based on the SSML phoneme format for
// representing a word. You can specify them in standard International Phonetic Alphabet (IPA) representation
//
//   <code>&lt;phoneme alphabet="ipa" ph="t&#601;m&#712;&#593;to"&gt;&lt;/phoneme&gt;</code>
//
//   or in the proprietary IBM Symbolic Phonetic Representation (SPR)
//
//   <code>&lt;phoneme alphabet="ibm" ph="1gAstroEntxrYFXs"&gt;&lt;/phoneme&gt;</code>
//
// **See also:**
// * [Updating a custom
// model](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customModels#cuModelsUpdate)
// * [Adding words to a Japanese custom
// model](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customWords#cuJapaneseAdd)
// * [Understanding
// customization](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customIntro#customIntro).
func (textToSpeech *TextToSpeechV1) UpdateCustomModel(updateCustomModelOptions *UpdateCustomModelOptions) (response *core.DetailedResponse, err error) {
	return textToSpeech.UpdateCustomModelWithContext(context.Background(), updateCustomModelOptions)
}

// UpdateCustomModelWithContext is an alternate form of the UpdateCustomModel method which supports a Context parameter
func (textToSpeech *TextToSpeechV1) UpdateCustomModelWithContext(ctx context.Context, updateCustomModelOptions *UpdateCustomModelOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(updateCustomModelOptions, "updateCustomModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(updateCustomModelOptions, "updateCustomModelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *updateCustomModelOptions.CustomizationID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = textToSpeech.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(textToSpeech.Service.Options.URL, `/v1/customizations/{customization_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range updateCustomModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "UpdateCustomModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if updateCustomModelOptions.Name != nil {
		body["name"] = updateCustomModelOptions.Name
	}
	if updateCustomModelOptions.Description != nil {
		body["description"] = updateCustomModelOptions.Description
	}
	if updateCustomModelOptions.Words != nil {
		body["words"] = updateCustomModelOptions.Words
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = textToSpeech.Service.Request(request, nil)

	return
}

// GetCustomModel : Get a custom model
// Gets all information about a specified custom model. In addition to metadata such as the name and description of the
// custom model, the output includes the words and their translations that are defined for the model, as well as any
// prompts that are defined for the model. To see just the metadata for a model, use the [List custom
// models](#listcustommodels) method.
//
// **See also:** [Querying a custom
// model](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customModels#cuModelsQuery).
func (textToSpeech *TextToSpeechV1) GetCustomModel(getCustomModelOptions *GetCustomModelOptions) (result *CustomModel, response *core.DetailedResponse, err error) {
	return textToSpeech.GetCustomModelWithContext(context.Background(), getCustomModelOptions)
}

// GetCustomModelWithContext is an alternate form of the GetCustomModel method which supports a Context parameter
func (textToSpeech *TextToSpeechV1) GetCustomModelWithContext(ctx context.Context, getCustomModelOptions *GetCustomModelOptions) (result *CustomModel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getCustomModelOptions, "getCustomModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getCustomModelOptions, "getCustomModelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *getCustomModelOptions.CustomizationID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = textToSpeech.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(textToSpeech.Service.Options.URL, `/v1/customizations/{customization_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCustomModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "GetCustomModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = textToSpeech.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalCustomModel)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteCustomModel : Delete a custom model
// Deletes the specified custom model. You must use credentials for the instance of the service that owns a model to
// delete it.
//
// **See also:** [Deleting a custom
// model](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customModels#cuModelsDelete).
func (textToSpeech *TextToSpeechV1) DeleteCustomModel(deleteCustomModelOptions *DeleteCustomModelOptions) (response *core.DetailedResponse, err error) {
	return textToSpeech.DeleteCustomModelWithContext(context.Background(), deleteCustomModelOptions)
}

// DeleteCustomModelWithContext is an alternate form of the DeleteCustomModel method which supports a Context parameter
func (textToSpeech *TextToSpeechV1) DeleteCustomModelWithContext(ctx context.Context, deleteCustomModelOptions *DeleteCustomModelOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteCustomModelOptions, "deleteCustomModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteCustomModelOptions, "deleteCustomModelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *deleteCustomModelOptions.CustomizationID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = textToSpeech.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(textToSpeech.Service.Options.URL, `/v1/customizations/{customization_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteCustomModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "DeleteCustomModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = textToSpeech.Service.Request(request, nil)

	return
}

// AddWords : Add custom words
// Adds one or more words and their translations to the specified custom model. Adding a new translation for a word that
// already exists in a custom model overwrites the word's existing translation. A custom model can contain no more than
// 20,000 entries. You must use credentials for the instance of the service that owns a model to add words to it.
//
// You can define sounds-like or phonetic translations for words. A sounds-like translation consists of one or more
// words that, when combined, sound like the word. Phonetic translations are based on the SSML phoneme format for
// representing a word. You can specify them in standard International Phonetic Alphabet (IPA) representation
//
//   <code>&lt;phoneme alphabet="ipa" ph="t&#601;m&#712;&#593;to"&gt;&lt;/phoneme&gt;</code>
//
//   or in the proprietary IBM Symbolic Phonetic Representation (SPR)
//
//   <code>&lt;phoneme alphabet="ibm" ph="1gAstroEntxrYFXs"&gt;&lt;/phoneme&gt;</code>
//
// **See also:**
// * [Adding multiple words to a custom
// model](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customWords#cuWordsAdd)
// * [Adding words to a Japanese custom
// model](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customWords#cuJapaneseAdd)
// * [Understanding
// customization](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customIntro#customIntro).
func (textToSpeech *TextToSpeechV1) AddWords(addWordsOptions *AddWordsOptions) (response *core.DetailedResponse, err error) {
	return textToSpeech.AddWordsWithContext(context.Background(), addWordsOptions)
}

// AddWordsWithContext is an alternate form of the AddWords method which supports a Context parameter
func (textToSpeech *TextToSpeechV1) AddWordsWithContext(ctx context.Context, addWordsOptions *AddWordsOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addWordsOptions, "addWordsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addWordsOptions, "addWordsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *addWordsOptions.CustomizationID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = textToSpeech.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(textToSpeech.Service.Options.URL, `/v1/customizations/{customization_id}/words`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range addWordsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "AddWords")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if addWordsOptions.Words != nil {
		body["words"] = addWordsOptions.Words
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = textToSpeech.Service.Request(request, nil)

	return
}

// ListWords : List custom words
// Lists all of the words and their translations for the specified custom model. The output shows the translations as
// they are defined in the model. You must use credentials for the instance of the service that owns a model to list its
// words.
//
// **See also:** [Querying all words from a custom
// model](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customWords#cuWordsQueryModel).
func (textToSpeech *TextToSpeechV1) ListWords(listWordsOptions *ListWordsOptions) (result *Words, response *core.DetailedResponse, err error) {
	return textToSpeech.ListWordsWithContext(context.Background(), listWordsOptions)
}

// ListWordsWithContext is an alternate form of the ListWords method which supports a Context parameter
func (textToSpeech *TextToSpeechV1) ListWordsWithContext(ctx context.Context, listWordsOptions *ListWordsOptions) (result *Words, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listWordsOptions, "listWordsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listWordsOptions, "listWordsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *listWordsOptions.CustomizationID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = textToSpeech.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(textToSpeech.Service.Options.URL, `/v1/customizations/{customization_id}/words`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listWordsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "ListWords")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = textToSpeech.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalWords)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// AddWord : Add a custom word
// Adds a single word and its translation to the specified custom model. Adding a new translation for a word that
// already exists in a custom model overwrites the word's existing translation. A custom model can contain no more than
// 20,000 entries. You must use credentials for the instance of the service that owns a model to add a word to it.
//
// You can define sounds-like or phonetic translations for words. A sounds-like translation consists of one or more
// words that, when combined, sound like the word. Phonetic translations are based on the SSML phoneme format for
// representing a word. You can specify them in standard International Phonetic Alphabet (IPA) representation
//
//   <code>&lt;phoneme alphabet="ipa" ph="t&#601;m&#712;&#593;to"&gt;&lt;/phoneme&gt;</code>
//
//   or in the proprietary IBM Symbolic Phonetic Representation (SPR)
//
//   <code>&lt;phoneme alphabet="ibm" ph="1gAstroEntxrYFXs"&gt;&lt;/phoneme&gt;</code>
//
// **See also:**
// * [Adding a single word to a custom
// model](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customWords#cuWordAdd)
// * [Adding words to a Japanese custom
// model](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customWords#cuJapaneseAdd)
// * [Understanding
// customization](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customIntro#customIntro).
func (textToSpeech *TextToSpeechV1) AddWord(addWordOptions *AddWordOptions) (response *core.DetailedResponse, err error) {
	return textToSpeech.AddWordWithContext(context.Background(), addWordOptions)
}

// AddWordWithContext is an alternate form of the AddWord method which supports a Context parameter
func (textToSpeech *TextToSpeechV1) AddWordWithContext(ctx context.Context, addWordOptions *AddWordOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addWordOptions, "addWordOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addWordOptions, "addWordOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *addWordOptions.CustomizationID,
		"word":             *addWordOptions.Word,
	}

	builder := core.NewRequestBuilder(core.PUT)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = textToSpeech.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(textToSpeech.Service.Options.URL, `/v1/customizations/{customization_id}/words/{word}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range addWordOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "AddWord")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Content-Type", "application/json")

	body := make(map[string]interface{})
	if addWordOptions.Translation != nil {
		body["translation"] = addWordOptions.Translation
	}
	if addWordOptions.PartOfSpeech != nil {
		body["part_of_speech"] = addWordOptions.PartOfSpeech
	}
	_, err = builder.SetBodyContentJSON(body)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = textToSpeech.Service.Request(request, nil)

	return
}

// GetWord : Get a custom word
// Gets the translation for a single word from the specified custom model. The output shows the translation as it is
// defined in the model. You must use credentials for the instance of the service that owns a model to list its words.
//
// **See also:** [Querying a single word from a custom
// model](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customWords#cuWordQueryModel).
func (textToSpeech *TextToSpeechV1) GetWord(getWordOptions *GetWordOptions) (result *Translation, response *core.DetailedResponse, err error) {
	return textToSpeech.GetWordWithContext(context.Background(), getWordOptions)
}

// GetWordWithContext is an alternate form of the GetWord method which supports a Context parameter
func (textToSpeech *TextToSpeechV1) GetWordWithContext(ctx context.Context, getWordOptions *GetWordOptions) (result *Translation, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getWordOptions, "getWordOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getWordOptions, "getWordOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *getWordOptions.CustomizationID,
		"word":             *getWordOptions.Word,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = textToSpeech.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(textToSpeech.Service.Options.URL, `/v1/customizations/{customization_id}/words/{word}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getWordOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "GetWord")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = textToSpeech.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalTranslation)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteWord : Delete a custom word
// Deletes a single word from the specified custom model. You must use credentials for the instance of the service that
// owns a model to delete its words.
//
// **See also:** [Deleting a word from a custom
// model](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-customWords#cuWordDelete).
func (textToSpeech *TextToSpeechV1) DeleteWord(deleteWordOptions *DeleteWordOptions) (response *core.DetailedResponse, err error) {
	return textToSpeech.DeleteWordWithContext(context.Background(), deleteWordOptions)
}

// DeleteWordWithContext is an alternate form of the DeleteWord method which supports a Context parameter
func (textToSpeech *TextToSpeechV1) DeleteWordWithContext(ctx context.Context, deleteWordOptions *DeleteWordOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteWordOptions, "deleteWordOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteWordOptions, "deleteWordOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *deleteWordOptions.CustomizationID,
		"word":             *deleteWordOptions.Word,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = textToSpeech.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(textToSpeech.Service.Options.URL, `/v1/customizations/{customization_id}/words/{word}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteWordOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "DeleteWord")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = textToSpeech.Service.Request(request, nil)

	return
}

// ListCustomPrompts : List custom prompts
// Lists information about all custom prompts that are defined for a custom model. The information includes the prompt
// ID, prompt text, status, and optional speaker ID for each prompt of the custom model. You must use credentials for
// the instance of the service that owns the custom model. The same information about all of the prompts for a custom
// model is also provided by the [Get a custom model](#getcustommodel) method. That method provides complete details
// about a specified custom model, including its language, owner, custom words, and more. Custom prompts are supported
// only for use with US English custom models and voices.
//
// **See also:** [Listing custom
// prompts](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-tbe-custom-prompts#tbe-custom-prompts-list).
func (textToSpeech *TextToSpeechV1) ListCustomPrompts(listCustomPromptsOptions *ListCustomPromptsOptions) (result *Prompts, response *core.DetailedResponse, err error) {
	return textToSpeech.ListCustomPromptsWithContext(context.Background(), listCustomPromptsOptions)
}

// ListCustomPromptsWithContext is an alternate form of the ListCustomPrompts method which supports a Context parameter
func (textToSpeech *TextToSpeechV1) ListCustomPromptsWithContext(ctx context.Context, listCustomPromptsOptions *ListCustomPromptsOptions) (result *Prompts, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(listCustomPromptsOptions, "listCustomPromptsOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(listCustomPromptsOptions, "listCustomPromptsOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *listCustomPromptsOptions.CustomizationID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = textToSpeech.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(textToSpeech.Service.Options.URL, `/v1/customizations/{customization_id}/prompts`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range listCustomPromptsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "ListCustomPrompts")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = textToSpeech.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPrompts)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// AddCustomPrompt : Add a custom prompt
// Adds a custom prompt to a custom model. A prompt is defined by the text that is to be spoken, the audio for that
// text, a unique user-specified ID for the prompt, and an optional speaker ID. The information is used to generate
// prosodic data that is not visible to the user. This data is used by the service to produce the synthesized audio upon
// request. You must use credentials for the instance of the service that owns a custom model to add a prompt to it. You
// can add a maximum of 1000 custom prompts to a single custom model.
//
// You are recommended to assign meaningful values for prompt IDs. For example, use `goodbye` to identify a prompt that
// speaks a farewell message. Prompt IDs must be unique within a given custom model. You cannot define two prompts with
// the same name for the same custom model. If you provide the ID of an existing prompt, the previously uploaded prompt
// is replaced by the new information. The existing prompt is reprocessed by using the new text and audio and, if
// provided, new speaker model, and the prosody data associated with the prompt is updated.
//
// The quality of a prompt is undefined if the language of a prompt does not match the language of its custom model.
// This is consistent with any text or SSML that is specified for a speech synthesis request. The service makes a
// best-effort attempt to render the specified text for the prompt; it does not validate that the language of the text
// matches the language of the model.
//
// Adding a prompt is an asynchronous operation. Although it accepts less audio than speaker enrollment, the service
// must align the audio with the provided text. The time that it takes to process a prompt depends on the prompt itself.
// The processing time for a reasonably sized prompt generally matches the length of the audio (for example, it takes 20
// seconds to process a 20-second prompt).
//
// For shorter prompts, you can wait for a reasonable amount of time and then check the status of the prompt with the
// [Get a custom prompt](#getcustomprompt) method. For longer prompts, consider using that method to poll the service
// every few seconds to determine when the prompt becomes available. No prompt can be used for speech synthesis if it is
// in the `processing` or `failed` state. Only prompts that are in the `available` state can be used for speech
// synthesis.
//
// When it processes a request, the service attempts to align the text and the audio that are provided for the prompt.
// The text that is passed with a prompt must match the spoken audio as closely as possible. Optimally, the text and
// audio match exactly. The service does its best to align the specified text with the audio, and it can often
// compensate for mismatches between the two. But if the service cannot effectively align the text and the audio,
// possibly because the magnitude of mismatches between the two is too great, processing of the prompt fails.
//
// ### Evaluating a prompt
//
//  Always listen to and evaluate a prompt to determine its quality before using it in production. To evaluate a prompt,
// include only the single prompt in a speech synthesis request by using the following SSML extension, in this case for
// a prompt whose ID is `goodbye`:
//
// `<ibm:prompt id="goodbye"/>`
//
// In some cases, you might need to rerecord and resubmit a prompt as many as five times to address the following
// possible problems:
// * The service might fail to detect a mismatch between the prompt’s text and audio. The longer the prompt, the greater
// the chance for misalignment between its text and audio. Therefore, multiple shorter prompts are preferable to a
// single long prompt.
// * The text of a prompt might include a word that the service does not recognize. In this case, you can create a
// custom word and pronunciation pair to tell the service how to pronounce the word. You must then re-create the prompt.
// * The quality of the input audio might be insufficient or the service’s processing of the audio might fail to detect
// the intended prosody. Submitting new audio for the prompt can correct these issues.
//
// If a prompt that is created without a speaker ID does not adequately reflect the intended prosody, enrolling the
// speaker and providing a speaker ID for the prompt is one recommended means of potentially improving the quality of
// the prompt. This is especially important for shorter prompts such as "good-bye" or "thank you," where less audio data
// makes it more difficult to match the prosody of the speaker. Custom prompts are supported only for use with US
// English custom models and voices.
//
// **See also:**
// * [Add a custom
// prompt](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-tbe-create#tbe-create-add-prompt)
// * [Evaluate a custom
// prompt](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-tbe-create#tbe-create-evaluate-prompt)
// * [Rules for creating custom
// prompts](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-tbe-rules#tbe-rules-prompts).
func (textToSpeech *TextToSpeechV1) AddCustomPrompt(addCustomPromptOptions *AddCustomPromptOptions) (result *Prompt, response *core.DetailedResponse, err error) {
	return textToSpeech.AddCustomPromptWithContext(context.Background(), addCustomPromptOptions)
}

// AddCustomPromptWithContext is an alternate form of the AddCustomPrompt method which supports a Context parameter
func (textToSpeech *TextToSpeechV1) AddCustomPromptWithContext(ctx context.Context, addCustomPromptOptions *AddCustomPromptOptions) (result *Prompt, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(addCustomPromptOptions, "addCustomPromptOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(addCustomPromptOptions, "addCustomPromptOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *addCustomPromptOptions.CustomizationID,
		"prompt_id":        *addCustomPromptOptions.PromptID,
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = textToSpeech.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(textToSpeech.Service.Options.URL, `/v1/customizations/{customization_id}/prompts/{prompt_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range addCustomPromptOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "AddCustomPrompt")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	builder.AddFormData("metadata", "", "application/json", addCustomPromptOptions.Metadata)
	builder.AddFormData("file", "filename",
		"audio/wav", addCustomPromptOptions.File)

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = textToSpeech.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPrompt)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetCustomPrompt : Get a custom prompt
// Gets information about a specified custom prompt for a specified custom model. The information includes the prompt
// ID, prompt text, status, and optional speaker ID for each prompt of the custom model. You must use credentials for
// the instance of the service that owns the custom model. Custom prompts are supported only for use with US English
// custom models and voices.
//
// **See also:** [Listing custom
// prompts](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-tbe-custom-prompts#tbe-custom-prompts-list).
func (textToSpeech *TextToSpeechV1) GetCustomPrompt(getCustomPromptOptions *GetCustomPromptOptions) (result *Prompt, response *core.DetailedResponse, err error) {
	return textToSpeech.GetCustomPromptWithContext(context.Background(), getCustomPromptOptions)
}

// GetCustomPromptWithContext is an alternate form of the GetCustomPrompt method which supports a Context parameter
func (textToSpeech *TextToSpeechV1) GetCustomPromptWithContext(ctx context.Context, getCustomPromptOptions *GetCustomPromptOptions) (result *Prompt, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getCustomPromptOptions, "getCustomPromptOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getCustomPromptOptions, "getCustomPromptOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *getCustomPromptOptions.CustomizationID,
		"prompt_id":        *getCustomPromptOptions.PromptID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = textToSpeech.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(textToSpeech.Service.Options.URL, `/v1/customizations/{customization_id}/prompts/{prompt_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getCustomPromptOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "GetCustomPrompt")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = textToSpeech.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalPrompt)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteCustomPrompt : Delete a custom prompt
// Deletes an existing custom prompt from a custom model. The service deletes the prompt with the specified ID. You must
// use credentials for the instance of the service that owns the custom model from which the prompt is to be deleted.
//
// **Caution:** Deleting a custom prompt elicits a 400 response code from synthesis requests that attempt to use the
// prompt. Make sure that you do not attempt to use a deleted prompt in a production application. Custom prompts are
// supported only for use with US English custom models and voices.
//
// **See also:** [Deleting a custom
// prompt](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-tbe-custom-prompts#tbe-custom-prompts-delete).
func (textToSpeech *TextToSpeechV1) DeleteCustomPrompt(deleteCustomPromptOptions *DeleteCustomPromptOptions) (response *core.DetailedResponse, err error) {
	return textToSpeech.DeleteCustomPromptWithContext(context.Background(), deleteCustomPromptOptions)
}

// DeleteCustomPromptWithContext is an alternate form of the DeleteCustomPrompt method which supports a Context parameter
func (textToSpeech *TextToSpeechV1) DeleteCustomPromptWithContext(ctx context.Context, deleteCustomPromptOptions *DeleteCustomPromptOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteCustomPromptOptions, "deleteCustomPromptOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteCustomPromptOptions, "deleteCustomPromptOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"customization_id": *deleteCustomPromptOptions.CustomizationID,
		"prompt_id":        *deleteCustomPromptOptions.PromptID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = textToSpeech.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(textToSpeech.Service.Options.URL, `/v1/customizations/{customization_id}/prompts/{prompt_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteCustomPromptOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "DeleteCustomPrompt")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = textToSpeech.Service.Request(request, nil)

	return
}

// ListSpeakerModels : List speaker models
// Lists information about all speaker models that are defined for a service instance. The information includes the
// speaker ID and speaker name of each defined speaker. You must use credentials for the instance of a service to list
// its speakers. Speaker models and the custom prompts with which they are used are supported only for use with US
// English custom models and voices.
//
// **See also:** [Listing speaker
// models](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-tbe-speaker-models#tbe-speaker-models-list).
func (textToSpeech *TextToSpeechV1) ListSpeakerModels(listSpeakerModelsOptions *ListSpeakerModelsOptions) (result *Speakers, response *core.DetailedResponse, err error) {
	return textToSpeech.ListSpeakerModelsWithContext(context.Background(), listSpeakerModelsOptions)
}

// ListSpeakerModelsWithContext is an alternate form of the ListSpeakerModels method which supports a Context parameter
func (textToSpeech *TextToSpeechV1) ListSpeakerModelsWithContext(ctx context.Context, listSpeakerModelsOptions *ListSpeakerModelsOptions) (result *Speakers, response *core.DetailedResponse, err error) {
	err = core.ValidateStruct(listSpeakerModelsOptions, "listSpeakerModelsOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = textToSpeech.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(textToSpeech.Service.Options.URL, `/v1/speakers`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range listSpeakerModelsOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "ListSpeakerModels")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = textToSpeech.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSpeakers)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// CreateSpeakerModel : Create a speaker model
// Creates a new speaker model, which is an optional enrollment token for users who are to add prompts to custom models.
// A speaker model contains information about a user's voice. The service extracts this information from a WAV audio
// sample that you pass as the body of the request. Associating a speaker model with a prompt is optional, but the
// information that is extracted from the speaker model helps the service learn about the speaker's voice.
//
// A speaker model can make an appreciable difference in the quality of prompts, especially short prompts with
// relatively little audio, that are associated with that speaker. A speaker model can help the service produce a prompt
// with more confidence; the lack of a speaker model can potentially compromise the quality of a prompt.
//
// The gender of the speaker who creates a speaker model does not need to match the gender of a voice that is used with
// prompts that are associated with that speaker model. For example, a speaker model that is created by a male speaker
// can be associated with prompts that are spoken by female voices.
//
// You create a speaker model for a given instance of the service. The new speaker model is owned by the service
// instance whose credentials are used to create it. That same speaker can then be used to create prompts for all custom
// models within that service instance. No language is associated with a speaker model, but each custom model has a
// single specified language. You can add prompts only to US English models.
//
// You specify a name for the speaker when you create it. The name must be unique among all speaker names for the owning
// service instance. To re-create a speaker model for an existing speaker name, you must first delete the existing
// speaker model that has that name.
//
// Speaker enrollment is a synchronous operation. Although it accepts more audio data than a prompt, the process of
// adding a speaker is very fast. The service simply extracts information about the speaker’s voice from the audio.
// Unlike prompts, speaker models neither need nor accept a transcription of the audio. When the call returns, the audio
// is fully processed and the speaker enrollment is complete.
//
// The service returns a speaker ID with the request. A speaker ID is globally unique identifier (GUID) that you use to
// identify the speaker in subsequent requests to the service. Speaker models and the custom prompts with which they are
// used are supported only for use with US English custom models and voices.
//
// **See also:**
// * [Create a speaker
// model](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-tbe-create#tbe-create-speaker-model)
// * [Rules for creating speaker
// models](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-tbe-rules#tbe-rules-speakers).
func (textToSpeech *TextToSpeechV1) CreateSpeakerModel(createSpeakerModelOptions *CreateSpeakerModelOptions) (result *SpeakerModel, response *core.DetailedResponse, err error) {
	return textToSpeech.CreateSpeakerModelWithContext(context.Background(), createSpeakerModelOptions)
}

// CreateSpeakerModelWithContext is an alternate form of the CreateSpeakerModel method which supports a Context parameter
func (textToSpeech *TextToSpeechV1) CreateSpeakerModelWithContext(ctx context.Context, createSpeakerModelOptions *CreateSpeakerModelOptions) (result *SpeakerModel, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(createSpeakerModelOptions, "createSpeakerModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(createSpeakerModelOptions, "createSpeakerModelOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.POST)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = textToSpeech.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(textToSpeech.Service.Options.URL, `/v1/speakers`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range createSpeakerModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "CreateSpeakerModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")
	builder.AddHeader("Content-Type", "audio/wav")

	builder.AddQuery("speaker_name", fmt.Sprint(*createSpeakerModelOptions.SpeakerName))

	_, err = builder.SetBodyContent("audio/wav", nil, nil, createSpeakerModelOptions.Audio)
	if err != nil {
		return
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = textToSpeech.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSpeakerModel)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// GetSpeakerModel : Get a speaker model
// Gets information about all prompts that are defined by a specified speaker for all custom models that are owned by a
// service instance. The information is grouped by the customization IDs of the custom models. For each custom model,
// the information lists information about each prompt that is defined for that custom model by the speaker. You must
// use credentials for the instance of the service that owns a speaker model to list its prompts. Speaker models and the
// custom prompts with which they are used are supported only for use with US English custom models and voices.
//
// **See also:** [Listing the custom prompts for a speaker
// model](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-tbe-speaker-models#tbe-speaker-models-list-prompts).
func (textToSpeech *TextToSpeechV1) GetSpeakerModel(getSpeakerModelOptions *GetSpeakerModelOptions) (result *SpeakerCustomModels, response *core.DetailedResponse, err error) {
	return textToSpeech.GetSpeakerModelWithContext(context.Background(), getSpeakerModelOptions)
}

// GetSpeakerModelWithContext is an alternate form of the GetSpeakerModel method which supports a Context parameter
func (textToSpeech *TextToSpeechV1) GetSpeakerModelWithContext(ctx context.Context, getSpeakerModelOptions *GetSpeakerModelOptions) (result *SpeakerCustomModels, response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(getSpeakerModelOptions, "getSpeakerModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(getSpeakerModelOptions, "getSpeakerModelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"speaker_id": *getSpeakerModelOptions.SpeakerID,
	}

	builder := core.NewRequestBuilder(core.GET)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = textToSpeech.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(textToSpeech.Service.Options.URL, `/v1/speakers/{speaker_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range getSpeakerModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "GetSpeakerModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}
	builder.AddHeader("Accept", "application/json")

	request, err := builder.Build()
	if err != nil {
		return
	}

	var rawResponse map[string]json.RawMessage
	response, err = textToSpeech.Service.Request(request, &rawResponse)
	if err != nil {
		return
	}
	if rawResponse != nil {
		err = core.UnmarshalModel(rawResponse, "", &result, UnmarshalSpeakerCustomModels)
		if err != nil {
			return
		}
		response.Result = result
	}

	return
}

// DeleteSpeakerModel : Delete a speaker model
// Deletes an existing speaker model from the service instance. The service deletes the enrolled speaker with the
// specified speaker ID. You must use credentials for the instance of the service that owns a speaker model to delete
// the speaker.
//
// Any prompts that are associated with the deleted speaker are not affected by the speaker's deletion. The prosodic
// data that defines the quality of a prompt is established when the prompt is created. A prompt is static and remains
// unaffected by deletion of its associated speaker. However, the prompt cannot be resubmitted or updated with its
// original speaker once that speaker is deleted. Speaker models and the custom prompts with which they are used are
// supported only for use with US English custom models and voices.
//
// **See also:** [Deleting a speaker
// model](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-tbe-speaker-models#tbe-speaker-models-delete).
func (textToSpeech *TextToSpeechV1) DeleteSpeakerModel(deleteSpeakerModelOptions *DeleteSpeakerModelOptions) (response *core.DetailedResponse, err error) {
	return textToSpeech.DeleteSpeakerModelWithContext(context.Background(), deleteSpeakerModelOptions)
}

// DeleteSpeakerModelWithContext is an alternate form of the DeleteSpeakerModel method which supports a Context parameter
func (textToSpeech *TextToSpeechV1) DeleteSpeakerModelWithContext(ctx context.Context, deleteSpeakerModelOptions *DeleteSpeakerModelOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteSpeakerModelOptions, "deleteSpeakerModelOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteSpeakerModelOptions, "deleteSpeakerModelOptions")
	if err != nil {
		return
	}

	pathParamsMap := map[string]string{
		"speaker_id": *deleteSpeakerModelOptions.SpeakerID,
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = textToSpeech.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(textToSpeech.Service.Options.URL, `/v1/speakers/{speaker_id}`, pathParamsMap)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteSpeakerModelOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "DeleteSpeakerModel")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = textToSpeech.Service.Request(request, nil)

	return
}

// DeleteUserData : Delete labeled data
// Deletes all data that is associated with a specified customer ID. The method deletes all data for the customer ID,
// regardless of the method by which the information was added. The method has no effect if no data is associated with
// the customer ID. You must issue the request with credentials for the same instance of the service that was used to
// associate the customer ID with the data. You associate a customer ID with data by passing the `X-Watson-Metadata`
// header with a request that passes the data.
//
// **Note:** If you delete an instance of the service from the service console, all data associated with that service
// instance is automatically deleted. This includes all custom models and word/translation pairs, and all data related
// to speech synthesis requests.
//
// **See also:** [Information
// security](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-information-security#information-security).
func (textToSpeech *TextToSpeechV1) DeleteUserData(deleteUserDataOptions *DeleteUserDataOptions) (response *core.DetailedResponse, err error) {
	return textToSpeech.DeleteUserDataWithContext(context.Background(), deleteUserDataOptions)
}

// DeleteUserDataWithContext is an alternate form of the DeleteUserData method which supports a Context parameter
func (textToSpeech *TextToSpeechV1) DeleteUserDataWithContext(ctx context.Context, deleteUserDataOptions *DeleteUserDataOptions) (response *core.DetailedResponse, err error) {
	err = core.ValidateNotNil(deleteUserDataOptions, "deleteUserDataOptions cannot be nil")
	if err != nil {
		return
	}
	err = core.ValidateStruct(deleteUserDataOptions, "deleteUserDataOptions")
	if err != nil {
		return
	}

	builder := core.NewRequestBuilder(core.DELETE)
	builder = builder.WithContext(ctx)
	builder.EnableGzipCompression = textToSpeech.GetEnableGzipCompression()
	_, err = builder.ResolveRequestURL(textToSpeech.Service.Options.URL, `/v1/user_data`, nil)
	if err != nil {
		return
	}

	for headerName, headerValue := range deleteUserDataOptions.Headers {
		builder.AddHeader(headerName, headerValue)
	}

	sdkHeaders := common.GetSdkHeaders("text_to_speech", "V1", "DeleteUserData")
	for headerName, headerValue := range sdkHeaders {
		builder.AddHeader(headerName, headerValue)
	}

	builder.AddQuery("customer_id", fmt.Sprint(*deleteUserDataOptions.CustomerID))

	request, err := builder.Build()
	if err != nil {
		return
	}

	response, err = textToSpeech.Service.Request(request, nil)

	return
}

// AddCustomPromptOptions : The AddCustomPrompt options.
type AddCustomPromptOptions struct {
	// The customization ID (GUID) of the custom model. You must make the request with credentials for the instance of the
	// service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// The identifier of the prompt that is to be added to the custom model:
	// * Include a maximum of 49 characters in the ID.
	// * Include only alphanumeric characters and `_` (underscores) in the ID.
	// * Do not include XML sensitive characters (double quotes, single quotes, ampersands, angle brackets, and slashes) in
	// the ID.
	// * To add a new prompt, the ID must be unique for the specified custom model. Otherwise, the new information for the
	// prompt overwrites the existing prompt that has that ID.
	PromptID *string `json:"prompt_id" validate:"required,ne="`

	// Information about the prompt that is to be added to a custom model. The following example of a `PromptMetadata`
	// object includes both the required prompt text and an optional speaker model ID:
	//
	// `{ "prompt_text": "Thank you and good-bye!", "speaker_id": "823068b2-ed4e-11ea-b6e0-7b6456aa95cc" }`.
	Metadata *PromptMetadata `json:"metadata" validate:"required"`

	// An audio file that speaks the text of the prompt with intonation and prosody that matches how you would like the
	// prompt to be spoken.
	// * The prompt audio must be in WAV format and must have a minimum sampling rate of 16 kHz. The service accepts audio
	// with higher sampling rates. The service transcodes all audio to 16 kHz before processing it.
	// * The length of the prompt audio is limited to 30 seconds.
	File io.ReadCloser `json:"file" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewAddCustomPromptOptions : Instantiate AddCustomPromptOptions
func (*TextToSpeechV1) NewAddCustomPromptOptions(customizationID string, promptID string, metadata *PromptMetadata, file io.ReadCloser) *AddCustomPromptOptions {
	return &AddCustomPromptOptions{
		CustomizationID: core.StringPtr(customizationID),
		PromptID:        core.StringPtr(promptID),
		Metadata:        metadata,
		File:            file,
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *AddCustomPromptOptions) SetCustomizationID(customizationID string) *AddCustomPromptOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetPromptID : Allow user to set PromptID
func (_options *AddCustomPromptOptions) SetPromptID(promptID string) *AddCustomPromptOptions {
	_options.PromptID = core.StringPtr(promptID)
	return _options
}

// SetMetadata : Allow user to set Metadata
func (_options *AddCustomPromptOptions) SetMetadata(metadata *PromptMetadata) *AddCustomPromptOptions {
	_options.Metadata = metadata
	return _options
}

// SetFile : Allow user to set File
func (_options *AddCustomPromptOptions) SetFile(file io.ReadCloser) *AddCustomPromptOptions {
	_options.File = file
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *AddCustomPromptOptions) SetHeaders(param map[string]string) *AddCustomPromptOptions {
	options.Headers = param
	return options
}

// AddWordOptions : The AddWord options.
type AddWordOptions struct {
	// The customization ID (GUID) of the custom model. You must make the request with credentials for the instance of the
	// service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// The word that is to be added or updated for the custom model.
	Word *string `json:"word" validate:"required,ne="`

	// The phonetic or sounds-like translation for the word. A phonetic translation is based on the SSML format for
	// representing the phonetic string of a word either as an IPA translation or as an IBM SPR translation. The Arabic,
	// Chinese, Dutch, Australian English, and Korean languages support only IPA. A sounds-like is one or more words that,
	// when combined, sound like the word.
	Translation *string `json:"translation" validate:"required"`

	// **Japanese only.** The part of speech for the word. The service uses the value to produce the correct intonation for
	// the word. You can create only a single entry, with or without a single part of speech, for any word; you cannot
	// create multiple entries with different parts of speech for the same word. For more information, see [Working with
	// Japanese entries](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-rules#jaNotes).
	PartOfSpeech *string `json:"part_of_speech,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the AddWordOptions.PartOfSpeech property.
// **Japanese only.** The part of speech for the word. The service uses the value to produce the correct intonation for
// the word. You can create only a single entry, with or without a single part of speech, for any word; you cannot
// create multiple entries with different parts of speech for the same word. For more information, see [Working with
// Japanese entries](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-rules#jaNotes).
const (
	AddWordOptionsPartOfSpeechDosiConst = "Dosi"
	AddWordOptionsPartOfSpeechFukuConst = "Fuku"
	AddWordOptionsPartOfSpeechGobiConst = "Gobi"
	AddWordOptionsPartOfSpeechHokaConst = "Hoka"
	AddWordOptionsPartOfSpeechJodoConst = "Jodo"
	AddWordOptionsPartOfSpeechJosiConst = "Josi"
	AddWordOptionsPartOfSpeechKatoConst = "Kato"
	AddWordOptionsPartOfSpeechKedoConst = "Kedo"
	AddWordOptionsPartOfSpeechKeyoConst = "Keyo"
	AddWordOptionsPartOfSpeechKigoConst = "Kigo"
	AddWordOptionsPartOfSpeechKoyuConst = "Koyu"
	AddWordOptionsPartOfSpeechMesiConst = "Mesi"
	AddWordOptionsPartOfSpeechRetaConst = "Reta"
	AddWordOptionsPartOfSpeechStbiConst = "Stbi"
	AddWordOptionsPartOfSpeechSttoConst = "Stto"
	AddWordOptionsPartOfSpeechStzoConst = "Stzo"
	AddWordOptionsPartOfSpeechSujiConst = "Suji"
)

// NewAddWordOptions : Instantiate AddWordOptions
func (*TextToSpeechV1) NewAddWordOptions(customizationID string, word string, translation string) *AddWordOptions {
	return &AddWordOptions{
		CustomizationID: core.StringPtr(customizationID),
		Word:            core.StringPtr(word),
		Translation:     core.StringPtr(translation),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *AddWordOptions) SetCustomizationID(customizationID string) *AddWordOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetWord : Allow user to set Word
func (_options *AddWordOptions) SetWord(word string) *AddWordOptions {
	_options.Word = core.StringPtr(word)
	return _options
}

// SetTranslation : Allow user to set Translation
func (_options *AddWordOptions) SetTranslation(translation string) *AddWordOptions {
	_options.Translation = core.StringPtr(translation)
	return _options
}

// SetPartOfSpeech : Allow user to set PartOfSpeech
func (_options *AddWordOptions) SetPartOfSpeech(partOfSpeech string) *AddWordOptions {
	_options.PartOfSpeech = core.StringPtr(partOfSpeech)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *AddWordOptions) SetHeaders(param map[string]string) *AddWordOptions {
	options.Headers = param
	return options
}

// AddWordsOptions : The AddWords options.
type AddWordsOptions struct {
	// The customization ID (GUID) of the custom model. You must make the request with credentials for the instance of the
	// service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// The [Add custom words](#addwords) method accepts an array of `Word` objects. Each object provides a word that is to
	// be added or updated for the custom model and the word's translation.
	//
	// The [List custom words](#listwords) method returns an array of `Word` objects. Each object shows a word and its
	// translation from the custom model. The words are listed in alphabetical order, with uppercase letters listed before
	// lowercase letters. The array is empty if the custom model contains no words.
	Words []Word `json:"words" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewAddWordsOptions : Instantiate AddWordsOptions
func (*TextToSpeechV1) NewAddWordsOptions(customizationID string, words []Word) *AddWordsOptions {
	return &AddWordsOptions{
		CustomizationID: core.StringPtr(customizationID),
		Words:           words,
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *AddWordsOptions) SetCustomizationID(customizationID string) *AddWordsOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetWords : Allow user to set Words
func (_options *AddWordsOptions) SetWords(words []Word) *AddWordsOptions {
	_options.Words = words
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *AddWordsOptions) SetHeaders(param map[string]string) *AddWordsOptions {
	options.Headers = param
	return options
}

// CreateCustomModelOptions : The CreateCustomModel options.
type CreateCustomModelOptions struct {
	// The name of the new custom model.
	Name *string `json:"name" validate:"required"`

	// The language of the new custom model. You create a custom model for a specific language, not for a specific voice. A
	// custom model can be used with any voice for its specified language. Omit the parameter to use the the default
	// language, `en-US`.
	//
	// **Important:** If you are using the service on IBM Cloud Pak for Data _and_ you install the neural voices, the
	// `language`parameter is required. You must specify the language for the custom model in the indicated format (for
	// example, `en-AU` for Australian English). The request fails if you do not specify a language.
	Language *string `json:"language,omitempty"`

	// A description of the new custom model. Specifying a description is recommended.
	Description *string `json:"description,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the CreateCustomModelOptions.Language property.
// The language of the new custom model. You create a custom model for a specific language, not for a specific voice. A
// custom model can be used with any voice for its specified language. Omit the parameter to use the the default
// language, `en-US`.
//
// **Important:** If you are using the service on IBM Cloud Pak for Data _and_ you install the neural voices, the
// `language`parameter is required. You must specify the language for the custom model in the indicated format (for
// example, `en-AU` for Australian English). The request fails if you do not specify a language.
const (
	CreateCustomModelOptionsLanguageArMsConst = "ar-MS"
	CreateCustomModelOptionsLanguageCsCzConst = "cs-CZ"
	CreateCustomModelOptionsLanguageDeDeConst = "de-DE"
	CreateCustomModelOptionsLanguageEnAuConst = "en-AU"
	CreateCustomModelOptionsLanguageEnGbConst = "en-GB"
	CreateCustomModelOptionsLanguageEnUsConst = "en-US"
	CreateCustomModelOptionsLanguageEsEsConst = "es-ES"
	CreateCustomModelOptionsLanguageEsLaConst = "es-LA"
	CreateCustomModelOptionsLanguageEsUsConst = "es-US"
	CreateCustomModelOptionsLanguageFrCaConst = "fr-CA"
	CreateCustomModelOptionsLanguageFrFrConst = "fr-FR"
	CreateCustomModelOptionsLanguageItItConst = "it-IT"
	CreateCustomModelOptionsLanguageJaJpConst = "ja-JP"
	CreateCustomModelOptionsLanguageKoKrConst = "ko-KR"
	CreateCustomModelOptionsLanguageNlBeConst = "nl-BE"
	CreateCustomModelOptionsLanguageNlNlConst = "nl-NL"
	CreateCustomModelOptionsLanguagePtBrConst = "pt-BR"
	CreateCustomModelOptionsLanguageSvSeConst = "sv-SE"
	CreateCustomModelOptionsLanguageZhCnConst = "zh-CN"
)

// NewCreateCustomModelOptions : Instantiate CreateCustomModelOptions
func (*TextToSpeechV1) NewCreateCustomModelOptions(name string) *CreateCustomModelOptions {
	return &CreateCustomModelOptions{
		Name: core.StringPtr(name),
	}
}

// SetName : Allow user to set Name
func (_options *CreateCustomModelOptions) SetName(name string) *CreateCustomModelOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetLanguage : Allow user to set Language
func (_options *CreateCustomModelOptions) SetLanguage(language string) *CreateCustomModelOptions {
	_options.Language = core.StringPtr(language)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *CreateCustomModelOptions) SetDescription(description string) *CreateCustomModelOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateCustomModelOptions) SetHeaders(param map[string]string) *CreateCustomModelOptions {
	options.Headers = param
	return options
}

// CreateSpeakerModelOptions : The CreateSpeakerModel options.
type CreateSpeakerModelOptions struct {
	// The name of the speaker that is to be added to the service instance.
	// * Include a maximum of 49 characters in the name.
	// * Include only alphanumeric characters and `_` (underscores) in the name.
	// * Do not include XML sensitive characters (double quotes, single quotes, ampersands, angle brackets, and slashes) in
	// the name.
	// * Do not use the name of an existing speaker that is already defined for the service instance.
	SpeakerName *string `json:"speaker_name" validate:"required"`

	// An enrollment audio file that contains a sample of the speaker’s voice.
	// * The enrollment audio must be in WAV format and must have a minimum sampling rate of 16 kHz. The service accepts
	// audio with higher sampling rates. It transcodes all audio to 16 kHz before processing it.
	// * The length of the enrollment audio is limited to 1 minute. Speaking one or two paragraphs of text that include
	// five to ten sentences is recommended.
	Audio io.ReadCloser `json:"audio" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewCreateSpeakerModelOptions : Instantiate CreateSpeakerModelOptions
func (*TextToSpeechV1) NewCreateSpeakerModelOptions(speakerName string, audio io.ReadCloser) *CreateSpeakerModelOptions {
	return &CreateSpeakerModelOptions{
		SpeakerName: core.StringPtr(speakerName),
		Audio:       audio,
	}
}

// SetSpeakerName : Allow user to set SpeakerName
func (_options *CreateSpeakerModelOptions) SetSpeakerName(speakerName string) *CreateSpeakerModelOptions {
	_options.SpeakerName = core.StringPtr(speakerName)
	return _options
}

// SetAudio : Allow user to set Audio
func (_options *CreateSpeakerModelOptions) SetAudio(audio io.ReadCloser) *CreateSpeakerModelOptions {
	_options.Audio = audio
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *CreateSpeakerModelOptions) SetHeaders(param map[string]string) *CreateSpeakerModelOptions {
	options.Headers = param
	return options
}

// CustomModel : Information about an existing custom model.
type CustomModel struct {
	// The customization ID (GUID) of the custom model. The [Create a custom model](#createcustommodel) method returns only
	// this field. It does not not return the other fields of this object.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// The name of the custom model.
	Name *string `json:"name,omitempty"`

	// The language identifier of the custom model (for example, `en-US`).
	Language *string `json:"language,omitempty"`

	// The GUID of the credentials for the instance of the service that owns the custom model.
	Owner *string `json:"owner,omitempty"`

	// The date and time in Coordinated Universal Time (UTC) at which the custom model was created. The value is provided
	// in full ISO 8601 format (`YYYY-MM-DDThh:mm:ss.sTZD`).
	Created *string `json:"created,omitempty"`

	// The date and time in Coordinated Universal Time (UTC) at which the custom model was last modified. The `created` and
	// `updated` fields are equal when a model is first added but has yet to be updated. The value is provided in full ISO
	// 8601 format (`YYYY-MM-DDThh:mm:ss.sTZD`).
	LastModified *string `json:"last_modified,omitempty"`

	// The description of the custom model.
	Description *string `json:"description,omitempty"`

	// An array of `Word` objects that lists the words and their translations from the custom model. The words are listed
	// in alphabetical order, with uppercase letters listed before lowercase letters. The array is empty if no words are
	// defined for the custom model. This field is returned only by the [Get a custom model](#getcustommodel) method.
	Words []Word `json:"words,omitempty"`

	// An array of `Prompt` objects that provides information about the prompts that are defined for the specified custom
	// model. The array is empty if no prompts are defined for the custom model. This field is returned only by the [Get a
	// custom model](#getcustommodel) method.
	Prompts []Prompt `json:"prompts,omitempty"`
}

// UnmarshalCustomModel unmarshals an instance of CustomModel from the specified map of raw messages.
func UnmarshalCustomModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CustomModel)
	err = core.UnmarshalPrimitive(m, "customization_id", &obj.CustomizationID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "language", &obj.Language)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "owner", &obj.Owner)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "created", &obj.Created)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "last_modified", &obj.LastModified)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "words", &obj.Words, UnmarshalWord)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "prompts", &obj.Prompts, UnmarshalPrompt)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// CustomModels : Information about existing custom models.
type CustomModels struct {
	// An array of `CustomModel` objects that provides information about each available custom model. The array is empty if
	// the requesting credentials own no custom models (if no language is specified) or own no custom models for the
	// specified language.
	Customizations []CustomModel `json:"customizations" validate:"required"`
}

// UnmarshalCustomModels unmarshals an instance of CustomModels from the specified map of raw messages.
func UnmarshalCustomModels(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(CustomModels)
	err = core.UnmarshalModel(m, "customizations", &obj.Customizations, UnmarshalCustomModel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// DeleteCustomModelOptions : The DeleteCustomModel options.
type DeleteCustomModelOptions struct {
	// The customization ID (GUID) of the custom model. You must make the request with credentials for the instance of the
	// service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteCustomModelOptions : Instantiate DeleteCustomModelOptions
func (*TextToSpeechV1) NewDeleteCustomModelOptions(customizationID string) *DeleteCustomModelOptions {
	return &DeleteCustomModelOptions{
		CustomizationID: core.StringPtr(customizationID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *DeleteCustomModelOptions) SetCustomizationID(customizationID string) *DeleteCustomModelOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteCustomModelOptions) SetHeaders(param map[string]string) *DeleteCustomModelOptions {
	options.Headers = param
	return options
}

// DeleteCustomPromptOptions : The DeleteCustomPrompt options.
type DeleteCustomPromptOptions struct {
	// The customization ID (GUID) of the custom model. You must make the request with credentials for the instance of the
	// service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// The identifier (name) of the prompt that is to be deleted.
	PromptID *string `json:"prompt_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteCustomPromptOptions : Instantiate DeleteCustomPromptOptions
func (*TextToSpeechV1) NewDeleteCustomPromptOptions(customizationID string, promptID string) *DeleteCustomPromptOptions {
	return &DeleteCustomPromptOptions{
		CustomizationID: core.StringPtr(customizationID),
		PromptID:        core.StringPtr(promptID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *DeleteCustomPromptOptions) SetCustomizationID(customizationID string) *DeleteCustomPromptOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetPromptID : Allow user to set PromptID
func (_options *DeleteCustomPromptOptions) SetPromptID(promptID string) *DeleteCustomPromptOptions {
	_options.PromptID = core.StringPtr(promptID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteCustomPromptOptions) SetHeaders(param map[string]string) *DeleteCustomPromptOptions {
	options.Headers = param
	return options
}

// DeleteSpeakerModelOptions : The DeleteSpeakerModel options.
type DeleteSpeakerModelOptions struct {
	// The speaker ID (GUID) of the speaker model. You must make the request with service credentials for the instance of
	// the service that owns the speaker model.
	SpeakerID *string `json:"speaker_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteSpeakerModelOptions : Instantiate DeleteSpeakerModelOptions
func (*TextToSpeechV1) NewDeleteSpeakerModelOptions(speakerID string) *DeleteSpeakerModelOptions {
	return &DeleteSpeakerModelOptions{
		SpeakerID: core.StringPtr(speakerID),
	}
}

// SetSpeakerID : Allow user to set SpeakerID
func (_options *DeleteSpeakerModelOptions) SetSpeakerID(speakerID string) *DeleteSpeakerModelOptions {
	_options.SpeakerID = core.StringPtr(speakerID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteSpeakerModelOptions) SetHeaders(param map[string]string) *DeleteSpeakerModelOptions {
	options.Headers = param
	return options
}

// DeleteUserDataOptions : The DeleteUserData options.
type DeleteUserDataOptions struct {
	// The customer ID for which all data is to be deleted.
	CustomerID *string `json:"customer_id" validate:"required"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteUserDataOptions : Instantiate DeleteUserDataOptions
func (*TextToSpeechV1) NewDeleteUserDataOptions(customerID string) *DeleteUserDataOptions {
	return &DeleteUserDataOptions{
		CustomerID: core.StringPtr(customerID),
	}
}

// SetCustomerID : Allow user to set CustomerID
func (_options *DeleteUserDataOptions) SetCustomerID(customerID string) *DeleteUserDataOptions {
	_options.CustomerID = core.StringPtr(customerID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteUserDataOptions) SetHeaders(param map[string]string) *DeleteUserDataOptions {
	options.Headers = param
	return options
}

// DeleteWordOptions : The DeleteWord options.
type DeleteWordOptions struct {
	// The customization ID (GUID) of the custom model. You must make the request with credentials for the instance of the
	// service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// The word that is to be deleted from the custom model.
	Word *string `json:"word" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewDeleteWordOptions : Instantiate DeleteWordOptions
func (*TextToSpeechV1) NewDeleteWordOptions(customizationID string, word string) *DeleteWordOptions {
	return &DeleteWordOptions{
		CustomizationID: core.StringPtr(customizationID),
		Word:            core.StringPtr(word),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *DeleteWordOptions) SetCustomizationID(customizationID string) *DeleteWordOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetWord : Allow user to set Word
func (_options *DeleteWordOptions) SetWord(word string) *DeleteWordOptions {
	_options.Word = core.StringPtr(word)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *DeleteWordOptions) SetHeaders(param map[string]string) *DeleteWordOptions {
	options.Headers = param
	return options
}

// GetCustomModelOptions : The GetCustomModel options.
type GetCustomModelOptions struct {
	// The customization ID (GUID) of the custom model. You must make the request with credentials for the instance of the
	// service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCustomModelOptions : Instantiate GetCustomModelOptions
func (*TextToSpeechV1) NewGetCustomModelOptions(customizationID string) *GetCustomModelOptions {
	return &GetCustomModelOptions{
		CustomizationID: core.StringPtr(customizationID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *GetCustomModelOptions) SetCustomizationID(customizationID string) *GetCustomModelOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetCustomModelOptions) SetHeaders(param map[string]string) *GetCustomModelOptions {
	options.Headers = param
	return options
}

// GetCustomPromptOptions : The GetCustomPrompt options.
type GetCustomPromptOptions struct {
	// The customization ID (GUID) of the custom model. You must make the request with credentials for the instance of the
	// service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// The identifier (name) of the prompt.
	PromptID *string `json:"prompt_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetCustomPromptOptions : Instantiate GetCustomPromptOptions
func (*TextToSpeechV1) NewGetCustomPromptOptions(customizationID string, promptID string) *GetCustomPromptOptions {
	return &GetCustomPromptOptions{
		CustomizationID: core.StringPtr(customizationID),
		PromptID:        core.StringPtr(promptID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *GetCustomPromptOptions) SetCustomizationID(customizationID string) *GetCustomPromptOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetPromptID : Allow user to set PromptID
func (_options *GetCustomPromptOptions) SetPromptID(promptID string) *GetCustomPromptOptions {
	_options.PromptID = core.StringPtr(promptID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetCustomPromptOptions) SetHeaders(param map[string]string) *GetCustomPromptOptions {
	options.Headers = param
	return options
}

// GetPronunciationOptions : The GetPronunciation options.
type GetPronunciationOptions struct {
	// The word for which the pronunciation is requested.
	Text *string `json:"text" validate:"required"`

	// A voice that specifies the language in which the pronunciation is to be returned. All voices for the same language
	// (for example, `en-US`) return the same translation.
	Voice *string `json:"voice,omitempty"`

	// The phoneme format in which to return the pronunciation. The Arabic, Chinese, Dutch, Australian English, and Korean
	// languages support only IPA. Omit the parameter to obtain the pronunciation in the default format.
	Format *string `json:"format,omitempty"`

	// The customization ID (GUID) of a custom model for which the pronunciation is to be returned. The language of a
	// specified custom model must match the language of the specified voice. If the word is not defined in the specified
	// custom model, the service returns the default translation for the custom model's language. You must make the request
	// with credentials for the instance of the service that owns the custom model. Omit the parameter to see the
	// translation for the specified voice with no customization.
	CustomizationID *string `json:"customization_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetPronunciationOptions.Voice property.
// A voice that specifies the language in which the pronunciation is to be returned. All voices for the same language
// (for example, `en-US`) return the same translation.
const (
	GetPronunciationOptionsVoiceArArOmarvoiceConst        = "ar-AR_OmarVoice"
	GetPronunciationOptionsVoiceArMsOmarvoiceConst        = "ar-MS_OmarVoice"
	GetPronunciationOptionsVoiceCsCzAlenavoiceConst       = "cs-CZ_AlenaVoice"
	GetPronunciationOptionsVoiceDeDeBirgitv3voiceConst    = "de-DE_BirgitV3Voice"
	GetPronunciationOptionsVoiceDeDeBirgitvoiceConst      = "de-DE_BirgitVoice"
	GetPronunciationOptionsVoiceDeDeDieterv3voiceConst    = "de-DE_DieterV3Voice"
	GetPronunciationOptionsVoiceDeDeDietervoiceConst      = "de-DE_DieterVoice"
	GetPronunciationOptionsVoiceDeDeErikav3voiceConst     = "de-DE_ErikaV3Voice"
	GetPronunciationOptionsVoiceEnAuCraigvoiceConst       = "en-AU_CraigVoice"
	GetPronunciationOptionsVoiceEnAuMadisonvoiceConst     = "en-AU_MadisonVoice"
	GetPronunciationOptionsVoiceEnAuStevevoiceConst       = "en-AU_SteveVoice"
	GetPronunciationOptionsVoiceEnGbCharlottev3voiceConst = "en-GB_CharlotteV3Voice"
	GetPronunciationOptionsVoiceEnGbJamesv3voiceConst     = "en-GB_JamesV3Voice"
	GetPronunciationOptionsVoiceEnGbKatev3voiceConst      = "en-GB_KateV3Voice"
	GetPronunciationOptionsVoiceEnGbKatevoiceConst        = "en-GB_KateVoice"
	GetPronunciationOptionsVoiceEnUsAllisonv3voiceConst   = "en-US_AllisonV3Voice"
	GetPronunciationOptionsVoiceEnUsAllisonvoiceConst     = "en-US_AllisonVoice"
	GetPronunciationOptionsVoiceEnUsEmilyv3voiceConst     = "en-US_EmilyV3Voice"
	GetPronunciationOptionsVoiceEnUsHenryv3voiceConst     = "en-US_HenryV3Voice"
	GetPronunciationOptionsVoiceEnUsKevinv3voiceConst     = "en-US_KevinV3Voice"
	GetPronunciationOptionsVoiceEnUsLisav3voiceConst      = "en-US_LisaV3Voice"
	GetPronunciationOptionsVoiceEnUsLisavoiceConst        = "en-US_LisaVoice"
	GetPronunciationOptionsVoiceEnUsMichaelv3voiceConst   = "en-US_MichaelV3Voice"
	GetPronunciationOptionsVoiceEnUsMichaelvoiceConst     = "en-US_MichaelVoice"
	GetPronunciationOptionsVoiceEnUsOliviav3voiceConst    = "en-US_OliviaV3Voice"
	GetPronunciationOptionsVoiceEsEsEnriquev3voiceConst   = "es-ES_EnriqueV3Voice"
	GetPronunciationOptionsVoiceEsEsEnriquevoiceConst     = "es-ES_EnriqueVoice"
	GetPronunciationOptionsVoiceEsEsLaurav3voiceConst     = "es-ES_LauraV3Voice"
	GetPronunciationOptionsVoiceEsEsLauravoiceConst       = "es-ES_LauraVoice"
	GetPronunciationOptionsVoiceEsLaSofiav3voiceConst     = "es-LA_SofiaV3Voice"
	GetPronunciationOptionsVoiceEsLaSofiavoiceConst       = "es-LA_SofiaVoice"
	GetPronunciationOptionsVoiceEsUsSofiav3voiceConst     = "es-US_SofiaV3Voice"
	GetPronunciationOptionsVoiceEsUsSofiavoiceConst       = "es-US_SofiaVoice"
	GetPronunciationOptionsVoiceFrCaLouisev3voiceConst    = "fr-CA_LouiseV3Voice"
	GetPronunciationOptionsVoiceFrFrNicolasv3voiceConst   = "fr-FR_NicolasV3Voice"
	GetPronunciationOptionsVoiceFrFrReneev3voiceConst     = "fr-FR_ReneeV3Voice"
	GetPronunciationOptionsVoiceFrFrReneevoiceConst       = "fr-FR_ReneeVoice"
	GetPronunciationOptionsVoiceItItFrancescav3voiceConst = "it-IT_FrancescaV3Voice"
	GetPronunciationOptionsVoiceItItFrancescavoiceConst   = "it-IT_FrancescaVoice"
	GetPronunciationOptionsVoiceJaJpEmiv3voiceConst       = "ja-JP_EmiV3Voice"
	GetPronunciationOptionsVoiceJaJpEmivoiceConst         = "ja-JP_EmiVoice"
	GetPronunciationOptionsVoiceKoKrHyunjunvoiceConst     = "ko-KR_HyunjunVoice"
	GetPronunciationOptionsVoiceKoKrSiwoovoiceConst       = "ko-KR_SiWooVoice"
	GetPronunciationOptionsVoiceKoKrYoungmivoiceConst     = "ko-KR_YoungmiVoice"
	GetPronunciationOptionsVoiceKoKrYunavoiceConst        = "ko-KR_YunaVoice"
	GetPronunciationOptionsVoiceNlBeAdelevoiceConst       = "nl-BE_AdeleVoice"
	GetPronunciationOptionsVoiceNlBeBramvoiceConst        = "nl-BE_BramVoice"
	GetPronunciationOptionsVoiceNlNlEmmavoiceConst        = "nl-NL_EmmaVoice"
	GetPronunciationOptionsVoiceNlNlLiamvoiceConst        = "nl-NL_LiamVoice"
	GetPronunciationOptionsVoicePtBrIsabelav3voiceConst   = "pt-BR_IsabelaV3Voice"
	GetPronunciationOptionsVoicePtBrIsabelavoiceConst     = "pt-BR_IsabelaVoice"
	GetPronunciationOptionsVoiceSvSeIngridvoiceConst      = "sv-SE_IngridVoice"
	GetPronunciationOptionsVoiceZhCnLinavoiceConst        = "zh-CN_LiNaVoice"
	GetPronunciationOptionsVoiceZhCnWangweivoiceConst     = "zh-CN_WangWeiVoice"
	GetPronunciationOptionsVoiceZhCnZhangjingvoiceConst   = "zh-CN_ZhangJingVoice"
)

// Constants associated with the GetPronunciationOptions.Format property.
// The phoneme format in which to return the pronunciation. The Arabic, Chinese, Dutch, Australian English, and Korean
// languages support only IPA. Omit the parameter to obtain the pronunciation in the default format.
const (
	GetPronunciationOptionsFormatIBMConst = "ibm"
	GetPronunciationOptionsFormatIpaConst = "ipa"
)

// NewGetPronunciationOptions : Instantiate GetPronunciationOptions
func (*TextToSpeechV1) NewGetPronunciationOptions(text string) *GetPronunciationOptions {
	return &GetPronunciationOptions{
		Text: core.StringPtr(text),
	}
}

// SetText : Allow user to set Text
func (_options *GetPronunciationOptions) SetText(text string) *GetPronunciationOptions {
	_options.Text = core.StringPtr(text)
	return _options
}

// SetVoice : Allow user to set Voice
func (_options *GetPronunciationOptions) SetVoice(voice string) *GetPronunciationOptions {
	_options.Voice = core.StringPtr(voice)
	return _options
}

// SetFormat : Allow user to set Format
func (_options *GetPronunciationOptions) SetFormat(format string) *GetPronunciationOptions {
	_options.Format = core.StringPtr(format)
	return _options
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *GetPronunciationOptions) SetCustomizationID(customizationID string) *GetPronunciationOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetPronunciationOptions) SetHeaders(param map[string]string) *GetPronunciationOptions {
	options.Headers = param
	return options
}

// GetSpeakerModelOptions : The GetSpeakerModel options.
type GetSpeakerModelOptions struct {
	// The speaker ID (GUID) of the speaker model. You must make the request with service credentials for the instance of
	// the service that owns the speaker model.
	SpeakerID *string `json:"speaker_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetSpeakerModelOptions : Instantiate GetSpeakerModelOptions
func (*TextToSpeechV1) NewGetSpeakerModelOptions(speakerID string) *GetSpeakerModelOptions {
	return &GetSpeakerModelOptions{
		SpeakerID: core.StringPtr(speakerID),
	}
}

// SetSpeakerID : Allow user to set SpeakerID
func (_options *GetSpeakerModelOptions) SetSpeakerID(speakerID string) *GetSpeakerModelOptions {
	_options.SpeakerID = core.StringPtr(speakerID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetSpeakerModelOptions) SetHeaders(param map[string]string) *GetSpeakerModelOptions {
	options.Headers = param
	return options
}

// GetVoiceOptions : The GetVoice options.
type GetVoiceOptions struct {
	// The voice for which information is to be returned.
	Voice *string `json:"voice" validate:"required,ne="`

	// The customization ID (GUID) of a custom model for which information is to be returned. You must make the request
	// with credentials for the instance of the service that owns the custom model. Omit the parameter to see information
	// about the specified voice with no customization.
	CustomizationID *string `json:"customization_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the GetVoiceOptions.Voice property.
// The voice for which information is to be returned.
const (
	GetVoiceOptionsVoiceArArOmarvoiceConst        = "ar-AR_OmarVoice"
	GetVoiceOptionsVoiceArMsOmarvoiceConst        = "ar-MS_OmarVoice"
	GetVoiceOptionsVoiceCsCzAlenavoiceConst       = "cs-CZ_AlenaVoice"
	GetVoiceOptionsVoiceDeDeBirgitv3voiceConst    = "de-DE_BirgitV3Voice"
	GetVoiceOptionsVoiceDeDeBirgitvoiceConst      = "de-DE_BirgitVoice"
	GetVoiceOptionsVoiceDeDeDieterv3voiceConst    = "de-DE_DieterV3Voice"
	GetVoiceOptionsVoiceDeDeDietervoiceConst      = "de-DE_DieterVoice"
	GetVoiceOptionsVoiceDeDeErikav3voiceConst     = "de-DE_ErikaV3Voice"
	GetVoiceOptionsVoiceEnAuCraigvoiceConst       = "en-AU_CraigVoice"
	GetVoiceOptionsVoiceEnAuMadisonvoiceConst     = "en-AU_MadisonVoice"
	GetVoiceOptionsVoiceEnAuStevevoiceConst       = "en-AU_SteveVoice"
	GetVoiceOptionsVoiceEnGbCharlottev3voiceConst = "en-GB_CharlotteV3Voice"
	GetVoiceOptionsVoiceEnGbJamesv3voiceConst     = "en-GB_JamesV3Voice"
	GetVoiceOptionsVoiceEnGbKatev3voiceConst      = "en-GB_KateV3Voice"
	GetVoiceOptionsVoiceEnGbKatevoiceConst        = "en-GB_KateVoice"
	GetVoiceOptionsVoiceEnUsAllisonv3voiceConst   = "en-US_AllisonV3Voice"
	GetVoiceOptionsVoiceEnUsAllisonvoiceConst     = "en-US_AllisonVoice"
	GetVoiceOptionsVoiceEnUsEmilyv3voiceConst     = "en-US_EmilyV3Voice"
	GetVoiceOptionsVoiceEnUsHenryv3voiceConst     = "en-US_HenryV3Voice"
	GetVoiceOptionsVoiceEnUsKevinv3voiceConst     = "en-US_KevinV3Voice"
	GetVoiceOptionsVoiceEnUsLisav3voiceConst      = "en-US_LisaV3Voice"
	GetVoiceOptionsVoiceEnUsLisavoiceConst        = "en-US_LisaVoice"
	GetVoiceOptionsVoiceEnUsMichaelv3voiceConst   = "en-US_MichaelV3Voice"
	GetVoiceOptionsVoiceEnUsMichaelvoiceConst     = "en-US_MichaelVoice"
	GetVoiceOptionsVoiceEnUsOliviav3voiceConst    = "en-US_OliviaV3Voice"
	GetVoiceOptionsVoiceEsEsEnriquev3voiceConst   = "es-ES_EnriqueV3Voice"
	GetVoiceOptionsVoiceEsEsEnriquevoiceConst     = "es-ES_EnriqueVoice"
	GetVoiceOptionsVoiceEsEsLaurav3voiceConst     = "es-ES_LauraV3Voice"
	GetVoiceOptionsVoiceEsEsLauravoiceConst       = "es-ES_LauraVoice"
	GetVoiceOptionsVoiceEsLaSofiav3voiceConst     = "es-LA_SofiaV3Voice"
	GetVoiceOptionsVoiceEsLaSofiavoiceConst       = "es-LA_SofiaVoice"
	GetVoiceOptionsVoiceEsUsSofiav3voiceConst     = "es-US_SofiaV3Voice"
	GetVoiceOptionsVoiceEsUsSofiavoiceConst       = "es-US_SofiaVoice"
	GetVoiceOptionsVoiceFrCaLouisev3voiceConst    = "fr-CA_LouiseV3Voice"
	GetVoiceOptionsVoiceFrFrNicolasv3voiceConst   = "fr-FR_NicolasV3Voice"
	GetVoiceOptionsVoiceFrFrReneev3voiceConst     = "fr-FR_ReneeV3Voice"
	GetVoiceOptionsVoiceFrFrReneevoiceConst       = "fr-FR_ReneeVoice"
	GetVoiceOptionsVoiceItItFrancescav3voiceConst = "it-IT_FrancescaV3Voice"
	GetVoiceOptionsVoiceItItFrancescavoiceConst   = "it-IT_FrancescaVoice"
	GetVoiceOptionsVoiceJaJpEmiv3voiceConst       = "ja-JP_EmiV3Voice"
	GetVoiceOptionsVoiceJaJpEmivoiceConst         = "ja-JP_EmiVoice"
	GetVoiceOptionsVoiceKoKrHyunjunvoiceConst     = "ko-KR_HyunjunVoice"
	GetVoiceOptionsVoiceKoKrSiwoovoiceConst       = "ko-KR_SiWooVoice"
	GetVoiceOptionsVoiceKoKrYoungmivoiceConst     = "ko-KR_YoungmiVoice"
	GetVoiceOptionsVoiceKoKrYunavoiceConst        = "ko-KR_YunaVoice"
	GetVoiceOptionsVoiceNlBeAdelevoiceConst       = "nl-BE_AdeleVoice"
	GetVoiceOptionsVoiceNlBeBramvoiceConst        = "nl-BE_BramVoice"
	GetVoiceOptionsVoiceNlNlEmmavoiceConst        = "nl-NL_EmmaVoice"
	GetVoiceOptionsVoiceNlNlLiamvoiceConst        = "nl-NL_LiamVoice"
	GetVoiceOptionsVoicePtBrIsabelav3voiceConst   = "pt-BR_IsabelaV3Voice"
	GetVoiceOptionsVoicePtBrIsabelavoiceConst     = "pt-BR_IsabelaVoice"
	GetVoiceOptionsVoiceSvSeIngridvoiceConst      = "sv-SE_IngridVoice"
	GetVoiceOptionsVoiceZhCnLinavoiceConst        = "zh-CN_LiNaVoice"
	GetVoiceOptionsVoiceZhCnWangweivoiceConst     = "zh-CN_WangWeiVoice"
	GetVoiceOptionsVoiceZhCnZhangjingvoiceConst   = "zh-CN_ZhangJingVoice"
)

// NewGetVoiceOptions : Instantiate GetVoiceOptions
func (*TextToSpeechV1) NewGetVoiceOptions(voice string) *GetVoiceOptions {
	return &GetVoiceOptions{
		Voice: core.StringPtr(voice),
	}
}

// SetVoice : Allow user to set Voice
func (_options *GetVoiceOptions) SetVoice(voice string) *GetVoiceOptions {
	_options.Voice = core.StringPtr(voice)
	return _options
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *GetVoiceOptions) SetCustomizationID(customizationID string) *GetVoiceOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetVoiceOptions) SetHeaders(param map[string]string) *GetVoiceOptions {
	options.Headers = param
	return options
}

// GetWordOptions : The GetWord options.
type GetWordOptions struct {
	// The customization ID (GUID) of the custom model. You must make the request with credentials for the instance of the
	// service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// The word that is to be queried from the custom model.
	Word *string `json:"word" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewGetWordOptions : Instantiate GetWordOptions
func (*TextToSpeechV1) NewGetWordOptions(customizationID string, word string) *GetWordOptions {
	return &GetWordOptions{
		CustomizationID: core.StringPtr(customizationID),
		Word:            core.StringPtr(word),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *GetWordOptions) SetCustomizationID(customizationID string) *GetWordOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetWord : Allow user to set Word
func (_options *GetWordOptions) SetWord(word string) *GetWordOptions {
	_options.Word = core.StringPtr(word)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *GetWordOptions) SetHeaders(param map[string]string) *GetWordOptions {
	options.Headers = param
	return options
}

// ListCustomModelsOptions : The ListCustomModels options.
type ListCustomModelsOptions struct {
	// The language for which custom models that are owned by the requesting credentials are to be returned. Omit the
	// parameter to see all custom models that are owned by the requester.
	Language *string `json:"language,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the ListCustomModelsOptions.Language property.
// The language for which custom models that are owned by the requesting credentials are to be returned. Omit the
// parameter to see all custom models that are owned by the requester.
const (
	ListCustomModelsOptionsLanguageArMsConst = "ar-MS"
	ListCustomModelsOptionsLanguageCsCzConst = "cs-CZ"
	ListCustomModelsOptionsLanguageDeDeConst = "de-DE"
	ListCustomModelsOptionsLanguageEnAuConst = "en-AU"
	ListCustomModelsOptionsLanguageEnGbConst = "en-GB"
	ListCustomModelsOptionsLanguageEnUsConst = "en-US"
	ListCustomModelsOptionsLanguageEsEsConst = "es-ES"
	ListCustomModelsOptionsLanguageEsLaConst = "es-LA"
	ListCustomModelsOptionsLanguageEsUsConst = "es-US"
	ListCustomModelsOptionsLanguageFrCaConst = "fr-CA"
	ListCustomModelsOptionsLanguageFrFrConst = "fr-FR"
	ListCustomModelsOptionsLanguageItItConst = "it-IT"
	ListCustomModelsOptionsLanguageJaJpConst = "ja-JP"
	ListCustomModelsOptionsLanguageKoKrConst = "ko-KR"
	ListCustomModelsOptionsLanguageNlBeConst = "nl-BE"
	ListCustomModelsOptionsLanguageNlNlConst = "nl-NL"
	ListCustomModelsOptionsLanguagePtBrConst = "pt-BR"
	ListCustomModelsOptionsLanguageSvSeConst = "sv-SE"
	ListCustomModelsOptionsLanguageZhCnConst = "zh-CN"
)

// NewListCustomModelsOptions : Instantiate ListCustomModelsOptions
func (*TextToSpeechV1) NewListCustomModelsOptions() *ListCustomModelsOptions {
	return &ListCustomModelsOptions{}
}

// SetLanguage : Allow user to set Language
func (_options *ListCustomModelsOptions) SetLanguage(language string) *ListCustomModelsOptions {
	_options.Language = core.StringPtr(language)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListCustomModelsOptions) SetHeaders(param map[string]string) *ListCustomModelsOptions {
	options.Headers = param
	return options
}

// ListCustomPromptsOptions : The ListCustomPrompts options.
type ListCustomPromptsOptions struct {
	// The customization ID (GUID) of the custom model. You must make the request with credentials for the instance of the
	// service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListCustomPromptsOptions : Instantiate ListCustomPromptsOptions
func (*TextToSpeechV1) NewListCustomPromptsOptions(customizationID string) *ListCustomPromptsOptions {
	return &ListCustomPromptsOptions{
		CustomizationID: core.StringPtr(customizationID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *ListCustomPromptsOptions) SetCustomizationID(customizationID string) *ListCustomPromptsOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListCustomPromptsOptions) SetHeaders(param map[string]string) *ListCustomPromptsOptions {
	options.Headers = param
	return options
}

// ListSpeakerModelsOptions : The ListSpeakerModels options.
type ListSpeakerModelsOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListSpeakerModelsOptions : Instantiate ListSpeakerModelsOptions
func (*TextToSpeechV1) NewListSpeakerModelsOptions() *ListSpeakerModelsOptions {
	return &ListSpeakerModelsOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListSpeakerModelsOptions) SetHeaders(param map[string]string) *ListSpeakerModelsOptions {
	options.Headers = param
	return options
}

// ListVoicesOptions : The ListVoices options.
type ListVoicesOptions struct {

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListVoicesOptions : Instantiate ListVoicesOptions
func (*TextToSpeechV1) NewListVoicesOptions() *ListVoicesOptions {
	return &ListVoicesOptions{}
}

// SetHeaders : Allow user to set Headers
func (options *ListVoicesOptions) SetHeaders(param map[string]string) *ListVoicesOptions {
	options.Headers = param
	return options
}

// ListWordsOptions : The ListWords options.
type ListWordsOptions struct {
	// The customization ID (GUID) of the custom model. You must make the request with credentials for the instance of the
	// service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewListWordsOptions : Instantiate ListWordsOptions
func (*TextToSpeechV1) NewListWordsOptions(customizationID string) *ListWordsOptions {
	return &ListWordsOptions{
		CustomizationID: core.StringPtr(customizationID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *ListWordsOptions) SetCustomizationID(customizationID string) *ListWordsOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *ListWordsOptions) SetHeaders(param map[string]string) *ListWordsOptions {
	options.Headers = param
	return options
}

// Prompt : Information about a custom prompt.
type Prompt struct {
	// The user-specified text of the prompt.
	Prompt *string `json:"prompt" validate:"required"`

	// The user-specified identifier (name) of the prompt.
	PromptID *string `json:"prompt_id" validate:"required"`

	// The status of the prompt:
	// * `processing`: The service received the request to add the prompt and is analyzing the validity of the prompt.
	// * `available`: The service successfully validated the prompt, which is now ready for use in a speech synthesis
	// request.
	// * `failed`: The service's validation of the prompt failed. The status of the prompt includes an `error` field that
	// describes the reason for the failure.
	Status *string `json:"status" validate:"required"`

	// If the status of the prompt is `failed`, an error message that describes the reason for the failure. The field is
	// omitted if no error occurred.
	Error *string `json:"error,omitempty"`

	// The speaker ID (GUID) of the speaker for which the prompt was defined. The field is omitted if no speaker ID was
	// specified.
	SpeakerID *string `json:"speaker_id,omitempty"`
}

// UnmarshalPrompt unmarshals an instance of Prompt from the specified map of raw messages.
func UnmarshalPrompt(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Prompt)
	err = core.UnmarshalPrimitive(m, "prompt", &obj.Prompt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "prompt_id", &obj.PromptID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "error", &obj.Error)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "speaker_id", &obj.SpeakerID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// PromptMetadata : Information about the prompt that is to be added to a custom model. The following example of a `PromptMetadata`
// object includes both the required prompt text and an optional speaker model ID:
//
// `{ "prompt_text": "Thank you and good-bye!", "speaker_id": "823068b2-ed4e-11ea-b6e0-7b6456aa95cc" }`.
type PromptMetadata struct {
	// The required written text of the spoken prompt. The length of a prompt's text is limited to a few sentences.
	// Speaking one or two sentences of text is the recommended limit. A prompt cannot contain more than 1000 characters of
	// text. Escape any XML control characters (double quotes, single quotes, ampersands, angle brackets, and slashes) that
	// appear in the text of the prompt.
	PromptText *string `json:"prompt_text" validate:"required"`

	// The optional speaker ID (GUID) of a previously defined speaker model that is to be associated with the prompt.
	SpeakerID *string `json:"speaker_id,omitempty"`
}

// NewPromptMetadata : Instantiate PromptMetadata (Generic Model Constructor)
func (*TextToSpeechV1) NewPromptMetadata(promptText string) (_model *PromptMetadata, err error) {
	_model = &PromptMetadata{
		PromptText: core.StringPtr(promptText),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalPromptMetadata unmarshals an instance of PromptMetadata from the specified map of raw messages.
func UnmarshalPromptMetadata(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(PromptMetadata)
	err = core.UnmarshalPrimitive(m, "prompt_text", &obj.PromptText)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "speaker_id", &obj.SpeakerID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Prompts : Information about the custom prompts that are defined for a custom model.
type Prompts struct {
	// An array of `Prompt` objects that provides information about the prompts that are defined for the specified custom
	// model. The array is empty if no prompts are defined for the custom model.
	Prompts []Prompt `json:"prompts" validate:"required"`
}

// UnmarshalPrompts unmarshals an instance of Prompts from the specified map of raw messages.
func UnmarshalPrompts(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Prompts)
	err = core.UnmarshalModel(m, "prompts", &obj.Prompts, UnmarshalPrompt)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Pronunciation : The pronunciation of the specified text.
type Pronunciation struct {
	// The pronunciation of the specified text in the requested voice and format. If a custom model is specified, the
	// pronunciation also reflects that custom model.
	Pronunciation *string `json:"pronunciation" validate:"required"`
}

// UnmarshalPronunciation unmarshals an instance of Pronunciation from the specified map of raw messages.
func UnmarshalPronunciation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Pronunciation)
	err = core.UnmarshalPrimitive(m, "pronunciation", &obj.Pronunciation)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Speaker : Information about a speaker model.
type Speaker struct {
	// The speaker ID (GUID) of the speaker.
	SpeakerID *string `json:"speaker_id" validate:"required"`

	// The user-defined name of the speaker.
	Name *string `json:"name" validate:"required"`
}

// UnmarshalSpeaker unmarshals an instance of Speaker from the specified map of raw messages.
func UnmarshalSpeaker(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Speaker)
	err = core.UnmarshalPrimitive(m, "speaker_id", &obj.SpeakerID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SpeakerCustomModel : A custom models for which the speaker has defined prompts.
type SpeakerCustomModel struct {
	// The customization ID (GUID) of a custom model for which the speaker has defined one or more prompts.
	CustomizationID *string `json:"customization_id" validate:"required"`

	// An array of `SpeakerPrompt` objects that provides information about each prompt that the user has defined for the
	// custom model.
	Prompts []SpeakerPrompt `json:"prompts" validate:"required"`
}

// UnmarshalSpeakerCustomModel unmarshals an instance of SpeakerCustomModel from the specified map of raw messages.
func UnmarshalSpeakerCustomModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SpeakerCustomModel)
	err = core.UnmarshalPrimitive(m, "customization_id", &obj.CustomizationID)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "prompts", &obj.Prompts, UnmarshalSpeakerPrompt)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SpeakerCustomModels : Custom models for which the speaker has defined prompts.
type SpeakerCustomModels struct {
	// An array of `SpeakerCustomModel` objects. Each object provides information about the prompts that are defined for a
	// specified speaker in the custom models that are owned by a specified service instance. The array is empty if no
	// prompts are defined for the speaker.
	Customizations []SpeakerCustomModel `json:"customizations" validate:"required"`
}

// UnmarshalSpeakerCustomModels unmarshals an instance of SpeakerCustomModels from the specified map of raw messages.
func UnmarshalSpeakerCustomModels(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SpeakerCustomModels)
	err = core.UnmarshalModel(m, "customizations", &obj.Customizations, UnmarshalSpeakerCustomModel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SpeakerModel : The speaker ID of the speaker model.
type SpeakerModel struct {
	// The speaker ID (GUID) of the speaker model.
	SpeakerID *string `json:"speaker_id" validate:"required"`
}

// UnmarshalSpeakerModel unmarshals an instance of SpeakerModel from the specified map of raw messages.
func UnmarshalSpeakerModel(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SpeakerModel)
	err = core.UnmarshalPrimitive(m, "speaker_id", &obj.SpeakerID)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SpeakerPrompt : A prompt that a speaker has defined for a custom model.
type SpeakerPrompt struct {
	// The user-specified text of the prompt.
	Prompt *string `json:"prompt" validate:"required"`

	// The user-specified identifier (name) of the prompt.
	PromptID *string `json:"prompt_id" validate:"required"`

	// The status of the prompt:
	// * `processing`: The service received the request to add the prompt and is analyzing the validity of the prompt.
	// * `available`: The service successfully validated the prompt, which is now ready for use in a speech synthesis
	// request.
	// * `failed`: The service's validation of the prompt failed. The status of the prompt includes an `error` field that
	// describes the reason for the failure.
	Status *string `json:"status" validate:"required"`

	// If the status of the prompt is `failed`, an error message that describes the reason for the failure. The field is
	// omitted if no error occurred.
	Error *string `json:"error,omitempty"`
}

// UnmarshalSpeakerPrompt unmarshals an instance of SpeakerPrompt from the specified map of raw messages.
func UnmarshalSpeakerPrompt(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SpeakerPrompt)
	err = core.UnmarshalPrimitive(m, "prompt", &obj.Prompt)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "prompt_id", &obj.PromptID)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "status", &obj.Status)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "error", &obj.Error)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Speakers : Information about all speaker models for the service instance.
type Speakers struct {
	// An array of `Speaker` objects that provides information about the speakers for the service instance. The array is
	// empty if the service instance has no speakers.
	Speakers []Speaker `json:"speakers" validate:"required"`
}

// UnmarshalSpeakers unmarshals an instance of Speakers from the specified map of raw messages.
func UnmarshalSpeakers(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Speakers)
	err = core.UnmarshalModel(m, "speakers", &obj.Speakers, UnmarshalSpeaker)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SupportedFeatures : Additional service features that are supported with the voice.
type SupportedFeatures struct {
	// If `true`, the voice can be customized; if `false`, the voice cannot be customized. (Same as `customizable`.).
	CustomPronunciation *bool `json:"custom_pronunciation" validate:"required"`

	// If `true`, the voice can be transformed by using the SSML &lt;voice-transformation&gt; element; if `false`, the
	// voice cannot be transformed. The feature was available only for the now-deprecated standard voices. You cannot use
	// the feature with neural voices.
	VoiceTransformation *bool `json:"voice_transformation" validate:"required"`
}

// UnmarshalSupportedFeatures unmarshals an instance of SupportedFeatures from the specified map of raw messages.
func UnmarshalSupportedFeatures(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(SupportedFeatures)
	err = core.UnmarshalPrimitive(m, "custom_pronunciation", &obj.CustomPronunciation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "voice_transformation", &obj.VoiceTransformation)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// SynthesizeOptions : The Synthesize options.
type SynthesizeOptions struct {
	// The text to synthesize.
	Text *string `json:"text" validate:"required"`

	// The requested format (MIME type) of the audio. You can use the `Accept` header or the `accept` parameter to specify
	// the audio format. For more information about specifying an audio format, see **Audio formats (accept types)** in the
	// method description.
	Accept *string `json:"Accept,omitempty"`

	// The voice to use for synthesis. If you omit the `voice` parameter, the service uses a default voice, which depends
	// on the version of the service that you are using:
	// * _For IBM Cloud,_ the service always uses the US English `en-US_MichaelV3Voice` by default.
	// * _For IBM Cloud Pak for Data,_ the default voice depends on the voices that you installed. If you installed the
	// _enhanced neural voices_, the service uses the US English `en-US_MichaelV3Voice` by default; if that voice is not
	// installed, you must specify a voice. If you installed the _neural voices_, the service always uses the Australian
	// English `en-AU_MadisonVoice` by default.
	//
	// **See also:** See also [Using languages and
	// voices](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-voices).
	Voice *string `json:"voice,omitempty"`

	// The customization ID (GUID) of a custom model to use for the synthesis. If a custom model is specified, it works
	// only if it matches the language of the indicated voice. You must make the request with credentials for the instance
	// of the service that owns the custom model. Omit the parameter to use the specified voice with no customization.
	CustomizationID *string `json:"customization_id,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// Constants associated with the SynthesizeOptions.Voice property.
// The voice to use for synthesis. If you omit the `voice` parameter, the service uses a default voice, which depends on
// the version of the service that you are using:
// * _For IBM Cloud,_ the service always uses the US English `en-US_MichaelV3Voice` by default.
// * _For IBM Cloud Pak for Data,_ the default voice depends on the voices that you installed. If you installed the
// _enhanced neural voices_, the service uses the US English `en-US_MichaelV3Voice` by default; if that voice is not
// installed, you must specify a voice. If you installed the _neural voices_, the service always uses the Australian
// English `en-AU_MadisonVoice` by default.
//
// **See also:** See also [Using languages and
// voices](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-voices).
const (
	SynthesizeOptionsVoiceArArOmarvoiceConst        = "ar-AR_OmarVoice"
	SynthesizeOptionsVoiceArMsOmarvoiceConst        = "ar-MS_OmarVoice"
	SynthesizeOptionsVoiceCsCzAlenavoiceConst       = "cs-CZ_AlenaVoice"
	SynthesizeOptionsVoiceDeDeBirgitv3voiceConst    = "de-DE_BirgitV3Voice"
	SynthesizeOptionsVoiceDeDeBirgitvoiceConst      = "de-DE_BirgitVoice"
	SynthesizeOptionsVoiceDeDeDieterv3voiceConst    = "de-DE_DieterV3Voice"
	SynthesizeOptionsVoiceDeDeDietervoiceConst      = "de-DE_DieterVoice"
	SynthesizeOptionsVoiceDeDeErikav3voiceConst     = "de-DE_ErikaV3Voice"
	SynthesizeOptionsVoiceEnAuCraigvoiceConst       = "en-AU_CraigVoice"
	SynthesizeOptionsVoiceEnAuMadisonvoiceConst     = "en-AU_MadisonVoice"
	SynthesizeOptionsVoiceEnAuStevevoiceConst       = "en-AU_SteveVoice"
	SynthesizeOptionsVoiceEnGbCharlottev3voiceConst = "en-GB_CharlotteV3Voice"
	SynthesizeOptionsVoiceEnGbJamesv3voiceConst     = "en-GB_JamesV3Voice"
	SynthesizeOptionsVoiceEnGbKatev3voiceConst      = "en-GB_KateV3Voice"
	SynthesizeOptionsVoiceEnGbKatevoiceConst        = "en-GB_KateVoice"
	SynthesizeOptionsVoiceEnUsAllisonv3voiceConst   = "en-US_AllisonV3Voice"
	SynthesizeOptionsVoiceEnUsAllisonvoiceConst     = "en-US_AllisonVoice"
	SynthesizeOptionsVoiceEnUsEmilyv3voiceConst     = "en-US_EmilyV3Voice"
	SynthesizeOptionsVoiceEnUsHenryv3voiceConst     = "en-US_HenryV3Voice"
	SynthesizeOptionsVoiceEnUsKevinv3voiceConst     = "en-US_KevinV3Voice"
	SynthesizeOptionsVoiceEnUsLisav3voiceConst      = "en-US_LisaV3Voice"
	SynthesizeOptionsVoiceEnUsLisavoiceConst        = "en-US_LisaVoice"
	SynthesizeOptionsVoiceEnUsMichaelv3voiceConst   = "en-US_MichaelV3Voice"
	SynthesizeOptionsVoiceEnUsMichaelvoiceConst     = "en-US_MichaelVoice"
	SynthesizeOptionsVoiceEnUsOliviav3voiceConst    = "en-US_OliviaV3Voice"
	SynthesizeOptionsVoiceEsEsEnriquev3voiceConst   = "es-ES_EnriqueV3Voice"
	SynthesizeOptionsVoiceEsEsEnriquevoiceConst     = "es-ES_EnriqueVoice"
	SynthesizeOptionsVoiceEsEsLaurav3voiceConst     = "es-ES_LauraV3Voice"
	SynthesizeOptionsVoiceEsEsLauravoiceConst       = "es-ES_LauraVoice"
	SynthesizeOptionsVoiceEsLaSofiav3voiceConst     = "es-LA_SofiaV3Voice"
	SynthesizeOptionsVoiceEsLaSofiavoiceConst       = "es-LA_SofiaVoice"
	SynthesizeOptionsVoiceEsUsSofiav3voiceConst     = "es-US_SofiaV3Voice"
	SynthesizeOptionsVoiceEsUsSofiavoiceConst       = "es-US_SofiaVoice"
	SynthesizeOptionsVoiceFrCaLouisev3voiceConst    = "fr-CA_LouiseV3Voice"
	SynthesizeOptionsVoiceFrFrNicolasv3voiceConst   = "fr-FR_NicolasV3Voice"
	SynthesizeOptionsVoiceFrFrReneev3voiceConst     = "fr-FR_ReneeV3Voice"
	SynthesizeOptionsVoiceFrFrReneevoiceConst       = "fr-FR_ReneeVoice"
	SynthesizeOptionsVoiceItItFrancescav3voiceConst = "it-IT_FrancescaV3Voice"
	SynthesizeOptionsVoiceItItFrancescavoiceConst   = "it-IT_FrancescaVoice"
	SynthesizeOptionsVoiceJaJpEmiv3voiceConst       = "ja-JP_EmiV3Voice"
	SynthesizeOptionsVoiceJaJpEmivoiceConst         = "ja-JP_EmiVoice"
	SynthesizeOptionsVoiceKoKrHyunjunvoiceConst     = "ko-KR_HyunjunVoice"
	SynthesizeOptionsVoiceKoKrSiwoovoiceConst       = "ko-KR_SiWooVoice"
	SynthesizeOptionsVoiceKoKrYoungmivoiceConst     = "ko-KR_YoungmiVoice"
	SynthesizeOptionsVoiceKoKrYunavoiceConst        = "ko-KR_YunaVoice"
	SynthesizeOptionsVoiceNlBeAdelevoiceConst       = "nl-BE_AdeleVoice"
	SynthesizeOptionsVoiceNlBeBramvoiceConst        = "nl-BE_BramVoice"
	SynthesizeOptionsVoiceNlNlEmmavoiceConst        = "nl-NL_EmmaVoice"
	SynthesizeOptionsVoiceNlNlLiamvoiceConst        = "nl-NL_LiamVoice"
	SynthesizeOptionsVoicePtBrIsabelav3voiceConst   = "pt-BR_IsabelaV3Voice"
	SynthesizeOptionsVoicePtBrIsabelavoiceConst     = "pt-BR_IsabelaVoice"
	SynthesizeOptionsVoiceSvSeIngridvoiceConst      = "sv-SE_IngridVoice"
	SynthesizeOptionsVoiceZhCnLinavoiceConst        = "zh-CN_LiNaVoice"
	SynthesizeOptionsVoiceZhCnWangweivoiceConst     = "zh-CN_WangWeiVoice"
	SynthesizeOptionsVoiceZhCnZhangjingvoiceConst   = "zh-CN_ZhangJingVoice"
)

// NewSynthesizeOptions : Instantiate SynthesizeOptions
func (*TextToSpeechV1) NewSynthesizeOptions(text string) *SynthesizeOptions {
	return &SynthesizeOptions{
		Text: core.StringPtr(text),
	}
}

// SetText : Allow user to set Text
func (_options *SynthesizeOptions) SetText(text string) *SynthesizeOptions {
	_options.Text = core.StringPtr(text)
	return _options
}

// SetAccept : Allow user to set Accept
func (_options *SynthesizeOptions) SetAccept(accept string) *SynthesizeOptions {
	_options.Accept = core.StringPtr(accept)
	return _options
}

// SetVoice : Allow user to set Voice
func (_options *SynthesizeOptions) SetVoice(voice string) *SynthesizeOptions {
	_options.Voice = core.StringPtr(voice)
	return _options
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *SynthesizeOptions) SetCustomizationID(customizationID string) *SynthesizeOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *SynthesizeOptions) SetHeaders(param map[string]string) *SynthesizeOptions {
	options.Headers = param
	return options
}

// Translation : Information about the translation for the specified text.
type Translation struct {
	// The phonetic or sounds-like translation for the word. A phonetic translation is based on the SSML format for
	// representing the phonetic string of a word either as an IPA translation or as an IBM SPR translation. The Arabic,
	// Chinese, Dutch, Australian English, and Korean languages support only IPA. A sounds-like is one or more words that,
	// when combined, sound like the word.
	Translation *string `json:"translation" validate:"required"`

	// **Japanese only.** The part of speech for the word. The service uses the value to produce the correct intonation for
	// the word. You can create only a single entry, with or without a single part of speech, for any word; you cannot
	// create multiple entries with different parts of speech for the same word. For more information, see [Working with
	// Japanese entries](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-rules#jaNotes).
	PartOfSpeech *string `json:"part_of_speech,omitempty"`
}

// Constants associated with the Translation.PartOfSpeech property.
// **Japanese only.** The part of speech for the word. The service uses the value to produce the correct intonation for
// the word. You can create only a single entry, with or without a single part of speech, for any word; you cannot
// create multiple entries with different parts of speech for the same word. For more information, see [Working with
// Japanese entries](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-rules#jaNotes).
const (
	TranslationPartOfSpeechDosiConst = "Dosi"
	TranslationPartOfSpeechFukuConst = "Fuku"
	TranslationPartOfSpeechGobiConst = "Gobi"
	TranslationPartOfSpeechHokaConst = "Hoka"
	TranslationPartOfSpeechJodoConst = "Jodo"
	TranslationPartOfSpeechJosiConst = "Josi"
	TranslationPartOfSpeechKatoConst = "Kato"
	TranslationPartOfSpeechKedoConst = "Kedo"
	TranslationPartOfSpeechKeyoConst = "Keyo"
	TranslationPartOfSpeechKigoConst = "Kigo"
	TranslationPartOfSpeechKoyuConst = "Koyu"
	TranslationPartOfSpeechMesiConst = "Mesi"
	TranslationPartOfSpeechRetaConst = "Reta"
	TranslationPartOfSpeechStbiConst = "Stbi"
	TranslationPartOfSpeechSttoConst = "Stto"
	TranslationPartOfSpeechStzoConst = "Stzo"
	TranslationPartOfSpeechSujiConst = "Suji"
)

// NewTranslation : Instantiate Translation (Generic Model Constructor)
func (*TextToSpeechV1) NewTranslation(translation string) (_model *Translation, err error) {
	_model = &Translation{
		Translation: core.StringPtr(translation),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalTranslation unmarshals an instance of Translation from the specified map of raw messages.
func UnmarshalTranslation(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Translation)
	err = core.UnmarshalPrimitive(m, "translation", &obj.Translation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "part_of_speech", &obj.PartOfSpeech)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// UpdateCustomModelOptions : The UpdateCustomModel options.
type UpdateCustomModelOptions struct {
	// The customization ID (GUID) of the custom model. You must make the request with credentials for the instance of the
	// service that owns the custom model.
	CustomizationID *string `json:"customization_id" validate:"required,ne="`

	// A new name for the custom model.
	Name *string `json:"name,omitempty"`

	// A new description for the custom model.
	Description *string `json:"description,omitempty"`

	// An array of `Word` objects that provides the words and their translations that are to be added or updated for the
	// custom model. Pass an empty array to make no additions or updates.
	Words []Word `json:"words,omitempty"`

	// Allows users to set headers on API requests
	Headers map[string]string
}

// NewUpdateCustomModelOptions : Instantiate UpdateCustomModelOptions
func (*TextToSpeechV1) NewUpdateCustomModelOptions(customizationID string) *UpdateCustomModelOptions {
	return &UpdateCustomModelOptions{
		CustomizationID: core.StringPtr(customizationID),
	}
}

// SetCustomizationID : Allow user to set CustomizationID
func (_options *UpdateCustomModelOptions) SetCustomizationID(customizationID string) *UpdateCustomModelOptions {
	_options.CustomizationID = core.StringPtr(customizationID)
	return _options
}

// SetName : Allow user to set Name
func (_options *UpdateCustomModelOptions) SetName(name string) *UpdateCustomModelOptions {
	_options.Name = core.StringPtr(name)
	return _options
}

// SetDescription : Allow user to set Description
func (_options *UpdateCustomModelOptions) SetDescription(description string) *UpdateCustomModelOptions {
	_options.Description = core.StringPtr(description)
	return _options
}

// SetWords : Allow user to set Words
func (_options *UpdateCustomModelOptions) SetWords(words []Word) *UpdateCustomModelOptions {
	_options.Words = words
	return _options
}

// SetHeaders : Allow user to set Headers
func (options *UpdateCustomModelOptions) SetHeaders(param map[string]string) *UpdateCustomModelOptions {
	options.Headers = param
	return options
}

// Voice : Information about an available voice.
type Voice struct {
	// The URI of the voice.
	URL *string `json:"url" validate:"required"`

	// The gender of the voice: `male` or `female`.
	Gender *string `json:"gender" validate:"required"`

	// The name of the voice. Use this as the voice identifier in all requests.
	Name *string `json:"name" validate:"required"`

	// The language and region of the voice (for example, `en-US`).
	Language *string `json:"language" validate:"required"`

	// A textual description of the voice.
	Description *string `json:"description" validate:"required"`

	// If `true`, the voice can be customized; if `false`, the voice cannot be customized. (Same as `custom_pronunciation`;
	// maintained for backward compatibility.).
	Customizable *bool `json:"customizable" validate:"required"`

	// Additional service features that are supported with the voice.
	SupportedFeatures *SupportedFeatures `json:"supported_features" validate:"required"`

	// Returns information about a specified custom model. This field is returned only by the [Get a voice](#getvoice)
	// method and only when you specify the customization ID of a custom model.
	Customization *CustomModel `json:"customization,omitempty"`
}

// UnmarshalVoice unmarshals an instance of Voice from the specified map of raw messages.
func UnmarshalVoice(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Voice)
	err = core.UnmarshalPrimitive(m, "url", &obj.URL)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "gender", &obj.Gender)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "name", &obj.Name)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "language", &obj.Language)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "description", &obj.Description)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "customizable", &obj.Customizable)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "supported_features", &obj.SupportedFeatures, UnmarshalSupportedFeatures)
	if err != nil {
		return
	}
	err = core.UnmarshalModel(m, "customization", &obj.Customization, UnmarshalCustomModel)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Voices : Information about all available voices.
type Voices struct {
	// A list of available voices.
	Voices []Voice `json:"voices" validate:"required"`
}

// UnmarshalVoices unmarshals an instance of Voices from the specified map of raw messages.
func UnmarshalVoices(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Voices)
	err = core.UnmarshalModel(m, "voices", &obj.Voices, UnmarshalVoice)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Word : Information about a word for the custom model.
type Word struct {
	// The word for the custom model. The maximum length of a word is 49 characters.
	Word *string `json:"word" validate:"required"`

	// The phonetic or sounds-like translation for the word. A phonetic translation is based on the SSML format for
	// representing the phonetic string of a word either as an IPA or IBM SPR translation. The Arabic, Chinese, Dutch,
	// Australian English, and Korean languages support only IPA. A sounds-like translation consists of one or more words
	// that, when combined, sound like the word. The maximum length of a translation is 499 characters.
	Translation *string `json:"translation" validate:"required"`

	// **Japanese only.** The part of speech for the word. The service uses the value to produce the correct intonation for
	// the word. You can create only a single entry, with or without a single part of speech, for any word; you cannot
	// create multiple entries with different parts of speech for the same word. For more information, see [Working with
	// Japanese entries](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-rules#jaNotes).
	PartOfSpeech *string `json:"part_of_speech,omitempty"`
}

// Constants associated with the Word.PartOfSpeech property.
// **Japanese only.** The part of speech for the word. The service uses the value to produce the correct intonation for
// the word. You can create only a single entry, with or without a single part of speech, for any word; you cannot
// create multiple entries with different parts of speech for the same word. For more information, see [Working with
// Japanese entries](https://cloud.ibm.com/docs/text-to-speech?topic=text-to-speech-rules#jaNotes).
const (
	WordPartOfSpeechDosiConst = "Dosi"
	WordPartOfSpeechFukuConst = "Fuku"
	WordPartOfSpeechGobiConst = "Gobi"
	WordPartOfSpeechHokaConst = "Hoka"
	WordPartOfSpeechJodoConst = "Jodo"
	WordPartOfSpeechJosiConst = "Josi"
	WordPartOfSpeechKatoConst = "Kato"
	WordPartOfSpeechKedoConst = "Kedo"
	WordPartOfSpeechKeyoConst = "Keyo"
	WordPartOfSpeechKigoConst = "Kigo"
	WordPartOfSpeechKoyuConst = "Koyu"
	WordPartOfSpeechMesiConst = "Mesi"
	WordPartOfSpeechRetaConst = "Reta"
	WordPartOfSpeechStbiConst = "Stbi"
	WordPartOfSpeechSttoConst = "Stto"
	WordPartOfSpeechStzoConst = "Stzo"
	WordPartOfSpeechSujiConst = "Suji"
)

// NewWord : Instantiate Word (Generic Model Constructor)
func (*TextToSpeechV1) NewWord(word string, translation string) (_model *Word, err error) {
	_model = &Word{
		Word:        core.StringPtr(word),
		Translation: core.StringPtr(translation),
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalWord unmarshals an instance of Word from the specified map of raw messages.
func UnmarshalWord(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Word)
	err = core.UnmarshalPrimitive(m, "word", &obj.Word)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "translation", &obj.Translation)
	if err != nil {
		return
	}
	err = core.UnmarshalPrimitive(m, "part_of_speech", &obj.PartOfSpeech)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}

// Words : For the [Add custom words](#addwords) method, one or more words that are to be added or updated for the custom model
// and the translation for each specified word.
//
// For the [List custom words](#listwords) method, the words and their translations from the custom model.
type Words struct {
	// The [Add custom words](#addwords) method accepts an array of `Word` objects. Each object provides a word that is to
	// be added or updated for the custom model and the word's translation.
	//
	// The [List custom words](#listwords) method returns an array of `Word` objects. Each object shows a word and its
	// translation from the custom model. The words are listed in alphabetical order, with uppercase letters listed before
	// lowercase letters. The array is empty if the custom model contains no words.
	Words []Word `json:"words" validate:"required"`
}

// NewWords : Instantiate Words (Generic Model Constructor)
func (*TextToSpeechV1) NewWords(words []Word) (_model *Words, err error) {
	_model = &Words{
		Words: words,
	}
	err = core.ValidateStruct(_model, "required parameters")
	return
}

// UnmarshalWords unmarshals an instance of Words from the specified map of raw messages.
func UnmarshalWords(m map[string]json.RawMessage, result interface{}) (err error) {
	obj := new(Words)
	err = core.UnmarshalModel(m, "words", &obj.Words, UnmarshalWord)
	if err != nil {
		return
	}
	reflect.ValueOf(result).Elem().Set(reflect.ValueOf(obj))
	return
}
