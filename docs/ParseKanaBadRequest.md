# ParseKanaBadRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Text** | **string** |  | 
**ErrorName** | **string** | |name|description| |---|---| | UNKNOWN_TEXT | 判別できない読み仮名があります: {text} | | ACCENT_TOP | 句頭にアクセントは置けません: {text} | | ACCENT_TWICE | 1つのアクセント句に二つ以上のアクセントは置けません: {text} | | ACCENT_NOTFOUND | アクセントを指定していないアクセント句があります: {text} | | EMPTY_PHRASE | {position}番目のアクセント句が空白です | | INTERROGATION_MARK_NOT_AT_END | アクセント句末以外に「？」は置けません: {text} | | INFINITE_LOOP | 処理時に無限ループになってしまいました...バグ報告をお願いします。 | | 
**ErrorArgs** | **map[string]string** |  | 

## Methods

### NewParseKanaBadRequest

`func NewParseKanaBadRequest(text string, errorName string, errorArgs map[string]string, ) *ParseKanaBadRequest`

NewParseKanaBadRequest instantiates a new ParseKanaBadRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewParseKanaBadRequestWithDefaults

`func NewParseKanaBadRequestWithDefaults() *ParseKanaBadRequest`

NewParseKanaBadRequestWithDefaults instantiates a new ParseKanaBadRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetText

`func (o *ParseKanaBadRequest) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *ParseKanaBadRequest) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *ParseKanaBadRequest) SetText(v string)`

SetText sets Text field to given value.


### GetErrorName

`func (o *ParseKanaBadRequest) GetErrorName() string`

GetErrorName returns the ErrorName field if non-nil, zero value otherwise.

### GetErrorNameOk

`func (o *ParseKanaBadRequest) GetErrorNameOk() (*string, bool)`

GetErrorNameOk returns a tuple with the ErrorName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorName

`func (o *ParseKanaBadRequest) SetErrorName(v string)`

SetErrorName sets ErrorName field to given value.


### GetErrorArgs

`func (o *ParseKanaBadRequest) GetErrorArgs() map[string]string`

GetErrorArgs returns the ErrorArgs field if non-nil, zero value otherwise.

### GetErrorArgsOk

`func (o *ParseKanaBadRequest) GetErrorArgsOk() (*map[string]string, bool)`

GetErrorArgsOk returns a tuple with the ErrorArgs field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetErrorArgs

`func (o *ParseKanaBadRequest) SetErrorArgs(v map[string]string)`

SetErrorArgs sets ErrorArgs field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


