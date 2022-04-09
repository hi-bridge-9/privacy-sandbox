# First-Party Sets＆SameParty属性の活用例（Cookie設定要求〜埋め込みコンテンツ返却）

## 通常の場合
#### 3rd Party Cookieが規制されている状態において、b.exampleはサードパーティコンテキストにあたるため、Cookieにアクセスすることができない。

 →⑤のリクエストヘッダーにCookieが含まれない
```mermaid
sequenceDiagram

    participant A as ブラウザ
    participant B as Webサイト（a.example）
    participant C as 埋め込みコンテンツ（b.example）

    Note over A, C: Cookie設定用タグ発火
    A->>C: ①HTTPリクエスト
    C-->>A: ②HTTPレスポンス（Cookie設定）
    Note left of C: Set-Cookie: id=123（Secure、SameSite=Lax）

    Note over A, C: サブリソースからCookieの操作
    A->>B: ③HTTPリクエスト
    B-->>A: ④HTTPレスポンス（HTMLファイル返却）
    A->>C: ⑤HTTPリクエスト（Cookieなし）
    C-->>A: ⑥HTTPレスポンス（コンテンツ返却）
```


## 機能適応後の場合
#### First-Party Setsの宣言とSameParty属性によって、b.exampleはファーストコンテキストという扱いになり、Cookieにアクセスすることができる。

→⑤のリクエストヘッダーにCookieが含まれる
```mermaid
%%{init:{'theme':'base'}}%%
sequenceDiagram

    participant A as ブラウザ
    participant B as Webサイト（a.example）
    participant C as 埋め込みコンテンツ（b.example）

    Note over A, C: Cookie設定用タグ発火
    A->>C: ①HTTPリクエスト
    C-->>A: ②HTTPレスポンス（Cookie設定）
    Note left of C: Set-Cookie: id=123（Secure、SameSite=Lax、 SameParty）

    Note over A, C: サブリソースからCookieの操作
    A->>B: ③HTTPリクエスト
    B-->>A: ④HTTPレスポンス（HTMLファイル返却）
    A->>C: ⑤HTTPリクエスト（Cookieあり）
    Note left of C: Cookie: id=123
    C-->>A: ⑥HTTPレスポンス（コンテンツ返却）
```