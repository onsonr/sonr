//
//  Error.swift
//  Bridge
//
//  Created by Antoine Eddi on 11/21/19.
//

import Foundation

/// IPFSError is a base error for IPFS Bridge module (subclass of NSError)
public class IPFSError: NSError {
    private static let code: Int = 1
    private static let domain: String = "IPFS"

    private static func createUserInfo(_ description: String, _ optCause: NSError?) -> [String: Any] {
        var userInfo = ["NSLocalizedDescription": description]

        if let cause = optCause {
            userInfo["NSLocalizedFailureReason"] = cause.localizedDescription
        }

        return userInfo
    }

    required init(_ description: String, _ optCause: NSError? = nil) {
        super.init(
            domain: IPFSError.domain,
            code: IPFSError.code,
            userInfo: IPFSError.createUserInfo(description, optCause)
        )
    }

    init(_ description: String, _ optCause: NSError?, _ subdomain: String, _ code: Int?) {
        super.init(
            domain: "\(IPFSError.domain).\(subdomain)",
            code: code ?? IPFSError.code,
            userInfo: IPFSError.createUserInfo(description, optCause)
        )
    }

    public var localizedFullDescription: String {
        if let reason = self.localizedFailureReason {
            return "\(self.domain)(\(self.localizedDescription): \(reason))"
        }

        return "\(self.domain)(\(self.localizedDescription))"
    }

    required init?(coder: NSCoder) {
        super.init(coder: coder)
    }
}
