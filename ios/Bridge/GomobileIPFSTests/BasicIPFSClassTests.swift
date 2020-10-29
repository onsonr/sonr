//
//  BasicIPFSClassTests.swift
//  GomobileIPFSTests
//
//  Created by Antoine Eddi on 2/23/20.
//  Copyright © 2020 Antoine Eddi. All rights reserved.
//

import XCTest
@testable import GomobileIPFS

class BasicIPFSClassTests: XCTestCase {

    private var ipfs: IPFS!

    override func setUp() {
        do {
            ipfs = try IPFS()
            try ipfs!.start()
        } catch _ {
            XCTFail("IPFS initialization failed")
        }
    }

    func testDefault() throws {
        // Tests on started IPFS instance
        try makeRequest()

        XCTAssertTrue(
            ipfs.isStarted(),
            "IPFS should be started"
        )

        XCTAssertTrue(
            FileManager.default.fileExists(
                atPath: ipfs.getAbsoluteRepoPath().appendingPathComponent("/config").path
            ),
            "config file doesn't exist in repo"
        )
        XCTAssertTrue(
            FileManager.default.fileExists(
                atPath: ipfs.getAbsoluteRepoPath().appendingPathComponent("/version").path
            ),
            "version file doesn't exist in repo"
        )
        XCTAssertTrue(
            FileManager.default.fileExists(
                atPath: ipfs.getAbsoluteRepoPath().appendingPathComponent("/repo.lock").path
            ),
            "repo.lock file doesn't exist in repo"
        )

        do {
            try ipfs.start()
            XCTFail("Calling start() on a started IPFS instance should throw")
        } catch _ {}


        // Tests on stopped IPFS instance
        try ipfs.stop()

        XCTAssertFalse(
            ipfs.isStarted(),
            "IPFS should be stopped"
        )

        do {
            try ipfs.stop()
            XCTFail("Calling stop() on a stopped IPFS instance should throw")
        } catch _ {}

        do {
            try ipfs.restart()
            XCTFail("Calling restart() on a stopped IPFS instance should throw")
        } catch _ {}

        do {
            try makeRequest()
            XCTFail("Making request on a stopped IPFS instance should throw")
        } catch _ {}


        // Tests on started IPFS instance (after stop)
        try ipfs.start()
        try makeRequest()

        XCTAssertTrue(
            ipfs.isStarted(),
            "IPFS should be started"
        )


        // Tests on restarted IPFS instance
        try ipfs.restart()
        try makeRequest()

        XCTAssertTrue(
            ipfs.isStarted(),
            "IPFS should be started"
        )
    }

    func makeRequest() throws {
        let res = try ipfs.newRequest("id").sendToDict()

        // TODO: improve these checks
        guard let peerID = res["ID"] as? String else {
            XCTFail("error while casting value associated to \"ID\" key")
            return
        }
        guard let publicKey = res["PublicKey"] as? String else {
            XCTFail("error while casting value associated to \"PublicKey\" key")
            return
        }
        let index = peerID.index(peerID.startIndex, offsetBy: 2)

        XCTAssertEqual(
            peerID[..<index],
            "Qm",
            "Invalid peerID"
        )
        XCTAssertNotEqual(
            publicKey,
            "",
            "Empty public key"
        )
    }
}
