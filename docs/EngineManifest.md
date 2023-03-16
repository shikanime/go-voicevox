# EngineManifest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ManifestVersion** | **string** |  | 
**Name** | **string** |  | 
**BrandName** | **string** |  | 
**Uuid** | **string** |  | 
**Url** | **string** |  | 
**Icon** | **string** |  | 
**DefaultSamplingRate** | **int32** |  | 
**TermsOfService** | **string** |  | 
**UpdateInfos** | [**[]UpdateInfo**](UpdateInfo.md) |  | 
**DependencyLicenses** | [**[]LicenseInfo**](LicenseInfo.md) |  | 
**SupportedFeatures** | [**SupportedFeatures**](SupportedFeatures.md) |  | 

## Methods

### NewEngineManifest

`func NewEngineManifest(manifestVersion string, name string, brandName string, uuid string, url string, icon string, defaultSamplingRate int32, termsOfService string, updateInfos []UpdateInfo, dependencyLicenses []LicenseInfo, supportedFeatures SupportedFeatures, ) *EngineManifest`

NewEngineManifest instantiates a new EngineManifest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewEngineManifestWithDefaults

`func NewEngineManifestWithDefaults() *EngineManifest`

NewEngineManifestWithDefaults instantiates a new EngineManifest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetManifestVersion

`func (o *EngineManifest) GetManifestVersion() string`

GetManifestVersion returns the ManifestVersion field if non-nil, zero value otherwise.

### GetManifestVersionOk

`func (o *EngineManifest) GetManifestVersionOk() (*string, bool)`

GetManifestVersionOk returns a tuple with the ManifestVersion field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetManifestVersion

`func (o *EngineManifest) SetManifestVersion(v string)`

SetManifestVersion sets ManifestVersion field to given value.


### GetName

`func (o *EngineManifest) GetName() string`

GetName returns the Name field if non-nil, zero value otherwise.

### GetNameOk

`func (o *EngineManifest) GetNameOk() (*string, bool)`

GetNameOk returns a tuple with the Name field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetName

`func (o *EngineManifest) SetName(v string)`

SetName sets Name field to given value.


### GetBrandName

`func (o *EngineManifest) GetBrandName() string`

GetBrandName returns the BrandName field if non-nil, zero value otherwise.

### GetBrandNameOk

`func (o *EngineManifest) GetBrandNameOk() (*string, bool)`

GetBrandNameOk returns a tuple with the BrandName field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetBrandName

`func (o *EngineManifest) SetBrandName(v string)`

SetBrandName sets BrandName field to given value.


### GetUuid

`func (o *EngineManifest) GetUuid() string`

GetUuid returns the Uuid field if non-nil, zero value otherwise.

### GetUuidOk

`func (o *EngineManifest) GetUuidOk() (*string, bool)`

GetUuidOk returns a tuple with the Uuid field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUuid

`func (o *EngineManifest) SetUuid(v string)`

SetUuid sets Uuid field to given value.


### GetUrl

`func (o *EngineManifest) GetUrl() string`

GetUrl returns the Url field if non-nil, zero value otherwise.

### GetUrlOk

`func (o *EngineManifest) GetUrlOk() (*string, bool)`

GetUrlOk returns a tuple with the Url field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUrl

`func (o *EngineManifest) SetUrl(v string)`

SetUrl sets Url field to given value.


### GetIcon

`func (o *EngineManifest) GetIcon() string`

GetIcon returns the Icon field if non-nil, zero value otherwise.

### GetIconOk

`func (o *EngineManifest) GetIconOk() (*string, bool)`

GetIconOk returns a tuple with the Icon field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetIcon

`func (o *EngineManifest) SetIcon(v string)`

SetIcon sets Icon field to given value.


### GetDefaultSamplingRate

`func (o *EngineManifest) GetDefaultSamplingRate() int32`

GetDefaultSamplingRate returns the DefaultSamplingRate field if non-nil, zero value otherwise.

### GetDefaultSamplingRateOk

`func (o *EngineManifest) GetDefaultSamplingRateOk() (*int32, bool)`

GetDefaultSamplingRateOk returns a tuple with the DefaultSamplingRate field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDefaultSamplingRate

`func (o *EngineManifest) SetDefaultSamplingRate(v int32)`

SetDefaultSamplingRate sets DefaultSamplingRate field to given value.


### GetTermsOfService

`func (o *EngineManifest) GetTermsOfService() string`

GetTermsOfService returns the TermsOfService field if non-nil, zero value otherwise.

### GetTermsOfServiceOk

`func (o *EngineManifest) GetTermsOfServiceOk() (*string, bool)`

GetTermsOfServiceOk returns a tuple with the TermsOfService field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetTermsOfService

`func (o *EngineManifest) SetTermsOfService(v string)`

SetTermsOfService sets TermsOfService field to given value.


### GetUpdateInfos

`func (o *EngineManifest) GetUpdateInfos() []UpdateInfo`

GetUpdateInfos returns the UpdateInfos field if non-nil, zero value otherwise.

### GetUpdateInfosOk

`func (o *EngineManifest) GetUpdateInfosOk() (*[]UpdateInfo, bool)`

GetUpdateInfosOk returns a tuple with the UpdateInfos field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetUpdateInfos

`func (o *EngineManifest) SetUpdateInfos(v []UpdateInfo)`

SetUpdateInfos sets UpdateInfos field to given value.


### GetDependencyLicenses

`func (o *EngineManifest) GetDependencyLicenses() []LicenseInfo`

GetDependencyLicenses returns the DependencyLicenses field if non-nil, zero value otherwise.

### GetDependencyLicensesOk

`func (o *EngineManifest) GetDependencyLicensesOk() (*[]LicenseInfo, bool)`

GetDependencyLicensesOk returns a tuple with the DependencyLicenses field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDependencyLicenses

`func (o *EngineManifest) SetDependencyLicenses(v []LicenseInfo)`

SetDependencyLicenses sets DependencyLicenses field to given value.


### GetSupportedFeatures

`func (o *EngineManifest) GetSupportedFeatures() SupportedFeatures`

GetSupportedFeatures returns the SupportedFeatures field if non-nil, zero value otherwise.

### GetSupportedFeaturesOk

`func (o *EngineManifest) GetSupportedFeaturesOk() (*SupportedFeatures, bool)`

GetSupportedFeaturesOk returns a tuple with the SupportedFeatures field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetSupportedFeatures

`func (o *EngineManifest) SetSupportedFeatures(v SupportedFeatures)`

SetSupportedFeatures sets SupportedFeatures field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


