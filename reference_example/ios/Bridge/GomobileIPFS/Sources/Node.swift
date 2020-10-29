//
//  Node.swift
//  Bridge
//
//  Created by Guilhem Fanton on 07/11/2019.
//

import Foundation
import Core

/// NodeError is a Node specific error (subclass of IPFSError)
public class NodeError: IPFSError {
    private static var code: Int = 3
    private static var subdomain: String = "NodeManager"

    required init(_ description: String, _ optCause: NSError? = nil) {
        super.init(description, optCause, NodeError.subdomain, NodeError.code)
    }

    required init?(coder: NSCoder) {
        super.init(coder: coder)
    }
}

/// Node is a class that wraps a golang `node` object
///
/// **Should not be used on its own**
public class Node {
    let node: CoreNode

    /// Class constructor using repo passed as parameter as node repo
    /// - Parameter repo: The repo of the node
    /// - Throws: `NodeError`: If the creation of the node failed
    public init(_ repo: Repo) throws {
        var err: NSError?

        if let node = CoreNewNode(repo.goRepo, &err) {
            self.node = node
        } else {
            throw NodeError("creation failed", err)
        }
    }

    /// Closes this node instance
    /// - Throws: `NodeError`: If the closing of the node failed
    public func close() throws {
        do {
            try self.node.close()
        } catch let error as NSError {
            throw NodeError("closing failed", error)
        }
    }

    /// Serves node API over UDS
    /// - Parameter onUDS: The UDS path to serve on
    /// - Throws: `NodeError`: If the node failed to serve
    public func serve(onUDS: String) throws {
        do {
            try self.node.serveUnixSocketAPI(onUDS)
        } catch let error as NSError {
            throw NodeError("unable to serve api on UDS", error)
        }
    }

    /// Serves node API over TCP
    /// - Parameter onTCPPort: The TCP port to serve on
    /// - Throws: `NodeError`: If the node failed to serve
    /// - Returns: The TCP/IP MultiAddr the node is serving on
    public func serve(onTCPPort: String) throws -> String {
        var err: NSError?

        let maddr = self.node.serveTCPAPI(onTCPPort, error: &err)

        if err != nil {
            throw NodeError("unable to serve api on TCP", err)
        }

        return maddr
    }
}
