# SupportedDevicesInfo

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**Cpu** | **bool** |  | 
**Cuda** | **bool** |  | 
**Dml** | **bool** |  | 

## Methods

### NewSupportedDevicesInfo

`func NewSupportedDevicesInfo(cpu bool, cuda bool, dml bool, ) *SupportedDevicesInfo`

NewSupportedDevicesInfo instantiates a new SupportedDevicesInfo object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewSupportedDevicesInfoWithDefaults

`func NewSupportedDevicesInfoWithDefaults() *SupportedDevicesInfo`

NewSupportedDevicesInfoWithDefaults instantiates a new SupportedDevicesInfo object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetCpu

`func (o *SupportedDevicesInfo) GetCpu() bool`

GetCpu returns the Cpu field if non-nil, zero value otherwise.

### GetCpuOk

`func (o *SupportedDevicesInfo) GetCpuOk() (*bool, bool)`

GetCpuOk returns a tuple with the Cpu field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCpu

`func (o *SupportedDevicesInfo) SetCpu(v bool)`

SetCpu sets Cpu field to given value.


### GetCuda

`func (o *SupportedDevicesInfo) GetCuda() bool`

GetCuda returns the Cuda field if non-nil, zero value otherwise.

### GetCudaOk

`func (o *SupportedDevicesInfo) GetCudaOk() (*bool, bool)`

GetCudaOk returns a tuple with the Cuda field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetCuda

`func (o *SupportedDevicesInfo) SetCuda(v bool)`

SetCuda sets Cuda field to given value.


### GetDml

`func (o *SupportedDevicesInfo) GetDml() bool`

GetDml returns the Dml field if non-nil, zero value otherwise.

### GetDmlOk

`func (o *SupportedDevicesInfo) GetDmlOk() (*bool, bool)`

GetDmlOk returns a tuple with the Dml field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetDml

`func (o *SupportedDevicesInfo) SetDml(v bool)`

SetDml sets Dml field to given value.



[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


