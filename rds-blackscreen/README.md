Script to resolve endless black screen after successful rds login.

Whats the problem: Bitmap Cache gets stuck after login.

What it does:

- Deletes cache files under C:\Users\%AppData%\Local\Microsoft\Terminal Server Client\Cache

- Modify registry record to don't allow bitmap cache.

[HKEY_CURRENT_USER\Software\Microsoft\Terminal Server Client\Default]
"bitmapCacheSize"=dword:00000000

Actually working on batch only.
