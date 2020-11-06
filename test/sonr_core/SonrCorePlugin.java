package io.sonr.sonr_core;

import androidx.annotation.NonNull;

import io.flutter.embedding.engine.plugins.FlutterPlugin;
import io.flutter.plugin.common.MethodCall;
import io.flutter.plugin.common.MethodChannel;
import io.flutter.plugin.common.MethodChannel.MethodCallHandler;
import io.flutter.plugin.common.MethodChannel.Result;
import io.flutter.plugin.common.PluginRegistry.Registrar;

import sonr.*;

/** SonrCorePlugin */
public class SonrCorePlugin implements FlutterPlugin, MethodCallHandler {
  /// The MethodChannel that will the communication between Flutter and native Android
  ///
  /// This local reference serves to register the plugin with the Flutter Engine and unregister it
  /// when the Flutter Engine is detached from the Activity
  private MethodChannel channel;
  private SonrProxy sonr = new SonrProxy();

  @Override
  public void onAttachedToEngine(@NonNull FlutterPluginBinding flutterPluginBinding) {
    channel = new MethodChannel(flutterPluginBinding.getBinaryMessenger(), "sonr_core");
    channel.setMethodCallHandler(this);
  }

  @Override
  public void onMethodCall(@NonNull MethodCall call, @NonNull Result result) {
    // Switch by Call Method
    switch (call.method) {
      // ** Initialize Node ** //
      case "connect":
        if (call.arguments instanceof String) {
          Node temp = Sonr.start(call.arguments.toString(), sonr.callback);
          sonr.initialize(temp);
          result.success(sonr.getUser());
        } else {
          result.error("BAR_ARGS", "Wrong Argument types", null);
        }
        break;
      // ** Send a Message ** //
      case "send":
        if (call.arguments instanceof String) {
          // Send Message, Check if Fail
          if(sonr.send(call.arguments.toString()) == false) {
            result.error("FAIL", "Message couldnt be sent", null);
          }
        } else {
          result.error("BAR_ARGS", "Wrong Argument types", null);
        }
        break;
      // ** Send a Update ** //
      case "update":
        if (call.arguments instanceof String) {
          // Send Update, Check if Fail
          if(sonr.update(call.arguments.toString()) == false) {
            result.error("FAIL", "Update couldnt be sent", null);
          }
        } else {
          result.error("BAR_ARGS", "Wrong Argument types", null);
        }
        break;
      // ** Get Platform Version ** //
      case "getPlatformVersion":
        result.success("Android " + android.os.Build.VERSION.RELEASE);
        break;

      default:
        result.notImplemented();
    }
  }

  @Override
  public void onDetachedFromEngine(@NonNull FlutterPluginBinding binding) {
    channel.setMethodCallHandler(null);
  }
}
