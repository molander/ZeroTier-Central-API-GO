/*
 * zerotiercentralapi_lib
 *
 * This file was automatically generated by APIMATIC v2.0 ( https://apimatic.io )
 */
package network_pkg


import(
	"encoding/json"
	"zerotiercentralapi_lib/models_pkg"
	"github.com/apimatic/unirest-go"
	"zerotiercentralapi_lib"
	"zerotiercentralapi_lib/apihelper_pkg"
)
/*
 * Client structure as interface implementation
 */
type NETWORK_IMPL struct { }

/**
 * Retrieve a Network
 * @param    string        networkId     parameter: Required
 * @return	Returns the *models_pkg.Network response from the API call
 */
func (me *NETWORK_IMPL) RetrieveANetwork (
            networkId string) (*models_pkg.Network, error) {
        //the base uri for api requests
    _queryBuilder := zerotiercentralapi_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/api/network/{networkId}"

    //variable to hold errors
    var err error = nil
    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithTemplateParameters(_queryBuilder, map[string]interface{} {
        "networkId" : networkId,
    })
    if err != nil {
        //error in template param handling
        return nil, err
    }

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "accept" : "application/json",
    }

    //prepare API request
    _request := unirest.GetWithAuth(_queryBuilder, headers, zerotiercentralapi_lib.Config.BasicAuthUserName, zerotiercentralapi_lib.Config.BasicAuthPassword)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models_pkg.Network = &models_pkg.Network{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

/**
 * Update or create a Network
 * @param    string                     networkId     parameter: Required
 * @param    *models_pkg.Network        body          parameter: Required
 * @return	Returns the *models_pkg.Network response from the API call
 */
func (me *NETWORK_IMPL) UpdateOrCreateANetwork (
            networkId string,
            body *models_pkg.Network) (*models_pkg.Network, error) {
        //the base uri for api requests
    _queryBuilder := zerotiercentralapi_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/api/network/{networkId}"

    //variable to hold errors
    var err error = nil
    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithTemplateParameters(_queryBuilder, map[string]interface{} {
        "networkId" : networkId,
    })
    if err != nil {
        //error in template param handling
        return nil, err
    }

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "accept" : "application/json",
        "content-type" : "application/json; charset=utf-8",
    }

    //prepare API request
    _request := unirest.PostWithAuth(_queryBuilder, headers, body, zerotiercentralapi_lib.Config.BasicAuthUserName, zerotiercentralapi_lib.Config.BasicAuthPassword)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal *models_pkg.Network = &models_pkg.Network{}
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

/**
 * Delete a Network
 * @param    string        networkId     parameter: Required
 * @return	Returns the  response from the API call
 */
func (me *NETWORK_IMPL) DeleteANetwork (
            networkId string) (error) {
        //the base uri for api requests
    _queryBuilder := zerotiercentralapi_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/api/network/{networkId}"

    //variable to hold errors
    var err error = nil
    //process optional query parameters
    _queryBuilder, err = apihelper_pkg.AppendUrlWithTemplateParameters(_queryBuilder, map[string]interface{} {
        "networkId" : networkId,
    })
    if err != nil {
        //error in template param handling
        return err
    }

    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
    }

    //prepare API request
    _request := unirest.DeleteWithAuth(_queryBuilder, headers, nil, zerotiercentralapi_lib.Config.BasicAuthUserName, zerotiercentralapi_lib.Config.BasicAuthPassword)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return err
    }

    //returning the response
    return nil
}

/**
 * Get All Viewable Networks
 * @return	Returns the []*models_pkg.Network response from the API call
 */
func (me *NETWORK_IMPL) GetAllViewableNetworks () ([]*models_pkg.Network, error) {
        //the base uri for api requests
    _queryBuilder := zerotiercentralapi_lib.BASEURI;

    //prepare query string for API call
   _queryBuilder = _queryBuilder + "/api/network"

    //variable to hold errors
    var err error = nil
    //validate and preprocess url
    _queryBuilder, err = apihelper_pkg.CleanUrl(_queryBuilder)
    if err != nil {
        //error in url validation or cleaning
        return nil, err
    }

    //prepare headers for the outgoing request
    headers := map[string]interface{} {
        "user-agent" : "APIMATIC 2.0",
        "accept" : "application/json",
    }

    //prepare API request
    _request := unirest.GetWithAuth(_queryBuilder, headers, zerotiercentralapi_lib.Config.BasicAuthUserName, zerotiercentralapi_lib.Config.BasicAuthPassword)
    //and invoke the API call request to fetch the response
    _response, err := unirest.AsString(_request);
    if err != nil {
        //error in API invocation
        return nil, err
    }

    //error handling using HTTP status codes
    if (_response.Code < 200) || (_response.Code > 206) { //[200,206] = HTTP OK
        err = apihelper_pkg.NewAPIError("HTTP Response Not OK" , _response.Code, _response.RawBody)
    }
    if(err != nil) {
        //error detected in status code validation
        return nil, err
    }

    //returning the response
    var retVal []*models_pkg.Network
    err = json.Unmarshal(_response.RawBody, &retVal)

    if err != nil {
        //error in parsing
        return nil, err
    }
    return retVal, nil
}

