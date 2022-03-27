package com.bitzquad.spindlydemo

import android.content.Intent
import android.util.Log
import java.io.*


class MainActivity : com.google.androidbrowserhelper.trusted.LauncherActivity() {

    override fun onResume() {
        super.onResume()

        val cwd = cacheDir.absolutePath + "/"
        copyFileOrDir("WebApp", cwd)
        startSpindlyOnAnotherThread()

    }

    private fun startSpindlyOnAnotherThread() {

        Intent(this, SpindlyService::class.java).also { intent ->
            startService(intent)
        }

//
//        val goThread = Thread {
//            GoApp.mainMobile(cwd)
//        }
//        goThread.start()
    }

    private fun copyFileOrDir(path: String, dest: String) {

        val assetManager = this.assets
        val assets: Array<String>?
        try {
            assets = assetManager.list(path)
            if (assets!!.isEmpty()) {
                copyFile(path, dest)
            } else {
                val fullPath = dest + path
                val dir = File(fullPath)
                if (!dir.exists()) dir.mkdir()
                for (i in assets.indices) {
                    copyFileOrDir(path + "/" + assets[i], dest)
                }
            }
        } catch (ex: IOException) {
            Log.e("tag", "I/O Exception", ex)
        }
    }

    private fun copyFile(filename: String, dest: String) {


        val assetManager = this.assets
        val inputStrm: InputStream?
        val out: OutputStream?
        try {
            inputStrm = assetManager.open(filename)
            val newFileName = dest + filename

            Log.d("Copy", "Copy $filename to $newFileName")

            out = FileOutputStream(newFileName)
            val buffer = ByteArray(1024)
            var read: Int
            while (inputStrm.read(buffer).also { read = it } != -1) {
                out.write(buffer, 0, read)
            }
            inputStrm.close()
            out.flush()
            out.close()
        } catch (e: Exception) {
            Log.e("tag", e.message!!)
        }
    }
}

