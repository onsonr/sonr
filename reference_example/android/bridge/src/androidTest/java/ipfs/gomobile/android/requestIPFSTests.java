package ipfs.gomobile.android;

import android.content.Context;

import androidx.test.ext.junit.runners.AndroidJUnit4;
import androidx.test.platform.app.InstrumentationRegistry;

import org.json.JSONException;
import org.json.JSONObject;
import org.junit.Before;
import org.junit.Rule;
import org.junit.Test;
import org.junit.rules.Timeout;
import org.junit.runner.RunWith;

import static org.junit.Assert.*;

/**
* Instrumented test, which will execute on an Android device.
*
* @see <a href="http://d.android.com/tools/testing">Testing documentation</a>
*/
@RunWith(AndroidJUnit4.class)
public class requestIPFSTests {
    private IPFS ipfs;

    @Rule
    public Timeout globalTimeout = Timeout.seconds(600);

    @Before
    public void setup() throws Exception {
        Context appContext = InstrumentationRegistry.getInstrumentation().getTargetContext();
        ipfs = new IPFS(appContext);
        ipfs.start();
    }

    @Test
    public void testDNSRequest() throws Exception {
        String domain = "website.ipfs.io";

        JSONObject resolveResp = ipfs.newRequest("resolve")
                .withArgument("/ipns/" + domain)
                .sendToJSONList()
                .get(0);
        JSONObject dnsResp = ipfs.newRequest("dns")
                .withArgument(domain)
                .sendToJSONList()
                .get(0);

        String resolvePath = resolveResp.getString("Path");
        String dnsPath = dnsResp.getString("Path");

        assertEquals(
            "resolve and dns request should return the same result",
            resolvePath,
            dnsPath
        );

        assertEquals(
            "response should start with \"/ipfs/\"",
            dnsPath.substring(0, 6),
            "/ipfs/"
        );
    }

    @Test
    public void testCatFile() throws Exception {
        byte[] latestRaw = ipfs.newRequest("cat")
                .withArgument("/ipns/xkcd.hacdias.com/latest/info.json")
                .send();

        try {
            new JSONObject(new String(latestRaw));
        } catch (JSONException e) {
            fail("error while parsing fetched JSON: " + new String(latestRaw));
        }
    }
}
