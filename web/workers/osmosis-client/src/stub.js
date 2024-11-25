import { DurableObject } from "cloudflare:workers";

/** A Durable Object's behavior is defined in an exported Javascript class */
export class OsmosisDurableClient extends DurableObject {
  /**
   * The constructor is invoked once upon creation of the Durable Object, i.e. the first call to
   * 	`DurableObjectStub::get` for a given identifier (no-op constructors can be omitted)
   *
   * @param {DurableObjectState} ctx - The interface for interacting with Durable Object state
   * @param {Env} env - The interface to reference bindings declared in wrangler.toml
   */
  constructor(ctx, env) {
    super(ctx, env);
  }

  /**
   * The Durable Object exposes an RPC method sayHello which will be invoked when when a Durable
   *  Object instance receives a request from a Worker via the same method invocation on the stub
   *
   * @param {string} name - The name provided to a Durable Object instance from a Worker
   * @returns {Promise<string>} The greeting to be sent back to the Worker
   */
  async sayHello(name) {
    return `Hello, ${name}!`;
  }
}
