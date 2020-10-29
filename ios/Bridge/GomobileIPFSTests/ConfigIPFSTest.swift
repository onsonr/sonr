//
//  ConfigIPFSTest.swift
//  GomobileIPFSTests
//
//  Created by Antoine Eddi on 4/16/20.
//  Copyright © 2020 Antoine Eddi. All rights reserved.
//

import XCTest
@testable import GomobileIPFS

class ConfigIPFSTests: XCTestCase {

    private var ipfs: IPFS!

    override func setUp() {
        do {
            ipfs = try IPFS()
        } catch _ {
            XCTFail("IPFS initialization failed")
        }
    }

    func testConfig() throws {
        // Reset to default config
        try ipfs.setConfig(nil)

        let backup = try ipfs.getConfig()

        XCTAssertTrue(
            try getMDNSStateFromConfig(),
            "MDNS state should be enabled on default config"
        )
        try setMDNSStateToConfig(false)
        XCTAssertFalse(
            try getMDNSStateFromConfig(),
            "MDNS state should be disabled after setting it in config"
        )
        XCTAssertFalse(
            try (ipfs.getConfig() as NSDictionary).isEqual(backup),
            "current IPFS config and backup should not be equals"
        )

        try ipfs.start()

        XCTAssertFalse(
            try getMDNSStateFromConfig(),
            "MDNS state should be still disabled after starting IPFS"
        )
        try setMDNSStateToConfig(true)
        XCTAssertTrue(
            try getMDNSStateFromConfig(),
            "MDNS state should be enabled after setting it in config"
        )

        try ipfs.restart()

        XCTAssertTrue(
            try getMDNSStateFromConfig(),
            "MDNS state should be still enabled after restarting IPFS"
        )
        try setMDNSStateToConfig(false)
        XCTAssertFalse(
            try getMDNSStateFromConfig(),
            "MDNS state should be disabled after setting it in config"
        )

        try ipfs.stop()

        XCTAssertFalse(
            try getMDNSStateFromConfig(),
            "MDNS state should be still disabled after stopping IPFS"
        )
        try setMDNSStateToConfig(true)
        XCTAssertTrue(
            try getMDNSStateFromConfig(),
            "MDNS state should be enabled after setting it in config"
        )

        let mdnsCfg = ["MDNS": ["Enabled": true, "Interval": 10]]
        try ipfs.setConfigKey("Discovery", mdnsCfg);

        XCTAssertTrue(
            try (ipfs.getConfig() as NSDictionary).isEqual(backup),
            "current IPFS config and backup should be equals"
        )

        // Reset config
        try ipfs.setConfig(nil)

        XCTAssertFalse(
            try (ipfs.getConfig() as NSDictionary).isEqual(backup),
            "current IPFS config and backup should not be equals (Identity changed)"
        )
    }

    func getMDNSStateFromConfig() throws -> Bool {
        let config = try ipfs.getConfig()

        guard let discoveryCfg = config["Discovery"] as? [String: Any] else {
            XCTFail("error while casting value associated to \"Discovery\" key")
            return false
        }
        guard let mdnsCfg = discoveryCfg["MDNS"] as? [String: Any] else {
            XCTFail("error while casting value associated to \"MDNS\" key")
            return false
        }
        guard let state = mdnsCfg["Enabled"] as? Bool else {
            XCTFail("error while casting value associated to \"Enabled\" key")
            return false
        }

        return state
    }

    func setMDNSStateToConfig(_ state: Bool) throws {
        let mdnsCfg = ["MDNS": ["Enabled": state]]
        try ipfs.setConfigKey("Discovery", mdnsCfg);
    }
}
