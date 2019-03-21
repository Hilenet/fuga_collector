# collector

## about
収集

## purpose
* DDD
* cleanArchitectureっぽいの

## memo
* repositoryパターン
* usecaseが無駄に厚く感じたので，interactor+DAO/DTOをサボって直実装+entity依存
  * DAO(InputPort)は欲しいかもしれん
  * まあ所詮外部packageに依存してるし
* repository, slack周りはもっと依存切れそう

### messageの対応
Message(ユーザによる個々のポスト)はSlackのObjectTypesに含まれず，Eventの形態として存在する．
[reference](https://api.slack.com/types)

Message {
  base-props: require,

  attatchments: json,

  subtype: some,
  subtype-entity: {some}
}

DomainModelの設計，DB上での取り扱い
* 利用形態が不明，怠いのでなるべく軽くしたい
* attachmentsとsubtypeがそれぞれ排他的option?
* hiddenは？
* star, pinn, reactionは放棄


## TODO
* Dockerはマルチステージングビルドで分離
* DM，Group(DMgroup?)の扱い

