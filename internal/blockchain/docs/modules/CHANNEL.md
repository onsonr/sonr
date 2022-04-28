# Sonr - Channel Module

The Sonr channel module is used to store the records of the active pubsub topics associated with Applications powered by the Sonr Network. Each record contains an `ChannelDoc` which describes the Topic configuration and status of the channel. Each channel is required to have a set RegisteredType to pass as a payload with ChannelMessages.

## Overview

The record type utilized in the **Channel module** is the `HowIs` type. This type provides both an interface to utilize VerifiableCredentials and modify the ChannelDoc type definition

## Usage

> Blockchain Methods supplied by Channel Module. Full implementation is still a work in progress.

#### `CreateChannel()` - Records a new Channel configuration for a specified application on Sonr.

    - (`string`) Creator                : The Account Address signing this message
    - (`Session`) Session               : The Session for the authenticated user
    - (`string`) Label                  : Name of the channel defined by developer
    - (`string`) Description            : Description of the channel defined by developer
    - (`ObjectDoc`) ObjectToRegister    : The registered verified type to be sent in channel messages

## Status Codes

WIP
