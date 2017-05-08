# Getting started

## How to Build


* In order to successfully build and run your SDK files, you are required to have the following setup in your system:
    * **Go**  (Visit [https://golang.org/doc/install](https://golang.org/doc/install) for more details on how to install Go)
    * **Java VM** Version 8 or later
    * **Eclipse 4.6 (Neon)** or later ([http://www.eclipse.org/neon/](http://www.eclipse.org/neon/))
    * **GoClipse** setup within above installed Eclipse (Follow the instructions at [this link](https://github.com/GoClipse/goclipse/blob/latest/documentation/Installation.md#instructions) to setup GoClipse)
* Ensure that ```GOPATH``` environment variable is set in the system variables. If not, set it to your workspace directory where you will be adding your Go projects.
* The generated code uses unirest-go http library. Therefore, you will need internet access to resolve this dependency. If Go is properly installed and configured, run the following command to pull the dependency:

```Go
go get github.com/apimatic/unirest-go
```

This will install unirest-go in the ```GOPATH``` you specified in the system variables.

Now follow the steps mentioned below to build your SDK:

1. Open eclipse in the Go language perspective and click on the ```Import``` option in ```File``` menu.

![Importing SDK into Eclipse - Step 1](https://apidocs.io/illustration/go?step=import0)

2. Select ```General -> Existing Projects into Workspace``` option from the tree list.

![Importing SDK into Eclipse - Step 2](https://apidocs.io/illustration/go?step=import1)

3. In ```Select root directory```, provide path to the unzipped archive for the generated code. Once the path is set and the Project becomes visible under ```Projects``` click ```Finish```

![Importing SDK into Eclipse - Step 3](https://apidocs.io/illustration/go?step=import2&workspaceFolder=ZeroTier%20Central%20API-GoLang&projectName=zerotiercentralapi_lib)

4. The Go library will be imported and its files will be visible in the Project Explorer

![Importing SDK into Eclipse - Step 4](https://apidocs.io/illustration/go?step=import3&projectName=zerotiercentralapi_lib)

## How to Use

The following section explains how to use the ZerotiercentralapiLib library in a new project.

### 1. Add a new Test Project

Create a new project in Eclipse by ```File``` -> ```New``` -> ```Go Project```

![Add a new project in Eclipse](https://apidocs.io/illustration/go?step=createNewProject0)

Name the Project as ```Test``` and click ```Finish```

![Create a new Maven Project - Step 1](https://apidocs.io/illustration/go?step=createNewProject1)

Create a new directory in the ```src``` directory of this project

![Create a new Maven Project - Step 2](https://apidocs.io/illustration/go?step=createNewProject2&projectName=zerotiercentralapi_lib)

Name it ```test.com```

![Create a new Maven Project - Step 3](https://apidocs.io/illustration/go?step=createNewProject3&projectName=zerotiercentralapi_lib)

Now create a new file inside ```src/test.com```

![Create a new Maven Project - Step 4](https://apidocs.io/illustration/go?step=createNewProject4&projectName=zerotiercentralapi_lib)

Name it ```testsdk.go```

![Create a new Maven Project - Step 5](https://apidocs.io/illustration/go?step=createNewProject5&projectName=zerotiercentralapi_lib)

In this Go file, you can start adding code to initialize the client library. Sample code to initialize the client library and using its methods is given in the subsequent sections.

### 2. Configure the Test Project

You need to import your generated library in this project in order to make use of its functions. In order to import the library, you can add its path in the ```GOPATH``` for this project. Follow the below steps:

Right click on the project name and click on ```Properties```

![Adding dependency to the client library - Step 1](https://apidocs.io/illustration/go?step=testProject0&projectName=zerotiercentralapi_lib)

Choose ```Go Compiler``` from the side menu. Check ```Use project specific settings``` and uncheck ```Use same value as the GOPATH environment variable.```. By default, the GOPATH value from the environment variables will be visible in ```Eclipse GOPATH```. Do not remove this as this points to the unirest dependency.

![Adding dependency to the client library - Step 2](https://apidocs.io/illustration/go?step=testProject1)

Append the library path to this GOPATH

![Adding dependency to the client library - Step 3](https://apidocs.io/illustration/go?step=testProject2&workspaceFolder=ZeroTier%20Central%20API-GoLang)

Once the path is appended, click on ```OK```

### 3. Build the Test Project

Right click on the project name and click on ```Build Project```

![Build Project](https://apidocs.io/illustration/go?step=buildProject&projectName=zerotiercentralapi_lib)

### 4. Run the Test Project

If the build is successful, right click on your Go file and click on ```Run As``` -> ```Go Application```

![Run Project](https://apidocs.io/illustration/go?step=runProject&projectName=zerotiercentralapi_lib)

## Initialization

### Authentication
In order to setup authentication of the API client, you need the following information.

| Parameter | Description |
|-----------|-------------|
| basicAuthUserName | The username to use with basic authentication |
| basicAuthPassword | The password to use with basic authentication |


To configure these for your generated code, open the file "Configuration.go" and edit it's contents.


# Class Reference

## <a name="list_of_controllers"></a>List of Controllers

* [generalqueries_pkg](#generalqueries_pkg)
* [user_pkg](#user_pkg)
* [network_pkg](#network_pkg)
* [member_pkg](#member_pkg)

## <a name="generalqueries_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".generalqueries_pkg") generalqueries_pkg

### Get instance

Factory for the ``` GENERALQUERIES ``` interface can be accessed from the package generalqueries_pkg.

```go
generalQueries := generalqueries_pkg.NewGENERALQUERIES()
```

### <a name="get_status_and_configuration_information"></a>![Method: ](https://apidocs.io/img/method.png ".generalqueries_pkg.GetStatusAndConfigurationInformation") GetStatusAndConfigurationInformation

> Get Status and Configuration Information


```go
func (me *GENERALQUERIES_IMPL) GetStatusAndConfigurationInformation()(*models_pkg.Status,error)
```

#### Example Usage

```go

var result *models_pkg.Status
result,_ = generalQueries.GetStatusAndConfigurationInformation()

```


### <a name="get_currently_authenticated_user"></a>![Method: ](https://apidocs.io/img/method.png ".generalqueries_pkg.GetCurrentlyAuthenticatedUser") GetCurrentlyAuthenticatedUser

> Get Currently Authenticated User


```go
func (me *GENERALQUERIES_IMPL) GetCurrentlyAuthenticatedUser()(string,error)
```

#### Example Usage

```go

var result string
result,_ = generalQueries.GetCurrentlyAuthenticatedUser()

```


### <a name="get_generate_a_random_token"></a>![Method: ](https://apidocs.io/img/method.png ".generalqueries_pkg.GetGenerateARandomToken") GetGenerateARandomToken

> Generate a Random Token


```go
func (me *GENERALQUERIES_IMPL) GetGenerateARandomToken()(*models_pkg.RandomToken,error)
```

#### Example Usage

```go

var result *models_pkg.RandomToken
result,_ = generalQueries.GetGenerateARandomToken()

```


### <a name="create_terminate_current_session"></a>![Method: ](https://apidocs.io/img/method.png ".generalqueries_pkg.CreateTerminateCurrentSession") CreateTerminateCurrentSession

> Terminate Current Session


```go
func (me *GENERALQUERIES_IMPL) CreateTerminateCurrentSession()(,error)
```

#### Example Usage

```go

var result 
result,_ = generalQueries.CreateTerminateCurrentSession()

```


[Back to List of Controllers](#list_of_controllers)

## <a name="user_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".user_pkg") user_pkg

### Get instance

Factory for the ``` USER ``` interface can be accessed from the package user_pkg.

```go
user := user_pkg.NewUSER()
```

### <a name="retrieve_a_user"></a>![Method: ](https://apidocs.io/img/method.png ".user_pkg.RetrieveAUser") RetrieveAUser

> Retrieve a User


```go
func (me *USER_IMPL) RetrieveAUser(userId string)(*models_pkg.User,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| userId |  ``` Required ```  | 0000-0000-0000-000000000000 (required,string) - Internal user ID (GUID) |


#### Example Usage

```go
userId := "00000000"

var result *models_pkg.User
result,_ = user.RetrieveAUser(userId)

```


### <a name="update_a_user"></a>![Method: ](https://apidocs.io/img/method.png ".user_pkg.UpdateAUser") UpdateAUser

> Update a User


```go
func (me *USER_IMPL) UpdateAUser(userId string)(*models_pkg.User,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| userId |  ``` Required ```  | 0000-0000-0000-000000000000 (required,string) - Internal user ID (GUID) |


#### Example Usage

```go
userId := "00000000"

var result *models_pkg.User
result,_ = user.UpdateAUser(userId)

```


### <a name="get_all_viewable_users"></a>![Method: ](https://apidocs.io/img/method.png ".user_pkg.GetAllViewableUsers") GetAllViewableUsers

> Get All Viewable Users


```go
func (me *USER_IMPL) GetAllViewableUsers()([]*models_pkg.User,error)
```

#### Example Usage

```go

var result []*models_pkg.User
result,_ = user.GetAllViewableUsers()

```


[Back to List of Controllers](#list_of_controllers)

## <a name="network_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".network_pkg") network_pkg

### Get instance

Factory for the ``` NETWORK ``` interface can be accessed from the package network_pkg.

```go
network := network_pkg.NewNETWORK()
```

### <a name="retrieve_a_network"></a>![Method: ](https://apidocs.io/img/method.png ".network_pkg.RetrieveANetwork") RetrieveANetwork

> Retrieve a Network


```go
func (me *NETWORK_IMPL) RetrieveANetwork(networkId string)(*models_pkg.Network,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| networkId |  ``` Required ```  | 16-digit ZeroTier network ID |


#### Example Usage

```go
networkId := "0000000000000000"

var result *models_pkg.Network
result,_ = network.RetrieveANetwork(networkId)

```


### <a name="update_or_create_a_network"></a>![Method: ](https://apidocs.io/img/method.png ".network_pkg.UpdateOrCreateANetwork") UpdateOrCreateANetwork

> Update or create a Network


```go
func (me *NETWORK_IMPL) UpdateOrCreateANetwork(
            networkId string,
            body *models_pkg.Network)(*models_pkg.Network,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| networkId |  ``` Required ```  | 16-digit ZeroTier network ID |
| body |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
networkId := "0000000000000000"
bodyValue := []byte("{  \"id\": \"\",  \"type\": \"\",  \"clock\": 0,  \"ui\": {},  \"rulesSource\": \"\",  \"description\": \"\",  \"permissions\": {    \"{id}\": {      \"r\": false,      \"m\": false,      \"d\": false,      \"a\": false,      \"o\": false,      \"t\": \"\"    }  },  \"onlineMemberCount\": 0,  \"capabilitiesByName\": {},  \"tagsByName\": {},  \"circuitTestEvery\": 0,  \"config\": {    \"id\": \"\",    \"nwid\": \"\",    \"name\": \"\",    \"objtype\": \"\",    \"private\": false,    \"creationTime\": 0,    \"revision\": 0,    \"lastModified\": 0,    \"multicastLimit\": 0,    \"routes\": [],    \"rules\": [],    \"tags\": [],    \"capabilities\": [],    \"totalMemberCount\": 0,    \"activeMemberCount\": 0,    \"authTokens\": [],    \"authorizedMemberCount\": 0,    \"v4AssignMode\": {},    \"v6AssignMode\": {}  }}")
var body *models_pkg.Network
json.Unmarshal(bodyValue,&body)

var result *models_pkg.Network
result,_ = network.UpdateOrCreateANetwork(networkId, body)

```


### <a name="delete_a_network"></a>![Method: ](https://apidocs.io/img/method.png ".network_pkg.DeleteANetwork") DeleteANetwork

> Delete a Network


```go
func (me *NETWORK_IMPL) DeleteANetwork(networkId string)(,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| networkId |  ``` Required ```  | 16-digit ZeroTier network ID |


#### Example Usage

```go
networkId := "0000000000000000"

var result 
result,_ = network.DeleteANetwork(networkId)

```


### <a name="get_all_viewable_networks"></a>![Method: ](https://apidocs.io/img/method.png ".network_pkg.GetAllViewableNetworks") GetAllViewableNetworks

> Get All Viewable Networks


```go
func (me *NETWORK_IMPL) GetAllViewableNetworks()([]*models_pkg.Network,error)
```

#### Example Usage

```go

var result []*models_pkg.Network
result,_ = network.GetAllViewableNetworks()

```


[Back to List of Controllers](#list_of_controllers)

## <a name="member_pkg"></a>![Class: ](https://apidocs.io/img/class.png ".member_pkg") member_pkg

### Get instance

Factory for the ``` MEMBER ``` interface can be accessed from the package member_pkg.

```go
member := member_pkg.NewMEMBER()
```

### <a name="retrieve_a_member"></a>![Method: ](https://apidocs.io/img/method.png ".member_pkg.RetrieveAMember") RetrieveAMember

> Retrieve a Member


```go
func (me *MEMBER_IMPL) RetrieveAMember(
            networkId string,
            nodeId string)(*models_pkg.Member,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| networkId |  ``` Required ```  | 16-digit ZeroTier network ID |
| nodeId |  ``` Required ```  | 10-digit ZeroTier node ID (a.k.a. ZeroTier address) |


#### Example Usage

```go
networkId := "networkId"
nodeId := "nodeId"

var result *models_pkg.Member
result,_ = member.RetrieveAMember(networkId, nodeId)

```


### <a name="update_or_add_a_member"></a>![Method: ](https://apidocs.io/img/method.png ".member_pkg.UpdateOrAddAMember") UpdateOrAddAMember

> Update or add a Member


```go
func (me *MEMBER_IMPL) UpdateOrAddAMember(
            networkId string,
            nodeId string,
            body *models_pkg.Member)(*models_pkg.Member,error)
```

#### Parameters

| Parameter | Tags | Description |
|-----------|------|-------------|
| networkId |  ``` Required ```  | 16-digit ZeroTier network ID |
| nodeId |  ``` Required ```  | 10-digit ZeroTier node ID (a.k.a. ZeroTier address) |
| body |  ``` Required ```  | TODO: Add a parameter description |


#### Example Usage

```go
networkId := "networkId"
nodeId := "nodeId"
var body *models_pkg.Member

var result *models_pkg.Member
result,_ = member.UpdateOrAddAMember(networkId, nodeId, body)

```


[Back to List of Controllers](#list_of_controllers)



