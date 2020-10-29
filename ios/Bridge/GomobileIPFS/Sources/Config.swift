//
//  Config.swift
//  Bridge
//
//  Created by Guilhem Fanton on 07/11/2019.
//

import Foundation
import Core

/// ConfigError is a Config specific error (subclass of IPFSError)
public class ConfigError: IPFSError {
    private static var code: Int = 2
    private static var subdomain: String = "Config"

    required init(_ description: String, _ optCause: NSError? = nil) {
        super.init(description, optCause, ConfigError.subdomain, ConfigError.code)
    }

    required init?(coder: NSCoder) {
        super.init(coder: coder)
    }
}

/// Config is a class that wraps a golang `config` object
///
/// **Should not be used on its own**
public class Config {
    let goConfig: CoreConfig

    /// Class constructor using a golang `config` object passed as parameter as config instance
    /// - Parameter config: A golang `config` object
    public init(_ config: CoreConfig) {
        self.goConfig = config
    }

    /// Returns a default golang `config` object
    /// - Throws: `ConfigError`: If the creation of the config failed
    /// - Returns: A golang `config` object
    public class func defaultConfig() throws -> Config {
        var err: NSError?

        if let config = CoreNewDefaultConfig(&err) {
            return Config(config)
        } else {
            throw ConfigError("default config creation failed", err)
        }
    }

    /// Returns an empty golang `config` object
    /// - Throws: `ConfigError`: If the creation of the config failed
    /// - Returns: A golang `config` object
    public class func emptyConfig() throws -> Config {
        var err: NSError?

        if let config = CoreNewConfig("{}".data(using: .utf8), &err) {
            return Config(config)
        } else {
            throw ConfigError("empty config creation failed", err)
        }
    }

    /// Returns a golang `config` object based on the dict passed as parameter
    /// - Parameter dict: The dict containing the config to create
    /// - Throws: `ConfigError`: If the creation of the config failed
    /// - Returns: A golang `config` object
    public class func configFromDict(_ dict: [String: Any]) throws -> Config {
        var err: NSError?

        let json = try JSONSerialization.data(withJSONObject: dict)

        if let config = CoreNewConfig(json, &err) {
            return Config(config)
        } else {
            throw ConfigError("config from dict creation failed", err)
        }
    }

    /// Gets the config as a dict
    /// - Throws: `ConfigError`: If the getting of the config failed
    /// - Returns: A dict containing the config
    public func get() throws -> [String: Any] {
        do {
            let rawJson = try self.goConfig.get()
            let json = try JSONSerialization.jsonObject(with: rawJson, options: [])
            return (json as? [String: Any])!
        } catch let error as NSError {
            throw ConfigError("getting failed", error)
        }
    }

    /// Sets a key and its value in the config
    /// - Parameters:
    ///   - key: The key to set
    ///   - value: A dict containing the value to set
    /// - Throws: `ConfigError`: If the setting of the key failed
    public func setKey(_ key: String, _ value: [String: Any]) throws {
        do {
            let json = try JSONSerialization.data(withJSONObject: value)
            try self.goConfig.setKey(key, raw_value: json)
        } catch let error as NSError {
            throw ConfigError("setting key failed", error)
        }
    }

    /// Gets the value associated to the key passed as parameter in the config
    /// - Parameter key: The key to get
    /// - Throws: `ConfigError`: If the getting of the key failed
    /// - Returns: A dict containing the value associated to the key passed as parameter
    public func getKey(_ key: String) throws -> [String: Any] {
        do {
            let rawJson = try self.goConfig.getKey(key)
            let json = try JSONSerialization.jsonObject(with: rawJson, options: [])
            return (json as? [String: Any])!
        } catch let error as NSError {
            throw ConfigError("getting key failed", error)
        }
    }
}
