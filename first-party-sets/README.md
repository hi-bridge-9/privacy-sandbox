# First-Party Sets

## はじめに
この機能は同じ組織・団体が運営する異なるドメイン郡を同じファーストパーティとして宣言可能にするもの。

例えば、Googleは国やサービス、ブランド、システムなどで「google.com、google.co.uk、youtube.com」といったドメインを使い分けているが、それらはブラウザ上で関連のないサードパーティとして扱われている。

しかし、この状況はCookieやキャッシュを利用したサイト移動時のサインインや埋め込みコンテンツの提供といったの機能の実現の妨げにつながり、ユーザーの信頼を損なう可能性がある。

そこでFirst-Party Setsは同じエンティティが管理するドメイン郡に限って、
一部の機能におけるサードパーティへの規制を緩和することを目標としている。

そのため、関連のないドメイン群でのサインインやCookieの共有、広告効果の計測の実現は非目標とされている。


## 利用条件（検討中）
- ファーストパーティとして宣言するドメイン群が、First-Party Setsが搭載されているブラウザの**UAポリシー**を満たしている
  - Chromium: https://github.com/michaelkleber/privacy-model/blob/master/README.md
  - Edge: https://blogs.windows.com/msedgedev/2019/06/27/tracking-prevention-microsoft-edge-preview/
  - Mozilla: https://wiki.mozilla.org/Security/Anti_tracking_policy
  - Webkit: https://webkit.org/tracking-prevention-policy/
- ドメイン群のプロトコルがhttpsである
- 上記を証明する情報やドメイン群のリストをパブリックトラッカーに提出し、承認される
  - この際に「オーナードメインは1つ」、「ドメインは2つ以上のセットに属していない」などの制約がある

## 実装項目（検討中）
- ドメイン群のオーナーに設定するドメインの`/.well-known/first-party-set`で、以下のようなJSONファイルを提供する
  ```
  例:
  オーナー: a.example
  メンバー: b.example、c.example

    {
        "owner": "https://a.example",
        "members": ["https://b.example", "https://c.example"]
    }
  ```

- ドメイン群のメンバーに設定する全てのドメインの`/.well-known/first-party-set`で、以下のようなJSONファイルを提供す
  ```
    {
        "owner": "https://a.example"
    }
  ```

## CookieのSameParty属性との組み合わせ
First-Party Setsの活用例としてCookieのSameParty属性との組み合わせが紹介されている。

これはChrome89でChromeに搭載されたもので、First-Party Setsで宣言したファーストパーティコンテキストの状況でCookieの読み書きを可能にする。

以下のようにCookieを設定する際にSameParty属性を含めることで、
サードパーティの埋め込みコンテンツなどやPOSTリクエストでの画面遷移の際にCookieを操作することができる。
```
例：
Set-Cookie: id=123; Secure; SameSite=Lax; SameParty
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

またCookieの設定時に以下の条件を満たす必要がある
- Secure属性が含まれている
- SameSite属性がNone、またはLaxに設定されている


## テスト方法

## 参考資料
### 概要
- **First-Party Sets**
https://developer.chrome.com/ja/docs/privacy-sandbox/first-party-sets/


### 詳細
- **First-Party Sets Proposal**
https://github.com/privacycg/first-party-sets

- **First-Party Sets**
https://www.chromium.org/updates/first-party-sets/

### オリジントライアル登録ページ
- **Trial for First Party Sets & SameParty**
https://developer.chrome.com/origintrials/#/view_trial/988540118207823873

### 問い合わせ先
- **Intent to Prototype**
https://groups.google.com/u/1/a/chromium.org/g/blink-dev/c/XkWbQKrBzMg



### デモサイト
- **First-Party Sets - Demo Site**
https://fps-member1.glitch.me/

### コード
特になし


## 関連資料
- **First-Party Sets and the SameParty attribute**
https://developer.chrome.com/blog/first-party-sets-sameparty/

- **SameParty Cookies and First-Party Sets**
https://github.com/privacycg/first-party-sets#sameparty-cookies-and-first-party-sets

- **SameParty cookie attribute explainer**
https://github.com/cfredric/sameparty



