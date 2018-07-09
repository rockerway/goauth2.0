# goauth2.0
Auth 2.0 by Go lang


## Response 的規格

Status Code: 200 OK

Header 裡面一定要有這些：

Pragma: no-cache
Cache-Control: no-store (如果有 Token, Credential 或是其他敏感資料)

## Access Token

### 參數

參數名 | 必/選 | 填什麼/意義
--- | --- | ---
access_token | 必 | 由 Authorization Server 核發的 Access Token 。
token_type | 必 | Token 的類型，例如 Bearer （見系列文第 6 篇）。
expires_in | 建議有 | 幾秒過期，如 3600 表示 1 小時。若要省略，最好在文件裡註明。
scope | 必* | Access Token 的授權範圍 (scopes)。
refresh_token | 選 | 就是 Refresh Token

## Refresh Token

### 參數

參數名 | 必/選 | 填什麼/意義
--- | --- | ---
grant_type | 必 | refresh_token
refresh_token | 必 | 就填 Refresh Token
scope | 選 | 申請的存取範圍

## 發生錯誤時的回應方式
Status Code: 400 Bad Request (若無特別規定就用這個)

Header 和 Format (JSON) 跟核發的時候一樣。

### 參數

參數名 | 必/選 | 填什麼/意義
--- | --- | ---
error | 必 | 錯誤代碼，其值後述。
error_description | 選 | 人可讀的錯誤訊息，給 Client 開發者看的，不是給 End User 看的．ASCII 可見字元，除了雙引號和反斜線之外。
error_uri | 選 | 一個 URI ，指向載有錯誤細節的網頁，要符合 URI 的格式。

而 error 的值是以下的其中一個：

值 | 意義/用途
--- | ---
invalid_request | 欠缺必要的參數、有不正確的參數、有重複的參數、或其他原因導致無法解讀。
invalid_client | Client 認證失敗，如 Client 未知、沒送出 Client 認證、使用了 Server 不支援的認證方式。
invalid_grant | 提出的 Grant 或是 Refresh Token 不正確、過期、被撤銷、Redirection URI 不符、不是給你這個 Client。
unauthorized_client	Client | 沒有被授權可以使用這種方法來取得 Authorization Code。
unsupported_grant_type | Authorization Server 不支援使用這種 Grant Type (例如不支援 MAC)。
invalid_scope | 所要求的 scope 不正確、未知、無法解讀。
