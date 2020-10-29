//
//  PeerCountUpdater.swift
//  Example
//
//  Created by CI Agent on 24/11/2019.
//  Copyright © 2019 CocoaPods. All rights reserved.
//

import Foundation
import GomobileIPFS

public class PeerCountUpdater: NSObject {
    let interval: Double
    private var running: Bool = false

    public init(_ interval: Double = 1.0) {
        self.interval = interval
        super.init()
    }

    public func start() {
        if self.running != true {
            self.running = true
            self.updatePeerCount(0.0)
        }
    }

    private func updatePeerCount(_ delay: Double) {
        DispatchQueue.global(qos: .background).asyncAfter(deadline: .now() + delay, execute: {
            if self.running && ViewController.ipfs!.isStarted() {
                var peerCount: Int = 0
                do {
                    let res = try ViewController.ipfs!.newRequest("/swarm/peers").sendToDict()
                    let peerList = res["Peers"] as? NSArray
                    peerCount = peerList?.count ?? 0
                } catch let error {
                    print(error)
                }
                DispatchQueue.main.async {
                    NotificationCenter.default.post(
                        name: Notification.Name("updatePeerCount"),
                        object: nil,
                        userInfo: ["peerCount": peerCount]
                    )
                    self.updatePeerCount(self.interval)
                }
            }
        })
    }

    public func stop() {
        self.running = false
    }
}
