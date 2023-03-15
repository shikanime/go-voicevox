# LicenseInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Name** | **string** |  | 
**Version** | Pointer to **string** |  | [optional] 
**License** | Pointer to **string** |  | [optional] 
**Text** | **string** |  | 

## Methods

### NewLicenseInfo

`func NewLicenseInfo(name string, text string, ) *LicenseInfo`

NewLicenseInfo instantiates a new LicenseInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewLicenseInfoWithDefaults

`func NewLicenseInfoWithDefaults() *LicenseInfo`

NewLicenseInfoWithDefaults instantiates a new LicenseInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetName

`func (o *LicenseInfo) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *LicenseInfo) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *LicenseInfo) SetName(v string)`

SetName sets Name field to given value.


### GetVersion

`func (o *LicenseInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *LicenseInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *LicenseInfo) SetVersion(v string)`

SetVersion sets Version field to given value.

### HasVersion

`func (o *LicenseInfo) HasVersion() bool`

HasVersion returns a boolean if a field has been set.

### GetLicense

`func (o *LicenseInfo) GetLicense() string`

GetLicense returns the License field if non-nil, zero value otherwise.

### GetLicenseOk

`func (o *LicenseInfo) GetLicenseOk() (*string, bool)`

GetLicenseOk returns a tuple with the License field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetLicense

`func (o *LicenseInfo) SetLicense(v string)`

SetLicense sets License field to given value.

### HasLicense

`func (o *LicenseInfo) HasLicense() bool`

HasLicense returns a boolean if a field has been set.

### GetText

`func (o *LicenseInfo) GetText() string`

GetText returns the Text field if non-nil, zero value otherwise.

### GetTextOk

`func (o *LicenseInfo) GetTextOk() (*string, bool)`

GetTextOk returns a tuple with the Text field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetText

`func (o *LicenseInfo) SetText(v string)`

SetText sets Text field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


