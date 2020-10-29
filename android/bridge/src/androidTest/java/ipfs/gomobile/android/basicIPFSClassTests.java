package ipfs.gomobile.android;

import android.content.Context;

import androidx.test.platform.app.InstrumentationRegistry;
import androidx.test.ext.junit.runners.AndroidJUnit4;

import org.json.JSONObject;
import org.junit.Rule;
import org.junit.Test;
import org.junit.rules.Timeout;
import org.junit.runner.RunWith;

import java.io.File;

import static org.junit.Assert.*;

/**
* Instrumented test, which will execute on an Android device.
*
* @see <a href="http://d.android.com/tools/testing">Testing documentation</a>
*/
@RunWith(AndroidJUnit4.class)
public class basicIPFSClassTests {
    private final String internalDir = "/data/user/0/ipfs.gomobile.android.test/files";
    private final String externalDir = "/storage/emulated/0/Android/data/ipfs.gomobile.android.test/files";

    @Rule
    public Timeout globalTimeout = Timeout.seconds(600);

    @Test
    public void testDefaultPathInternal() throws Exception {
        Context appContext = InstrumentationRegistry.getInstrumentation().getTargetContext();
        String defaultRepoPath = "/ipfs/repo";

        IPFS ipfs = new IPFS(appContext);

        testIPFSInstance(ipfs, internalDir + defaultRepoPath);
    }

    @Test
    public void testCustomPathInternal() throws Exception {
        Context appContext = InstrumentationRegistry.getInstrumentation().getTargetContext();
        String customPath = "////////foo/";

        IPFS ipfs = new IPFS(appContext, customPath);

        testIPFSInstance(ipfs, internalDir + "/foo");
    }

    @Test
    public void testCustomPathExternal() throws Exception {
        Context appContext = InstrumentationRegistry.getInstrumentation().getTargetContext();
        String customPath = "////////foo/";

        IPFS ipfs = new IPFS(appContext, customPath, false);

        testIPFSInstance(ipfs, externalDir + "/foo");
    }

    @Test
    public void testNullParams() throws Exception {
        Context appContext = InstrumentationRegistry.getInstrumentation().getTargetContext();
        IPFS ipfs = new IPFS(appContext);
        ipfs.start();

        // Constructor tests
        try {
            new IPFS(null);
            fail("new IPFS() should fail with a null context");
        } catch (Exception e) { /* ignore */ }

        try {
            new IPFS(appContext, null);
            fail("new IPFS() should fail with a null repoPath");
        } catch (Exception e) { /* ignore */ }


        // DNS pair setter tests
        try {
            IPFS.setDNSPair(null, "foo");
            fail("setDNSPair() should fail with a null primary");
        } catch (Exception e) { /* ignore */ }

        try {
            IPFS.setDNSPair("foo", null);
            fail("setDNSPair() should fail with a null secondary");
        } catch (Exception e) { /* ignore */ }


        // Request tests
        try {
            ipfs.newRequest(null);
            fail("newRequest() should fail with a null command");
        } catch (Exception e) { /* ignore */ }

        try {
            ipfs.newRequest("foo").withArgument(null);
            fail("RequestBuilder.withArgument() should fail with a null argument");
        } catch (Exception e) { /* ignore */ }

        try {
            ipfs.newRequest("foo").withOption(null, true);
            fail("RequestBuilder.withOption() should fail with a null option");
        } catch (Exception e) { /* ignore */ }

        try {
            ipfs.newRequest("foo").withOption(null, "foo");
            fail("RequestBuilder.withOption() should fail with a null option");
        } catch (Exception e) { /* ignore */ }

        try {
            ipfs.newRequest("foo").withOption(null, "foo".getBytes());
            fail("RequestBuilder.withOption() should fail with a null option");
        } catch (Exception e) { /* ignore */ }

        try {
            String s = null;
            ipfs.newRequest("foo").withOption("foo", s);
            fail("RequestBuilder.withOption() should fail with a null value");
        } catch (Exception e) { /* ignore */ }

        try {
            byte[] b = null;
            ipfs.newRequest("foo").withOption("foo", b);
            fail("RequestBuilder.withOption() should fail with a null value");
        } catch (Exception e) { /* ignore */ }

        try {
            String s = null;
            ipfs.newRequest("foo").withBody(s);
            fail("RequestBuilder.withBody() should fail with a null body");
        } catch (Exception e) { /* ignore */ }

        try {
            byte[] b = null;
            ipfs.newRequest("foo").withBody(b);
            fail("RequestBuilder.withBody() should fail with a null body");
        } catch (Exception e) { /* ignore */ }

        try {
            ipfs.newRequest("foo").withHeader(null, "foo");
            fail("RequestBuilder.withHeader() should fail with a null key");
        } catch (Exception e) { /* ignore */ }

        try {
            ipfs.newRequest("foo").withHeader("foo", null);
            fail("RequestBuilder.withHeader() should fail with a null value");
        } catch (Exception e) { /* ignore */ }


        // Config tests
        try {
            ipfs.getConfigKey(null);
            fail("getConfigKey() should fail with a null key");
        } catch (Exception e) { /* ignore */ }

        try {
            ipfs.setConfigKey(null, new JSONObject("{\"foo\":\"bar\"}"));
            fail("setConfigKey() should fail with a null key");
        } catch (Exception e) { /* ignore */ }

        try {
            ipfs.setConfigKey("foo", null);
            fail("setConfigKey() should fail with a null value");
        } catch (Exception e) { /* ignore */ }
    }

    public void testIPFSInstance(IPFS ipfs, String expectedPath) throws Exception {
        // Tests on started IPFS instance
        ipfs.start();
        testRequest(ipfs);

        assertTrue(
            "IPFS should be started",
            ipfs.isStarted()
        );
        assertEquals(
            "Repo path mismatch",
            ipfs.getRepoAbsolutePath(),
            expectedPath
        );

        assertTrue(
            "config file doesn't exist in repo",
            new File(ipfs.getRepoAbsolutePath() + "/config").exists()
        );
        assertTrue(
            "version file doesn't exist in repo",
            new File(ipfs.getRepoAbsolutePath() + "/version").exists()
        );
        assertTrue(
            "repo.lock file doesn't exist in repo",
            new File(ipfs.getRepoAbsolutePath() + "/repo.lock").exists()
        );

        try {
            ipfs.start();
            fail("Calling start() on a started IPFS instance should throw");
        } catch (Exception e) { /* ignore */ }


        // Tests on stopped IPFS instance
        ipfs.stop();

        assertFalse(
            "IPFS should be stopped",
            ipfs.isStarted()
        );

        try {
            ipfs.stop();
            fail("Calling stop() on a stopped IPFS instance should throw");
        } catch (Exception e) { /* ignore */ }

        try {
            ipfs.restart();
            fail("Calling restart() on a stopped IPFS instance should throw");
        } catch (Exception e) { /* ignore */ }

        try {
            testRequest(ipfs);
            fail("Making request on a stopped IPFS instance should throw");
        } catch (Exception e) { /* ignore */ }


        // Tests on started IPFS instance (after stop)
        ipfs.start();
        testRequest(ipfs);

        assertTrue(
                "IPFS should be started",
                ipfs.isStarted()
        );


        // Tests on restarted IPFS instance
        ipfs.restart();
        testRequest(ipfs);

        assertTrue(
                "IPFS should be started",
                ipfs.isStarted()
        );
    }

    public void testRequest(IPFS ipfs) throws Exception {
        JSONObject response = ipfs.newRequest("id").sendToJSONList().get(0);
        // TODO: improve these checks
        assertEquals(
            "Invalid peerID",
            response.getString("ID").substring(0, 2),
            "Qm"
        );
        assertNotEquals(
            "Empty public key",
            response.getString("PublicKey"),
            ""
        );
    }
}
