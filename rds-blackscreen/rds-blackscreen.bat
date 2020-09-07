DEL /Q "%userprofile%\AppData\Local\Microsoft\Terminal Server Client\Cache\*.*"
REG ADD "HKEY_CURRENT_USER\Software\Microsoft\Terminal Server Client\Default" /v bitmapCacheSize /t REG_DWORD /d 0x00000000 /f
