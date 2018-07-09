# goauth2.0
Auth 2.0 by Go lang


## Response 的規格

Status Code: 200 OK

Header 裡面一定要有這些：

Pragma: no-cache
Cache-Control: no-store (如果有 Token, Credential 或是其他敏感資料)

### 參數

參數名 | 必/選 | 填什麼/意義
--- | --- | ---
access_token | 必 | 由 Authorization Server 核發的 Access Token 。
token_type | 必 | Token 的類型，例如 Bearer （見系列文第 6 篇）。
expires_in | 建議有 | 幾秒過期，如 3600 表示 1 小時。若要省略，最好在文件裡註明。
scope | 必* | Access Token 的授權範圍 (scopes)。
refresh_token | 選 | 就是 Refresh Token