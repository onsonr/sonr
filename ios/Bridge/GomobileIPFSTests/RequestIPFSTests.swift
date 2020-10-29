//
//  RequestIPFSTests.swift
//  GomobileIPFSTests
//
//  Created by Antoine Eddi on 4/15/20.
//  Copyright © 2020 Antoine Eddi. All rights reserved.
//

import XCTest
@testable import GomobileIPFS

class RequestIPFSTests: XCTestCase {

    private var ipfs: IPFS!

    override func setUp() {
        do {
            ipfs = try IPFS()
            try ipfs!.start()
        } catch _ {
            XCTFail("IPFS initialization failed")
        }
    }

    func testDNSRequest() throws {
        let domain = "website.ipfs.io"

        let resolveResp = try ipfs.newRequest("resolve")
            .with(argument: "/ipns/\(domain)")
            .sendToDict()
        let dnsResp = try ipfs.newRequest("dns")
            .with(argument: domain)
            .sendToDict()

        guard let resolvePath = resolveResp["Path"] as? String else {
            XCTFail("error while casting value associated to \"Path\" key")
            return
        }
        guard let dnsPath = dnsResp["Path"] as? String else {
            XCTFail("error while casting value associated to \"Path\" key")
            return
        }
        let index = dnsPath.index(dnsPath.startIndex, offsetBy: 6)

        XCTAssertEqual(
            resolvePath,
            dnsPath,
            "resolve and dns request should return the same result"
        )

        XCTAssertEqual(
            dnsPath[..<index],
            "/ipfs/",
            "response should start with \"/ipfs/\""
        )
    }

    func testCatFile() throws {
        let latestRaw = try ipfs.newRequest("cat")
            .with(argument: "/ipns/xkcd.hacdias.com/latest/info.json")
            .send()

        do {
            try JSONSerialization.jsonObject(with: latestRaw, options: [])
        } catch _ {
            XCTFail("error while parsing fetched JSON:  \(String(decoding: latestRaw, as: UTF8.self))")
        }
    }
}
