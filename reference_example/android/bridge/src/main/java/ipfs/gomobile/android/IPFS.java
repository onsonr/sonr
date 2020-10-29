package ipfs.gomobile.android;

import android.content.Context;
import androidx.annotation.NonNull;

import java.io.File;
import java.util.Objects;
import org.apache.commons.io.FilenameUtils;
import org.json.JSONObject;

// Import gomobile-ipfs core
import core.Core;
import core.Config;
import core.Repo;
import core.Node;
import core.Shell;
import core.SockManager;

/**
* IPFS is a class that wraps a go-ipfs node and its shell over UDS.
*/
public class IPFS {

    // Paths
    private static final String defaultRepoPath = "/ipfs/repo";
    private final String absRepoPath;
    private final String absSockPath;

    // Go objects
    private static SockManager sockmanager;
    private Node node;
    private Repo repo;
    private Shell shell;

    /**
    * Class constructor using defaultRepoPath "/ipfs/repo" on internal storage.
    *
    * @param context The application context
    * @throws ConfigCreationException If the creation of the config failed
    * @throws RepoInitException If the initialization of the repo failed
    * @throws SockManagerException If the initialization of SockManager failed
    */
    public IPFS(@NonNull Context context)
        throws ConfigCreationException, RepoInitException, SockManagerException {
        this(context, defaultRepoPath, true);
    }

    /**
    * Class constructor using repoPath passed as parameter on internal storage.
    *
    * @param context The application context
    * @param repoPath The path of the go-ipfs repo (relative to internal root)
    * @throws ConfigCreationException If the creation of the config failed
    * @throws RepoInitException If the initialization of the repo failed
    * @throws SockManagerException If the initialization of SockManager failed
    */
    public IPFS(@NonNull Context context, @NonNull String repoPath)
        throws ConfigCreationException, RepoInitException, SockManagerException {
        this(context, repoPath, true);
    }

    /**
    * Class constructor using repoPath and storage location passed as parameters.
    *
    * @param context The application context
    * @param repoPath The path of the go-ipfs repo (relative to internal/external root)
    * @param internalStorage true, if the desired storage location for the repo path is internal
    *                        false, if the desired storage location for the repo path is external
    * @throws ConfigCreationException If the creation of the config failed
    * @throws RepoInitException If the initialization of the repo failed
    * @throws SockManagerException If the initialization of SockManager failed
    */
    public IPFS(@NonNull Context context, @NonNull String repoPath, boolean internalStorage)
        throws ConfigCreationException, RepoInitException, SockManagerException {
        Objects.requireNonNull(context, "context should not be null");
        Objects.requireNonNull(repoPath, "repoPath should not be null");

        String absPath;

        if (internalStorage) {
            absPath = context.getFilesDir().getAbsolutePath();
        } else {
            File externalDir = context.getExternalFilesDir(null);

            if (externalDir == null) {
                throw new RepoInitException("No external storage available");
            }
            absPath = externalDir.getAbsolutePath();
        }
        absRepoPath = FilenameUtils.normalizeNoEndSeparator(absPath + "/" + repoPath);

        synchronized (IPFS.class) {
            if (sockmanager == null) {
                try {
                    sockmanager = Core.newSockManager(context.getCacheDir().getAbsolutePath());
                } catch (Exception e) {
                    throw new SockManagerException("Socket manager initialization failed", e);
                }
            }
        }

        try {
            absSockPath = sockmanager.newSockPath();
        } catch (Exception e) {
            throw new SockManagerException("API socket creation failed", e);
        }

        if (!Core.repoIsInitialized(absRepoPath)) {
            Config config;

            try {
                config = Core.newDefaultConfig();
            } catch (Exception e) {
                throw new ConfigCreationException("Config creation failed", e);
            }

            final File repoDir = new File(absRepoPath);
            if (!repoDir.exists()) {
                if (!repoDir.mkdirs()) {
                    throw new RepoInitException("Repo directory creation failed: " + absRepoPath);
                }
            }
            try {
                Core.initRepo(absRepoPath, config);
            } catch (Exception e) {
                throw new RepoInitException("Repo initialization failed", e);
            }
        }
    }

    /**
    * Returns the repo absolute path as a string.
    *
    * @return The repo absolute path
    */
    synchronized public String getRepoAbsolutePath() {
        return absRepoPath;
    }

    /**
    * Returns true if this IPFS instance is "started" by checking if the underlying go-ipfs node
    * is instantiated.
    *
    * @return true, if this IPFS instance is started
    */
    synchronized public boolean isStarted() {
        return node != null;
    }

    /**
    * Starts this IPFS instance.
    *
    * @throws NodeStartException If the node is already started or if its startup fails
    */
    synchronized public void start() throws NodeStartException {
        if (isStarted()) {
            throw new NodeStartException("Node already started");
        }

        try {
            openRepoIfClosed();
            node = Core.newNode(repo);
            node.serveUnixSocketAPI(absSockPath);
        } catch (Exception e) {
            throw new NodeStartException("Node start failed", e);
        }

        shell = Core.newUDSShell(absSockPath);
    }

    /**
    * Stops this IPFS instance.
    *
    * @throws NodeStopException If the node is already stopped or if its stop fails
    */
    synchronized public void stop() throws NodeStopException {
        if (!isStarted()) {
            throw new NodeStopException("Node not started yet");
        }

        try {
            node.close();
            node = null;
            repo = null;
        } catch (Exception e) {
            throw new NodeStopException("Node stop failed", e);
        }
    }

    /**
    * Restarts this IPFS instance.
    *
    * @throws NodeStopException If the node is already stopped or if its stop fails
    */
    synchronized public void restart() throws NodeStopException {
        stop();
        try { start(); } catch(NodeStartException ignore) { /* Should never happen */ }
    }

    /**
    * Enable PubSub experimental feature on an IPFS node instance.
    * <b>A started instance must be restarted for this feature to be enabled.</b>
    *
    * @throws ExtraOptionException If enable the extra option failed
    * @see <a href="https://github.com/ipfs/go-ipfs/blob/master/docs/experimental-features.md#ipfs-pubsub">Experimental features of IPFS</a>
    */
    synchronized public void enablePubsubExperiment() throws ExtraOptionException {
        try {
            openRepoIfClosed();
            repo.enablePubsubExperiment();
        } catch (Exception e) {
            throw new ExtraOptionException("Enable pubsub experiment failed", e);
        }
    }

    /**
    * Enable PubSub experimental feature and IPNS record distribution through PubSub.
    * <b>A started instance must be restarted for this feature to be enabled.</b>
    *
    * @throws ExtraOptionException If enable the extra option failed
    * @see <a href="https://github.com/ipfs/go-ipfs/blob/master/docs/experimental-features.md#ipns-pubsub">Experimental features of IPFS</a>
    */
    synchronized public void enableNamesysPubsub() throws ExtraOptionException {
        try {
            openRepoIfClosed();
            repo.enableNamesysPubsub();
        } catch (Exception e) {
            throw new ExtraOptionException("Enable namesys pubsub failed", e);
        }
    }

    /**
    * Gets the IPFS instance config as a JSON.
    *
    * @return The IPFS instance config as a JSON
    * @throws ConfigGettingException If the getting of the config failed
    * @see <a href="https://github.com/ipfs/go-ipfs/blob/master/docs/config.md">IPFS Config Doc</a>
    */
    synchronized public JSONObject getConfig() throws ConfigGettingException {
        try {
            openRepoIfClosed();
            byte[] rawConfig = repo.getConfig().get();
            return new JSONObject(new String(rawConfig));
        } catch (Exception e) {
            throw new ConfigGettingException("Config getting failed", e);
        }
    }

    /**
    * Sets JSON config passed as parameter as IPFS config or reset to default config (with a new
    * identity) if the config parameter is null.<br>
    * <b>A started instance must be restarted for its config to be applied.</b>
    *
    * @param config The IPFS instance JSON config to set (if null, default config will be used)
    * @throws ConfigSettingException If the setting of the config failed
    * @see <a href="https://github.com/ipfs/go-ipfs/blob/master/docs/config.md">IPFS Config Doc</a>
    */
    synchronized public void setConfig(JSONObject config) throws ConfigSettingException {
        try {
            Config goConfig;

            if (config != null) {
                goConfig = Core.newConfig(config.toString().getBytes());
            } else {
                goConfig = Core.newDefaultConfig();
            }

            openRepoIfClosed();
            repo.setConfig(goConfig);
        } catch (Exception e) {
            throw new ConfigSettingException("Config setting failed", e);
        }
    }

    /**
    * Gets the JSON value associated to the key passed as parameter in the IPFS instance config.
    *
    * @param key The key associated to the value to get in the IPFS config
    * @return The JSON value associated to the key passed as parameter in the IPFS instance config
    * @throws ConfigGettingException If the getting of the config value failed
    * @see <a href="https://github.com/ipfs/go-ipfs/blob/master/docs/config.md">IPFS Config Doc</a>
    */
    synchronized public JSONObject getConfigKey(@NonNull String key) throws ConfigGettingException {
        Objects.requireNonNull(key, "key should not be null");

        try {
            openRepoIfClosed();
            byte[] rawValue = repo.getConfig().getKey(key);
            return new JSONObject(new String(rawValue));
        } catch (Exception e) {
            throw new ConfigGettingException("Config value getting failed", e);
        }
    }

    /**
    * Sets JSON config value to the key passed as parameters in the IPFS instance config.<br>
    * <b>A started instance must be restarted for its config to be applied.</b>
    *
    * @param key The key associated to the value to set in the IPFS instance config
    * @param value The JSON value associated to the key to set in the IPFS instance config
    * @throws ConfigSettingException If the setting of the config value failed
    * @see <a href="https://github.com/ipfs/go-ipfs/blob/master/docs/config.md">IPFS Config Doc</a>
    */
    synchronized public void setConfigKey(@NonNull String key, @NonNull JSONObject value)
        throws ConfigSettingException {
        Objects.requireNonNull(key, "key should not be null");
        Objects.requireNonNull(value, "value should not be null");

        try {
            openRepoIfClosed();
            Config ipfsConfig = repo.getConfig();
            ipfsConfig.setKey(key, value.toString().getBytes());
            repo.setConfig(ipfsConfig);
        } catch (Exception e) {
            throw new ConfigSettingException("Config setting failed", e);
        }
    }

    /**
    * Creates and returns a RequestBuilder associated to this IPFS instance shell.
    *
    * @param command The command of the request
    * @return A RequestBuilder based on the command passed as parameter
    * @throws ShellRequestException If this IPFS instance is not started
    * @see <a href="https://docs.ipfs.io/reference/api/http/">IPFS API Doc</a>
    */
    synchronized public RequestBuilder newRequest(@NonNull String command)
        throws ShellRequestException {
        Objects.requireNonNull(command, "command should not be null");

        if (!this.isStarted()) {
            throw new ShellRequestException("Shell request failed: node isn't started");
        }

        core.RequestBuilder requestBuilder = this.shell.newRequest(command);
        return new RequestBuilder(requestBuilder);
    }

    /**
    * Sets the primary and secondary DNS for gomobile (hacky, will be removed in future version)
    *
    * @param primary The primary DNS address in the form {@code <ip4>:<port>}
    * @param secondary The secondary DNS address in the form {@code <ip4>:<port>}
    */
    public static void setDNSPair(@NonNull String primary, @NonNull String secondary) {
        Objects.requireNonNull(primary, "primary should not be null");
        Objects.requireNonNull(secondary, "secondary should not be null");

        Core.setDNSPair(primary, secondary, false);
    }

    /**
    * Internal helper that opens the repo if it is closed.
    *
    * @throws RepoOpenException If the opening of the repo failed
    */
    synchronized private void openRepoIfClosed() throws RepoOpenException {
        if (repo == null) {
            try {
                repo = Core.openRepo(absRepoPath);
            } catch (Exception e) {
                throw new RepoOpenException("Repo opening failed", e);
            }
        }
    }

    // Exceptions
    public static class ExtraOptionException extends Exception {
        ExtraOptionException(String message, Throwable err) { super(message, err); }
    }

    public static class ConfigCreationException extends Exception {
        ConfigCreationException(String message, Throwable err) { super(message, err); }
    }

    public static class ConfigGettingException extends Exception {
        ConfigGettingException(String message, Throwable err) { super(message, err); }
    }

    public static class ConfigSettingException extends Exception {
        ConfigSettingException(String message, Throwable err) { super(message, err); }
    }

    public static class NodeStartException extends Exception {
        NodeStartException(String message) { super(message); }
        NodeStartException(String message, Throwable err) { super(message, err); }
    }

    public static class NodeStopException extends Exception {
        NodeStopException(String message) { super(message); }
        NodeStopException(String message, Throwable err) { super(message, err); }
    }

    public static class SockManagerException extends Exception {
        SockManagerException(String message, Throwable err) { super(message, err); }
    }

    public static class RepoInitException extends Exception {
        RepoInitException(String message) { super(message); }
        RepoInitException(String message, Throwable err) { super(message, err); }
    }

    public static class RepoOpenException extends Exception {
        RepoOpenException(String message, Throwable err) { super(message, err); }
    }

    public static class ShellRequestException extends Exception {
        ShellRequestException(String message) { super(message); }
    }
}
