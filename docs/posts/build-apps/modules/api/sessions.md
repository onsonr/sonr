---
title: Sessions
slug: GYaU-sessions
createdAt: 2022-04-27T18:21:38.000Z
updatedAt: 2022-04-27T18:21:43.000Z
---
#Sessions
##Overview
A `Session` is used to denote a user's authenication status with each request made. In order for a request to be processed it must be associated with a valid session in relation to the channel.  In essence, the Session makes sure that a user is valid to further interact with the Sonr Network. For more information on sessions see `Registry`

##Usage

```
Session {
	string base_did = 1;        // Base DID is the current Account or Application whois DID url
	WhoIs whois = 2; 	        // WhoIs is the current Document for the DID
	Credential credential = 3; 	// Credential is the current Credential for the DID
}
```


