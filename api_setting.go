/*
VOICEVOX Engine

VOICEVOXの音声合成エンジンです。

API version: latest
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package voicevox

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)

type SettingApi interface {

	/*
		SettingGetSettingGet Setting Get

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return SettingApiSettingGetSettingGetRequest
	*/
	SettingGetSettingGet(ctx context.Context) SettingApiSettingGetSettingGetRequest

	// SettingGetSettingGetExecute executes the request
	//  @return string
	SettingGetSettingGetExecute(r SettingApiSettingGetSettingGetRequest) (string, *http.Response, error)

	/*
		SettingPostSettingPost Setting Post

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return SettingApiSettingPostSettingPostRequest
	*/
	SettingPostSettingPost(ctx context.Context) SettingApiSettingPostSettingPostRequest

	// SettingPostSettingPostExecute executes the request
	//  @return string
	SettingPostSettingPostExecute(r SettingApiSettingPostSettingPostRequest) (string, *http.Response, error)
}

// SettingApiService SettingApi service
type SettingApiService service

type SettingApiSettingGetSettingGetRequest struct {
	ctx        context.Context
	ApiService SettingApi
}

func (r SettingApiSettingGetSettingGetRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.SettingGetSettingGetExecute(r)
}

/*
SettingGetSettingGet Setting Get

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SettingApiSettingGetSettingGetRequest
*/
func (a *SettingApiService) SettingGetSettingGet(ctx context.Context) SettingApiSettingGetSettingGetRequest {
	return SettingApiSettingGetSettingGetRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return string
func (a *SettingApiService) SettingGetSettingGetExecute(r SettingApiSettingGetSettingGetRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodGet
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SettingApiService.SettingGetSettingGet")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/setting"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/html"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type SettingApiSettingPostSettingPostRequest struct {
	ctx            context.Context
	ApiService     SettingApi
	corsPolicyMode *string
	allowOrigin    *string
}

func (r SettingApiSettingPostSettingPostRequest) CorsPolicyMode(corsPolicyMode string) SettingApiSettingPostSettingPostRequest {
	r.corsPolicyMode = &corsPolicyMode
	return r
}

func (r SettingApiSettingPostSettingPostRequest) AllowOrigin(allowOrigin string) SettingApiSettingPostSettingPostRequest {
	r.allowOrigin = &allowOrigin
	return r
}

func (r SettingApiSettingPostSettingPostRequest) Execute() (string, *http.Response, error) {
	return r.ApiService.SettingPostSettingPostExecute(r)
}

/*
SettingPostSettingPost Setting Post

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return SettingApiSettingPostSettingPostRequest
*/
func (a *SettingApiService) SettingPostSettingPost(ctx context.Context) SettingApiSettingPostSettingPostRequest {
	return SettingApiSettingPostSettingPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return string
func (a *SettingApiService) SettingPostSettingPostExecute(r SettingApiSettingPostSettingPostRequest) (string, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue string
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "SettingApiService.SettingPostSettingPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/setting"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/x-www-form-urlencoded"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"text/html", "application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.corsPolicyMode != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "cors_policy_mode", r.corsPolicyMode, "")
	}
	if r.allowOrigin != nil {
		parameterAddToHeaderOrQuery(localVarFormParams, "allow_origin", r.allowOrigin, "")
	}
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 422 {
			var v HTTPValidationError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}