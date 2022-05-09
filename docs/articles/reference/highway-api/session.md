---
title: Session
slug: dsi2-session
createdAt: 2022-04-26T14:49:33.000Z
updatedAt: 2022-04-27T21:55:57.000Z
---
#Sessions
##Overview
`Sessions` are used to approve a user's authenication status with each request made. For a request to be processed it must be associated with a valid session in relation to the channel. In essence, the Session makes sure that a user is approved to further interact with the Sonr Network. For more information on sessions see `Registry`

##Usage

```
Session {
	string base_did;        // Base DID is the current Account or Application whois DID url
	WhoIs whois; 	        // WhoIs is the current Document for the DID
	Credential credential; 	// Credential is the current Credential for the DID
}
```


