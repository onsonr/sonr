//
//  RequestBuilder.swift
//  Bridge
//
//  Created by Guilhem Fanton on 14/01/2020.
//

import Foundation
import Core

/// Enum of the different option types: bool, string and bytes
public enum RequestOption {
    case bool(Bool)
    case string(String)
    case bytes(Data)
}

/// Enum of the different body types: string and bytes
public enum RequestBody {
    case string(String)
    case bytes(Data)
}

/// RequestBuilderError is a RequestBuilder specific error (subclass of IPFSError)
public class RequestBuilderError: IPFSError {
    private static var code: Int = 6
    private static var subdomain: String = "RequestBuilder"

    required init(_ description: String, _ optCause: NSError? = nil) {
        super.init(description, optCause, RequestBuilderError.subdomain, RequestBuilderError.code)
    }

    required init?(coder: NSCoder) {
        super.init(coder: coder)
    }
}

/// RequestBuilder is an IPFS command request builder
public class RequestBuilder {
    private let requestBuilder: CoreRequestBuilder

    internal init(_ requestBuilder: CoreRequestBuilder) {
        self.requestBuilder = requestBuilder
    }

    /// Adds an argument to the request
    /// - Parameter argument: The argument to add
    /// - Returns: This instance of RequestBuilder
    /// - seealso: [IPFS API Doc](https://docs.ipfs.io/reference/api/http/)
    public func with(argument: String) -> RequestBuilder {
        self.requestBuilder.argument(argument)
        return self
    }

    /// Adds an option to the request
    /// - Parameters:
    ///   - option: The name of the option to add
    ///   - value: The value of the option to add
    /// - Returns: This instance of RequestBuilder
    /// - seealso: [IPFS API Doc](https://docs.ipfs.io/reference/api/http/)
    public func with(option: String, value: RequestOption) -> RequestBuilder {
        switch value {
        case .bool(let bool):
            self.requestBuilder.boolOptions(option, value: bool)
        case .string(let string):
            self.requestBuilder.stringOptions(option, value: string)
        case .bytes(let data):
            self.requestBuilder.byteOptions(option, value: data)
        }

        return self
    }

    /// Adds a body to the request
    /// - Parameter body: The value of body to add
    /// - Returns: This instance of RequestBuilder
    /// - seealso: [IPFS API Doc](https://docs.ipfs.io/reference/api/http/)
    public func with(body: RequestBody) -> RequestBuilder {
        switch body {
        case .bytes(let data):
            self.requestBuilder.bodyBytes(data)
        case .string(let string):
            self.requestBuilder.bodyString(string)
        }

        return self
    }

    /// Adds a header to the request
    /// - Parameters:
    ///   - header: The key of the header to add
    ///   - value: The value of the header to add
    /// - Returns: This instance of RequestBuilder
    /// - seealso: [IPFS API Doc](https://docs.ipfs.io/reference/api/http/)
    public func with(header: String, value: String) -> RequestBuilder {
        self.requestBuilder.header(header, value: value)
        return self
    }

    /// Sends the request to the underlying go-ipfs node
    /// - Throws: `RequestBuilderError`: If sending the request failed
    /// - Returns: A Data object containing the response
    /// - seealso: [IPFS API Doc](https://docs.ipfs.io/reference/api/http/)
    public func send() throws -> Data {
        do {
            return try self.requestBuilder.send()
        } catch let error as NSError {
            throw RequestBuilderError("sending request failed", error)
        }
    }

    /// Sends the request to the underlying go-ipfs node and returns a dict
    /// - Throws: `RequestBuilderError`: If sending the request or converting the response failed
    /// - Returns: A dict containing the response
    /// - seealso: [IPFS API Doc](https://docs.ipfs.io/reference/api/http/)
    public func sendToDict() throws -> [String: Any] {
        let res = try self.requestBuilder.send()

        do {
            let json = try JSONSerialization.jsonObject(with: res, options: [])
            return (json as? [String: Any])!
        } catch let error as NSError {
            throw RequestBuilderError("converting response to dict failed", error)
        }
    }
}
