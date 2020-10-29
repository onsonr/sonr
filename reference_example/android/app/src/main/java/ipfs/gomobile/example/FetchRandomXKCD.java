package ipfs.gomobile.example;

import android.content.Intent;
import android.os.AsyncTask;
import android.util.Log;

import org.json.JSONObject;

import java.lang.ref.WeakReference;
import java.util.Random;

import ipfs.gomobile.android.IPFS;

final class FetchRandomXKCD extends AsyncTask<Void, Void, String> {
    private static final String TAG = "FetchRandomXKCD";

    private static final String XKCDIPNS = "/ipns/xkcd.hacdias.com";

    private static Random random = new Random();

    private static int XKCDLatest = -1;

    private final WeakReference<MainActivity> activityRef;
    private boolean backgroundError;
    private byte[] fetchedData;

    FetchRandomXKCD(MainActivity activity) {
        activityRef = new WeakReference<>(activity);
    }

    @Override
    protected void onPreExecute() {
        MainActivity activity = activityRef.get();
        if (activity == null || activity.isFinishing()) return;

        activity.displayFetchProgress();
    }

    @Override
    protected String doInBackground(Void... v) {
        MainActivity activity = activityRef.get();
        if (activity == null || activity.isFinishing()) {
            cancel(true);
            return null;
        }

        IPFS ipfs = activity.getIpfs();

        try {
            if (XKCDLatest == -1) {
                String address = String.format("%s/latest/info.json", XKCDIPNS);
                byte[] latestRaw = ipfs.newRequest("cat")
                    .withArgument(address)
                    .send();

                XKCDLatest = new JSONObject(new String(latestRaw)).getInt("num");
            }

            int randomIndex = random.nextInt(XKCDLatest) + 1;
            String formattedIndex = String.format("%04d", randomIndex);

            byte[] infoRaw = ipfs.newRequest("cat")
                .withArgument(String.format("%s/%s/info.json", XKCDIPNS, formattedIndex))
                .send();
            JSONObject infoJSON = new JSONObject(new String(infoRaw));

            String title = infoJSON.getString("title");

            String imgURL = infoJSON.getString("img");
            String[] imgURLSplit = imgURL.split("\\.");
            String imgExt = imgURLSplit[imgURLSplit.length - 1].contains("png") ? "png" : "jpg";

            fetchedData = ipfs.newRequest("cat")
                .withArgument(String.format("%s/%s/image.%s", XKCDIPNS, formattedIndex, imgExt))
                .send();

            return String.format("%d. %s", randomIndex, title);
        } catch (Exception err) {
            backgroundError = true;
            return MainActivity.exceptionToString(err);
        }
    }

    protected void onPostExecute(String result) {
        MainActivity activity = activityRef.get();
        if (activity == null || activity.isFinishing()) return;

        if (backgroundError) {
            activity.displayFetchError(result);
            Log.e(TAG, "XKCD fetch error: " + result);
        } else {
            activity.displayFetchSuccess();

            Intent intent = new Intent(activity, DisplayImageActivity.class);
            intent.putExtra("ImageData", fetchedData);
            intent.putExtra("Title", result);
            activity.startActivity(intent);
        }
    }
}
