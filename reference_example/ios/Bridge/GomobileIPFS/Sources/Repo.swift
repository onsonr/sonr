//
//  Repo.swift
//  Bridge
//
//  Created by Guilhem Fanton on 07/11/2019.
//

import Foundation
import Core

/// RepoError is a Repo specific error (subclass of IPFSError)
public class RepoError: IPFSError {
    private static var code: Int = 4
    private static var subdomain: String = "Repo"

    required init(_ description: String, _ optCause: NSError? = nil) {
        super.init(description, optCause, RepoError.subdomain, RepoError.code)
    }

    required init?(coder: NSCoder) {
        super.init(coder: coder)
    }
}

/// Repo is a class that wraps a golang `repo` object
///
/// **Should not be used on its own**
public class Repo {
    let goRepo: CoreRepo

    private let url: URL

    /// Class constructor using url passed as parameter as repo path
    /// - Parameter url: The path of the repo
    /// - Throws: `RepoError`: If the opening of the repo failed
    public init(_ url: URL) throws {
        var err: NSError?

        if let repo = CoreOpenRepo(url.path, &err) {
            self.url = url
            self.goRepo = repo
        } else {
            throw RepoError("openning failed", err)
        }
    }

    /// Returns True if the repo is initialized
    /// - Parameter url: The path of the repo
    /// - Throws: `RepoError`: If the checking failed
    /// - Returns: True, if the repo is initialized
    public static func isInitialized(_ url: URL) -> Bool {
        return CoreRepoIsInitialized(url.path)
    }

    /// Initializes the repo using the path and the config passed as parameters
    /// - Parameters:
    ///     - url: The path of the repo
    ///     - config: The config of the repo
    /// - Throws: `RepoError`: If the initialization of the repo failed
    public static func initialize(_ url: URL, _ config: Config) throws {
        var err: NSError?
        var isDirectory: ObjCBool = true
        let exist = FileManager.default.fileExists(atPath: url.path, isDirectory: &isDirectory)
        if !exist {
            try FileManager.default.createDirectory(
                atPath: url.path,
                withIntermediateDirectories: true,
                attributes: nil
            )
        }

        CoreInitRepo(url.path, config.goConfig, &err)
        if err != nil {
            throw RepoError("initialization failed", err)
        }
    }

    /// Gets the config of the repo
    /// - Throws: `RepoError`: If the getting of the config failed
    public func getConfig() throws -> Config {
        do {
            let goconfig = try self.goRepo.getConfig()
            return Config(goconfig)
        } catch let error as NSError {
            throw RepoError("getting configuration failed", error)
        }
    }

    /// Sets the config of the repo
    /// - Parameter config: The config to set
    /// - Throws: `RepoError`: If the setting of the config failed
    public func setConfig(_ config: Config) throws {
        do {
            try self.goRepo.setConfig(config.goConfig)
        } catch let error as NSError {
            throw RepoError("setting configuration failed", error)
        }
    }
}
