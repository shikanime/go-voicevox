# UpdateInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Version** | **string** |  | 
**Descriptions** | **[]string** |  | 
**Contributors** | Pointer to **[]string** |  | [optional] 

## Methods

### NewUpdateInfo

`func NewUpdateInfo(version string, descriptions []string, ) *UpdateInfo`

NewUpdateInfo instantiates a new UpdateInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewUpdateInfoWithDefaults

`func NewUpdateInfoWithDefaults() *UpdateInfo`

NewUpdateInfoWithDefaults instantiates a new UpdateInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetVersion

`func (o *UpdateInfo) GetVersion() string`

GetVersion returns the Version field if non-nil, zero value otherwise.

### GetVersionOk

`func (o *UpdateInfo) GetVersionOk() (*string, bool)`

GetVersionOk returns a tuple with the Version field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetVersion

`func (o *UpdateInfo) SetVersion(v string)`

SetVersion sets Version field to given value.


### GetDescriptions

`func (o *UpdateInfo) GetDescriptions() []string`

GetDescriptions returns the Descriptions field if non-nil, zero value otherwise.

### GetDescriptionsOk

`func (o *UpdateInfo) GetDescriptionsOk() (*[]string, bool)`

GetDescriptionsOk returns a tuple with the Descriptions field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDescriptions

`func (o *UpdateInfo) SetDescriptions(v []string)`

SetDescriptions sets Descriptions field to given value.


### GetContributors

`func (o *UpdateInfo) GetContributors() []string`

GetContributors returns the Contributors field if non-nil, zero value otherwise.

### GetContributorsOk

`func (o *UpdateInfo) GetContributorsOk() (*[]string, bool)`

GetContributorsOk returns a tuple with the Contributors field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetContributors

`func (o *UpdateInfo) SetContributors(v []string)`

SetContributors sets Contributors field to given value.

### HasContributors

`func (o *UpdateInfo) HasContributors() bool`

HasContributors returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


