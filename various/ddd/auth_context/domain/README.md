# domainパッケージ

## ■概略
企業全体のビジネスルールをカプセル化した業務ロジックを実装する層。

DDDにおける「Entities」とごっちゃになる恐れからか、「domain」という名にする事例が多い。

## ■[The Clean Architecture](https://blog.cleancoder.com/uncle-bob/2012/08/13/the-clean-architecture.html)対応

「Entities」に当たる。

> エンティティは、企業全体のビジネスルールをカプセル化します。
> エンティティは、メソッドを持つオブジェクトにすることも、データ構造と関数のセットにすることもできます。
> エンティティが企業内のさまざまなアプリケーションで使用される可能性がある限り、問題ではありません。

> あなたが企業を持っておらず、ただ一つのアプリケーションを書いているだけなら、これらのエンティティはアプリケーションのビジネスオブジェクトです。
> それらは最も一般的で高レベルのルールをカプセル化しています。
> 何かが外的に変化したときに変化する可能性が最も低いです。
> たとえば、これらのオブジェクトがページナビゲーションやセキュリティの変更による影響を受けるとは思われません。
> 特定のアプリケーションに対する操作上の変更は、エンティティ層に影響を与えません。

## ■[DDD](http://domainlanguage.com/wp-content/uploads/2016/05/DDD_Reference_2015-03.pdf)対応
