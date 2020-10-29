package ipfs.gomobile.example;

import android.os.AsyncTask;
import android.util.Log;

import org.json.JSONObject;

import java.lang.ref.WeakReference;
import java.util.ArrayList;

import ipfs.gomobile.android.IPFS;

final class StartIPFS extends AsyncTask<Void, Void, String> {
    private static final String TAG = "StartIPFS";

    private WeakReference<MainActivity> activityRef;
    private boolean backgroundError;

    StartIPFS(MainActivity activity) {
        activityRef = new WeakReference<>(activity);
    }

    @Override
    protected void onPreExecute() {}

    @Override
    protected String doInBackground(Void... v) {
        MainActivity activity = activityRef.get();
        if (activity == null || activity.isFinishing()) {
            cancel(true);
            return null;
        }

        try {
            IPFS ipfs = new IPFS(activity.getApplicationContext());
            ipfs.start();

            ArrayList<JSONObject> jsonList = ipfs.newRequest("id").sendToJSONList();

            activity.setIpfs(ipfs);
            return jsonList.get(0).getString("ID");
        } catch (Exception err) {
            backgroundError = true;
            return MainActivity.exceptionToString(err);
        }
    }

    protected void onPostExecute(String result) {
        MainActivity activity = activityRef.get();
        if (activity == null || activity.isFinishing()) return;

        if (backgroundError) {
            activity.displayPeerIDError(result);
            Log.e(TAG, "IPFS start error: " + result);
        } else {
            activity.displayPeerIDResult(result);
            Log.i(TAG, "Your PeerID is: " + result);
        }
    }
}
