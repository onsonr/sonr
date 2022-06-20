---
title: Registry
slug: VWTR-registry
createdAt: 2022-04-29T18:20:10.000Z
updatedAt: 2022-05-02T21:55:51.000Z
---
#Registry
##Overview
The Registry keeps track of plain text names, which are then used to create a `WhoIs` instance. A `WhoIs` repersents either an  application or user registered on the Sonr network. On registration the provided name is postifxed with `.snr` . this name is unique once registered.  Names cannot be unregistered. meaning once they are entered into the registry, it then a registered snr. Registration of names will be available on `main` network only. on `test` nework, users will be given an alphanumeric string postfixed with `snr` as their developer domain name. these names will not accrue a cost and will be free for application developers.

##Usage

### Registering a New Application

The following is an example of a request to register a new application

```
MsgRegisterApplication {
  string creator;                   // Creator is the account address of the creator of the Application. 
  Credential credential;            // Client side JSON Web Token for AssertionMethod
  string application_name;          // Application Name is the endpoint of the Application.
  string application_description;   //optional  // Application Description is the description of the Application.
  string application_url;           //optional // Application URL is the URL of the Application.
  string application_category;      //optional // Application Category is the category of the Application.
}
```



### Response Definition from RegisterApplication

The following is an example of a response after registering a new application

```
MsgRegisterApplicationResponse {
    int32 code = 1;// Code of the response
    string message = 2;// Message of the response
    WhoIs who_is = 3;// WhoIs for the registered name
    Session session = 4;// Session returns the session for the name
}
```



### Registering a New Name

The following is an example of a request to register a new name

```
MsgRegisterName {
  string creator;                               // Account address of the name owner  
  string name_to_register;                      // Selected Name to register
  Credential credential;                        // Client side JSON Web Token for AssertionMethod   
  map<string, string> metadata; // optional     // The Updated Metadata
}
```



### Response Definition from RegisterName

The following is an example of a response after registering a new name

```
MsgRegisterNameResponse {
    int32 code;        // Code of the response    
    string message;    // Message of the response
    WhoIs who_is;      // WhoIs for the registered name  
    Session session;   // Session returns the session for the name
}
```



### How to Access a Name

The following is an example of a request to access a name

```
MsgAccessName {
  string creator;           // The account that is accessing the name  
  string name;              // The name to access  
  Credential credential;    // Client side JSON Web Token for AssertionMethod
}
```



### Response Definition from AccessName

The following is an example of a response after accessing a name

```
MsgAccessNameResponse {    
    int32 code;         // Code of the response
    string message;     // Message of the response
    WhoIs who_is;       // WhoIs for the registered name   
    Session session;    // Session returns the session for the name
}
```





### Updating a Name

The following is an example of a request to update a name

```
MsgUpdateName {
  string creator;                              // The account that owns the name.
  string did;                                  // The did of the peer to update the name of 
  Credential credential;                       // Client side JSON Web Token for AssertionMethod. For additional devices being linked.
  map<string, string> metadata; // optional    // The Updated Metadata  
  Session session;                             // Session returns the session for the name
}
```





### Response Definition from UpdateName

The following is an example of a response after updating a name

```
  MsgUpdateNameResponse {
    int32 code;        // Code of the response    
    string message;    // Message of the response    
    WhoIs who_is;      // WhoIs for the registered name
}
```



### How to Access an Application

The following is an example of a request to access an application

```
MsgAccessApplication {
  string creator;           // The account that is accessing the Application
  string app_name;          // The name of the Application to access
  Credential credential;    // Client side JSON Web Token for AssertionMethod
}
```



### Response Definition from AcessApplication

The following is an example of a response after accessing an application

```
MsgAccessApplicationResponse {
    int32 code;                      // Code of the response    
    string message;                  // Message of the response   
    map<string, string> metadata;    // Data of the response
    WhoIs who_is;                    // WhoIs for the registered name
    Session session;                 // Session returns the session for the name
}
```



### Updating an Application

The following is an example of a request to update an application

```
MsgUpdateApplication {
  string creator;                  // The account that owns the name.
  string did;                      // The name of the peer to update the Application details of
  map<string, string> metadata;    // The updated configuration for the Application
  Session session;                 // Session returns the session for the name
}
```





### Response Definition from UpdateApplication

The following is an example of a response after updating an application

```
MsgUpdateApplicationResponse {    
    int32 code;                      // Code of the response
    string message;                  // Message of the response 
    map<string, string> metadata;    // Data of the response
    WhoIs who_is;                    // WhoIs for the registered name
}
```


