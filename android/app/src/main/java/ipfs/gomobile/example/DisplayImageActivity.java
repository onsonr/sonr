package ipfs.gomobile.example;

import android.content.Intent;
import android.graphics.Bitmap;
import android.graphics.BitmapFactory;
import android.os.Bundle;
import android.util.Log;
import android.widget.ImageView;

import androidx.appcompat.app.AppCompatActivity;

public class DisplayImageActivity extends AppCompatActivity {
    private static final String TAG = "DisplayImageActivity";

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_display_image);

        Intent intent = getIntent();

        try {
            String title = intent.getExtras().getString("Title");
            getSupportActionBar().setTitle(title);
        } catch (NullPointerException err) {
            Log.e(TAG, "Error: can't set title");
        }

        try {
            byte[] data = intent.getExtras().getByteArray("ImageData");
            Bitmap bitmap = BitmapFactory.decodeByteArray(data, 0, data.length);

            ImageView imageView = findViewById(R.id.imageView);
            imageView.setImageBitmap(bitmap);
        } catch (Exception err) {
            Log.e(TAG, "Error: can't display image");
        }
    }
}
