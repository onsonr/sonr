package ipfs.gomobile.android;

import android.content.Context;

import androidx.test.ext.junit.runners.AndroidJUnit4;
import androidx.test.platform.app.InstrumentationRegistry;

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
public class configIPFSTests {
    private IPFS ipfs;

    @Rule
    public Timeout globalTimeout = Timeout.seconds(240);

    @Before
    public void setup() throws Exception {
        Context appContext = InstrumentationRegistry.getInstrumentation().getTargetContext();
        ipfs = new IPFS(appContext);
    }

    @Test
    public void testConfig() throws Exception {
        // Reset to default config
        ipfs.setConfig(null);

        // Backup current config
        JSONObject backup = ipfs.getConfig();

        assertTrue(
            "MDNS state should be enabled on default config",
            getMDNSStateFromConfig()
        );
        setMDNSStateToConfig(false);
        assertFalse(
            "MDNS state should be disabled after setting it in config",
            getMDNSStateFromConfig()
        );
        assertFalse(
            "current IPFS config and backup should not be equals",
            ipfs.getConfig().toString().equals(backup.toString())
        );

        ipfs.start();

        assertFalse(
                "MDNS state should be still disabled after starting IPFS",
                getMDNSStateFromConfig()
        );
        setMDNSStateToConfig(true);
        assertTrue(
                "MDNS state should be enabled after setting it in config",
                getMDNSStateFromConfig()
        );

        ipfs.restart();

        assertTrue(
                "MDNS state should be still enabled after restarting IPFS",
                getMDNSStateFromConfig()
        );
        setMDNSStateToConfig(false);
        assertFalse(
                "MDNS state should be disabled after setting it in config",
                getMDNSStateFromConfig()
        );

        ipfs.stop();

        assertFalse(
                "MDNS state should be still disabled after stopping IPFS",
                getMDNSStateFromConfig()
        );
        setMDNSStateToConfig(true);
        assertTrue(
                "MDNS state should be enabled after setting it in config",
                getMDNSStateFromConfig()
        );

        JSONObject mdnsCfg = new JSONObject("{\"MDNS\":{\"Enabled\":true,\"Interval\":10}}");
        ipfs.setConfigKey("Discovery", mdnsCfg);

        assertTrue(
                "current IPFS config and backup should be equals",
                ipfs.getConfig().toString().equals(backup.toString())
        );

        // Reset config
        ipfs.setConfig(null);

        assertFalse(
                "current IPFS config and backup should not be equals (Identity changed)",
                ipfs.getConfig().toString().equals(backup.toString())
        );
    }

    public boolean getMDNSStateFromConfig() throws Exception {
        return ipfs.getConfig()
            .getJSONObject("Discovery")
            .getJSONObject("MDNS")
            .getBoolean("Enabled");
    }

    public void setMDNSStateToConfig(boolean state) throws Exception {
        JSONObject mdnsCfg = new JSONObject("{\"MDNS\":{\"Enabled\":" + state + "}}");
        ipfs.setConfigKey("Discovery", mdnsCfg);
    }
}
