//
//  IPFS.swift
//  Bridge
//
//  Created by Guilhem Fanton on 08/11/2019.
//

import Foundation
import Core

extension FileManager {
    public var compatTemporaryDirectory: URL {
        if #available(iOS 10.0, *) {
            return temporaryDirectory
        } else {
            return (try? url(
                for: .itemReplacementDirectory,
                in: .userDomainMask,
                appropriateFor: nil,
                create: true)
            ) ?? URL(fileURLWithPath: NSTemporaryDirectory())
        }
    }
}

/// IPFS is a class that wraps a go-ipfs node and its shell over UDS
public class IPFS {
    public static let defaultRepoPath = "ipfs/repo"

    private static var sockManager: SockManager?

    private var node: Node?
    private var shell: CoreShell?
    private var repo: Repo?

    private let absRepoURL: URL
    private let absSockPath: String

    /// Class constructor using repoPath passed as parameter on internal storage
    /// - Parameter repoPath: The path of the go-ipfs repo (default: `ipfs/repo`)
    /// - Throws:
    ///     - `SockManagerError`: If the initialization of SockManager failed
    ///     - `ConfigError`: If the creation of the config failed
    ///     - `RepoError`: If the initialization of the repo failed
    public init(_ repoPath: String = defaultRepoPath) throws {
        let absUserUrl = FileManager.default.urls(for: .documentDirectory, in: .userDomainMask).first!
        self.absRepoURL = absUserUrl.appendingPathComponent(repoPath, isDirectory: true)

        // Instantiate sockManager singleton if needed
        #if !targetEnvironment(simulator)
        if IPFS.sockManager == nil {
            let absTmpURL = FileManager.default.compatTemporaryDirectory
            IPFS.sockManager = try SockManager(absTmpURL)
        }

        self.absSockPath = try IPFS.sockManager!.newSockPath()
        #else // On simulator we can't create an UDS, see comment below
        self.absSockPath = ""
        #endif

        // Init IPFS Repo if not already initialized
        if !Repo.isInitialized(absRepoURL) {
            let config = try Config.defaultConfig()
            try Repo.initialize(absRepoURL, config)
        }
    }

    /// Returns the absolute repo path as an URL
    /// - Returns: The absolute repo path
    public func getAbsoluteRepoPath() -> URL {
        return self.absRepoURL
    }

    /// Returns True if this IPFS instance is started by checking if the underlying go-ipfs node is instantiated
    /// - Returns: True, if this IPFS instance is started
    public func isStarted() -> Bool {
        return self.node != nil
    }

    /// Starts this IPFS instance
    /// - Throws:
    ///     - `RepoError`: If the opening of the repo failed
    ///     - `NodeError`: If the node is already started or if its startup fails
    public func start() throws {
        if self.isStarted() {
            throw IPFSError("node already started")
        }

        // Open go-ipfs repo
        try openRepoIfClosed()

        // Instanciate the node
        self.node = try Node(self.repo!)

        // Create a shell over UDS on physical device
        #if !targetEnvironment(simulator)
        try self.node!.serve(onUDS: self.absSockPath)
        self.shell = CoreNewUDSShell(self.absSockPath)
        /*
        ** On iOS simulator, temporary directory's absolute path exceeds
        ** the length limit for Unix Domain Socket, since simulator is
        ** only used for debug, we can safely fallback on shell over TCP
        */
        #else
        let maddr: String = try self.node!.serve(onTCPPort: "0")
        self.shell = CoreNewShell(maddr)
        #endif
    }

    /// Stops this IPFS instance
    /// - Throws: `IPFSError`: If the node is already stopped or if its stop fails
    public func stop() throws {
        if !self.isStarted() {
            throw IPFSError("node already stopped")
        }

        try self.node?.close()
        self.node = nil
        self.repo = nil
    }

    /// Restarts this IPFS instance
    /// - Throws:
    ///     - `IPFSError`: If the node is already stopped or if its stop fails
    ///     - `RepoError`: If the opening of the repo failed
    public func restart() throws {
        try self.stop()
        try self.start()
    }

    /// Enable PubSub experimental feature on an IPFS node instance.
    /// - Attention: A started instance must be restarted for this feature to be enabled
    /// - Throws:
    ///     - `RepoError`: If the opening of the repo failed
    public func enablePubsubExperiment() throws {
        try openRepoIfClosed()
        self.repo!.goRepo.enablePubsubExperiment()
    }

    /// Enable PubSub experimental feature and IPNS record distribution through PubSub.
    /// - Attention: A started instance must be restarted for this feature to be enabled
    /// - Throws:
    ///     - `RepoError`: If the opening of the repo failed
    public func enableNamesysPubsub() throws {
        try openRepoIfClosed()
        self.repo!.goRepo.enableNamesysPubsub()
    }

    /// Gets the IPFS instance config as a dict
    /// - Throws:
    ///     - `RepoError`: If the opening of the repo or the getting of its config failed
    ///     - `ConfigError`: If the getting of the config as a dict failed
    /// - Returns: The IPFS instance config as a dict
    /// - seealso: [IPFS Config Doc](https://github.com/ipfs/go-ipfs/blob/master/docs/config.md)
    public func getConfig() throws -> [String: Any] {
        try openRepoIfClosed()

        return try repo!.getConfig().get()
    }

    /// Sets dict config passed as parameter as IPFS config or reset to default config (with a new identity)
    /// if the config parameter is nil
    /// - Attention: A started IPFS instance must be restarted for its config to be applied
    /// - Parameter config: The IPFS instance dict config to set (if nil, default config will be used)
    /// - Throws:
    ///     - `RepoError`: If the opening of the repo or the setting of its config failed
    ///     - `ConfigError`: If the setting of the config as a dict failed
    /// - seealso: [IPFS Config Doc](https://github.com/ipfs/go-ipfs/blob/master/docs/config.md)
    public func setConfig(_ config: [String: Any]? = nil) throws {
        try openRepoIfClosed()

        if let config = config {
            try repo!.setConfig(Config.configFromDict(config))
        } else {
            try repo!.setConfig(Config.defaultConfig())
        }
    }

    /// Gets the dict value associated to the key passed as parameter in the IPFS instance config
    /// - Parameter key: The key associated to the value to get in the IPFS config
    /// - Throws:
    ///     - `RepoError`: If the opening of the repo or the getting of its config failed
    ///     - `ConfigError`: If the getting of the config value as a dict failed
    /// - Returns: The dict value associated to the key passed as parameter in the IPFS instance config
    /// - seealso: [IPFS Config Doc](https://github.com/ipfs/go-ipfs/blob/master/docs/config.md)
    public func getConfigKey(_ key: String) throws -> [String: Any] {
        try openRepoIfClosed()

        return try repo!.getConfig().getKey(key)
    }

    /// Sets dict config value to the key passed as parameters in the IPFS instance config
    /// - Attention: A started IPFS instance must be restarted for its config to be applied
    /// - Parameters:
    ///     - key: The key associated to the value to set in the IPFS instance config
    ///     - value: The dict value associated to the key to set in the IPFS instance config
    /// - Throws:
    ///     - `RepoError`: If the opening of the repo or the getting/setting of its config failed
    ///     - `ConfigError`: If the setting of the config value as a dict failed
    /// - seealso: [IPFS Config Doc](https://github.com/ipfs/go-ipfs/blob/master/docs/config.md)
    public func setConfigKey(_ key: String, _ value: [String: Any]) throws {
        try openRepoIfClosed()

        let config = try repo!.getConfig()
        try config.setKey(key, value)
        try repo!.setConfig(config)
    }

    /// Creates and returns a RequestBuilder associated to this IPFS instance shell
    /// - Parameter command: The command of the request
    /// - Throws: `IPFSError`: If the request creaton failed
    /// - Returns: A RequestBuilder based on the command passed as parameter
    public func newRequest(_ command: String) throws -> RequestBuilder {
        guard let requestBuilder = self.shell?.newRequest(command) else {
            throw IPFSError("unable to get shell, is the node started?")
        }

        return RequestBuilder(requestBuilder)
    }

    /// Sets the primary and secondary DNS for gomobile (hacky, will be removed in future version)
    /// - Parameters:
    ///   - primary: The primary DNS address in the form `<ip4>:<port>`
    ///   - secondary: The secondary DNS address in the form `<ip4>:<port>`
    public class func setDNSPair(_ primary: String, _ secondary: String) {
        CoreSetDNSPair(primary, secondary, false)
    }

    /// Internal helper that opens the repo if it is closed
    /// - Throws: `RepoError`: If the opening of the repo failed
    private func openRepoIfClosed() throws {
        if self.repo == nil {
            self.repo = try Repo(self.absRepoURL)
        }
    }
}
