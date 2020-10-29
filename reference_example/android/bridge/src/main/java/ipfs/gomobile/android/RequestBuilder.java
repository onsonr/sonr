package ipfs.gomobile.android;

import androidx.annotation.NonNull;

import org.json.JSONException;
import org.json.JSONObject;

import java.util.ArrayList;
import java.util.Objects;
import java.util.Scanner;

/**
* RequestBuilder is an IPFS command request builder.
*/
public class RequestBuilder {

    private core.RequestBuilder requestBuilder;

    /**
    * Package-Private class constructor using RequestBuilder passed by IPFS.newRequest method.
    * @param requestBuilder A go-ipfs requestBuilder object
    */
    RequestBuilder(@NonNull core.RequestBuilder requestBuilder) {
        Objects.requireNonNull(requestBuilder, "requestBuilder should not be null");

        this.requestBuilder = requestBuilder;
    }

    // Send methods
    /**
    * Sends the request to the underlying go-ipfs node.
    *
    * @return A byte array containing the response
    * @throws RequestBuilderException If sending the request failed
    * @see <a href="https://docs.ipfs.io/reference/api/http/">IPFS API Doc</a>
    */
    public byte[] send() throws RequestBuilderException {
        try {
            return requestBuilder.send();
        } catch (Exception err) {
            throw new RequestBuilderException("Failed to send request", err);
        }
    }
    /**
    * Sends the request to the underlying go-ipfs node and returns an array of JSONObject.
    *
    * @return An ArrayList of JSONObject generated from the response
    * @throws RequestBuilderException If sending the request failed
    * @throws JSONException If converting the response to JSONObject failed
    * @see <a href="https://docs.ipfs.io/reference/api/http/">IPFS API Doc</a>
    */
    public ArrayList<JSONObject> sendToJSONList() throws RequestBuilderException, JSONException {
        String raw = new String(this.send());

        ArrayList<JSONObject> jsonList = new ArrayList<>();
        Scanner scanner = new Scanner(raw);
        while (scanner.hasNextLine()) {
            jsonList.add(new JSONObject(scanner.nextLine()));
        }

        return jsonList;
    }

    // Argument method
    /**
    * Adds an argument to the request.
    *
    * @param argument The argument to add
    * @return This instance of RequestBuilder
    * @see <a href="https://docs.ipfs.io/reference/api/http/">IPFS API Doc</a>
    */
    public RequestBuilder withArgument(@NonNull String argument) {
        Objects.requireNonNull(argument, "argument should not be null");

        requestBuilder.argument(argument);
        return this;
    }

    // Option methods
    /**
    * Adds a boolean option to the request.
    *
    * @param option The name of the option to add
    * @param value The boolean value of the option to add
    * @return This instance of RequestBuilder
    * @see <a href="https://docs.ipfs.io/reference/api/http/">IPFS API Doc</a>
    */
    public RequestBuilder withOption(@NonNull String option, boolean value) {
        Objects.requireNonNull(option, "option should not be null");

        requestBuilder.boolOptions(option, value);
        return this;
    }
    /**
    * Adds a string option to the request.
    *
    * @param option The name of the option to add
    * @param value The string value of the option to add
    * @return This instance of RequestBuilder
    * @see <a href="https://docs.ipfs.io/reference/api/http/">IPFS API Doc</a>
    */
    public RequestBuilder withOption(@NonNull String option, @NonNull String value) {
        Objects.requireNonNull(option, "option should not be null");
        Objects.requireNonNull(value, "value should not be null");

        requestBuilder.stringOptions(option, value);
        return this;
    }
    /**
    * Adds a byte array option to the request.
    *
    * @param option The name of the option to add
    * @param value The byte array value of the option to add
    * @return This instance of RequestBuilder
    * @see <a href="https://docs.ipfs.io/reference/api/http/">IPFS API Doc</a>
    */
    public RequestBuilder withOption(@NonNull String option, @NonNull byte[] value) {
        Objects.requireNonNull(option, "option should not be null");
        Objects.requireNonNull(value, "value should not be null");

        requestBuilder.byteOptions(option, value);
        return this;
    }

    // Body methods
    /**
    * Adds a string body to the request.
    *
    * @param body The string value of the body to add
    * @return This instance of RequestBuilder
    * @see <a href="https://docs.ipfs.io/reference/api/http/">IPFS API Doc</a>
    */
    public RequestBuilder withBody(@NonNull String body) {
        Objects.requireNonNull(body, "body should not be null");

        requestBuilder.bodyString(body);
        return this;
    }
    /**
    * Adds a byte array body to the request.
    *
    * @param body The byte array value of the body to add
    * @return This instance of RequestBuilder
    * @see <a href="https://docs.ipfs.io/reference/api/http/">IPFS API Doc</a>
    */
    public RequestBuilder withBody(@NonNull byte[] body) {
        Objects.requireNonNull(body, "body should not be null");

        requestBuilder.bodyBytes(body);
        return this;
    }

    // Header method
    /**
    * Adds a header to the request.
    *
    * @param key The key of the header to add
    * @param value The value of the header to add
    * @return This instance of RequestBuilder
    * @see <a href="https://docs.ipfs.io/reference/api/http/">IPFS API Doc</a>
    */
    public RequestBuilder withHeader(@NonNull String key, @NonNull String value) {
        Objects.requireNonNull(key, "key should not be null");
        Objects.requireNonNull(value, "value should not be null");

        requestBuilder.header(key, value);
        return this;
    }

    public static class RequestBuilderException extends Exception {
        RequestBuilderException(String message, Throwable err) { super(message, err); }
    }
}
