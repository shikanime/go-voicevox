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

type QueryEditApi interface {

	/*
			AccentPhrasesAccentPhrasesPost テキストからアクセント句を得る

			テキストからアクセント句を得ます。
		is_kanaが`true`のとき、テキストは次のようなAquesTalkライクな記法に従う読み仮名として処理されます。デフォルトは`false`です。
		* 全てのカナはカタカナで記述される
		* アクセント句は`/`または`、`で区切る。`、`で区切った場合に限り無音区間が挿入される。
		* カナの手前に`_`を入れるとそのカナは無声化される
		* アクセント位置を`'`で指定する。全てのアクセント句にはアクセント位置を1つ指定する必要がある。
		* アクセント句末に`？`(全角)を入れることにより疑問文の発音ができる。

			@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
			@return QueryEditApiAccentPhrasesAccentPhrasesPostRequest
	*/
	AccentPhrasesAccentPhrasesPost(ctx context.Context) QueryEditApiAccentPhrasesAccentPhrasesPostRequest

	// AccentPhrasesAccentPhrasesPostExecute executes the request
	//  @return []AccentPhrase
	AccentPhrasesAccentPhrasesPostExecute(r QueryEditApiAccentPhrasesAccentPhrasesPostRequest) ([]AccentPhrase, *http.Response, error)

	/*
		MoraDataMoraDataPost アクセント句から音高・音素長を得る

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return QueryEditApiMoraDataMoraDataPostRequest
	*/
	MoraDataMoraDataPost(ctx context.Context) QueryEditApiMoraDataMoraDataPostRequest

	// MoraDataMoraDataPostExecute executes the request
	//  @return []AccentPhrase
	MoraDataMoraDataPostExecute(r QueryEditApiMoraDataMoraDataPostRequest) ([]AccentPhrase, *http.Response, error)

	/*
		MoraLengthMoraLengthPost アクセント句から音素長を得る

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return QueryEditApiMoraLengthMoraLengthPostRequest
	*/
	MoraLengthMoraLengthPost(ctx context.Context) QueryEditApiMoraLengthMoraLengthPostRequest

	// MoraLengthMoraLengthPostExecute executes the request
	//  @return []AccentPhrase
	MoraLengthMoraLengthPostExecute(r QueryEditApiMoraLengthMoraLengthPostRequest) ([]AccentPhrase, *http.Response, error)

	/*
		MoraPitchMoraPitchPost アクセント句から音高を得る

		@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
		@return QueryEditApiMoraPitchMoraPitchPostRequest
	*/
	MoraPitchMoraPitchPost(ctx context.Context) QueryEditApiMoraPitchMoraPitchPostRequest

	// MoraPitchMoraPitchPostExecute executes the request
	//  @return []AccentPhrase
	MoraPitchMoraPitchPostExecute(r QueryEditApiMoraPitchMoraPitchPostRequest) ([]AccentPhrase, *http.Response, error)
}

// QueryEditApiService QueryEditApi service
type QueryEditApiService service

type QueryEditApiAccentPhrasesAccentPhrasesPostRequest struct {
	ctx         context.Context
	ApiService  QueryEditApi
	text        *string
	speaker     *int32
	isKana      *bool
	coreVersion *string
}

func (r QueryEditApiAccentPhrasesAccentPhrasesPostRequest) Text(text string) QueryEditApiAccentPhrasesAccentPhrasesPostRequest {
	r.text = &text
	return r
}

func (r QueryEditApiAccentPhrasesAccentPhrasesPostRequest) Speaker(speaker int32) QueryEditApiAccentPhrasesAccentPhrasesPostRequest {
	r.speaker = &speaker
	return r
}

func (r QueryEditApiAccentPhrasesAccentPhrasesPostRequest) IsKana(isKana bool) QueryEditApiAccentPhrasesAccentPhrasesPostRequest {
	r.isKana = &isKana
	return r
}

func (r QueryEditApiAccentPhrasesAccentPhrasesPostRequest) CoreVersion(coreVersion string) QueryEditApiAccentPhrasesAccentPhrasesPostRequest {
	r.coreVersion = &coreVersion
	return r
}

func (r QueryEditApiAccentPhrasesAccentPhrasesPostRequest) Execute() ([]AccentPhrase, *http.Response, error) {
	return r.ApiService.AccentPhrasesAccentPhrasesPostExecute(r)
}

/*
AccentPhrasesAccentPhrasesPost テキストからアクセント句を得る

テキストからアクセント句を得ます。
is_kanaが`true`のとき、テキストは次のようなAquesTalkライクな記法に従う読み仮名として処理されます。デフォルトは`false`です。
* 全てのカナはカタカナで記述される
* アクセント句は`/`または`、`で区切る。`、`で区切った場合に限り無音区間が挿入される。
* カナの手前に`_`を入れるとそのカナは無声化される
* アクセント位置を`'`で指定する。全てのアクセント句にはアクセント位置を1つ指定する必要がある。
* アクセント句末に`？`(全角)を入れることにより疑問文の発音ができる。

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return QueryEditApiAccentPhrasesAccentPhrasesPostRequest
*/
func (a *QueryEditApiService) AccentPhrasesAccentPhrasesPost(ctx context.Context) QueryEditApiAccentPhrasesAccentPhrasesPostRequest {
	return QueryEditApiAccentPhrasesAccentPhrasesPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []AccentPhrase
func (a *QueryEditApiService) AccentPhrasesAccentPhrasesPostExecute(r QueryEditApiAccentPhrasesAccentPhrasesPostRequest) ([]AccentPhrase, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []AccentPhrase
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QueryEditApiService.AccentPhrasesAccentPhrasesPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/accent_phrases"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.text == nil {
		return localVarReturnValue, nil, reportError("text is required and must be specified")
	}
	if r.speaker == nil {
		return localVarReturnValue, nil, reportError("speaker is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "text", r.text, "")
	parameterAddToHeaderOrQuery(localVarQueryParams, "speaker", r.speaker, "")
	if r.isKana != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "is_kana", r.isKana, "")
	}
	if r.coreVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "core_version", r.coreVersion, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

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
		if localVarHTTPResponse.StatusCode == 400 {
			var v ParseKanaBadRequest
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
			newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
			newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
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

type QueryEditApiMoraDataMoraDataPostRequest struct {
	ctx          context.Context
	ApiService   QueryEditApi
	speaker      *int32
	accentPhrase *[]AccentPhrase
	coreVersion  *string
}

func (r QueryEditApiMoraDataMoraDataPostRequest) Speaker(speaker int32) QueryEditApiMoraDataMoraDataPostRequest {
	r.speaker = &speaker
	return r
}

func (r QueryEditApiMoraDataMoraDataPostRequest) AccentPhrase(accentPhrase []AccentPhrase) QueryEditApiMoraDataMoraDataPostRequest {
	r.accentPhrase = &accentPhrase
	return r
}

func (r QueryEditApiMoraDataMoraDataPostRequest) CoreVersion(coreVersion string) QueryEditApiMoraDataMoraDataPostRequest {
	r.coreVersion = &coreVersion
	return r
}

func (r QueryEditApiMoraDataMoraDataPostRequest) Execute() ([]AccentPhrase, *http.Response, error) {
	return r.ApiService.MoraDataMoraDataPostExecute(r)
}

/*
MoraDataMoraDataPost アクセント句から音高・音素長を得る

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return QueryEditApiMoraDataMoraDataPostRequest
*/
func (a *QueryEditApiService) MoraDataMoraDataPost(ctx context.Context) QueryEditApiMoraDataMoraDataPostRequest {
	return QueryEditApiMoraDataMoraDataPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []AccentPhrase
func (a *QueryEditApiService) MoraDataMoraDataPostExecute(r QueryEditApiMoraDataMoraDataPostRequest) ([]AccentPhrase, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []AccentPhrase
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QueryEditApiService.MoraDataMoraDataPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mora_data"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.speaker == nil {
		return localVarReturnValue, nil, reportError("speaker is required and must be specified")
	}
	if r.accentPhrase == nil {
		return localVarReturnValue, nil, reportError("accentPhrase is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "speaker", r.speaker, "")
	if r.coreVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "core_version", r.coreVersion, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.accentPhrase
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

type QueryEditApiMoraLengthMoraLengthPostRequest struct {
	ctx          context.Context
	ApiService   QueryEditApi
	speaker      *int32
	accentPhrase *[]AccentPhrase
	coreVersion  *string
}

func (r QueryEditApiMoraLengthMoraLengthPostRequest) Speaker(speaker int32) QueryEditApiMoraLengthMoraLengthPostRequest {
	r.speaker = &speaker
	return r
}

func (r QueryEditApiMoraLengthMoraLengthPostRequest) AccentPhrase(accentPhrase []AccentPhrase) QueryEditApiMoraLengthMoraLengthPostRequest {
	r.accentPhrase = &accentPhrase
	return r
}

func (r QueryEditApiMoraLengthMoraLengthPostRequest) CoreVersion(coreVersion string) QueryEditApiMoraLengthMoraLengthPostRequest {
	r.coreVersion = &coreVersion
	return r
}

func (r QueryEditApiMoraLengthMoraLengthPostRequest) Execute() ([]AccentPhrase, *http.Response, error) {
	return r.ApiService.MoraLengthMoraLengthPostExecute(r)
}

/*
MoraLengthMoraLengthPost アクセント句から音素長を得る

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return QueryEditApiMoraLengthMoraLengthPostRequest
*/
func (a *QueryEditApiService) MoraLengthMoraLengthPost(ctx context.Context) QueryEditApiMoraLengthMoraLengthPostRequest {
	return QueryEditApiMoraLengthMoraLengthPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []AccentPhrase
func (a *QueryEditApiService) MoraLengthMoraLengthPostExecute(r QueryEditApiMoraLengthMoraLengthPostRequest) ([]AccentPhrase, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []AccentPhrase
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QueryEditApiService.MoraLengthMoraLengthPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mora_length"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.speaker == nil {
		return localVarReturnValue, nil, reportError("speaker is required and must be specified")
	}
	if r.accentPhrase == nil {
		return localVarReturnValue, nil, reportError("accentPhrase is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "speaker", r.speaker, "")
	if r.coreVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "core_version", r.coreVersion, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.accentPhrase
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

type QueryEditApiMoraPitchMoraPitchPostRequest struct {
	ctx          context.Context
	ApiService   QueryEditApi
	speaker      *int32
	accentPhrase *[]AccentPhrase
	coreVersion  *string
}

func (r QueryEditApiMoraPitchMoraPitchPostRequest) Speaker(speaker int32) QueryEditApiMoraPitchMoraPitchPostRequest {
	r.speaker = &speaker
	return r
}

func (r QueryEditApiMoraPitchMoraPitchPostRequest) AccentPhrase(accentPhrase []AccentPhrase) QueryEditApiMoraPitchMoraPitchPostRequest {
	r.accentPhrase = &accentPhrase
	return r
}

func (r QueryEditApiMoraPitchMoraPitchPostRequest) CoreVersion(coreVersion string) QueryEditApiMoraPitchMoraPitchPostRequest {
	r.coreVersion = &coreVersion
	return r
}

func (r QueryEditApiMoraPitchMoraPitchPostRequest) Execute() ([]AccentPhrase, *http.Response, error) {
	return r.ApiService.MoraPitchMoraPitchPostExecute(r)
}

/*
MoraPitchMoraPitchPost アクセント句から音高を得る

	@param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
	@return QueryEditApiMoraPitchMoraPitchPostRequest
*/
func (a *QueryEditApiService) MoraPitchMoraPitchPost(ctx context.Context) QueryEditApiMoraPitchMoraPitchPostRequest {
	return QueryEditApiMoraPitchMoraPitchPostRequest{
		ApiService: a,
		ctx:        ctx,
	}
}

// Execute executes the request
//
//	@return []AccentPhrase
func (a *QueryEditApiService) MoraPitchMoraPitchPostExecute(r QueryEditApiMoraPitchMoraPitchPostRequest) ([]AccentPhrase, *http.Response, error) {
	var (
		localVarHTTPMethod  = http.MethodPost
		localVarPostBody    interface{}
		formFiles           []formFile
		localVarReturnValue []AccentPhrase
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "QueryEditApiService.MoraPitchMoraPitchPost")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/mora_pitch"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.speaker == nil {
		return localVarReturnValue, nil, reportError("speaker is required and must be specified")
	}
	if r.accentPhrase == nil {
		return localVarReturnValue, nil, reportError("accentPhrase is required and must be specified")
	}

	parameterAddToHeaderOrQuery(localVarQueryParams, "speaker", r.speaker, "")
	if r.coreVersion != nil {
		parameterAddToHeaderOrQuery(localVarQueryParams, "core_version", r.coreVersion, "")
	}
	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	// body params
	localVarPostBody = r.accentPhrase
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
