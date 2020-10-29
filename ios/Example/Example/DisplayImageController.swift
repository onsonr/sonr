//
//  DisplayImageController.swift
//  Example
//
//  Created by CI Agent on 24/11/2019.
//  Copyright © 2019 CocoaPods. All rights reserved.
//

import UIKit

class DisplayImageController: UIViewController {
    @IBOutlet weak var XKCDTitle: UINavigationItem!
    @IBOutlet weak var XKCDImage: UIImageView!

    var xkcdTitle: String = ""
    var xkcdImage: UIImage = UIImage()


    override func viewDidLoad() {
        super.viewDidLoad()

        XKCDTitle.title = xkcdTitle
        XKCDImage.image = xkcdImage
    }

    override var supportedInterfaceOrientations: UIInterfaceOrientationMask {
        get {
            return .all
        }
    }
}
