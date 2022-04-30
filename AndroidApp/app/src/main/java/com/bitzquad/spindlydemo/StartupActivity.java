package com.bitzquad.spindlydemo;

import android.animation.ObjectAnimator;
import android.annotation.SuppressLint;
import android.bluetooth.BluetoothAdapter;
import android.content.Intent;
import android.content.pm.PackageManager;
import android.os.Bundle;
import android.os.Handler;
import android.widget.ProgressBar;
import android.widget.TextView;
import android.widget.Toast;

import androidx.appcompat.app.AppCompatActivity;

public class StartupActivity extends AppCompatActivity {

    boolean hasTWAOpened = false;
    private final int permissionDeniedCount = 0;

    ProgressBar splashProgress;
    TextView txt;
    private static final int SPLASH_TIME = 300;

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);
        setContentView(R.layout.activity_startup);

        splashProgress = findViewById(R.id.splashProgress);
        txt = findViewById(R.id.txtinfo);
        if (requestAllPermissions()) {
            openTWA();
        }
        playProgress();

    }

    private void playProgress() {
        ObjectAnimator.ofInt(splashProgress, "progress", 0, 50, 70)
                .setDuration(2000)
                .start();
    }

    @SuppressLint("MissingPermission")
    void openTWA() {
        txt.setText("Setting Up...");
        splashProgress.setProgress(100);

        if (hasTWAOpened) {
            return;
        }
        hasTWAOpened = true;

        new Handler(getMainLooper()).postDelayed(new Runnable() {
            @Override
            public void run() {
                //Do any action here. Now we are moving to next page
                Intent intent = new Intent(StartupActivity.this, MainActivity.class);
                startActivity(intent);

                //This 'finish()' is for exiting the app when back button pressed from Home page which is ActivityHome
                finish();

            }
        }, SPLASH_TIME);

    }


    boolean requestAllPermissions() {

        if (permissionDeniedCount > 12) {
            onBackPressed();
            return false;
        }

        return true;

    }

}