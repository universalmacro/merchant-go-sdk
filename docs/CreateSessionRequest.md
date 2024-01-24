# CreateSessionRequest

## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**ShortMerchantId** | Pointer to **string** |  | [optional] 
**Account** | Pointer to **string** |  | [optional] 
**Password** | Pointer to **string** |  | [optional] 

## Methods

### NewCreateSessionRequest

`func NewCreateSessionRequest() *CreateSessionRequest`

NewCreateSessionRequest instantiates a new CreateSessionRequest object
This constructor will assign default values to properties that have it defined,
and makes sure properties required by API are set, but the set of arguments
will change when the set of required properties is changed

### NewCreateSessionRequestWithDefaults

`func NewCreateSessionRequestWithDefaults() *CreateSessionRequest`

NewCreateSessionRequestWithDefaults instantiates a new CreateSessionRequest object
This constructor will only assign default values to properties that have it defined,
but it doesn't guarantee that properties required by API are set

### GetShortMerchantId

`func (o *CreateSessionRequest) GetShortMerchantId() string`

GetShortMerchantId returns the ShortMerchantId field if non-nil, zero value otherwise.

### GetShortMerchantIdOk

`func (o *CreateSessionRequest) GetShortMerchantIdOk() (*string, bool)`

GetShortMerchantIdOk returns a tuple with the ShortMerchantId field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetShortMerchantId

`func (o *CreateSessionRequest) SetShortMerchantId(v string)`

SetShortMerchantId sets ShortMerchantId field to given value.

### HasShortMerchantId

`func (o *CreateSessionRequest) HasShortMerchantId() bool`

HasShortMerchantId returns a boolean if a field has been set.

### GetAccount

`func (o *CreateSessionRequest) GetAccount() string`

GetAccount returns the Account field if non-nil, zero value otherwise.

### GetAccountOk

`func (o *CreateSessionRequest) GetAccountOk() (*string, bool)`

GetAccountOk returns a tuple with the Account field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetAccount

`func (o *CreateSessionRequest) SetAccount(v string)`

SetAccount sets Account field to given value.

### HasAccount

`func (o *CreateSessionRequest) HasAccount() bool`

HasAccount returns a boolean if a field has been set.

### GetPassword

`func (o *CreateSessionRequest) GetPassword() string`

GetPassword returns the Password field if non-nil, zero value otherwise.

### GetPasswordOk

`func (o *CreateSessionRequest) GetPasswordOk() (*string, bool)`

GetPasswordOk returns a tuple with the Password field if it's non-nil, zero value otherwise
and a boolean to check if the value has been set.

### SetPassword

`func (o *CreateSessionRequest) SetPassword(v string)`

SetPassword sets Password field to given value.

### HasPassword

`func (o *CreateSessionRequest) HasPassword() bool`

HasPassword returns a boolean if a field has been set.


[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


