# First-Party Sets

## はじめに
この機能は同じ組織・団体が運営する異なるドメイン郡を同じファーストパーティとして宣言可能にするもの。

例えば、Googleは国やサービス、ブランド、システムなどで「google.com、google.co.uk、youtube.com」といったドメインを使い分けているが、それらはブラウザ上で関連のないサードパーティとして扱われている。

しかし、この状況はCookieやキャッシュを利用したサイト移動時のサインインや埋め込みコンテンツの提供といったの機能の実現の妨げにつながり、ユーザーの信頼を損なう可能性がある。

そこでFirst-Party Setsは同じエンティティが管理するドメイン郡に限って、
一部の機能におけるサードパーティへの規制を緩和することを目標としている。

そのため、関連のないドメイン群でのサインインやCookieの共有、広告効果の計測の実現は非目標とされている。


参考：https://developer.chrome.com/ja/docs/privacy-sandbox/first-party-sets/#first-party-sets-%E3%81%AE%E4%BB%95%E7%B5%84%E3%81%BF

## 利用条件（未確定）
- ファーストパーティとして宣言するドメイン群が、First-Party Setsが搭載されているブラウザのUAポリシーを満たしている
  - Chromium: https://github.com/michaelkleber/privacy-model/blob/master/README.md
  - Edge: https://blogs.windows.com/msedgedev/2019/06/27/tracking-prevention-microsoft-edge-preview/
  - Mozilla: https://wiki.mozilla.org/Security/Anti_tracking_policy
  - Webkit: https://webkit.org/tracking-prevention-policy/
- ドメイン群のプロトコルがhttpsである
- 上記を証明する情報やドメイン群のリストをパブリックトラッカーに提出し、承認される
  - この際に「オーナードメインは1つ」、「ドメインは2つ以上のセットに属していない」などの制約がある

参考：https://github.com/privacycg/first-party-sets#acceptance-process

## 実装項目（未確定）
- ドメイン群のオーナーに設定するドメインの`/.well-known/first-party-set`で、以下のようなJSONファイルを提供する
  ```
  例:
  オーナー: a.example
  メンバー: b.example、c.example

    {
        "owner": "https://a.example",
        "members": ["https://b.example", "https://c.example"],
        ...
    }
  ```

- ドメイン群のメンバーに設定する全てのドメインの`/.well-known/first-party-set`で、以下のようなJSONファイルを提供する
  ```
    {
        "owner": "https://a.example"
    }
  ```

参考：https://developer.chrome.com/ja/docs/privacy-sandbox/first-party-sets/#first-party-sets-%E3%81%AE%E4%BB%95%E7%B5%84%E3%81%BF

## CookieのSameParty属性との組み合わせ
First-Party Setsの活用例としてCookieのSameParty属性との組み合わせが紹介されている。

これはFirst-Party Setsで宣言したファーストパーティコンテキストの状況でCookieの読み書きを可能にする。

以下のようにCookieを設定する際にSameParty属性を含めることで、
サードパーティの埋め込みコンテンツなどやPOSTリクエストでの画面遷移の際にCookieを操作することができる。
```
例：
Set-Cookie: id=123; Secure; SameSite=Lax; SameParty

＊この時、以下の条件を合わせて満たす必要がある
　- Secure属性が含まれている
　- SameSite属性がNone、またはLaxに設定されている
```

ただ注意するポイントとして、
事前に設定した自身のドメインに紐づくCookieにアクセス可能になるだけで、
異なるドメイン間でCookieを共有することは依然としてで不可能である。


```
OKパターン
1. a.exampleがCookieを設定
2. b.exampleがCookieを設定
3. a.exampleのサイト上に埋め込まれたb.exampleのリソースから、Cookie（b.exampleに紐づく）にアクセスする

NGパターン
1. a.exampleがCookieを設定
2. b.exampleがCookieを設定
3. a.exampleのサイト上に埋め込まれたb.exampleのリソースから、Cookie（a.exampleに紐づく）にアクセスする
```

参考：https://github.com/privacycg/first-party-sets#sameparty-cookies-and-first-party-sets

## 試験的利用手順
### 方法1：Chromeのフラグ機能を利用（GUI）
#### ブラウザから
1. Chromeの検索バーから`chrome://flags/#use-first-party-set`を開き、ステータスを「Disabled」 → 「Enabled」に変更する
2. 空欄に宣言したいドメイン群をオリジン形式で以下のようにカンマ区切りで入力する
    ```
    https://fps-owner.example,https://fps-member1.example,https://fps-member2.example
    ```
3. 手順1と同じように検索バーから`chrome://flags/#sameparty-cookies-considered-first-party`を開き、ステータスを「Disabled」 → 「Enabled」に変更する
4. 全てのウインドウを閉じて終了し、再度起動する
5. 以下のようにSameParty属性を含んだCookieを設定する
    ```
    Set-Cookie: id=123; Secure; SameSite=Lax; SameParty
    ```
6. サードパーティコンテキストからCookieを操作できるか確認する

#### ターミナルから
1. Chromeブラウザの全てのウインドウを閉じ、終了する
1. ターミナルにて以下のようなコマンドを実行する（ドメインをオリジン形式でカンマ区切りに並べる）
    ```sh
    MacOSの場合

    /Applications/Google\ Chrome.app/Contents/MacOS/Google\ Chrome  --use-first-party-set="https://fps-owner.example,https://fps-member1.example,https://fps-member2.example" --sameparty-cookies-considered-first-party
    ```

2. 以下のようにSameParty属性を含んだCookieを設定する
    ```
    Set-Cookie: id=123; Secure; SameSite=Lax; SameParty
    ```
3. サードパーティコンテキストからCookieを操作できるか確認する


参考：https://www.chromium.org/updates/first-party-sets/


### 方法2：Origin Trialsに登録し、トークンを使用する
1. Origin Trialsの[登録ページ](https://developer.chrome.com/origintrials/#/view_trial/988540118207823873)から、宣言したいドメイン毎にトークンを取得する（ドメインの数が3つなら、3回登録を行う）
2. ドメインを割り当てるサイトのHTMLやAPIのレスポンスヘッダーにトークンを設定する
    ```html
    HTMLの場合
    <meta http-equiv="origin-trial" content="[ここにトークンを入れる]">

    APIの場合
    Origin-Trial: [ここにトークンを入れる]
    ```
3. 以下のようにSameParty属性を含んだCookieを設定する
    ```
    Set-Cookie: id=123; Secure; SameSite=Lax; SameParty
    ```
4. サードパーティコンテキストからCookieを操作できるか確認する

参考：https://github.com/GoogleChrome/OriginTrials/blob/gh-pages/developer-guide.md#how-do-i-enable-an-experimental-feature-on-my-origin







## 参考資料
### 概要
- **First-Party Sets**
  - https://developer.chrome.com/ja/docs/privacy-sandbox/first-party-sets/

- **First-Party Sets and the SameParty attribute**
  - https://developer.chrome.com/blog/first-party-sets-sameparty/


### 詳細
- **First-Party Sets Proposal**
  - https://github.com/privacycg/first-party-sets

- **First-Party Sets**
  - https://www.chromium.org/updates/first-party-sets/


- **SameParty Cookies and First-Party Sets**
  - https://github.com/privacycg/first-party-sets#sameparty-cookies-and-first-party-sets

- **SameParty cookie attribute explainer**
  - https://github.com/cfredric/sameparty

- **First-Party Sets & SameParty Prototype Design Doc**
  - https://docs.google.com/document/d/16m5IfppdmmL-Zwk9zW8tJD4iHTVGJOLRP7g-QwBwX5c/edit



### オリジントライアル登録ページ
- **Trial for First Party Sets & SameParty**
  - https://developer.chrome.com/origintrials/#/view_trial/988540118207823873

### 問い合わせ先
- **Intent to Prototype**
  - https://groups.google.com/u/1/a/chromium.org/g/blink-dev/c/XkWbQKrBzMg


### デモサイト
- **First-Party Sets - Demo Site**
  - https://fps-member1.glitch.me/


### コード
特になし




