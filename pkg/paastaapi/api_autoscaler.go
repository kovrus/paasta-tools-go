/*
 * Paasta API
 *
 * No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package paastaapi

import (
	_context "context"
	_ioutil "io/ioutil"
	_nethttp "net/http"
	_neturl "net/url"
	"strings"
)

// Linger please
var (
	_ _context.Context
)

// AutoscalerApiService AutoscalerApi service
type AutoscalerApiService service

type apiGetAutoscalerCountRequest struct {
	ctx _context.Context
	apiService *AutoscalerApiService
	service string
	instance string
}

/*
GetAutoscalerCount Get status of service_name.instance_name
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param service Service name
 * @param instance Instance name
@return apiGetAutoscalerCountRequest
*/
func (a *AutoscalerApiService) GetAutoscalerCount(ctx _context.Context, service string, instance string) apiGetAutoscalerCountRequest {
	return apiGetAutoscalerCountRequest{
		apiService: a,
		ctx: ctx,
		service: service,
		instance: instance,
	}
}

/*
Execute executes the request
@return InlineResponse2002
*/
func (r apiGetAutoscalerCountRequest) Execute() (InlineResponse2002, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodGet
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse2002
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "AutoscalerApiService.GetAutoscalerCount")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/services/{service}/{instance}/autoscaler"
	localVarPath = strings.Replace(localVarPath, "{"+"service"+"}", _neturl.PathEscape(parameterToString(r.service, "")) , -1)
	localVarPath = strings.Replace(localVarPath, "{"+"instance"+"}", _neturl.PathEscape(parameterToString(r.instance, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}

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
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = r.apiService.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
type apiUpdateAutoscalerCountRequest struct {
	ctx _context.Context
	apiService *AutoscalerApiService
	service string
	instance string
	inlineObject1 *InlineObject1
}

func (r apiUpdateAutoscalerCountRequest) InlineObject1(inlineObject1 InlineObject1) apiUpdateAutoscalerCountRequest {
	r.inlineObject1 = &inlineObject1
	return r
}
/*
UpdateAutoscalerCount Get status of service_name.instance_name
 * @param ctx _context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 * @param service Service name
 * @param instance Instance name
@return apiUpdateAutoscalerCountRequest
*/
func (a *AutoscalerApiService) UpdateAutoscalerCount(ctx _context.Context, service string, instance string) apiUpdateAutoscalerCountRequest {
	return apiUpdateAutoscalerCountRequest{
		apiService: a,
		ctx: ctx,
		service: service,
		instance: instance,
	}
}

/*
Execute executes the request
@return InlineResponse202
*/
func (r apiUpdateAutoscalerCountRequest) Execute() (InlineResponse202, *_nethttp.Response, error) {
	var (
		localVarHTTPMethod   = _nethttp.MethodPost
		localVarPostBody     interface{}
		localVarFormFileName string
		localVarFileName     string
		localVarFileBytes    []byte
		localVarReturnValue  InlineResponse202
	)

	localBasePath, err := r.apiService.client.cfg.ServerURLWithContext(r.ctx, "AutoscalerApiService.UpdateAutoscalerCount")
	if err != nil {
		return localVarReturnValue, nil, GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/services/{service}/{instance}/autoscaler"
	localVarPath = strings.Replace(localVarPath, "{"+"service"+"}", _neturl.PathEscape(parameterToString(r.service, "")) , -1)
	localVarPath = strings.Replace(localVarPath, "{"+"instance"+"}", _neturl.PathEscape(parameterToString(r.instance, "")) , -1)

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := _neturl.Values{}
	localVarFormParams := _neturl.Values{}
	if r.inlineObject1 == nil {
		return localVarReturnValue, nil, reportError("inlineObject1 is required and must be specified")
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
	localVarPostBody = r.inlineObject1
	req, err := r.apiService.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, localVarFormFileName, localVarFileName, localVarFileBytes)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := r.apiService.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := _ioutil.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = r.apiService.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
