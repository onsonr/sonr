package io.sonr.sonr_core;

import sonr.*;

public class SonrProxy {
    // ** Dependencies ** //
    public SonrProxyCallback callback = new SonrProxyCallback();
    private Node node = new Node();

    // ** Functional Methods ** //
    public void initialize(Node proxyNode){
        if(proxyNode != null){
            this.node = proxyNode;
        }
   }

   // ** Get Properties ** //
   // Return Node ID
  public String getId() {
       // Check Proxy
       if (this.node != null) {
           return node.getPeerID();
       }
       return "Invalid";
   }

   // Return User Info
   public String getUser() {
        // Check Proxy
       if (this.node != null) {
           return node.getUser();
       }
       return "Invalid";
   }

   // ** Event Emission ** //
   // Send General Message
  public boolean send(String msg) {
       if (this.node != null) {
         boolean result = this.node.send(msg);
         return result;
       }
       return false;
   }

   // Update Profile Values
  public boolean update(String msg) {
       if (this.node != null) {
         boolean result = this.node.update(msg);
         return result;
       }
       return false;
   }
}