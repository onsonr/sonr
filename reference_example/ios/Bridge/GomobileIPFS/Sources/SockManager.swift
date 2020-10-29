//
//  SockManager.swift
//  Bridge
//
//  Created by Guilhem Fanton on 18/11/2019.
//

import Foundation
import Core

/// SockManagerError is a SockManger specific error (subclass of IPFSError)
public class SockManagerError: IPFSError {
    private static var code: Int = 5
    private static var subdomain: String = "SockManager"

    required init(_ description: String, _ optCause: NSError? = nil) {
        super.init(description, optCause, SockManagerError.subdomain, SockManagerError.code)
    }

    required init?(coder: NSCoder) {
        super.init(coder: coder)
    }
}

/// SockManager is a class that wraps a golang `sockManager` object
///
/// **Should not be used on its own**
public class SockManager {
    let sockManager: CoreSockManager

    /// Class constructor using sockBasePath passed as parameter as UDS base path
    /// - Parameter sockBasePath: The path where the UDS will be created
    /// - Throws: `SockManagerError`: If the initialization of the socket manager failed
    public init(_ sockBasePath: URL) throws {
        var err: NSError?

        if let sman = CoreNewSockManager(sockBasePath.path, &err) {
            self.sockManager = sman
        } else {
            throw SockManagerError("initialization failed", err)
        }
    }

    /// Creates an UDS and returns its path
    /// - Throws: `SockManagerError`: If the socket creation failed
    /// - Returns: The path of the created socket
    public func newSockPath() throws -> String {
        var err: NSError?

        let path = self.sockManager.newSockPath(&err)

        if err != nil {
            throw SockManagerError("socket path creation failed", err)
        }

        return path
    }
}
