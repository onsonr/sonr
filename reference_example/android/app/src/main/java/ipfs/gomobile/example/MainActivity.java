package ipfs.gomobile.example;

import androidx.appcompat.app.AppCompatActivity;

import android.graphics.Color;
import android.os.Bundle;
import android.view.View;
import android.widget.Button;
import android.widget.ProgressBar;
import android.widget.TextView;

import ipfs.gomobile.android.IPFS;

public class MainActivity extends AppCompatActivity {
    private IPFS ipfs;

    private TextView ipfsTitle;
    private ProgressBar ipfsProgress;
    private TextView ipfsResult;

    private TextView peerCounter;

    private Button xkcdButton;
    private TextView xkcdStatus;
    private ProgressBar xkcdProgress;
    private TextView xkcdError;

    private PeerCounter peerCounterUpdater;

    void setIpfs(IPFS ipfs) {
        this.ipfs = ipfs;
    }

    IPFS getIpfs() {
        return ipfs;
    }

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_main);

        ipfsTitle = findViewById(R.id.ipfsTitle);
        ipfsProgress = findViewById(R.id.ipfsProgress);
        ipfsResult = findViewById(R.id.ipfsResult);

        peerCounter = findViewById(R.id.peerCounter);

        xkcdButton = findViewById(R.id.xkcdButton);
        xkcdStatus = findViewById(R.id.xkcdStatus);
        xkcdProgress = findViewById(R.id.xkcdProgress);
        xkcdError = findViewById(R.id.xkcdError);

        new StartIPFS(this).execute();

        final MainActivity activity = this;
        xkcdButton.setOnClickListener(new View.OnClickListener() {
            @Override
            public void onClick(View v) {
                new FetchRandomXKCD(activity).execute();
            }
        });
    }

    @Override
    protected void onPause() {
        super.onPause();

        if (peerCounterUpdater != null) {
            peerCounterUpdater.stop();
        }
    }

    @Override
    protected void onResume() {
        super.onResume();

        if (peerCounterUpdater != null) {
            peerCounterUpdater.start();
        }
    }

    void displayPeerIDError(String error) {
        ipfsTitle.setTextColor(Color.RED);
        ipfsResult.setTextColor(Color.RED);

        ipfsTitle.setText(getString(R.string.titlePeerIDErr));
        ipfsResult.setText(error);
        ipfsProgress.setVisibility(View.INVISIBLE);
    }

    void displayPeerIDResult(String peerID) {
        ipfsTitle.setText(getString(R.string.titlePeerID));
        ipfsResult.setText(peerID);
        ipfsProgress.setVisibility(View.INVISIBLE);

        updatePeerCount(0);
        peerCounter.setVisibility(View.VISIBLE);
        xkcdButton.setVisibility(View.VISIBLE);

        peerCounterUpdater = new PeerCounter(this, 1000);
        peerCounterUpdater.start();
    }

    void updatePeerCount(int count) {
        peerCounter.setText(getString(R.string.titlePeerCon, count));
    }

    void displayFetchProgress() {
        xkcdStatus.setTextColor(ipfsTitle.getCurrentTextColor());
        xkcdStatus.setText(R.string.titleFetching);
        xkcdStatus.setVisibility(View.VISIBLE);
        xkcdError.setVisibility(View.INVISIBLE);
        xkcdProgress.setVisibility(View.VISIBLE);

        xkcdButton.setAlpha(0.5f);
        xkcdButton.setClickable(false);
    }

    void displayFetchSuccess() {
        xkcdStatus.setVisibility(View.INVISIBLE);
        xkcdProgress.setVisibility(View.INVISIBLE);

        xkcdButton.setAlpha(1);
        xkcdButton.setClickable(true);
    }

    void displayFetchError(String error) {
        xkcdStatus.setTextColor(Color.RED);
        xkcdStatus.setText(R.string.titleFetchingErr);

        xkcdProgress.setVisibility(View.INVISIBLE);
        xkcdError.setVisibility(View.VISIBLE);
        xkcdError.setText(error);

        xkcdButton.setAlpha(1);
        xkcdButton.setClickable(true);
    }

    static String exceptionToString(Exception error) {
        String string = error.getMessage();

        if (error.getCause() != null) {
            string += ": " + error.getCause().getMessage();
        }

        return string;
    }
}
