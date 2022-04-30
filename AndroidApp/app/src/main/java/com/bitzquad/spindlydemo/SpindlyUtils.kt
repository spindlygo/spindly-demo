package com.bitzquad.spindlydemo

import android.content.res.AssetManager
import android.util.Log
import java.io.*

class SpindlyUtils {

    companion object {

        @JvmStatic
        fun initializeStatic(cacheDir: String, assetManager: AssetManager) {
            val cwd = "$cacheDir/"
            copyStaticFiles("WebApp", cwd, assetManager)
        }

        private fun copyStaticFiles(path: String, dest: String, assetManager: AssetManager) {

            val assets: Array<String>?
            try {
                assets = assetManager.list(path)
                if (assets!!.isEmpty()) {
                    copyFile(path, dest, assetManager)
                } else {
                    val fullPath = dest + path
                    val dir = File(fullPath)
                    if (!dir.exists()) dir.mkdir()
                    for (i in assets.indices) {
                        copyStaticFiles(path + "/" + assets[i], dest, assetManager)
                    }
                }
            } catch (ex: IOException) {
                Log.e("tag", "I/O Exception", ex)
            }
        }


        private fun copyFile(filename: String, dest: String, assetManager: AssetManager) {

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

}