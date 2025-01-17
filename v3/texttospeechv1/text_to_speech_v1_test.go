/**
 * (C) Copyright IBM Corp. 2018, 2022.
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

package texttospeechv1_test

import (
	"bytes"
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"os"
	"time"

	"github.com/IBM/go-sdk-core/v5/core"
	"github.com/go-openapi/strfmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"github.com/watson-developer-cloud/go-sdk/v3/texttospeechv1"
)

var _ = Describe(`TextToSpeechV1`, func() {
	var testServer *httptest.Server
	Describe(`Service constructor tests`, func() {
		It(`Instantiate service client`, func() {
			textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
				Authenticator: &core.NoAuthAuthenticator{},
			})
			Expect(textToSpeechService).ToNot(BeNil())
			Expect(serviceErr).To(BeNil())
			Expect(textToSpeechService.Service.IsSSLDisabled()).To(BeFalse())
			textToSpeechService.DisableSSLVerification()
			Expect(textToSpeechService.Service.IsSSLDisabled()).To(BeTrue())
		})
		It(`Instantiate service client with error: Invalid URL`, func() {
			textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
				URL: "{BAD_URL_STRING",
			})
			Expect(textToSpeechService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
		It(`Instantiate service client with error: Invalid Auth`, func() {
			textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
				URL: "https://texttospeechv1/api",
				Authenticator: &core.BasicAuthenticator{
					Username: "",
					Password: "",
				},
			})
			Expect(textToSpeechService).To(BeNil())
			Expect(serviceErr).ToNot(BeNil())
		})
	})
	Describe(`Service constructor tests using external config`, func() {
		Context(`Using external config, construct service client instances`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"TEXT_TO_SPEECH_URL":       "https://texttospeechv1/api",
				"TEXT_TO_SPEECH_AUTH_TYPE": "noauth",
			}

			It(`Create service client using external config successfully`, func() {
				SetTestEnvironment(testEnvironment)
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{})
				Expect(textToSpeechService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				ClearTestEnvironment(testEnvironment)

				clone := textToSpeechService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != textToSpeechService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(textToSpeechService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(textToSpeechService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url from constructor successfully`, func() {
				SetTestEnvironment(testEnvironment)
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL: "https://testService/api",
				})
				Expect(textToSpeechService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := textToSpeechService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != textToSpeechService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(textToSpeechService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(textToSpeechService.Service.Options.Authenticator))
			})
			It(`Create service client using external config and set url programatically successfully`, func() {
				SetTestEnvironment(testEnvironment)
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{})
				err := textToSpeechService.SetServiceURL("https://testService/api")
				Expect(err).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService.Service.GetServiceURL()).To(Equal("https://testService/api"))
				ClearTestEnvironment(testEnvironment)

				clone := textToSpeechService.Clone()
				Expect(clone).ToNot(BeNil())
				Expect(clone.Service != textToSpeechService.Service).To(BeTrue())
				Expect(clone.GetServiceURL()).To(Equal(textToSpeechService.GetServiceURL()))
				Expect(clone.Service.Options.Authenticator).To(Equal(textToSpeechService.Service.Options.Authenticator))
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid Auth`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"TEXT_TO_SPEECH_URL":       "https://texttospeechv1/api",
				"TEXT_TO_SPEECH_AUTH_TYPE": "someOtherAuth",
			}

			SetTestEnvironment(testEnvironment)
			textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{})

			It(`Instantiate service client with error`, func() {
				Expect(textToSpeechService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
		Context(`Using external config, construct service client instances with error: Invalid URL`, func() {
			// Map containing environment variables used in testing.
			var testEnvironment = map[string]string{
				"TEXT_TO_SPEECH_AUTH_TYPE": "NOAuth",
			}

			SetTestEnvironment(testEnvironment)
			textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
				URL: "{BAD_URL_STRING",
			})

			It(`Instantiate service client with error`, func() {
				Expect(textToSpeechService).To(BeNil())
				Expect(serviceErr).ToNot(BeNil())
				ClearTestEnvironment(testEnvironment)
			})
		})
	})
	Describe(`Regional endpoint tests`, func() {
		It(`GetServiceURLForRegion(region string)`, func() {
			var url string
			var err error
			url, err = texttospeechv1.GetServiceURLForRegion("INVALID_REGION")
			Expect(url).To(BeEmpty())
			Expect(err).ToNot(BeNil())
			fmt.Fprintf(GinkgoWriter, "Expected error: %s\n", err.Error())
		})
	})
	Describe(`ListVoices(listVoicesOptions *ListVoicesOptions) - Operation response error`, func() {
		listVoicesPath := "/v1/voices"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listVoicesPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListVoices with error: Operation response processing error`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the ListVoicesOptions model
				listVoicesOptionsModel := new(texttospeechv1.ListVoicesOptions)
				listVoicesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := textToSpeechService.ListVoices(listVoicesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				textToSpeechService.EnableRetries(0, 0)
				result, response, operationErr = textToSpeechService.ListVoices(listVoicesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListVoices(listVoicesOptions *ListVoicesOptions)`, func() {
		listVoicesPath := "/v1/voices"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listVoicesPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"voices": [{"url": "URL", "gender": "Gender", "name": "Name", "language": "Language", "description": "Description", "customizable": true, "supported_features": {"custom_pronunciation": false, "voice_transformation": false}, "customization": {"customization_id": "CustomizationID", "name": "Name", "language": "Language", "owner": "Owner", "created": "Created", "last_modified": "LastModified", "description": "Description", "words": [{"word": "Word", "translation": "Translation", "part_of_speech": "Dosi"}], "prompts": [{"prompt": "Prompt", "prompt_id": "PromptID", "status": "Status", "error": "Error", "speaker_id": "SpeakerID"}]}}]}`)
				}))
			})
			It(`Invoke ListVoices successfully with retries`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())
				textToSpeechService.EnableRetries(0, 0)

				// Construct an instance of the ListVoicesOptions model
				listVoicesOptionsModel := new(texttospeechv1.ListVoicesOptions)
				listVoicesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := textToSpeechService.ListVoicesWithContext(ctx, listVoicesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				textToSpeechService.DisableRetries()
				result, response, operationErr := textToSpeechService.ListVoices(listVoicesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = textToSpeechService.ListVoicesWithContext(ctx, listVoicesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listVoicesPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"voices": [{"url": "URL", "gender": "Gender", "name": "Name", "language": "Language", "description": "Description", "customizable": true, "supported_features": {"custom_pronunciation": false, "voice_transformation": false}, "customization": {"customization_id": "CustomizationID", "name": "Name", "language": "Language", "owner": "Owner", "created": "Created", "last_modified": "LastModified", "description": "Description", "words": [{"word": "Word", "translation": "Translation", "part_of_speech": "Dosi"}], "prompts": [{"prompt": "Prompt", "prompt_id": "PromptID", "status": "Status", "error": "Error", "speaker_id": "SpeakerID"}]}}]}`)
				}))
			})
			It(`Invoke ListVoices successfully`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := textToSpeechService.ListVoices(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListVoicesOptions model
				listVoicesOptionsModel := new(texttospeechv1.ListVoicesOptions)
				listVoicesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = textToSpeechService.ListVoices(listVoicesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListVoices with error: Operation request error`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the ListVoicesOptions model
				listVoicesOptionsModel := new(texttospeechv1.ListVoicesOptions)
				listVoicesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := textToSpeechService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := textToSpeechService.ListVoices(listVoicesOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListVoices successfully`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the ListVoicesOptions model
				listVoicesOptionsModel := new(texttospeechv1.ListVoicesOptions)
				listVoicesOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := textToSpeechService.ListVoices(listVoicesOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetVoice(getVoiceOptions *GetVoiceOptions) - Operation response error`, func() {
		getVoicePath := "/v1/voices/ar-AR_OmarVoice"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getVoicePath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["customization_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetVoice with error: Operation response processing error`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the GetVoiceOptions model
				getVoiceOptionsModel := new(texttospeechv1.GetVoiceOptions)
				getVoiceOptionsModel.Voice = core.StringPtr("ar-AR_OmarVoice")
				getVoiceOptionsModel.CustomizationID = core.StringPtr("testString")
				getVoiceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := textToSpeechService.GetVoice(getVoiceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				textToSpeechService.EnableRetries(0, 0)
				result, response, operationErr = textToSpeechService.GetVoice(getVoiceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetVoice(getVoiceOptions *GetVoiceOptions)`, func() {
		getVoicePath := "/v1/voices/ar-AR_OmarVoice"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getVoicePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["customization_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"url": "URL", "gender": "Gender", "name": "Name", "language": "Language", "description": "Description", "customizable": true, "supported_features": {"custom_pronunciation": false, "voice_transformation": false}, "customization": {"customization_id": "CustomizationID", "name": "Name", "language": "Language", "owner": "Owner", "created": "Created", "last_modified": "LastModified", "description": "Description", "words": [{"word": "Word", "translation": "Translation", "part_of_speech": "Dosi"}], "prompts": [{"prompt": "Prompt", "prompt_id": "PromptID", "status": "Status", "error": "Error", "speaker_id": "SpeakerID"}]}}`)
				}))
			})
			It(`Invoke GetVoice successfully with retries`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())
				textToSpeechService.EnableRetries(0, 0)

				// Construct an instance of the GetVoiceOptions model
				getVoiceOptionsModel := new(texttospeechv1.GetVoiceOptions)
				getVoiceOptionsModel.Voice = core.StringPtr("ar-AR_OmarVoice")
				getVoiceOptionsModel.CustomizationID = core.StringPtr("testString")
				getVoiceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := textToSpeechService.GetVoiceWithContext(ctx, getVoiceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				textToSpeechService.DisableRetries()
				result, response, operationErr := textToSpeechService.GetVoice(getVoiceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = textToSpeechService.GetVoiceWithContext(ctx, getVoiceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getVoicePath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["customization_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"url": "URL", "gender": "Gender", "name": "Name", "language": "Language", "description": "Description", "customizable": true, "supported_features": {"custom_pronunciation": false, "voice_transformation": false}, "customization": {"customization_id": "CustomizationID", "name": "Name", "language": "Language", "owner": "Owner", "created": "Created", "last_modified": "LastModified", "description": "Description", "words": [{"word": "Word", "translation": "Translation", "part_of_speech": "Dosi"}], "prompts": [{"prompt": "Prompt", "prompt_id": "PromptID", "status": "Status", "error": "Error", "speaker_id": "SpeakerID"}]}}`)
				}))
			})
			It(`Invoke GetVoice successfully`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := textToSpeechService.GetVoice(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetVoiceOptions model
				getVoiceOptionsModel := new(texttospeechv1.GetVoiceOptions)
				getVoiceOptionsModel.Voice = core.StringPtr("ar-AR_OmarVoice")
				getVoiceOptionsModel.CustomizationID = core.StringPtr("testString")
				getVoiceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = textToSpeechService.GetVoice(getVoiceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetVoice with error: Operation validation and request error`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the GetVoiceOptions model
				getVoiceOptionsModel := new(texttospeechv1.GetVoiceOptions)
				getVoiceOptionsModel.Voice = core.StringPtr("ar-AR_OmarVoice")
				getVoiceOptionsModel.CustomizationID = core.StringPtr("testString")
				getVoiceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := textToSpeechService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := textToSpeechService.GetVoice(getVoiceOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetVoiceOptions model with no property values
				getVoiceOptionsModelNew := new(texttospeechv1.GetVoiceOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = textToSpeechService.GetVoice(getVoiceOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetVoice successfully`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the GetVoiceOptions model
				getVoiceOptionsModel := new(texttospeechv1.GetVoiceOptions)
				getVoiceOptionsModel.Voice = core.StringPtr("ar-AR_OmarVoice")
				getVoiceOptionsModel.CustomizationID = core.StringPtr("testString")
				getVoiceOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := textToSpeechService.GetVoice(getVoiceOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Synthesize(synthesizeOptions *SynthesizeOptions)`, func() {
		synthesizePath := "/v1/synthesize"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(synthesizePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "audio/ogg;codecs=opus")))
					Expect(req.URL.Query()["voice"]).To(Equal([]string{"en-US_MichaelV3Voice"}))
					Expect(req.URL.Query()["customization_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "audio/basic")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke Synthesize successfully with retries`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())
				textToSpeechService.EnableRetries(0, 0)

				// Construct an instance of the SynthesizeOptions model
				synthesizeOptionsModel := new(texttospeechv1.SynthesizeOptions)
				synthesizeOptionsModel.Text = core.StringPtr("testString")
				synthesizeOptionsModel.Accept = core.StringPtr("audio/ogg;codecs=opus")
				synthesizeOptionsModel.Voice = core.StringPtr("en-US_MichaelV3Voice")
				synthesizeOptionsModel.CustomizationID = core.StringPtr("testString")
				synthesizeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := textToSpeechService.SynthesizeWithContext(ctx, synthesizeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				textToSpeechService.DisableRetries()
				result, response, operationErr := textToSpeechService.Synthesize(synthesizeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = textToSpeechService.SynthesizeWithContext(ctx, synthesizeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(synthesizePath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.Header["Accept"]).ToNot(BeNil())
					Expect(req.Header["Accept"][0]).To(Equal(fmt.Sprintf("%v", "audio/ogg;codecs=opus")))
					Expect(req.URL.Query()["voice"]).To(Equal([]string{"en-US_MichaelV3Voice"}))
					Expect(req.URL.Query()["customization_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "audio/basic")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `This is a mock binary response.`)
				}))
			})
			It(`Invoke Synthesize successfully`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := textToSpeechService.Synthesize(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the SynthesizeOptions model
				synthesizeOptionsModel := new(texttospeechv1.SynthesizeOptions)
				synthesizeOptionsModel.Text = core.StringPtr("testString")
				synthesizeOptionsModel.Accept = core.StringPtr("audio/ogg;codecs=opus")
				synthesizeOptionsModel.Voice = core.StringPtr("en-US_MichaelV3Voice")
				synthesizeOptionsModel.CustomizationID = core.StringPtr("testString")
				synthesizeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = textToSpeechService.Synthesize(synthesizeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke Synthesize with error: Operation validation and request error`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the SynthesizeOptions model
				synthesizeOptionsModel := new(texttospeechv1.SynthesizeOptions)
				synthesizeOptionsModel.Text = core.StringPtr("testString")
				synthesizeOptionsModel.Accept = core.StringPtr("audio/ogg;codecs=opus")
				synthesizeOptionsModel.Voice = core.StringPtr("en-US_MichaelV3Voice")
				synthesizeOptionsModel.CustomizationID = core.StringPtr("testString")
				synthesizeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := textToSpeechService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := textToSpeechService.Synthesize(synthesizeOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the SynthesizeOptions model with no property values
				synthesizeOptionsModelNew := new(texttospeechv1.SynthesizeOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = textToSpeechService.Synthesize(synthesizeOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke Synthesize successfully`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the SynthesizeOptions model
				synthesizeOptionsModel := new(texttospeechv1.SynthesizeOptions)
				synthesizeOptionsModel.Text = core.StringPtr("testString")
				synthesizeOptionsModel.Accept = core.StringPtr("audio/ogg;codecs=opus")
				synthesizeOptionsModel.Voice = core.StringPtr("en-US_MichaelV3Voice")
				synthesizeOptionsModel.CustomizationID = core.StringPtr("testString")
				synthesizeOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := textToSpeechService.Synthesize(synthesizeOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify empty byte buffer.
				Expect(result).ToNot(BeNil())
				buffer, operationErr := ioutil.ReadAll(result)
				Expect(operationErr).To(BeNil())
				Expect(buffer).ToNot(BeNil())
				Expect(len(buffer)).To(Equal(0))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPronunciation(getPronunciationOptions *GetPronunciationOptions) - Operation response error`, func() {
		getPronunciationPath := "/v1/pronunciation"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPronunciationPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["text"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["voice"]).To(Equal([]string{"en-US_MichaelV3Voice"}))
					Expect(req.URL.Query()["format"]).To(Equal([]string{"ipa"}))
					Expect(req.URL.Query()["customization_id"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetPronunciation with error: Operation response processing error`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the GetPronunciationOptions model
				getPronunciationOptionsModel := new(texttospeechv1.GetPronunciationOptions)
				getPronunciationOptionsModel.Text = core.StringPtr("testString")
				getPronunciationOptionsModel.Voice = core.StringPtr("en-US_MichaelV3Voice")
				getPronunciationOptionsModel.Format = core.StringPtr("ipa")
				getPronunciationOptionsModel.CustomizationID = core.StringPtr("testString")
				getPronunciationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := textToSpeechService.GetPronunciation(getPronunciationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				textToSpeechService.EnableRetries(0, 0)
				result, response, operationErr = textToSpeechService.GetPronunciation(getPronunciationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetPronunciation(getPronunciationOptions *GetPronunciationOptions)`, func() {
		getPronunciationPath := "/v1/pronunciation"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPronunciationPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["text"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["voice"]).To(Equal([]string{"en-US_MichaelV3Voice"}))
					Expect(req.URL.Query()["format"]).To(Equal([]string{"ipa"}))
					Expect(req.URL.Query()["customization_id"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"pronunciation": "Pronunciation"}`)
				}))
			})
			It(`Invoke GetPronunciation successfully with retries`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())
				textToSpeechService.EnableRetries(0, 0)

				// Construct an instance of the GetPronunciationOptions model
				getPronunciationOptionsModel := new(texttospeechv1.GetPronunciationOptions)
				getPronunciationOptionsModel.Text = core.StringPtr("testString")
				getPronunciationOptionsModel.Voice = core.StringPtr("en-US_MichaelV3Voice")
				getPronunciationOptionsModel.Format = core.StringPtr("ipa")
				getPronunciationOptionsModel.CustomizationID = core.StringPtr("testString")
				getPronunciationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := textToSpeechService.GetPronunciationWithContext(ctx, getPronunciationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				textToSpeechService.DisableRetries()
				result, response, operationErr := textToSpeechService.GetPronunciation(getPronunciationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = textToSpeechService.GetPronunciationWithContext(ctx, getPronunciationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getPronunciationPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["text"]).To(Equal([]string{"testString"}))
					Expect(req.URL.Query()["voice"]).To(Equal([]string{"en-US_MichaelV3Voice"}))
					Expect(req.URL.Query()["format"]).To(Equal([]string{"ipa"}))
					Expect(req.URL.Query()["customization_id"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"pronunciation": "Pronunciation"}`)
				}))
			})
			It(`Invoke GetPronunciation successfully`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := textToSpeechService.GetPronunciation(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetPronunciationOptions model
				getPronunciationOptionsModel := new(texttospeechv1.GetPronunciationOptions)
				getPronunciationOptionsModel.Text = core.StringPtr("testString")
				getPronunciationOptionsModel.Voice = core.StringPtr("en-US_MichaelV3Voice")
				getPronunciationOptionsModel.Format = core.StringPtr("ipa")
				getPronunciationOptionsModel.CustomizationID = core.StringPtr("testString")
				getPronunciationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = textToSpeechService.GetPronunciation(getPronunciationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetPronunciation with error: Operation validation and request error`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the GetPronunciationOptions model
				getPronunciationOptionsModel := new(texttospeechv1.GetPronunciationOptions)
				getPronunciationOptionsModel.Text = core.StringPtr("testString")
				getPronunciationOptionsModel.Voice = core.StringPtr("en-US_MichaelV3Voice")
				getPronunciationOptionsModel.Format = core.StringPtr("ipa")
				getPronunciationOptionsModel.CustomizationID = core.StringPtr("testString")
				getPronunciationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := textToSpeechService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := textToSpeechService.GetPronunciation(getPronunciationOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetPronunciationOptions model with no property values
				getPronunciationOptionsModelNew := new(texttospeechv1.GetPronunciationOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = textToSpeechService.GetPronunciation(getPronunciationOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetPronunciation successfully`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the GetPronunciationOptions model
				getPronunciationOptionsModel := new(texttospeechv1.GetPronunciationOptions)
				getPronunciationOptionsModel.Text = core.StringPtr("testString")
				getPronunciationOptionsModel.Voice = core.StringPtr("en-US_MichaelV3Voice")
				getPronunciationOptionsModel.Format = core.StringPtr("ipa")
				getPronunciationOptionsModel.CustomizationID = core.StringPtr("testString")
				getPronunciationOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := textToSpeechService.GetPronunciation(getPronunciationOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateCustomModel(createCustomModelOptions *CreateCustomModelOptions) - Operation response error`, func() {
		createCustomModelPath := "/v1/customizations"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCustomModelPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateCustomModel with error: Operation response processing error`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the CreateCustomModelOptions model
				createCustomModelOptionsModel := new(texttospeechv1.CreateCustomModelOptions)
				createCustomModelOptionsModel.Name = core.StringPtr("testString")
				createCustomModelOptionsModel.Language = core.StringPtr("en-US")
				createCustomModelOptionsModel.Description = core.StringPtr("testString")
				createCustomModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := textToSpeechService.CreateCustomModel(createCustomModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				textToSpeechService.EnableRetries(0, 0)
				result, response, operationErr = textToSpeechService.CreateCustomModel(createCustomModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateCustomModel(createCustomModelOptions *CreateCustomModelOptions)`, func() {
		createCustomModelPath := "/v1/customizations"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCustomModelPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"customization_id": "CustomizationID", "name": "Name", "language": "Language", "owner": "Owner", "created": "Created", "last_modified": "LastModified", "description": "Description", "words": [{"word": "Word", "translation": "Translation", "part_of_speech": "Dosi"}], "prompts": [{"prompt": "Prompt", "prompt_id": "PromptID", "status": "Status", "error": "Error", "speaker_id": "SpeakerID"}]}`)
				}))
			})
			It(`Invoke CreateCustomModel successfully with retries`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())
				textToSpeechService.EnableRetries(0, 0)

				// Construct an instance of the CreateCustomModelOptions model
				createCustomModelOptionsModel := new(texttospeechv1.CreateCustomModelOptions)
				createCustomModelOptionsModel.Name = core.StringPtr("testString")
				createCustomModelOptionsModel.Language = core.StringPtr("en-US")
				createCustomModelOptionsModel.Description = core.StringPtr("testString")
				createCustomModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := textToSpeechService.CreateCustomModelWithContext(ctx, createCustomModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				textToSpeechService.DisableRetries()
				result, response, operationErr := textToSpeechService.CreateCustomModel(createCustomModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = textToSpeechService.CreateCustomModelWithContext(ctx, createCustomModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createCustomModelPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"customization_id": "CustomizationID", "name": "Name", "language": "Language", "owner": "Owner", "created": "Created", "last_modified": "LastModified", "description": "Description", "words": [{"word": "Word", "translation": "Translation", "part_of_speech": "Dosi"}], "prompts": [{"prompt": "Prompt", "prompt_id": "PromptID", "status": "Status", "error": "Error", "speaker_id": "SpeakerID"}]}`)
				}))
			})
			It(`Invoke CreateCustomModel successfully`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := textToSpeechService.CreateCustomModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateCustomModelOptions model
				createCustomModelOptionsModel := new(texttospeechv1.CreateCustomModelOptions)
				createCustomModelOptionsModel.Name = core.StringPtr("testString")
				createCustomModelOptionsModel.Language = core.StringPtr("en-US")
				createCustomModelOptionsModel.Description = core.StringPtr("testString")
				createCustomModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = textToSpeechService.CreateCustomModel(createCustomModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateCustomModel with error: Operation validation and request error`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the CreateCustomModelOptions model
				createCustomModelOptionsModel := new(texttospeechv1.CreateCustomModelOptions)
				createCustomModelOptionsModel.Name = core.StringPtr("testString")
				createCustomModelOptionsModel.Language = core.StringPtr("en-US")
				createCustomModelOptionsModel.Description = core.StringPtr("testString")
				createCustomModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := textToSpeechService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := textToSpeechService.CreateCustomModel(createCustomModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateCustomModelOptions model with no property values
				createCustomModelOptionsModelNew := new(texttospeechv1.CreateCustomModelOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = textToSpeechService.CreateCustomModel(createCustomModelOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateCustomModel successfully`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the CreateCustomModelOptions model
				createCustomModelOptionsModel := new(texttospeechv1.CreateCustomModelOptions)
				createCustomModelOptionsModel.Name = core.StringPtr("testString")
				createCustomModelOptionsModel.Language = core.StringPtr("en-US")
				createCustomModelOptionsModel.Description = core.StringPtr("testString")
				createCustomModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := textToSpeechService.CreateCustomModel(createCustomModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListCustomModels(listCustomModelsOptions *ListCustomModelsOptions) - Operation response error`, func() {
		listCustomModelsPath := "/v1/customizations"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCustomModelsPath))
					Expect(req.Method).To(Equal("GET"))
					Expect(req.URL.Query()["language"]).To(Equal([]string{"ar-MS"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListCustomModels with error: Operation response processing error`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the ListCustomModelsOptions model
				listCustomModelsOptionsModel := new(texttospeechv1.ListCustomModelsOptions)
				listCustomModelsOptionsModel.Language = core.StringPtr("ar-MS")
				listCustomModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := textToSpeechService.ListCustomModels(listCustomModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				textToSpeechService.EnableRetries(0, 0)
				result, response, operationErr = textToSpeechService.ListCustomModels(listCustomModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListCustomModels(listCustomModelsOptions *ListCustomModelsOptions)`, func() {
		listCustomModelsPath := "/v1/customizations"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCustomModelsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["language"]).To(Equal([]string{"ar-MS"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"customizations": [{"customization_id": "CustomizationID", "name": "Name", "language": "Language", "owner": "Owner", "created": "Created", "last_modified": "LastModified", "description": "Description", "words": [{"word": "Word", "translation": "Translation", "part_of_speech": "Dosi"}], "prompts": [{"prompt": "Prompt", "prompt_id": "PromptID", "status": "Status", "error": "Error", "speaker_id": "SpeakerID"}]}]}`)
				}))
			})
			It(`Invoke ListCustomModels successfully with retries`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())
				textToSpeechService.EnableRetries(0, 0)

				// Construct an instance of the ListCustomModelsOptions model
				listCustomModelsOptionsModel := new(texttospeechv1.ListCustomModelsOptions)
				listCustomModelsOptionsModel.Language = core.StringPtr("ar-MS")
				listCustomModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := textToSpeechService.ListCustomModelsWithContext(ctx, listCustomModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				textToSpeechService.DisableRetries()
				result, response, operationErr := textToSpeechService.ListCustomModels(listCustomModelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = textToSpeechService.ListCustomModelsWithContext(ctx, listCustomModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCustomModelsPath))
					Expect(req.Method).To(Equal("GET"))

					Expect(req.URL.Query()["language"]).To(Equal([]string{"ar-MS"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"customizations": [{"customization_id": "CustomizationID", "name": "Name", "language": "Language", "owner": "Owner", "created": "Created", "last_modified": "LastModified", "description": "Description", "words": [{"word": "Word", "translation": "Translation", "part_of_speech": "Dosi"}], "prompts": [{"prompt": "Prompt", "prompt_id": "PromptID", "status": "Status", "error": "Error", "speaker_id": "SpeakerID"}]}]}`)
				}))
			})
			It(`Invoke ListCustomModels successfully`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := textToSpeechService.ListCustomModels(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListCustomModelsOptions model
				listCustomModelsOptionsModel := new(texttospeechv1.ListCustomModelsOptions)
				listCustomModelsOptionsModel.Language = core.StringPtr("ar-MS")
				listCustomModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = textToSpeechService.ListCustomModels(listCustomModelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListCustomModels with error: Operation request error`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the ListCustomModelsOptions model
				listCustomModelsOptionsModel := new(texttospeechv1.ListCustomModelsOptions)
				listCustomModelsOptionsModel.Language = core.StringPtr("ar-MS")
				listCustomModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := textToSpeechService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := textToSpeechService.ListCustomModels(listCustomModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListCustomModels successfully`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the ListCustomModelsOptions model
				listCustomModelsOptionsModel := new(texttospeechv1.ListCustomModelsOptions)
				listCustomModelsOptionsModel.Language = core.StringPtr("ar-MS")
				listCustomModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := textToSpeechService.ListCustomModels(listCustomModelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`UpdateCustomModel(updateCustomModelOptions *UpdateCustomModelOptions)`, func() {
		updateCustomModelPath := "/v1/customizations/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(updateCustomModelPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					res.WriteHeader(200)
				}))
			})
			It(`Invoke UpdateCustomModel successfully`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := textToSpeechService.UpdateCustomModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the Word model
				wordModel := new(texttospeechv1.Word)
				wordModel.Word = core.StringPtr("testString")
				wordModel.Translation = core.StringPtr("testString")
				wordModel.PartOfSpeech = core.StringPtr("Dosi")

				// Construct an instance of the UpdateCustomModelOptions model
				updateCustomModelOptionsModel := new(texttospeechv1.UpdateCustomModelOptions)
				updateCustomModelOptionsModel.CustomizationID = core.StringPtr("testString")
				updateCustomModelOptionsModel.Name = core.StringPtr("testString")
				updateCustomModelOptionsModel.Description = core.StringPtr("testString")
				updateCustomModelOptionsModel.Words = []texttospeechv1.Word{*wordModel}
				updateCustomModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = textToSpeechService.UpdateCustomModel(updateCustomModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke UpdateCustomModel with error: Operation validation and request error`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the Word model
				wordModel := new(texttospeechv1.Word)
				wordModel.Word = core.StringPtr("testString")
				wordModel.Translation = core.StringPtr("testString")
				wordModel.PartOfSpeech = core.StringPtr("Dosi")

				// Construct an instance of the UpdateCustomModelOptions model
				updateCustomModelOptionsModel := new(texttospeechv1.UpdateCustomModelOptions)
				updateCustomModelOptionsModel.CustomizationID = core.StringPtr("testString")
				updateCustomModelOptionsModel.Name = core.StringPtr("testString")
				updateCustomModelOptionsModel.Description = core.StringPtr("testString")
				updateCustomModelOptionsModel.Words = []texttospeechv1.Word{*wordModel}
				updateCustomModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := textToSpeechService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := textToSpeechService.UpdateCustomModel(updateCustomModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the UpdateCustomModelOptions model with no property values
				updateCustomModelOptionsModelNew := new(texttospeechv1.UpdateCustomModelOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = textToSpeechService.UpdateCustomModel(updateCustomModelOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCustomModel(getCustomModelOptions *GetCustomModelOptions) - Operation response error`, func() {
		getCustomModelPath := "/v1/customizations/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCustomModelPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCustomModel with error: Operation response processing error`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the GetCustomModelOptions model
				getCustomModelOptionsModel := new(texttospeechv1.GetCustomModelOptions)
				getCustomModelOptionsModel.CustomizationID = core.StringPtr("testString")
				getCustomModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := textToSpeechService.GetCustomModel(getCustomModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				textToSpeechService.EnableRetries(0, 0)
				result, response, operationErr = textToSpeechService.GetCustomModel(getCustomModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCustomModel(getCustomModelOptions *GetCustomModelOptions)`, func() {
		getCustomModelPath := "/v1/customizations/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCustomModelPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"customization_id": "CustomizationID", "name": "Name", "language": "Language", "owner": "Owner", "created": "Created", "last_modified": "LastModified", "description": "Description", "words": [{"word": "Word", "translation": "Translation", "part_of_speech": "Dosi"}], "prompts": [{"prompt": "Prompt", "prompt_id": "PromptID", "status": "Status", "error": "Error", "speaker_id": "SpeakerID"}]}`)
				}))
			})
			It(`Invoke GetCustomModel successfully with retries`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())
				textToSpeechService.EnableRetries(0, 0)

				// Construct an instance of the GetCustomModelOptions model
				getCustomModelOptionsModel := new(texttospeechv1.GetCustomModelOptions)
				getCustomModelOptionsModel.CustomizationID = core.StringPtr("testString")
				getCustomModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := textToSpeechService.GetCustomModelWithContext(ctx, getCustomModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				textToSpeechService.DisableRetries()
				result, response, operationErr := textToSpeechService.GetCustomModel(getCustomModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = textToSpeechService.GetCustomModelWithContext(ctx, getCustomModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCustomModelPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"customization_id": "CustomizationID", "name": "Name", "language": "Language", "owner": "Owner", "created": "Created", "last_modified": "LastModified", "description": "Description", "words": [{"word": "Word", "translation": "Translation", "part_of_speech": "Dosi"}], "prompts": [{"prompt": "Prompt", "prompt_id": "PromptID", "status": "Status", "error": "Error", "speaker_id": "SpeakerID"}]}`)
				}))
			})
			It(`Invoke GetCustomModel successfully`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := textToSpeechService.GetCustomModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCustomModelOptions model
				getCustomModelOptionsModel := new(texttospeechv1.GetCustomModelOptions)
				getCustomModelOptionsModel.CustomizationID = core.StringPtr("testString")
				getCustomModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = textToSpeechService.GetCustomModel(getCustomModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetCustomModel with error: Operation validation and request error`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the GetCustomModelOptions model
				getCustomModelOptionsModel := new(texttospeechv1.GetCustomModelOptions)
				getCustomModelOptionsModel.CustomizationID = core.StringPtr("testString")
				getCustomModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := textToSpeechService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := textToSpeechService.GetCustomModel(getCustomModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetCustomModelOptions model with no property values
				getCustomModelOptionsModelNew := new(texttospeechv1.GetCustomModelOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = textToSpeechService.GetCustomModel(getCustomModelOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetCustomModel successfully`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the GetCustomModelOptions model
				getCustomModelOptionsModel := new(texttospeechv1.GetCustomModelOptions)
				getCustomModelOptionsModel.CustomizationID = core.StringPtr("testString")
				getCustomModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := textToSpeechService.GetCustomModel(getCustomModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteCustomModel(deleteCustomModelOptions *DeleteCustomModelOptions)`, func() {
		deleteCustomModelPath := "/v1/customizations/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteCustomModelPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteCustomModel successfully`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := textToSpeechService.DeleteCustomModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteCustomModelOptions model
				deleteCustomModelOptionsModel := new(texttospeechv1.DeleteCustomModelOptions)
				deleteCustomModelOptionsModel.CustomizationID = core.StringPtr("testString")
				deleteCustomModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = textToSpeechService.DeleteCustomModel(deleteCustomModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteCustomModel with error: Operation validation and request error`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the DeleteCustomModelOptions model
				deleteCustomModelOptionsModel := new(texttospeechv1.DeleteCustomModelOptions)
				deleteCustomModelOptionsModel.CustomizationID = core.StringPtr("testString")
				deleteCustomModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := textToSpeechService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := textToSpeechService.DeleteCustomModel(deleteCustomModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteCustomModelOptions model with no property values
				deleteCustomModelOptionsModelNew := new(texttospeechv1.DeleteCustomModelOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = textToSpeechService.DeleteCustomModel(deleteCustomModelOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AddWords(addWordsOptions *AddWordsOptions)`, func() {
		addWordsPath := "/v1/customizations/testString/words"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addWordsPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					res.WriteHeader(200)
				}))
			})
			It(`Invoke AddWords successfully`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := textToSpeechService.AddWords(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the Word model
				wordModel := new(texttospeechv1.Word)
				wordModel.Word = core.StringPtr("testString")
				wordModel.Translation = core.StringPtr("testString")
				wordModel.PartOfSpeech = core.StringPtr("Dosi")

				// Construct an instance of the AddWordsOptions model
				addWordsOptionsModel := new(texttospeechv1.AddWordsOptions)
				addWordsOptionsModel.CustomizationID = core.StringPtr("testString")
				addWordsOptionsModel.Words = []texttospeechv1.Word{*wordModel}
				addWordsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = textToSpeechService.AddWords(addWordsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke AddWords with error: Operation validation and request error`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the Word model
				wordModel := new(texttospeechv1.Word)
				wordModel.Word = core.StringPtr("testString")
				wordModel.Translation = core.StringPtr("testString")
				wordModel.PartOfSpeech = core.StringPtr("Dosi")

				// Construct an instance of the AddWordsOptions model
				addWordsOptionsModel := new(texttospeechv1.AddWordsOptions)
				addWordsOptionsModel.CustomizationID = core.StringPtr("testString")
				addWordsOptionsModel.Words = []texttospeechv1.Word{*wordModel}
				addWordsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := textToSpeechService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := textToSpeechService.AddWords(addWordsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the AddWordsOptions model with no property values
				addWordsOptionsModelNew := new(texttospeechv1.AddWordsOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = textToSpeechService.AddWords(addWordsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListWords(listWordsOptions *ListWordsOptions) - Operation response error`, func() {
		listWordsPath := "/v1/customizations/testString/words"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listWordsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListWords with error: Operation response processing error`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the ListWordsOptions model
				listWordsOptionsModel := new(texttospeechv1.ListWordsOptions)
				listWordsOptionsModel.CustomizationID = core.StringPtr("testString")
				listWordsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := textToSpeechService.ListWords(listWordsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				textToSpeechService.EnableRetries(0, 0)
				result, response, operationErr = textToSpeechService.ListWords(listWordsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListWords(listWordsOptions *ListWordsOptions)`, func() {
		listWordsPath := "/v1/customizations/testString/words"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listWordsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"words": [{"word": "Word", "translation": "Translation", "part_of_speech": "Dosi"}]}`)
				}))
			})
			It(`Invoke ListWords successfully with retries`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())
				textToSpeechService.EnableRetries(0, 0)

				// Construct an instance of the ListWordsOptions model
				listWordsOptionsModel := new(texttospeechv1.ListWordsOptions)
				listWordsOptionsModel.CustomizationID = core.StringPtr("testString")
				listWordsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := textToSpeechService.ListWordsWithContext(ctx, listWordsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				textToSpeechService.DisableRetries()
				result, response, operationErr := textToSpeechService.ListWords(listWordsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = textToSpeechService.ListWordsWithContext(ctx, listWordsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listWordsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"words": [{"word": "Word", "translation": "Translation", "part_of_speech": "Dosi"}]}`)
				}))
			})
			It(`Invoke ListWords successfully`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := textToSpeechService.ListWords(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListWordsOptions model
				listWordsOptionsModel := new(texttospeechv1.ListWordsOptions)
				listWordsOptionsModel.CustomizationID = core.StringPtr("testString")
				listWordsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = textToSpeechService.ListWords(listWordsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListWords with error: Operation validation and request error`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the ListWordsOptions model
				listWordsOptionsModel := new(texttospeechv1.ListWordsOptions)
				listWordsOptionsModel.CustomizationID = core.StringPtr("testString")
				listWordsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := textToSpeechService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := textToSpeechService.ListWords(listWordsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListWordsOptions model with no property values
				listWordsOptionsModelNew := new(texttospeechv1.ListWordsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = textToSpeechService.ListWords(listWordsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListWords successfully`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the ListWordsOptions model
				listWordsOptionsModel := new(texttospeechv1.ListWordsOptions)
				listWordsOptionsModel.CustomizationID = core.StringPtr("testString")
				listWordsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := textToSpeechService.ListWords(listWordsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AddWord(addWordOptions *AddWordOptions)`, func() {
		addWordPath := "/v1/customizations/testString/words/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addWordPath))
					Expect(req.Method).To(Equal("PUT"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					res.WriteHeader(200)
				}))
			})
			It(`Invoke AddWord successfully`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := textToSpeechService.AddWord(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the AddWordOptions model
				addWordOptionsModel := new(texttospeechv1.AddWordOptions)
				addWordOptionsModel.CustomizationID = core.StringPtr("testString")
				addWordOptionsModel.Word = core.StringPtr("testString")
				addWordOptionsModel.Translation = core.StringPtr("testString")
				addWordOptionsModel.PartOfSpeech = core.StringPtr("Dosi")
				addWordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = textToSpeechService.AddWord(addWordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke AddWord with error: Operation validation and request error`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the AddWordOptions model
				addWordOptionsModel := new(texttospeechv1.AddWordOptions)
				addWordOptionsModel.CustomizationID = core.StringPtr("testString")
				addWordOptionsModel.Word = core.StringPtr("testString")
				addWordOptionsModel.Translation = core.StringPtr("testString")
				addWordOptionsModel.PartOfSpeech = core.StringPtr("Dosi")
				addWordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := textToSpeechService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := textToSpeechService.AddWord(addWordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the AddWordOptions model with no property values
				addWordOptionsModelNew := new(texttospeechv1.AddWordOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = textToSpeechService.AddWord(addWordOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetWord(getWordOptions *GetWordOptions) - Operation response error`, func() {
		getWordPath := "/v1/customizations/testString/words/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWordPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetWord with error: Operation response processing error`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the GetWordOptions model
				getWordOptionsModel := new(texttospeechv1.GetWordOptions)
				getWordOptionsModel.CustomizationID = core.StringPtr("testString")
				getWordOptionsModel.Word = core.StringPtr("testString")
				getWordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := textToSpeechService.GetWord(getWordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				textToSpeechService.EnableRetries(0, 0)
				result, response, operationErr = textToSpeechService.GetWord(getWordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetWord(getWordOptions *GetWordOptions)`, func() {
		getWordPath := "/v1/customizations/testString/words/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWordPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"translation": "Translation", "part_of_speech": "Dosi"}`)
				}))
			})
			It(`Invoke GetWord successfully with retries`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())
				textToSpeechService.EnableRetries(0, 0)

				// Construct an instance of the GetWordOptions model
				getWordOptionsModel := new(texttospeechv1.GetWordOptions)
				getWordOptionsModel.CustomizationID = core.StringPtr("testString")
				getWordOptionsModel.Word = core.StringPtr("testString")
				getWordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := textToSpeechService.GetWordWithContext(ctx, getWordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				textToSpeechService.DisableRetries()
				result, response, operationErr := textToSpeechService.GetWord(getWordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = textToSpeechService.GetWordWithContext(ctx, getWordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getWordPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"translation": "Translation", "part_of_speech": "Dosi"}`)
				}))
			})
			It(`Invoke GetWord successfully`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := textToSpeechService.GetWord(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetWordOptions model
				getWordOptionsModel := new(texttospeechv1.GetWordOptions)
				getWordOptionsModel.CustomizationID = core.StringPtr("testString")
				getWordOptionsModel.Word = core.StringPtr("testString")
				getWordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = textToSpeechService.GetWord(getWordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetWord with error: Operation validation and request error`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the GetWordOptions model
				getWordOptionsModel := new(texttospeechv1.GetWordOptions)
				getWordOptionsModel.CustomizationID = core.StringPtr("testString")
				getWordOptionsModel.Word = core.StringPtr("testString")
				getWordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := textToSpeechService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := textToSpeechService.GetWord(getWordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetWordOptions model with no property values
				getWordOptionsModelNew := new(texttospeechv1.GetWordOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = textToSpeechService.GetWord(getWordOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetWord successfully`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the GetWordOptions model
				getWordOptionsModel := new(texttospeechv1.GetWordOptions)
				getWordOptionsModel.CustomizationID = core.StringPtr("testString")
				getWordOptionsModel.Word = core.StringPtr("testString")
				getWordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := textToSpeechService.GetWord(getWordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteWord(deleteWordOptions *DeleteWordOptions)`, func() {
		deleteWordPath := "/v1/customizations/testString/words/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteWordPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteWord successfully`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := textToSpeechService.DeleteWord(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteWordOptions model
				deleteWordOptionsModel := new(texttospeechv1.DeleteWordOptions)
				deleteWordOptionsModel.CustomizationID = core.StringPtr("testString")
				deleteWordOptionsModel.Word = core.StringPtr("testString")
				deleteWordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = textToSpeechService.DeleteWord(deleteWordOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteWord with error: Operation validation and request error`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the DeleteWordOptions model
				deleteWordOptionsModel := new(texttospeechv1.DeleteWordOptions)
				deleteWordOptionsModel.CustomizationID = core.StringPtr("testString")
				deleteWordOptionsModel.Word = core.StringPtr("testString")
				deleteWordOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := textToSpeechService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := textToSpeechService.DeleteWord(deleteWordOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteWordOptions model with no property values
				deleteWordOptionsModelNew := new(texttospeechv1.DeleteWordOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = textToSpeechService.DeleteWord(deleteWordOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListCustomPrompts(listCustomPromptsOptions *ListCustomPromptsOptions) - Operation response error`, func() {
		listCustomPromptsPath := "/v1/customizations/testString/prompts"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCustomPromptsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListCustomPrompts with error: Operation response processing error`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the ListCustomPromptsOptions model
				listCustomPromptsOptionsModel := new(texttospeechv1.ListCustomPromptsOptions)
				listCustomPromptsOptionsModel.CustomizationID = core.StringPtr("testString")
				listCustomPromptsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := textToSpeechService.ListCustomPrompts(listCustomPromptsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				textToSpeechService.EnableRetries(0, 0)
				result, response, operationErr = textToSpeechService.ListCustomPrompts(listCustomPromptsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListCustomPrompts(listCustomPromptsOptions *ListCustomPromptsOptions)`, func() {
		listCustomPromptsPath := "/v1/customizations/testString/prompts"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCustomPromptsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"prompts": [{"prompt": "Prompt", "prompt_id": "PromptID", "status": "Status", "error": "Error", "speaker_id": "SpeakerID"}]}`)
				}))
			})
			It(`Invoke ListCustomPrompts successfully with retries`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())
				textToSpeechService.EnableRetries(0, 0)

				// Construct an instance of the ListCustomPromptsOptions model
				listCustomPromptsOptionsModel := new(texttospeechv1.ListCustomPromptsOptions)
				listCustomPromptsOptionsModel.CustomizationID = core.StringPtr("testString")
				listCustomPromptsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := textToSpeechService.ListCustomPromptsWithContext(ctx, listCustomPromptsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				textToSpeechService.DisableRetries()
				result, response, operationErr := textToSpeechService.ListCustomPrompts(listCustomPromptsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = textToSpeechService.ListCustomPromptsWithContext(ctx, listCustomPromptsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listCustomPromptsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"prompts": [{"prompt": "Prompt", "prompt_id": "PromptID", "status": "Status", "error": "Error", "speaker_id": "SpeakerID"}]}`)
				}))
			})
			It(`Invoke ListCustomPrompts successfully`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := textToSpeechService.ListCustomPrompts(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListCustomPromptsOptions model
				listCustomPromptsOptionsModel := new(texttospeechv1.ListCustomPromptsOptions)
				listCustomPromptsOptionsModel.CustomizationID = core.StringPtr("testString")
				listCustomPromptsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = textToSpeechService.ListCustomPrompts(listCustomPromptsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListCustomPrompts with error: Operation validation and request error`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the ListCustomPromptsOptions model
				listCustomPromptsOptionsModel := new(texttospeechv1.ListCustomPromptsOptions)
				listCustomPromptsOptionsModel.CustomizationID = core.StringPtr("testString")
				listCustomPromptsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := textToSpeechService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := textToSpeechService.ListCustomPrompts(listCustomPromptsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the ListCustomPromptsOptions model with no property values
				listCustomPromptsOptionsModelNew := new(texttospeechv1.ListCustomPromptsOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = textToSpeechService.ListCustomPrompts(listCustomPromptsOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListCustomPrompts successfully`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the ListCustomPromptsOptions model
				listCustomPromptsOptionsModel := new(texttospeechv1.ListCustomPromptsOptions)
				listCustomPromptsOptionsModel.CustomizationID = core.StringPtr("testString")
				listCustomPromptsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := textToSpeechService.ListCustomPrompts(listCustomPromptsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AddCustomPrompt(addCustomPromptOptions *AddCustomPromptOptions) - Operation response error`, func() {
		addCustomPromptPath := "/v1/customizations/testString/prompts/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addCustomPromptPath))
					Expect(req.Method).To(Equal("POST"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke AddCustomPrompt with error: Operation response processing error`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the PromptMetadata model
				promptMetadataModel := new(texttospeechv1.PromptMetadata)
				promptMetadataModel.PromptText = core.StringPtr("testString")
				promptMetadataModel.SpeakerID = core.StringPtr("testString")

				// Construct an instance of the AddCustomPromptOptions model
				addCustomPromptOptionsModel := new(texttospeechv1.AddCustomPromptOptions)
				addCustomPromptOptionsModel.CustomizationID = core.StringPtr("testString")
				addCustomPromptOptionsModel.PromptID = core.StringPtr("testString")
				addCustomPromptOptionsModel.Metadata = promptMetadataModel
				addCustomPromptOptionsModel.File = CreateMockReader("This is a mock file.")
				addCustomPromptOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := textToSpeechService.AddCustomPrompt(addCustomPromptOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				textToSpeechService.EnableRetries(0, 0)
				result, response, operationErr = textToSpeechService.AddCustomPrompt(addCustomPromptOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`AddCustomPrompt(addCustomPromptOptions *AddCustomPromptOptions)`, func() {
		addCustomPromptPath := "/v1/customizations/testString/prompts/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addCustomPromptPath))
					Expect(req.Method).To(Equal("POST"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"prompt": "Prompt", "prompt_id": "PromptID", "status": "Status", "error": "Error", "speaker_id": "SpeakerID"}`)
				}))
			})
			It(`Invoke AddCustomPrompt successfully with retries`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())
				textToSpeechService.EnableRetries(0, 0)

				// Construct an instance of the PromptMetadata model
				promptMetadataModel := new(texttospeechv1.PromptMetadata)
				promptMetadataModel.PromptText = core.StringPtr("testString")
				promptMetadataModel.SpeakerID = core.StringPtr("testString")

				// Construct an instance of the AddCustomPromptOptions model
				addCustomPromptOptionsModel := new(texttospeechv1.AddCustomPromptOptions)
				addCustomPromptOptionsModel.CustomizationID = core.StringPtr("testString")
				addCustomPromptOptionsModel.PromptID = core.StringPtr("testString")
				addCustomPromptOptionsModel.Metadata = promptMetadataModel
				addCustomPromptOptionsModel.File = CreateMockReader("This is a mock file.")
				addCustomPromptOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := textToSpeechService.AddCustomPromptWithContext(ctx, addCustomPromptOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				textToSpeechService.DisableRetries()
				result, response, operationErr := textToSpeechService.AddCustomPrompt(addCustomPromptOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = textToSpeechService.AddCustomPromptWithContext(ctx, addCustomPromptOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(addCustomPromptPath))
					Expect(req.Method).To(Equal("POST"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"prompt": "Prompt", "prompt_id": "PromptID", "status": "Status", "error": "Error", "speaker_id": "SpeakerID"}`)
				}))
			})
			It(`Invoke AddCustomPrompt successfully`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := textToSpeechService.AddCustomPrompt(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the PromptMetadata model
				promptMetadataModel := new(texttospeechv1.PromptMetadata)
				promptMetadataModel.PromptText = core.StringPtr("testString")
				promptMetadataModel.SpeakerID = core.StringPtr("testString")

				// Construct an instance of the AddCustomPromptOptions model
				addCustomPromptOptionsModel := new(texttospeechv1.AddCustomPromptOptions)
				addCustomPromptOptionsModel.CustomizationID = core.StringPtr("testString")
				addCustomPromptOptionsModel.PromptID = core.StringPtr("testString")
				addCustomPromptOptionsModel.Metadata = promptMetadataModel
				addCustomPromptOptionsModel.File = CreateMockReader("This is a mock file.")
				addCustomPromptOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = textToSpeechService.AddCustomPrompt(addCustomPromptOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke AddCustomPrompt with error: Operation validation and request error`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the PromptMetadata model
				promptMetadataModel := new(texttospeechv1.PromptMetadata)
				promptMetadataModel.PromptText = core.StringPtr("testString")
				promptMetadataModel.SpeakerID = core.StringPtr("testString")

				// Construct an instance of the AddCustomPromptOptions model
				addCustomPromptOptionsModel := new(texttospeechv1.AddCustomPromptOptions)
				addCustomPromptOptionsModel.CustomizationID = core.StringPtr("testString")
				addCustomPromptOptionsModel.PromptID = core.StringPtr("testString")
				addCustomPromptOptionsModel.Metadata = promptMetadataModel
				addCustomPromptOptionsModel.File = CreateMockReader("This is a mock file.")
				addCustomPromptOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := textToSpeechService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := textToSpeechService.AddCustomPrompt(addCustomPromptOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the AddCustomPromptOptions model with no property values
				addCustomPromptOptionsModelNew := new(texttospeechv1.AddCustomPromptOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = textToSpeechService.AddCustomPrompt(addCustomPromptOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke AddCustomPrompt successfully`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the PromptMetadata model
				promptMetadataModel := new(texttospeechv1.PromptMetadata)
				promptMetadataModel.PromptText = core.StringPtr("testString")
				promptMetadataModel.SpeakerID = core.StringPtr("testString")

				// Construct an instance of the AddCustomPromptOptions model
				addCustomPromptOptionsModel := new(texttospeechv1.AddCustomPromptOptions)
				addCustomPromptOptionsModel.CustomizationID = core.StringPtr("testString")
				addCustomPromptOptionsModel.PromptID = core.StringPtr("testString")
				addCustomPromptOptionsModel.Metadata = promptMetadataModel
				addCustomPromptOptionsModel.File = CreateMockReader("This is a mock file.")
				addCustomPromptOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := textToSpeechService.AddCustomPrompt(addCustomPromptOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCustomPrompt(getCustomPromptOptions *GetCustomPromptOptions) - Operation response error`, func() {
		getCustomPromptPath := "/v1/customizations/testString/prompts/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCustomPromptPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetCustomPrompt with error: Operation response processing error`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the GetCustomPromptOptions model
				getCustomPromptOptionsModel := new(texttospeechv1.GetCustomPromptOptions)
				getCustomPromptOptionsModel.CustomizationID = core.StringPtr("testString")
				getCustomPromptOptionsModel.PromptID = core.StringPtr("testString")
				getCustomPromptOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := textToSpeechService.GetCustomPrompt(getCustomPromptOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				textToSpeechService.EnableRetries(0, 0)
				result, response, operationErr = textToSpeechService.GetCustomPrompt(getCustomPromptOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetCustomPrompt(getCustomPromptOptions *GetCustomPromptOptions)`, func() {
		getCustomPromptPath := "/v1/customizations/testString/prompts/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCustomPromptPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"prompt": "Prompt", "prompt_id": "PromptID", "status": "Status", "error": "Error", "speaker_id": "SpeakerID"}`)
				}))
			})
			It(`Invoke GetCustomPrompt successfully with retries`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())
				textToSpeechService.EnableRetries(0, 0)

				// Construct an instance of the GetCustomPromptOptions model
				getCustomPromptOptionsModel := new(texttospeechv1.GetCustomPromptOptions)
				getCustomPromptOptionsModel.CustomizationID = core.StringPtr("testString")
				getCustomPromptOptionsModel.PromptID = core.StringPtr("testString")
				getCustomPromptOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := textToSpeechService.GetCustomPromptWithContext(ctx, getCustomPromptOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				textToSpeechService.DisableRetries()
				result, response, operationErr := textToSpeechService.GetCustomPrompt(getCustomPromptOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = textToSpeechService.GetCustomPromptWithContext(ctx, getCustomPromptOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getCustomPromptPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"prompt": "Prompt", "prompt_id": "PromptID", "status": "Status", "error": "Error", "speaker_id": "SpeakerID"}`)
				}))
			})
			It(`Invoke GetCustomPrompt successfully`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := textToSpeechService.GetCustomPrompt(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetCustomPromptOptions model
				getCustomPromptOptionsModel := new(texttospeechv1.GetCustomPromptOptions)
				getCustomPromptOptionsModel.CustomizationID = core.StringPtr("testString")
				getCustomPromptOptionsModel.PromptID = core.StringPtr("testString")
				getCustomPromptOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = textToSpeechService.GetCustomPrompt(getCustomPromptOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetCustomPrompt with error: Operation validation and request error`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the GetCustomPromptOptions model
				getCustomPromptOptionsModel := new(texttospeechv1.GetCustomPromptOptions)
				getCustomPromptOptionsModel.CustomizationID = core.StringPtr("testString")
				getCustomPromptOptionsModel.PromptID = core.StringPtr("testString")
				getCustomPromptOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := textToSpeechService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := textToSpeechService.GetCustomPrompt(getCustomPromptOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetCustomPromptOptions model with no property values
				getCustomPromptOptionsModelNew := new(texttospeechv1.GetCustomPromptOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = textToSpeechService.GetCustomPrompt(getCustomPromptOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetCustomPrompt successfully`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the GetCustomPromptOptions model
				getCustomPromptOptionsModel := new(texttospeechv1.GetCustomPromptOptions)
				getCustomPromptOptionsModel.CustomizationID = core.StringPtr("testString")
				getCustomPromptOptionsModel.PromptID = core.StringPtr("testString")
				getCustomPromptOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := textToSpeechService.GetCustomPrompt(getCustomPromptOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteCustomPrompt(deleteCustomPromptOptions *DeleteCustomPromptOptions)`, func() {
		deleteCustomPromptPath := "/v1/customizations/testString/prompts/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteCustomPromptPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteCustomPrompt successfully`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := textToSpeechService.DeleteCustomPrompt(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteCustomPromptOptions model
				deleteCustomPromptOptionsModel := new(texttospeechv1.DeleteCustomPromptOptions)
				deleteCustomPromptOptionsModel.CustomizationID = core.StringPtr("testString")
				deleteCustomPromptOptionsModel.PromptID = core.StringPtr("testString")
				deleteCustomPromptOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = textToSpeechService.DeleteCustomPrompt(deleteCustomPromptOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteCustomPrompt with error: Operation validation and request error`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the DeleteCustomPromptOptions model
				deleteCustomPromptOptionsModel := new(texttospeechv1.DeleteCustomPromptOptions)
				deleteCustomPromptOptionsModel.CustomizationID = core.StringPtr("testString")
				deleteCustomPromptOptionsModel.PromptID = core.StringPtr("testString")
				deleteCustomPromptOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := textToSpeechService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := textToSpeechService.DeleteCustomPrompt(deleteCustomPromptOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteCustomPromptOptions model with no property values
				deleteCustomPromptOptionsModelNew := new(texttospeechv1.DeleteCustomPromptOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = textToSpeechService.DeleteCustomPrompt(deleteCustomPromptOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListSpeakerModels(listSpeakerModelsOptions *ListSpeakerModelsOptions) - Operation response error`, func() {
		listSpeakerModelsPath := "/v1/speakers"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSpeakerModelsPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke ListSpeakerModels with error: Operation response processing error`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the ListSpeakerModelsOptions model
				listSpeakerModelsOptionsModel := new(texttospeechv1.ListSpeakerModelsOptions)
				listSpeakerModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := textToSpeechService.ListSpeakerModels(listSpeakerModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				textToSpeechService.EnableRetries(0, 0)
				result, response, operationErr = textToSpeechService.ListSpeakerModels(listSpeakerModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`ListSpeakerModels(listSpeakerModelsOptions *ListSpeakerModelsOptions)`, func() {
		listSpeakerModelsPath := "/v1/speakers"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSpeakerModelsPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"speakers": [{"speaker_id": "SpeakerID", "name": "Name"}]}`)
				}))
			})
			It(`Invoke ListSpeakerModels successfully with retries`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())
				textToSpeechService.EnableRetries(0, 0)

				// Construct an instance of the ListSpeakerModelsOptions model
				listSpeakerModelsOptionsModel := new(texttospeechv1.ListSpeakerModelsOptions)
				listSpeakerModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := textToSpeechService.ListSpeakerModelsWithContext(ctx, listSpeakerModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				textToSpeechService.DisableRetries()
				result, response, operationErr := textToSpeechService.ListSpeakerModels(listSpeakerModelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = textToSpeechService.ListSpeakerModelsWithContext(ctx, listSpeakerModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(listSpeakerModelsPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"speakers": [{"speaker_id": "SpeakerID", "name": "Name"}]}`)
				}))
			})
			It(`Invoke ListSpeakerModels successfully`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := textToSpeechService.ListSpeakerModels(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the ListSpeakerModelsOptions model
				listSpeakerModelsOptionsModel := new(texttospeechv1.ListSpeakerModelsOptions)
				listSpeakerModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = textToSpeechService.ListSpeakerModels(listSpeakerModelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke ListSpeakerModels with error: Operation request error`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the ListSpeakerModelsOptions model
				listSpeakerModelsOptionsModel := new(texttospeechv1.ListSpeakerModelsOptions)
				listSpeakerModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := textToSpeechService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := textToSpeechService.ListSpeakerModels(listSpeakerModelsOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke ListSpeakerModels successfully`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the ListSpeakerModelsOptions model
				listSpeakerModelsOptionsModel := new(texttospeechv1.ListSpeakerModelsOptions)
				listSpeakerModelsOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := textToSpeechService.ListSpeakerModels(listSpeakerModelsOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateSpeakerModel(createSpeakerModelOptions *CreateSpeakerModelOptions) - Operation response error`, func() {
		createSpeakerModelPath := "/v1/speakers"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSpeakerModelPath))
					Expect(req.Method).To(Equal("POST"))
					Expect(req.URL.Query()["speaker_name"]).To(Equal([]string{"testString"}))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke CreateSpeakerModel with error: Operation response processing error`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the CreateSpeakerModelOptions model
				createSpeakerModelOptionsModel := new(texttospeechv1.CreateSpeakerModelOptions)
				createSpeakerModelOptionsModel.SpeakerName = core.StringPtr("testString")
				createSpeakerModelOptionsModel.Audio = CreateMockReader("This is a mock file.")
				createSpeakerModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := textToSpeechService.CreateSpeakerModel(createSpeakerModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				textToSpeechService.EnableRetries(0, 0)
				result, response, operationErr = textToSpeechService.CreateSpeakerModel(createSpeakerModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`CreateSpeakerModel(createSpeakerModelOptions *CreateSpeakerModelOptions)`, func() {
		createSpeakerModelPath := "/v1/speakers"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSpeakerModelPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["speaker_name"]).To(Equal([]string{"testString"}))
					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"speaker_id": "SpeakerID"}`)
				}))
			})
			It(`Invoke CreateSpeakerModel successfully with retries`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())
				textToSpeechService.EnableRetries(0, 0)

				// Construct an instance of the CreateSpeakerModelOptions model
				createSpeakerModelOptionsModel := new(texttospeechv1.CreateSpeakerModelOptions)
				createSpeakerModelOptionsModel.SpeakerName = core.StringPtr("testString")
				createSpeakerModelOptionsModel.Audio = CreateMockReader("This is a mock file.")
				createSpeakerModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := textToSpeechService.CreateSpeakerModelWithContext(ctx, createSpeakerModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				textToSpeechService.DisableRetries()
				result, response, operationErr := textToSpeechService.CreateSpeakerModel(createSpeakerModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = textToSpeechService.CreateSpeakerModelWithContext(ctx, createSpeakerModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(createSpeakerModelPath))
					Expect(req.Method).To(Equal("POST"))

					// For gzip-disabled operation, verify Content-Encoding is not set.
					Expect(req.Header.Get("Content-Encoding")).To(BeEmpty())

					// If there is a body, then make sure we can read it
					bodyBuf := new(bytes.Buffer)
					if req.Header.Get("Content-Encoding") == "gzip" {
						body, err := core.NewGzipDecompressionReader(req.Body)
						Expect(err).To(BeNil())
						_, err = bodyBuf.ReadFrom(body)
						Expect(err).To(BeNil())
					} else {
						_, err := bodyBuf.ReadFrom(req.Body)
						Expect(err).To(BeNil())
					}
					fmt.Fprintf(GinkgoWriter, "  Request body: %s", bodyBuf.String())

					Expect(req.URL.Query()["speaker_name"]).To(Equal([]string{"testString"}))
					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(201)
					fmt.Fprintf(res, "%s", `{"speaker_id": "SpeakerID"}`)
				}))
			})
			It(`Invoke CreateSpeakerModel successfully`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := textToSpeechService.CreateSpeakerModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the CreateSpeakerModelOptions model
				createSpeakerModelOptionsModel := new(texttospeechv1.CreateSpeakerModelOptions)
				createSpeakerModelOptionsModel.SpeakerName = core.StringPtr("testString")
				createSpeakerModelOptionsModel.Audio = CreateMockReader("This is a mock file.")
				createSpeakerModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = textToSpeechService.CreateSpeakerModel(createSpeakerModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke CreateSpeakerModel with error: Operation validation and request error`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the CreateSpeakerModelOptions model
				createSpeakerModelOptionsModel := new(texttospeechv1.CreateSpeakerModelOptions)
				createSpeakerModelOptionsModel.SpeakerName = core.StringPtr("testString")
				createSpeakerModelOptionsModel.Audio = CreateMockReader("This is a mock file.")
				createSpeakerModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := textToSpeechService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := textToSpeechService.CreateSpeakerModel(createSpeakerModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the CreateSpeakerModelOptions model with no property values
				createSpeakerModelOptionsModelNew := new(texttospeechv1.CreateSpeakerModelOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = textToSpeechService.CreateSpeakerModel(createSpeakerModelOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(201)
				}))
			})
			It(`Invoke CreateSpeakerModel successfully`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the CreateSpeakerModelOptions model
				createSpeakerModelOptionsModel := new(texttospeechv1.CreateSpeakerModelOptions)
				createSpeakerModelOptionsModel.SpeakerName = core.StringPtr("testString")
				createSpeakerModelOptionsModel.Audio = CreateMockReader("This is a mock file.")
				createSpeakerModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := textToSpeechService.CreateSpeakerModel(createSpeakerModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSpeakerModel(getSpeakerModelOptions *GetSpeakerModelOptions) - Operation response error`, func() {
		getSpeakerModelPath := "/v1/speakers/testString"
		Context(`Using mock server endpoint with invalid JSON response`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSpeakerModelPath))
					Expect(req.Method).To(Equal("GET"))
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprint(res, `} this is not valid json {`)
				}))
			})
			It(`Invoke GetSpeakerModel with error: Operation response processing error`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the GetSpeakerModelOptions model
				getSpeakerModelOptionsModel := new(texttospeechv1.GetSpeakerModelOptions)
				getSpeakerModelOptionsModel.SpeakerID = core.StringPtr("testString")
				getSpeakerModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Expect response parsing to fail since we are receiving a text/plain response
				result, response, operationErr := textToSpeechService.GetSpeakerModel(getSpeakerModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())

				// Enable retries and test again
				textToSpeechService.EnableRetries(0, 0)
				result, response, operationErr = textToSpeechService.GetSpeakerModel(getSpeakerModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`GetSpeakerModel(getSpeakerModelOptions *GetSpeakerModelOptions)`, func() {
		getSpeakerModelPath := "/v1/speakers/testString"
		Context(`Using mock server endpoint with timeout`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSpeakerModelPath))
					Expect(req.Method).To(Equal("GET"))

					// Sleep a short time to support a timeout test
					time.Sleep(100 * time.Millisecond)

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"customizations": [{"customization_id": "CustomizationID", "prompts": [{"prompt": "Prompt", "prompt_id": "PromptID", "status": "Status", "error": "Error"}]}]}`)
				}))
			})
			It(`Invoke GetSpeakerModel successfully with retries`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())
				textToSpeechService.EnableRetries(0, 0)

				// Construct an instance of the GetSpeakerModelOptions model
				getSpeakerModelOptionsModel := new(texttospeechv1.GetSpeakerModelOptions)
				getSpeakerModelOptionsModel.SpeakerID = core.StringPtr("testString")
				getSpeakerModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with a Context to test a timeout error
				ctx, cancelFunc := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc()
				_, _, operationErr := textToSpeechService.GetSpeakerModelWithContext(ctx, getSpeakerModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))

				// Disable retries and test again
				textToSpeechService.DisableRetries()
				result, response, operationErr := textToSpeechService.GetSpeakerModel(getSpeakerModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

				// Re-test the timeout error with retries disabled
				ctx, cancelFunc2 := context.WithTimeout(context.Background(), 80*time.Millisecond)
				defer cancelFunc2()
				_, _, operationErr = textToSpeechService.GetSpeakerModelWithContext(ctx, getSpeakerModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring("deadline exceeded"))
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(getSpeakerModelPath))
					Expect(req.Method).To(Equal("GET"))

					// Set mock response
					res.Header().Set("Content-type", "application/json")
					res.WriteHeader(200)
					fmt.Fprintf(res, "%s", `{"customizations": [{"customization_id": "CustomizationID", "prompts": [{"prompt": "Prompt", "prompt_id": "PromptID", "status": "Status", "error": "Error"}]}]}`)
				}))
			})
			It(`Invoke GetSpeakerModel successfully`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				result, response, operationErr := textToSpeechService.GetSpeakerModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())

				// Construct an instance of the GetSpeakerModelOptions model
				getSpeakerModelOptionsModel := new(texttospeechv1.GetSpeakerModelOptions)
				getSpeakerModelOptionsModel.SpeakerID = core.StringPtr("testString")
				getSpeakerModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				result, response, operationErr = textToSpeechService.GetSpeakerModel(getSpeakerModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
				Expect(result).ToNot(BeNil())

			})
			It(`Invoke GetSpeakerModel with error: Operation validation and request error`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the GetSpeakerModelOptions model
				getSpeakerModelOptionsModel := new(texttospeechv1.GetSpeakerModelOptions)
				getSpeakerModelOptionsModel.SpeakerID = core.StringPtr("testString")
				getSpeakerModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := textToSpeechService.SetServiceURL("")
				Expect(err).To(BeNil())
				result, response, operationErr := textToSpeechService.GetSpeakerModel(getSpeakerModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
				// Construct a second instance of the GetSpeakerModelOptions model with no property values
				getSpeakerModelOptionsModelNew := new(texttospeechv1.GetSpeakerModelOptions)
				// Invoke operation with invalid model (negative test)
				result, response, operationErr = textToSpeechService.GetSpeakerModel(getSpeakerModelOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
		Context(`Using mock server endpoint with missing response body`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Set success status code with no respoonse body
					res.WriteHeader(200)
				}))
			})
			It(`Invoke GetSpeakerModel successfully`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the GetSpeakerModelOptions model
				getSpeakerModelOptionsModel := new(texttospeechv1.GetSpeakerModelOptions)
				getSpeakerModelOptionsModel.SpeakerID = core.StringPtr("testString")
				getSpeakerModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation
				result, response, operationErr := textToSpeechService.GetSpeakerModel(getSpeakerModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())

				// Verify a nil result
				Expect(result).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteSpeakerModel(deleteSpeakerModelOptions *DeleteSpeakerModelOptions)`, func() {
		deleteSpeakerModelPath := "/v1/speakers/testString"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteSpeakerModelPath))
					Expect(req.Method).To(Equal("DELETE"))

					res.WriteHeader(204)
				}))
			})
			It(`Invoke DeleteSpeakerModel successfully`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := textToSpeechService.DeleteSpeakerModel(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteSpeakerModelOptions model
				deleteSpeakerModelOptionsModel := new(texttospeechv1.DeleteSpeakerModelOptions)
				deleteSpeakerModelOptionsModel.SpeakerID = core.StringPtr("testString")
				deleteSpeakerModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = textToSpeechService.DeleteSpeakerModel(deleteSpeakerModelOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteSpeakerModel with error: Operation validation and request error`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the DeleteSpeakerModelOptions model
				deleteSpeakerModelOptionsModel := new(texttospeechv1.DeleteSpeakerModelOptions)
				deleteSpeakerModelOptionsModel.SpeakerID = core.StringPtr("testString")
				deleteSpeakerModelOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := textToSpeechService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := textToSpeechService.DeleteSpeakerModel(deleteSpeakerModelOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteSpeakerModelOptions model with no property values
				deleteSpeakerModelOptionsModelNew := new(texttospeechv1.DeleteSpeakerModelOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = textToSpeechService.DeleteSpeakerModel(deleteSpeakerModelOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`DeleteUserData(deleteUserDataOptions *DeleteUserDataOptions)`, func() {
		deleteUserDataPath := "/v1/user_data"
		Context(`Using mock server endpoint`, func() {
			BeforeEach(func() {
				testServer = httptest.NewServer(http.HandlerFunc(func(res http.ResponseWriter, req *http.Request) {
					defer GinkgoRecover()

					// Verify the contents of the request
					Expect(req.URL.EscapedPath()).To(Equal(deleteUserDataPath))
					Expect(req.Method).To(Equal("DELETE"))

					Expect(req.URL.Query()["customer_id"]).To(Equal([]string{"testString"}))
					res.WriteHeader(200)
				}))
			})
			It(`Invoke DeleteUserData successfully`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Invoke operation with nil options model (negative test)
				response, operationErr := textToSpeechService.DeleteUserData(nil)
				Expect(operationErr).NotTo(BeNil())
				Expect(response).To(BeNil())

				// Construct an instance of the DeleteUserDataOptions model
				deleteUserDataOptionsModel := new(texttospeechv1.DeleteUserDataOptions)
				deleteUserDataOptionsModel.CustomerID = core.StringPtr("testString")
				deleteUserDataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}

				// Invoke operation with valid options model (positive test)
				response, operationErr = textToSpeechService.DeleteUserData(deleteUserDataOptionsModel)
				Expect(operationErr).To(BeNil())
				Expect(response).ToNot(BeNil())
			})
			It(`Invoke DeleteUserData with error: Operation validation and request error`, func() {
				textToSpeechService, serviceErr := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
					URL:           testServer.URL,
					Authenticator: &core.NoAuthAuthenticator{},
				})
				Expect(serviceErr).To(BeNil())
				Expect(textToSpeechService).ToNot(BeNil())

				// Construct an instance of the DeleteUserDataOptions model
				deleteUserDataOptionsModel := new(texttospeechv1.DeleteUserDataOptions)
				deleteUserDataOptionsModel.CustomerID = core.StringPtr("testString")
				deleteUserDataOptionsModel.Headers = map[string]string{"x-custom-header": "x-custom-value"}
				// Invoke operation with empty URL (negative test)
				err := textToSpeechService.SetServiceURL("")
				Expect(err).To(BeNil())
				response, operationErr := textToSpeechService.DeleteUserData(deleteUserDataOptionsModel)
				Expect(operationErr).ToNot(BeNil())
				Expect(operationErr.Error()).To(ContainSubstring(core.ERRORMSG_SERVICE_URL_MISSING))
				Expect(response).To(BeNil())
				// Construct a second instance of the DeleteUserDataOptions model with no property values
				deleteUserDataOptionsModelNew := new(texttospeechv1.DeleteUserDataOptions)
				// Invoke operation with invalid model (negative test)
				response, operationErr = textToSpeechService.DeleteUserData(deleteUserDataOptionsModelNew)
				Expect(operationErr).ToNot(BeNil())
				Expect(response).To(BeNil())
			})
			AfterEach(func() {
				testServer.Close()
			})
		})
	})
	Describe(`Model constructor tests`, func() {
		Context(`Using a service client instance`, func() {
			textToSpeechService, _ := texttospeechv1.NewTextToSpeechV1(&texttospeechv1.TextToSpeechV1Options{
				URL:           "http://texttospeechv1modelgenerator.com",
				Authenticator: &core.NoAuthAuthenticator{},
			})
			It(`Invoke NewAddCustomPromptOptions successfully`, func() {
				// Construct an instance of the PromptMetadata model
				promptMetadataModel := new(texttospeechv1.PromptMetadata)
				Expect(promptMetadataModel).ToNot(BeNil())
				promptMetadataModel.PromptText = core.StringPtr("testString")
				promptMetadataModel.SpeakerID = core.StringPtr("testString")
				Expect(promptMetadataModel.PromptText).To(Equal(core.StringPtr("testString")))
				Expect(promptMetadataModel.SpeakerID).To(Equal(core.StringPtr("testString")))

				// Construct an instance of the AddCustomPromptOptions model
				customizationID := "testString"
				promptID := "testString"
				var metadata *texttospeechv1.PromptMetadata = nil
				file := CreateMockReader("This is a mock file.")
				addCustomPromptOptionsModel := textToSpeechService.NewAddCustomPromptOptions(customizationID, promptID, metadata, file)
				addCustomPromptOptionsModel.SetCustomizationID("testString")
				addCustomPromptOptionsModel.SetPromptID("testString")
				addCustomPromptOptionsModel.SetMetadata(promptMetadataModel)
				addCustomPromptOptionsModel.SetFile(CreateMockReader("This is a mock file."))
				addCustomPromptOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(addCustomPromptOptionsModel).ToNot(BeNil())
				Expect(addCustomPromptOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(addCustomPromptOptionsModel.PromptID).To(Equal(core.StringPtr("testString")))
				Expect(addCustomPromptOptionsModel.Metadata).To(Equal(promptMetadataModel))
				Expect(addCustomPromptOptionsModel.File).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(addCustomPromptOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewAddWordOptions successfully`, func() {
				// Construct an instance of the AddWordOptions model
				customizationID := "testString"
				word := "testString"
				addWordOptionsTranslation := "testString"
				addWordOptionsModel := textToSpeechService.NewAddWordOptions(customizationID, word, addWordOptionsTranslation)
				addWordOptionsModel.SetCustomizationID("testString")
				addWordOptionsModel.SetWord("testString")
				addWordOptionsModel.SetTranslation("testString")
				addWordOptionsModel.SetPartOfSpeech("Dosi")
				addWordOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(addWordOptionsModel).ToNot(BeNil())
				Expect(addWordOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(addWordOptionsModel.Word).To(Equal(core.StringPtr("testString")))
				Expect(addWordOptionsModel.Translation).To(Equal(core.StringPtr("testString")))
				Expect(addWordOptionsModel.PartOfSpeech).To(Equal(core.StringPtr("Dosi")))
				Expect(addWordOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewAddWordsOptions successfully`, func() {
				// Construct an instance of the Word model
				wordModel := new(texttospeechv1.Word)
				Expect(wordModel).ToNot(BeNil())
				wordModel.Word = core.StringPtr("testString")
				wordModel.Translation = core.StringPtr("testString")
				wordModel.PartOfSpeech = core.StringPtr("Dosi")
				Expect(wordModel.Word).To(Equal(core.StringPtr("testString")))
				Expect(wordModel.Translation).To(Equal(core.StringPtr("testString")))
				Expect(wordModel.PartOfSpeech).To(Equal(core.StringPtr("Dosi")))

				// Construct an instance of the AddWordsOptions model
				customizationID := "testString"
				addWordsOptionsWords := []texttospeechv1.Word{}
				addWordsOptionsModel := textToSpeechService.NewAddWordsOptions(customizationID, addWordsOptionsWords)
				addWordsOptionsModel.SetCustomizationID("testString")
				addWordsOptionsModel.SetWords([]texttospeechv1.Word{*wordModel})
				addWordsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(addWordsOptionsModel).ToNot(BeNil())
				Expect(addWordsOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(addWordsOptionsModel.Words).To(Equal([]texttospeechv1.Word{*wordModel}))
				Expect(addWordsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateCustomModelOptions successfully`, func() {
				// Construct an instance of the CreateCustomModelOptions model
				createCustomModelOptionsName := "testString"
				createCustomModelOptionsModel := textToSpeechService.NewCreateCustomModelOptions(createCustomModelOptionsName)
				createCustomModelOptionsModel.SetName("testString")
				createCustomModelOptionsModel.SetLanguage("en-US")
				createCustomModelOptionsModel.SetDescription("testString")
				createCustomModelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createCustomModelOptionsModel).ToNot(BeNil())
				Expect(createCustomModelOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(createCustomModelOptionsModel.Language).To(Equal(core.StringPtr("en-US")))
				Expect(createCustomModelOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(createCustomModelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewCreateSpeakerModelOptions successfully`, func() {
				// Construct an instance of the CreateSpeakerModelOptions model
				speakerName := "testString"
				audio := CreateMockReader("This is a mock file.")
				createSpeakerModelOptionsModel := textToSpeechService.NewCreateSpeakerModelOptions(speakerName, audio)
				createSpeakerModelOptionsModel.SetSpeakerName("testString")
				createSpeakerModelOptionsModel.SetAudio(CreateMockReader("This is a mock file."))
				createSpeakerModelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(createSpeakerModelOptionsModel).ToNot(BeNil())
				Expect(createSpeakerModelOptionsModel.SpeakerName).To(Equal(core.StringPtr("testString")))
				Expect(createSpeakerModelOptionsModel.Audio).To(Equal(CreateMockReader("This is a mock file.")))
				Expect(createSpeakerModelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteCustomModelOptions successfully`, func() {
				// Construct an instance of the DeleteCustomModelOptions model
				customizationID := "testString"
				deleteCustomModelOptionsModel := textToSpeechService.NewDeleteCustomModelOptions(customizationID)
				deleteCustomModelOptionsModel.SetCustomizationID("testString")
				deleteCustomModelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteCustomModelOptionsModel).ToNot(BeNil())
				Expect(deleteCustomModelOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCustomModelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteCustomPromptOptions successfully`, func() {
				// Construct an instance of the DeleteCustomPromptOptions model
				customizationID := "testString"
				promptID := "testString"
				deleteCustomPromptOptionsModel := textToSpeechService.NewDeleteCustomPromptOptions(customizationID, promptID)
				deleteCustomPromptOptionsModel.SetCustomizationID("testString")
				deleteCustomPromptOptionsModel.SetPromptID("testString")
				deleteCustomPromptOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteCustomPromptOptionsModel).ToNot(BeNil())
				Expect(deleteCustomPromptOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCustomPromptOptionsModel.PromptID).To(Equal(core.StringPtr("testString")))
				Expect(deleteCustomPromptOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteSpeakerModelOptions successfully`, func() {
				// Construct an instance of the DeleteSpeakerModelOptions model
				speakerID := "testString"
				deleteSpeakerModelOptionsModel := textToSpeechService.NewDeleteSpeakerModelOptions(speakerID)
				deleteSpeakerModelOptionsModel.SetSpeakerID("testString")
				deleteSpeakerModelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteSpeakerModelOptionsModel).ToNot(BeNil())
				Expect(deleteSpeakerModelOptionsModel.SpeakerID).To(Equal(core.StringPtr("testString")))
				Expect(deleteSpeakerModelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteUserDataOptions successfully`, func() {
				// Construct an instance of the DeleteUserDataOptions model
				customerID := "testString"
				deleteUserDataOptionsModel := textToSpeechService.NewDeleteUserDataOptions(customerID)
				deleteUserDataOptionsModel.SetCustomerID("testString")
				deleteUserDataOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteUserDataOptionsModel).ToNot(BeNil())
				Expect(deleteUserDataOptionsModel.CustomerID).To(Equal(core.StringPtr("testString")))
				Expect(deleteUserDataOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewDeleteWordOptions successfully`, func() {
				// Construct an instance of the DeleteWordOptions model
				customizationID := "testString"
				word := "testString"
				deleteWordOptionsModel := textToSpeechService.NewDeleteWordOptions(customizationID, word)
				deleteWordOptionsModel.SetCustomizationID("testString")
				deleteWordOptionsModel.SetWord("testString")
				deleteWordOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(deleteWordOptionsModel).ToNot(BeNil())
				Expect(deleteWordOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(deleteWordOptionsModel.Word).To(Equal(core.StringPtr("testString")))
				Expect(deleteWordOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCustomModelOptions successfully`, func() {
				// Construct an instance of the GetCustomModelOptions model
				customizationID := "testString"
				getCustomModelOptionsModel := textToSpeechService.NewGetCustomModelOptions(customizationID)
				getCustomModelOptionsModel.SetCustomizationID("testString")
				getCustomModelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCustomModelOptionsModel).ToNot(BeNil())
				Expect(getCustomModelOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(getCustomModelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetCustomPromptOptions successfully`, func() {
				// Construct an instance of the GetCustomPromptOptions model
				customizationID := "testString"
				promptID := "testString"
				getCustomPromptOptionsModel := textToSpeechService.NewGetCustomPromptOptions(customizationID, promptID)
				getCustomPromptOptionsModel.SetCustomizationID("testString")
				getCustomPromptOptionsModel.SetPromptID("testString")
				getCustomPromptOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getCustomPromptOptionsModel).ToNot(BeNil())
				Expect(getCustomPromptOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(getCustomPromptOptionsModel.PromptID).To(Equal(core.StringPtr("testString")))
				Expect(getCustomPromptOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetPronunciationOptions successfully`, func() {
				// Construct an instance of the GetPronunciationOptions model
				text := "testString"
				getPronunciationOptionsModel := textToSpeechService.NewGetPronunciationOptions(text)
				getPronunciationOptionsModel.SetText("testString")
				getPronunciationOptionsModel.SetVoice("en-US_MichaelV3Voice")
				getPronunciationOptionsModel.SetFormat("ipa")
				getPronunciationOptionsModel.SetCustomizationID("testString")
				getPronunciationOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getPronunciationOptionsModel).ToNot(BeNil())
				Expect(getPronunciationOptionsModel.Text).To(Equal(core.StringPtr("testString")))
				Expect(getPronunciationOptionsModel.Voice).To(Equal(core.StringPtr("en-US_MichaelV3Voice")))
				Expect(getPronunciationOptionsModel.Format).To(Equal(core.StringPtr("ipa")))
				Expect(getPronunciationOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(getPronunciationOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetSpeakerModelOptions successfully`, func() {
				// Construct an instance of the GetSpeakerModelOptions model
				speakerID := "testString"
				getSpeakerModelOptionsModel := textToSpeechService.NewGetSpeakerModelOptions(speakerID)
				getSpeakerModelOptionsModel.SetSpeakerID("testString")
				getSpeakerModelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getSpeakerModelOptionsModel).ToNot(BeNil())
				Expect(getSpeakerModelOptionsModel.SpeakerID).To(Equal(core.StringPtr("testString")))
				Expect(getSpeakerModelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetVoiceOptions successfully`, func() {
				// Construct an instance of the GetVoiceOptions model
				voice := "ar-AR_OmarVoice"
				getVoiceOptionsModel := textToSpeechService.NewGetVoiceOptions(voice)
				getVoiceOptionsModel.SetVoice("ar-AR_OmarVoice")
				getVoiceOptionsModel.SetCustomizationID("testString")
				getVoiceOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getVoiceOptionsModel).ToNot(BeNil())
				Expect(getVoiceOptionsModel.Voice).To(Equal(core.StringPtr("ar-AR_OmarVoice")))
				Expect(getVoiceOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(getVoiceOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewGetWordOptions successfully`, func() {
				// Construct an instance of the GetWordOptions model
				customizationID := "testString"
				word := "testString"
				getWordOptionsModel := textToSpeechService.NewGetWordOptions(customizationID, word)
				getWordOptionsModel.SetCustomizationID("testString")
				getWordOptionsModel.SetWord("testString")
				getWordOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(getWordOptionsModel).ToNot(BeNil())
				Expect(getWordOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(getWordOptionsModel.Word).To(Equal(core.StringPtr("testString")))
				Expect(getWordOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListCustomModelsOptions successfully`, func() {
				// Construct an instance of the ListCustomModelsOptions model
				listCustomModelsOptionsModel := textToSpeechService.NewListCustomModelsOptions()
				listCustomModelsOptionsModel.SetLanguage("ar-MS")
				listCustomModelsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listCustomModelsOptionsModel).ToNot(BeNil())
				Expect(listCustomModelsOptionsModel.Language).To(Equal(core.StringPtr("ar-MS")))
				Expect(listCustomModelsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListCustomPromptsOptions successfully`, func() {
				// Construct an instance of the ListCustomPromptsOptions model
				customizationID := "testString"
				listCustomPromptsOptionsModel := textToSpeechService.NewListCustomPromptsOptions(customizationID)
				listCustomPromptsOptionsModel.SetCustomizationID("testString")
				listCustomPromptsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listCustomPromptsOptionsModel).ToNot(BeNil())
				Expect(listCustomPromptsOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(listCustomPromptsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListSpeakerModelsOptions successfully`, func() {
				// Construct an instance of the ListSpeakerModelsOptions model
				listSpeakerModelsOptionsModel := textToSpeechService.NewListSpeakerModelsOptions()
				listSpeakerModelsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listSpeakerModelsOptionsModel).ToNot(BeNil())
				Expect(listSpeakerModelsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListVoicesOptions successfully`, func() {
				// Construct an instance of the ListVoicesOptions model
				listVoicesOptionsModel := textToSpeechService.NewListVoicesOptions()
				listVoicesOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listVoicesOptionsModel).ToNot(BeNil())
				Expect(listVoicesOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewListWordsOptions successfully`, func() {
				// Construct an instance of the ListWordsOptions model
				customizationID := "testString"
				listWordsOptionsModel := textToSpeechService.NewListWordsOptions(customizationID)
				listWordsOptionsModel.SetCustomizationID("testString")
				listWordsOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(listWordsOptionsModel).ToNot(BeNil())
				Expect(listWordsOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(listWordsOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewPromptMetadata successfully`, func() {
				promptText := "testString"
				_model, err := textToSpeechService.NewPromptMetadata(promptText)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewSynthesizeOptions successfully`, func() {
				// Construct an instance of the SynthesizeOptions model
				synthesizeOptionsText := "testString"
				synthesizeOptionsModel := textToSpeechService.NewSynthesizeOptions(synthesizeOptionsText)
				synthesizeOptionsModel.SetText("testString")
				synthesizeOptionsModel.SetAccept("audio/ogg;codecs=opus")
				synthesizeOptionsModel.SetVoice("en-US_MichaelV3Voice")
				synthesizeOptionsModel.SetCustomizationID("testString")
				synthesizeOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(synthesizeOptionsModel).ToNot(BeNil())
				Expect(synthesizeOptionsModel.Text).To(Equal(core.StringPtr("testString")))
				Expect(synthesizeOptionsModel.Accept).To(Equal(core.StringPtr("audio/ogg;codecs=opus")))
				Expect(synthesizeOptionsModel.Voice).To(Equal(core.StringPtr("en-US_MichaelV3Voice")))
				Expect(synthesizeOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(synthesizeOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewTranslation successfully`, func() {
				translation := "testString"
				_model, err := textToSpeechService.NewTranslation(translation)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewUpdateCustomModelOptions successfully`, func() {
				// Construct an instance of the Word model
				wordModel := new(texttospeechv1.Word)
				Expect(wordModel).ToNot(BeNil())
				wordModel.Word = core.StringPtr("testString")
				wordModel.Translation = core.StringPtr("testString")
				wordModel.PartOfSpeech = core.StringPtr("Dosi")
				Expect(wordModel.Word).To(Equal(core.StringPtr("testString")))
				Expect(wordModel.Translation).To(Equal(core.StringPtr("testString")))
				Expect(wordModel.PartOfSpeech).To(Equal(core.StringPtr("Dosi")))

				// Construct an instance of the UpdateCustomModelOptions model
				customizationID := "testString"
				updateCustomModelOptionsModel := textToSpeechService.NewUpdateCustomModelOptions(customizationID)
				updateCustomModelOptionsModel.SetCustomizationID("testString")
				updateCustomModelOptionsModel.SetName("testString")
				updateCustomModelOptionsModel.SetDescription("testString")
				updateCustomModelOptionsModel.SetWords([]texttospeechv1.Word{*wordModel})
				updateCustomModelOptionsModel.SetHeaders(map[string]string{"foo": "bar"})
				Expect(updateCustomModelOptionsModel).ToNot(BeNil())
				Expect(updateCustomModelOptionsModel.CustomizationID).To(Equal(core.StringPtr("testString")))
				Expect(updateCustomModelOptionsModel.Name).To(Equal(core.StringPtr("testString")))
				Expect(updateCustomModelOptionsModel.Description).To(Equal(core.StringPtr("testString")))
				Expect(updateCustomModelOptionsModel.Words).To(Equal([]texttospeechv1.Word{*wordModel}))
				Expect(updateCustomModelOptionsModel.Headers).To(Equal(map[string]string{"foo": "bar"}))
			})
			It(`Invoke NewWord successfully`, func() {
				word := "testString"
				translation := "testString"
				_model, err := textToSpeechService.NewWord(word, translation)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
			It(`Invoke NewWords successfully`, func() {
				words := []texttospeechv1.Word{}
				_model, err := textToSpeechService.NewWords(words)
				Expect(_model).ToNot(BeNil())
				Expect(err).To(BeNil())
			})
		})
	})
	Describe(`Utility function tests`, func() {
		It(`Invoke CreateMockByteArray() successfully`, func() {
			mockByteArray := CreateMockByteArray("This is a test")
			Expect(mockByteArray).ToNot(BeNil())
		})
		It(`Invoke CreateMockUUID() successfully`, func() {
			mockUUID := CreateMockUUID("9fab83da-98cb-4f18-a7ba-b6f0435c9673")
			Expect(mockUUID).ToNot(BeNil())
		})
		It(`Invoke CreateMockReader() successfully`, func() {
			mockReader := CreateMockReader("This is a test.")
			Expect(mockReader).ToNot(BeNil())
		})
		It(`Invoke CreateMockDate() successfully`, func() {
			mockDate := CreateMockDate("2019-01-01")
			Expect(mockDate).ToNot(BeNil())
		})
		It(`Invoke CreateMockDateTime() successfully`, func() {
			mockDateTime := CreateMockDateTime("2019-01-01T12:00:00.000Z")
			Expect(mockDateTime).ToNot(BeNil())
		})
	})
})

//
// Utility functions used by the generated test code
//

func CreateMockByteArray(mockData string) *[]byte {
	ba := make([]byte, 0)
	ba = append(ba, mockData...)
	return &ba
}

func CreateMockUUID(mockData string) *strfmt.UUID {
	uuid := strfmt.UUID(mockData)
	return &uuid
}

func CreateMockReader(mockData string) io.ReadCloser {
	return ioutil.NopCloser(bytes.NewReader([]byte(mockData)))
}

func CreateMockDate(mockData string) *strfmt.Date {
	d, err := core.ParseDate(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func CreateMockDateTime(mockData string) *strfmt.DateTime {
	d, err := core.ParseDateTime(mockData)
	if err != nil {
		return nil
	}
	return &d
}

func SetTestEnvironment(testEnvironment map[string]string) {
	for key, value := range testEnvironment {
		os.Setenv(key, value)
	}
}

func ClearTestEnvironment(testEnvironment map[string]string) {
	for key := range testEnvironment {
		os.Unsetenv(key)
	}
}
