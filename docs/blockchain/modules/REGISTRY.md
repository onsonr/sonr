# Sonr - Registry Module

The Sonr registry module is used to store the records of user accounts and applications. Each record contains a DIDDocument and additional WebAuthn credential information.

## Overview

The record type utilized in the **Registry module** is the `WhoIs` type. This type provides both an interface to utilize WebAuthn credentials and a means to access the record's DIDDocument.

## Usage

> Blockchain Methods supplied by Registry Module. Full implementation is still a work in progress.

#### `RegisterName()` - Register's a new '.snr' domain name for an account

    - (`string`) NameToRegister     : The name to register
    - (`string`) Creator            : The Account Address signing this message
    - (`Credential`) Credential     : Webauthn credential to use for registration
    - (`map`) Metadata              : Metadata to attach to the `WhoIs` record

#### `RegisterApplication()` - Register's a new Application for the Sonr Network

    - (`string`) Creator                : The Account Address signing this message
    - (`Credential`) Credential         : Webauthn credential to use for registration
    - (`string`) ApplicationName        : The Name of the Application being registered
    - (`string`) ApplicationDescription : Short about description of the App
    - (`string`) ApplicationURL         : Website/Homepage of the App
    - (`string`) ApplicationCategory    : Category of the Application Type

#### `AccessName()` - Accesses a particular name essentially a "Login" function

    - (`string`) Creator            : The Account Address signing this message
    - (`Credential`) Credential     : Webauthn credential to use for registration
    - (`string`) Name               : The name to authenticate and retreive data

#### `AccessApplication()` - Accesses a particular application essentially a "Register" function

    - (`string`) Creator                : The Account Address signing this message
    - (`string`) AppName                : The Name of the Application being accessed
    - (`Credential`) Credential         : Webauthn Credential of the authenticated user

## Record Type: `WhoIs`

WIP

## Status Codes

WIP
