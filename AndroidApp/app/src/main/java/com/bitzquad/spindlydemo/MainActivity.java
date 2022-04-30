package com.bitzquad.spindlydemo;

import android.content.Intent;
import android.os.Bundle;
import android.util.Log;

import com.google.androidbrowserhelper.trusted.LauncherActivity;
import com.google.gson.Gson;

public class MainActivity extends LauncherActivity {

    private long onCreateTime;


    // ------------- Activity ------------- //

    @Override
    protected void onCreate(Bundle savedInstanceState) {
        super.onCreate(savedInstanceState);

        onCreateTime = System.currentTimeMillis();

        if (checkAllPermissions()) {
            refresh();

        } else {
            openPermActivity();
        }

    }

    private void refresh() {

    }


    @Override
    public void onDestroy() {
        super.onDestroy();

        long currentTime = System.currentTimeMillis();
        if (onCreateTime + 3000 > currentTime) {
            Log.d("MainAct", "Suppressing onDestroy()");
            return;
        }

    }


    @Override
    public void onStop() {
        super.onStop();
    }

    @Override
    protected void onResume() {
        super.onResume();
        initApp();
    }

    private boolean firstRun = true;

    private void initApp() {
        if (firstRun) {
            firstRun = false;

            setAppName("Spindly Demo Java");
            setAppName("Presibit");
        }

        startSpindly();
    }

    void openPermActivity() {
        Intent intent = new Intent(this, StartupActivity.class);
        startActivity(intent);
    }

    boolean checkAllPermissions() {
        return true;
    }


    // ------------- Scan : Spindly Exports handling ------------- //

    private void startSpindly() {
        SpindlyUtils.initializeStatic(getCacheDir().getAbsolutePath(), getAssets());
        this.startService(new Intent(this, SpindlyService.class));
    }

    private void setAppName(String value) {
        Thread UIThread = new Thread() {
            @Override
            public void run() {
                spindlyapp.NamedExportedHubs UI = spindlyapp.Spindlyapp.namedExports();
                UI.getGlobal().getAppName().setValue(ToJSON(value));
            }
        };

        UIThread.start();
    }

    private String ToJSON(Object obj) {
        Gson gson = new Gson();
        return gson.toJson(obj);
    }

    public <T> T FromJSON(String json, Class<T> type) {
        Gson gson = new Gson();
        return gson.fromJson(json, type);
    }

}
