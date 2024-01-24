# \MerchantAPI

All URIs are relative to *https://uat.api.universalmacro.com/core*

Method | HTTP request | Description
------------- | ------------- | -------------
[**CreateMerchant**](MerchantAPI.md#CreateMerchant) | **Post** /merchants | Create merchant



## CreateMerchant

> Merchant CreateMerchant(ctx).CreateMerchantRequest(createMerchantRequest).Execute()

Create merchant

### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	createMerchantRequest := *openapiclient.NewCreateMerchantRequest("ShortMerchantId_example", "Account_example", "Password_example") // CreateMerchantRequest |  (optional)

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.MerchantAPI.CreateMerchant(context.Background()).CreateMerchantRequest(createMerchantRequest).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `MerchantAPI.CreateMerchant``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `CreateMerchant`: Merchant
	fmt.Fprintf(os.Stdout, "Response from `MerchantAPI.CreateMerchant`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiCreateMerchantRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **createMerchantRequest** | [**CreateMerchantRequest**](CreateMerchantRequest.md) |  | 

### Return type

[**Merchant**](Merchant.md)

### Authorization

[ApiKeyAuth](../README.md#ApiKeyAuth)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

