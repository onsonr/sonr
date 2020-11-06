package io.sonr.sonr_core;

import sonr.*;

public class SonrProxyCallback implements Callback {
    public void onMessage(String data) {
        System.out.println("Refreshed: " + data);
    }

    public void onRefresh(String data) {
        System.out.println("Refreshed: " + data);
    }

    public void onRequested(String data) {
        // System.out.println("Refreshed: ", s!)
    }

    public void onAccepted(String data) {
        // System.out.println("Refreshed: ", s!)
    }

    public void onDenied(String data) {
        // System.out.println("Refreshed: ", s!)
    }

    public void onProgress(String data) {
        // System.out.println("Refreshed: ", s!)
    }

    public void onComplete(String data) {
        // System.out.println("Refreshed: ", s!)
    }
}